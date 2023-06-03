---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "discord_guild Data Source - terraform-provider-discord"
subcategory: ""
description: |-
  Retrieves a Discord guild, which represents an isolated collection of users and channels and are often referred to as "servers" in the UI.
---

# discord_guild (Data Source)

Retrieves a Discord guild, which represents an isolated collection of users and channels and are often referred to as "servers" in the UI.

## Example Usage

```terraform
data "discord_guild" "guild" {
  id = "12345"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Guild ID

### Read-Only

- `name` (String) Guild name

