// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure BFD.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceRouterBfd() *schema.Resource {
	return &schema.Resource{
		Description: "Configure BFD.",

		CreateContext: resourceRouterBfdCreate,
		ReadContext:   resourceRouterBfdRead,
		UpdateContext: resourceRouterBfdUpdate,
		DeleteContext: resourceRouterBfdDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"multihop_template": {
				Type:        schema.TypeList,
				Description: "BFD multi-hop template table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "md5"}, false),

							Description: "Authentication mode.",
							Optional:    true,
							Computed:    true,
						},
						"bfd_desired_min_tx": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(100, 30000),

							Description: "BFD desired minimal transmit interval (milliseconds).",
							Optional:    true,
							Computed:    true,
						},
						"bfd_detect_mult": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 50),

							Description: "BFD detection multiplier.",
							Optional:    true,
							Computed:    true,
						},
						"bfd_required_min_rx": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(100, 30000),

							Description: "BFD required minimal receive interval (milliseconds).",
							Optional:    true,
							Computed:    true,
						},
						"dst": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Classnet,

							Description: "Destination prefix.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"md5_key": {
							Type: schema.TypeString,

							Description: "MD5 key of key ID 1.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"src": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Classnet,

							Description: "Source prefix.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "Neighbor.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IPv4 address of the BFD neighbor.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceRouterBfdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterBfd(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterBfd(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBfd")
	}

	return resourceRouterBfdRead(ctx, d, meta)
}

func resourceRouterBfdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterBfd(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterBfd(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterBfd resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBfd")
	}

	return resourceRouterBfdRead(ctx, d, meta)
}

func resourceRouterBfdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectRouterBfd(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateRouterBfd(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterBfd resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBfdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterBfd(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBfd resource: %v", err)
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

	diags := refreshObjectRouterBfd(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterBfdMultihopTemplate(d *schema.ResourceData, v *[]models.RouterBfdMultihopTemplate, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AuthMode; tmp != nil {
				v["auth_mode"] = *tmp
			}

			if tmp := cfg.BfdDesiredMinTx; tmp != nil {
				v["bfd_desired_min_tx"] = *tmp
			}

			if tmp := cfg.BfdDetectMult; tmp != nil {
				v["bfd_detect_mult"] = *tmp
			}

			if tmp := cfg.BfdRequiredMinRx; tmp != nil {
				v["bfd_required_min_rx"] = *tmp
			}

			if tmp := cfg.Dst; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.dst", prefix, i), *tmp); tmp != nil {
					v["dst"] = *tmp
				}
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Md5Key; tmp != nil {
				v["md5_key"] = *tmp
			}

			if tmp := cfg.Src; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.src", prefix, i), *tmp); tmp != nil {
					v["src"] = *tmp
				}
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterBfdNeighbor(d *schema.ResourceData, v *[]models.RouterBfdNeighbor, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "ip")
	}

	return flat
}

func refreshObjectRouterBfd(d *schema.ResourceData, o *models.RouterBfd, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.MultihopTemplate != nil {
		if err = d.Set("multihop_template", flattenRouterBfdMultihopTemplate(d, o.MultihopTemplate, "multihop_template", sort)); err != nil {
			return diag.Errorf("error reading multihop_template: %v", err)
		}
	}

	if o.Neighbor != nil {
		if err = d.Set("neighbor", flattenRouterBfdNeighbor(d, o.Neighbor, "neighbor", sort)); err != nil {
			return diag.Errorf("error reading neighbor: %v", err)
		}
	}

	return nil
}

func expandRouterBfdMultihopTemplate(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBfdMultihopTemplate, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBfdMultihopTemplate

	for i := range l {
		tmp := models.RouterBfdMultihopTemplate{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bfd_desired_min_tx", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BfdDesiredMinTx = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bfd_detect_mult", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BfdDetectMult = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bfd_required_min_rx", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BfdRequiredMinRx = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dst", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dst = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.md5_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Md5Key = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.src", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Src = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBfdNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBfdNeighbor, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBfdNeighbor

	for i := range l {
		tmp := models.RouterBfdNeighbor{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectRouterBfd(d *schema.ResourceData, sv string) (*models.RouterBfd, diag.Diagnostics) {
	obj := models.RouterBfd{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("multihop_template"); ok {
		if !utils.CheckVer(sv, "v7.2.0", "") {
			e := utils.AttributeVersionWarning("multihop_template", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBfdMultihopTemplate(d, v, "multihop_template", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.MultihopTemplate = t
		}
	} else if d.HasChange("multihop_template") {
		old, new := d.GetChange("multihop_template")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.MultihopTemplate = &[]models.RouterBfdMultihopTemplate{}
		}
	}
	if v, ok := d.GetOk("neighbor"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("neighbor", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBfdNeighbor(d, v, "neighbor", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Neighbor = t
		}
	} else if d.HasChange("neighbor") {
		old, new := d.GetChange("neighbor")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Neighbor = &[]models.RouterBfdNeighbor{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectRouterBfd(d *schema.ResourceData, sv string) (*models.RouterBfd, diag.Diagnostics) {
	obj := models.RouterBfd{}
	diags := diag.Diagnostics{}

	obj.MultihopTemplate = &[]models.RouterBfdMultihopTemplate{}
	obj.Neighbor = &[]models.RouterBfdNeighbor{}

	return &obj, diags
}
