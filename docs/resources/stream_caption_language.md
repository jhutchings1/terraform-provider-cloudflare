---
page_title: "cloudflare_stream_caption_language Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_stream_caption_language (Resource)



## Example Usage

```terraform
resource "cloudflare_stream_caption_language" "example_stream_caption_language" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  identifier = "ea95132c15732412d22c1476fa83f27a"
  language = "tr"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier.
- `identifier` (String) A Cloudflare-generated unique identifier for a media item.
- `language` (String) The language tag in BCP 47 format.

### Optional

- `file` (String) The WebVTT file containing the caption or subtitle content.

### Read-Only

- `generated` (Boolean) Whether the caption was generated via AI.
- `label` (String) The language label displayed in the native language to users.
- `status` (String) The status of a generated caption.
Available values: "ready", "inprogress", "error".


