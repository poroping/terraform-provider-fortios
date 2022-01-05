// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WTP groups.

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

func resourceWirelessControllerWtpGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WTP groups.",

		CreateContext: resourceWirelessControllerWtpGroupCreate,
		ReadContext:   resourceWirelessControllerWtpGroupRead,
		UpdateContext: resourceWirelessControllerWtpGroupUpdate,
		DeleteContext: resourceWirelessControllerWtpGroupDelete,

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
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WTP group name.",
				ForceNew:    true,
				Required:    true,
			},
			"platform_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"AP-11N", "220B", "210B", "222B", "112B", "320B", "11C", "14C", "223B", "28C", "320C", "221C", "25D", "222C", "224D", "214B", "21D", "24D", "112D", "223C", "321C", "C220C", "C225C", "C23JD", "C24JE", "S321C", "S322C", "S323C", "S311C", "S313C", "S321CR", "S322CR", "S323CR", "S421E", "S422E", "S423E", "421E", "423E", "221E", "222E", "223E", "224E", "231E", "S221E", "S223E", "321E", "431F", "432F", "433F", "231F", "234F", "23JF", "831F", "U421E", "U422EV", "U423E", "U221EV", "U223EV", "U24JEV", "U321EV", "U323EV", "U431F", "U433F", "U231F", "U234F", "U432F"}, false),

				Description: "FortiAP models to define the WTP group platform type.",
				Optional:    true,
				Computed:    true,
			},
			"wtps": {
				Type:        schema.TypeList,
				Description: "WTP list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"wtp_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "WTP ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessControllerWtpGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerWtpGroup resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerWtpGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerWtpGroup(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerWtpGroup")
	}

	return resourceWirelessControllerWtpGroupRead(ctx, d, meta)
}

func resourceWirelessControllerWtpGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerWtpGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerWtpGroup(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerWtpGroup resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerWtpGroup")
	}

	return resourceWirelessControllerWtpGroupRead(ctx, d, meta)
}

func resourceWirelessControllerWtpGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerWtpGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerWtpGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerWtpGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerWtpGroup resource: %v", err)
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

	diags := refreshObjectWirelessControllerWtpGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerWtpGroupWtps(v *[]models.WirelessControllerWtpGroupWtps, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.WtpId; tmp != nil {
				v["wtp_id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "wtp_id")
	}

	return flat
}

func refreshObjectWirelessControllerWtpGroup(d *schema.ResourceData, o *models.WirelessControllerWtpGroup, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PlatformType != nil {
		v := *o.PlatformType

		if err = d.Set("platform_type", v); err != nil {
			return diag.Errorf("error reading platform_type: %v", err)
		}
	}

	if o.Wtps != nil {
		if err = d.Set("wtps", flattenWirelessControllerWtpGroupWtps(o.Wtps, sort)); err != nil {
			return diag.Errorf("error reading wtps: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerWtpGroupWtps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpGroupWtps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpGroupWtps

	for i := range l {
		tmp := models.WirelessControllerWtpGroupWtps{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.wtp_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.WtpId = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerWtpGroup(d *schema.ResourceData, sv string) (*models.WirelessControllerWtpGroup, diag.Diagnostics) {
	obj := models.WirelessControllerWtpGroup{}
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
	if v1, ok := d.GetOk("platform_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("platform_type", sv)
				diags = append(diags, e)
			}
			obj.PlatformType = &v2
		}
	}
	if v, ok := d.GetOk("wtps"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("wtps", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpGroupWtps(d, v, "wtps", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Wtps = t
		}
	} else if d.HasChange("wtps") {
		old, new := d.GetChange("wtps")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Wtps = &[]models.WirelessControllerWtpGroupWtps{}
		}
	}
	return &obj, diags
}
