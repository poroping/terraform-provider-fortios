// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Kerberos keytab entries.

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

func resourceUserKrbKeytab() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Kerberos keytab entries.",

		CreateContext: resourceUserKrbKeytabCreate,
		ReadContext:   resourceUserKrbKeytabRead,
		UpdateContext: resourceUserKrbKeytabUpdate,
		DeleteContext: resourceUserKrbKeytabDelete,

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
			"keytab": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 8191),

				Description: "base64 coded keytab file containing a pre-shared key.",
				Optional:    true,
				Computed:    true,
			},
			"ldap_server": {
				Type:        schema.TypeList,
				Description: "LDAP server name(s).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "LDAP server name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Kerberos keytab entry name.",
				ForceNew:    true,
				Required:    true,
			},
			"pac_data": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable parsing PAC data in the ticket.",
				Optional:    true,
				Computed:    true,
			},
			"principal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),

				Description: "Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserKrbKeytabCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserKrbKeytab resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserKrbKeytab(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserKrbKeytab(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserKrbKeytab")
	}

	return resourceUserKrbKeytabRead(ctx, d, meta)
}

func resourceUserKrbKeytabUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserKrbKeytab(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserKrbKeytab(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserKrbKeytab resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserKrbKeytab")
	}

	return resourceUserKrbKeytabRead(ctx, d, meta)
}

func resourceUserKrbKeytabDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserKrbKeytab(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserKrbKeytab resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserKrbKeytabRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserKrbKeytab(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserKrbKeytab resource: %v", err)
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

	diags := refreshObjectUserKrbKeytab(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenUserKrbKeytabLdapServer(v *[]models.UserKrbKeytabLdapServer, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func refreshObjectUserKrbKeytab(d *schema.ResourceData, o *models.UserKrbKeytab, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Keytab != nil {
		v := *o.Keytab

		if err = d.Set("keytab", v); err != nil {
			return diag.Errorf("error reading keytab: %v", err)
		}
	}

	if o.LdapServer != nil {
		if err = d.Set("ldap_server", flattenUserKrbKeytabLdapServer(o.LdapServer, sort)); err != nil {
			return diag.Errorf("error reading ldap_server: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PacData != nil {
		v := *o.PacData

		if err = d.Set("pac_data", v); err != nil {
			return diag.Errorf("error reading pac_data: %v", err)
		}
	}

	if o.Principal != nil {
		v := *o.Principal

		if err = d.Set("principal", v); err != nil {
			return diag.Errorf("error reading principal: %v", err)
		}
	}

	return nil
}

func expandUserKrbKeytabLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserKrbKeytabLdapServer, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserKrbKeytabLdapServer

	for i := range l {
		tmp := models.UserKrbKeytabLdapServer{}
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

func getObjectUserKrbKeytab(d *schema.ResourceData, sv string) (*models.UserKrbKeytab, diag.Diagnostics) {
	obj := models.UserKrbKeytab{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("keytab"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keytab", sv)
				diags = append(diags, e)
			}
			obj.Keytab = &v2
		}
	}
	if v, ok := d.GetOk("ldap_server"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ldap_server", sv)
			diags = append(diags, e)
		}
		t, err := expandUserKrbKeytabLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.LdapServer = t
		}
	} else if d.HasChange("ldap_server") {
		old, new := d.GetChange("ldap_server")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.LdapServer = &[]models.UserKrbKeytabLdapServer{}
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
	if v1, ok := d.GetOk("pac_data"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pac_data", sv)
				diags = append(diags, e)
			}
			obj.PacData = &v2
		}
	}
	if v1, ok := d.GetOk("principal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("principal", sv)
				diags = append(diags, e)
			}
			obj.Principal = &v2
		}
	}
	return &obj, diags
}
