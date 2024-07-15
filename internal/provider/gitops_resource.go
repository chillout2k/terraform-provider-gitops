package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/chillout2k/gitopsclient"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &gitopsInstanceResource{}
	_ resource.ResourceWithConfigure   = &gitopsInstanceResource{}
	_ resource.ResourceWithImportState = &gitopsInstanceResource{}
)

// NewOrderResource is a helper function to simplify the provider implementation.
func NewGitopsInstanceResource() resource.Resource {
	return &gitopsInstanceResource{}
}

// gitopsInstanceResource is the resource implementation.
type gitopsInstanceResource struct {
	client *gitopsclient.GitopsClient
}

// gitopsInstanceResourceModel maps the resource schema data.
type gitopsInstanceResourceModel struct {
	Instance_id   types.String `tfsdk:"instance_id"`
	Order_time    types.String `tfsdk:"order_time"`
	Stage         types.String `tfsdk:"stage"`
	Instance_name types.String `tfsdk:"instance_name"`
	Orderer_id    types.String `tfsdk:"orderer_id"`
	Bits_account  types.Int64  `tfsdk:"bits_account"`
	Service_id    types.Int64  `tfsdk:"service_id"`
	Replica_count types.Int64  `tfsdk:"replica_count"`
	Version       types.String `tfsdk:"version"`
	Some_value    types.String `tfsdk:"some_value"`
	LastUpdated   types.String `tfsdk:"last_updated"`
}

// Configure adds the provider configured client to the resource.
func (r *gitopsInstanceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*gitopsclient.GitopsClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *GitopsClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Metadata returns the resource type name.
func (r *gitopsInstanceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_instance"
}

// Schema defines the schema for the resource.
func (r *gitopsInstanceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"instance_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"order_time": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"stage": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: false,
				Required: true,
			},
			"orderer_id": schema.StringAttribute{
				Computed: false,
				Required: true,
			},
			"bits_account": schema.Int64Attribute{
				Computed: false,
				Required: true,
			},
			"service_id": schema.Int64Attribute{
				Computed: false,
				Required: true,
			},
			"replica_count": schema.Int64Attribute{
				Computed: false,
				Required: true,
			},
			"version": schema.StringAttribute{
				Computed: false,
				Required: true,
			},
			"some_value": schema.StringAttribute{
				Computed: false,
				Required: true,
			},
		},
	}
}

// Create a new resource.
func (r *gitopsInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan gitopsInstanceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var instance_order gitopsclient.InstanceOrder
	instance_order.Instance_name = plan.Instance_name.ValueString()
	instance_order.Orderer_id = plan.Orderer_id.ValueString()
	instance_order.Bits_account = uint64(plan.Bits_account.ValueInt64())
	instance_order.Service_id = uint64(plan.Service_id.ValueInt64())
	instance_order.Replica_count = uint64(plan.Replica_count.ValueInt64())
	instance_order.Version = plan.Version.ValueString()
	instance_order.Some_value = plan.Some_value.ValueString()

	// Create new gitopsInstance
	gitopsInstance, err := r.client.PostInstanceOrder(instance_order)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating instance gitopsInstance",
			"Could not create gitopsInstance, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.Instance_id = types.StringValue(gitopsInstance.Instance_id)
	plan.Order_time = types.StringValue(gitopsInstance.Order_time)
	plan.Stage = types.StringValue(gitopsInstance.Stage)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *gitopsInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state gitopsInstanceResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from HashiCups
	gitopsInstance, err := r.client.GetInstance(state.Instance_id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading gitopsInstance",
			"Could not read gitopsInstance ID "+state.Instance_id.ValueString()+": "+err.Error(),
		)
		return
	}

	state.Instance_id = types.StringValue(gitopsInstance.Instance_id)
	state.Instance_name = types.StringValue(gitopsInstance.Instance_name)
	state.Bits_account = types.Int64Value(int64(gitopsInstance.Bits_account))
	state.Order_time = types.StringValue(gitopsInstance.Order_time)
	state.Orderer_id = types.StringValue(gitopsInstance.Orderer_id)
	state.Service_id = types.Int64Value(int64(gitopsInstance.Service_id))
	state.Stage = types.StringValue(gitopsInstance.Stage)
	state.Replica_count = types.Int64Value(int64(gitopsInstance.Replica_count))
	state.Version = types.StringValue(gitopsInstance.Version)
	state.Some_value = types.StringValue(gitopsInstance.Some_value)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *gitopsInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan gitopsInstanceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var instanceUpdate gitopsclient.InstanceUpdate
	instanceUpdate.Instance_name = plan.Instance_name.ValueString()
	instanceUpdate.Bits_account = uint64(plan.Bits_account.ValueInt64())
	instanceUpdate.Service_id = uint64(plan.Service_id.ValueInt64())
	instanceUpdate.Replica_count = uint64(plan.Replica_count.ValueInt64())
	instanceUpdate.Version = plan.Version.ValueString()
	instanceUpdate.Some_value = plan.Some_value.ValueString()

	// Update existing gitopsInstance
	gitopsInstance, err := r.client.PutInstance(plan.Instance_id.ValueString(), instanceUpdate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Gitpos Instance",
			"Could not update order, unexpected error: "+err.Error(),
		)
		return
	}
	plan.Bits_account = types.Int64Value(int64(gitopsInstance.Bits_account))
	plan.Service_id = types.Int64Value(int64(gitopsInstance.Service_id))
	plan.Instance_name = types.StringValue(gitopsInstance.Instance_name)
	plan.Replica_count = types.Int64Value(int64(gitopsInstance.Replica_count))
	plan.Version = types.StringValue(gitopsInstance.Version)
	plan.Some_value = types.StringValue(gitopsInstance.Some_value)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *gitopsInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state gitopsInstanceResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	err := r.client.DeleteInstance(state.Instance_id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting gitopsInstance"+state.Instance_id.ValueString(),
			"Could not delete gitopsInstance, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *gitopsInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("instance_id"), req, resp)
}
