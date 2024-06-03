// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS VDOM parameter.

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

func resourceIpsSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS VDOM parameter.",

		CreateContext: resourceIpsSettingsCreate,
		ReadContext:   resourceIpsSettingsRead,
		UpdateContext: resourceIpsSettingsUpdate,
		DeleteContext: resourceIpsSettingsDelete,

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
			"ips_packet_quota": {
				Type: schema.TypeInt,

				Description: "Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.",
				Optional:    true,
				Computed:    true,
			},
			"packet_log_history": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).",
				Optional:    true,
				Computed:    true,
			},
			"packet_log_memory": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(64, 8192),

				Description: "Maximum memory can be used by packet log (64 - 8192 kB).",
				Optional:    true,
				Computed:    true,
			},
			"packet_log_post_attack": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Number of packets to log after the IPS signature is detected (0 - 255).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceIpsSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectIpsSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateIpsSettings(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IpsSettings")
	}

	return resourceIpsSettingsRead(ctx, d, meta)
}

func resourceIpsSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectIpsSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateIpsSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating IpsSettings resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IpsSettings")
	}

	return resourceIpsSettingsRead(ctx, d, meta)
}

func resourceIpsSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectIpsSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateIpsSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting IpsSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadIpsSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IpsSettings resource: %v", err)
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

	diags := refreshObjectIpsSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectIpsSettings(d *schema.ResourceData, o *models.IpsSettings, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.IpsPacketQuota != nil {
		v := *o.IpsPacketQuota

		if err = d.Set("ips_packet_quota", v); err != nil {
			return diag.Errorf("error reading ips_packet_quota: %v", err)
		}
	}

	if o.PacketLogHistory != nil {
		v := *o.PacketLogHistory

		if err = d.Set("packet_log_history", v); err != nil {
			return diag.Errorf("error reading packet_log_history: %v", err)
		}
	}

	if o.PacketLogMemory != nil {
		v := *o.PacketLogMemory

		if err = d.Set("packet_log_memory", v); err != nil {
			return diag.Errorf("error reading packet_log_memory: %v", err)
		}
	}

	if o.PacketLogPostAttack != nil {
		v := *o.PacketLogPostAttack

		if err = d.Set("packet_log_post_attack", v); err != nil {
			return diag.Errorf("error reading packet_log_post_attack: %v", err)
		}
	}

	return nil
}

func getObjectIpsSettings(d *schema.ResourceData, sv string) (*models.IpsSettings, diag.Diagnostics) {
	obj := models.IpsSettings{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("ips_packet_quota"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ips_packet_quota", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IpsPacketQuota = &tmp
		}
	}
	if v1, ok := d.GetOk("packet_log_history"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("packet_log_history", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PacketLogHistory = &tmp
		}
	}
	if v1, ok := d.GetOk("packet_log_memory"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("packet_log_memory", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PacketLogMemory = &tmp
		}
	}
	if v1, ok := d.GetOk("packet_log_post_attack"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("packet_log_post_attack", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PacketLogPostAttack = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectIpsSettings(d *schema.ResourceData, sv string) (*models.IpsSettings, diag.Diagnostics) {
	obj := models.IpsSettings{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
