// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package testing

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: DeepCopy_testing_AdmissionPluginTestConfig, InType: reflect.TypeOf(&AdmissionPluginTestConfig{})},
	}
}

// DeepCopy_testing_AdmissionPluginTestConfig is an autogenerated deepcopy function.
func DeepCopy_testing_AdmissionPluginTestConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AdmissionPluginTestConfig)
		out := out.(*AdmissionPluginTestConfig)
		*out = *in
		return nil
	}
}
