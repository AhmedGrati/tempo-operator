//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuiltInCertManagement) DeepCopyInto(out *BuiltInCertManagement) {
	*out = *in
	out.CACertValidity = in.CACertValidity
	out.CACertRefresh = in.CACertRefresh
	out.CertValidity = in.CertValidity
	out.CertRefresh = in.CertRefresh
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuiltInCertManagement.
func (in *BuiltInCertManagement) DeepCopy() *BuiltInCertManagement {
	if in == nil {
		return nil
	}
	out := new(BuiltInCertManagement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureGates) DeepCopyInto(out *FeatureGates) {
	*out = *in
	out.OpenShift = in.OpenShift
	out.BuiltInCertManagement = in.BuiltInCertManagement
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureGates.
func (in *FeatureGates) DeepCopy() *FeatureGates {
	if in == nil {
		return nil
	}
	out := new(FeatureGates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagesSpec) DeepCopyInto(out *ImagesSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagesSpec.
func (in *ImagesSpec) DeepCopy() *ImagesSpec {
	if in == nil {
		return nil
	}
	out := new(ImagesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftFeatureGates) DeepCopyInto(out *OpenShiftFeatureGates) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftFeatureGates.
func (in *OpenShiftFeatureGates) DeepCopy() *OpenShiftFeatureGates {
	if in == nil {
		return nil
	}
	out := new(OpenShiftFeatureGates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectConfig) DeepCopyInto(out *ProjectConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
	out.DefaultImages = in.DefaultImages
	out.Gates = in.Gates
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectConfig.
func (in *ProjectConfig) DeepCopy() *ProjectConfig {
	if in == nil {
		return nil
	}
	out := new(ProjectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
