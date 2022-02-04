// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure 3GPP public land mobile network (PLMN).

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

func resourceWirelessControllerHotspot20Anqp3gppCellular() *schema.Resource {
	return &schema.Resource{
		Description: "Configure 3GPP public land mobile network (PLMN).",

		CreateContext: resourceWirelessControllerHotspot20Anqp3gppCellularCreate,
		ReadContext:   resourceWirelessControllerHotspot20Anqp3gppCellularRead,
		UpdateContext: resourceWirelessControllerHotspot20Anqp3gppCellularUpdate,
		DeleteContext: resourceWirelessControllerHotspot20Anqp3gppCellularDelete,

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
			"mcc_mnc_list": {
				Type:        schema.TypeList,
				Description: "Mobile Country Code and Mobile Network Code configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 6),

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"mcc": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 3),

							Description: "Mobile country code.",
							Optional:    true,
							Computed:    true,
						},
						"mnc": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 3),

							Description: "Mobile network code.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "3GPP PLMN name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20Anqp3gppCellularCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerHotspot20Anqp3gppCellular resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerHotspot20Anqp3gppCellular(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerHotspot20Anqp3gppCellular(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20Anqp3gppCellular")
	}

	return resourceWirelessControllerHotspot20Anqp3gppCellularRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20Anqp3gppCellularUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerHotspot20Anqp3gppCellular(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerHotspot20Anqp3gppCellular(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerHotspot20Anqp3gppCellular resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20Anqp3gppCellular")
	}

	return resourceWirelessControllerHotspot20Anqp3gppCellularRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20Anqp3gppCellularDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerHotspot20Anqp3gppCellular(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerHotspot20Anqp3gppCellular resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20Anqp3gppCellularRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20Anqp3gppCellular(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20Anqp3gppCellular resource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20Anqp3gppCellular(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerHotspot20Anqp3gppCellularMccMncList(v *[]models.WirelessControllerHotspot20Anqp3gppCellularMccMncList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Mcc; tmp != nil {
				v["mcc"] = *tmp
			}

			if tmp := cfg.Mnc; tmp != nil {
				v["mnc"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectWirelessControllerHotspot20Anqp3gppCellular(d *schema.ResourceData, o *models.WirelessControllerHotspot20Anqp3gppCellular, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.MccMncList != nil {
		if err = d.Set("mcc_mnc_list", flattenWirelessControllerHotspot20Anqp3gppCellularMccMncList(o.MccMncList, sort)); err != nil {
			return diag.Errorf("error reading mcc_mnc_list: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerHotspot20Anqp3gppCellularMccMncList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerHotspot20Anqp3gppCellularMccMncList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerHotspot20Anqp3gppCellularMccMncList

	for i := range l {
		tmp := models.WirelessControllerHotspot20Anqp3gppCellularMccMncList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mcc", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mcc = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mnc", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mnc = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerHotspot20Anqp3gppCellular(d *schema.ResourceData, sv string) (*models.WirelessControllerHotspot20Anqp3gppCellular, diag.Diagnostics) {
	obj := models.WirelessControllerHotspot20Anqp3gppCellular{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("mcc_mnc_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("mcc_mnc_list", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerHotspot20Anqp3gppCellularMccMncList(d, v, "mcc_mnc_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.MccMncList = t
		}
	} else if d.HasChange("mcc_mnc_list") {
		old, new := d.GetChange("mcc_mnc_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.MccMncList = &[]models.WirelessControllerHotspot20Anqp3gppCellularMccMncList{}
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
	return &obj, diags
}
