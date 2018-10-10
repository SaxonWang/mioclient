// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	system "github.com/hidevopsio/hiboot/pkg/system"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildConfig) DeepCopyInto(out *BuildConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildConfig.
func (in *BuildConfig) DeepCopy() *BuildConfig {
	if in == nil {
		return nil
	}
	out := new(BuildConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildConfigList) DeepCopyInto(out *BuildConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BuildConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildConfigList.
func (in *BuildConfigList) DeepCopy() *BuildConfigList {
	if in == nil {
		return nil
	}
	out := new(BuildConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildConfigSpec) DeepCopyInto(out *BuildConfigSpec) {
	*out = *in
	out.CloneConfig = in.CloneConfig
	if in.CompileCmd != nil {
		in, out := &in.CompileCmd, &out.CompileCmd
		*out = make([]CMD, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DockerFile != nil {
		in, out := &in.DockerFile, &out.DockerFile
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildConfigSpec.
func (in *BuildConfigSpec) DeepCopy() *BuildConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BuildConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildConfigs) DeepCopyInto(out *BuildConfigs) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]system.Env, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildConfigs.
func (in *BuildConfigs) DeepCopy() *BuildConfigs {
	if in == nil {
		return nil
	}
	out := new(BuildConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CMD) DeepCopyInto(out *CMD) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CMD.
func (in *CMD) DeepCopy() *CMD {
	if in == nil {
		return nil
	}
	out := new(CMD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneConfig) DeepCopyInto(out *CloneConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneConfig.
func (in *CloneConfig) DeepCopy() *CloneConfig {
	if in == nil {
		return nil
	}
	out := new(CloneConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfigs) DeepCopyInto(out *DeploymentConfigs) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]system.Env, len(*in))
		copy(*out, *in)
	}
	out.Labels = in.Labels
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfigs.
func (in *DeploymentConfigs) DeepCopy() *DeploymentConfigs {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayConfigs) DeepCopyInto(out *GatewayConfigs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayConfigs.
func (in *GatewayConfigs) DeepCopy() *GatewayConfigs {
	if in == nil {
		return nil
	}
	out := new(GatewayConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Labels) DeepCopyInto(out *Labels) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Labels.
func (in *Labels) DeepCopy() *Labels {
	if in == nil {
		return nil
	}
	out := new(Labels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pipeline.
func (in *Pipeline) DeepCopy() *Pipeline {
	if in == nil {
		return nil
	}
	out := new(Pipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineList) DeepCopyInto(out *PipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineList.
func (in *PipelineList) DeepCopy() *PipelineList {
	if in == nil {
		return nil
	}
	out := new(PipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpec) DeepCopyInto(out *PipelineSpec) {
	*out = *in
	out.Scm = in.Scm
	if in.Identifiers != nil {
		in, out := &in.Identifiers, &out.Identifiers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigFiles != nil {
		in, out := &in.ConfigFiles, &out.ConfigFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]Ports, len(*in))
		copy(*out, *in)
	}
	in.BuildConfigs.DeepCopyInto(&out.BuildConfigs)
	in.DeploymentConfigs.DeepCopyInto(&out.DeploymentConfigs)
	out.GatewayConfigs = in.GatewayConfigs
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpec.
func (in *PipelineSpec) DeepCopy() *PipelineSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ports) DeepCopyInto(out *Ports) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ports.
func (in *Ports) DeepCopy() *Ports {
	if in == nil {
		return nil
	}
	out := new(Ports)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scm) DeepCopyInto(out *Scm) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scm.
func (in *Scm) DeepCopy() *Scm {
	if in == nil {
		return nil
	}
	out := new(Scm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCode) DeepCopyInto(out *SourceCode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCode.
func (in *SourceCode) DeepCopy() *SourceCode {
	if in == nil {
		return nil
	}
	out := new(SourceCode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceConfig) DeepCopyInto(out *SourceConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceConfig.
func (in *SourceConfig) DeepCopy() *SourceConfig {
	if in == nil {
		return nil
	}
	out := new(SourceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceConfigList) DeepCopyInto(out *SourceConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SourceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceConfigList.
func (in *SourceConfigList) DeepCopy() *SourceConfigList {
	if in == nil {
		return nil
	}
	out := new(SourceConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceConfigScm) DeepCopyInto(out *SourceConfigScm) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceConfigScm.
func (in *SourceConfigScm) DeepCopy() *SourceConfigScm {
	if in == nil {
		return nil
	}
	out := new(SourceConfigScm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceConfigSpec) DeepCopyInto(out *SourceConfigSpec) {
	*out = *in
	out.Scm = in.Scm
	if in.SourceCode != nil {
		in, out := &in.SourceCode, &out.SourceCode
		*out = make([]SourceCode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceConfigSpec.
func (in *SourceConfigSpec) DeepCopy() *SourceConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SourceConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusDetails) DeepCopyInto(out *StatusDetails) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusDetails.
func (in *StatusDetails) DeepCopy() *StatusDetails {
	if in == nil {
		return nil
	}
	out := new(StatusDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TypeStatus) DeepCopyInto(out *TypeStatus) {
	*out = *in
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		if *in == nil {
			*out = nil
		} else {
			*out = new(StatusDetails)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TypeStatus.
func (in *TypeStatus) DeepCopy() *TypeStatus {
	if in == nil {
		return nil
	}
	out := new(TypeStatus)
	in.DeepCopyInto(out)
	return out
}
