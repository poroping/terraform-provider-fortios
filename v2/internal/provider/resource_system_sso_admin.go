// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SSO admin users.

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

func resourceSystemSsoAdmin() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SSO admin users.",

		CreateContext: resourceSystemSsoAdminCreate,
		ReadContext:   resourceSystemSsoAdminRead,
		UpdateContext: resourceSystemSsoAdminUpdate,
		DeleteContext: resourceSystemSsoAdminDelete,

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
			"accprofile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SSO admin user access profile.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "SSO admin name.",
				ForceNew:    true,
				Required:    true,
			},
			"vdom": {
				Type:        schema.TypeList,
				Description: "Virtual domain(s) that the administrator can access.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Virtual domain name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemSsoAdminCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemSsoAdmin resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemSsoAdmin(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSsoAdmin(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSsoAdmin")
	}

	return resourceSystemSsoAdminRead(ctx, d, meta)
}

func resourceSystemSsoAdminUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSsoAdmin(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSsoAdmin(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSsoAdmin resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSsoAdmin")
	}

	return resourceSystemSsoAdminRead(ctx, d, meta)
}

func resourceSystemSsoAdminDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSsoAdmin(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSsoAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSsoAdminRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSsoAdmin(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSsoAdmin resource: %v", err)
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

	diags := refreshObjectSystemSsoAdmin(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemSsoAdminVdom(d *schema.ResourceData, v *[]models.SystemSsoAdminVdom, prefix string, sort bool) interface{} {
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

func refreshObjectSystemSsoAdmin(d *schema.ResourceData, o *models.SystemSsoAdmin, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Accprofile != nil {
		v := *o.Accprofile

		if err = d.Set("accprofile", v); err != nil {
			return diag.Errorf("error reading accprofile: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Vdom != nil {
		if err = d.Set("vdom", flattenSystemSsoAdminVdom(d, o.Vdom, "vdom", sort)); err != nil {
			return diag.Errorf("error reading vdom: %v", err)
		}
	}

	return nil
}

func expandSystemSsoAdminVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSsoAdminVdom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSsoAdminVdom

	for i := range l {
		tmp := models.SystemSsoAdminVdom{}
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

func getObjectSystemSsoAdmin(d *schema.ResourceData, sv string) (*models.SystemSsoAdmin, diag.Diagnostics) {
	obj := models.SystemSsoAdmin{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("accprofile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("accprofile", sv)
				diags = append(diags, e)
			}
			obj.Accprofile = &v2
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
	if v, ok := d.GetOk("vdom"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("vdom", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSsoAdminVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Vdom = t
		}
	} else if d.HasChange("vdom") {
		old, new := d.GetChange("vdom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Vdom = &[]models.SystemSsoAdminVdom{}
		}
	}
	return &obj, diags
}
