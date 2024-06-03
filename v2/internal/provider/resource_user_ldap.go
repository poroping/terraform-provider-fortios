// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure LDAP server entries.

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

func resourceUserLdap() *schema.Resource {
	return &schema.Resource{
		Description: "Configure LDAP server entries.",

		CreateContext: resourceUserLdapCreate,
		ReadContext:   resourceUserLdapRead,
		UpdateContext: resourceUserLdapUpdate,
		DeleteContext: resourceUserLdapDelete,

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
			"account_key_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),

				Description: "Account key filter, using the UPN as the search filter.",
				Optional:    true,
				Computed:    true,
			},
			"account_key_processing": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"same", "strip"}, false),

				Description: "Account key processing operation, either keep or strip domain string of UPN in the token.",
				Optional:    true,
				Computed:    true,
			},
			"account_key_upn_san": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"othername", "rfc822name", "dnsname"}, false),

				Description: "Define SAN in certificate for user principle name matching.",
				Optional:    true,
				Computed:    true,
			},
			"antiphish": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable AntiPhishing credential backend.",
				Optional:    true,
				Computed:    true,
			},
			"ca_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "CA certificate name.",
				Optional:    true,
				Computed:    true,
			},
			"client_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Client certificate name.",
				Optional:    true,
				Computed:    true,
			},
			"client_cert_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable using client certificate for TLS authentication.",
				Optional:    true,
				Computed:    true,
			},
			"cnid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 20),

				Description: "Common name identifier for the LDAP server. The common name identifier for most LDAP servers is \"cn\".",
				Optional:    true,
				Computed:    true,
			},
			"dn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),

				Description: "Distinguished name used to look up entries on the LDAP server.",
				Optional:    true,
				Computed:    true,
			},
			"group_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),

				Description: "Filter used for group matching.",
				Optional:    true,
				Computed:    true,
			},
			"group_member_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"user-attr", "group-object", "posix-group-object"}, false),

				Description: "Group member checking methods.",
				Optional:    true,
				Computed:    true,
			},
			"group_object_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),

				Description: "Filter used for group searching.",
				Optional:    true,
				Computed:    true,
			},
			"group_search_base": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),

				Description: "Search base used for group searching.",
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
			"member_attr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Name of attribute from which to get group membership.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "LDAP server entry name.",
				ForceNew:    true,
				Required:    true,
			},
			"obtain_user_info": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable obtaining of user information.",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Password for initial binding.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"password_attr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of attribute to get password hash.",
				Optional:    true,
				Computed:    true,
			},
			"password_expiry_warning": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable password expiry warnings.",
				Optional:    true,
				Computed:    true,
			},
			"password_renewal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable online password renewal.",
				Optional:    true,
				Computed:    true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port to be used for communication with the LDAP server (default = 389).",
				Optional:    true,
				Computed:    true,
			},
			"search_type": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Search type.",
				Optional:         true,
				Computed:         true,
			},
			"secondary_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Secondary LDAP server CN domain name or IP.",
				Optional:    true,
				Computed:    true,
			},
			"secure": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "starttls", "ldaps"}, false),

				Description: "Port to be used for authentication.",
				Optional:    true,
				Computed:    true,
			},
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "LDAP server CN domain name or IP.",
				Optional:    true,
				Computed:    true,
			},
			"server_identity_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate).",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "FortiGate IP address to be used for communication with the LDAP server.",
				Optional:    true,
				Computed:    true,
			},
			"source_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Source port to be used for communication with the LDAP server.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "SSLv3", "TLSv1", "TLSv1-1", "TLSv1-2"}, false),

				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
				Optional:    true,
				Computed:    true,
			},
			"tertiary_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Tertiary LDAP server CN domain name or IP.",
				Optional:    true,
				Computed:    true,
			},
			"two_factor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "fortitoken-cloud"}, false),

				Description: "Enable/disable two-factor authentication.",
				Optional:    true,
				Computed:    true,
			},
			"two_factor_authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fortitoken", "email", "sms"}, false),

				Description: "Authentication method by FortiToken Cloud.",
				Optional:    true,
				Computed:    true,
			},
			"two_factor_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),

				Description: "Filter used to synchronize users to FortiToken Cloud.",
				Optional:    true,
				Computed:    true,
			},
			"two_factor_notification": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"email", "sms"}, false),

				Description: "Notification method for user activation by FortiToken Cloud.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"simple", "anonymous", "regular"}, false),

				Description: "Authentication type for LDAP searches.",
				Optional:    true,
				Computed:    true,
			},
			"user_info_exchange_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "MS Exchange server from which to fetch user information.",
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),

				Description: "Username (full DN) for initial binding.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserLdapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserLdap resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserLdap(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserLdap(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserLdap")
	}

	return resourceUserLdapRead(ctx, d, meta)
}

func resourceUserLdapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserLdap(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserLdap(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserLdap resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserLdap")
	}

	return resourceUserLdapRead(ctx, d, meta)
}

func resourceUserLdapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserLdap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserLdap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserLdapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserLdap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserLdap resource: %v", err)
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

	diags := refreshObjectUserLdap(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectUserLdap(d *schema.ResourceData, o *models.UserLdap, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AccountKeyFilter != nil {
		v := *o.AccountKeyFilter

		if err = d.Set("account_key_filter", v); err != nil {
			return diag.Errorf("error reading account_key_filter: %v", err)
		}
	}

	if o.AccountKeyProcessing != nil {
		v := *o.AccountKeyProcessing

		if err = d.Set("account_key_processing", v); err != nil {
			return diag.Errorf("error reading account_key_processing: %v", err)
		}
	}

	if o.AccountKeyUpnSan != nil {
		v := *o.AccountKeyUpnSan

		if err = d.Set("account_key_upn_san", v); err != nil {
			return diag.Errorf("error reading account_key_upn_san: %v", err)
		}
	}

	if o.Antiphish != nil {
		v := *o.Antiphish

		if err = d.Set("antiphish", v); err != nil {
			return diag.Errorf("error reading antiphish: %v", err)
		}
	}

	if o.CaCert != nil {
		v := *o.CaCert

		if err = d.Set("ca_cert", v); err != nil {
			return diag.Errorf("error reading ca_cert: %v", err)
		}
	}

	if o.ClientCert != nil {
		v := *o.ClientCert

		if err = d.Set("client_cert", v); err != nil {
			return diag.Errorf("error reading client_cert: %v", err)
		}
	}

	if o.ClientCertAuth != nil {
		v := *o.ClientCertAuth

		if err = d.Set("client_cert_auth", v); err != nil {
			return diag.Errorf("error reading client_cert_auth: %v", err)
		}
	}

	if o.Cnid != nil {
		v := *o.Cnid

		if err = d.Set("cnid", v); err != nil {
			return diag.Errorf("error reading cnid: %v", err)
		}
	}

	if o.Dn != nil {
		v := *o.Dn

		if err = d.Set("dn", v); err != nil {
			return diag.Errorf("error reading dn: %v", err)
		}
	}

	if o.GroupFilter != nil {
		v := *o.GroupFilter

		if err = d.Set("group_filter", v); err != nil {
			return diag.Errorf("error reading group_filter: %v", err)
		}
	}

	if o.GroupMemberCheck != nil {
		v := *o.GroupMemberCheck

		if err = d.Set("group_member_check", v); err != nil {
			return diag.Errorf("error reading group_member_check: %v", err)
		}
	}

	if o.GroupObjectFilter != nil {
		v := *o.GroupObjectFilter

		if err = d.Set("group_object_filter", v); err != nil {
			return diag.Errorf("error reading group_object_filter: %v", err)
		}
	}

	if o.GroupSearchBase != nil {
		v := *o.GroupSearchBase

		if err = d.Set("group_search_base", v); err != nil {
			return diag.Errorf("error reading group_search_base: %v", err)
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

	if o.MemberAttr != nil {
		v := *o.MemberAttr

		if err = d.Set("member_attr", v); err != nil {
			return diag.Errorf("error reading member_attr: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ObtainUserInfo != nil {
		v := *o.ObtainUserInfo

		if err = d.Set("obtain_user_info", v); err != nil {
			return diag.Errorf("error reading obtain_user_info: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.PasswordAttr != nil {
		v := *o.PasswordAttr

		if err = d.Set("password_attr", v); err != nil {
			return diag.Errorf("error reading password_attr: %v", err)
		}
	}

	if o.PasswordExpiryWarning != nil {
		v := *o.PasswordExpiryWarning

		if err = d.Set("password_expiry_warning", v); err != nil {
			return diag.Errorf("error reading password_expiry_warning: %v", err)
		}
	}

	if o.PasswordRenewal != nil {
		v := *o.PasswordRenewal

		if err = d.Set("password_renewal", v); err != nil {
			return diag.Errorf("error reading password_renewal: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.SearchType != nil {
		v := *o.SearchType

		if err = d.Set("search_type", v); err != nil {
			return diag.Errorf("error reading search_type: %v", err)
		}
	}

	if o.SecondaryServer != nil {
		v := *o.SecondaryServer

		if err = d.Set("secondary_server", v); err != nil {
			return diag.Errorf("error reading secondary_server: %v", err)
		}
	}

	if o.Secure != nil {
		v := *o.Secure

		if err = d.Set("secure", v); err != nil {
			return diag.Errorf("error reading secure: %v", err)
		}
	}

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.ServerIdentityCheck != nil {
		v := *o.ServerIdentityCheck

		if err = d.Set("server_identity_check", v); err != nil {
			return diag.Errorf("error reading server_identity_check: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.SourcePort != nil {
		v := *o.SourcePort

		if err = d.Set("source_port", v); err != nil {
			return diag.Errorf("error reading source_port: %v", err)
		}
	}

	if o.SslMinProtoVersion != nil {
		v := *o.SslMinProtoVersion

		if err = d.Set("ssl_min_proto_version", v); err != nil {
			return diag.Errorf("error reading ssl_min_proto_version: %v", err)
		}
	}

	if o.TertiaryServer != nil {
		v := *o.TertiaryServer

		if err = d.Set("tertiary_server", v); err != nil {
			return diag.Errorf("error reading tertiary_server: %v", err)
		}
	}

	if o.TwoFactor != nil {
		v := *o.TwoFactor

		if err = d.Set("two_factor", v); err != nil {
			return diag.Errorf("error reading two_factor: %v", err)
		}
	}

	if o.TwoFactorAuthentication != nil {
		v := *o.TwoFactorAuthentication

		if err = d.Set("two_factor_authentication", v); err != nil {
			return diag.Errorf("error reading two_factor_authentication: %v", err)
		}
	}

	if o.TwoFactorFilter != nil {
		v := *o.TwoFactorFilter

		if err = d.Set("two_factor_filter", v); err != nil {
			return diag.Errorf("error reading two_factor_filter: %v", err)
		}
	}

	if o.TwoFactorNotification != nil {
		v := *o.TwoFactorNotification

		if err = d.Set("two_factor_notification", v); err != nil {
			return diag.Errorf("error reading two_factor_notification: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.UserInfoExchangeServer != nil {
		v := *o.UserInfoExchangeServer

		if err = d.Set("user_info_exchange_server", v); err != nil {
			return diag.Errorf("error reading user_info_exchange_server: %v", err)
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

func getObjectUserLdap(d *schema.ResourceData, sv string) (*models.UserLdap, diag.Diagnostics) {
	obj := models.UserLdap{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("account_key_filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("account_key_filter", sv)
				diags = append(diags, e)
			}
			obj.AccountKeyFilter = &v2
		}
	}
	if v1, ok := d.GetOk("account_key_processing"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("account_key_processing", sv)
				diags = append(diags, e)
			}
			obj.AccountKeyProcessing = &v2
		}
	}
	if v1, ok := d.GetOk("account_key_upn_san"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("account_key_upn_san", sv)
				diags = append(diags, e)
			}
			obj.AccountKeyUpnSan = &v2
		}
	}
	if v1, ok := d.GetOk("antiphish"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("antiphish", sv)
				diags = append(diags, e)
			}
			obj.Antiphish = &v2
		}
	}
	if v1, ok := d.GetOk("ca_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ca_cert", sv)
				diags = append(diags, e)
			}
			obj.CaCert = &v2
		}
	}
	if v1, ok := d.GetOk("client_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("client_cert", sv)
				diags = append(diags, e)
			}
			obj.ClientCert = &v2
		}
	}
	if v1, ok := d.GetOk("client_cert_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("client_cert_auth", sv)
				diags = append(diags, e)
			}
			obj.ClientCertAuth = &v2
		}
	}
	if v1, ok := d.GetOk("cnid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cnid", sv)
				diags = append(diags, e)
			}
			obj.Cnid = &v2
		}
	}
	if v1, ok := d.GetOk("dn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dn", sv)
				diags = append(diags, e)
			}
			obj.Dn = &v2
		}
	}
	if v1, ok := d.GetOk("group_filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_filter", sv)
				diags = append(diags, e)
			}
			obj.GroupFilter = &v2
		}
	}
	if v1, ok := d.GetOk("group_member_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_member_check", sv)
				diags = append(diags, e)
			}
			obj.GroupMemberCheck = &v2
		}
	}
	if v1, ok := d.GetOk("group_object_filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_object_filter", sv)
				diags = append(diags, e)
			}
			obj.GroupObjectFilter = &v2
		}
	}
	if v1, ok := d.GetOk("group_search_base"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_search_base", sv)
				diags = append(diags, e)
			}
			obj.GroupSearchBase = &v2
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
	if v1, ok := d.GetOk("member_attr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("member_attr", sv)
				diags = append(diags, e)
			}
			obj.MemberAttr = &v2
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
	if v1, ok := d.GetOk("obtain_user_info"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("obtain_user_info", sv)
				diags = append(diags, e)
			}
			obj.ObtainUserInfo = &v2
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
	if v1, ok := d.GetOk("password_attr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("password_attr", sv)
				diags = append(diags, e)
			}
			obj.PasswordAttr = &v2
		}
	}
	if v1, ok := d.GetOk("password_expiry_warning"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password_expiry_warning", sv)
				diags = append(diags, e)
			}
			obj.PasswordExpiryWarning = &v2
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
	if v1, ok := d.GetOk("search_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("search_type", sv)
				diags = append(diags, e)
			}
			obj.SearchType = &v2
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
	if v1, ok := d.GetOk("secure"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secure", sv)
				diags = append(diags, e)
			}
			obj.Secure = &v2
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
	if v1, ok := d.GetOk("server_identity_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_identity_check", sv)
				diags = append(diags, e)
			}
			obj.ServerIdentityCheck = &v2
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
	if v1, ok := d.GetOk("source_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("source_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SourcePort = &tmp
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
	if v1, ok := d.GetOk("tertiary_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tertiary_server", sv)
				diags = append(diags, e)
			}
			obj.TertiaryServer = &v2
		}
	}
	if v1, ok := d.GetOk("two_factor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor", sv)
				diags = append(diags, e)
			}
			obj.TwoFactor = &v2
		}
	}
	if v1, ok := d.GetOk("two_factor_authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor_authentication", sv)
				diags = append(diags, e)
			}
			obj.TwoFactorAuthentication = &v2
		}
	}
	if v1, ok := d.GetOk("two_factor_filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.1", "") {
				e := utils.AttributeVersionWarning("two_factor_filter", sv)
				diags = append(diags, e)
			}
			obj.TwoFactorFilter = &v2
		}
	}
	if v1, ok := d.GetOk("two_factor_notification"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor_notification", sv)
				diags = append(diags, e)
			}
			obj.TwoFactorNotification = &v2
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
	if v1, ok := d.GetOk("user_info_exchange_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_info_exchange_server", sv)
				diags = append(diags, e)
			}
			obj.UserInfoExchangeServer = &v2
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
