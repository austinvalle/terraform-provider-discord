package provider

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &GuildDataSource{}

func NewGuildDataSource() datasource.DataSource {
	return &GuildDataSource{}
}

type GuildDataSource struct {
	session *discordgo.Session
}

// https://discord.com/developers/docs/resources/guild#guild-object
type GuildDataSourceModel struct {
	// TODO: snowflake ID
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (d *GuildDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_guild"
}

func (d *GuildDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves a Discord guild, which represents an isolated collection of users and channels and are often referred to as \"servers\" in the UI.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Guild ID",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Guild name",
				Computed:            true,
			},
		},
	}
}

func (d *GuildDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	session, ok := req.ProviderData.(*discordgo.Session)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *discordgo.Session, got: %T. Please report this issue to the provider developer.", req.ProviderData),
		)

		return
	}

	d.session = session
}

func (d *GuildDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GuildDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "reading guild", map[string]interface{}{
		"id": data.Id.ValueString(),
	})

	guild, err := d.session.Guild(data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read guild: %s", err))
		return
	}

	data.Id = types.StringValue(guild.ID)
	data.Name = types.StringValue(guild.Name)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
