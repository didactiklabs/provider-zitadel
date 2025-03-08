package triggeraction

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_trigger_actions", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "zitadel"
		r.Kind = "TriggerActions"
		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
		r.References["actions_ids"] = config.Reference{
			TerraformName: "zitadel_action",
		}
	})
}
