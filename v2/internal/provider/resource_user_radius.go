// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure RADIUS server entries.

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

func resourceUserRadius() *schema.Resource {
	return &schema.Resource{
		Description: "Configure RADIUS server entries.",

		CreateContext: resourceUserRadiusCreate,
		ReadContext:   resourceUserRadiusRead,
		UpdateContext: resourceUserRadiusUpdate,
		DeleteContext: resourceUserRadiusDelete,

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
			"accounting_server": {
				Type:        schema.TypeList,
				Description: "Additional accounting servers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "ID (0 - 4294967295).",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Specify outgoing interface to reach server.",
							Optional:    true,
							Computed:    true,
						},
						"interface_select_method": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

							Description: "Specify how to select outgoing interface to reach server.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "RADIUS accounting port number.",
							Optional:    true,
							Computed:    true,
						},
						"secret": {
							Type: schema.TypeString,

							Description: "Secret key.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Server CN domain name or IP address.",
							Optional:    true,
							Computed:    true,
						},
						"source_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Source IP address for communications to the RADIUS server.",
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
			"acct_all_servers": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sending of accounting messages to all configured servers (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"acct_interim_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),

				Description: "Time in seconds between each accounting interim update message.",
				Optional:    true,
				Computed:    true,
			},
			"all_usergroup": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable automatically including this RADIUS server in all user groups.",
				Optional:    true,
				Computed:    true,
			},
			"auth_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "ms_chap_v2", "ms_chap", "chap", "pap"}, false),

				Description: "Authentication methods/protocols permitted for this RADIUS server.",
				Optional:    true,
				Computed:    true,
			},
			"class": {
				Type:        schema.TypeList,
				Description: "Class attribute name(s).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Class name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"group_override_attr_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"filter-Id", "class"}, false),

				Description: "RADIUS attribute type to override user group information.",
				Optional:    true,
				Computed:    true,
			},
			"h3c_compatibility": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable compatibility with the H3C, a mechanism that performs security checking for authentication.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Specify outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"interface_select_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

				Description: "Specify how to select outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "RADIUS server entry name.",
				ForceNew:    true,
				Required:    true,
			},
			"nas_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address used to communicate with the RADIUS server and used as NAS-IP-Address and Called-Station-ID attributes.",
				Optional:    true,
				Computed:    true,
			},
			"password_encoding": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "ISO-8859-1"}, false),

				Description: "Password encoding.",
				Optional:    true,
				Computed:    true,
			},
			"password_renewal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable password renewal.",
				Optional:    true,
				Computed:    true,
			},
			"radius_coa": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to allow a mechanism to change the attributes of an authentication, authorization, and accounting session after it is authenticated.",
				Optional:    true,
				Computed:    true,
			},
			"radius_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "RADIUS service port number.",
				Optional:    true,
				Computed:    true,
			},
			"rsso": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable RADIUS based single sign on feature.",
				Optional:    true,
				Computed:    true,
			},
			"rsso_context_timeout": {
				Type: schema.TypeInt,

				Description: "Time in seconds before the logged out user is removed from the \"user context list\" of logged on users.",
				Optional:    true,
				Computed:    true,
			},
			"rsso_endpoint_attribute": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"User-Name", "NAS-IP-Address", "Framed-IP-Address", "Framed-IP-Netmask", "Filter-Id", "Login-IP-Host", "Reply-Message", "Callback-Number", "Callback-Id", "Framed-Route", "Framed-IPX-Network", "Class", "Called-Station-Id", "Calling-Station-Id", "NAS-Identifier", "Proxy-State", "Login-LAT-Service", "Login-LAT-Node", "Login-LAT-Group", "Framed-AppleTalk-Zone", "Acct-Session-Id", "Acct-Multi-Session-Id"}, false),

				Description: "RADIUS attributes used to extract the user end point identifier from the RADIUS Start record.",
				Optional:    true,
				Computed:    true,
			},
			"rsso_endpoint_block_attribute": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"User-Name", "NAS-IP-Address", "Framed-IP-Address", "Framed-IP-Netmask", "Filter-Id", "Login-IP-Host", "Reply-Message", "Callback-Number", "Callback-Id", "Framed-Route", "Framed-IPX-Network", "Class", "Called-Station-Id", "Calling-Station-Id", "NAS-Identifier", "Proxy-State", "Login-LAT-Service", "Login-LAT-Node", "Login-LAT-Group", "Framed-AppleTalk-Zone", "Acct-Session-Id", "Acct-Multi-Session-Id"}, false),

				Description: "RADIUS attributes used to block a user.",
				Optional:    true,
				Computed:    true,
			},
			"rsso_ep_one_ip_only": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the replacement of old IP addresses with new ones for the same endpoint on RADIUS accounting Start messages.",
				Optional:    true,
				Computed:    true,
			},
			"rsso_flush_ip_session": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable flushing user IP sessions on RADIUS accounting Stop messages.",
				Optional:    true,
				Computed:    true,
			},
			"rsso_log_flags": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Events to log.",
				Optional:         true,
				Computed:         true,
			},
			"rsso_log_period": {
				Type: schema.TypeInt,

				Description: "Time interval in seconds that group event log messages will be generated for dynamic profile events.",
				Optional:    true,
				Computed:    true,
			},
			"rsso_radius_response": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sending RADIUS response packets after receiving Start and Stop records.",
				Optional:    true,
				Computed:    true,
			},
			"rsso_radius_server_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "UDP port to listen on for RADIUS Start and Stop records.",
				Optional:    true,
				Computed:    true,
			},
			"rsso_secret": {
				Type: schema.TypeString,

				Description: "RADIUS secret used by the RADIUS accounting server.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"rsso_validate_request_secret": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable validating the RADIUS request shared secret in the Start or End record.",
				Optional:    true,
				Computed:    true,
			},
			"secondary_secret": {
				Type: schema.TypeString,

				Description: "Secret key to access the secondary server.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"secondary_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Secondary RADIUS CN domain name or IP address.",
				Optional:    true,
				Computed:    true,
			},
			"secret": {
				Type: schema.TypeString,

				Description: "Pre-shared secret key used to access the primary RADIUS server.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Primary RADIUS server CN domain name or IP address.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Source IP address for communications to the RADIUS server.",
				Optional:    true,
				Computed:    true,
			},
			"sso_attribute": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"User-Name", "NAS-IP-Address", "Framed-IP-Address", "Framed-IP-Netmask", "Filter-Id", "Login-IP-Host", "Reply-Message", "Callback-Number", "Callback-Id", "Framed-Route", "Framed-IPX-Network", "Class", "Called-Station-Id", "Calling-Station-Id", "NAS-Identifier", "Proxy-State", "Login-LAT-Service", "Login-LAT-Node", "Login-LAT-Group", "Framed-AppleTalk-Zone", "Acct-Session-Id", "Acct-Multi-Session-Id"}, false),

				Description: "RADIUS attribute that contains the profile group name to be extracted from the RADIUS Start record.",
				Optional:    true,
				Computed:    true,
			},
			"sso_attribute_key": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Key prefix for SSO group value in the SSO attribute.",
				Optional:    true,
				Computed:    true,
			},
			"sso_attribute_value_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable override old attribute value with new value for the same endpoint.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_acct_fast_framedip_detect": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 600),

				Description: "Switch controller accounting message Framed-IP detection from DHCP snooping (seconds, default=2).",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_service_type": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "RADIUS service type.",
				Optional:         true,
				Computed:         true,
			},
			"tertiary_secret": {
				Type: schema.TypeString,

				Description: "Secret key to access the tertiary server.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"tertiary_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Tertiary RADIUS CN domain name or IP address.",
				Optional:    true,
				Computed:    true,
			},
			"timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),

				Description: "Time in seconds between re-sending authentication requests.",
				Optional:    true,
				Computed:    true,
			},
			"use_management_vdom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable using management VDOM to send requests.",
				Optional:    true,
				Computed:    true,
			},
			"username_case_sensitive": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable case sensitive user names.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserRadiusCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserRadius resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserRadius(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserRadius(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserRadius")
	}

	return resourceUserRadiusRead(ctx, d, meta)
}

func resourceUserRadiusUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserRadius(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserRadius(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserRadius resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserRadius")
	}

	return resourceUserRadiusRead(ctx, d, meta)
}

func resourceUserRadiusDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserRadius(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserRadius resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserRadiusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserRadius(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserRadius resource: %v", err)
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

	diags := refreshObjectUserRadius(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenUserRadiusAccountingServer(d *schema.ResourceData, v *[]models.UserRadiusAccountingServer, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.InterfaceSelectMethod; tmp != nil {
				v["interface_select_method"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.Secret; tmp != nil {
				v["secret"] = *tmp
			}

			if tmp := cfg.Server; tmp != nil {
				v["server"] = *tmp
			}

			if tmp := cfg.SourceIp; tmp != nil {
				v["source_ip"] = *tmp
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

func flattenUserRadiusClass(d *schema.ResourceData, v *[]models.UserRadiusClass, prefix string, sort bool) interface{} {
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

func refreshObjectUserRadius(d *schema.ResourceData, o *models.UserRadius, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AccountingServer != nil {
		if err = d.Set("accounting_server", flattenUserRadiusAccountingServer(d, o.AccountingServer, "accounting_server", sort)); err != nil {
			return diag.Errorf("error reading accounting_server: %v", err)
		}
	}

	if o.AcctAllServers != nil {
		v := *o.AcctAllServers

		if err = d.Set("acct_all_servers", v); err != nil {
			return diag.Errorf("error reading acct_all_servers: %v", err)
		}
	}

	if o.AcctInterimInterval != nil {
		v := *o.AcctInterimInterval

		if err = d.Set("acct_interim_interval", v); err != nil {
			return diag.Errorf("error reading acct_interim_interval: %v", err)
		}
	}

	if o.AllUsergroup != nil {
		v := *o.AllUsergroup

		if err = d.Set("all_usergroup", v); err != nil {
			return diag.Errorf("error reading all_usergroup: %v", err)
		}
	}

	if o.AuthType != nil {
		v := *o.AuthType

		if err = d.Set("auth_type", v); err != nil {
			return diag.Errorf("error reading auth_type: %v", err)
		}
	}

	if o.Class != nil {
		if err = d.Set("class", flattenUserRadiusClass(d, o.Class, "class", sort)); err != nil {
			return diag.Errorf("error reading class: %v", err)
		}
	}

	if o.GroupOverrideAttrType != nil {
		v := *o.GroupOverrideAttrType

		if err = d.Set("group_override_attr_type", v); err != nil {
			return diag.Errorf("error reading group_override_attr_type: %v", err)
		}
	}

	if o.H3cCompatibility != nil {
		v := *o.H3cCompatibility

		if err = d.Set("h3c_compatibility", v); err != nil {
			return diag.Errorf("error reading h3c_compatibility: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.InterfaceSelectMethod != nil {
		v := *o.InterfaceSelectMethod

		if err = d.Set("interface_select_method", v); err != nil {
			return diag.Errorf("error reading interface_select_method: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.NasIp != nil {
		v := *o.NasIp

		if err = d.Set("nas_ip", v); err != nil {
			return diag.Errorf("error reading nas_ip: %v", err)
		}
	}

	if o.PasswordEncoding != nil {
		v := *o.PasswordEncoding

		if err = d.Set("password_encoding", v); err != nil {
			return diag.Errorf("error reading password_encoding: %v", err)
		}
	}

	if o.PasswordRenewal != nil {
		v := *o.PasswordRenewal

		if err = d.Set("password_renewal", v); err != nil {
			return diag.Errorf("error reading password_renewal: %v", err)
		}
	}

	if o.RadiusCoa != nil {
		v := *o.RadiusCoa

		if err = d.Set("radius_coa", v); err != nil {
			return diag.Errorf("error reading radius_coa: %v", err)
		}
	}

	if o.RadiusPort != nil {
		v := *o.RadiusPort

		if err = d.Set("radius_port", v); err != nil {
			return diag.Errorf("error reading radius_port: %v", err)
		}
	}

	if o.Rsso != nil {
		v := *o.Rsso

		if err = d.Set("rsso", v); err != nil {
			return diag.Errorf("error reading rsso: %v", err)
		}
	}

	if o.RssoContextTimeout != nil {
		v := *o.RssoContextTimeout

		if err = d.Set("rsso_context_timeout", v); err != nil {
			return diag.Errorf("error reading rsso_context_timeout: %v", err)
		}
	}

	if o.RssoEndpointAttribute != nil {
		v := *o.RssoEndpointAttribute

		if err = d.Set("rsso_endpoint_attribute", v); err != nil {
			return diag.Errorf("error reading rsso_endpoint_attribute: %v", err)
		}
	}

	if o.RssoEndpointBlockAttribute != nil {
		v := *o.RssoEndpointBlockAttribute

		if err = d.Set("rsso_endpoint_block_attribute", v); err != nil {
			return diag.Errorf("error reading rsso_endpoint_block_attribute: %v", err)
		}
	}

	if o.RssoEpOneIpOnly != nil {
		v := *o.RssoEpOneIpOnly

		if err = d.Set("rsso_ep_one_ip_only", v); err != nil {
			return diag.Errorf("error reading rsso_ep_one_ip_only: %v", err)
		}
	}

	if o.RssoFlushIpSession != nil {
		v := *o.RssoFlushIpSession

		if err = d.Set("rsso_flush_ip_session", v); err != nil {
			return diag.Errorf("error reading rsso_flush_ip_session: %v", err)
		}
	}

	if o.RssoLogFlags != nil {
		v := *o.RssoLogFlags

		if err = d.Set("rsso_log_flags", v); err != nil {
			return diag.Errorf("error reading rsso_log_flags: %v", err)
		}
	}

	if o.RssoLogPeriod != nil {
		v := *o.RssoLogPeriod

		if err = d.Set("rsso_log_period", v); err != nil {
			return diag.Errorf("error reading rsso_log_period: %v", err)
		}
	}

	if o.RssoRadiusResponse != nil {
		v := *o.RssoRadiusResponse

		if err = d.Set("rsso_radius_response", v); err != nil {
			return diag.Errorf("error reading rsso_radius_response: %v", err)
		}
	}

	if o.RssoRadiusServerPort != nil {
		v := *o.RssoRadiusServerPort

		if err = d.Set("rsso_radius_server_port", v); err != nil {
			return diag.Errorf("error reading rsso_radius_server_port: %v", err)
		}
	}

	if o.RssoSecret != nil {
		v := *o.RssoSecret

		if v == "ENC XXXX" {
		} else if err = d.Set("rsso_secret", v); err != nil {
			return diag.Errorf("error reading rsso_secret: %v", err)
		}
	}

	if o.RssoValidateRequestSecret != nil {
		v := *o.RssoValidateRequestSecret

		if err = d.Set("rsso_validate_request_secret", v); err != nil {
			return diag.Errorf("error reading rsso_validate_request_secret: %v", err)
		}
	}

	if o.SecondarySecret != nil {
		v := *o.SecondarySecret

		if v == "ENC XXXX" {
		} else if err = d.Set("secondary_secret", v); err != nil {
			return diag.Errorf("error reading secondary_secret: %v", err)
		}
	}

	if o.SecondaryServer != nil {
		v := *o.SecondaryServer

		if err = d.Set("secondary_server", v); err != nil {
			return diag.Errorf("error reading secondary_server: %v", err)
		}
	}

	if o.Secret != nil {
		v := *o.Secret

		if v == "ENC XXXX" {
		} else if err = d.Set("secret", v); err != nil {
			return diag.Errorf("error reading secret: %v", err)
		}
	}

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.SsoAttribute != nil {
		v := *o.SsoAttribute

		if err = d.Set("sso_attribute", v); err != nil {
			return diag.Errorf("error reading sso_attribute: %v", err)
		}
	}

	if o.SsoAttributeKey != nil {
		v := *o.SsoAttributeKey

		if err = d.Set("sso_attribute_key", v); err != nil {
			return diag.Errorf("error reading sso_attribute_key: %v", err)
		}
	}

	if o.SsoAttributeValueOverride != nil {
		v := *o.SsoAttributeValueOverride

		if err = d.Set("sso_attribute_value_override", v); err != nil {
			return diag.Errorf("error reading sso_attribute_value_override: %v", err)
		}
	}

	if o.SwitchControllerAcctFastFramedipDetect != nil {
		v := *o.SwitchControllerAcctFastFramedipDetect

		if err = d.Set("switch_controller_acct_fast_framedip_detect", v); err != nil {
			return diag.Errorf("error reading switch_controller_acct_fast_framedip_detect: %v", err)
		}
	}

	if o.SwitchControllerServiceType != nil {
		v := *o.SwitchControllerServiceType

		if err = d.Set("switch_controller_service_type", v); err != nil {
			return diag.Errorf("error reading switch_controller_service_type: %v", err)
		}
	}

	if o.TertiarySecret != nil {
		v := *o.TertiarySecret

		if v == "ENC XXXX" {
		} else if err = d.Set("tertiary_secret", v); err != nil {
			return diag.Errorf("error reading tertiary_secret: %v", err)
		}
	}

	if o.TertiaryServer != nil {
		v := *o.TertiaryServer

		if err = d.Set("tertiary_server", v); err != nil {
			return diag.Errorf("error reading tertiary_server: %v", err)
		}
	}

	if o.Timeout != nil {
		v := *o.Timeout

		if err = d.Set("timeout", v); err != nil {
			return diag.Errorf("error reading timeout: %v", err)
		}
	}

	if o.UseManagementVdom != nil {
		v := *o.UseManagementVdom

		if err = d.Set("use_management_vdom", v); err != nil {
			return diag.Errorf("error reading use_management_vdom: %v", err)
		}
	}

	if o.UsernameCaseSensitive != nil {
		v := *o.UsernameCaseSensitive

		if err = d.Set("username_case_sensitive", v); err != nil {
			return diag.Errorf("error reading username_case_sensitive: %v", err)
		}
	}

	return nil
}

func expandUserRadiusAccountingServer(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserRadiusAccountingServer, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserRadiusAccountingServer

	for i := range l {
		tmp := models.UserRadiusAccountingServer{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface_select_method", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterfaceSelectMethod = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Port = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.secret", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Secret = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Server = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SourceIp = &v2
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

func expandUserRadiusClass(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserRadiusClass, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserRadiusClass

	for i := range l {
		tmp := models.UserRadiusClass{}
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

func getObjectUserRadius(d *schema.ResourceData, sv string) (*models.UserRadius, diag.Diagnostics) {
	obj := models.UserRadius{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("accounting_server"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("accounting_server", sv)
			diags = append(diags, e)
		}
		t, err := expandUserRadiusAccountingServer(d, v, "accounting_server", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AccountingServer = t
		}
	} else if d.HasChange("accounting_server") {
		old, new := d.GetChange("accounting_server")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AccountingServer = &[]models.UserRadiusAccountingServer{}
		}
	}
	if v1, ok := d.GetOk("acct_all_servers"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("acct_all_servers", sv)
				diags = append(diags, e)
			}
			obj.AcctAllServers = &v2
		}
	}
	if v1, ok := d.GetOk("acct_interim_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("acct_interim_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AcctInterimInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("all_usergroup"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("all_usergroup", sv)
				diags = append(diags, e)
			}
			obj.AllUsergroup = &v2
		}
	}
	if v1, ok := d.GetOk("auth_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_type", sv)
				diags = append(diags, e)
			}
			obj.AuthType = &v2
		}
	}
	if v, ok := d.GetOk("class"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("class", sv)
			diags = append(diags, e)
		}
		t, err := expandUserRadiusClass(d, v, "class", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Class = t
		}
	} else if d.HasChange("class") {
		old, new := d.GetChange("class")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Class = &[]models.UserRadiusClass{}
		}
	}
	if v1, ok := d.GetOk("group_override_attr_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("group_override_attr_type", sv)
				diags = append(diags, e)
			}
			obj.GroupOverrideAttrType = &v2
		}
	}
	if v1, ok := d.GetOk("h3c_compatibility"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("h3c_compatibility", sv)
				diags = append(diags, e)
			}
			obj.H3cCompatibility = &v2
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface_select_method", sv)
				diags = append(diags, e)
			}
			obj.InterfaceSelectMethod = &v2
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
	if v1, ok := d.GetOk("nas_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nas_ip", sv)
				diags = append(diags, e)
			}
			obj.NasIp = &v2
		}
	}
	if v1, ok := d.GetOk("password_encoding"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password_encoding", sv)
				diags = append(diags, e)
			}
			obj.PasswordEncoding = &v2
		}
	}
	if v1, ok := d.GetOk("password_renewal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password_renewal", sv)
				diags = append(diags, e)
			}
			obj.PasswordRenewal = &v2
		}
	}
	if v1, ok := d.GetOk("radius_coa"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radius_coa", sv)
				diags = append(diags, e)
			}
			obj.RadiusCoa = &v2
		}
	}
	if v1, ok := d.GetOk("radius_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radius_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RadiusPort = &tmp
		}
	}
	if v1, ok := d.GetOk("rsso"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsso", sv)
				diags = append(diags, e)
			}
			obj.Rsso = &v2
		}
	}
	if v1, ok := d.GetOk("rsso_context_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsso_context_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RssoContextTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("rsso_endpoint_attribute"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsso_endpoint_attribute", sv)
				diags = append(diags, e)
			}
			obj.RssoEndpointAttribute = &v2
		}
	}
	if v1, ok := d.GetOk("rsso_endpoint_block_attribute"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsso_endpoint_block_attribute", sv)
				diags = append(diags, e)
			}
			obj.RssoEndpointBlockAttribute = &v2
		}
	}
	if v1, ok := d.GetOk("rsso_ep_one_ip_only"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsso_ep_one_ip_only", sv)
				diags = append(diags, e)
			}
			obj.RssoEpOneIpOnly = &v2
		}
	}
	if v1, ok := d.GetOk("rsso_flush_ip_session"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsso_flush_ip_session", sv)
				diags = append(diags, e)
			}
			obj.RssoFlushIpSession = &v2
		}
	}
	if v1, ok := d.GetOk("rsso_log_flags"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsso_log_flags", sv)
				diags = append(diags, e)
			}
			obj.RssoLogFlags = &v2
		}
	}
	if v1, ok := d.GetOk("rsso_log_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsso_log_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RssoLogPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("rsso_radius_response"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsso_radius_response", sv)
				diags = append(diags, e)
			}
			obj.RssoRadiusResponse = &v2
		}
	}
	if v1, ok := d.GetOk("rsso_radius_server_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsso_radius_server_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RssoRadiusServerPort = &tmp
		}
	}
	if v1, ok := d.GetOk("rsso_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsso_secret", sv)
				diags = append(diags, e)
			}
			obj.RssoSecret = &v2
		}
	}
	if v1, ok := d.GetOk("rsso_validate_request_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsso_validate_request_secret", sv)
				diags = append(diags, e)
			}
			obj.RssoValidateRequestSecret = &v2
		}
	}
	if v1, ok := d.GetOk("secondary_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secondary_secret", sv)
				diags = append(diags, e)
			}
			obj.SecondarySecret = &v2
		}
	}
	if v1, ok := d.GetOk("secondary_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secondary_server", sv)
				diags = append(diags, e)
			}
			obj.SecondaryServer = &v2
		}
	}
	if v1, ok := d.GetOk("secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secret", sv)
				diags = append(diags, e)
			}
			obj.Secret = &v2
		}
	}
	if v1, ok := d.GetOk("server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server", sv)
				diags = append(diags, e)
			}
			obj.Server = &v2
		}
	}
	if v1, ok := d.GetOk("source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip", sv)
				diags = append(diags, e)
			}
			obj.SourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("sso_attribute"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sso_attribute", sv)
				diags = append(diags, e)
			}
			obj.SsoAttribute = &v2
		}
	}
	if v1, ok := d.GetOk("sso_attribute_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sso_attribute_key", sv)
				diags = append(diags, e)
			}
			obj.SsoAttributeKey = &v2
		}
	}
	if v1, ok := d.GetOk("sso_attribute_value_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sso_attribute_value_override", sv)
				diags = append(diags, e)
			}
			obj.SsoAttributeValueOverride = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_acct_fast_framedip_detect"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("switch_controller_acct_fast_framedip_detect", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SwitchControllerAcctFastFramedipDetect = &tmp
		}
	}
	if v1, ok := d.GetOk("switch_controller_service_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("switch_controller_service_type", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerServiceType = &v2
		}
	}
	if v1, ok := d.GetOk("tertiary_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tertiary_secret", sv)
				diags = append(diags, e)
			}
			obj.TertiarySecret = &v2
		}
	}
	if v1, ok := d.GetOk("tertiary_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tertiary_server", sv)
				diags = append(diags, e)
			}
			obj.TertiaryServer = &v2
		}
	}
	if v1, ok := d.GetOk("timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Timeout = &tmp
		}
	}
	if v1, ok := d.GetOk("use_management_vdom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("use_management_vdom", sv)
				diags = append(diags, e)
			}
			obj.UseManagementVdom = &v2
		}
	}
	if v1, ok := d.GetOk("username_case_sensitive"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("username_case_sensitive", sv)
				diags = append(diags, e)
			}
			obj.UsernameCaseSensitive = &v2
		}
	}
	return &obj, diags
}
