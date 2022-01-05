// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Designate cache-service for wan-optimization and webcache.

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

func resourceWanoptCacheService() *schema.Resource {
	return &schema.Resource{
		Description: "Designate cache-service for wan-optimization and webcache.",

		CreateContext: resourceWanoptCacheServiceCreate,
		ReadContext:   resourceWanoptCacheServiceRead,
		UpdateContext: resourceWanoptCacheServiceUpdate,
		DeleteContext: resourceWanoptCacheServiceDelete,

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
			"acceptable_connections": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"any", "peers"}, false),

				Description: "Set strategy when accepting cache collaboration connection.",
				Optional:    true,
				Computed:    true,
			},
			"collaboration": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable cache-collaboration between cache-service clusters.",
				Optional:    true,
				Computed:    true,
			},
			"device_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Set identifier for this cache device.",
				Optional:    true,
				Computed:    true,
			},
			"dst_peer": {
				Type:        schema.TypeList,
				Description: "Modify cache-service destination peer list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Set authentication type for this peer.",
							Optional:    true,
							Computed:    true,
						},
						"device_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Device ID of this peer.",
							Optional:    true,
							Computed:    true,
						},
						"encode_type": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Set encode type for this peer.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Set cluster IP address of this peer.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Set priority for this peer.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"prefer_scenario": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"balance", "prefer-speed", "prefer-cache"}, false),

				Description: "Set the preferred cache behavior towards the balance between latency and hit-ratio.",
				Optional:    true,
				Computed:    true,
			},
			"src_peer": {
				Type:        schema.TypeList,
				Description: "Modify cache-service source peer list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Set authentication type for this peer.",
							Optional:    true,
							Computed:    true,
						},
						"device_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Device ID of this peer.",
							Optional:    true,
							Computed:    true,
						},
						"encode_type": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Set encode type for this peer.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Set cluster IP address of this peer.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Set priority for this peer.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceWanoptCacheServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWanoptCacheService(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWanoptCacheService(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptCacheService")
	}

	return resourceWanoptCacheServiceRead(ctx, d, meta)
}

func resourceWanoptCacheServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWanoptCacheService(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWanoptCacheService(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WanoptCacheService resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptCacheService")
	}

	return resourceWanoptCacheServiceRead(ctx, d, meta)
}

func resourceWanoptCacheServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectWanoptCacheService(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateWanoptCacheService(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WanoptCacheService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptCacheServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWanoptCacheService(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WanoptCacheService resource: %v", err)
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

	diags := refreshObjectWanoptCacheService(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWanoptCacheServiceDstPeer(v *[]models.WanoptCacheServiceDstPeer, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AuthType; tmp != nil {
				v["auth_type"] = *tmp
			}

			if tmp := cfg.DeviceId; tmp != nil {
				v["device_id"] = *tmp
			}

			if tmp := cfg.EncodeType; tmp != nil {
				v["encode_type"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "device_id")
	}

	return flat
}

func flattenWanoptCacheServiceSrcPeer(v *[]models.WanoptCacheServiceSrcPeer, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AuthType; tmp != nil {
				v["auth_type"] = *tmp
			}

			if tmp := cfg.DeviceId; tmp != nil {
				v["device_id"] = *tmp
			}

			if tmp := cfg.EncodeType; tmp != nil {
				v["encode_type"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "device_id")
	}

	return flat
}

func refreshObjectWanoptCacheService(d *schema.ResourceData, o *models.WanoptCacheService, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AcceptableConnections != nil {
		v := *o.AcceptableConnections

		if err = d.Set("acceptable_connections", v); err != nil {
			return diag.Errorf("error reading acceptable_connections: %v", err)
		}
	}

	if o.Collaboration != nil {
		v := *o.Collaboration

		if err = d.Set("collaboration", v); err != nil {
			return diag.Errorf("error reading collaboration: %v", err)
		}
	}

	if o.DeviceId != nil {
		v := *o.DeviceId

		if err = d.Set("device_id", v); err != nil {
			return diag.Errorf("error reading device_id: %v", err)
		}
	}

	if o.DstPeer != nil {
		if err = d.Set("dst_peer", flattenWanoptCacheServiceDstPeer(o.DstPeer, sort)); err != nil {
			return diag.Errorf("error reading dst_peer: %v", err)
		}
	}

	if o.PreferScenario != nil {
		v := *o.PreferScenario

		if err = d.Set("prefer_scenario", v); err != nil {
			return diag.Errorf("error reading prefer_scenario: %v", err)
		}
	}

	if o.SrcPeer != nil {
		if err = d.Set("src_peer", flattenWanoptCacheServiceSrcPeer(o.SrcPeer, sort)); err != nil {
			return diag.Errorf("error reading src_peer: %v", err)
		}
	}

	return nil
}

func expandWanoptCacheServiceDstPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WanoptCacheServiceDstPeer, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptCacheServiceDstPeer

	for i := range l {
		tmp := models.WanoptCacheServiceDstPeer{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.AuthType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.device_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DeviceId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.encode_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.EncodeType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Priority = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWanoptCacheServiceSrcPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WanoptCacheServiceSrcPeer, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptCacheServiceSrcPeer

	for i := range l {
		tmp := models.WanoptCacheServiceSrcPeer{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.AuthType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.device_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DeviceId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.encode_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.EncodeType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Priority = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWanoptCacheService(d *schema.ResourceData, sv string) (*models.WanoptCacheService, diag.Diagnostics) {
	obj := models.WanoptCacheService{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("acceptable_connections"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("acceptable_connections", sv)
				diags = append(diags, e)
			}
			obj.AcceptableConnections = &v2
		}
	}
	if v1, ok := d.GetOk("collaboration"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("collaboration", sv)
				diags = append(diags, e)
			}
			obj.Collaboration = &v2
		}
	}
	if v1, ok := d.GetOk("device_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("device_id", sv)
				diags = append(diags, e)
			}
			obj.DeviceId = &v2
		}
	}
	if v, ok := d.GetOk("dst_peer"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dst_peer", sv)
			diags = append(diags, e)
		}
		t, err := expandWanoptCacheServiceDstPeer(d, v, "dst_peer", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DstPeer = t
		}
	} else if d.HasChange("dst_peer") {
		old, new := d.GetChange("dst_peer")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DstPeer = &[]models.WanoptCacheServiceDstPeer{}
		}
	}
	if v1, ok := d.GetOk("prefer_scenario"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("prefer_scenario", sv)
				diags = append(diags, e)
			}
			obj.PreferScenario = &v2
		}
	}
	if v, ok := d.GetOk("src_peer"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("src_peer", sv)
			diags = append(diags, e)
		}
		t, err := expandWanoptCacheServiceSrcPeer(d, v, "src_peer", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SrcPeer = t
		}
	} else if d.HasChange("src_peer") {
		old, new := d.GetChange("src_peer")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SrcPeer = &[]models.WanoptCacheServiceSrcPeer{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectWanoptCacheService(d *schema.ResourceData, sv string) (*models.WanoptCacheService, diag.Diagnostics) {
	obj := models.WanoptCacheService{}
	diags := diag.Diagnostics{}

	obj.DstPeer = &[]models.WanoptCacheServiceDstPeer{}
	obj.SrcPeer = &[]models.WanoptCacheServiceSrcPeer{}

	return &obj, diags
}
