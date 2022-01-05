// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Forward Error Correction (FEC) mapping profiles.

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

func resourceVpnIpsecFec() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Forward Error Correction (FEC) mapping profiles.",

		CreateContext: resourceVpnIpsecFecCreate,
		ReadContext:   resourceVpnIpsecFecRead,
		UpdateContext: resourceVpnIpsecFecUpdate,
		DeleteContext: resourceVpnIpsecFecDelete,

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
			"mappings": {
				Type:        schema.TypeList,
				Description: "FEC redundancy mapping table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bandwidth_bi_threshold": {
							Type: schema.TypeInt,

							Description: "Apply FEC parameters when available bi-bandwidth is >= threshold (kbps, 0 means no threshold).",
							Optional:    true,
							Computed:    true,
						},
						"bandwidth_down_threshold": {
							Type: schema.TypeInt,

							Description: "Apply FEC parameters when available down bandwidth is >= threshold (kbps, 0 means no threshold).",
							Optional:    true,
							Computed:    true,
						},
						"bandwidth_up_threshold": {
							Type: schema.TypeInt,

							Description: "Apply FEC parameters when available up bandwidth is >= threshold (kbps, 0 means no threshold).",
							Optional:    true,
							Computed:    true,
						},
						"base": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 20),

							Description: "Number of base FEC packets (1 - 20).",
							Optional:    true,
							Computed:    true,
						},
						"latency_threshold": {
							Type: schema.TypeInt,

							Description: "Apply FEC parameters when latency is <= threshold (0 means no threshold).",
							Optional:    true,
							Computed:    true,
						},
						"packet_loss_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),

							Description: "Apply FEC parameters when packet loss is >= threshold (0 - 100, 0 means no threshold).",
							Optional:    true,
							Computed:    true,
						},
						"redundant": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 5),

							Description: "Number of redundant FEC packets (1 - 5).",
							Optional:    true,
							Computed:    true,
						},
						"seqno": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 64),

							Description: "Sequence number (1 - 64).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Profile name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceVpnIpsecFecCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating VpnIpsecFec resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnIpsecFec(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnIpsecFec(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnIpsecFec")
	}

	return resourceVpnIpsecFecRead(ctx, d, meta)
}

func resourceVpnIpsecFecUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnIpsecFec(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnIpsecFec(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnIpsecFec resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnIpsecFec")
	}

	return resourceVpnIpsecFecRead(ctx, d, meta)
}

func resourceVpnIpsecFecDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVpnIpsecFec(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnIpsecFec resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecFecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnIpsecFec(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnIpsecFec resource: %v", err)
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

	diags := refreshObjectVpnIpsecFec(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenVpnIpsecFecMappings(v *[]models.VpnIpsecFecMappings, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.BandwidthBiThreshold; tmp != nil {
				v["bandwidth_bi_threshold"] = *tmp
			}

			if tmp := cfg.BandwidthDownThreshold; tmp != nil {
				v["bandwidth_down_threshold"] = *tmp
			}

			if tmp := cfg.BandwidthUpThreshold; tmp != nil {
				v["bandwidth_up_threshold"] = *tmp
			}

			if tmp := cfg.Base; tmp != nil {
				v["base"] = *tmp
			}

			if tmp := cfg.LatencyThreshold; tmp != nil {
				v["latency_threshold"] = *tmp
			}

			if tmp := cfg.PacketLossThreshold; tmp != nil {
				v["packet_loss_threshold"] = *tmp
			}

			if tmp := cfg.Redundant; tmp != nil {
				v["redundant"] = *tmp
			}

			if tmp := cfg.Seqno; tmp != nil {
				v["seqno"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "seqno")
	}

	return flat
}

func refreshObjectVpnIpsecFec(d *schema.ResourceData, o *models.VpnIpsecFec, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Mappings != nil {
		if err = d.Set("mappings", flattenVpnIpsecFecMappings(o.Mappings, sort)); err != nil {
			return diag.Errorf("error reading mappings: %v", err)
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

func expandVpnIpsecFecMappings(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnIpsecFecMappings, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnIpsecFecMappings

	for i := range l {
		tmp := models.VpnIpsecFecMappings{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.bandwidth_bi_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.BandwidthBiThreshold = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bandwidth_down_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.BandwidthDownThreshold = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bandwidth_up_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.BandwidthUpThreshold = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.base", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Base = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.latency_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.LatencyThreshold = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.packet_loss_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PacketLossThreshold = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.redundant", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Redundant = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.seqno", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Seqno = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectVpnIpsecFec(d *schema.ResourceData, sv string) (*models.VpnIpsecFec, diag.Diagnostics) {
	obj := models.VpnIpsecFec{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("mappings"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("mappings", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnIpsecFecMappings(d, v, "mappings", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Mappings = t
		}
	} else if d.HasChange("mappings") {
		old, new := d.GetChange("mappings")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Mappings = &[]models.VpnIpsecFecMappings{}
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
