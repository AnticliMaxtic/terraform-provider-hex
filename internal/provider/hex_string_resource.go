// Copyright (c) 2025 AnticliMaxtic
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/hex"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &HexStringResource{}
var _ resource.ResourceWithImportState = &HexStringResource{}

func NewHexStringResource() resource.Resource {
	return &HexStringResource{}
}

// HexStringResource defines the resource implementation.
type HexStringResource struct{}

// HexStringResourceModel describes the resource data model.
type HexStringResourceModel struct {
	Data   types.String `tfsdk:"data"`
	Result types.String `tfsdk:"result"`
	Id     types.String `tfsdk:"id"`
}

func (r *HexStringResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_string"
}

func (r *HexStringResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Hex string resource - converts input data to hexadecimal representation",

		Attributes: map[string]schema.Attribute{
			"data": schema.StringAttribute{
				MarkdownDescription: "Input data to convert to hexadecimal",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"result": schema.StringAttribute{
				MarkdownDescription: "Hexadecimal representation of the input data",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Unique identifier for this resource",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *HexStringResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// This resource doesn't need any provider-level configuration
}

func (r *HexStringResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HexStringResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Convert the input data to hexadecimal
	inputData := data.Data.ValueString()
	hexResult := hex.EncodeToString([]byte(inputData))

	// Set the computed values
	data.Result = types.StringValue(hexResult)
	data.Id = types.StringValue(hexResult) // Use hex result as ID

	// Write logs using the tflog package
	tflog.Trace(ctx, "created hex string resource", map[string]interface{}{
		"input": inputData,
		"hex":   hexResult,
	})

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HexStringResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HexStringResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// For hex string resource, we don't need to read from external API
	// The values are computed from the input data and should remain unchanged
	// Save current state back into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HexStringResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data HexStringResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Convert the updated input data to hexadecimal
	inputData := data.Data.ValueString()
	hexResult := hex.EncodeToString([]byte(inputData))

	// Set the computed values
	data.Result = types.StringValue(hexResult)
	data.Id = types.StringValue(hexResult) // Use hex result as ID

	// Write logs using the tflog package
	tflog.Trace(ctx, "updated hex string resource", map[string]interface{}{
		"input": inputData,
		"hex":   hexResult,
	})

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HexStringResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HexStringResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Nothing to delete for hex string resource - it's just a computation
	// The resource will be removed from state automatically
}

func (r *HexStringResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
