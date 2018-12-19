package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LoadTestSpec defines the desired state of LoadTest
type LoadTestSpec struct {
	URL      string `json:"url"`
	Duration string `json:"duration"`
	Header	 string `json:"header"`
	User	 string `json:"user"`
	Password string `json:"password"`
	Qps	 string `json:"qps"`
	Threads  string `json:"threads"`
	Action   string `json:"action"`
	
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// LoadTestStatus defines the observed state of LoadTest
type LoadTestStatus struct {
	Condition []LoadTestCondition	`json:"condition"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

type LoadTestCondition struct {
	Target50	string	`json:"50%"`
	Target75	string	`json:"75%"`
	Target90	string	`json:"90%"`
	Target99	string	`json:"99%"`
	Target999	string	`json:"99.9%"`
}
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LoadTest is the Schema for the loadtests API
// +k8s:openapi-gen=true
type LoadTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoadTestSpec   `json:"spec,omitempty"`
	Status LoadTestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LoadTestList contains a list of LoadTest
type LoadTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadTest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoadTest{}, &LoadTestList{})
}