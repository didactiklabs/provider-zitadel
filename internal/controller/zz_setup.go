// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	instancegithub "github.com/didactiklabs/provider-zitadel/internal/controller/idp/instancegithub"
	orggithub "github.com/didactiklabs/provider-zitadel/internal/controller/idp/orggithub"
	providerconfig "github.com/didactiklabs/provider-zitadel/internal/controller/providerconfig"
	org "github.com/didactiklabs/provider-zitadel/internal/controller/zitadel/org"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		instancegithub.Setup,
		orggithub.Setup,
		providerconfig.Setup,
		org.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
