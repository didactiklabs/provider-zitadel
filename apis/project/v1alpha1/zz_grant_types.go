// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GrantInitParameters struct {

	// (String) ID of the organization granted the project
	// ID of the organization granted the project
	GrantedOrgID *string `json:"grantedOrgId,omitempty" tf:"granted_org_id,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	// +crossplane:generate:reference:type=github.com/didactiklabs/provider-zitadel/apis/zitadel/v1alpha1.Org
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Org in zitadel to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDRef *v1.Reference `json:"orgIdRef,omitempty" tf:"-"`

	// Selector for a Org in zitadel to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDSelector *v1.Selector `json:"orgIdSelector,omitempty" tf:"-"`

	// (String) ID of the project
	// ID of the project
	// +crossplane:generate:reference:type=github.com/didactiklabs/provider-zitadel/apis/zitadel/v1alpha1.Project
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in zitadel to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in zitadel to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// (Set of String) List of roles granted
	// List of roles granted
	// +listType=set
	RoleKeys []*string `json:"roleKeys,omitempty" tf:"role_keys,omitempty"`
}

type GrantObservation struct {

	// (String) ID of the organization granted the project
	// ID of the organization granted the project
	GrantedOrgID *string `json:"grantedOrgId,omitempty" tf:"granted_org_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (String) ID of the project
	// ID of the project
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Set of String) List of roles granted
	// List of roles granted
	// +listType=set
	RoleKeys []*string `json:"roleKeys,omitempty" tf:"role_keys,omitempty"`
}

type GrantParameters struct {

	// (String) ID of the organization granted the project
	// ID of the organization granted the project
	// +kubebuilder:validation:Optional
	GrantedOrgID *string `json:"grantedOrgId,omitempty" tf:"granted_org_id,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	// +crossplane:generate:reference:type=github.com/didactiklabs/provider-zitadel/apis/zitadel/v1alpha1.Org
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Org in zitadel to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDRef *v1.Reference `json:"orgIdRef,omitempty" tf:"-"`

	// Selector for a Org in zitadel to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDSelector *v1.Selector `json:"orgIdSelector,omitempty" tf:"-"`

	// (String) ID of the project
	// ID of the project
	// +crossplane:generate:reference:type=github.com/didactiklabs/provider-zitadel/apis/zitadel/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in zitadel to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in zitadel to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// (Set of String) List of roles granted
	// List of roles granted
	// +kubebuilder:validation:Optional
	// +listType=set
	RoleKeys []*string `json:"roleKeys,omitempty" tf:"role_keys,omitempty"`
}

// GrantSpec defines the desired state of Grant
type GrantSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GrantParameters `json:"forProvider"`
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
	InitProvider GrantInitParameters `json:"initProvider,omitempty"`
}

// GrantStatus defines the observed state of Grant.
type GrantStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GrantObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Grant is the Schema for the Grants API. Resource representing the grant of a project to a different organization, also containing the available roles which can be given to the members of the projectgrant.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type Grant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.grantedOrgId) || (has(self.initProvider) && has(self.initProvider.grantedOrgId))",message="spec.forProvider.grantedOrgId is a required parameter"
	Spec   GrantSpec   `json:"spec"`
	Status GrantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GrantList contains a list of Grants
type GrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Grant `json:"items"`
}

// Repository type metadata.
var (
	Grant_Kind             = "Grant"
	Grant_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Grant_Kind}.String()
	Grant_KindAPIVersion   = Grant_Kind + "." + CRDGroupVersion.String()
	Grant_GroupVersionKind = CRDGroupVersion.WithKind(Grant_Kind)
)

func init() {
	SchemeBuilder.Register(&Grant{}, &GrantList{})
}
