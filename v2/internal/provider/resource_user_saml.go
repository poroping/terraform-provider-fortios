// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SAML server entry configuration.

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

func resourceUserSaml() *schema.Resource {
	return &schema.Resource{
		Description: "SAML server entry configuration.",

		CreateContext: resourceUserSamlCreate,
		ReadContext:   resourceUserSamlRead,
		UpdateContext: resourceUserSamlUpdate,
		DeleteContext: resourceUserSamlDelete,

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
			"adfs_claim": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Certificate to sign SAML messages.",
				Optional:    true,
				Computed:    true,
			},
			"digest_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"sha1", "sha256"}, false),

				Description: "Digest Method Algorithm. (default = sha1).",
				Optional:    true,
				Computed:    true,
			},
			"entity_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "SP entity ID.",
				Optional:    true,
				Computed:    true,
			},
			"group_claim_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"email", "given-name", "name", "upn", "common-name", "email-adfs-1x", "group", "upn-adfs-1x", "role", "sur-name", "ppid", "name-identifier", "authentication-method", "deny-only-group-sid", "deny-only-primary-sid", "deny-only-primary-group-sid", "group-sid", "primary-group-sid", "primary-sid", "windows-account-name"}, false),

				Description: "Group claim in assertion statement.",
				Optional:    true,
				Computed:    true,
			},
			"group_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Group name in assertion statement.",
				Optional:    true,
				Computed:    true,
			},
			"idp_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "IDP Certificate name.",
				Optional:    true,
				Computed:    true,
			},
			"idp_entity_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "IDP entity ID.",
				Optional:    true,
				Computed:    true,
			},
			"idp_single_logout_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "IDP single logout url.",
				Optional:    true,
				Computed:    true,
			},
			"idp_single_sign_on_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "IDP single sign-on URL.",
				Optional:    true,
				Computed:    true,
			},
			"limit_relaystate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SAML server entry name.",
				ForceNew:    true,
				Required:    true,
			},
			"single_logout_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "SP single logout URL.",
				Optional:    true,
				Computed:    true,
			},
			"single_sign_on_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "SP single sign-on URL.",
				Optional:    true,
				Computed:    true,
			},
			"user_claim_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"email", "given-name", "name", "upn", "common-name", "email-adfs-1x", "group", "upn-adfs-1x", "role", "sur-name", "ppid", "name-identifier", "authentication-method", "deny-only-group-sid", "deny-only-primary-sid", "deny-only-primary-group-sid", "group-sid", "primary-group-sid", "primary-sid", "windows-account-name"}, false),

				Description: "User name claim in assertion statement.",
				Optional:    true,
				Computed:    true,
			},
			"user_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "User name in assertion statement.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserSamlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserSaml resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserSaml(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserSaml(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserSaml")
	}

	return resourceUserSamlRead(ctx, d, meta)
}

func resourceUserSamlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserSaml(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserSaml(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserSaml resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserSaml")
	}

	return resourceUserSamlRead(ctx, d, meta)
}

func resourceUserSamlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserSaml(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserSaml resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserSamlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserSaml(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserSaml resource: %v", err)
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

	diags := refreshObjectUserSaml(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectUserSaml(d *schema.ResourceData, o *models.UserSaml, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AdfsClaim != nil {
		v := *o.AdfsClaim

		if err = d.Set("adfs_claim", v); err != nil {
			return diag.Errorf("error reading adfs_claim: %v", err)
		}
	}

	if o.Cert != nil {
		v := *o.Cert

		if err = d.Set("cert", v); err != nil {
			return diag.Errorf("error reading cert: %v", err)
		}
	}

	if o.DigestMethod != nil {
		v := *o.DigestMethod

		if err = d.Set("digest_method", v); err != nil {
			return diag.Errorf("error reading digest_method: %v", err)
		}
	}

	if o.EntityId != nil {
		v := *o.EntityId

		if err = d.Set("entity_id", v); err != nil {
			return diag.Errorf("error reading entity_id: %v", err)
		}
	}

	if o.GroupClaimType != nil {
		v := *o.GroupClaimType

		if err = d.Set("group_claim_type", v); err != nil {
			return diag.Errorf("error reading group_claim_type: %v", err)
		}
	}

	if o.GroupName != nil {
		v := *o.GroupName

		if err = d.Set("group_name", v); err != nil {
			return diag.Errorf("error reading group_name: %v", err)
		}
	}

	if o.IdpCert != nil {
		v := *o.IdpCert

		if err = d.Set("idp_cert", v); err != nil {
			return diag.Errorf("error reading idp_cert: %v", err)
		}
	}

	if o.IdpEntityId != nil {
		v := *o.IdpEntityId

		if err = d.Set("idp_entity_id", v); err != nil {
			return diag.Errorf("error reading idp_entity_id: %v", err)
		}
	}

	if o.IdpSingleLogoutUrl != nil {
		v := *o.IdpSingleLogoutUrl

		if err = d.Set("idp_single_logout_url", v); err != nil {
			return diag.Errorf("error reading idp_single_logout_url: %v", err)
		}
	}

	if o.IdpSingleSignOnUrl != nil {
		v := *o.IdpSingleSignOnUrl

		if err = d.Set("idp_single_sign_on_url", v); err != nil {
			return diag.Errorf("error reading idp_single_sign_on_url: %v", err)
		}
	}

	if o.LimitRelaystate != nil {
		v := *o.LimitRelaystate

		if err = d.Set("limit_relaystate", v); err != nil {
			return diag.Errorf("error reading limit_relaystate: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.SingleLogoutUrl != nil {
		v := *o.SingleLogoutUrl

		if err = d.Set("single_logout_url", v); err != nil {
			return diag.Errorf("error reading single_logout_url: %v", err)
		}
	}

	if o.SingleSignOnUrl != nil {
		v := *o.SingleSignOnUrl

		if err = d.Set("single_sign_on_url", v); err != nil {
			return diag.Errorf("error reading single_sign_on_url: %v", err)
		}
	}

	if o.UserClaimType != nil {
		v := *o.UserClaimType

		if err = d.Set("user_claim_type", v); err != nil {
			return diag.Errorf("error reading user_claim_type: %v", err)
		}
	}

	if o.UserName != nil {
		v := *o.UserName

		if err = d.Set("user_name", v); err != nil {
			return diag.Errorf("error reading user_name: %v", err)
		}
	}

	return nil
}

func getObjectUserSaml(d *schema.ResourceData, sv string) (*models.UserSaml, diag.Diagnostics) {
	obj := models.UserSaml{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("adfs_claim"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("adfs_claim", sv)
				diags = append(diags, e)
			}
			obj.AdfsClaim = &v2
		}
	}
	if v1, ok := d.GetOk("cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cert", sv)
				diags = append(diags, e)
			}
			obj.Cert = &v2
		}
	}
	if v1, ok := d.GetOk("digest_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("digest_method", sv)
				diags = append(diags, e)
			}
			obj.DigestMethod = &v2
		}
	}
	if v1, ok := d.GetOk("entity_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("entity_id", sv)
				diags = append(diags, e)
			}
			obj.EntityId = &v2
		}
	}
	if v1, ok := d.GetOk("group_claim_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("group_claim_type", sv)
				diags = append(diags, e)
			}
			obj.GroupClaimType = &v2
		}
	}
	if v1, ok := d.GetOk("group_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_name", sv)
				diags = append(diags, e)
			}
			obj.GroupName = &v2
		}
	}
	if v1, ok := d.GetOk("idp_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idp_cert", sv)
				diags = append(diags, e)
			}
			obj.IdpCert = &v2
		}
	}
	if v1, ok := d.GetOk("idp_entity_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idp_entity_id", sv)
				diags = append(diags, e)
			}
			obj.IdpEntityId = &v2
		}
	}
	if v1, ok := d.GetOk("idp_single_logout_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idp_single_logout_url", sv)
				diags = append(diags, e)
			}
			obj.IdpSingleLogoutUrl = &v2
		}
	}
	if v1, ok := d.GetOk("idp_single_sign_on_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idp_single_sign_on_url", sv)
				diags = append(diags, e)
			}
			obj.IdpSingleSignOnUrl = &v2
		}
	}
	if v1, ok := d.GetOk("limit_relaystate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("limit_relaystate", sv)
				diags = append(diags, e)
			}
			obj.LimitRelaystate = &v2
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
	if v1, ok := d.GetOk("single_logout_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("single_logout_url", sv)
				diags = append(diags, e)
			}
			obj.SingleLogoutUrl = &v2
		}
	}
	if v1, ok := d.GetOk("single_sign_on_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("single_sign_on_url", sv)
				diags = append(diags, e)
			}
			obj.SingleSignOnUrl = &v2
		}
	}
	if v1, ok := d.GetOk("user_claim_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("user_claim_type", sv)
				diags = append(diags, e)
			}
			obj.UserClaimType = &v2
		}
	}
	if v1, ok := d.GetOk("user_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_name", sv)
				diags = append(diags, e)
			}
			obj.UserName = &v2
		}
	}
	return &obj, diags
}
