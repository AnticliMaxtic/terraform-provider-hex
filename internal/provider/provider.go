// Copyright (c) 2025 AnticliMaxtic
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure HexProvider satisfies various provider interfaces.
var _ provider.Provider = &HexProvider{}
var _ provider.ProviderWithFunctions = &HexProvider{}
var _ provider.ProviderWithEphemeralResources = &HexProvider{}

// HexProvider defines the provider implementation.
type HexProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// HexProviderModel describes the provider data model.
type HexProviderModel struct {
}

func (p *HexProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "hex"
	resp.Version = p.version
}

func (p *HexProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *HexProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data HexProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *HexProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewHexStringResource,
	}
}

func (p *HexProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		// No ephemeral resources in this provider
	}
}

func (p *HexProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// No data sources in this provider yet
	}
}

func (p *HexProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		// No functions in this provider yet
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &HexProvider{
			version: version,
		}
	}
}
