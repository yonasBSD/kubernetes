//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1

import (
	resourcev1 "k8s.io/api/resource/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&resourcev1.ResourceClaim{}, func(obj interface{}) { SetObjectDefaults_ResourceClaim(obj.(*resourcev1.ResourceClaim)) })
	scheme.AddTypeDefaultingFunc(&resourcev1.ResourceClaimList{}, func(obj interface{}) { SetObjectDefaults_ResourceClaimList(obj.(*resourcev1.ResourceClaimList)) })
	scheme.AddTypeDefaultingFunc(&resourcev1.ResourceClaimTemplate{}, func(obj interface{}) {
		SetObjectDefaults_ResourceClaimTemplate(obj.(*resourcev1.ResourceClaimTemplate))
	})
	scheme.AddTypeDefaultingFunc(&resourcev1.ResourceClaimTemplateList{}, func(obj interface{}) {
		SetObjectDefaults_ResourceClaimTemplateList(obj.(*resourcev1.ResourceClaimTemplateList))
	})
	scheme.AddTypeDefaultingFunc(&resourcev1.ResourceSlice{}, func(obj interface{}) { SetObjectDefaults_ResourceSlice(obj.(*resourcev1.ResourceSlice)) })
	scheme.AddTypeDefaultingFunc(&resourcev1.ResourceSliceList{}, func(obj interface{}) { SetObjectDefaults_ResourceSliceList(obj.(*resourcev1.ResourceSliceList)) })
	return nil
}

func SetObjectDefaults_ResourceClaim(in *resourcev1.ResourceClaim) {
	for i := range in.Spec.Devices.Requests {
		a := &in.Spec.Devices.Requests[i]
		if a.Exactly != nil {
			SetDefaults_ExactDeviceRequest(a.Exactly)
			for j := range a.Exactly.Tolerations {
				b := &a.Exactly.Tolerations[j]
				if b.Operator == "" {
					b.Operator = "Equal"
				}
			}
		}
		for j := range a.FirstAvailable {
			b := &a.FirstAvailable[j]
			SetDefaults_DeviceSubRequest(b)
			for k := range b.Tolerations {
				c := &b.Tolerations[k]
				if c.Operator == "" {
					c.Operator = "Equal"
				}
			}
		}
	}
	if in.Status.Allocation != nil {
		for i := range in.Status.Allocation.Devices.Results {
			a := &in.Status.Allocation.Devices.Results[i]
			for j := range a.Tolerations {
				b := &a.Tolerations[j]
				if b.Operator == "" {
					b.Operator = "Equal"
				}
			}
		}
	}
}

func SetObjectDefaults_ResourceClaimList(in *resourcev1.ResourceClaimList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ResourceClaim(a)
	}
}

func SetObjectDefaults_ResourceClaimTemplate(in *resourcev1.ResourceClaimTemplate) {
	for i := range in.Spec.Spec.Devices.Requests {
		a := &in.Spec.Spec.Devices.Requests[i]
		if a.Exactly != nil {
			SetDefaults_ExactDeviceRequest(a.Exactly)
			for j := range a.Exactly.Tolerations {
				b := &a.Exactly.Tolerations[j]
				if b.Operator == "" {
					b.Operator = "Equal"
				}
			}
		}
		for j := range a.FirstAvailable {
			b := &a.FirstAvailable[j]
			SetDefaults_DeviceSubRequest(b)
			for k := range b.Tolerations {
				c := &b.Tolerations[k]
				if c.Operator == "" {
					c.Operator = "Equal"
				}
			}
		}
	}
}

func SetObjectDefaults_ResourceClaimTemplateList(in *resourcev1.ResourceClaimTemplateList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ResourceClaimTemplate(a)
	}
}

func SetObjectDefaults_ResourceSlice(in *resourcev1.ResourceSlice) {
	for i := range in.Spec.Devices {
		a := &in.Spec.Devices[i]
		for j := range a.Taints {
			b := &a.Taints[j]
			SetDefaults_DeviceTaint(b)
		}
	}
}

func SetObjectDefaults_ResourceSliceList(in *resourcev1.ResourceSliceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ResourceSlice(a)
	}
}
