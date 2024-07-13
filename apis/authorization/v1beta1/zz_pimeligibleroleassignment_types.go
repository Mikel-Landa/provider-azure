// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PimEligibleRoleAssignmentInitParameters struct {

	// The justification of the role assignment. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The justification of the eligible role assignment.
	Justification *string `json:"justification,omitempty" tf:"justification,omitempty"`

	// The principal id. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The principal id.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The role definition id. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The role definition id.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// A schedule block as defined below. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The schedule details of this eligible role assignment.
	Schedule *PimEligibleRoleAssignmentScheduleInitParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The scope. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The scope.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/management/v1beta1.ManagementGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Reference to a ManagementGroup in management to populate scope.
	// +kubebuilder:validation:Optional
	ScopeRef *v1.Reference `json:"scopeRef,omitempty" tf:"-"`

	// Selector for a ManagementGroup in management to populate scope.
	// +kubebuilder:validation:Optional
	ScopeSelector *v1.Selector `json:"scopeSelector,omitempty" tf:"-"`

	// A ticket block as defined below. Changing this forces a new Pim Eligible Role Assignment to be created.
	// Ticket details relating to the assignment.
	Ticket *PimEligibleRoleAssignmentTicketInitParameters `json:"ticket,omitempty" tf:"ticket,omitempty"`
}

type PimEligibleRoleAssignmentObservation struct {

	// The ID of the Pim Eligible Role Assignment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The justification of the role assignment. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The justification of the eligible role assignment.
	Justification *string `json:"justification,omitempty" tf:"justification,omitempty"`

	// The principal id. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The principal id.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The type of principal.
	// The type of principal.
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`

	// The role definition id. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The role definition id.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// A schedule block as defined below. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The schedule details of this eligible role assignment.
	Schedule *PimEligibleRoleAssignmentScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The scope. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The scope.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// A ticket block as defined below. Changing this forces a new Pim Eligible Role Assignment to be created.
	// Ticket details relating to the assignment.
	Ticket *PimEligibleRoleAssignmentTicketObservation `json:"ticket,omitempty" tf:"ticket,omitempty"`
}

type PimEligibleRoleAssignmentParameters struct {

	// The justification of the role assignment. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The justification of the eligible role assignment.
	// +kubebuilder:validation:Optional
	Justification *string `json:"justification,omitempty" tf:"justification,omitempty"`

	// The principal id. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The principal id.
	// +kubebuilder:validation:Optional
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The role definition id. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The role definition id.
	// +kubebuilder:validation:Optional
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// A schedule block as defined below. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The schedule details of this eligible role assignment.
	// +kubebuilder:validation:Optional
	Schedule *PimEligibleRoleAssignmentScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The scope. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The scope.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/management/v1beta1.ManagementGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Reference to a ManagementGroup in management to populate scope.
	// +kubebuilder:validation:Optional
	ScopeRef *v1.Reference `json:"scopeRef,omitempty" tf:"-"`

	// Selector for a ManagementGroup in management to populate scope.
	// +kubebuilder:validation:Optional
	ScopeSelector *v1.Selector `json:"scopeSelector,omitempty" tf:"-"`

	// A ticket block as defined below. Changing this forces a new Pim Eligible Role Assignment to be created.
	// Ticket details relating to the assignment.
	// +kubebuilder:validation:Optional
	Ticket *PimEligibleRoleAssignmentTicketParameters `json:"ticket,omitempty" tf:"ticket,omitempty"`
}

type PimEligibleRoleAssignmentScheduleInitParameters struct {

	// A expiration block as defined above.
	Expiration *ScheduleExpirationInitParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The start date time of the role assignment. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The start date time.
	StartDateTime *string `json:"startDateTime,omitempty" tf:"start_date_time,omitempty"`
}

type PimEligibleRoleAssignmentScheduleObservation struct {

	// A expiration block as defined above.
	Expiration *ScheduleExpirationObservation `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The start date time of the role assignment. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The start date time.
	StartDateTime *string `json:"startDateTime,omitempty" tf:"start_date_time,omitempty"`
}

type PimEligibleRoleAssignmentScheduleParameters struct {

	// A expiration block as defined above.
	// +kubebuilder:validation:Optional
	Expiration *ScheduleExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The start date time of the role assignment. Changing this forces a new Pim Eligible Role Assignment to be created.
	// The start date time.
	// +kubebuilder:validation:Optional
	StartDateTime *string `json:"startDateTime,omitempty" tf:"start_date_time,omitempty"`
}

type PimEligibleRoleAssignmentTicketInitParameters struct {

	// The ticket number.
	// The ticket number.
	Number *string `json:"number,omitempty" tf:"number,omitempty"`

	// The ticket system.
	// The ticket system.
	System *string `json:"system,omitempty" tf:"system,omitempty"`
}

type PimEligibleRoleAssignmentTicketObservation struct {

	// The ticket number.
	// The ticket number.
	Number *string `json:"number,omitempty" tf:"number,omitempty"`

	// The ticket system.
	// The ticket system.
	System *string `json:"system,omitempty" tf:"system,omitempty"`
}

type PimEligibleRoleAssignmentTicketParameters struct {

	// The ticket number.
	// The ticket number.
	// +kubebuilder:validation:Optional
	Number *string `json:"number,omitempty" tf:"number,omitempty"`

	// The ticket system.
	// The ticket system.
	// +kubebuilder:validation:Optional
	System *string `json:"system,omitempty" tf:"system,omitempty"`
}

type ScheduleExpirationInitParameters struct {

	// The duration of the role assignment in days. Conflicts with schedule[0].expiration[0].duration_hours,schedule[0].expiration[0].end_date_time Changing this forces a new Pim Eligible Role Assignment to be created.
	// The duration of the assignment in days.
	DurationDays *float64 `json:"durationDays,omitempty" tf:"duration_days,omitempty"`

	// The duration of the role assignment in hours. Conflicts with schedule[0].expiration[0].duration_days,schedule[0].expiration[0].end_date_time Changing this forces a new Pim Eligible Role Assignment to be created.
	// The duration of the assignment in hours.
	DurationHours *float64 `json:"durationHours,omitempty" tf:"duration_hours,omitempty"`

	// The end date time of the role assignment. Conflicts with schedule[0].expiration[0].duration_days,schedule[0].expiration[0].duration_hours Changing this forces a new Pim Eligible Role Assignment to be created.
	// The end date time of the assignment.
	EndDateTime *string `json:"endDateTime,omitempty" tf:"end_date_time,omitempty"`
}

type ScheduleExpirationObservation struct {

	// The duration of the role assignment in days. Conflicts with schedule[0].expiration[0].duration_hours,schedule[0].expiration[0].end_date_time Changing this forces a new Pim Eligible Role Assignment to be created.
	// The duration of the assignment in days.
	DurationDays *float64 `json:"durationDays,omitempty" tf:"duration_days,omitempty"`

	// The duration of the role assignment in hours. Conflicts with schedule[0].expiration[0].duration_days,schedule[0].expiration[0].end_date_time Changing this forces a new Pim Eligible Role Assignment to be created.
	// The duration of the assignment in hours.
	DurationHours *float64 `json:"durationHours,omitempty" tf:"duration_hours,omitempty"`

	// The end date time of the role assignment. Conflicts with schedule[0].expiration[0].duration_days,schedule[0].expiration[0].duration_hours Changing this forces a new Pim Eligible Role Assignment to be created.
	// The end date time of the assignment.
	EndDateTime *string `json:"endDateTime,omitempty" tf:"end_date_time,omitempty"`
}

type ScheduleExpirationParameters struct {

	// The duration of the role assignment in days. Conflicts with schedule[0].expiration[0].duration_hours,schedule[0].expiration[0].end_date_time Changing this forces a new Pim Eligible Role Assignment to be created.
	// The duration of the assignment in days.
	// +kubebuilder:validation:Optional
	DurationDays *float64 `json:"durationDays,omitempty" tf:"duration_days,omitempty"`

	// The duration of the role assignment in hours. Conflicts with schedule[0].expiration[0].duration_days,schedule[0].expiration[0].end_date_time Changing this forces a new Pim Eligible Role Assignment to be created.
	// The duration of the assignment in hours.
	// +kubebuilder:validation:Optional
	DurationHours *float64 `json:"durationHours,omitempty" tf:"duration_hours,omitempty"`

	// The end date time of the role assignment. Conflicts with schedule[0].expiration[0].duration_days,schedule[0].expiration[0].duration_hours Changing this forces a new Pim Eligible Role Assignment to be created.
	// The end date time of the assignment.
	// +kubebuilder:validation:Optional
	EndDateTime *string `json:"endDateTime,omitempty" tf:"end_date_time,omitempty"`
}

// PimEligibleRoleAssignmentSpec defines the desired state of PimEligibleRoleAssignment
type PimEligibleRoleAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PimEligibleRoleAssignmentParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider PimEligibleRoleAssignmentInitParameters `json:"initProvider,omitempty"`
}

// PimEligibleRoleAssignmentStatus defines the observed state of PimEligibleRoleAssignment.
type PimEligibleRoleAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PimEligibleRoleAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PimEligibleRoleAssignment is the Schema for the PimEligibleRoleAssignments API. Manages a Pim Eligible Role Assignment.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PimEligibleRoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principalId) || (has(self.initProvider) && has(self.initProvider.principalId))",message="spec.forProvider.principalId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleDefinitionId) || (has(self.initProvider) && has(self.initProvider.roleDefinitionId))",message="spec.forProvider.roleDefinitionId is a required parameter"
	Spec   PimEligibleRoleAssignmentSpec   `json:"spec"`
	Status PimEligibleRoleAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PimEligibleRoleAssignmentList contains a list of PimEligibleRoleAssignments
type PimEligibleRoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PimEligibleRoleAssignment `json:"items"`
}

// Repository type metadata.
var (
	PimEligibleRoleAssignment_Kind             = "PimEligibleRoleAssignment"
	PimEligibleRoleAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PimEligibleRoleAssignment_Kind}.String()
	PimEligibleRoleAssignment_KindAPIVersion   = PimEligibleRoleAssignment_Kind + "." + CRDGroupVersion.String()
	PimEligibleRoleAssignment_GroupVersionKind = CRDGroupVersion.WithKind(PimEligibleRoleAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&PimEligibleRoleAssignment{}, &PimEligibleRoleAssignmentList{})
}