// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure venue name duple.

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

func resourceWirelessControllerhotspot20AnqpVenueName() *schema.Resource {
	return &schema.Resource{
		Description: "Configure venue name duple.",

		CreateContext: resourceWirelessControllerhotspot20AnqpVenueNameCreate,
		ReadContext:   resourceWirelessControllerhotspot20AnqpVenueNameRead,
		UpdateContext: resourceWirelessControllerhotspot20AnqpVenueNameUpdate,
		DeleteContext: resourceWirelessControllerhotspot20AnqpVenueNameDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of venue name duple.",
				ForceNew:    true,
				Required:    true,
			},
			"value_list": {
				Type:        schema.TypeList,
				Description: "Name list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),

							Description: "Value index.",
							Optional:    true,
							Computed:    true,
						},
						"lang": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 3),

							Description: "Language code.",
							Optional:    true,
							Computed:    true,
						},
						"value": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 252),

							Description: "Venue name value.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessControllerhotspot20AnqpVenueNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerhotspot20AnqpVenueName resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerhotspot20AnqpVenueName(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerhotspot20AnqpVenueName(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerhotspot20AnqpVenueName")
	}

	return resourceWirelessControllerhotspot20AnqpVenueNameRead(ctx, d, meta)
}

func resourceWirelessControllerhotspot20AnqpVenueNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerhotspot20AnqpVenueName(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerhotspot20AnqpVenueName(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerhotspot20AnqpVenueName resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerhotspot20AnqpVenueName")
	}

	return resourceWirelessControllerhotspot20AnqpVenueNameRead(ctx, d, meta)
}

func resourceWirelessControllerhotspot20AnqpVenueNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerhotspot20AnqpVenueName(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerhotspot20AnqpVenueName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerhotspot20AnqpVenueNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerhotspot20AnqpVenueName(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerhotspot20AnqpVenueName resource: %v", err)
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

	diags := refreshObjectWirelessControllerhotspot20AnqpVenueName(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerhotspot20AnqpVenueNameValueList(v *[]models.WirelessControllerhotspot20AnqpVenueNameValueList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Index; tmp != nil {
				v["index"] = *tmp
			}

			if tmp := cfg.Lang; tmp != nil {
				v["lang"] = *tmp
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "index")
	}

	return flat
}

func refreshObjectWirelessControllerhotspot20AnqpVenueName(d *schema.ResourceData, o *models.WirelessControllerhotspot20AnqpVenueName, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ValueList != nil {
		if err = d.Set("value_list", flattenWirelessControllerhotspot20AnqpVenueNameValueList(o.ValueList, sort)); err != nil {
			return diag.Errorf("error reading value_list: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerhotspot20AnqpVenueNameValueList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerhotspot20AnqpVenueNameValueList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerhotspot20AnqpVenueNameValueList

	for i := range l {
		tmp := models.WirelessControllerhotspot20AnqpVenueNameValueList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.index", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Index = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.lang", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Lang = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerhotspot20AnqpVenueName(d *schema.ResourceData, sv string) (*models.WirelessControllerhotspot20AnqpVenueName, diag.Diagnostics) {
	obj := models.WirelessControllerhotspot20AnqpVenueName{}
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
	if v, ok := d.GetOk("value_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("value_list", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerhotspot20AnqpVenueNameValueList(d, v, "value_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ValueList = t
		}
	} else if d.HasChange("value_list") {
		old, new := d.GetChange("value_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ValueList = &[]models.WirelessControllerhotspot20AnqpVenueNameValueList{}
		}
	}
	return &obj, diags
}
