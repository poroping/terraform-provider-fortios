// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SSL-VPN host check software.

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

func resourceVpnSslWebHostCheckSoftware() *schema.Resource {
	return &schema.Resource{
		Description: "SSL-VPN host check software.",

		CreateContext: resourceVpnSslWebHostCheckSoftwareCreate,
		ReadContext:   resourceVpnSslWebHostCheckSoftwareRead,
		UpdateContext: resourceVpnSslWebHostCheckSoftwareUpdate,
		DeleteContext: resourceVpnSslWebHostCheckSoftwareDelete,

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
			"check_item_list": {
				Type:        schema.TypeList,
				Description: "Check item list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"require", "deny"}, false),

							Description: "Action.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "ID (0 - 4294967295).",
							Optional:    true,
							Computed:    true,
						},
						"md5s": {
							Type:        schema.TypeList,
							Description: "MD5 checksum.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 32),

										Description: "Hex string of MD5 checksum.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"target": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Target.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"file", "registry", "process"}, false),

							Description: "Type.",
							Optional:    true,
							Computed:    true,
						},
						"version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Version.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"guid": {
				Type: schema.TypeString,

				Description: "Globally unique ID.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"os_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"windows", "macos"}, false),

				Description: "OS type.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"av", "fw"}, false),

				Description: "Type.",
				Optional:    true,
				Computed:    true,
			},
			"version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Version.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpnSslWebHostCheckSoftwareCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating VpnSslWebHostCheckSoftware resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnSslWebHostCheckSoftware(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnSslWebHostCheckSoftware(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnSslWebHostCheckSoftware")
	}

	return resourceVpnSslWebHostCheckSoftwareRead(ctx, d, meta)
}

func resourceVpnSslWebHostCheckSoftwareUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnSslWebHostCheckSoftware(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnSslWebHostCheckSoftware(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnSslWebHostCheckSoftware resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnSslWebHostCheckSoftware")
	}

	return resourceVpnSslWebHostCheckSoftwareRead(ctx, d, meta)
}

func resourceVpnSslWebHostCheckSoftwareDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVpnSslWebHostCheckSoftware(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnSslWebHostCheckSoftware resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebHostCheckSoftwareRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnSslWebHostCheckSoftware(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnSslWebHostCheckSoftware resource: %v", err)
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

	diags := refreshObjectVpnSslWebHostCheckSoftware(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenVpnSslWebHostCheckSoftwareCheckItemList(v *[]models.VpnSslWebHostCheckSoftwareCheckItemList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Md5s; tmp != nil {
				v["md5s"] = flattenVpnSslWebHostCheckSoftwareCheckItemListMd5s(tmp, sort)
			}

			if tmp := cfg.Target; tmp != nil {
				v["target"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.Version; tmp != nil {
				v["version"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenVpnSslWebHostCheckSoftwareCheckItemListMd5s(v *[]models.VpnSslWebHostCheckSoftwareCheckItemListMd5s, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectVpnSslWebHostCheckSoftware(d *schema.ResourceData, o *models.VpnSslWebHostCheckSoftware, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CheckItemList != nil {
		if err = d.Set("check_item_list", flattenVpnSslWebHostCheckSoftwareCheckItemList(o.CheckItemList, sort)); err != nil {
			return diag.Errorf("error reading check_item_list: %v", err)
		}
	}

	if o.Guid != nil {
		v := *o.Guid

		if err = d.Set("guid", v); err != nil {
			return diag.Errorf("error reading guid: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.OsType != nil {
		v := *o.OsType

		if err = d.Set("os_type", v); err != nil {
			return diag.Errorf("error reading os_type: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.Version != nil {
		v := *o.Version

		if err = d.Set("version", v); err != nil {
			return diag.Errorf("error reading version: %v", err)
		}
	}

	return nil
}

func expandVpnSslWebHostCheckSoftwareCheckItemList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebHostCheckSoftwareCheckItemList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebHostCheckSoftwareCheckItemList

	for i := range l {
		tmp := models.VpnSslWebHostCheckSoftwareCheckItemList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.md5s", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnSslWebHostCheckSoftwareCheckItemListMd5s(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnSslWebHostCheckSoftwareCheckItemListMd5s
			// 	}
			tmp.Md5s = v2

		}

		pre_append = fmt.Sprintf("%s.%d.target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Target = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Version = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnSslWebHostCheckSoftwareCheckItemListMd5s(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebHostCheckSoftwareCheckItemListMd5s, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebHostCheckSoftwareCheckItemListMd5s

	for i := range l {
		tmp := models.VpnSslWebHostCheckSoftwareCheckItemListMd5s{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectVpnSslWebHostCheckSoftware(d *schema.ResourceData, sv string) (*models.VpnSslWebHostCheckSoftware, diag.Diagnostics) {
	obj := models.VpnSslWebHostCheckSoftware{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("check_item_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("check_item_list", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslWebHostCheckSoftwareCheckItemList(d, v, "check_item_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.CheckItemList = t
		}
	} else if d.HasChange("check_item_list") {
		old, new := d.GetChange("check_item_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.CheckItemList = &[]models.VpnSslWebHostCheckSoftwareCheckItemList{}
		}
	}
	if v1, ok := d.GetOk("guid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("guid", sv)
				diags = append(diags, e)
			}
			obj.Guid = &v2
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
	if v1, ok := d.GetOk("os_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("os_type", sv)
				diags = append(diags, e)
			}
			obj.OsType = &v2
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("version", sv)
				diags = append(diags, e)
			}
			obj.Version = &v2
		}
	}
	return &obj, diags
}
