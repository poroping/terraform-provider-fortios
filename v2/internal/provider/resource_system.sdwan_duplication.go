// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Create SD-WAN duplication rule.

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

func resourceSystemsdwanDuplication() *schema.Resource {
	return &schema.Resource{
		Description: "Create SD-WAN duplication rule.",

		CreateContext: resourceSystemsdwanDuplicationCreate,
		ReadContext:   resourceSystemsdwanDuplicationRead,
		UpdateContext: resourceSystemsdwanDuplicationUpdate,
		DeleteContext: resourceSystemsdwanDuplicationDelete,

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
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination address or address group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address or address group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dstaddr6": {
				Type:        schema.TypeList,
				Description: "Destination address6 or address6 group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address6 or address6 group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dstintf": {
				Type:        schema.TypeList,
				Description: "Outgoing (egress) interfaces or zones.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface, zone or SDWAN zone name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Duplication rule ID (1 - 255).",
				Optional:    true,
				Computed:    true,
			},
			"packet_de_duplication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable discarding of packets that have been duplicated.",
				Optional:    true,
				Computed:    true,
			},
			"packet_duplication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "force", "on-demand"}, false),

				Description: "Configure packet duplication method.",
				Optional:    true,
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Service and service group name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Service and service group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"service_id": {
				Type:        schema.TypeList,
				Description: "SD-WAN service rule ID list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "SD-WAN service rule ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address or address group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address or address group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"srcaddr6": {
				Type:        schema.TypeList,
				Description: "Source address6 or address6 group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address6 or address6 group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"srcintf": {
				Type:        schema.TypeList,
				Description: "Incoming (ingress) interfaces or zones.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface, zone or SDWAN zone name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemsdwanDuplicationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemsdwanDuplication resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemsdwanDuplication(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemsdwanDuplication(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemsdwanDuplication")
	}

	return resourceSystemsdwanDuplicationRead(ctx, d, meta)
}

func resourceSystemsdwanDuplicationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemsdwanDuplication(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemsdwanDuplication(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemsdwanDuplication resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemsdwanDuplication")
	}

	return resourceSystemsdwanDuplicationRead(ctx, d, meta)
}

func resourceSystemsdwanDuplicationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemsdwanDuplication(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemsdwanDuplication resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemsdwanDuplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemsdwanDuplication(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemsdwanDuplication resource: %v", err)
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

	diags := refreshObjectSystemsdwanDuplication(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemsdwanDuplicationDstaddr(v *[]models.SystemsdwanDuplicationDstaddr, sort bool) interface{} {
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

func flattenSystemsdwanDuplicationDstaddr6(v *[]models.SystemsdwanDuplicationDstaddr6, sort bool) interface{} {
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

func flattenSystemsdwanDuplicationDstintf(v *[]models.SystemsdwanDuplicationDstintf, sort bool) interface{} {
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

func flattenSystemsdwanDuplicationService(v *[]models.SystemsdwanDuplicationService, sort bool) interface{} {
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

func flattenSystemsdwanDuplicationServiceId(v *[]models.SystemsdwanDuplicationServiceId, sort bool) interface{} {
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

func flattenSystemsdwanDuplicationSrcaddr(v *[]models.SystemsdwanDuplicationSrcaddr, sort bool) interface{} {
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

func flattenSystemsdwanDuplicationSrcaddr6(v *[]models.SystemsdwanDuplicationSrcaddr6, sort bool) interface{} {
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

func flattenSystemsdwanDuplicationSrcintf(v *[]models.SystemsdwanDuplicationSrcintf, sort bool) interface{} {
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

func refreshObjectSystemsdwanDuplication(d *schema.ResourceData, o *models.SystemsdwanDuplication, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Dstaddr != nil {
		if err = d.Set("dstaddr", flattenSystemsdwanDuplicationDstaddr(o.Dstaddr, sort)); err != nil {
			return diag.Errorf("error reading dstaddr: %v", err)
		}
	}

	if o.Dstaddr6 != nil {
		if err = d.Set("dstaddr6", flattenSystemsdwanDuplicationDstaddr6(o.Dstaddr6, sort)); err != nil {
			return diag.Errorf("error reading dstaddr6: %v", err)
		}
	}

	if o.Dstintf != nil {
		if err = d.Set("dstintf", flattenSystemsdwanDuplicationDstintf(o.Dstintf, sort)); err != nil {
			return diag.Errorf("error reading dstintf: %v", err)
		}
	}

	if o.Fosid != nil {
		v := *o.Fosid

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.PacketDeDuplication != nil {
		v := *o.PacketDeDuplication

		if err = d.Set("packet_de_duplication", v); err != nil {
			return diag.Errorf("error reading packet_de_duplication: %v", err)
		}
	}

	if o.PacketDuplication != nil {
		v := *o.PacketDuplication

		if err = d.Set("packet_duplication", v); err != nil {
			return diag.Errorf("error reading packet_duplication: %v", err)
		}
	}

	if o.Service != nil {
		if err = d.Set("service", flattenSystemsdwanDuplicationService(o.Service, sort)); err != nil {
			return diag.Errorf("error reading service: %v", err)
		}
	}

	if o.ServiceId != nil {
		if err = d.Set("service_id", flattenSystemsdwanDuplicationServiceId(o.ServiceId, sort)); err != nil {
			return diag.Errorf("error reading service_id: %v", err)
		}
	}

	if o.Srcaddr != nil {
		if err = d.Set("srcaddr", flattenSystemsdwanDuplicationSrcaddr(o.Srcaddr, sort)); err != nil {
			return diag.Errorf("error reading srcaddr: %v", err)
		}
	}

	if o.Srcaddr6 != nil {
		if err = d.Set("srcaddr6", flattenSystemsdwanDuplicationSrcaddr6(o.Srcaddr6, sort)); err != nil {
			return diag.Errorf("error reading srcaddr6: %v", err)
		}
	}

	if o.Srcintf != nil {
		if err = d.Set("srcintf", flattenSystemsdwanDuplicationSrcintf(o.Srcintf, sort)); err != nil {
			return diag.Errorf("error reading srcintf: %v", err)
		}
	}

	return nil
}

func expandSystemsdwanDuplicationDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanDuplicationDstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanDuplicationDstaddr

	for i := range l {
		tmp := models.SystemsdwanDuplicationDstaddr{}
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

func expandSystemsdwanDuplicationDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanDuplicationDstaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanDuplicationDstaddr6

	for i := range l {
		tmp := models.SystemsdwanDuplicationDstaddr6{}
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

func expandSystemsdwanDuplicationDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanDuplicationDstintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanDuplicationDstintf

	for i := range l {
		tmp := models.SystemsdwanDuplicationDstintf{}
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

func expandSystemsdwanDuplicationService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanDuplicationService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanDuplicationService

	for i := range l {
		tmp := models.SystemsdwanDuplicationService{}
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

func expandSystemsdwanDuplicationServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanDuplicationServiceId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanDuplicationServiceId

	for i := range l {
		tmp := models.SystemsdwanDuplicationServiceId{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemsdwanDuplicationSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanDuplicationSrcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanDuplicationSrcaddr

	for i := range l {
		tmp := models.SystemsdwanDuplicationSrcaddr{}
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

func expandSystemsdwanDuplicationSrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanDuplicationSrcaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanDuplicationSrcaddr6

	for i := range l {
		tmp := models.SystemsdwanDuplicationSrcaddr6{}
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

func expandSystemsdwanDuplicationSrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanDuplicationSrcintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanDuplicationSrcintf

	for i := range l {
		tmp := models.SystemsdwanDuplicationSrcintf{}
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

func getObjectSystemsdwanDuplication(d *schema.ResourceData, sv string) (*models.SystemsdwanDuplication, diag.Diagnostics) {
	obj := models.SystemsdwanDuplication{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("dstaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanDuplicationDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr = t
		}
	} else if d.HasChange("dstaddr") {
		old, new := d.GetChange("dstaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr = &[]models.SystemsdwanDuplicationDstaddr{}
		}
	}
	if v, ok := d.GetOk("dstaddr6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanDuplicationDstaddr6(d, v, "dstaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr6 = t
		}
	} else if d.HasChange("dstaddr6") {
		old, new := d.GetChange("dstaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr6 = &[]models.SystemsdwanDuplicationDstaddr6{}
		}
	}
	if v, ok := d.GetOk("dstintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstintf", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanDuplicationDstintf(d, v, "dstintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstintf = t
		}
	} else if d.HasChange("dstintf") {
		old, new := d.GetChange("dstintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstintf = &[]models.SystemsdwanDuplicationDstintf{}
		}
	}
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
	if v1, ok := d.GetOk("packet_de_duplication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("packet_de_duplication", sv)
				diags = append(diags, e)
			}
			obj.PacketDeDuplication = &v2
		}
	}
	if v1, ok := d.GetOk("packet_duplication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("packet_duplication", sv)
				diags = append(diags, e)
			}
			obj.PacketDuplication = &v2
		}
	}
	if v, ok := d.GetOk("service"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("service", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanDuplicationService(d, v, "service", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Service = t
		}
	} else if d.HasChange("service") {
		old, new := d.GetChange("service")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Service = &[]models.SystemsdwanDuplicationService{}
		}
	}
	if v, ok := d.GetOk("service_id"); ok {
		if !utils.CheckVer(sv, "v6.4.3", "") {
			e := utils.AttributeVersionWarning("service_id", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanDuplicationServiceId(d, v, "service_id", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ServiceId = t
		}
	} else if d.HasChange("service_id") {
		old, new := d.GetChange("service_id")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ServiceId = &[]models.SystemsdwanDuplicationServiceId{}
		}
	}
	if v, ok := d.GetOk("srcaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanDuplicationSrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr = t
		}
	} else if d.HasChange("srcaddr") {
		old, new := d.GetChange("srcaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr = &[]models.SystemsdwanDuplicationSrcaddr{}
		}
	}
	if v, ok := d.GetOk("srcaddr6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanDuplicationSrcaddr6(d, v, "srcaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr6 = t
		}
	} else if d.HasChange("srcaddr6") {
		old, new := d.GetChange("srcaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr6 = &[]models.SystemsdwanDuplicationSrcaddr6{}
		}
	}
	if v, ok := d.GetOk("srcintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcintf", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanDuplicationSrcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcintf = t
		}
	} else if d.HasChange("srcintf") {
		old, new := d.GetChange("srcintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcintf = &[]models.SystemsdwanDuplicationSrcintf{}
		}
	}
	return &obj, diags
}
