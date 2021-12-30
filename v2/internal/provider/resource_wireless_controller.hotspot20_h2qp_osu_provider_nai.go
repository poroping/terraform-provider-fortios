// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure online sign up (OSU) provider NAI list.

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

func resourceWirelessControllerhotspot20H2qpOsuProviderNai() *schema.Resource {
	return &schema.Resource{
		Description: "Configure online sign up (OSU) provider NAI list.",

		CreateContext: resourceWirelessControllerhotspot20H2qpOsuProviderNaiCreate,
		ReadContext:   resourceWirelessControllerhotspot20H2qpOsuProviderNaiRead,
		UpdateContext: resourceWirelessControllerhotspot20H2qpOsuProviderNaiUpdate,
		DeleteContext: resourceWirelessControllerhotspot20H2qpOsuProviderNaiDelete,

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
			"nai_list": {
				Type:        schema.TypeList,
				Description: "OSU NAI list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "OSU NAI ID.",
							Optional:    true,
							Computed:    true,
						},
						"osu_nai": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "OSU NAI.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "OSU provider NAI ID.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceWirelessControllerhotspot20H2qpOsuProviderNaiCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerhotspot20H2qpOsuProviderNai resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerhotspot20H2qpOsuProviderNai(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerhotspot20H2qpOsuProviderNai(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerhotspot20H2qpOsuProviderNai")
	}

	return resourceWirelessControllerhotspot20H2qpOsuProviderNaiRead(ctx, d, meta)
}

func resourceWirelessControllerhotspot20H2qpOsuProviderNaiUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerhotspot20H2qpOsuProviderNai(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerhotspot20H2qpOsuProviderNai(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerhotspot20H2qpOsuProviderNai resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerhotspot20H2qpOsuProviderNai")
	}

	return resourceWirelessControllerhotspot20H2qpOsuProviderNaiRead(ctx, d, meta)
}

func resourceWirelessControllerhotspot20H2qpOsuProviderNaiDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerhotspot20H2qpOsuProviderNai(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerhotspot20H2qpOsuProviderNai resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerhotspot20H2qpOsuProviderNaiRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerhotspot20H2qpOsuProviderNai(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerhotspot20H2qpOsuProviderNai resource: %v", err)
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

	diags := refreshObjectWirelessControllerhotspot20H2qpOsuProviderNai(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerhotspot20H2qpOsuProviderNaiNaiList(v *[]models.WirelessControllerhotspot20H2qpOsuProviderNaiNaiList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.OsuNai; tmp != nil {
				v["osu_nai"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectWirelessControllerhotspot20H2qpOsuProviderNai(d *schema.ResourceData, o *models.WirelessControllerhotspot20H2qpOsuProviderNai, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.NaiList != nil {
		if err = d.Set("nai_list", flattenWirelessControllerhotspot20H2qpOsuProviderNaiNaiList(o.NaiList, sort)); err != nil {
			return diag.Errorf("error reading nai_list: %v", err)
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

func expandWirelessControllerhotspot20H2qpOsuProviderNaiNaiList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerhotspot20H2qpOsuProviderNaiNaiList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerhotspot20H2qpOsuProviderNaiNaiList

	for i := range l {
		tmp := models.WirelessControllerhotspot20H2qpOsuProviderNaiNaiList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.osu_nai", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OsuNai = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerhotspot20H2qpOsuProviderNai(d *schema.ResourceData, sv string) (*models.WirelessControllerhotspot20H2qpOsuProviderNai, diag.Diagnostics) {
	obj := models.WirelessControllerhotspot20H2qpOsuProviderNai{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("nai_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("nai_list", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerhotspot20H2qpOsuProviderNaiNaiList(d, v, "nai_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NaiList = t
		}
	} else if d.HasChange("nai_list") {
		old, new := d.GetChange("nai_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NaiList = &[]models.WirelessControllerhotspot20H2qpOsuProviderNaiNaiList{}
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
