// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch LLDP settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSwitchControllerLldpSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch LLDP settings.",

		CreateContext: resourceSwitchControllerLldpSettingsCreate,
		ReadContext:   resourceSwitchControllerLldpSettingsRead,
		UpdateContext: resourceSwitchControllerLldpSettingsUpdate,
		DeleteContext: resourceSwitchControllerLldpSettingsDelete,

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
			"device_detection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable dynamic detection of LLDP neighbor devices for VLAN assignment.",
				Optional:    true,
				Computed:    true,
			},
			"fast_start_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Frequency of LLDP PDU transmission from FortiSwitch for the first 4 packets when the link is up (2 - 5 sec, default = 2, 0 = disable fast start).",
				Optional:    true,
				Computed:    true,
			},
			"management_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"internal", "mgmt"}, false),

				Description: "Primary management interface to be advertised in LLDP and CDP PDUs.",
				Optional:    true,
				Computed:    true,
			},
			"tx_hold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16),

				Description: "Number of tx-intervals before local LLDP data expires (1 - 16, default = 4). Packet TTL is tx-hold * tx-interval.",
				Optional:    true,
				Computed:    true,
			},
			"tx_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 4095),

				Description: "Frequency of LLDP PDU transmission from FortiSwitch (5 - 4095 sec, default = 30). Packet TTL is tx-hold * tx-interval.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSwitchControllerLldpSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerLldpSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerLldpSettings(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerLldpSettings")
	}

	return resourceSwitchControllerLldpSettingsRead(ctx, d, meta)
}

func resourceSwitchControllerLldpSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerLldpSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerLldpSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerLldpSettings resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerLldpSettings")
	}

	return resourceSwitchControllerLldpSettingsRead(ctx, d, meta)
}

func resourceSwitchControllerLldpSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSwitchControllerLldpSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerLldpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLldpSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerLldpSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerLldpSettings resource: %v", err)
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

	diags := refreshObjectSwitchControllerLldpSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSwitchControllerLldpSettings(d *schema.ResourceData, o *models.SwitchControllerLldpSettings, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DeviceDetection != nil {
		v := *o.DeviceDetection

		if err = d.Set("device_detection", v); err != nil {
			return diag.Errorf("error reading device_detection: %v", err)
		}
	}

	if o.FastStartInterval != nil {
		v := *o.FastStartInterval

		if err = d.Set("fast_start_interval", v); err != nil {
			return diag.Errorf("error reading fast_start_interval: %v", err)
		}
	}

	if o.ManagementInterface != nil {
		v := *o.ManagementInterface

		if err = d.Set("management_interface", v); err != nil {
			return diag.Errorf("error reading management_interface: %v", err)
		}
	}

	if o.TxHold != nil {
		v := *o.TxHold

		if err = d.Set("tx_hold", v); err != nil {
			return diag.Errorf("error reading tx_hold: %v", err)
		}
	}

	if o.TxInterval != nil {
		v := *o.TxInterval

		if err = d.Set("tx_interval", v); err != nil {
			return diag.Errorf("error reading tx_interval: %v", err)
		}
	}

	return nil
}

func getObjectSwitchControllerLldpSettings(d *schema.ResourceData, sv string) (*models.SwitchControllerLldpSettings, diag.Diagnostics) {
	obj := models.SwitchControllerLldpSettings{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("device_detection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("device_detection", sv)
				diags = append(diags, e)
			}
			obj.DeviceDetection = &v2
		}
	}
	if v1, ok := d.GetOk("fast_start_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fast_start_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FastStartInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("management_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("management_interface", sv)
				diags = append(diags, e)
			}
			obj.ManagementInterface = &v2
		}
	}
	if v1, ok := d.GetOk("tx_hold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tx_hold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TxHold = &tmp
		}
	}
	if v1, ok := d.GetOk("tx_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tx_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TxInterval = &tmp
		}
	}
	return &obj, diags
}
