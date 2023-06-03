package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGuildDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `data "discord_guild" "test" {
					id = "768357607931904022"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.discord_guild.test", "id", "768357607931904022"),
					resource.TestCheckResourceAttr("data.discord_guild.test", "name", "dev"),
				),
			},
		},
	})
}
