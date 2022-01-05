// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure MS Exchange server entries.

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

func resourceUserExchange() *schema.Resource {
	return &schema.Resource{
		Description: "Configure MS Exchange server entries.",

		CreateContext: resourceUserExchangeCreate,
		ReadContext:   resourceUserExchangeRead,
		UpdateContext: resourceUserExchangeUpdate,
		DeleteContext: resourceUserExchangeDelete,

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
			"auth_level": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"connect", "call", "packet", "integrity", "privacy"}, false),

				Description: "Authentication security level used for the RPC protocol layer.",
				Optional:    true,
				Computed:    true,
			},
			"auth_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"spnego", "ntlm", "kerberos"}, false),

				Description: "Authentication security type used for the RPC protocol layer.",
				Optional:    true,
				Computed:    true,
			},
			"auto_discover_kdc": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic discovery of KDC IP addresses.",
				Optional:    true,
				Computed:    true,
			},
			"connect_protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"rpc-over-tcp", "rpc-over-http", "rpc-over-https"}, false),

				Description: "Connection protocol used to connect to MS Exchange service.",
				Optional:    true,
				Computed:    true,
			},
			"domain_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "MS Exchange server fully qualified domain name.",
				Optional:    true,
				Computed:    true,
			},
			"http_auth_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"basic", "ntlm"}, false),

				Description: "Authentication security type used for the HTTP transport.",
				Optional:    true,
				Computed:    true,
			},
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Server IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"kdc_ip": {
				Type:        schema.TypeList,
				Description: "KDC IPv4 addresses for Kerberos authentication.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "KDC IPv4 addresses for Kerberos authentication.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "MS Exchange server entry name.",
				ForceNew:    true,
				Required:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Password for the specified username.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"server_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "MS Exchange server hostname.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "SSLv3", "TLSv1", "TLSv1-1", "TLSv1-2"}, false),

				Description: "Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting).",
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "User name used to sign in to the server. Must have proper permissions for service.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserExchangeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserExchange resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserExchange(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserExchange(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserExchange")
	}

	return resourceUserExchangeRead(ctx, d, meta)
}

func resourceUserExchangeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserExchange(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserExchange(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserExchange resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserExchange")
	}

	return resourceUserExchangeRead(ctx, d, meta)
}

func resourceUserExchangeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserExchange(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserExchange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserExchangeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserExchange(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserExchange resource: %v", err)
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

	diags := refreshObjectUserExchange(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenUserExchangeKdcIp(v *[]models.UserExchangeKdcIp, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Ipv4; tmp != nil {
				v["ipv4"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "ipv4")
	}

	return flat
}

func refreshObjectUserExchange(d *schema.ResourceData, o *models.UserExchange, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthLevel != nil {
		v := *o.AuthLevel

		if err = d.Set("auth_level", v); err != nil {
			return diag.Errorf("error reading auth_level: %v", err)
		}
	}

	if o.AuthType != nil {
		v := *o.AuthType

		if err = d.Set("auth_type", v); err != nil {
			return diag.Errorf("error reading auth_type: %v", err)
		}
	}

	if o.AutoDiscoverKdc != nil {
		v := *o.AutoDiscoverKdc

		if err = d.Set("auto_discover_kdc", v); err != nil {
			return diag.Errorf("error reading auto_discover_kdc: %v", err)
		}
	}

	if o.ConnectProtocol != nil {
		v := *o.ConnectProtocol

		if err = d.Set("connect_protocol", v); err != nil {
			return diag.Errorf("error reading connect_protocol: %v", err)
		}
	}

	if o.DomainName != nil {
		v := *o.DomainName

		if err = d.Set("domain_name", v); err != nil {
			return diag.Errorf("error reading domain_name: %v", err)
		}
	}

	if o.HttpAuthType != nil {
		v := *o.HttpAuthType

		if err = d.Set("http_auth_type", v); err != nil {
			return diag.Errorf("error reading http_auth_type: %v", err)
		}
	}

	if o.Ip != nil {
		v := *o.Ip

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.KdcIp != nil {
		if err = d.Set("kdc_ip", flattenUserExchangeKdcIp(o.KdcIp, sort)); err != nil {
			return diag.Errorf("error reading kdc_ip: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.ServerName != nil {
		v := *o.ServerName

		if err = d.Set("server_name", v); err != nil {
			return diag.Errorf("error reading server_name: %v", err)
		}
	}

	if o.SslMinProtoVersion != nil {
		v := *o.SslMinProtoVersion

		if err = d.Set("ssl_min_proto_version", v); err != nil {
			return diag.Errorf("error reading ssl_min_proto_version: %v", err)
		}
	}

	if o.Username != nil {
		v := *o.Username

		if err = d.Set("username", v); err != nil {
			return diag.Errorf("error reading username: %v", err)
		}
	}

	return nil
}

func expandUserExchangeKdcIp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserExchangeKdcIp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserExchangeKdcIp

	for i := range l {
		tmp := models.UserExchangeKdcIp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ipv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectUserExchange(d *schema.ResourceData, sv string) (*models.UserExchange, diag.Diagnostics) {
	obj := models.UserExchange{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth_level"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_level", sv)
				diags = append(diags, e)
			}
			obj.AuthLevel = &v2
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
	if v1, ok := d.GetOk("auto_discover_kdc"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("auto_discover_kdc", sv)
				diags = append(diags, e)
			}
			obj.AutoDiscoverKdc = &v2
		}
	}
	if v1, ok := d.GetOk("connect_protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("connect_protocol", sv)
				diags = append(diags, e)
			}
			obj.ConnectProtocol = &v2
		}
	}
	if v1, ok := d.GetOk("domain_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("domain_name", sv)
				diags = append(diags, e)
			}
			obj.DomainName = &v2
		}
	}
	if v1, ok := d.GetOk("http_auth_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_auth_type", sv)
				diags = append(diags, e)
			}
			obj.HttpAuthType = &v2
		}
	}
	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
		}
	}
	if v, ok := d.GetOk("kdc_ip"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("kdc_ip", sv)
			diags = append(diags, e)
		}
		t, err := expandUserExchangeKdcIp(d, v, "kdc_ip", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.KdcIp = t
		}
	} else if d.HasChange("kdc_ip") {
		old, new := d.GetChange("kdc_ip")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.KdcIp = &[]models.UserExchangeKdcIp{}
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
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("server_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_name", sv)
				diags = append(diags, e)
			}
			obj.ServerName = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_min_proto_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_min_proto_version", sv)
				diags = append(diags, e)
			}
			obj.SslMinProtoVersion = &v2
		}
	}
	if v1, ok := d.GetOk("username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("username", sv)
				diags = append(diags, e)
			}
			obj.Username = &v2
		}
	}
	return &obj, diags
}
