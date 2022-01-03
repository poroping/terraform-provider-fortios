// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Fortinet Single Sign On (FSSO) agents.

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

func resourceUserFsso() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Fortinet Single Sign On (FSSO) agents.",

		CreateContext: resourceUserFssoCreate,
		ReadContext:   resourceUserFssoRead,
		UpdateContext: resourceUserFssoUpdate,
		DeleteContext: resourceUserFssoDelete,

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
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"group_poll_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2880),

				Description: "Interval in minutes within to fetch groups from FSSO server, or unset to disable.",
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
			"ldap_poll": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic fetching of groups from LDAP server.",
				Optional:    true,
				Computed:    true,
			},
			"ldap_poll_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),

				Description: "Filter used to fetch groups.",
				Optional:    true,
				Computed:    true,
			},
			"ldap_poll_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2880),

				Description: "Interval in minutes within to fetch groups from LDAP server.",
				Optional:    true,
				Computed:    true,
			},
			"ldap_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "LDAP server to get group information.",
				Optional:    true,
				Computed:    true,
			},
			"logon_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2880),

				Description: "Interval in minutes to keep logons after FSSO server down.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Password of the first FSSO collector agent.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"password2": {
				Type: schema.TypeString,

				Description: "Password of the second FSSO collector agent.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"password3": {
				Type: schema.TypeString,

				Description: "Password of the third FSSO collector agent.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"password4": {
				Type: schema.TypeString,

				Description: "Password of the fourth FSSO collector agent.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"password5": {
				Type: schema.TypeString,

				Description: "Password of the fifth FSSO collector agent.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port of the first FSSO collector agent.",
				Optional:    true,
				Computed:    true,
			},
			"port2": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port of the second FSSO collector agent.",
				Optional:    true,
				Computed:    true,
			},
			"port3": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port of the third FSSO collector agent.",
				Optional:    true,
				Computed:    true,
			},
			"port4": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port of the fourth FSSO collector agent.",
				Optional:    true,
				Computed:    true,
			},
			"port5": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port of the fifth FSSO collector agent.",
				Optional:    true,
				Computed:    true,
			},
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Domain name or IP address of the first FSSO collector agent.",
				Optional:    true,
				Computed:    true,
			},
			"server2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Domain name or IP address of the second FSSO collector agent.",
				Optional:    true,
				Computed:    true,
			},
			"server3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Domain name or IP address of the third FSSO collector agent.",
				Optional:    true,
				Computed:    true,
			},
			"server4": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Domain name or IP address of the fourth FSSO collector agent.",
				Optional:    true,
				Computed:    true,
			},
			"server5": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Domain name or IP address of the fifth FSSO collector agent.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP for communications to FSSO agent.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 source for communications to FSSO agent.",
				Optional:         true,
				Computed:         true,
			},
			"ssl": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of SSL.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_server_host_ip_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable server host/IP verification.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_trusted_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Trusted server certificate or CA certificate.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "fortinac"}, false),

				Description: "Server type.",
				Optional:    true,
				Computed:    true,
			},
			"user_info_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "LDAP server to get user information.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserFssoCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserFsso resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserFsso(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserFsso(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserFsso")
	}

	return resourceUserFssoRead(ctx, d, meta)
}

func resourceUserFssoUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserFsso(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserFsso(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserFsso resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserFsso")
	}

	return resourceUserFssoRead(ctx, d, meta)
}

func resourceUserFssoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserFsso(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserFsso resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserFssoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserFsso(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserFsso resource: %v", err)
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

	diags := refreshObjectUserFsso(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectUserFsso(d *schema.ResourceData, o *models.UserFsso, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.GroupPollInterval != nil {
		v := *o.GroupPollInterval

		if err = d.Set("group_poll_interval", v); err != nil {
			return diag.Errorf("error reading group_poll_interval: %v", err)
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

	if o.LdapPoll != nil {
		v := *o.LdapPoll

		if err = d.Set("ldap_poll", v); err != nil {
			return diag.Errorf("error reading ldap_poll: %v", err)
		}
	}

	if o.LdapPollFilter != nil {
		v := *o.LdapPollFilter

		if err = d.Set("ldap_poll_filter", v); err != nil {
			return diag.Errorf("error reading ldap_poll_filter: %v", err)
		}
	}

	if o.LdapPollInterval != nil {
		v := *o.LdapPollInterval

		if err = d.Set("ldap_poll_interval", v); err != nil {
			return diag.Errorf("error reading ldap_poll_interval: %v", err)
		}
	}

	if o.LdapServer != nil {
		v := *o.LdapServer

		if err = d.Set("ldap_server", v); err != nil {
			return diag.Errorf("error reading ldap_server: %v", err)
		}
	}

	if o.LogonTimeout != nil {
		v := *o.LogonTimeout

		if err = d.Set("logon_timeout", v); err != nil {
			return diag.Errorf("error reading logon_timeout: %v", err)
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

		if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.Password2 != nil {
		v := *o.Password2

		if err = d.Set("password2", v); err != nil {
			return diag.Errorf("error reading password2: %v", err)
		}
	}

	if o.Password3 != nil {
		v := *o.Password3

		if err = d.Set("password3", v); err != nil {
			return diag.Errorf("error reading password3: %v", err)
		}
	}

	if o.Password4 != nil {
		v := *o.Password4

		if err = d.Set("password4", v); err != nil {
			return diag.Errorf("error reading password4: %v", err)
		}
	}

	if o.Password5 != nil {
		v := *o.Password5

		if err = d.Set("password5", v); err != nil {
			return diag.Errorf("error reading password5: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.Port2 != nil {
		v := *o.Port2

		if err = d.Set("port2", v); err != nil {
			return diag.Errorf("error reading port2: %v", err)
		}
	}

	if o.Port3 != nil {
		v := *o.Port3

		if err = d.Set("port3", v); err != nil {
			return diag.Errorf("error reading port3: %v", err)
		}
	}

	if o.Port4 != nil {
		v := *o.Port4

		if err = d.Set("port4", v); err != nil {
			return diag.Errorf("error reading port4: %v", err)
		}
	}

	if o.Port5 != nil {
		v := *o.Port5

		if err = d.Set("port5", v); err != nil {
			return diag.Errorf("error reading port5: %v", err)
		}
	}

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.Server2 != nil {
		v := *o.Server2

		if err = d.Set("server2", v); err != nil {
			return diag.Errorf("error reading server2: %v", err)
		}
	}

	if o.Server3 != nil {
		v := *o.Server3

		if err = d.Set("server3", v); err != nil {
			return diag.Errorf("error reading server3: %v", err)
		}
	}

	if o.Server4 != nil {
		v := *o.Server4

		if err = d.Set("server4", v); err != nil {
			return diag.Errorf("error reading server4: %v", err)
		}
	}

	if o.Server5 != nil {
		v := *o.Server5

		if err = d.Set("server5", v); err != nil {
			return diag.Errorf("error reading server5: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.SourceIp6 != nil {
		v := *o.SourceIp6

		if err = d.Set("source_ip6", v); err != nil {
			return diag.Errorf("error reading source_ip6: %v", err)
		}
	}

	if o.Ssl != nil {
		v := *o.Ssl

		if err = d.Set("ssl", v); err != nil {
			return diag.Errorf("error reading ssl: %v", err)
		}
	}

	if o.SslServerHostIpCheck != nil {
		v := *o.SslServerHostIpCheck

		if err = d.Set("ssl_server_host_ip_check", v); err != nil {
			return diag.Errorf("error reading ssl_server_host_ip_check: %v", err)
		}
	}

	if o.SslTrustedCert != nil {
		v := *o.SslTrustedCert

		if err = d.Set("ssl_trusted_cert", v); err != nil {
			return diag.Errorf("error reading ssl_trusted_cert: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.UserInfoServer != nil {
		v := *o.UserInfoServer

		if err = d.Set("user_info_server", v); err != nil {
			return diag.Errorf("error reading user_info_server: %v", err)
		}
	}

	return nil
}

func getObjectUserFsso(d *schema.ResourceData, sv string) (*models.UserFsso, diag.Diagnostics) {
	obj := models.UserFsso{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("group_poll_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_poll_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GroupPollInterval = &tmp
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
	if v1, ok := d.GetOk("ldap_poll"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldap_poll", sv)
				diags = append(diags, e)
			}
			obj.LdapPoll = &v2
		}
	}
	if v1, ok := d.GetOk("ldap_poll_filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldap_poll_filter", sv)
				diags = append(diags, e)
			}
			obj.LdapPollFilter = &v2
		}
	}
	if v1, ok := d.GetOk("ldap_poll_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldap_poll_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LdapPollInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("ldap_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldap_server", sv)
				diags = append(diags, e)
			}
			obj.LdapServer = &v2
		}
	}
	if v1, ok := d.GetOk("logon_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.7", "v7.0.0") {
				e := utils.AttributeVersionWarning("logon_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LogonTimeout = &tmp
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
	if v1, ok := d.GetOk("password2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password2", sv)
				diags = append(diags, e)
			}
			obj.Password2 = &v2
		}
	}
	if v1, ok := d.GetOk("password3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password3", sv)
				diags = append(diags, e)
			}
			obj.Password3 = &v2
		}
	}
	if v1, ok := d.GetOk("password4"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password4", sv)
				diags = append(diags, e)
			}
			obj.Password4 = &v2
		}
	}
	if v1, ok := d.GetOk("password5"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password5", sv)
				diags = append(diags, e)
			}
			obj.Password5 = &v2
		}
	}
	if v1, ok := d.GetOk("port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port = &tmp
		}
	}
	if v1, ok := d.GetOk("port2"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port2", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port2 = &tmp
		}
	}
	if v1, ok := d.GetOk("port3"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port3", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port3 = &tmp
		}
	}
	if v1, ok := d.GetOk("port4"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port4", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port4 = &tmp
		}
	}
	if v1, ok := d.GetOk("port5"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port5", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port5 = &tmp
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
	if v1, ok := d.GetOk("server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server2", sv)
				diags = append(diags, e)
			}
			obj.Server2 = &v2
		}
	}
	if v1, ok := d.GetOk("server3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server3", sv)
				diags = append(diags, e)
			}
			obj.Server3 = &v2
		}
	}
	if v1, ok := d.GetOk("server4"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server4", sv)
				diags = append(diags, e)
			}
			obj.Server4 = &v2
		}
	}
	if v1, ok := d.GetOk("server5"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server5", sv)
				diags = append(diags, e)
			}
			obj.Server5 = &v2
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
	if v1, ok := d.GetOk("source_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip6", sv)
				diags = append(diags, e)
			}
			obj.SourceIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("ssl"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl", sv)
				diags = append(diags, e)
			}
			obj.Ssl = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_server_host_ip_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("ssl_server_host_ip_check", sv)
				diags = append(diags, e)
			}
			obj.SslServerHostIpCheck = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_trusted_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_trusted_cert", sv)
				diags = append(diags, e)
			}
			obj.SslTrustedCert = &v2
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("user_info_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_info_server", sv)
				diags = append(diags, e)
			}
			obj.UserInfoServer = &v2
		}
	}
	return &obj, diags
}
