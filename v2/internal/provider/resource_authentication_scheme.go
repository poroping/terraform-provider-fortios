// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Authentication Schemes.

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

func resourceAuthenticationScheme() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Authentication Schemes.",

		CreateContext: resourceAuthenticationSchemeCreate,
		ReadContext:   resourceAuthenticationSchemeRead,
		UpdateContext: resourceAuthenticationSchemeUpdate,
		DeleteContext: resourceAuthenticationSchemeDelete,

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
			"domain_controller": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Domain controller setting.",
				Optional:    true,
				Computed:    true,
			},
			"ems_device_owner": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSH public-key authentication with device owner (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"fsso_agent_for_ntlm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "FSSO agent to use for NTLM authentication.",
				Optional:    true,
				Computed:    true,
			},
			"fsso_guest": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable user fsso-guest authentication (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"kerberos_keytab": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Kerberos keytab setting.",
				Optional:    true,
				Computed:    true,
			},
			"method": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Authentication methods (default = basic).",
				Optional:         true,
				Computed:         true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Authentication scheme name.",
				ForceNew:    true,
				Required:    true,
			},
			"negotiate_ntlm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable negotiate authentication for NTLM (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"require_tfa": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable two-factor authentication (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"saml_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SAML configuration.",
				Optional:    true,
				Computed:    true,
			},
			"saml_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 1200),

				Description: "SAML authentication timeout in seconds.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_ca": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SSH CA name.",
				Optional:    true,
				Computed:    true,
			},
			"user_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable authentication with user certificate (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"user_database": {
				Type:        schema.TypeList,
				Description: "Authentication server to contain user information; \"local\" (default) or \"123\" (for LDAP).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Authentication server name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceAuthenticationSchemeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating AuthenticationScheme resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectAuthenticationScheme(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateAuthenticationScheme(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AuthenticationScheme")
	}

	return resourceAuthenticationSchemeRead(ctx, d, meta)
}

func resourceAuthenticationSchemeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectAuthenticationScheme(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateAuthenticationScheme(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating AuthenticationScheme resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AuthenticationScheme")
	}

	return resourceAuthenticationSchemeRead(ctx, d, meta)
}

func resourceAuthenticationSchemeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteAuthenticationScheme(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting AuthenticationScheme resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAuthenticationSchemeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadAuthenticationScheme(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AuthenticationScheme resource: %v", err)
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

	diags := refreshObjectAuthenticationScheme(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenAuthenticationSchemeUserDatabase(d *schema.ResourceData, v *[]models.AuthenticationSchemeUserDatabase, prefix string, sort bool) interface{} {
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

func refreshObjectAuthenticationScheme(d *schema.ResourceData, o *models.AuthenticationScheme, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DomainController != nil {
		v := *o.DomainController

		if err = d.Set("domain_controller", v); err != nil {
			return diag.Errorf("error reading domain_controller: %v", err)
		}
	}

	if o.EmsDeviceOwner != nil {
		v := *o.EmsDeviceOwner

		if err = d.Set("ems_device_owner", v); err != nil {
			return diag.Errorf("error reading ems_device_owner: %v", err)
		}
	}

	if o.FssoAgentForNtlm != nil {
		v := *o.FssoAgentForNtlm

		if err = d.Set("fsso_agent_for_ntlm", v); err != nil {
			return diag.Errorf("error reading fsso_agent_for_ntlm: %v", err)
		}
	}

	if o.FssoGuest != nil {
		v := *o.FssoGuest

		if err = d.Set("fsso_guest", v); err != nil {
			return diag.Errorf("error reading fsso_guest: %v", err)
		}
	}

	if o.KerberosKeytab != nil {
		v := *o.KerberosKeytab

		if err = d.Set("kerberos_keytab", v); err != nil {
			return diag.Errorf("error reading kerberos_keytab: %v", err)
		}
	}

	if o.Method != nil {
		v := *o.Method

		if err = d.Set("method", v); err != nil {
			return diag.Errorf("error reading method: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.NegotiateNtlm != nil {
		v := *o.NegotiateNtlm

		if err = d.Set("negotiate_ntlm", v); err != nil {
			return diag.Errorf("error reading negotiate_ntlm: %v", err)
		}
	}

	if o.RequireTfa != nil {
		v := *o.RequireTfa

		if err = d.Set("require_tfa", v); err != nil {
			return diag.Errorf("error reading require_tfa: %v", err)
		}
	}

	if o.SamlServer != nil {
		v := *o.SamlServer

		if err = d.Set("saml_server", v); err != nil {
			return diag.Errorf("error reading saml_server: %v", err)
		}
	}

	if o.SamlTimeout != nil {
		v := *o.SamlTimeout

		if err = d.Set("saml_timeout", v); err != nil {
			return diag.Errorf("error reading saml_timeout: %v", err)
		}
	}

	if o.SshCa != nil {
		v := *o.SshCa

		if err = d.Set("ssh_ca", v); err != nil {
			return diag.Errorf("error reading ssh_ca: %v", err)
		}
	}

	if o.UserCert != nil {
		v := *o.UserCert

		if err = d.Set("user_cert", v); err != nil {
			return diag.Errorf("error reading user_cert: %v", err)
		}
	}

	if o.UserDatabase != nil {
		if err = d.Set("user_database", flattenAuthenticationSchemeUserDatabase(d, o.UserDatabase, "user_database", sort)); err != nil {
			return diag.Errorf("error reading user_database: %v", err)
		}
	}

	return nil
}

func expandAuthenticationSchemeUserDatabase(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.AuthenticationSchemeUserDatabase, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AuthenticationSchemeUserDatabase

	for i := range l {
		tmp := models.AuthenticationSchemeUserDatabase{}
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

func getObjectAuthenticationScheme(d *schema.ResourceData, sv string) (*models.AuthenticationScheme, diag.Diagnostics) {
	obj := models.AuthenticationScheme{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("domain_controller"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("domain_controller", sv)
				diags = append(diags, e)
			}
			obj.DomainController = &v2
		}
	}
	if v1, ok := d.GetOk("ems_device_owner"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "v7.0.1") {
				e := utils.AttributeVersionWarning("ems_device_owner", sv)
				diags = append(diags, e)
			}
			obj.EmsDeviceOwner = &v2
		}
	}
	if v1, ok := d.GetOk("fsso_agent_for_ntlm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fsso_agent_for_ntlm", sv)
				diags = append(diags, e)
			}
			obj.FssoAgentForNtlm = &v2
		}
	}
	if v1, ok := d.GetOk("fsso_guest"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fsso_guest", sv)
				diags = append(diags, e)
			}
			obj.FssoGuest = &v2
		}
	}
	if v1, ok := d.GetOk("kerberos_keytab"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("kerberos_keytab", sv)
				diags = append(diags, e)
			}
			obj.KerberosKeytab = &v2
		}
	}
	if v1, ok := d.GetOk("method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("method", sv)
				diags = append(diags, e)
			}
			obj.Method = &v2
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
	if v1, ok := d.GetOk("negotiate_ntlm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("negotiate_ntlm", sv)
				diags = append(diags, e)
			}
			obj.NegotiateNtlm = &v2
		}
	}
	if v1, ok := d.GetOk("require_tfa"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("require_tfa", sv)
				diags = append(diags, e)
			}
			obj.RequireTfa = &v2
		}
	}
	if v1, ok := d.GetOk("saml_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("saml_server", sv)
				diags = append(diags, e)
			}
			obj.SamlServer = &v2
		}
	}
	if v1, ok := d.GetOk("saml_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("saml_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SamlTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("ssh_ca"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssh_ca", sv)
				diags = append(diags, e)
			}
			obj.SshCa = &v2
		}
	}
	if v1, ok := d.GetOk("user_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("user_cert", sv)
				diags = append(diags, e)
			}
			obj.UserCert = &v2
		}
	}
	if v, ok := d.GetOk("user_database"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("user_database", sv)
			diags = append(diags, e)
		}
		t, err := expandAuthenticationSchemeUserDatabase(d, v, "user_database", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.UserDatabase = t
		}
	} else if d.HasChange("user_database") {
		old, new := d.GetChange("user_database")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.UserDatabase = &[]models.AuthenticationSchemeUserDatabase{}
		}
	}
	return &obj, diags
}
