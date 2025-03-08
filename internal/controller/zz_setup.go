// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	oidc "github.com/didactiklabs/provider-zitadel/internal/controller/application/oidc"
	instancegithub "github.com/didactiklabs/provider-zitadel/internal/controller/idp/instancegithub"
	orggithub "github.com/didactiklabs/provider-zitadel/internal/controller/idp/orggithub"
	policy "github.com/didactiklabs/provider-zitadel/internal/controller/login/policy"
	grant "github.com/didactiklabs/provider-zitadel/internal/controller/project/grant"
	role "github.com/didactiklabs/provider-zitadel/internal/controller/project/role"
	providerconfig "github.com/didactiklabs/provider-zitadel/internal/controller/providerconfig"
	action "github.com/didactiklabs/provider-zitadel/internal/controller/zitadel/action"
	actions "github.com/didactiklabs/provider-zitadel/internal/controller/zitadel/actions"
	org "github.com/didactiklabs/provider-zitadel/internal/controller/zitadel/org"
	project "github.com/didactiklabs/provider-zitadel/internal/controller/zitadel/project"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		oidc.Setup,
		instancegithub.Setup,
		orggithub.Setup,
		policy.Setup,
		grant.Setup,
		role.Setup,
		providerconfig.Setup,
		action.Setup,
		actions.Setup,
		org.Setup,
		project.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
