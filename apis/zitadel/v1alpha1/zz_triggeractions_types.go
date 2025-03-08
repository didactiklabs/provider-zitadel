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

type TriggerActionsInitParameters struct {

	// (Set of String) IDs of the triggered actions
	// IDs of the triggered actions
	// +listType=set
	ActionIds []*string `json:"actionIds,omitempty" tf:"action_ids,omitempty"`

	// (String) Type of the flow to which the action triggers belong, supported values: FLOW_TYPE_EXTERNAL_AUTHENTICATION, FLOW_TYPE_CUSTOMISE_TOKEN, FLOW_TYPE_INTERNAL_AUTHENTICATION, FLOW_TYPE_SAML_RESPONSE
	// Type of the flow to which the action triggers belong, supported values: FLOW_TYPE_EXTERNAL_AUTHENTICATION, FLOW_TYPE_CUSTOMISE_TOKEN, FLOW_TYPE_INTERNAL_AUTHENTICATION, FLOW_TYPE_SAML_RESPONSE
	FlowType *string `json:"flowType,omitempty" tf:"flow_type,omitempty"`

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

	// (String) Trigger type on when the actions get triggered, supported values: TRIGGER_TYPE_POST_AUTHENTICATION, TRIGGER_TYPE_PRE_CREATION, TRIGGER_TYPE_POST_CREATION, TRIGGER_TYPE_PRE_USERINFO_CREATION, TRIGGER_TYPE_PRE_ACCESS_TOKEN_CREATION, TRIGGER_TYPE_PRE_SAML_RESPONSE_CREATION
	// Trigger type on when the actions get triggered, supported values: TRIGGER_TYPE_POST_AUTHENTICATION, TRIGGER_TYPE_PRE_CREATION, TRIGGER_TYPE_POST_CREATION, TRIGGER_TYPE_PRE_USERINFO_CREATION, TRIGGER_TYPE_PRE_ACCESS_TOKEN_CREATION, TRIGGER_TYPE_PRE_SAML_RESPONSE_CREATION
	TriggerType *string `json:"triggerType,omitempty" tf:"trigger_type,omitempty"`
}

type TriggerActionsObservation struct {

	// (Set of String) IDs of the triggered actions
	// IDs of the triggered actions
	// +listType=set
	ActionIds []*string `json:"actionIds,omitempty" tf:"action_ids,omitempty"`

	// (String) Type of the flow to which the action triggers belong, supported values: FLOW_TYPE_EXTERNAL_AUTHENTICATION, FLOW_TYPE_CUSTOMISE_TOKEN, FLOW_TYPE_INTERNAL_AUTHENTICATION, FLOW_TYPE_SAML_RESPONSE
	// Type of the flow to which the action triggers belong, supported values: FLOW_TYPE_EXTERNAL_AUTHENTICATION, FLOW_TYPE_CUSTOMISE_TOKEN, FLOW_TYPE_INTERNAL_AUTHENTICATION, FLOW_TYPE_SAML_RESPONSE
	FlowType *string `json:"flowType,omitempty" tf:"flow_type,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (String) Trigger type on when the actions get triggered, supported values: TRIGGER_TYPE_POST_AUTHENTICATION, TRIGGER_TYPE_PRE_CREATION, TRIGGER_TYPE_POST_CREATION, TRIGGER_TYPE_PRE_USERINFO_CREATION, TRIGGER_TYPE_PRE_ACCESS_TOKEN_CREATION, TRIGGER_TYPE_PRE_SAML_RESPONSE_CREATION
	// Trigger type on when the actions get triggered, supported values: TRIGGER_TYPE_POST_AUTHENTICATION, TRIGGER_TYPE_PRE_CREATION, TRIGGER_TYPE_POST_CREATION, TRIGGER_TYPE_PRE_USERINFO_CREATION, TRIGGER_TYPE_PRE_ACCESS_TOKEN_CREATION, TRIGGER_TYPE_PRE_SAML_RESPONSE_CREATION
	TriggerType *string `json:"triggerType,omitempty" tf:"trigger_type,omitempty"`
}

type TriggerActionsParameters struct {

	// (Set of String) IDs of the triggered actions
	// IDs of the triggered actions
	// +kubebuilder:validation:Optional
	// +listType=set
	ActionIds []*string `json:"actionIds,omitempty" tf:"action_ids,omitempty"`

	// (String) Type of the flow to which the action triggers belong, supported values: FLOW_TYPE_EXTERNAL_AUTHENTICATION, FLOW_TYPE_CUSTOMISE_TOKEN, FLOW_TYPE_INTERNAL_AUTHENTICATION, FLOW_TYPE_SAML_RESPONSE
	// Type of the flow to which the action triggers belong, supported values: FLOW_TYPE_EXTERNAL_AUTHENTICATION, FLOW_TYPE_CUSTOMISE_TOKEN, FLOW_TYPE_INTERNAL_AUTHENTICATION, FLOW_TYPE_SAML_RESPONSE
	// +kubebuilder:validation:Optional
	FlowType *string `json:"flowType,omitempty" tf:"flow_type,omitempty"`

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

	// (String) Trigger type on when the actions get triggered, supported values: TRIGGER_TYPE_POST_AUTHENTICATION, TRIGGER_TYPE_PRE_CREATION, TRIGGER_TYPE_POST_CREATION, TRIGGER_TYPE_PRE_USERINFO_CREATION, TRIGGER_TYPE_PRE_ACCESS_TOKEN_CREATION, TRIGGER_TYPE_PRE_SAML_RESPONSE_CREATION
	// Trigger type on when the actions get triggered, supported values: TRIGGER_TYPE_POST_AUTHENTICATION, TRIGGER_TYPE_PRE_CREATION, TRIGGER_TYPE_POST_CREATION, TRIGGER_TYPE_PRE_USERINFO_CREATION, TRIGGER_TYPE_PRE_ACCESS_TOKEN_CREATION, TRIGGER_TYPE_PRE_SAML_RESPONSE_CREATION
	// +kubebuilder:validation:Optional
	TriggerType *string `json:"triggerType,omitempty" tf:"trigger_type,omitempty"`
}

// TriggerActionsSpec defines the desired state of TriggerActions
type TriggerActionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerActionsParameters `json:"forProvider"`
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
	InitProvider TriggerActionsInitParameters `json:"initProvider,omitempty"`
}

// TriggerActionsStatus defines the observed state of TriggerActions.
type TriggerActionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerActionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TriggerActions is the Schema for the TriggerActionss API. Resource representing triggers, when actions get started
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type TriggerActions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.actionIds) || (has(self.initProvider) && has(self.initProvider.actionIds))",message="spec.forProvider.actionIds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.flowType) || (has(self.initProvider) && has(self.initProvider.flowType))",message="spec.forProvider.flowType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.triggerType) || (has(self.initProvider) && has(self.initProvider.triggerType))",message="spec.forProvider.triggerType is a required parameter"
	Spec   TriggerActionsSpec   `json:"spec"`
	Status TriggerActionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerActionsList contains a list of TriggerActionss
type TriggerActionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TriggerActions `json:"items"`
}

// Repository type metadata.
var (
	TriggerActions_Kind             = "TriggerActions"
	TriggerActions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TriggerActions_Kind}.String()
	TriggerActions_KindAPIVersion   = TriggerActions_Kind + "." + CRDGroupVersion.String()
	TriggerActions_GroupVersionKind = CRDGroupVersion.WithKind(TriggerActions_Kind)
)

func init() {
	SchemeBuilder.Register(&TriggerActions{}, &TriggerActionsList{})
}
