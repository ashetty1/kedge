package openshift
import (

	e2e "github.com/kedgeproject/kedge/tests/e2e"
)

// Hardcoding the location of the binary, which is in root of project directory
var ProjectPath = "$GOPATH/src/github.com/kedgeproject/kedge/"
var BinaryLocation = ProjectPath + "kedge"
var BinaryCommand = []string{"create", "-n"}

const (
	jobTimeout = 10 * time.Minute
)


// These structs create a specific name as well as port to ping
type ServicePort struct {
	Name string
	Port int32
}

// Here we will test all of our test data!
type testData struct {
	TestName         string
	Namespace        string
	InputFiles       []string
	PodStarted       []string
	NodePortServices []ServicePort
}

// The "bread and butter" of the test-suite. We will iterate through
// each test that is required and make sure that not only are the pods started
// but that each test is pingable / is accessable.
func main(t *testing.T) {
	clientset, err := e2e.createClient()
	if err != nil {
		t.Fatalf("error getting kube client: %v", err)
	}

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

	_, err = clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Kubernetes cluster is not running or not accessible: %v", err)
	}

	for _, test := range tests {
		test := test // capture range variable
		t.Run(test.TestName, func(t *testing.T) {
			t.Parallel()
			// create a namespace
			_, err := createNS(clientset, test.Namespace)
			if err != nil {
				t.Fatalf("error creating namespace: %v", err)
			}
			t.Logf("namespace %q created", test.Namespace)
			defer deleteNamespace(t, clientset, test.Namespace)

			// run kedge
			convertedOutput, err := RunBinary(test.InputFiles, test.Namespace)
			if err != nil {
				t.Fatalf("error running kedge: %v", err)
			}
			t.Log(string(convertedOutput))

			// see if the pods are running
			if err := PodsStarted(t, clientset, test.Namespace, test.PodStarted); err != nil {
				t.Fatalf("error finding running pods: %v", err)
			}

			// get endpoints for all services
			endPoints, err := getEndPoints(t, clientset, test.Namespace, test.NodePortServices)
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
