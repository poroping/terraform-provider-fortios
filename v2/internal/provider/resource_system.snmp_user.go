// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SNMP user configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemSnmpUser() *schema.Resource {
	return &schema.Resource{
		Description: "SNMP user configuration.",

		CreateContext: resourceSystemSnmpUserCreate,
		ReadContext:   resourceSystemSnmpUserRead,
		UpdateContext: resourceSystemSnmpUserUpdate,
		DeleteContext: resourceSystemSnmpUserDelete,

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
			"auth_proto": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"md5", "sha", "sha224", "sha256", "sha384", "sha512"}, false),

				Description: "Authentication protocol.",
				Optional:    true,
				Computed:    true,
			},
			"auth_pwd": {
				Type: schema.TypeString,

				Description: "Password for authentication protocol.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"events": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "SNMP notifications (traps) to send.",
				Optional:         true,
				Computed:         true,
			},
			"ha_direct": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable direct management of HA cluster members.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),

				Description: "SNMP user name.",
				ForceNew:    true,
				Required:    true,
			},
			"notify_hosts": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "SNMP managers to send notifications (traps) to.",
				Optional:    true,
				Computed:    true,
			},
			"notify_hosts6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 SNMP managers to send notifications (traps) to.",
				Optional:         true,
				Computed:         true,
			},
			"priv_proto": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"aes", "des", "aes256", "aes256cisco"}, false),

				Description: "Privacy (encryption) protocol.",
				Optional:    true,
				Computed:    true,
			},
			"priv_pwd": {
				Type: schema.TypeString,

				Description: "Password for privacy (encryption) protocol.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"queries": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SNMP queries for this user.",
				Optional:    true,
				Computed:    true,
			},
			"query_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "SNMPv3 query port (default = 161).",
				Optional:    true,
				Computed:    true,
			},
			"security_level": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"no-auth-no-priv", "auth-no-priv", "auth-priv"}, false),

				Description: "Security level for message authentication and encryption.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP for SNMP trap.",
				Optional:    true,
				Computed:    true,
			},
			"source_ipv6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Source IPv6 for SNMP trap.",
				Optional:         true,
				Computed:         true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this SNMP user.",
				Optional:    true,
				Computed:    true,
			},
			"trap_lport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "SNMPv3 local trap port (default = 162).",
				Optional:    true,
				Computed:    true,
			},
			"trap_rport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "SNMPv3 trap remote port (default = 162).",
				Optional:    true,
				Computed:    true,
			},
			"trap_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable traps for this SNMP user.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemSnmpUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemSnmpUser resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemSnmpUser(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSnmpUser(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSnmpUser")
	}

	return resourceSystemSnmpUserRead(ctx, d, meta)
}

func resourceSystemSnmpUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSnmpUser(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSnmpUser(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSnmpUser resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSnmpUser")
	}

	return resourceSystemSnmpUserRead(ctx, d, meta)
}

func resourceSystemSnmpUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSnmpUser(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSnmpUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSnmpUser(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSnmpUser resource: %v", err)
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

	diags := refreshObjectSystemSnmpUser(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemSnmpUser(d *schema.ResourceData, o *models.SystemSnmpUser, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthProto != nil {
		v := *o.AuthProto

		if err = d.Set("auth_proto", v); err != nil {
			return diag.Errorf("error reading auth_proto: %v", err)
		}
	}

	if o.AuthPwd != nil {
		v := *o.AuthPwd

		if err = d.Set("auth_pwd", v); err != nil {
			return diag.Errorf("error reading auth_pwd: %v", err)
		}
	}

	if o.Events != nil {
		v := *o.Events

		if err = d.Set("events", v); err != nil {
			return diag.Errorf("error reading events: %v", err)
		}
	}

	if o.HaDirect != nil {
		v := *o.HaDirect

		if err = d.Set("ha_direct", v); err != nil {
			return diag.Errorf("error reading ha_direct: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.NotifyHosts != nil {
		v := *o.NotifyHosts

		if err = d.Set("notify_hosts", v); err != nil {
			return diag.Errorf("error reading notify_hosts: %v", err)
		}
	}

	if o.NotifyHosts6 != nil {
		v := *o.NotifyHosts6

		if err = d.Set("notify_hosts6", v); err != nil {
			return diag.Errorf("error reading notify_hosts6: %v", err)
		}
	}

	if o.PrivProto != nil {
		v := *o.PrivProto

		if err = d.Set("priv_proto", v); err != nil {
			return diag.Errorf("error reading priv_proto: %v", err)
		}
	}

	if o.PrivPwd != nil {
		v := *o.PrivPwd

		if err = d.Set("priv_pwd", v); err != nil {
			return diag.Errorf("error reading priv_pwd: %v", err)
		}
	}

	if o.Queries != nil {
		v := *o.Queries

		if err = d.Set("queries", v); err != nil {
			return diag.Errorf("error reading queries: %v", err)
		}
	}

	if o.QueryPort != nil {
		v := *o.QueryPort

		if err = d.Set("query_port", v); err != nil {
			return diag.Errorf("error reading query_port: %v", err)
		}
	}

	if o.SecurityLevel != nil {
		v := *o.SecurityLevel

		if err = d.Set("security_level", v); err != nil {
			return diag.Errorf("error reading security_level: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.SourceIpv6 != nil {
		v := *o.SourceIpv6

		if err = d.Set("source_ipv6", v); err != nil {
			return diag.Errorf("error reading source_ipv6: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TrapLport != nil {
		v := *o.TrapLport

		if err = d.Set("trap_lport", v); err != nil {
			return diag.Errorf("error reading trap_lport: %v", err)
		}
	}

	if o.TrapRport != nil {
		v := *o.TrapRport

		if err = d.Set("trap_rport", v); err != nil {
			return diag.Errorf("error reading trap_rport: %v", err)
		}
	}

	if o.TrapStatus != nil {
		v := *o.TrapStatus

		if err = d.Set("trap_status", v); err != nil {
			return diag.Errorf("error reading trap_status: %v", err)
		}
	}

	return nil
}

func getObjectSystemSnmpUser(d *schema.ResourceData, sv string) (*models.SystemSnmpUser, diag.Diagnostics) {
	obj := models.SystemSnmpUser{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth_proto"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_proto", sv)
				diags = append(diags, e)
			}
			obj.AuthProto = &v2
		}
	}
	if v1, ok := d.GetOk("auth_pwd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_pwd", sv)
				diags = append(diags, e)
			}
			obj.AuthPwd = &v2
		}
	}
	if v1, ok := d.GetOk("events"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("events", sv)
				diags = append(diags, e)
			}
			obj.Events = &v2
		}
	}
	if v1, ok := d.GetOk("ha_direct"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ha_direct", sv)
				diags = append(diags, e)
			}
			obj.HaDirect = &v2
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
	if v1, ok := d.GetOk("notify_hosts"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("notify_hosts", sv)
				diags = append(diags, e)
			}
			obj.NotifyHosts = &v2
		}
	}
	if v1, ok := d.GetOk("notify_hosts6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("notify_hosts6", sv)
				diags = append(diags, e)
			}
			obj.NotifyHosts6 = &v2
		}
	}
	if v1, ok := d.GetOk("priv_proto"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priv_proto", sv)
				diags = append(diags, e)
			}
			obj.PrivProto = &v2
		}
	}
	if v1, ok := d.GetOk("priv_pwd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priv_pwd", sv)
				diags = append(diags, e)
			}
			obj.PrivPwd = &v2
		}
	}
	if v1, ok := d.GetOk("queries"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("queries", sv)
				diags = append(diags, e)
			}
			obj.Queries = &v2
		}
	}
	if v1, ok := d.GetOk("query_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("query_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.QueryPort = &tmp
		}
	}
	if v1, ok := d.GetOk("security_level"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_level", sv)
				diags = append(diags, e)
			}
			obj.SecurityLevel = &v2
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
	if v1, ok := d.GetOk("source_ipv6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ipv6", sv)
				diags = append(diags, e)
			}
			obj.SourceIpv6 = &v2
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
	if v1, ok := d.GetOk("trap_lport"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_lport", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TrapLport = &tmp
		}
	}
	if v1, ok := d.GetOk("trap_rport"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_rport", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TrapRport = &tmp
		}
	}
	if v1, ok := d.GetOk("trap_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_status", sv)
				diags = append(diags, e)
			}
			obj.TrapStatus = &v2
		}
	}
	return &obj, diags
}
