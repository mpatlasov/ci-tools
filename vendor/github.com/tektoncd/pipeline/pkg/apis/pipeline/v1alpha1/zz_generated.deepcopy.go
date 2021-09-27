// +build !ignore_autogenerated

/*
Copyright 2020 The Tekton Authors

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
	pod "github.com/tektoncd/pipeline/pkg/apis/pipeline/pod"
	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	resourcev1alpha1 "github.com/tektoncd/pipeline/pkg/apis/resource/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTask) DeepCopyInto(out *ClusterTask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTask.
func (in *ClusterTask) DeepCopy() *ClusterTask {
	if in == nil {
		return nil
	}
	out := new(ClusterTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterTask) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTaskList) DeepCopyInto(out *ClusterTaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTaskList.
func (in *ClusterTaskList) DeepCopy() *ClusterTaskList {
	if in == nil {
		return nil
	}
	out := new(ClusterTaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterTaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Condition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionCheck) DeepCopyInto(out *ConditionCheck) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionCheck.
func (in *ConditionCheck) DeepCopy() *ConditionCheck {
	if in == nil {
		return nil
	}
	out := new(ConditionCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionList) DeepCopyInto(out *ConditionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionList.
func (in *ConditionList) DeepCopy() *ConditionList {
	if in == nil {
		return nil
	}
	out := new(ConditionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConditionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionSpec) DeepCopyInto(out *ConditionSpec) {
	*out = *in
	in.Check.DeepCopyInto(&out.Check)
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]v1beta1.ParamSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]resourcev1alpha1.ResourceDeclaration, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionSpec.
func (in *ConditionSpec) DeepCopy() *ConditionSpec {
	if in == nil {
		return nil
	}
	out := new(ConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Inputs) DeepCopyInto(out *Inputs) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]v1beta1.TaskResource, len(*in))
		copy(*out, *in)
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]v1beta1.ParamSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Inputs.
func (in *Inputs) DeepCopy() *Inputs {
	if in == nil {
		return nil
	}
	out := new(Inputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Outputs) DeepCopyInto(out *Outputs) {
	*out = *in
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]TestResult, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]v1beta1.TaskResource, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Outputs.
func (in *Outputs) DeepCopy() *Outputs {
	if in == nil {
		return nil
	}
	out := new(Outputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(PipelineStatus)
		**out = **in
	}
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
	in.ListMeta.DeepCopyInto(&out.ListMeta)
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
func (in *PipelineRun) DeepCopyInto(out *PipelineRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineRun.
func (in *PipelineRun) DeepCopy() *PipelineRun {
	if in == nil {
		return nil
	}
	out := new(PipelineRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineRunList) DeepCopyInto(out *PipelineRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PipelineRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineRunList.
func (in *PipelineRunList) DeepCopy() *PipelineRunList {
	if in == nil {
		return nil
	}
	out := new(PipelineRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineRunSpec) DeepCopyInto(out *PipelineRunSpec) {
	*out = *in
	if in.PipelineRef != nil {
		in, out := &in.PipelineRef, &out.PipelineRef
		*out = new(v1beta1.PipelineRef)
		**out = **in
	}
	if in.PipelineSpec != nil {
		in, out := &in.PipelineSpec, &out.PipelineSpec
		*out = new(PipelineSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]v1beta1.PipelineResourceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]v1beta1.Param, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceAccountNames != nil {
		in, out := &in.ServiceAccountNames, &out.ServiceAccountNames
		*out = make([]v1beta1.PipelineRunSpecServiceAccountName, len(*in))
		copy(*out, *in)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.PodTemplate != nil {
		in, out := &in.PodTemplate, &out.PodTemplate
		*out = new(pod.Template)
		(*in).DeepCopyInto(*out)
	}
	if in.Workspaces != nil {
		in, out := &in.Workspaces, &out.Workspaces
		*out = make([]v1beta1.WorkspaceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TaskRunSpecs != nil {
		in, out := &in.TaskRunSpecs, &out.TaskRunSpecs
		*out = make([]PipelineTaskRunSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineRunSpec.
func (in *PipelineRunSpec) DeepCopy() *PipelineRunSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpec) DeepCopyInto(out *PipelineSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]v1beta1.PipelineDeclaredResource, len(*in))
		copy(*out, *in)
	}
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]PipelineTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]v1beta1.ParamSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Workspaces != nil {
		in, out := &in.Workspaces, &out.Workspaces
		*out = make([]v1beta1.PipelineWorkspaceDeclaration, len(*in))
		copy(*out, *in)
	}
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]v1beta1.PipelineResult, len(*in))
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
func (in *PipelineStatus) DeepCopyInto(out *PipelineStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineStatus.
func (in *PipelineStatus) DeepCopy() *PipelineStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineTask) DeepCopyInto(out *PipelineTask) {
	*out = *in
	if in.TaskRef != nil {
		in, out := &in.TaskRef, &out.TaskRef
		*out = new(v1beta1.TaskRef)
		**out = **in
	}
	if in.TaskSpec != nil {
		in, out := &in.TaskSpec, &out.TaskSpec
		*out = new(TaskSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1beta1.PipelineTaskCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RunAfter != nil {
		in, out := &in.RunAfter, &out.RunAfter
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1beta1.PipelineTaskResources)
		(*in).DeepCopyInto(*out)
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]v1beta1.Param, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Workspaces != nil {
		in, out := &in.Workspaces, &out.Workspaces
		*out = make([]v1beta1.WorkspacePipelineTaskBinding, len(*in))
		copy(*out, *in)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineTask.
func (in *PipelineTask) DeepCopy() *PipelineTask {
	if in == nil {
		return nil
	}
	out := new(PipelineTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PipelineTaskList) DeepCopyInto(out *PipelineTaskList) {
	{
		in := &in
		*out = make(PipelineTaskList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineTaskList.
func (in PipelineTaskList) DeepCopy() PipelineTaskList {
	if in == nil {
		return nil
	}
	out := new(PipelineTaskList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineTaskRunSpec) DeepCopyInto(out *PipelineTaskRunSpec) {
	*out = *in
	if in.TaskPodTemplate != nil {
		in, out := &in.TaskPodTemplate, &out.TaskPodTemplate
		*out = new(pod.Template)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineTaskRunSpec.
func (in *PipelineTaskRunSpec) DeepCopy() *PipelineTaskRunSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineTaskRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Run) DeepCopyInto(out *Run) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Run.
func (in *Run) DeepCopy() *Run {
	if in == nil {
		return nil
	}
	out := new(Run)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Run) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunList) DeepCopyInto(out *RunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Run, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunList.
func (in *RunList) DeepCopy() *RunList {
	if in == nil {
		return nil
	}
	out := new(RunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunSpec) DeepCopyInto(out *RunSpec) {
	*out = *in
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(v1beta1.TaskRef)
		**out = **in
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]v1beta1.Param, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunSpec.
func (in *RunSpec) DeepCopy() *RunSpec {
	if in == nil {
		return nil
	}
	out := new(RunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunStatus) DeepCopyInto(out *RunStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.RunStatusFields.DeepCopyInto(&out.RunStatusFields)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunStatus.
func (in *RunStatus) DeepCopy() *RunStatus {
	if in == nil {
		return nil
	}
	out := new(RunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunStatusFields) DeepCopyInto(out *RunStatusFields) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]v1beta1.TaskRunResult, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunStatusFields.
func (in *RunStatusFields) DeepCopy() *RunStatusFields {
	if in == nil {
		return nil
	}
	out := new(RunStatusFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Task) DeepCopyInto(out *Task) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Task.
func (in *Task) DeepCopy() *Task {
	if in == nil {
		return nil
	}
	out := new(Task)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Task) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskList) DeepCopyInto(out *TaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Task, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskList.
func (in *TaskList) DeepCopy() *TaskList {
	if in == nil {
		return nil
	}
	out := new(TaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskRun) DeepCopyInto(out *TaskRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskRun.
func (in *TaskRun) DeepCopy() *TaskRun {
	if in == nil {
		return nil
	}
	out := new(TaskRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskRunInputs) DeepCopyInto(out *TaskRunInputs) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]v1beta1.TaskResourceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]v1beta1.Param, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskRunInputs.
func (in *TaskRunInputs) DeepCopy() *TaskRunInputs {
	if in == nil {
		return nil
	}
	out := new(TaskRunInputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskRunList) DeepCopyInto(out *TaskRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TaskRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskRunList.
func (in *TaskRunList) DeepCopy() *TaskRunList {
	if in == nil {
		return nil
	}
	out := new(TaskRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskRunOutputs) DeepCopyInto(out *TaskRunOutputs) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]v1beta1.TaskResourceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskRunOutputs.
func (in *TaskRunOutputs) DeepCopy() *TaskRunOutputs {
	if in == nil {
		return nil
	}
	out := new(TaskRunOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskRunSpec) DeepCopyInto(out *TaskRunSpec) {
	*out = *in
	if in.TaskRef != nil {
		in, out := &in.TaskRef, &out.TaskRef
		*out = new(v1beta1.TaskRef)
		**out = **in
	}
	if in.TaskSpec != nil {
		in, out := &in.TaskSpec, &out.TaskSpec
		*out = new(TaskSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.PodTemplate != nil {
		in, out := &in.PodTemplate, &out.PodTemplate
		*out = new(pod.Template)
		(*in).DeepCopyInto(*out)
	}
	if in.Workspaces != nil {
		in, out := &in.Workspaces, &out.Workspaces
		*out = make([]v1beta1.WorkspaceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]v1beta1.Param, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1beta1.TaskRunResources)
		(*in).DeepCopyInto(*out)
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = new(TaskRunInputs)
		(*in).DeepCopyInto(*out)
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = new(TaskRunOutputs)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskRunSpec.
func (in *TaskRunSpec) DeepCopy() *TaskRunSpec {
	if in == nil {
		return nil
	}
	out := new(TaskRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskSpec) DeepCopyInto(out *TaskSpec) {
	*out = *in
	in.TaskSpec.DeepCopyInto(&out.TaskSpec)
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = new(Inputs)
		(*in).DeepCopyInto(*out)
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = new(Outputs)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSpec.
func (in *TaskSpec) DeepCopy() *TaskSpec {
	if in == nil {
		return nil
	}
	out := new(TaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestResult) DeepCopyInto(out *TestResult) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestResult.
func (in *TestResult) DeepCopy() *TestResult {
	if in == nil {
		return nil
	}
	out := new(TestResult)
	in.DeepCopyInto(out)
	return out
}
