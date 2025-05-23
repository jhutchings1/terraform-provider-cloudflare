---
page_title: "cloudflare_magic_transit_site_wans Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_magic_transit_site_wans (Data Source)



## Example Usage

```terraform
data "cloudflare_magic_transit_site_wans" "example_magic_transit_site_wans" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  site_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier
- `site_id` (String) Identifier

### Optional

- `max_items` (Number) Max items to fetch, default: 1000

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `health_check_rate` (String) Magic WAN health check rate for tunnels created on this link. The default value is `mid`.
Available values: "low", "mid", "high".
- `id` (String) Identifier
- `name` (String)
- `physport` (Number)
- `priority` (Number) Priority of WAN for traffic loadbalancing.
- `site_id` (String) Identifier
- `static_addressing` (Attributes) (optional) if omitted, use DHCP. Submit secondary_address when site is in high availability mode. (see [below for nested schema](#nestedatt--result--static_addressing))
- `vlan_tag` (Number) VLAN ID. Use zero for untagged.

<a id="nestedatt--result--static_addressing"></a>
### Nested Schema for `result.static_addressing`

Read-Only:

- `address` (String) A valid CIDR notation representing an IP range.
- `gateway_address` (String) A valid IPv4 address.
- `secondary_address` (String) A valid CIDR notation representing an IP range.


