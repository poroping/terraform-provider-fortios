// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Global configuration objects that can be configured independently across different ha peers for all VDOMs or for the defined VDOM scope.

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

func resourceSystemVdomException() *schema.Resource {
	return &schema.Resource{
		Description: "Global configuration objects that can be configured independently across different ha peers for all VDOMs or for the defined VDOM scope.",

		CreateContext: resourceSystemVdomExceptionCreate,
		ReadContext:   resourceSystemVdomExceptionRead,
		UpdateContext: resourceSystemVdomExceptionUpdate,
		DeleteContext: resourceSystemVdomExceptionDelete,

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
				Type:         schema.TypeBool,
				Description:  "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fosid": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4096),

				Description: "Index <1-4096>.",
				Optional:    true,
				Computed:    true,
			},
			"object": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"log.fortianalyzer.setting", "log.fortianalyzer.override-setting", "log.fortianalyzer2.setting", "log.fortianalyzer2.override-setting", "log.fortianalyzer3.setting", "log.fortianalyzer3.override-setting", "log.fortianalyzer-cloud.setting", "log.fortianalyzer-cloud.override-setting", "log.syslogd.setting", "log.syslogd.override-setting", "log.syslogd2.setting", "log.syslogd2.override-setting", "log.syslogd3.setting", "log.syslogd3.override-setting", "log.syslogd4.setting", "log.syslogd4.override-setting", "system.gre-tunnel", "system.central-management", "system.csf", "user.radius", "system.cluster-sync", "system.standalone-cluster", "system.interface", "vpn.ipsec.phase1-interface", "vpn.ipsec.phase2-interface", "router.bgp", "router.route-map", "router.prefix-list", "firewall.ippool", "firewall.ippool6", "router.static", "router.static6", "firewall.vip", "firewall.vip6", "system.sdwan", "system.saml", "router.policy", "router.policy6"}, false),

				Description: "Name of the configuration object that can be configured independently for all VDOMs.",
				Optional:    true,
				Computed:    true,
			},
			"scope": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"all", "inclusive", "exclusive"}, false),

				Description: "Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration.",
				Optional:    true,
				Computed:    true,
			},
			"vdom": {
				Type:        schema.TypeList,
				Description: "Names of the VDOMs.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "VDOM name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemVdomExceptionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemVdomException resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemVdomException(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemVdomException(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVdomException")
	}

	return resourceSystemVdomExceptionRead(ctx, d, meta)
}

func resourceSystemVdomExceptionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVdomException(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemVdomException(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemVdomException resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVdomException")
	}

	return resourceSystemVdomExceptionRead(ctx, d, meta)
}

func resourceSystemVdomExceptionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemVdomException(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemVdomException resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomExceptionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVdomException(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVdomException resource: %v", err)
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

	diags := refreshObjectSystemVdomException(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemVdomExceptionVdom(v *[]models.SystemVdomExceptionVdom, sort bool) interface{} {
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

func refreshObjectSystemVdomException(d *schema.ResourceData, o *models.SystemVdomException, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Fosid != nil {
		v := *o.Fosid

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Object != nil {
		v := *o.Object

		if err = d.Set("object", v); err != nil {
			return diag.Errorf("error reading object: %v", err)
		}
	}

	if o.Scope != nil {
		v := *o.Scope

		if err = d.Set("scope", v); err != nil {
			return diag.Errorf("error reading scope: %v", err)
		}
	}

	if o.Vdom != nil {
		if err = d.Set("vdom", flattenSystemVdomExceptionVdom(o.Vdom, sort)); err != nil {
			return diag.Errorf("error reading vdom: %v", err)
		}
	}

	return nil
}

func expandSystemVdomExceptionVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVdomExceptionVdom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVdomExceptionVdom

	for i := range l {
		tmp := models.SystemVdomExceptionVdom{}
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

func getObjectSystemVdomException(d *schema.ResourceData, sv string) (*models.SystemVdomException, diag.Diagnostics) {
	obj := models.SystemVdomException{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Fosid = &tmp
		}
	}
	if v1, ok := d.GetOk("object"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("object", sv)
				diags = append(diags, e)
			}
			obj.Object = &v2
		}
	}
	if v1, ok := d.GetOk("scope"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scope", sv)
				diags = append(diags, e)
			}
			obj.Scope = &v2
		}
	}
	if v, ok := d.GetOk("vdom"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("vdom", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemVdomExceptionVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Vdom = t
		}
	} else if d.HasChange("vdom") {
		old, new := d.GetChange("vdom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Vdom = &[]models.SystemVdomExceptionVdom{}
		}
	}
	return &obj, diags
}
