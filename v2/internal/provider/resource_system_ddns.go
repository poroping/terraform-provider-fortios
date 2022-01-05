// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DDNS.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemDdns() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DDNS.",

		CreateContext: resourceSystemDdnsCreate,
		ReadContext:   resourceSystemDdnsRead,
		UpdateContext: resourceSystemDdnsUpdate,
		DeleteContext: resourceSystemDdnsDelete,

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
			"addr_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),

				Description: "Address type of interface address in DDNS update.",
				Optional:    true,
				Computed:    true,
			},
			"bound_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 46),

				Description: "Bound IP address.",
				Optional:    true,
				Computed:    true,
			},
			"clear_text": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable use of clear text connections.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "tsig"}, false),

				Description: "Enable/disable TSIG authentication for your DDNS server.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "Your fully qualified domain name (for example, yourname.DDNS.com).",
				Optional:    true,
				Computed:    true,
			},
			"ddns_key": {
				Type: schema.TypeString,

				Description: "DDNS update key (base 64 encoding).",
				Optional:    true,
				Computed:    true,
			},
			"ddns_keyname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "DDNS update key name.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_password": {
				Type: schema.TypeString,

				Description: "DDNS password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"ddns_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"dyndns.org", "dyns.net", "tzo.com", "vavic.com", "dipdns.net", "now.net.cn", "dhs.org", "easydns.com", "genericDDNS", "FortiGuardDDNS", "noip.com"}, false),

				Description: "Select a DDNS service provider.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_server_addr": {
				Type:        schema.TypeList,
				Description: "Generic DDNS server IP/FQDN list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 256),

							Description: "IP address or FQDN of the server.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ddns_server_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Generic DDNS server IP.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_sn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "DDNS Serial Number.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),

				Description: "Time-to-live for DDNS packets.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "DDNS user name.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_zone": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "Zone of your domain name (for example, DDNS.com).",
				Optional:    true,
				Computed:    true,
			},
			"ddnsid": {
				Type: schema.TypeInt,

				Description: "DDNS ID.",
				ForceNew:    true,
				Required:    true,
			},
			"monitor_interface": {
				Type:        schema.TypeList,
				Description: "Monitored interface.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"server_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),

				Description: "Address type of the DDNS server.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of local certificate for SSL connections.",
				Optional:    true,
				Computed:    true,
			},
			"update_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 2592000),

				Description: "DDNS update interval (60 - 2592000 sec, default = 300).",
				Optional:    true,
				Computed:    true,
			},
			"use_public_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable use of public IP address.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemDdnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "ddnsid"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemDdns resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemDdns(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemDdns(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemDdns")
	}

	return resourceSystemDdnsRead(ctx, d, meta)
}

func resourceSystemDdnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemDdns(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemDdns(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemDdns resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemDdns")
	}

	return resourceSystemDdnsRead(ctx, d, meta)
}

func resourceSystemDdnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemDdns(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemDdns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDdnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemDdns(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemDdns resource: %v", err)
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

	diags := refreshObjectSystemDdns(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemDdnsDdnsServerAddr(v *[]models.SystemDdnsDdnsServerAddr, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Addr; tmp != nil {
				v["addr"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "addr")
	}

	return flat
}

func flattenSystemDdnsMonitorInterface(v *[]models.SystemDdnsMonitorInterface, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.InterfaceName; tmp != nil {
				v["interface_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "interface_name")
	}

	return flat
}

func refreshObjectSystemDdns(d *schema.ResourceData, o *models.SystemDdns, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddrType != nil {
		v := *o.AddrType

		if err = d.Set("addr_type", v); err != nil {
			return diag.Errorf("error reading addr_type: %v", err)
		}
	}

	if o.BoundIp != nil {
		v := *o.BoundIp

		if err = d.Set("bound_ip", v); err != nil {
			return diag.Errorf("error reading bound_ip: %v", err)
		}
	}

	if o.ClearText != nil {
		v := *o.ClearText

		if err = d.Set("clear_text", v); err != nil {
			return diag.Errorf("error reading clear_text: %v", err)
		}
	}

	if o.DdnsAuth != nil {
		v := *o.DdnsAuth

		if err = d.Set("ddns_auth", v); err != nil {
			return diag.Errorf("error reading ddns_auth: %v", err)
		}
	}

	if o.DdnsDomain != nil {
		v := *o.DdnsDomain

		if err = d.Set("ddns_domain", v); err != nil {
			return diag.Errorf("error reading ddns_domain: %v", err)
		}
	}

	if o.DdnsKey != nil {
		v := *o.DdnsKey

		if err = d.Set("ddns_key", v); err != nil {
			return diag.Errorf("error reading ddns_key: %v", err)
		}
	}

	if o.DdnsKeyname != nil {
		v := *o.DdnsKeyname

		if err = d.Set("ddns_keyname", v); err != nil {
			return diag.Errorf("error reading ddns_keyname: %v", err)
		}
	}

	if o.DdnsPassword != nil {
		v := *o.DdnsPassword

		if err = d.Set("ddns_password", v); err != nil {
			return diag.Errorf("error reading ddns_password: %v", err)
		}
	}

	if o.DdnsServer != nil {
		v := *o.DdnsServer

		if err = d.Set("ddns_server", v); err != nil {
			return diag.Errorf("error reading ddns_server: %v", err)
		}
	}

	if o.DdnsServerAddr != nil {
		if err = d.Set("ddns_server_addr", flattenSystemDdnsDdnsServerAddr(o.DdnsServerAddr, sort)); err != nil {
			return diag.Errorf("error reading ddns_server_addr: %v", err)
		}
	}

	if o.DdnsServerIp != nil {
		v := *o.DdnsServerIp

		if err = d.Set("ddns_server_ip", v); err != nil {
			return diag.Errorf("error reading ddns_server_ip: %v", err)
		}
	}

	if o.DdnsSn != nil {
		v := *o.DdnsSn

		if err = d.Set("ddns_sn", v); err != nil {
			return diag.Errorf("error reading ddns_sn: %v", err)
		}
	}

	if o.DdnsTtl != nil {
		v := *o.DdnsTtl

		if err = d.Set("ddns_ttl", v); err != nil {
			return diag.Errorf("error reading ddns_ttl: %v", err)
		}
	}

	if o.DdnsUsername != nil {
		v := *o.DdnsUsername

		if err = d.Set("ddns_username", v); err != nil {
			return diag.Errorf("error reading ddns_username: %v", err)
		}
	}

	if o.DdnsZone != nil {
		v := *o.DdnsZone

		if err = d.Set("ddns_zone", v); err != nil {
			return diag.Errorf("error reading ddns_zone: %v", err)
		}
	}

	if o.Ddnsid != nil {
		v := *o.Ddnsid

		if err = d.Set("ddnsid", v); err != nil {
			return diag.Errorf("error reading ddnsid: %v", err)
		}
	}

	if o.MonitorInterface != nil {
		if err = d.Set("monitor_interface", flattenSystemDdnsMonitorInterface(o.MonitorInterface, sort)); err != nil {
			return diag.Errorf("error reading monitor_interface: %v", err)
		}
	}

	if o.ServerType != nil {
		v := *o.ServerType

		if err = d.Set("server_type", v); err != nil {
			return diag.Errorf("error reading server_type: %v", err)
		}
	}

	if o.SslCertificate != nil {
		v := *o.SslCertificate

		if err = d.Set("ssl_certificate", v); err != nil {
			return diag.Errorf("error reading ssl_certificate: %v", err)
		}
	}

	if o.UpdateInterval != nil {
		v := *o.UpdateInterval

		if err = d.Set("update_interval", v); err != nil {
			return diag.Errorf("error reading update_interval: %v", err)
		}
	}

	if o.UsePublicIp != nil {
		v := *o.UsePublicIp

		if err = d.Set("use_public_ip", v); err != nil {
			return diag.Errorf("error reading use_public_ip: %v", err)
		}
	}

	return nil
}

func expandSystemDdnsDdnsServerAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDdnsDdnsServerAddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDdnsDdnsServerAddr

	for i := range l {
		tmp := models.SystemDdnsDdnsServerAddr{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.addr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Addr = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDdnsMonitorInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDdnsMonitorInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDdnsMonitorInterface

	for i := range l {
		tmp := models.SystemDdnsMonitorInterface{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterfaceName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemDdns(d *schema.ResourceData, sv string) (*models.SystemDdns, diag.Diagnostics) {
	obj := models.SystemDdns{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("addr_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.5", "") {
				e := utils.AttributeVersionWarning("addr_type", sv)
				diags = append(diags, e)
			}
			obj.AddrType = &v2
		}
	}
	if v1, ok := d.GetOk("bound_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bound_ip", sv)
				diags = append(diags, e)
			}
			obj.BoundIp = &v2
		}
	}
	if v1, ok := d.GetOk("clear_text"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("clear_text", sv)
				diags = append(diags, e)
			}
			obj.ClearText = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_auth", sv)
				diags = append(diags, e)
			}
			obj.DdnsAuth = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_domain", sv)
				diags = append(diags, e)
			}
			obj.DdnsDomain = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_key", sv)
				diags = append(diags, e)
			}
			obj.DdnsKey = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_keyname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_keyname", sv)
				diags = append(diags, e)
			}
			obj.DdnsKeyname = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_password", sv)
				diags = append(diags, e)
			}
			obj.DdnsPassword = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_server", sv)
				diags = append(diags, e)
			}
			obj.DdnsServer = &v2
		}
	}
	if v, ok := d.GetOk("ddns_server_addr"); ok {
		if !utils.CheckVer(sv, "v6.4.5", "") {
			e := utils.AttributeVersionWarning("ddns_server_addr", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemDdnsDdnsServerAddr(d, v, "ddns_server_addr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DdnsServerAddr = t
		}
	} else if d.HasChange("ddns_server_addr") {
		old, new := d.GetChange("ddns_server_addr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DdnsServerAddr = &[]models.SystemDdnsDdnsServerAddr{}
		}
	}
	if v1, ok := d.GetOk("ddns_server_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.5") {
				e := utils.AttributeVersionWarning("ddns_server_ip", sv)
				diags = append(diags, e)
			}
			obj.DdnsServerIp = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_sn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_sn", sv)
				diags = append(diags, e)
			}
			obj.DdnsSn = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DdnsTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("ddns_username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_username", sv)
				diags = append(diags, e)
			}
			obj.DdnsUsername = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_zone"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_zone", sv)
				diags = append(diags, e)
			}
			obj.DdnsZone = &v2
		}
	}
	if v1, ok := d.GetOk("ddnsid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddnsid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Ddnsid = &tmp
		}
	}
	if v, ok := d.GetOk("monitor_interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("monitor_interface", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemDdnsMonitorInterface(d, v, "monitor_interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.MonitorInterface = t
		}
	} else if d.HasChange("monitor_interface") {
		old, new := d.GetChange("monitor_interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.MonitorInterface = &[]models.SystemDdnsMonitorInterface{}
		}
	}
	if v1, ok := d.GetOk("server_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.5", "") {
				e := utils.AttributeVersionWarning("server_type", sv)
				diags = append(diags, e)
			}
			obj.ServerType = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_certificate", sv)
				diags = append(diags, e)
			}
			obj.SslCertificate = &v2
		}
	}
	if v1, ok := d.GetOk("update_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UpdateInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("use_public_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("use_public_ip", sv)
				diags = append(diags, e)
			}
			obj.UsePublicIp = &v2
		}
	}
	return &obj, diags
}
