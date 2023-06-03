package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &DiscordProvider{}

type DiscordProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *DiscordProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "discord"
	resp.Version = p.version
}

func (p *DiscordProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
}

func (p *DiscordProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	token, found := os.LookupEnv("DISCORD_TOKEN")
	if !found {
		resp.Diagnostics.AddError("Provider Configure Error", "DISCORD_TOKEN environment variable must be set")
		return
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		resp.Diagnostics.AddError("Provider Configure Error", fmt.Sprintf("Error creating discord session: %s", err))
		return
	}

	resp.DataSourceData = session
}

func (p *DiscordProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *DiscordProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewGuildDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &DiscordProvider{
			version: version,
		}
	}
}
