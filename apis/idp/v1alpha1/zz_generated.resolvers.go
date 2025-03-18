// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/didactiklabs/provider-zitadel/apis/zitadel/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this OrgGithub.
func (mg *OrgGithub) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrgIDRef,
		Selector:     mg.Spec.ForProvider.OrgIDSelector,
		To: reference.To{
			List:    &v1alpha1.OrgList{},
			Managed: &v1alpha1.Org{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrgIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.OrgID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.OrgIDRef,
		Selector:     mg.Spec.InitProvider.OrgIDSelector,
		To: reference.To{
			List:    &v1alpha1.OrgList{},
			Managed: &v1alpha1.Org{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.OrgID")
	}
	mg.Spec.InitProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.OrgIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OrgGoogle.
func (mg *OrgGoogle) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrgIDRef,
		Selector:     mg.Spec.ForProvider.OrgIDSelector,
		To: reference.To{
			List:    &v1alpha1.OrgList{},
			Managed: &v1alpha1.Org{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrgIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.OrgID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.OrgIDRef,
		Selector:     mg.Spec.InitProvider.OrgIDSelector,
		To: reference.To{
			List:    &v1alpha1.OrgList{},
			Managed: &v1alpha1.Org{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.OrgID")
	}
	mg.Spec.InitProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.OrgIDRef = rsp.ResolvedReference

	return nil
}
