// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Web application firewall configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWafProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Web application firewall configuration.",

		ReadContext: dataSourceWafProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"address_list": {
				Type:        schema.TypeList,
				Description: "Address block and allow lists.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"blocked_address": {
							Type:        schema.TypeList,
							Description: "Blocked address.",
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
						"blocked_log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging on blocked addresses.",
							Computed:    true,
						},
						"severity": {
							Type:        schema.TypeString,
							Description: "Severity.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Status.",
							Computed:    true,
						},
						"trusted_address": {
							Type:        schema.TypeList,
							Description: "Trusted address.",
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
					},
				},
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"constraint": {
				Type:        schema.TypeList,
				Description: "WAF HTTP protocol restrictions.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content_length": {
							Type:        schema.TypeList,
							Description: "HTTP content length in request.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"length": {
										Type:        schema.TypeInt,
										Description: "Length of HTTP content in bytes (0 to 2147483647).",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
						"exception": {
							Type:        schema.TypeList,
							Description: "HTTP constraint exception.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Description: "Host address.",
										Computed:    true,
									},
									"content_length": {
										Type:        schema.TypeString,
										Description: "HTTP content length in request.",
										Computed:    true,
									},
									"header_length": {
										Type:        schema.TypeString,
										Description: "HTTP header length in request.",
										Computed:    true,
									},
									"hostname": {
										Type:        schema.TypeString,
										Description: "Enable/disable hostname check.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Exception ID.",
										Computed:    true,
									},
									"line_length": {
										Type:        schema.TypeString,
										Description: "HTTP line length in request.",
										Computed:    true,
									},
									"malformed": {
										Type:        schema.TypeString,
										Description: "Enable/disable malformed HTTP request check.",
										Computed:    true,
									},
									"max_cookie": {
										Type:        schema.TypeString,
										Description: "Maximum number of cookies in HTTP request.",
										Computed:    true,
									},
									"max_header_line": {
										Type:        schema.TypeString,
										Description: "Maximum number of HTTP header line.",
										Computed:    true,
									},
									"max_range_segment": {
										Type:        schema.TypeString,
										Description: "Maximum number of range segments in HTTP range line.",
										Computed:    true,
									},
									"max_url_param": {
										Type:        schema.TypeString,
										Description: "Maximum number of parameters in URL.",
										Computed:    true,
									},
									"method": {
										Type:        schema.TypeString,
										Description: "Enable/disable HTTP method check.",
										Computed:    true,
									},
									"param_length": {
										Type:        schema.TypeString,
										Description: "Maximum length of parameter in URL, HTTP POST request or HTTP body.",
										Computed:    true,
									},
									"pattern": {
										Type:        schema.TypeString,
										Description: "URL pattern.",
										Computed:    true,
									},
									"regex": {
										Type:        schema.TypeString,
										Description: "Enable/disable regular expression based pattern match.",
										Computed:    true,
									},
									"url_param_length": {
										Type:        schema.TypeString,
										Description: "Maximum length of parameter in URL.",
										Computed:    true,
									},
									"version": {
										Type:        schema.TypeString,
										Description: "Enable/disable HTTP version check.",
										Computed:    true,
									},
								},
							},
						},
						"header_length": {
							Type:        schema.TypeList,
							Description: "HTTP header length in request.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"length": {
										Type:        schema.TypeInt,
										Description: "Length of HTTP header in bytes (0 to 2147483647).",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
						"hostname": {
							Type:        schema.TypeList,
							Description: "Enable/disable hostname check.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
						"line_length": {
							Type:        schema.TypeList,
							Description: "HTTP line length in request.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"length": {
										Type:        schema.TypeInt,
										Description: "Length of HTTP line in bytes (0 to 2147483647).",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
						"malformed": {
							Type:        schema.TypeList,
							Description: "Enable/disable malformed HTTP request check.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
						"max_cookie": {
							Type:        schema.TypeList,
							Description: "Maximum number of cookies in HTTP request.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"max_cookie": {
										Type:        schema.TypeInt,
										Description: "Maximum number of cookies in HTTP request (0 to 2147483647).",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
						"max_header_line": {
							Type:        schema.TypeList,
							Description: "Maximum number of HTTP header line.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"max_header_line": {
										Type:        schema.TypeInt,
										Description: "Maximum number HTTP header lines (0 to 2147483647).",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
						"max_range_segment": {
							Type:        schema.TypeList,
							Description: "Maximum number of range segments in HTTP range line.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"max_range_segment": {
										Type:        schema.TypeInt,
										Description: "Maximum number of range segments in HTTP range line (0 to 2147483647).",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
						"max_url_param": {
							Type:        schema.TypeList,
							Description: "Maximum number of parameters in URL.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"max_url_param": {
										Type:        schema.TypeInt,
										Description: "Maximum number of parameters in URL (0 to 2147483647).",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
						"method": {
							Type:        schema.TypeList,
							Description: "Enable/disable HTTP method check.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
						"param_length": {
							Type:        schema.TypeList,
							Description: "Maximum length of parameter in URL, HTTP POST request or HTTP body.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"length": {
										Type:        schema.TypeInt,
										Description: "Maximum length of parameter in URL, HTTP POST request or HTTP body in bytes (0 to 2147483647).",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
						"url_param_length": {
							Type:        schema.TypeList,
							Description: "Maximum length of parameter in URL.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"length": {
										Type:        schema.TypeInt,
										Description: "Maximum length of URL parameter in bytes (0 to 2147483647).",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
						"version": {
							Type:        schema.TypeList,
							Description: "Enable/disable HTTP version check.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable the constraint.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"extended_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable extended logging.",
				Computed:    true,
			},
			"external": {
				Type:        schema.TypeString,
				Description: "Disable/Enable external HTTP Inspection.",
				Computed:    true,
			},
			"method": {
				Type:        schema.TypeList,
				Description: "Method restriction.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_allowed_methods": {
							Type:        schema.TypeString,
							Description: "Methods.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"method_policy": {
							Type:        schema.TypeList,
							Description: "HTTP method policy.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Description: "Host address.",
										Computed:    true,
									},
									"allowed_methods": {
										Type:        schema.TypeString,
										Description: "Allowed Methods.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "HTTP method policy ID.",
										Computed:    true,
									},
									"pattern": {
										Type:        schema.TypeString,
										Description: "URL pattern.",
										Computed:    true,
									},
									"regex": {
										Type:        schema.TypeString,
										Description: "Enable/disable regular expression based pattern match.",
										Computed:    true,
									},
								},
							},
						},
						"severity": {
							Type:        schema.TypeString,
							Description: "Severity.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Status.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "WAF Profile name.",
				Required:    true,
			},
			"signature": {
				Type:        schema.TypeList,
				Description: "WAF signatures.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credit_card_detection_threshold": {
							Type:        schema.TypeInt,
							Description: "The minimum number of Credit cards to detect violation.",
							Computed:    true,
						},
						"custom_signature": {
							Type:        schema.TypeList,
							Description: "Custom signature.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"case_sensitivity": {
										Type:        schema.TypeString,
										Description: "Case sensitivity in pattern.",
										Computed:    true,
									},
									"direction": {
										Type:        schema.TypeString,
										Description: "Traffic direction.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Signature name.",
										Computed:    true,
									},
									"pattern": {
										Type:        schema.TypeString,
										Description: "Match pattern.",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Status.",
										Computed:    true,
									},
									"target": {
										Type:        schema.TypeString,
										Description: "Match HTTP target.",
										Computed:    true,
									},
								},
							},
						},
						"disabled_signature": {
							Type:        schema.TypeList,
							Description: "Disabled signatures",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Signature ID.",
										Computed:    true,
									},
								},
							},
						},
						"disabled_sub_class": {
							Type:        schema.TypeList,
							Description: "Disabled signature subclasses.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Signature subclass ID.",
										Computed:    true,
									},
								},
							},
						},
						"main_class": {
							Type:        schema.TypeList,
							Description: "Main signature class.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Main signature class ID.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"severity": {
										Type:        schema.TypeString,
										Description: "Severity.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Status.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"url_access": {
				Type:        schema.TypeList,
				Description: "URL access list",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_pattern": {
							Type:        schema.TypeList,
							Description: "URL access pattern.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "URL access pattern ID.",
										Computed:    true,
									},
									"negate": {
										Type:        schema.TypeString,
										Description: "Enable/disable match negation.",
										Computed:    true,
									},
									"pattern": {
										Type:        schema.TypeString,
										Description: "URL pattern.",
										Computed:    true,
									},
									"regex": {
										Type:        schema.TypeString,
										Description: "Enable/disable regular expression based pattern match.",
										Computed:    true,
									},
									"srcaddr": {
										Type:        schema.TypeString,
										Description: "Source address.",
										Computed:    true,
									},
								},
							},
						},
						"action": {
							Type:        schema.TypeString,
							Description: "Action.",
							Computed:    true,
						},
						"address": {
							Type:        schema.TypeString,
							Description: "Host address.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "URL access ID.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"severity": {
							Type:        schema.TypeString,
							Description: "Severity.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWafProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWafProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WafProfile dataSource: %v", err)
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

	diags := refreshObjectWafProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
