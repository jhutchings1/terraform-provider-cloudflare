---
page_title: "cloudflare_authenticated_origin_pulls_certificates Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_authenticated_origin_pulls_certificates (Data Source)



## Example Usage

```terraform
data "cloudflare_authenticated_origin_pulls_certificates" "example_authenticated_origin_pulls_certificates" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) Identifier

### Optional

- `max_items` (Number) Max items to fetch, default: 1000

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `certificate` (String) The zone's leaf certificate.
- `enabled` (Boolean) Indicates whether zone-level authenticated origin pulls is enabled.
- `expires_on` (String) When the certificate from the authority expires.
- `id` (String) Identifier
- `issuer` (String) The certificate authority that issued the certificate.
- `private_key` (String, Sensitive) The zone's private key.
- `signature` (String) The type of hash used for the certificate.
- `status` (String) Status of the certificate activation.
Available values: "initializing", "pending_deployment", "pending_deletion", "active", "deleted", "deployment_timed_out", "deletion_timed_out".
- `uploaded_on` (String) This is the time the certificate was uploaded.


