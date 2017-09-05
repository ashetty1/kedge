// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	reflect "reflect"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterRole, InType: reflect.TypeOf(&ClusterRole{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterRoleBinding, InType: reflect.TypeOf(&ClusterRoleBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterRoleBindingList, InType: reflect.TypeOf(&ClusterRoleBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterRoleList, InType: reflect.TypeOf(&ClusterRoleList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_PolicyRule, InType: reflect.TypeOf(&PolicyRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Role, InType: reflect.TypeOf(&Role{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_RoleBinding, InType: reflect.TypeOf(&RoleBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_RoleBindingList, InType: reflect.TypeOf(&RoleBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_RoleList, InType: reflect.TypeOf(&RoleList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_RoleRef, InType: reflect.TypeOf(&RoleRef{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Subject, InType: reflect.TypeOf(&Subject{})},
	)
}

func DeepCopy_v1alpha1_ClusterRole(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRole)
		out := out.(*ClusterRole)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_PolicyRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_ClusterRoleBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRoleBinding)
		out := out.(*ClusterRoleBinding)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.Subjects != nil {
			in, out := &in.Subjects, &out.Subjects
			*out = make([]Subject, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1alpha1_ClusterRoleBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRoleBindingList)
		out := out.(*ClusterRoleBindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterRoleBinding, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ClusterRoleBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_ClusterRoleList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRoleList)
		out := out.(*ClusterRoleList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterRole, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ClusterRole(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_PolicyRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyRule)
		out := out.(*PolicyRule)
		*out = *in
		if in.Verbs != nil {
			in, out := &in.Verbs, &out.Verbs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.APIGroups != nil {
			in, out := &in.APIGroups, &out.APIGroups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Resources != nil {
			in, out := &in.Resources, &out.Resources
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.ResourceNames != nil {
			in, out := &in.ResourceNames, &out.ResourceNames
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.NonResourceURLs != nil {
			in, out := &in.NonResourceURLs, &out.NonResourceURLs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1alpha1_Role(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Role)
		out := out.(*Role)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_PolicyRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_RoleBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBinding)
		out := out.(*RoleBinding)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.Subjects != nil {
			in, out := &in.Subjects, &out.Subjects
			*out = make([]Subject, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1alpha1_RoleBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingList)
		out := out.(*RoleBindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]RoleBinding, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_RoleBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_RoleList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleList)
		out := out.(*RoleList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Role, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_Role(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_RoleRef(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleRef)
		out := out.(*RoleRef)
		*out = *in
		return nil
	}
}

func DeepCopy_v1alpha1_Subject(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Subject)
		out := out.(*Subject)
		*out = *in
		return nil
	}
}
