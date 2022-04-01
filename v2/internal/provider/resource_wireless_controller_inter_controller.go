// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure inter wireless controller operation.

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

func resourceWirelessControllerInterController() *schema.Resource {
	return &schema.Resource{
		Description: "Configure inter wireless controller operation.",

		CreateContext: resourceWirelessControllerInterControllerCreate,
		ReadContext:   resourceWirelessControllerInterControllerRead,
		UpdateContext: resourceWirelessControllerInterControllerUpdate,
		DeleteContext: resourceWirelessControllerInterControllerDelete,

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
			"fast_failover_max": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 64),

				Description: "Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"fast_failover_wait": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 86400),

				Description: "Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"inter_controller_key": {
				Type: schema.TypeString,

				Description: "Secret key for inter-controller communications.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"inter_controller_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "l2-roaming", "1+1"}, false),

				Description: "Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"inter_controller_peer": {
				Type:        schema.TypeList,
				Description: "Fast failover peer wireless controller list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"peer_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Peer wireless controller's IP address.",
							Optional:    true,
							Computed:    true,
						},
						"peer_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1024, 49150),

							Description: "Port used by the wireless controller's for inter-controller communications (1024 - 49150, default = 5246).",
							Optional:    true,
							Computed:    true,
						},
						"peer_priority": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"primary", "secondary"}, false),

							Description: "Peer wireless controller's priority (primary or secondary, default = primary).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"inter_controller_pri": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"primary", "secondary"}, false),

				Description: "Configure inter-controller's priority (primary or secondary, default = primary).",
				Optional:    true,
				Computed:    true,
			},
			"l3_roaming": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable layer 3 roaming (default = disable).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerInterControllerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerInterController(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerInterController(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerInterController")
	}

	return resourceWirelessControllerInterControllerRead(ctx, d, meta)
}

func resourceWirelessControllerInterControllerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerInterController(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerInterController(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerInterController resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerInterController")
	}

	return resourceWirelessControllerInterControllerRead(ctx, d, meta)
}

func resourceWirelessControllerInterControllerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectWirelessControllerInterController(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateWirelessControllerInterController(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerInterController resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerInterControllerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerInterController(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerInterController resource: %v", err)
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

	diags := refreshObjectWirelessControllerInterController(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerInterControllerInterControllerPeer(d *schema.ResourceData, v *[]models.WirelessControllerInterControllerInterControllerPeer, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.PeerIp; tmp != nil {
				v["peer_ip"] = *tmp
			}

			if tmp := cfg.PeerPort; tmp != nil {
				v["peer_port"] = *tmp
			}

			if tmp := cfg.PeerPriority; tmp != nil {
				v["peer_priority"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectWirelessControllerInterController(d *schema.ResourceData, o *models.WirelessControllerInterController, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.FastFailoverMax != nil {
		v := *o.FastFailoverMax

		if err = d.Set("fast_failover_max", v); err != nil {
			return diag.Errorf("error reading fast_failover_max: %v", err)
		}
	}

	if o.FastFailoverWait != nil {
		v := *o.FastFailoverWait

		if err = d.Set("fast_failover_wait", v); err != nil {
			return diag.Errorf("error reading fast_failover_wait: %v", err)
		}
	}

	if o.InterControllerKey != nil {
		v := *o.InterControllerKey

		if v == "ENC XXXX" {
		} else if err = d.Set("inter_controller_key", v); err != nil {
			return diag.Errorf("error reading inter_controller_key: %v", err)
		}
	}

	if o.InterControllerMode != nil {
		v := *o.InterControllerMode

		if err = d.Set("inter_controller_mode", v); err != nil {
			return diag.Errorf("error reading inter_controller_mode: %v", err)
		}
	}

	if o.InterControllerPeer != nil {
		if err = d.Set("inter_controller_peer", flattenWirelessControllerInterControllerInterControllerPeer(d, o.InterControllerPeer, "inter_controller_peer", sort)); err != nil {
			return diag.Errorf("error reading inter_controller_peer: %v", err)
		}
	}

	if o.InterControllerPri != nil {
		v := *o.InterControllerPri

		if err = d.Set("inter_controller_pri", v); err != nil {
			return diag.Errorf("error reading inter_controller_pri: %v", err)
		}
	}

	if o.L3Roaming != nil {
		v := *o.L3Roaming

		if err = d.Set("l3_roaming", v); err != nil {
			return diag.Errorf("error reading l3_roaming: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerInterControllerInterControllerPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerInterControllerInterControllerPeer, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerInterControllerInterControllerPeer

	for i := range l {
		tmp := models.WirelessControllerInterControllerInterControllerPeer{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.peer_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PeerIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.peer_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PeerPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.peer_priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PeerPriority = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerInterController(d *schema.ResourceData, sv string) (*models.WirelessControllerInterController, diag.Diagnostics) {
	obj := models.WirelessControllerInterController{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("fast_failover_max"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fast_failover_max", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FastFailoverMax = &tmp
		}
	}
	if v1, ok := d.GetOk("fast_failover_wait"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fast_failover_wait", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FastFailoverWait = &tmp
		}
	}
	if v1, ok := d.GetOk("inter_controller_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("inter_controller_key", sv)
				diags = append(diags, e)
			}
			obj.InterControllerKey = &v2
		}
	}
	if v1, ok := d.GetOk("inter_controller_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("inter_controller_mode", sv)
				diags = append(diags, e)
			}
			obj.InterControllerMode = &v2
		}
	}
	if v, ok := d.GetOk("inter_controller_peer"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("inter_controller_peer", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerInterControllerInterControllerPeer(d, v, "inter_controller_peer", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InterControllerPeer = t
		}
	} else if d.HasChange("inter_controller_peer") {
		old, new := d.GetChange("inter_controller_peer")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InterControllerPeer = &[]models.WirelessControllerInterControllerInterControllerPeer{}
		}
	}
	if v1, ok := d.GetOk("inter_controller_pri"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("inter_controller_pri", sv)
				diags = append(diags, e)
			}
			obj.InterControllerPri = &v2
		}
	}
	if v1, ok := d.GetOk("l3_roaming"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("l3_roaming", sv)
				diags = append(diags, e)
			}
			obj.L3Roaming = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectWirelessControllerInterController(d *schema.ResourceData, sv string) (*models.WirelessControllerInterController, diag.Diagnostics) {
	obj := models.WirelessControllerInterController{}
	diags := diag.Diagnostics{}

	obj.InterControllerPeer = &[]models.WirelessControllerInterControllerInterControllerPeer{}

	return &obj, diags
}
