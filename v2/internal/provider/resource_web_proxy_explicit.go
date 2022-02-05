// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure explicit Web proxy settings.

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

func resourceWebProxyExplicit() *schema.Resource {
	return &schema.Resource{
		Description: "Configure explicit Web proxy settings.",

		CreateContext: resourceWebProxyExplicitCreate,
		ReadContext:   resourceWebProxyExplicitRead,
		UpdateContext: resourceWebProxyExplicitUpdate,
		DeleteContext: resourceWebProxyExplicitDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"ftp_incoming_port": {
				Type: schema.TypeString,

				Description: "Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).",
				Optional:    true,
				Computed:    true,
			},
			"ftp_over_http": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to proxy FTP-over-HTTP sessions sent from a web browser.",
				Optional:    true,
				Computed:    true,
			},
			"http_incoming_port": {
				Type: schema.TypeString,

				Description: "Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).",
				Optional:    true,
				Computed:    true,
			},
			"https_incoming_port": {
				Type: schema.TypeString,

				Description: "Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).",
				Optional:    true,
				Computed:    true,
			},
			"https_replacement_message": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sending the client a replacement message for HTTPS requests.",
				Optional:    true,
				Computed:    true,
			},
			"incoming_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.",
				Optional:    true,
				Computed:    true,
			},
			"incoming_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"ipv6_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command.",
				Optional:    true,
				Computed:    true,
			},
			"message_upon_server_error": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable displaying a replacement message when a server error is detected.",
				Optional:    true,
				Computed:    true,
			},
			"outgoing_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.",
				Optional:    true,
				Computed:    true,
			},
			"outgoing_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.",
				Optional:         true,
				Computed:         true,
			},
			"pac_file_data": {
				Type: schema.TypeString,

				Description: "PAC file contents enclosed in quotes (maximum of 256K bytes).",
				Optional:    true,
				Computed:    true,
			},
			"pac_file_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Pac file name.",
				Optional:    true,
				Computed:    true,
			},
			"pac_file_server_port": {
				Type: schema.TypeString,

				Description: "Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).",
				Optional:    true,
				Computed:    true,
			},
			"pac_file_server_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile.",
				Optional:    true,
				Computed:    true,
			},
			"pac_file_url": {
				Type: schema.TypeString,

				Description: "PAC file access URL.",
				Optional:    true,
				Computed:    true,
			},
			"pac_policy": {
				Type:        schema.TypeList,
				Description: "PAC policies.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comments": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1023),

							Description: "Optional comments.",
							Optional:    true,
							Computed:    true,
						},
						"dstaddr": {
							Type:        schema.TypeList,
							Description: "Destination address objects.",
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
						"pac_file_data": {
							Type: schema.TypeString,

							Description: "PAC file contents enclosed in quotes (maximum of 256K bytes).",
							Optional:    true,
							Computed:    true,
						},
						"pac_file_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Pac file name.",
							Optional:    true,
							Computed:    true,
						},
						"policyid": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Policy ID.",
							Optional:    true,
							Computed:    true,
						},
						"srcaddr": {
							Type:        schema.TypeList,
							Description: "Source address objects.",
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
						"srcaddr6": {
							Type:        schema.TypeList,
							Description: "Source address6 objects.",
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
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable policy.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"pref_dns_result": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),

				Description: "Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4).",
				Optional:    true,
				Computed:    true,
			},
			"realm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Authentication realm used to identify the explicit web proxy (maximum of 63 characters).",
				Optional:    true,
				Computed:    true,
			},
			"sec_default_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"accept", "deny"}, false),

				Description: "Accept or deny explicit web proxy sessions when no web proxy firewall policy exists.",
				Optional:    true,
				Computed:    true,
			},
			"socks": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the SOCKS proxy.",
				Optional:    true,
				Computed:    true,
			},
			"socks_incoming_port": {
				Type: schema.TypeString,

				Description: "Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).",
				Optional:    true,
				Computed:    true,
			},
			"ssl_algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

				Description: "Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the explicit Web proxy for HTTP and HTTPS session.",
				Optional:    true,
				Computed:    true,
			},
			"strict_guest": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable strict guest user checking by the explicit web proxy.",
				Optional:    true,
				Computed:    true,
			},
			"trace_auth_no_rsp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging timed-out authentication requests.",
				Optional:    true,
				Computed:    true,
			},
			"unknown_http_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"reject", "best-effort"}, false),

				Description: "How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWebProxyExplicitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebProxyExplicit(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWebProxyExplicit(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebProxyExplicit")
	}

	return resourceWebProxyExplicitRead(ctx, d, meta)
}

func resourceWebProxyExplicitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebProxyExplicit(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWebProxyExplicit(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WebProxyExplicit resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebProxyExplicit")
	}

	return resourceWebProxyExplicitRead(ctx, d, meta)
}

func resourceWebProxyExplicitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectWebProxyExplicit(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateWebProxyExplicit(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WebProxyExplicit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyExplicitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebProxyExplicit(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebProxyExplicit resource: %v", err)
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

	diags := refreshObjectWebProxyExplicit(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWebProxyExplicitPacPolicy(d *schema.ResourceData, v *[]models.WebProxyExplicitPacPolicy, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Comments; tmp != nil {
				v["comments"] = *tmp
			}

			if tmp := cfg.Dstaddr; tmp != nil {
				v["dstaddr"] = flattenWebProxyExplicitPacPolicyDstaddr(d, tmp, prefix+"dstaddr", sort)
			}

			if tmp := cfg.PacFileData; tmp != nil {
				v["pac_file_data"] = *tmp
			}

			if tmp := cfg.PacFileName; tmp != nil {
				v["pac_file_name"] = *tmp
			}

			if tmp := cfg.Policyid; tmp != nil {
				v["policyid"] = *tmp
			}

			if tmp := cfg.Srcaddr; tmp != nil {
				v["srcaddr"] = flattenWebProxyExplicitPacPolicySrcaddr(d, tmp, prefix+"srcaddr", sort)
			}

			if tmp := cfg.Srcaddr6; tmp != nil {
				v["srcaddr6"] = flattenWebProxyExplicitPacPolicySrcaddr6(d, tmp, prefix+"srcaddr6", sort)
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "policyid")
	}

	return flat
}

func flattenWebProxyExplicitPacPolicyDstaddr(d *schema.ResourceData, v *[]models.WebProxyExplicitPacPolicyDstaddr, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func flattenWebProxyExplicitPacPolicySrcaddr(d *schema.ResourceData, v *[]models.WebProxyExplicitPacPolicySrcaddr, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func flattenWebProxyExplicitPacPolicySrcaddr6(d *schema.ResourceData, v *[]models.WebProxyExplicitPacPolicySrcaddr6, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func refreshObjectWebProxyExplicit(d *schema.ResourceData, o *models.WebProxyExplicit, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.FtpIncomingPort != nil {
		v := *o.FtpIncomingPort

		if err = d.Set("ftp_incoming_port", v); err != nil {
			return diag.Errorf("error reading ftp_incoming_port: %v", err)
		}
	}

	if o.FtpOverHttp != nil {
		v := *o.FtpOverHttp

		if err = d.Set("ftp_over_http", v); err != nil {
			return diag.Errorf("error reading ftp_over_http: %v", err)
		}
	}

	if o.HttpIncomingPort != nil {
		v := *o.HttpIncomingPort

		if err = d.Set("http_incoming_port", v); err != nil {
			return diag.Errorf("error reading http_incoming_port: %v", err)
		}
	}

	if o.HttpsIncomingPort != nil {
		v := *o.HttpsIncomingPort

		if err = d.Set("https_incoming_port", v); err != nil {
			return diag.Errorf("error reading https_incoming_port: %v", err)
		}
	}

	if o.HttpsReplacementMessage != nil {
		v := *o.HttpsReplacementMessage

		if err = d.Set("https_replacement_message", v); err != nil {
			return diag.Errorf("error reading https_replacement_message: %v", err)
		}
	}

	if o.IncomingIp != nil {
		v := *o.IncomingIp

		if err = d.Set("incoming_ip", v); err != nil {
			return diag.Errorf("error reading incoming_ip: %v", err)
		}
	}

	if o.IncomingIp6 != nil {
		v := *o.IncomingIp6

		if err = d.Set("incoming_ip6", v); err != nil {
			return diag.Errorf("error reading incoming_ip6: %v", err)
		}
	}

	if o.Ipv6Status != nil {
		v := *o.Ipv6Status

		if err = d.Set("ipv6_status", v); err != nil {
			return diag.Errorf("error reading ipv6_status: %v", err)
		}
	}

	if o.MessageUponServerError != nil {
		v := *o.MessageUponServerError

		if err = d.Set("message_upon_server_error", v); err != nil {
			return diag.Errorf("error reading message_upon_server_error: %v", err)
		}
	}

	if o.OutgoingIp != nil {
		v := *o.OutgoingIp

		if err = d.Set("outgoing_ip", v); err != nil {
			return diag.Errorf("error reading outgoing_ip: %v", err)
		}
	}

	if o.OutgoingIp6 != nil {
		v := *o.OutgoingIp6

		if err = d.Set("outgoing_ip6", v); err != nil {
			return diag.Errorf("error reading outgoing_ip6: %v", err)
		}
	}

	if o.PacFileData != nil {
		v := *o.PacFileData

		if err = d.Set("pac_file_data", v); err != nil {
			return diag.Errorf("error reading pac_file_data: %v", err)
		}
	}

	if o.PacFileName != nil {
		v := *o.PacFileName

		if err = d.Set("pac_file_name", v); err != nil {
			return diag.Errorf("error reading pac_file_name: %v", err)
		}
	}

	if o.PacFileServerPort != nil {
		v := *o.PacFileServerPort

		if err = d.Set("pac_file_server_port", v); err != nil {
			return diag.Errorf("error reading pac_file_server_port: %v", err)
		}
	}

	if o.PacFileServerStatus != nil {
		v := *o.PacFileServerStatus

		if err = d.Set("pac_file_server_status", v); err != nil {
			return diag.Errorf("error reading pac_file_server_status: %v", err)
		}
	}

	if o.PacFileUrl != nil {
		v := *o.PacFileUrl

		if err = d.Set("pac_file_url", v); err != nil {
			return diag.Errorf("error reading pac_file_url: %v", err)
		}
	}

	if o.PacPolicy != nil {
		if err = d.Set("pac_policy", flattenWebProxyExplicitPacPolicy(d, o.PacPolicy, "pac_policy", sort)); err != nil {
			return diag.Errorf("error reading pac_policy: %v", err)
		}
	}

	if o.PrefDnsResult != nil {
		v := *o.PrefDnsResult

		if err = d.Set("pref_dns_result", v); err != nil {
			return diag.Errorf("error reading pref_dns_result: %v", err)
		}
	}

	if o.Realm != nil {
		v := *o.Realm

		if err = d.Set("realm", v); err != nil {
			return diag.Errorf("error reading realm: %v", err)
		}
	}

	if o.SecDefaultAction != nil {
		v := *o.SecDefaultAction

		if err = d.Set("sec_default_action", v); err != nil {
			return diag.Errorf("error reading sec_default_action: %v", err)
		}
	}

	if o.Socks != nil {
		v := *o.Socks

		if err = d.Set("socks", v); err != nil {
			return diag.Errorf("error reading socks: %v", err)
		}
	}

	if o.SocksIncomingPort != nil {
		v := *o.SocksIncomingPort

		if err = d.Set("socks_incoming_port", v); err != nil {
			return diag.Errorf("error reading socks_incoming_port: %v", err)
		}
	}

	if o.SslAlgorithm != nil {
		v := *o.SslAlgorithm

		if err = d.Set("ssl_algorithm", v); err != nil {
			return diag.Errorf("error reading ssl_algorithm: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.StrictGuest != nil {
		v := *o.StrictGuest

		if err = d.Set("strict_guest", v); err != nil {
			return diag.Errorf("error reading strict_guest: %v", err)
		}
	}

	if o.TraceAuthNoRsp != nil {
		v := *o.TraceAuthNoRsp

		if err = d.Set("trace_auth_no_rsp", v); err != nil {
			return diag.Errorf("error reading trace_auth_no_rsp: %v", err)
		}
	}

	if o.UnknownHttpVersion != nil {
		v := *o.UnknownHttpVersion

		if err = d.Set("unknown_http_version", v); err != nil {
			return diag.Errorf("error reading unknown_http_version: %v", err)
		}
	}

	return nil
}

func expandWebProxyExplicitPacPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebProxyExplicitPacPolicy, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebProxyExplicitPacPolicy

	for i := range l {
		tmp := models.WebProxyExplicitPacPolicy{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.comments", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comments = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dstaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebProxyExplicitPacPolicyDstaddr(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebProxyExplicitPacPolicyDstaddr
			// 	}
			tmp.Dstaddr = v2

		}

		pre_append = fmt.Sprintf("%s.%d.pac_file_data", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PacFileData = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pac_file_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PacFileName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.policyid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Policyid = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.srcaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebProxyExplicitPacPolicySrcaddr(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebProxyExplicitPacPolicySrcaddr
			// 	}
			tmp.Srcaddr = v2

		}

		pre_append = fmt.Sprintf("%s.%d.srcaddr6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebProxyExplicitPacPolicySrcaddr6(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebProxyExplicitPacPolicySrcaddr6
			// 	}
			tmp.Srcaddr6 = v2

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

func expandWebProxyExplicitPacPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebProxyExplicitPacPolicyDstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebProxyExplicitPacPolicyDstaddr

	for i := range l {
		tmp := models.WebProxyExplicitPacPolicyDstaddr{}
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

func expandWebProxyExplicitPacPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebProxyExplicitPacPolicySrcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebProxyExplicitPacPolicySrcaddr

	for i := range l {
		tmp := models.WebProxyExplicitPacPolicySrcaddr{}
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

func expandWebProxyExplicitPacPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebProxyExplicitPacPolicySrcaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebProxyExplicitPacPolicySrcaddr6

	for i := range l {
		tmp := models.WebProxyExplicitPacPolicySrcaddr6{}
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

func getObjectWebProxyExplicit(d *schema.ResourceData, sv string) (*models.WebProxyExplicit, diag.Diagnostics) {
	obj := models.WebProxyExplicit{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("ftp_incoming_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ftp_incoming_port", sv)
				diags = append(diags, e)
			}
			obj.FtpIncomingPort = &v2
		}
	}
	if v1, ok := d.GetOk("ftp_over_http"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ftp_over_http", sv)
				diags = append(diags, e)
			}
			obj.FtpOverHttp = &v2
		}
	}
	if v1, ok := d.GetOk("http_incoming_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_incoming_port", sv)
				diags = append(diags, e)
			}
			obj.HttpIncomingPort = &v2
		}
	}
	if v1, ok := d.GetOk("https_incoming_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("https_incoming_port", sv)
				diags = append(diags, e)
			}
			obj.HttpsIncomingPort = &v2
		}
	}
	if v1, ok := d.GetOk("https_replacement_message"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("https_replacement_message", sv)
				diags = append(diags, e)
			}
			obj.HttpsReplacementMessage = &v2
		}
	}
	if v1, ok := d.GetOk("incoming_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("incoming_ip", sv)
				diags = append(diags, e)
			}
			obj.IncomingIp = &v2
		}
	}
	if v1, ok := d.GetOk("incoming_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("incoming_ip6", sv)
				diags = append(diags, e)
			}
			obj.IncomingIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_status", sv)
				diags = append(diags, e)
			}
			obj.Ipv6Status = &v2
		}
	}
	if v1, ok := d.GetOk("message_upon_server_error"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("message_upon_server_error", sv)
				diags = append(diags, e)
			}
			obj.MessageUponServerError = &v2
		}
	}
	if v1, ok := d.GetOk("outgoing_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outgoing_ip", sv)
				diags = append(diags, e)
			}
			obj.OutgoingIp = &v2
		}
	}
	if v1, ok := d.GetOk("outgoing_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outgoing_ip6", sv)
				diags = append(diags, e)
			}
			obj.OutgoingIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("pac_file_data"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pac_file_data", sv)
				diags = append(diags, e)
			}
			obj.PacFileData = &v2
		}
	}
	if v1, ok := d.GetOk("pac_file_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pac_file_name", sv)
				diags = append(diags, e)
			}
			obj.PacFileName = &v2
		}
	}
	if v1, ok := d.GetOk("pac_file_server_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pac_file_server_port", sv)
				diags = append(diags, e)
			}
			obj.PacFileServerPort = &v2
		}
	}
	if v1, ok := d.GetOk("pac_file_server_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pac_file_server_status", sv)
				diags = append(diags, e)
			}
			obj.PacFileServerStatus = &v2
		}
	}
	if v1, ok := d.GetOk("pac_file_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pac_file_url", sv)
				diags = append(diags, e)
			}
			obj.PacFileUrl = &v2
		}
	}
	if v, ok := d.GetOk("pac_policy"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("pac_policy", sv)
			diags = append(diags, e)
		}
		t, err := expandWebProxyExplicitPacPolicy(d, v, "pac_policy", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PacPolicy = t
		}
	} else if d.HasChange("pac_policy") {
		old, new := d.GetChange("pac_policy")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PacPolicy = &[]models.WebProxyExplicitPacPolicy{}
		}
	}
	if v1, ok := d.GetOk("pref_dns_result"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pref_dns_result", sv)
				diags = append(diags, e)
			}
			obj.PrefDnsResult = &v2
		}
	}
	if v1, ok := d.GetOk("realm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("realm", sv)
				diags = append(diags, e)
			}
			obj.Realm = &v2
		}
	}
	if v1, ok := d.GetOk("sec_default_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sec_default_action", sv)
				diags = append(diags, e)
			}
			obj.SecDefaultAction = &v2
		}
	}
	if v1, ok := d.GetOk("socks"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("socks", sv)
				diags = append(diags, e)
			}
			obj.Socks = &v2
		}
	}
	if v1, ok := d.GetOk("socks_incoming_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("socks_incoming_port", sv)
				diags = append(diags, e)
			}
			obj.SocksIncomingPort = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_algorithm", sv)
				diags = append(diags, e)
			}
			obj.SslAlgorithm = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("strict_guest"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("strict_guest", sv)
				diags = append(diags, e)
			}
			obj.StrictGuest = &v2
		}
	}
	if v1, ok := d.GetOk("trace_auth_no_rsp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trace_auth_no_rsp", sv)
				diags = append(diags, e)
			}
			obj.TraceAuthNoRsp = &v2
		}
	}
	if v1, ok := d.GetOk("unknown_http_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("unknown_http_version", sv)
				diags = append(diags, e)
			}
			obj.UnknownHttpVersion = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectWebProxyExplicit(d *schema.ResourceData, sv string) (*models.WebProxyExplicit, diag.Diagnostics) {
	obj := models.WebProxyExplicit{}
	diags := diag.Diagnostics{}

	obj.PacPolicy = &[]models.WebProxyExplicitPacPolicy{}

	return &obj, diags
}
