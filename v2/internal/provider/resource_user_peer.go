// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure peer users.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceUserPeer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure peer users.",

		CreateContext: resourceUserPeerCreate,
		ReadContext:   resourceUserPeerRead,
		UpdateContext: resourceUserPeerUpdate,
		DeleteContext: resourceUserPeerDelete,

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
			"ca": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Name of the CA certificate.",
				Optional:    true,
				Computed:    true,
			},
			"cn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Peer certificate common name.",
				Optional:    true,
				Computed:    true,
			},
			"cn_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"string", "email", "FQDN", "ipv4", "ipv6"}, false),

				Description: "Peer certificate common name type.",
				Optional:    true,
				Computed:    true,
			},
			"ldap_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"password", "principal-name"}, false),

				Description: "Mode for LDAP peer authentication.",
				Optional:    true,
				Computed:    true,
			},
			"ldap_password": {
				Type: schema.TypeString,

				Description: "Password for LDAP server bind.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"ldap_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an LDAP server defined under the user ldap command. Performs client access rights check.",
				Optional:    true,
				Computed:    true,
			},
			"ldap_username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Username for LDAP server bind.",
				Optional:    true,
				Computed:    true,
			},
			"mandatory_ca_verify": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Peer name.",
				ForceNew:    true,
				Required:    true,
			},
			"ocsp_override_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Online Certificate Status Protocol (OCSP) server for certificate retrieval.",
				Optional:    true,
				Computed:    true,
			},
			"passwd": {
				Type: schema.TypeString,

				Description: "Peer's password used for two-factor authentication.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"subject": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Peer certificate name constraints.",
				Optional:    true,
				Computed:    true,
			},
			"two_factor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable two-factor authentication, applying certificate and password-based authentication.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserPeerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserPeer resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserPeer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserPeer(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserPeer")
	}

	return resourceUserPeerRead(ctx, d, meta)
}

func resourceUserPeerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserPeer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserPeer(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserPeer resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserPeer")
	}

	return resourceUserPeerRead(ctx, d, meta)
}

func resourceUserPeerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserPeer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserPeerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserPeer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserPeer resource: %v", err)
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

	diags := refreshObjectUserPeer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectUserPeer(d *schema.ResourceData, o *models.UserPeer, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Ca != nil {
		v := *o.Ca

		if err = d.Set("ca", v); err != nil {
			return diag.Errorf("error reading ca: %v", err)
		}
	}

	if o.Cn != nil {
		v := *o.Cn

		if err = d.Set("cn", v); err != nil {
			return diag.Errorf("error reading cn: %v", err)
		}
	}

	if o.CnType != nil {
		v := *o.CnType

		if err = d.Set("cn_type", v); err != nil {
			return diag.Errorf("error reading cn_type: %v", err)
		}
	}

	if o.LdapMode != nil {
		v := *o.LdapMode

		if err = d.Set("ldap_mode", v); err != nil {
			return diag.Errorf("error reading ldap_mode: %v", err)
		}
	}

	if o.LdapPassword != nil {
		v := *o.LdapPassword

		if err = d.Set("ldap_password", v); err != nil {
			return diag.Errorf("error reading ldap_password: %v", err)
		}
	}

	if o.LdapServer != nil {
		v := *o.LdapServer

		if err = d.Set("ldap_server", v); err != nil {
			return diag.Errorf("error reading ldap_server: %v", err)
		}
	}

	if o.LdapUsername != nil {
		v := *o.LdapUsername

		if err = d.Set("ldap_username", v); err != nil {
			return diag.Errorf("error reading ldap_username: %v", err)
		}
	}

	if o.MandatoryCaVerify != nil {
		v := *o.MandatoryCaVerify

		if err = d.Set("mandatory_ca_verify", v); err != nil {
			return diag.Errorf("error reading mandatory_ca_verify: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.OcspOverrideServer != nil {
		v := *o.OcspOverrideServer

		if err = d.Set("ocsp_override_server", v); err != nil {
			return diag.Errorf("error reading ocsp_override_server: %v", err)
		}
	}

	if o.Passwd != nil {
		v := *o.Passwd

		if err = d.Set("passwd", v); err != nil {
			return diag.Errorf("error reading passwd: %v", err)
		}
	}

	if o.Subject != nil {
		v := *o.Subject

		if err = d.Set("subject", v); err != nil {
			return diag.Errorf("error reading subject: %v", err)
		}
	}

	if o.TwoFactor != nil {
		v := *o.TwoFactor

		if err = d.Set("two_factor", v); err != nil {
			return diag.Errorf("error reading two_factor: %v", err)
		}
	}

	return nil
}

func getObjectUserPeer(d *schema.ResourceData, sv string) (*models.UserPeer, diag.Diagnostics) {
	obj := models.UserPeer{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("ca"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ca", sv)
				diags = append(diags, e)
			}
			obj.Ca = &v2
		}
	}
	if v1, ok := d.GetOk("cn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cn", sv)
				diags = append(diags, e)
			}
			obj.Cn = &v2
		}
	}
	if v1, ok := d.GetOk("cn_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cn_type", sv)
				diags = append(diags, e)
			}
			obj.CnType = &v2
		}
	}
	if v1, ok := d.GetOk("ldap_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldap_mode", sv)
				diags = append(diags, e)
			}
			obj.LdapMode = &v2
		}
	}
	if v1, ok := d.GetOk("ldap_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldap_password", sv)
				diags = append(diags, e)
			}
			obj.LdapPassword = &v2
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
	if v1, ok := d.GetOk("ldap_username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldap_username", sv)
				diags = append(diags, e)
			}
			obj.LdapUsername = &v2
		}
	}
	if v1, ok := d.GetOk("mandatory_ca_verify"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mandatory_ca_verify", sv)
				diags = append(diags, e)
			}
			obj.MandatoryCaVerify = &v2
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
	if v1, ok := d.GetOk("ocsp_override_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ocsp_override_server", sv)
				diags = append(diags, e)
			}
			obj.OcspOverrideServer = &v2
		}
	}
	if v1, ok := d.GetOk("passwd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("passwd", sv)
				diags = append(diags, e)
			}
			obj.Passwd = &v2
		}
	}
	if v1, ok := d.GetOk("subject"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("subject", sv)
				diags = append(diags, e)
			}
			obj.Subject = &v2
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
	return &obj, diags
}
