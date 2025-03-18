package orgidpgoogle

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_idp_google", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.Kind = "OrgGoogle"
		r.ShortGroup = "idp.zitadel"
		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}
