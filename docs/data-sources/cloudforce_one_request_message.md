---
page_title: "cloudflare_cloudforce_one_request_message Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_cloudforce_one_request_message (Data Source)



## Example Usage

```terraform
data "cloudflare_cloudforce_one_request_message" "example_cloudforce_one_request_message" {
  account_identifier = "023e105f4ecef8ad9ca31a8372d0c353"
  request_identifier = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_identifier` (String) Identifier
- `request_identifier` (String) UUID

### Read-Only

- `author` (String) Author of message
- `content` (String) Content of message
- `created` (String) Message creation time
- `id` (Number) Message ID
- `is_follow_on_request` (Boolean) Whether the message is a follow-on request
- `updated` (String) Message last updated time


