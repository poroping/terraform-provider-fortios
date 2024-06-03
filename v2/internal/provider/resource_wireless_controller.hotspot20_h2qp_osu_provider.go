// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure online sign up (OSU) provider list.

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

func resourceWirelessControllerHotspot20H2qpOsuProvider() *schema.Resource {
	return &schema.Resource{
		Description: "Configure online sign up (OSU) provider list.",

		CreateContext: resourceWirelessControllerHotspot20H2qpOsuProviderCreate,
		ReadContext:   resourceWirelessControllerHotspot20H2qpOsuProviderRead,
		UpdateContext: resourceWirelessControllerHotspot20H2qpOsuProviderUpdate,
		DeleteContext: resourceWirelessControllerHotspot20H2qpOsuProviderDelete,

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
			"friendly_name": {
				Type:        schema.TypeList,
				Description: "OSU provider friendly name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"friendly_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 252),

							Description: "OSU provider friendly name.",
							Optional:    true,
							Computed:    true,
						},
						"index": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),

							Description: "OSU provider friendly name index.",
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
					},
				},
			},
			"icon": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "OSU provider icon.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "OSU provider ID.",
				ForceNew:    true,
				Required:    true,
			},
			"osu_method": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "OSU method list.",
				Optional:         true,
				Computed:         true,
			},
			"osu_nai": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "OSU NAI.",
				Optional:    true,
				Computed:    true,
			},
			"server_uri": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Server URI.",
				Optional:    true,
				Computed:    true,
			},
			"service_description": {
				Type:        schema.TypeList,
				Description: "OSU service name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lang": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 3),

							Description: "Language code.",
							Optional:    true,
							Computed:    true,
						},
						"service_description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 252),

							Description: "Service description.",
							Optional:    true,
							Computed:    true,
						},
						"service_id": {
							Type: schema.TypeInt,

							Description: "OSU service ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessControllerHotspot20H2qpOsuProviderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerHotspot20H2qpOsuProvider resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerHotspot20H2qpOsuProvider(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerHotspot20H2qpOsuProvider(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20H2qpOsuProvider")
	}

	return resourceWirelessControllerHotspot20H2qpOsuProviderRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20H2qpOsuProviderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerHotspot20H2qpOsuProvider(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerHotspot20H2qpOsuProvider(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerHotspot20H2qpOsuProvider resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20H2qpOsuProvider")
	}

	return resourceWirelessControllerHotspot20H2qpOsuProviderRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20H2qpOsuProviderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerHotspot20H2qpOsuProvider(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerHotspot20H2qpOsuProvider resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2qpOsuProviderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20H2qpOsuProvider(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20H2qpOsuProvider resource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20H2qpOsuProvider(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerHotspot20H2qpOsuProviderFriendlyName(d *schema.ResourceData, v *[]models.WirelessControllerHotspot20H2qpOsuProviderFriendlyName, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.FriendlyName; tmp != nil {
				v["friendly_name"] = *tmp
			}

			if tmp := cfg.Index; tmp != nil {
				v["index"] = *tmp
			}

			if tmp := cfg.Lang; tmp != nil {
				v["lang"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "index")
	}

	return flat
}

func flattenWirelessControllerHotspot20H2qpOsuProviderServiceDescription(d *schema.ResourceData, v *[]models.WirelessControllerHotspot20H2qpOsuProviderServiceDescription, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Lang; tmp != nil {
				v["lang"] = *tmp
			}

			if tmp := cfg.ServiceDescription; tmp != nil {
				v["service_description"] = *tmp
			}

			if tmp := cfg.ServiceId; tmp != nil {
				v["service_id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "service_id")
	}

	return flat
}

func refreshObjectWirelessControllerHotspot20H2qpOsuProvider(d *schema.ResourceData, o *models.WirelessControllerHotspot20H2qpOsuProvider, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.FriendlyName != nil {
		if err = d.Set("friendly_name", flattenWirelessControllerHotspot20H2qpOsuProviderFriendlyName(d, o.FriendlyName, "friendly_name", sort)); err != nil {
			return diag.Errorf("error reading friendly_name: %v", err)
		}
	}

	if o.Icon != nil {
		v := *o.Icon

		if err = d.Set("icon", v); err != nil {
			return diag.Errorf("error reading icon: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.OsuMethod != nil {
		v := *o.OsuMethod

		if err = d.Set("osu_method", v); err != nil {
			return diag.Errorf("error reading osu_method: %v", err)
		}
	}

	if o.OsuNai != nil {
		v := *o.OsuNai

		if err = d.Set("osu_nai", v); err != nil {
			return diag.Errorf("error reading osu_nai: %v", err)
		}
	}

	if o.ServerUri != nil {
		v := *o.ServerUri

		if err = d.Set("server_uri", v); err != nil {
			return diag.Errorf("error reading server_uri: %v", err)
		}
	}

	if o.ServiceDescription != nil {
		if err = d.Set("service_description", flattenWirelessControllerHotspot20H2qpOsuProviderServiceDescription(d, o.ServiceDescription, "service_description", sort)); err != nil {
			return diag.Errorf("error reading service_description: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerHotspot20H2qpOsuProviderFriendlyName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerHotspot20H2qpOsuProviderFriendlyName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerHotspot20H2qpOsuProviderFriendlyName

	for i := range l {
		tmp := models.WirelessControllerHotspot20H2qpOsuProviderFriendlyName{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.friendly_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FriendlyName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.index", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Index = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.lang", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Lang = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerHotspot20H2qpOsuProviderServiceDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerHotspot20H2qpOsuProviderServiceDescription, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerHotspot20H2qpOsuProviderServiceDescription

	for i := range l {
		tmp := models.WirelessControllerHotspot20H2qpOsuProviderServiceDescription{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.lang", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Lang = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.service_description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ServiceDescription = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.service_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ServiceId = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerHotspot20H2qpOsuProvider(d *schema.ResourceData, sv string) (*models.WirelessControllerHotspot20H2qpOsuProvider, diag.Diagnostics) {
	obj := models.WirelessControllerHotspot20H2qpOsuProvider{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("friendly_name"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("friendly_name", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerHotspot20H2qpOsuProviderFriendlyName(d, v, "friendly_name", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FriendlyName = t
		}
	} else if d.HasChange("friendly_name") {
		old, new := d.GetChange("friendly_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FriendlyName = &[]models.WirelessControllerHotspot20H2qpOsuProviderFriendlyName{}
		}
	}
	if v1, ok := d.GetOk("icon"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("icon", sv)
				diags = append(diags, e)
			}
			obj.Icon = &v2
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
	if v1, ok := d.GetOk("osu_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("osu_method", sv)
				diags = append(diags, e)
			}
			obj.OsuMethod = &v2
		}
	}
	if v1, ok := d.GetOk("osu_nai"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("osu_nai", sv)
				diags = append(diags, e)
			}
			obj.OsuNai = &v2
		}
	}
	if v1, ok := d.GetOk("server_uri"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_uri", sv)
				diags = append(diags, e)
			}
			obj.ServerUri = &v2
		}
	}
	if v, ok := d.GetOk("service_description"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("service_description", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerHotspot20H2qpOsuProviderServiceDescription(d, v, "service_description", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ServiceDescription = t
		}
	} else if d.HasChange("service_description") {
		old, new := d.GetChange("service_description")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ServiceDescription = &[]models.WirelessControllerHotspot20H2qpOsuProviderServiceDescription{}
		}
	}
	return &obj, diags
}
