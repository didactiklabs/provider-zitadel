package applicationoidc

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_application_oidc", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "application.zitadel"
		r.References["project_id"] = config.Reference{
			TerraformName: "zitadel_project",
		}
		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}
