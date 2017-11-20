package openshift_test

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"
	"tests/e2e"
)

var ProjectPath = "$GOPATH/src/github.com/kedgeproject/kedge/"
var BinaryLocation = ProjectPath + "kedge"
var BinaryCommand = []string{"create", "-n"}

func runCmd(cmdS string) ([]byte, error) {
	var cmd *exec.Cmd
	var out, stdErr bytes.Buffer
	cmd = exec.Command("/bin/sh", "-c", cmdS)

	cmd.Stdout = &out
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error running command %v: %s", cmd, err)
	}
	return out.Bytes(), nil
}

func runBinary(files []string, namespace string) ([]byte, error) {
	args := append(BinaryCommand, namespace)
	for _, file := range files {
		args = append(args, "-f")
		args = append(args, os.ExpandEnv(file))
	}
	cmd := exec.Command(os.ExpandEnv(BinaryLocation), args...)

	var out, stdErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error running %q\n%s %s",
			fmt.Sprintf("command: %s", strings.Join(args, " ")),
			stdErr.String(), err)
	}
	return out.Bytes(), nil
}

func mapkeys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Check to see if specific pods have been started
func PodsStarted(t *testing.T, namespace string, podNames []string) error {
	// convert podNames to map
	podUp := make(map[string]int)
	for _, p := range podNames {
		podUp[p] = 0
	}

	// Timeouts after 9 minutes if the Pod has not yet started
	// 9 minute reasoning = 1 minute before 10-minute Golang test timeout.
	timeout := time.After(1 * time.Minute)
	tick := time.Tick(time.Second)

	for {
		select {
		case <-timeout:
			return fmt.Errorf("pods did not come up in given time: 5 minutes")
		case <-tick:
			t.Logf("pods not started yet: %q", strings.Join(mapkeys(podUp), " "))

			// iterate on all pods we care about
			for k := range podUp {
				podExists, err := runCmd("oc get pods --namespace=" + namespace + " --template=" +
					"'{{range .items}}{{ range .spec.containers }}" +
					"{{ eq .name \"" + k + "\"}}{{end}}{{end}}'")

				if err != nil {
					panic(err)
				}

				podStatus, err := runCmd("oc get pods --namespace=" + namespace + " --template=" +
					"'{{range .items}}{{ eq .status.phase \"Running\"}}{{end}}'")

				if err != nil {
					panic(err)
				}

				if bytes.Equal(podExists, podStatus) {
					t.Logf("Pod %q started!", k)
					delete(podUp, k)
				}
			}
		}

		if len(podUp) == 0 {
			break
		}
	}
	return nil
}

func getEndPoints(t *testing.T, namespace string, svcs []ServicePort) (map[string]string, error) {

	nodeIP, err := runCmd("oc get node --namespace=" + namespace + " --template=" +
		"'{{ range .items }}{{ ( index .status.addresses 0 ).address }}{{end}}'")
	if err != nil {
		return nil, errors.Wrap(err, "error while listing all nodes")
	}
	t.Logf("node ip address %s", nodeIP)

	// get all running services
	runningSvcs, err := runCmd("oc get svc --namespace=" + namespace + " --template=" +
		"'{{ range .items }}{{ .metadata.name }}{{end}}'")
	if err != nil {
		return nil, errors.Wrap(err, "error while listing all services")
	}

	endpoint := make(map[string]string)
	t.Logf("%s", runningSvcs)

	for _, svc := range svcs {
		if string(runningSvcs) == svc.Name {

			// Comparison with an integer port values value. Hence this workaround
			var portBug = strconv.Itoa(svc.Port) + ".00"

			getNodePort, err := runCmd("oc get svc --namespace=" + namespace + " --template=" +
				"'{{ range .items }}{{ range .spec.ports }}" +
				"{{ if eq .port " + portBug + "}}{{ .nodePort }}" +
				"{{end}}{{end}}{{end}}'")

			if err != nil {
				panic(err)
			}

			v := fmt.Sprintf("http://%s:%s", nodeIP, string(getNodePort))
			k := fmt.Sprintf("%s:%d", svc.Name, svc.Port)
			endpoint[k] = v
		}
	}
	t.Logf("endpoints: %#v", endpoint)
	return endpoint, nil
}

func pingEndPoints(t *testing.T, ep map[string]string) error {
	timeout := time.After(1 * time.Minute)
	tick := time.Tick(time.Second)

	for {
		select {
		case <-timeout:
			return fmt.Errorf("could not ping the specific service in given time: 5 minutes")
		case <-tick:
			for e, u := range ep {
				timeout := time.Duration(5 * time.Second)
				client := http.Client{
					Timeout: timeout,
				}
				respose, err := client.Get(u)
				if err != nil {
					t.Logf("error while making http request %q for service %q, err: %v", u, e, err)
					time.Sleep(1 * time.Second)
					continue
				}
				if respose.Status == "200 OK" {
					t.Logf("%q is running!", e)
					delete(ep, e)
				} else {
					return fmt.Errorf("for service %q got %q", e, respose.Status)
				}
			}
		}
		if len(ep) == 0 {
			break
		}
	}
	return nil
}

func deleteNS(t *testing.T, namespace string) {
	if _, err := runCmd("oc delete namespace " + namespace); err != nil {
		t.Logf("error deleting namespace %q: %v", namespace, err)
	}
	t.Logf("successfully deleted namespace: %q", namespace)
}

type ServicePort struct {
	Name string
	Port int
}

type testData struct {
	TestName         string
	Namespace        string
	InputFiles       []string
	PodStarted       []string
	NodePortServices []ServicePort
}

func TestOpenShift(t *testing.T) {

	tests := []testData{
		{
			TestName:  "Testing routes",
			Namespace: "testroutes",
			InputFiles: []string{
				ProjectPath + "docs/examples/routes/httpd.yml",
			},
			PodStarted: []string{"httpd"},
			NodePortServices: []ServicePort{
				{Name: "httpd", Port: 8080},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			t.Parallel()

			createNS, err := runCmd("oc create namespace " + test.Namespace)
			if err != nil {
				t.Fatalf("error creating namespace: %v", err)
			}
			t.Log(string(createNS))

			defer deleteNS(t, test.Namespace)

			convertedOutput, err := runBinary(test.InputFiles, test.Namespace)
			if err != nil {
				t.Fatalf("error running kedge: %v", err)
			}
			t.Log(string(convertedOutput))

			if err := PodsStarted(t, test.Namespace, test.PodStarted); err != nil {
				t.Fatalf("error finding running pods: %v", err)
			}

			// get endpoints for all services
			endPoints, err := getEndPoints(t, test.Namespace, test.NodePortServices)
			if err != nil {
				t.Fatalf("error getting nodes: %v", err)
			}

			if err := pingEndPoints(t, endPoints); err != nil {
				t.Fatalf("error pinging endpoint: %v", err)
			}
			t.Logf("Successfully pinged all endpoints!")

		})
	}
}
