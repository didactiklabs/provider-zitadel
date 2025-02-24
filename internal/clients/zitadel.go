/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/didactiklabs/provider-zitadel/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal zitadel credentials as JSON"

	// authentication
	keyInsecure       = "insecure"
	keyDomain         = "domain"
	keyJwtProfileFile = "jwt_profile_file"
	keyJwtProfileJson = "jwt_profile_json"
	keyPort           = "port"
	keyToken          = "token"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(
			ctx,
			pc.Spec.Credentials.Source,
			client,
			pc.Spec.Credentials.CommonCredentialSelectors,
		)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = map[string]any{}
		if v, ok := creds[keyInsecure]; ok {
			ps.Configuration[keyInsecure] = v
		}
		if v, ok := creds[keyDomain]; ok {
			ps.Configuration[keyDomain] = v
		}
		if v, ok := creds[keyJwtProfileFile]; ok {
			ps.Configuration[keyJwtProfileFile] = v
		}
		if v, ok := creds[keyJwtProfileJson]; ok {
			ps.Configuration[keyJwtProfileJson] = v
		}

		if v, ok := creds[keyPort]; ok {
			ps.Configuration[keyPort] = v
		}
		if v, ok := creds[keyToken]; ok {
			ps.Configuration[keyToken] = v
		}
		return ps, nil
	}
}
