// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure web proxy profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWebProxyProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure web proxy profiles.",

		ReadContext: dataSourceWebProxyProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"header_client_ip": {
				Type:        schema.TypeString,
				Description: "Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Computed:    true,
			},
			"header_front_end_https": {
				Type:        schema.TypeString,
				Description: "Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Computed:    true,
			},
			"header_via_request": {
				Type:        schema.TypeString,
				Description: "Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Computed:    true,
			},
			"header_via_response": {
				Type:        schema.TypeString,
				Description: "Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header.",
				Computed:    true,
			},
			"header_x_authenticated_groups": {
				Type:        schema.TypeString,
				Description: "Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Computed:    true,
			},
			"header_x_authenticated_user": {
				Type:        schema.TypeString,
				Description: "Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Computed:    true,
			},
			"header_x_forwarded_client_cert": {
				Type:        schema.TypeString,
				Description: "Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Computed:    true,
			},
			"header_x_forwarded_for": {
				Type:        schema.TypeString,
				Description: "Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Computed:    true,
			},
			"headers": {
				Type:        schema.TypeList,
				Description: "Configure HTTP forwarded requests headers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action when the HTTP header is forwarded.",
							Computed:    true,
						},
						"add_option": {
							Type:        schema.TypeString,
							Description: "Configure options to append content to existing HTTP header or add new HTTP header.",
							Computed:    true,
						},
						"base64_encoding": {
							Type:        schema.TypeString,
							Description: "Enable/disable use of base64 encoding of HTTP content.",
							Computed:    true,
						},
						"content": {
							Type:        schema.TypeString,
							Description: "HTTP header content.",
							Computed:    true,
						},
						"dstaddr": {
							Type:        schema.TypeList,
							Description: "Destination address and address group names.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Address name.",
										Computed:    true,
									},
								},
							},
						},
						"dstaddr6": {
							Type:        schema.TypeList,
							Description: "Destination address and address group names (IPv6).",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Address name.",
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "HTTP forwarded header id.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "HTTP forwarded header name.",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeString,
							Description: "Configure protocol(s) to take add-option action on (HTTP, HTTPS, or both).",
							Computed:    true,
						},
					},
				},
			},
			"log_header_change": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging HTTP header changes.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Profile name.",
				Required:    true,
			},
			"strip_encoding": {
				Type:        schema.TypeString,
				Description: "Enable/disable stripping unsupported encoding from the request header.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWebProxyProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	// c.Retries = 1

	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadWebProxyProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebProxyProfile dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectWebProxyProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
