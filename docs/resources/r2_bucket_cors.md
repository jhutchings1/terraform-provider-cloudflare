---
page_title: "cloudflare_r2_bucket_cors Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_r2_bucket_cors (Resource)



## Example Usage

```terraform
resource "cloudflare_r2_bucket_cors" "example_r2_bucket_cors" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  bucket_name = "example-bucket"
  rules = [{
    allowed = {
      methods = ["GET"]
      origins = ["http://localhost:3000"]
      headers = ["x-requested-by"]
    }
    id = "Allow Local Development"
    expose_headers = ["Content-Encoding"]
    max_age_seconds = 3600
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Account ID
- `bucket_name` (String) Name of the bucket

### Optional

- `jurisdiction` (String) Jurisdiction of the bucket
- `rules` (Attributes List) (see [below for nested schema](#nestedatt--rules))

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Required:

- `allowed` (Attributes) Object specifying allowed origins, methods and headers for this CORS rule. (see [below for nested schema](#nestedatt--rules--allowed))

Optional:

- `expose_headers` (List of String) Specifies the headers that can be exposed back, and accessed by, the JavaScript making the cross-origin request. If you need to access headers beyond the safelisted response headers, such as Content-Encoding or cf-cache-status, you must specify it here.
- `id` (String) Identifier for this rule
- `max_age_seconds` (Number) Specifies the amount of time (in seconds) browsers are allowed to cache CORS preflight responses. Browsers may limit this to 2 hours or less, even if the maximum value (86400) is specified.

<a id="nestedatt--rules--allowed"></a>
### Nested Schema for `rules.allowed`

Required:

- `methods` (List of String) Specifies the value for the Access-Control-Allow-Methods header R2 sets when requesting objects in a bucket from a browser.
- `origins` (List of String) Specifies the value for the Access-Control-Allow-Origin header R2 sets when requesting objects in a bucket from a browser.

Optional:

- `headers` (List of String) Specifies the value for the Access-Control-Allow-Headers header R2 sets when requesting objects in this bucket from a browser. Cross-origin requests that include custom headers (e.g. x-user-id) should specify these headers as AllowedHeaders.


