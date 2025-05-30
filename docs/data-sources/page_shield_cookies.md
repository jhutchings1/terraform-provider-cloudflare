---
page_title: "cloudflare_page_shield_cookies Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_page_shield_cookies (Data Source)



## Example Usage

```terraform
data "cloudflare_page_shield_cookies" "example_page_shield_cookies" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  cookie_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cookie_id` (String) Identifier
- `zone_id` (String) Identifier

### Read-Only

- `domain_attribute` (String)
- `expires_attribute` (String)
- `first_seen_at` (String)
- `host` (String)
- `http_only_attribute` (Boolean)
- `id` (String) Identifier
- `last_seen_at` (String)
- `max_age_attribute` (Number)
- `name` (String)
- `page_urls` (List of String)
- `path_attribute` (String)
- `same_site_attribute` (String) Available values: "lax", "strict", "none".
- `secure_attribute` (Boolean)
- `type` (String) Available values: "first_party", "unknown".


