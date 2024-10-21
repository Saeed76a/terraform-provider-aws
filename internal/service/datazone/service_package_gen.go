// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package datazone

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datazone"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceEnvironmentBlueprint,
			Name:    "Environment Blueprint",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceAssetType,
			Name:    "Asset Type",
		},
		{
			Factory: newResourceDomain,
			Name:    "Domain",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory: newResourceEnvironment,
			Name:    "Environment",
		},
		{
			Factory: newResourceEnvironmentBlueprintConfiguration,
			Name:    "Environment Blueprint Configuration",
		},
		{
			Factory: newResourceEnvironmentProfile,
			Name:    "Environment Profile",
		},
		{
			Factory: newResourceFormType,
			Name:    "Form Type",
		},
		{
			Factory: newResourceGlossary,
			Name:    "Glossary",
		},
		{
			Factory: newResourceGlossaryTerm,
			Name:    "Glossary Term",
		},
		{
			Factory: newResourceProject,
			Name:    "Project",
		},
		{
			Factory: newResourceUserProfile,
			Name:    "User Profile",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DataZone
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*datazone.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return datazone.NewFromConfig(cfg,
		datazone.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
