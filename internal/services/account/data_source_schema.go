// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Computed:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "Timestamp for the creation of the account",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Account name",
				Computed:    true,
			},
			"settings": schema.SingleNestedAttribute{
				Description: "Account settings",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[AccountSettingsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"abuse_contact_email": schema.StringAttribute{
						Description: "Sets an abuse contact email to notify for abuse reports.",
						Computed:    true,
					},
					"default_nameservers": schema.StringAttribute{
						Description:        "Specifies the default nameservers to be used for new zones added to this account.\n\n- `cloudflare.standard` for Cloudflare-branded nameservers\n- `custom.account` for account custom nameservers\n- `custom.tenant` for tenant custom nameservers\n\nSee [Custom Nameservers](https://developers.cloudflare.com/dns/additional-options/custom-nameservers/)\nfor more information.\n\nDeprecated in favor of [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).\nAvailable values: \"cloudflare.standard\", \"custom.account\", \"custom.tenant\".",
						Computed:           true,
						DeprecationMessage: "This attribute is deprecated.",
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"cloudflare.standard",
								"custom.account",
								"custom.tenant",
							),
						},
					},
					"enforce_twofactor": schema.BoolAttribute{
						Description: "Indicates whether membership in this account requires that\nTwo-Factor Authentication is enabled",
						Computed:    true,
					},
					"use_account_custom_ns_by_default": schema.BoolAttribute{
						Description:        "Indicates whether new zones should use the account-level custom\nnameservers by default.\n\nDeprecated in favor of [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).",
						Computed:           true,
						DeprecationMessage: "This attribute is deprecated.",
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "Direction to order results.\nAvailable values: \"asc\", \"desc\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of the account.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *AccountDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AccountDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("account_id"), path.MatchRoot("filter")),
	}
}
