package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/chillout2k/gitopsclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &gitopsProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &gitopsProvider{
			version: version,
		}
	}
}

// gitopsProviderModel maps provider schema data to a Go type.
type gitopsProviderModel struct {
	Host                types.String `tfsdk:"host"`
	CachePath           types.String `tfsdk:"cache_path"`
	Username            types.String `tfsdk:"username"`
	Password            types.String `tfsdk:"password"`
	ClientId            types.String `tfsdk:"client_id"`
	ClientSecret        types.String `tfsdk:"client_secret"`
	TokenURI            types.String `tfsdk:"token_uri"`
	JwksURI             types.String `tfsdk:"jwks_uri"`
	AuthURI             types.String `tfsdk:"auth_uri"`
	RedirectURI         types.String `tfsdk:"redirect_uri"`
	AuthzListenerSocket types.String `tfsdk:"authz_listener_socket"`
	Scopes              types.String `tfsdk:"scopes"`
	GrantType           types.String `tfsdk:"grant_type"`
	Debug               types.Bool   `tfsdk:"debug"`
}

// gitopsProvider is the provider implementation.
type gitopsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *gitopsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "gitops"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *gitopsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Required: true,
			},
			"cache_path": schema.StringAttribute{
				Required: true,
			},
			"username": schema.StringAttribute{
				Optional: true,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"client_id": schema.StringAttribute{
				Required: true,
			},
			"client_secret": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"token_uri": schema.StringAttribute{
				Required: true,
			},
			"auth_uri": schema.StringAttribute{
				Required: true,
			},
			"jwks_uri": schema.StringAttribute{
				Required: true,
			},
			"redirect_uri": schema.StringAttribute{
				Optional: true,
			},
			"authz_listener_socket": schema.StringAttribute{
				Optional: true,
			},
			"scopes": schema.StringAttribute{
				Optional: true,
			},
			"grant_type": schema.StringAttribute{
				Required: true,
			},
			"debug": schema.BoolAttribute{
				Optional: true,
			},
		},
	}
}

func (p *gitopsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Gitops client")

	// Retrieve provider data from configuration
	var config gitopsProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.
	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown gitops API Host",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API host.",
		)
	}
	if config.CachePath.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("cache_path"),
			"Unknown gitops API cache_path",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API cache_path.",
		)
	}
	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown gitops API Username",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API username.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown gitops API password",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API password.",
		)
	}

	if config.ClientId.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_id"),
			"Unknown gitops API client_id",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API client_id.",
		)
	}

	if config.ClientSecret.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Unknown gitops API client_secret",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API client_secret.",
		)
	}

	if config.TokenURI.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("token_uri"),
			"Unknown gitops API token_uri",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API token_uri.",
		)
	}

	if config.JwksURI.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("jwks_uri"),
			"Unknown gitops API jwks_uri",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API jwks_uri.",
		)
	}

	if config.AuthURI.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("auth_uri"),
			"Unknown gitops API auth_uri",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API auth_uri.",
		)
	}

	if config.RedirectURI.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("redirect_uri"),
			"Unknown gitops API redirect_uri",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API redirect_uri.",
		)
	}

	if config.AuthzListenerSocket.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("authz_listener_socket"),
			"Unknown gitops API authz_listener_socket",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API authz_listener_socket.",
		)
	}

	if config.Scopes.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("scopes"),
			"Unknown gitops API scopes",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API scopes.",
		)
	}

	if config.GrantType.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("grant_type"),
			"Unknown gitops API grant_type",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API grant_type.",
		)
	}

	if config.Debug.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("debug"),
			"Unknown gitops API debug",
			"The provider cannot create the gitops API client as there is an unknown configuration value for the gitops API debug.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("GITOPS_HOST")
	cache_path := os.Getenv("GITOPS_CACHEPATH")
	username := os.Getenv("GITOPS_USERNAME")
	password := os.Getenv("GITOPS_PASSWORD")
	client_id := os.Getenv("GITOPS_CLIENTID")
	client_secret := os.Getenv("GITOPS_CLIENTSECRET")
	token_uri := os.Getenv("GITOPS_TOKENURI")
	jwks_uri := os.Getenv("GITOPS_JWKSURI")
	auth_uri := os.Getenv("GITOPS_AUTHURI")
	redirect_uri := os.Getenv("GITOPS_REDIRECTURI")
	authz_listener_socket := os.Getenv("GITOPS_AUTHZLISTENERSOCKET")
	scopes := os.Getenv("GITOPS_SCOPES")
	grant_type := os.Getenv("GITOPS_GRANTTYPE")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.CachePath.IsNull() {
		cache_path = config.CachePath.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	if !config.ClientId.IsNull() {
		client_id = config.ClientId.ValueString()
	}

	if !config.ClientSecret.IsNull() {
		client_secret = config.ClientSecret.ValueString()
	}

	if !config.TokenURI.IsNull() {
		token_uri = config.TokenURI.ValueString()
	}

	if !config.AuthURI.IsNull() {
		auth_uri = config.AuthURI.ValueString()
	}

	if !config.JwksURI.IsNull() {
		jwks_uri = config.JwksURI.ValueString()
	}

	if !config.RedirectURI.IsNull() {
		redirect_uri = config.RedirectURI.ValueString()
	}

	if !config.AuthzListenerSocket.IsNull() {
		authz_listener_socket = config.AuthzListenerSocket.ValueString()
	}

	if !config.Scopes.IsNull() {
		scopes = config.Scopes.ValueString()
	}

	if !config.GrantType.IsNull() {
		grant_type = config.GrantType.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Gitops API Host",
			"The provider cannot create the gitops API client as there is a missing or empty value for the gitops API host. "+
				"Set the host value in the configuration or use the GITOPS_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if cache_path == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("cache_path"),
			"Missing Gitops API cache_path",
			"The provider cannot create the gitops API client as there is a missing or empty value for the gitops API cache_path. "+
				"Set the host value in the configuration or use the GITOPS_CACHEPATH environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	/*if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing Gitops API username",
			"The provider cannot create the gitops API client as there is a missing or empty value for the gitops API username. "+
				"Set the host value in the configuration or use the GITOPS_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}*/

	/*if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Gitops API password",
			"The provider cannot create the gitops API client as there is a missing or empty value for the gitops API password. "+
				"Set the host value in the configuration or use the GITOPS_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}*/

	if client_id == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_id"),
			"Missing Gitops API client_id",
			"The provider cannot create the gitops API client as there is a missing or empty value for the gitops API client_id. "+
				"Set the host value in the configuration or use the GITOPS_CLIENTID environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	/*if client_secret == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Missing Gitops API client_secret",
			"The provider cannot create the gitops API client as there is a missing or empty value for the gitops API client_secret. "+
				"Set the host value in the configuration or use the GITOPS_CLIENTSECRET environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}*/

	if token_uri == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("token_uri"),
			"Missing Gitops API token_uri",
			"The provider cannot create the gitops API client as there is a missing or empty value for the gitops API token_uri. "+
				"Set the host value in the configuration or use the GITOPS_TOKENURI environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if auth_uri == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("auth_uri"),
			"Missing Gitops API auth_uri",
			"The provider cannot create the gitops API client as there is a missing or empty value for the gitops API auth_uri. "+
				"Set the host value in the configuration or use the GITOPS_AUTHURI environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if jwks_uri == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("jwks_uri"),
			"Missing Gitops API jwks_uri",
			"The provider cannot create the gitops API client as there is a missing or empty value for the gitops API jwks_uri. "+
				"Set the host value in the configuration or use the GITOPS_JWKSURI environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if grant_type == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("grant_type"),
			"Missing Gitops API grant_type",
			"The provider cannot create the gitops API client as there is a missing or empty value for the gitops API grant_type. "+
				"Set the host value in the configuration or use the GITOPS_GRANTTYPE environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "gitops_host", host)
	ctx = tflog.SetField(ctx, "gitops_username", username)
	ctx = tflog.SetField(ctx, "gitops_clientid", client_id)
	ctx = tflog.SetField(ctx, "gitops_tokenuri", token_uri)

	tflog.Debug(ctx, "Creating Gitops client")

	// Create a new gitops client using the configuration values
	clientConfig := gitopsclient.GitopsClientConfig{
		GitopsApiURI:        host,
		CachePath:           cache_path,
		ClientId:            client_id,
		ClientSecret:        client_secret,
		TokenURI:            token_uri,
		AuthURI:             auth_uri,
		JwksURI:             jwks_uri,
		RedirectURI:         redirect_uri,
		AuthzListenerSocket: authz_listener_socket,
		Scopes:              scopes,
		GrantType:           grant_type,
		Username:            username,
		Password:            password,
	}

	client, err := gitopsclient.NewGitopsClient(clientConfig)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Gitops API Client",
			"An unexpected error occurred when creating the gitops API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Gitops Client Error: "+err.Error(),
		)
		return
	}
	// get access-token according to OAuth2 settings (grant_type, client_id, ...)
	var isTokenValid bool = false
	err = client.GetTokenFromCache("access")
	if err == nil {
		_, err = client.ParseToken(client.AccessToken)
		if err == nil {
			isTokenValid = true
		}
	}
	if !isTokenValid {
		err = client.GetToken()
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Create Gitops API access-token",
				"An unexpected error occurred when creating the gitops API access-token. "+
					"If the error is not clear, please contact the provider developers.\n\n"+
					"Gitops Client Error: "+err.Error(),
			)
			return
		}
	}
	tflog.Debug(ctx, "Gitops Client acces_token: "+client.AccessToken)
	tflog.Debug(ctx, "Gitops Client refresh_token: "+client.RefreshToken)

	// Make the gitops client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured Gitops client", map[string]any{"success": true})
}

// Resources defines the resources implemented in the provider.
func (p *gitopsProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewGitopsInstanceResource,
	}
}

// DataSources defines the data sources implemented in the provider.
func (p *gitopsProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewGitopsDataSource,
	}
}
