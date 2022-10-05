/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DiskAccessObservation struct {

	// The ID of the Disk Access resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DiskAccessParameters struct {

	// The Azure Region where the Disk Access should exist. Changing this forces a new Disk to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the Resource Group where the Disk Access should exist. Changing this forces a new Disk Access to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Disk Access.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DiskAccessSpec defines the desired state of DiskAccess
type DiskAccessSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskAccessParameters `json:"forProvider"`
}

// DiskAccessStatus defines the observed state of DiskAccess.
type DiskAccessStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskAccessObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DiskAccess is the Schema for the DiskAccesss API. Manages a Disk Access.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DiskAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskAccessSpec   `json:"spec"`
	Status            DiskAccessStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskAccessList contains a list of DiskAccesss
type DiskAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskAccess `json:"items"`
}

// Repository type metadata.
var (
	DiskAccess_Kind             = "DiskAccess"
	DiskAccess_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskAccess_Kind}.String()
	DiskAccess_KindAPIVersion   = DiskAccess_Kind + "." + CRDGroupVersion.String()
	DiskAccess_GroupVersionKind = CRDGroupVersion.WithKind(DiskAccess_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskAccess{}, &DiskAccessList{})
}