// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Web application firewall configuration.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceWafProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Web application firewall configuration.",

		CreateContext: resourceWafProfileCreate,
		ReadContext:   resourceWafProfileRead,
		UpdateContext: resourceWafProfileUpdate,
		DeleteContext: resourceWafProfileDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"address_list": {
				Type:        schema.TypeList,
				Description: "Address block and allow lists.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"blocked_address": {
							Type:        schema.TypeList,
							Description: "Blocked address.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Address name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"blocked_log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging on blocked addresses.",
							Optional:    true,
							Computed:    true,
						},
						"severity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

							Description: "Severity.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Status.",
							Optional:    true,
							Computed:    true,
						},
						"trusted_address": {
							Type:        schema.TypeList,
							Description: "Trusted address.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Address name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"constraint": {
				Type:        schema.TypeList,
				Description: "WAF HTTP protocol restrictions.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content_length": {
							Type:        schema.TypeList,
							Description: "HTTP content length in request.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"length": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 2147483647),

										Description: "Length of HTTP content in bytes (0 to 2147483647).",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"exception": {
							Type:        schema.TypeList,
							Description: "HTTP constraint exception.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Host address.",
										Optional:    true,
										Computed:    true,
									},
									"content_length": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "HTTP content length in request.",
										Optional:    true,
										Computed:    true,
									},
									"header_length": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "HTTP header length in request.",
										Optional:    true,
										Computed:    true,
									},
									"hostname": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable hostname check.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Exception ID.",
										Optional:    true,
										Computed:    true,
									},
									"line_length": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "HTTP line length in request.",
										Optional:    true,
										Computed:    true,
									},
									"malformed": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable malformed HTTP request check.",
										Optional:    true,
										Computed:    true,
									},
									"max_cookie": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Maximum number of cookies in HTTP request.",
										Optional:    true,
										Computed:    true,
									},
									"max_header_line": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Maximum number of HTTP header line.",
										Optional:    true,
										Computed:    true,
									},
									"max_range_segment": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Maximum number of range segments in HTTP range line.",
										Optional:    true,
										Computed:    true,
									},
									"max_url_param": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Maximum number of parameters in URL.",
										Optional:    true,
										Computed:    true,
									},
									"method": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable HTTP method check.",
										Optional:    true,
										Computed:    true,
									},
									"param_length": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Maximum length of parameter in URL, HTTP POST request or HTTP body.",
										Optional:    true,
										Computed:    true,
									},
									"pattern": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),

										Description: "URL pattern.",
										Optional:    true,
										Computed:    true,
									},
									"regex": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable regular expression based pattern match.",
										Optional:    true,
										Computed:    true,
									},
									"url_param_length": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Maximum length of parameter in URL.",
										Optional:    true,
										Computed:    true,
									},
									"version": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable HTTP version check.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"header_length": {
							Type:        schema.TypeList,
							Description: "HTTP header length in request.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"length": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 2147483647),

										Description: "Length of HTTP header in bytes (0 to 2147483647).",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"hostname": {
							Type:        schema.TypeList,
							Description: "Enable/disable hostname check.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"line_length": {
							Type:        schema.TypeList,
							Description: "HTTP line length in request.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"length": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 2147483647),

										Description: "Length of HTTP line in bytes (0 to 2147483647).",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"malformed": {
							Type:        schema.TypeList,
							Description: "Enable/disable malformed HTTP request check.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"max_cookie": {
							Type:        schema.TypeList,
							Description: "Maximum number of cookies in HTTP request.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"max_cookie": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 2147483647),

										Description: "Maximum number of cookies in HTTP request (0 to 2147483647).",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"max_header_line": {
							Type:        schema.TypeList,
							Description: "Maximum number of HTTP header line.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"max_header_line": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 2147483647),

										Description: "Maximum number HTTP header lines (0 to 2147483647).",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"max_range_segment": {
							Type:        schema.TypeList,
							Description: "Maximum number of range segments in HTTP range line.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"max_range_segment": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 2147483647),

										Description: "Maximum number of range segments in HTTP range line (0 to 2147483647).",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"max_url_param": {
							Type:        schema.TypeList,
							Description: "Maximum number of parameters in URL.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"max_url_param": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 2147483647),

										Description: "Maximum number of parameters in URL (0 to 2147483647).",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"method": {
							Type:        schema.TypeList,
							Description: "Enable/disable HTTP method check.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"param_length": {
							Type:        schema.TypeList,
							Description: "Maximum length of parameter in URL, HTTP POST request or HTTP body.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"length": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 2147483647),

										Description: "Maximum length of parameter in URL, HTTP POST request or HTTP body in bytes (0 to 2147483647).",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"url_param_length": {
							Type:        schema.TypeList,
							Description: "Maximum length of parameter in URL.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"length": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 2147483647),

										Description: "Maximum length of URL parameter in bytes (0 to 2147483647).",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"version": {
							Type:        schema.TypeList,
							Description: "Enable/disable HTTP version check.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the constraint.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"extended_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable extended logging.",
				Optional:    true,
				Computed:    true,
			},
			"external": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Disable/Enable external HTTP Inspection.",
				Optional:    true,
				Computed:    true,
			},
			"method": {
				Type:        schema.TypeList,
				Description: "Method restriction.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_allowed_methods": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Methods.",
							Optional:         true,
							Computed:         true,
						},
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging.",
							Optional:    true,
							Computed:    true,
						},
						"method_policy": {
							Type:        schema.TypeList,
							Description: "HTTP method policy.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Host address.",
										Optional:    true,
										Computed:    true,
									},
									"allowed_methods": {
										Type: schema.TypeString,

										DiffSuppressFunc: suppressors.DiffFakeListEqual,
										Description:      "Allowed Methods.",
										Optional:         true,
										Computed:         true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "HTTP method policy ID.",
										Optional:    true,
										Computed:    true,
									},
									"pattern": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),

										Description: "URL pattern.",
										Optional:    true,
										Computed:    true,
									},
									"regex": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable regular expression based pattern match.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"severity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

							Description: "Severity.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Status.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WAF Profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"signature": {
				Type:        schema.TypeList,
				Description: "WAF signatures.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credit_card_detection_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 128),

							Description: "The minimum number of Credit cards to detect violation.",
							Optional:    true,
							Computed:    true,
						},
						"custom_signature": {
							Type:        schema.TypeList,
							Description: "Custom signature.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block", "erase"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"case_sensitivity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "Case sensitivity in pattern.",
										Optional:    true,
										Computed:    true,
									},
									"direction": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"request", "response"}, false),

										Description: "Traffic direction.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Signature name.",
										Optional:    true,
										Computed:    true,
									},
									"pattern": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),

										Description: "Match pattern.",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Status.",
										Optional:    true,
										Computed:    true,
									},
									"target": {
										Type: schema.TypeString,

										DiffSuppressFunc: suppressors.DiffFakeListEqual,
										Description:      "Match HTTP target.",
										Optional:         true,
										Computed:         true,
									},
								},
							},
						},
						"disabled_signature": {
							Type:        schema.TypeList,
							Description: "Disabled signatures.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Signature ID.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"disabled_sub_class": {
							Type:        schema.TypeList,
							Description: "Disabled signature subclasses.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Signature subclass ID.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"main_class": {
							Type:        schema.TypeList,
							Description: "Main signature class.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "block", "erase"}, false),

										Description: "Action.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Main signature class ID.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"severity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

										Description: "Severity.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Status.",
										Optional:    true,
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
				Description: "URL access list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_pattern": {
							Type:        schema.TypeList,
							Description: "URL access pattern.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "URL access pattern ID.",
										Optional:    true,
										Computed:    true,
									},
									"negate": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable match negation.",
										Optional:    true,
										Computed:    true,
									},
									"pattern": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),

										Description: "URL pattern.",
										Optional:    true,
										Computed:    true,
									},
									"regex": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable regular expression based pattern match.",
										Optional:    true,
										Computed:    true,
									},
									"srcaddr": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Source address.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "permit", "block"}, false),

							Description: "Action.",
							Optional:    true,
							Computed:    true,
						},
						"address": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Host address.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "URL access ID.",
							Optional:    true,
							Computed:    true,
						},
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging.",
							Optional:    true,
							Computed:    true,
						},
						"severity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

							Description: "Severity.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceWafProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating WafProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWafProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWafProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WafProfile")
	}

	return resourceWafProfileRead(ctx, d, meta)
}

func resourceWafProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
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

	obj, diags := getObjectWafProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWafProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WafProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WafProfile")
	}

	return resourceWafProfileRead(ctx, d, meta)
}

func resourceWafProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	err := c.Cmdb.DeleteWafProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WafProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadWafProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WafProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
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
	return nil
}

func flattenWafProfileAddressList(v *[]models.WafProfileAddressList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.BlockedAddress; tmp != nil {
				v["blocked_address"] = flattenWafProfileAddressListBlockedAddress(tmp, sort)
			}

			if tmp := cfg.BlockedLog; tmp != nil {
				v["blocked_log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.TrustedAddress; tmp != nil {
				v["trusted_address"] = flattenWafProfileAddressListTrustedAddress(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileAddressListBlockedAddress(v *[]models.WafProfileAddressListBlockedAddress, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWafProfileAddressListTrustedAddress(v *[]models.WafProfileAddressListTrustedAddress, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWafProfileConstraint(v *[]models.WafProfileConstraint, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.ContentLength; tmp != nil {
				v["content_length"] = flattenWafProfileConstraintContentLength(tmp, sort)
			}

			if tmp := cfg.Exception; tmp != nil {
				v["exception"] = flattenWafProfileConstraintException(tmp, sort)
			}

			if tmp := cfg.HeaderLength; tmp != nil {
				v["header_length"] = flattenWafProfileConstraintHeaderLength(tmp, sort)
			}

			if tmp := cfg.Hostname; tmp != nil {
				v["hostname"] = flattenWafProfileConstraintHostname(tmp, sort)
			}

			if tmp := cfg.LineLength; tmp != nil {
				v["line_length"] = flattenWafProfileConstraintLineLength(tmp, sort)
			}

			if tmp := cfg.Malformed; tmp != nil {
				v["malformed"] = flattenWafProfileConstraintMalformed(tmp, sort)
			}

			if tmp := cfg.MaxCookie; tmp != nil {
				v["max_cookie"] = flattenWafProfileConstraintMaxCookie(tmp, sort)
			}

			if tmp := cfg.MaxHeaderLine; tmp != nil {
				v["max_header_line"] = flattenWafProfileConstraintMaxHeaderLine(tmp, sort)
			}

			if tmp := cfg.MaxRangeSegment; tmp != nil {
				v["max_range_segment"] = flattenWafProfileConstraintMaxRangeSegment(tmp, sort)
			}

			if tmp := cfg.MaxUrlParam; tmp != nil {
				v["max_url_param"] = flattenWafProfileConstraintMaxUrlParam(tmp, sort)
			}

			if tmp := cfg.Method; tmp != nil {
				v["method"] = flattenWafProfileConstraintMethod(tmp, sort)
			}

			if tmp := cfg.ParamLength; tmp != nil {
				v["param_length"] = flattenWafProfileConstraintParamLength(tmp, sort)
			}

			if tmp := cfg.UrlParamLength; tmp != nil {
				v["url_param_length"] = flattenWafProfileConstraintUrlParamLength(tmp, sort)
			}

			if tmp := cfg.Version; tmp != nil {
				v["version"] = flattenWafProfileConstraintVersion(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintContentLength(v *[]models.WafProfileConstraintContentLength, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Length; tmp != nil {
				v["length"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintException(v *[]models.WafProfileConstraintException, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Address; tmp != nil {
				v["address"] = *tmp
			}

			if tmp := cfg.ContentLength; tmp != nil {
				v["content_length"] = *tmp
			}

			if tmp := cfg.HeaderLength; tmp != nil {
				v["header_length"] = *tmp
			}

			if tmp := cfg.Hostname; tmp != nil {
				v["hostname"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.LineLength; tmp != nil {
				v["line_length"] = *tmp
			}

			if tmp := cfg.Malformed; tmp != nil {
				v["malformed"] = *tmp
			}

			if tmp := cfg.MaxCookie; tmp != nil {
				v["max_cookie"] = *tmp
			}

			if tmp := cfg.MaxHeaderLine; tmp != nil {
				v["max_header_line"] = *tmp
			}

			if tmp := cfg.MaxRangeSegment; tmp != nil {
				v["max_range_segment"] = *tmp
			}

			if tmp := cfg.MaxUrlParam; tmp != nil {
				v["max_url_param"] = *tmp
			}

			if tmp := cfg.Method; tmp != nil {
				v["method"] = *tmp
			}

			if tmp := cfg.ParamLength; tmp != nil {
				v["param_length"] = *tmp
			}

			if tmp := cfg.Pattern; tmp != nil {
				v["pattern"] = *tmp
			}

			if tmp := cfg.Regex; tmp != nil {
				v["regex"] = *tmp
			}

			if tmp := cfg.UrlParamLength; tmp != nil {
				v["url_param_length"] = *tmp
			}

			if tmp := cfg.Version; tmp != nil {
				v["version"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWafProfileConstraintHeaderLength(v *[]models.WafProfileConstraintHeaderLength, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Length; tmp != nil {
				v["length"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintHostname(v *[]models.WafProfileConstraintHostname, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintLineLength(v *[]models.WafProfileConstraintLineLength, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Length; tmp != nil {
				v["length"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintMalformed(v *[]models.WafProfileConstraintMalformed, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintMaxCookie(v *[]models.WafProfileConstraintMaxCookie, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.MaxCookie; tmp != nil {
				v["max_cookie"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintMaxHeaderLine(v *[]models.WafProfileConstraintMaxHeaderLine, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.MaxHeaderLine; tmp != nil {
				v["max_header_line"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintMaxRangeSegment(v *[]models.WafProfileConstraintMaxRangeSegment, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.MaxRangeSegment; tmp != nil {
				v["max_range_segment"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintMaxUrlParam(v *[]models.WafProfileConstraintMaxUrlParam, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.MaxUrlParam; tmp != nil {
				v["max_url_param"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintMethod(v *[]models.WafProfileConstraintMethod, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintParamLength(v *[]models.WafProfileConstraintParamLength, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Length; tmp != nil {
				v["length"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintUrlParamLength(v *[]models.WafProfileConstraintUrlParamLength, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Length; tmp != nil {
				v["length"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileConstraintVersion(v *[]models.WafProfileConstraintVersion, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileMethod(v *[]models.WafProfileMethod, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.DefaultAllowedMethods; tmp != nil {
				v["default_allowed_methods"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.MethodPolicy; tmp != nil {
				v["method_policy"] = flattenWafProfileMethodMethodPolicy(tmp, sort)
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileMethodMethodPolicy(v *[]models.WafProfileMethodMethodPolicy, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Address; tmp != nil {
				v["address"] = *tmp
			}

			if tmp := cfg.AllowedMethods; tmp != nil {
				v["allowed_methods"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Pattern; tmp != nil {
				v["pattern"] = *tmp
			}

			if tmp := cfg.Regex; tmp != nil {
				v["regex"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWafProfileSignature(v *[]models.WafProfileSignature, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CreditCardDetectionThreshold; tmp != nil {
				v["credit_card_detection_threshold"] = *tmp
			}

			if tmp := cfg.CustomSignature; tmp != nil {
				v["custom_signature"] = flattenWafProfileSignatureCustomSignature(tmp, sort)
			}

			if tmp := cfg.DisabledSignature; tmp != nil {
				v["disabled_signature"] = flattenWafProfileSignatureDisabledSignature(tmp, sort)
			}

			if tmp := cfg.DisabledSubClass; tmp != nil {
				v["disabled_sub_class"] = flattenWafProfileSignatureDisabledSubClass(tmp, sort)
			}

			if tmp := cfg.MainClass; tmp != nil {
				v["main_class"] = flattenWafProfileSignatureMainClass(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWafProfileSignatureCustomSignature(v *[]models.WafProfileSignatureCustomSignature, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.CaseSensitivity; tmp != nil {
				v["case_sensitivity"] = *tmp
			}

			if tmp := cfg.Direction; tmp != nil {
				v["direction"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Pattern; tmp != nil {
				v["pattern"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Target; tmp != nil {
				v["target"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWafProfileSignatureDisabledSignature(v *[]models.WafProfileSignatureDisabledSignature, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWafProfileSignatureDisabledSubClass(v *[]models.WafProfileSignatureDisabledSubClass, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWafProfileSignatureMainClass(v *[]models.WafProfileSignatureMainClass, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWafProfileUrlAccess(v *[]models.WafProfileUrlAccess, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AccessPattern; tmp != nil {
				v["access_pattern"] = flattenWafProfileUrlAccessAccessPattern(tmp, sort)
			}

			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Address; tmp != nil {
				v["address"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWafProfileUrlAccessAccessPattern(v *[]models.WafProfileUrlAccessAccessPattern, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Negate; tmp != nil {
				v["negate"] = *tmp
			}

			if tmp := cfg.Pattern; tmp != nil {
				v["pattern"] = *tmp
			}

			if tmp := cfg.Regex; tmp != nil {
				v["regex"] = *tmp
			}

			if tmp := cfg.Srcaddr; tmp != nil {
				v["srcaddr"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectWafProfile(d *schema.ResourceData, o *models.WafProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddressList != nil {
		if err = d.Set("address_list", flattenWafProfileAddressList(o.AddressList, sort)); err != nil {
			return diag.Errorf("error reading address_list: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Constraint != nil {
		if err = d.Set("constraint", flattenWafProfileConstraint(o.Constraint, sort)); err != nil {
			return diag.Errorf("error reading constraint: %v", err)
		}
	}

	if o.ExtendedLog != nil {
		v := *o.ExtendedLog

		if err = d.Set("extended_log", v); err != nil {
			return diag.Errorf("error reading extended_log: %v", err)
		}
	}

	if o.External != nil {
		v := *o.External

		if err = d.Set("external", v); err != nil {
			return diag.Errorf("error reading external: %v", err)
		}
	}

	if o.Method != nil {
		if err = d.Set("method", flattenWafProfileMethod(o.Method, sort)); err != nil {
			return diag.Errorf("error reading method: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Signature != nil {
		if err = d.Set("signature", flattenWafProfileSignature(o.Signature, sort)); err != nil {
			return diag.Errorf("error reading signature: %v", err)
		}
	}

	if o.UrlAccess != nil {
		if err = d.Set("url_access", flattenWafProfileUrlAccess(o.UrlAccess, sort)); err != nil {
			return diag.Errorf("error reading url_access: %v", err)
		}
	}

	return nil
}

func expandWafProfileAddressList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileAddressList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileAddressList

	for i := range l {
		tmp := models.WafProfileAddressList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.blocked_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileAddressListBlockedAddress(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileAddressListBlockedAddress
			// 	}
			tmp.BlockedAddress = v2

		}

		pre_append = fmt.Sprintf("%s.%d.blocked_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockedLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trusted_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileAddressListTrustedAddress(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileAddressListTrustedAddress
			// 	}
			tmp.TrustedAddress = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileAddressListBlockedAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileAddressListBlockedAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileAddressListBlockedAddress

	for i := range l {
		tmp := models.WafProfileAddressListBlockedAddress{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileAddressListTrustedAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileAddressListTrustedAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileAddressListTrustedAddress

	for i := range l {
		tmp := models.WafProfileAddressListTrustedAddress{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraint(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraint, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraint

	for i := range l {
		tmp := models.WafProfileConstraint{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.content_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintContentLength(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintContentLength
			// 	}
			tmp.ContentLength = v2

		}

		pre_append = fmt.Sprintf("%s.%d.exception", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintException(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintException
			// 	}
			tmp.Exception = v2

		}

		pre_append = fmt.Sprintf("%s.%d.header_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintHeaderLength(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintHeaderLength
			// 	}
			tmp.HeaderLength = v2

		}

		pre_append = fmt.Sprintf("%s.%d.hostname", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintHostname(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintHostname
			// 	}
			tmp.Hostname = v2

		}

		pre_append = fmt.Sprintf("%s.%d.line_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintLineLength(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintLineLength
			// 	}
			tmp.LineLength = v2

		}

		pre_append = fmt.Sprintf("%s.%d.malformed", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintMalformed(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintMalformed
			// 	}
			tmp.Malformed = v2

		}

		pre_append = fmt.Sprintf("%s.%d.max_cookie", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintMaxCookie(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintMaxCookie
			// 	}
			tmp.MaxCookie = v2

		}

		pre_append = fmt.Sprintf("%s.%d.max_header_line", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintMaxHeaderLine(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintMaxHeaderLine
			// 	}
			tmp.MaxHeaderLine = v2

		}

		pre_append = fmt.Sprintf("%s.%d.max_range_segment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintMaxRangeSegment(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintMaxRangeSegment
			// 	}
			tmp.MaxRangeSegment = v2

		}

		pre_append = fmt.Sprintf("%s.%d.max_url_param", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintMaxUrlParam(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintMaxUrlParam
			// 	}
			tmp.MaxUrlParam = v2

		}

		pre_append = fmt.Sprintf("%s.%d.method", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintMethod(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintMethod
			// 	}
			tmp.Method = v2

		}

		pre_append = fmt.Sprintf("%s.%d.param_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintParamLength(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintParamLength
			// 	}
			tmp.ParamLength = v2

		}

		pre_append = fmt.Sprintf("%s.%d.url_param_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintUrlParamLength(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintUrlParamLength
			// 	}
			tmp.UrlParamLength = v2

		}

		pre_append = fmt.Sprintf("%s.%d.version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileConstraintVersion(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileConstraintVersion
			// 	}
			tmp.Version = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintContentLength(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintContentLength, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintContentLength

	for i := range l {
		tmp := models.WafProfileConstraintContentLength{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Length = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintException(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintException, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintException

	for i := range l {
		tmp := models.WafProfileConstraintException{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Address = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.content_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ContentLength = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HeaderLength = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hostname", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Hostname = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.line_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LineLength = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Malformed = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_cookie", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MaxCookie = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_header_line", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MaxHeaderLine = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_range_segment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MaxRangeSegment = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_url_param", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MaxUrlParam = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.method", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Method = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.param_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ParamLength = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Pattern = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.regex", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Regex = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.url_param_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UrlParamLength = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Version = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintHeaderLength(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintHeaderLength, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintHeaderLength

	for i := range l {
		tmp := models.WafProfileConstraintHeaderLength{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Length = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintHostname, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintHostname

	for i := range l {
		tmp := models.WafProfileConstraintHostname{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintLineLength(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintLineLength, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintLineLength

	for i := range l {
		tmp := models.WafProfileConstraintLineLength{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Length = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintMalformed(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintMalformed, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintMalformed

	for i := range l {
		tmp := models.WafProfileConstraintMalformed{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintMaxCookie(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintMaxCookie, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintMaxCookie

	for i := range l {
		tmp := models.WafProfileConstraintMaxCookie{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_cookie", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxCookie = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintMaxHeaderLine(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintMaxHeaderLine, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintMaxHeaderLine

	for i := range l {
		tmp := models.WafProfileConstraintMaxHeaderLine{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_header_line", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxHeaderLine = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintMaxRangeSegment(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintMaxRangeSegment, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintMaxRangeSegment

	for i := range l {
		tmp := models.WafProfileConstraintMaxRangeSegment{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_range_segment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxRangeSegment = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintMaxUrlParam(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintMaxUrlParam, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintMaxUrlParam

	for i := range l {
		tmp := models.WafProfileConstraintMaxUrlParam{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_url_param", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxUrlParam = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintMethod, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintMethod

	for i := range l {
		tmp := models.WafProfileConstraintMethod{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintParamLength(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintParamLength, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintParamLength

	for i := range l {
		tmp := models.WafProfileConstraintParamLength{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Length = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintUrlParamLength(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintUrlParamLength, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintUrlParamLength

	for i := range l {
		tmp := models.WafProfileConstraintUrlParamLength{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Length = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileConstraintVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileConstraintVersion, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileConstraintVersion

	for i := range l {
		tmp := models.WafProfileConstraintVersion{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileMethod, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileMethod

	for i := range l {
		tmp := models.WafProfileMethod{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.default_allowed_methods", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DefaultAllowedMethods = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.method_policy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileMethodMethodPolicy(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileMethodMethodPolicy
			// 	}
			tmp.MethodPolicy = v2

		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileMethodMethodPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileMethodMethodPolicy, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileMethodMethodPolicy

	for i := range l {
		tmp := models.WafProfileMethodMethodPolicy{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Address = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.allowed_methods", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AllowedMethods = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Pattern = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.regex", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Regex = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileSignature(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileSignature, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileSignature

	for i := range l {
		tmp := models.WafProfileSignature{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.credit_card_detection_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.CreditCardDetectionThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.custom_signature", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileSignatureCustomSignature(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileSignatureCustomSignature
			// 	}
			tmp.CustomSignature = v2

		}

		pre_append = fmt.Sprintf("%s.%d.disabled_signature", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileSignatureDisabledSignature(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileSignatureDisabledSignature
			// 	}
			tmp.DisabledSignature = v2

		}

		pre_append = fmt.Sprintf("%s.%d.disabled_sub_class", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileSignatureDisabledSubClass(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileSignatureDisabledSubClass
			// 	}
			tmp.DisabledSubClass = v2

		}

		pre_append = fmt.Sprintf("%s.%d.main_class", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileSignatureMainClass(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileSignatureMainClass
			// 	}
			tmp.MainClass = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileSignatureCustomSignature(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileSignatureCustomSignature, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileSignatureCustomSignature

	for i := range l {
		tmp := models.WafProfileSignatureCustomSignature{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.case_sensitivity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CaseSensitivity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.direction", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Direction = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Pattern = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Target = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileSignatureDisabledSignature(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileSignatureDisabledSignature, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileSignatureDisabledSignature

	for i := range l {
		tmp := models.WafProfileSignatureDisabledSignature{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileSignatureDisabledSubClass(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileSignatureDisabledSubClass, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileSignatureDisabledSubClass

	for i := range l {
		tmp := models.WafProfileSignatureDisabledSubClass{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileSignatureMainClass(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileSignatureMainClass, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileSignatureMainClass

	for i := range l {
		tmp := models.WafProfileSignatureMainClass{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileUrlAccess(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileUrlAccess, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileUrlAccess

	for i := range l {
		tmp := models.WafProfileUrlAccess{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.access_pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWafProfileUrlAccessAccessPattern(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WafProfileUrlAccessAccessPattern
			// 	}
			tmp.AccessPattern = v2

		}

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Address = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWafProfileUrlAccessAccessPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WafProfileUrlAccessAccessPattern, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WafProfileUrlAccessAccessPattern

	for i := range l {
		tmp := models.WafProfileUrlAccessAccessPattern{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.negate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Negate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Pattern = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.regex", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Regex = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.srcaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Srcaddr = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWafProfile(d *schema.ResourceData, sv string) (*models.WafProfile, diag.Diagnostics) {
	obj := models.WafProfile{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("address_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("address_list", sv)
			diags = append(diags, e)
		}
		t, err := expandWafProfileAddressList(d, v, "address_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AddressList = t
		}
	} else if d.HasChange("address_list") {
		old, new := d.GetChange("address_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AddressList = &[]models.WafProfileAddressList{}
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v, ok := d.GetOk("constraint"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("constraint", sv)
			diags = append(diags, e)
		}
		t, err := expandWafProfileConstraint(d, v, "constraint", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Constraint = t
		}
	} else if d.HasChange("constraint") {
		old, new := d.GetChange("constraint")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Constraint = &[]models.WafProfileConstraint{}
		}
	}
	if v1, ok := d.GetOk("extended_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extended_log", sv)
				diags = append(diags, e)
			}
			obj.ExtendedLog = &v2
		}
	}
	if v1, ok := d.GetOk("external"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("external", sv)
				diags = append(diags, e)
			}
			obj.External = &v2
		}
	}
	if v, ok := d.GetOk("method"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("method", sv)
			diags = append(diags, e)
		}
		t, err := expandWafProfileMethod(d, v, "method", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Method = t
		}
	} else if d.HasChange("method") {
		old, new := d.GetChange("method")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Method = &[]models.WafProfileMethod{}
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v, ok := d.GetOk("signature"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("signature", sv)
			diags = append(diags, e)
		}
		t, err := expandWafProfileSignature(d, v, "signature", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Signature = t
		}
	} else if d.HasChange("signature") {
		old, new := d.GetChange("signature")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Signature = &[]models.WafProfileSignature{}
		}
	}
	if v, ok := d.GetOk("url_access"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("url_access", sv)
			diags = append(diags, e)
		}
		t, err := expandWafProfileUrlAccess(d, v, "url_access", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.UrlAccess = t
		}
	} else if d.HasChange("url_access") {
		old, new := d.GetChange("url_access")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.UrlAccess = &[]models.WafProfileUrlAccess{}
		}
	}
	return &obj, diags
}
