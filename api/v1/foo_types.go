package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FooSpec defines the desired state of Foo
type FooSpec struct {
	// Name of the friend Foo is looking for
	Name     string `json:"name"`
	Replicas int32  `json:"replicas"`
}

// FooStatus defines the observed state of Foo
type FooStatus struct {
	// Happy will be set to true if Foo found a friend
	Happy    bool   `json:"happy,omitempty"`
	Replicas int32  `json:"replicas"`
	Selector string `json:"selector"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas,selectorpath=.status.selector
// Foo is the Schema for the foos API
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FooSpec   `json:"spec,omitempty"`
	Status            FooStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FooList contains a list of Foo
type FooList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Foo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Foo{}, &FooList{})
}
