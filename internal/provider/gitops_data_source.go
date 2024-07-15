package provider

import (
	"context"
	"fmt"

	"github.com/chillout2k/gitopsclient"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &gitopsDataSource{}
	_ datasource.DataSourceWithConfigure = &gitopsDataSource{}
)

// NewGitopsDataSource is a helper function to simplify the provider implementation.
func NewGitopsDataSource() datasource.DataSource {
	return &gitopsDataSource{}
}

// gitopsDataSource is the data source implementation.
type gitopsDataSource struct {
	//client *hashicups.Client
	client *gitopsclient.GitopsClient
}

// gitopsDataSourceModel maps the data source schema data.
/*type gitopsDataSourceModel struct {
	Plans []gitopsPlanModel `tfsdk:"plans"`
}

// gitopsPlanModel maps gitops instance schema data.
type gitopsPlanModel struct {
	Instance_name string `tfsdk:"instance_name"`
	Orderer_id    string `tfsdk:"orderer_id"`
	Bits_account  uint64 `tfsdk:"bits_account"`
	Service_id    uint64 `tfsdk:"service_id"`
	Replica_count uint64 `tfsdk:"replica_count"`
	Version       string `tfsdk:"version"`
	Some_value    string `tfsdk:"some_value"`
}*/

// Metadata returns the data source type name.
func (d *gitopsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plans"
}

// Schema defines the schema for the data source.
func (d *gitopsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"plans": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"instance_name": schema.StringAttribute{
							Computed: true,
						},
						"orderer_id": schema.StringAttribute{
							Computed: true,
						},
						"bits_account": schema.Int64Attribute{
							Computed: true,
						},
						"service_id": schema.Int64Attribute{
							Computed: true,
						},
						"replica_count": schema.Int64Attribute{
							Computed: true,
						},
						"version": schema.StringAttribute{
							Computed: true,
						},
						"some_value": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *gitopsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	/*var state gitopsDataSourceModel

	plans, err := d.client.getPlans()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read gitops Plans",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, plan := range plans {
		planState := gitopsPlanModel{
			Instance_name: plan.Instance_id,
			Orderer_id:    plan.Order_time,
			Bits_account:  uint64(plan.Bits_account),
			Service_id:    uint64(plan.Service_id),
			Replica_count: uint64(plan.Replica_count),
			Version:       plan.Version,
			Some_value:    plan.Some_value,
		}

		state.Plans = append(state.Plans, planState)
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}*/
}

// Configure adds the provider configured client to the data source.
func (d *gitopsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*gitopsclient.GitopsClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *hashicups.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}
