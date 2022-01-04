/*
Copyright 2021.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WeebleSpec defines the desired state of Weeble
type WeebleSpec struct {
	// Foo is an example field of Weeble. Edit weeble_types.go to remove/update
	Foo string `json:"foo,omitempty"`

	//+kubebuilder:validation:Minimum=0
	// Size is the size of the Weeble deployment
	Size int32 `json:"size"`
}

// WeebleStatus defines the observed state of Weeble
type WeebleStatus struct {
	// Nodes are the names of the Weeble pods
	Nodes []string `json:"nodes"`
}



// Weeble is the Schema for the weebles API
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Alias",type=string,JSONPath=`.spec.alias`
//+kubebuilder:printcolumn:name="Rank",type=integer,JSONPath=`.spec.rank`
//+kubebuilder:printcolumn:name="Bravely Run Away",type=boolean,JSONPath=`.spec.knights[?(@ == "Sir Robin")]`,description="when danger rears its ugly head, he bravely turned his tail and fled",priority=10
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type Weeble struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WeebleSpec   `json:"spec,omitempty"`
	Status WeebleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// WeebleList contains a list of Weeble
type WeebleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Weeble `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Weeble{}, &WeebleList{})
}
