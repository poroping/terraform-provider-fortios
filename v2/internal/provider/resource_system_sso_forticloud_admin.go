// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiCloud SSO admin users.

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

func resourceSystemSsoForticloudAdmin() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiCloud SSO admin users.",

		CreateContext: resourceSystemSsoForticloudAdminCreate,
		ReadContext:   resourceSystemSsoForticloudAdminRead,
		UpdateContext: resourceSystemSsoForticloudAdminUpdate,
		DeleteContext: resourceSystemSsoForticloudAdminDelete,

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
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "FortiCloud SSO admin name.",
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

func resourceSystemSsoForticloudAdminCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemSsoForticloudAdmin resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemSsoForticloudAdmin(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSsoForticloudAdmin(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSsoForticloudAdmin")
	}

	return resourceSystemSsoForticloudAdminRead(ctx, d, meta)
}

func resourceSystemSsoForticloudAdminUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSsoForticloudAdmin(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSsoForticloudAdmin(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSsoForticloudAdmin resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSsoForticloudAdmin")
	}

	return resourceSystemSsoForticloudAdminRead(ctx, d, meta)
}

func resourceSystemSsoForticloudAdminDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSsoForticloudAdmin(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSsoForticloudAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSsoForticloudAdminRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSsoForticloudAdmin(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSsoForticloudAdmin resource: %v", err)
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

	diags := refreshObjectSystemSsoForticloudAdmin(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemSsoForticloudAdminVdom(v *[]models.SystemSsoForticloudAdminVdom, sort bool) interface{} {
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

func refreshObjectSystemSsoForticloudAdmin(d *schema.ResourceData, o *models.SystemSsoForticloudAdmin, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Vdom != nil {
		if err = d.Set("vdom", flattenSystemSsoForticloudAdminVdom(o.Vdom, sort)); err != nil {
			return diag.Errorf("error reading vdom: %v", err)
		}
	}

	return nil
}

func expandSystemSsoForticloudAdminVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSsoForticloudAdminVdom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSsoForticloudAdminVdom

	for i := range l {
		tmp := models.SystemSsoForticloudAdminVdom{}
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

func getObjectSystemSsoForticloudAdmin(d *schema.ResourceData, sv string) (*models.SystemSsoForticloudAdmin, diag.Diagnostics) {
	obj := models.SystemSsoForticloudAdmin{}
	diags := diag.Diagnostics{}

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
		t, err := expandSystemSsoForticloudAdminVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Vdom = t
		}
	} else if d.HasChange("vdom") {
		old, new := d.GetChange("vdom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Vdom = &[]models.SystemSsoForticloudAdminVdom{}
		}
	}
	return &obj, diags
}
