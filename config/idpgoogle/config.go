package idpgoogle

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_idp_google", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.Kind = "InstanceGoogle"
		r.ShortGroup = "idp.zitadel"
	})
}
