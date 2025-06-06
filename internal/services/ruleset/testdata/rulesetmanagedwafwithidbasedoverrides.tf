
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_managed"

  rules = [{
    action = "execute"
    action_parameters = {
      id = "efb7b8c949ac4650a09736fc376e9aee"
      overrides = {
        rules = [{
          id      = "5de7edfa648c4d6891dc3e7f84534ffa"
          action  = "log"
          enabled = true
          },
          {
            id      = "e3a567afc347477d9702d9047e97d760"
            action  = "log"
            enabled = true
        }]
      }
    }

    expression  = "true"
    description = "make 5de7edfa648c4d6891dc3e7f84534ffa and e3a567afc347477d9702d9047e97d760 log only"
    enabled     = false
  }]
}
