// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Coordinate federated upgrades within the Security Fabric.

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

func resourceSystemFederatedUpgrade() *schema.Resource {
	return &schema.Resource{
		Description: "Coordinate federated upgrades within the Security Fabric.",

		CreateContext: resourceSystemFederatedUpgradeCreate,
		ReadContext:   resourceSystemFederatedUpgradeRead,
		UpdateContext: resourceSystemFederatedUpgradeUpdate,
		DeleteContext: resourceSystemFederatedUpgradeDelete,

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
			"failure_device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Serial number of the node to include.",
				Optional:    true,
				Computed:    true,
			},
			"failure_reason": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "internal", "timeout", "device-type-unsupported", "download-failed", "device-missing", "version-unavailable", "staging-failed", "reboot-failed", "device-not-reconnected", "node-not-ready", "no-final-confirmation", "no-confirmation-query"}, false),

				Description: "Reason for upgrade failure.",
				Optional:    true,
				Computed:    true,
			},
			"next_path_index": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10),

				Description: "The index of the next image to upgrade to.",
				Optional:    true,
				Computed:    true,
			},
			"node_list": {
				Type:        schema.TypeList,
				Description: "Nodes which will be included in the upgrade.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"coordinating_fortigate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Serial number of the FortiGate unit that controls this device.",
							Optional:    true,
							Computed:    true,
						},
						"device_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"fortigate", "fortiswitch", "fortiap"}, false),

							Description: "What type of device this node represents.",
							Optional:    true,
							Computed:    true,
						},
						"serial": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Serial number of the node to include.",
							Optional:    true,
							Computed:    true,
						},
						"setup_time": {
							Type: schema.TypeString,

							Description: "When the upgrade was configured. Format hh:mm yyyy/mm/dd UTC.",
							Optional:    true,
							Computed:    true,
						},
						"time": {
							Type: schema.TypeString,

							Description: "Scheduled time for the upgrade. Format hh:mm yyyy/mm/dd UTC.",
							Optional:    true,
							Computed:    true,
						},
						"timing": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"immediate", "scheduled"}, false),

							Description: "Whether the upgrade should be run immediately, or at a scheduled time.",
							Optional:    true,
							Computed:    true,
						},
						"upgrade_path": {
							Type: schema.TypeString,

							Description: "Image IDs to upgrade through.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disabled", "initialized", "downloading", "device-disconnected", "ready", "staging", "final-check", "upgrade-devices", "cancelled", "confirmed", "done", "failed"}, false),

				Description: "Current status of the upgrade.",
				Optional:    true,
				Computed:    true,
			},
			"upgrade_id": {
				Type: schema.TypeInt,

				Description: "Unique identifier for this upgrade.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemFederatedUpgradeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemFederatedUpgrade(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemFederatedUpgrade(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemFederatedUpgrade")
	}

	return resourceSystemFederatedUpgradeRead(ctx, d, meta)
}

func resourceSystemFederatedUpgradeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemFederatedUpgrade(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemFederatedUpgrade(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemFederatedUpgrade resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemFederatedUpgrade")
	}

	return resourceSystemFederatedUpgradeRead(ctx, d, meta)
}

func resourceSystemFederatedUpgradeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemFederatedUpgrade(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemFederatedUpgrade(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemFederatedUpgrade resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFederatedUpgradeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemFederatedUpgrade(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemFederatedUpgrade resource: %v", err)
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

	diags := refreshObjectSystemFederatedUpgrade(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemFederatedUpgradeNodeList(v *[]models.SystemFederatedUpgradeNodeList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CoordinatingFortigate; tmp != nil {
				v["coordinating_fortigate"] = *tmp
			}

			if tmp := cfg.DeviceType; tmp != nil {
				v["device_type"] = *tmp
			}

			if tmp := cfg.Serial; tmp != nil {
				v["serial"] = *tmp
			}

			if tmp := cfg.SetupTime; tmp != nil {
				v["setup_time"] = *tmp
			}

			if tmp := cfg.Time; tmp != nil {
				v["time"] = *tmp
			}

			if tmp := cfg.Timing; tmp != nil {
				v["timing"] = *tmp
			}

			if tmp := cfg.UpgradePath; tmp != nil {
				v["upgrade_path"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "serial")
	}

	return flat
}

func refreshObjectSystemFederatedUpgrade(d *schema.ResourceData, o *models.SystemFederatedUpgrade, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.FailureDevice != nil {
		v := *o.FailureDevice

		if err = d.Set("failure_device", v); err != nil {
			return diag.Errorf("error reading failure_device: %v", err)
		}
	}

	if o.FailureReason != nil {
		v := *o.FailureReason

		if err = d.Set("failure_reason", v); err != nil {
			return diag.Errorf("error reading failure_reason: %v", err)
		}
	}

	if o.NextPathIndex != nil {
		v := *o.NextPathIndex

		if err = d.Set("next_path_index", v); err != nil {
			return diag.Errorf("error reading next_path_index: %v", err)
		}
	}

	if o.NodeList != nil {
		if err = d.Set("node_list", flattenSystemFederatedUpgradeNodeList(o.NodeList, sort)); err != nil {
			return diag.Errorf("error reading node_list: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.UpgradeId != nil {
		v := *o.UpgradeId

		if err = d.Set("upgrade_id", v); err != nil {
			return diag.Errorf("error reading upgrade_id: %v", err)
		}
	}

	return nil
}

func expandSystemFederatedUpgradeNodeList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemFederatedUpgradeNodeList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemFederatedUpgradeNodeList

	for i := range l {
		tmp := models.SystemFederatedUpgradeNodeList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.coordinating_fortigate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CoordinatingFortigate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.device_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DeviceType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.serial", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Serial = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.setup_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SetupTime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Time = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.timing", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Timing = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.upgrade_path", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UpgradePath = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemFederatedUpgrade(d *schema.ResourceData, sv string) (*models.SystemFederatedUpgrade, diag.Diagnostics) {
	obj := models.SystemFederatedUpgrade{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("failure_device"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("failure_device", sv)
				diags = append(diags, e)
			}
			obj.FailureDevice = &v2
		}
	}
	if v1, ok := d.GetOk("failure_reason"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("failure_reason", sv)
				diags = append(diags, e)
			}
			obj.FailureReason = &v2
		}
	}
	if v1, ok := d.GetOk("next_path_index"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("next_path_index", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NextPathIndex = &tmp
		}
	}
	if v, ok := d.GetOk("node_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("node_list", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemFederatedUpgradeNodeList(d, v, "node_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NodeList = t
		}
	} else if d.HasChange("node_list") {
		old, new := d.GetChange("node_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NodeList = &[]models.SystemFederatedUpgradeNodeList{}
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("upgrade_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("upgrade_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UpgradeId = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemFederatedUpgrade(d *schema.ResourceData, sv string) (*models.SystemFederatedUpgrade, diag.Diagnostics) {
	obj := models.SystemFederatedUpgrade{}
	diags := diag.Diagnostics{}

	obj.NodeList = &[]models.SystemFederatedUpgradeNodeList{}

	return &obj, diags
}
