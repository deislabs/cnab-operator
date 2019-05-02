package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BundleSpec defines the desired state of Bundle
// +k8s:openapi-gen=true
type BundleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// TODO - @radu-matei
	//
	// at this point, this is simply a placeholder.
	// We need to specifically agree on how the CRD should look like.

	// Name of the bundle
	Name string `json:"name,omitempty"`

	// URL of the bundle in a remote OCI registry
	URL string `json:"url,omitempty"`
}

// BundleStatus defines the observed state of Bundle
// +k8s:openapi-gen=true
type BundleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// TODO - @radu-matei
	//
	// Decide on how the status should look like.
	// Currently, this is a placeholder

	// Pulled indicates the controller successfully pulled the bundle from the registry
	Pulled bool `json:"pulled,omitempty"`

	// Processed indicates the controller successfully processed the pulled bundle
	Processed bool `json:"processed,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Bundle is the Schema for the bundles API
// +k8s:openapi-gen=true
type Bundle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BundleSpec   `json:"spec,omitempty"`
	Status BundleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BundleList contains a list of Bundle
type BundleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bundle `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bundle{}, &BundleList{})
}
