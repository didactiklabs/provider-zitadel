/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	_ "embed" // go:embed

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/didactiklabs/provider-zitadel/config/action"
	"github.com/didactiklabs/provider-zitadel/config/applicationoidc"
	"github.com/didactiklabs/provider-zitadel/config/idpgithub"
	"github.com/didactiklabs/provider-zitadel/config/loginpolicy"
	"github.com/didactiklabs/provider-zitadel/config/org"
	"github.com/didactiklabs/provider-zitadel/config/orgidpgithub"
	"github.com/didactiklabs/provider-zitadel/config/project"
	"github.com/didactiklabs/provider-zitadel/config/projectgrant"
	"github.com/didactiklabs/provider-zitadel/config/projectrole"
	"github.com/didactiklabs/provider-zitadel/config/triggeraction"
)

const (
	resourcePrefix = "zitadel"
	modulePath     = "github.com/didactiklabs/provider-zitadel"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("didactiklabs.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		idpgithub.Configure,
		orgidpgithub.Configure,
		org.Configure,
		action.Configure,
		triggeraction.Configure,
		project.Configure,
		projectrole.Configure,
		projectgrant.Configure,
		applicationoidc.Configure,
		loginpolicy.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
