// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Bluetooth Low Energy profile.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceWirelessControllerBleProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Bluetooth Low Energy profile.",

		CreateContext: resourceWirelessControllerBleProfileCreate,
		ReadContext:   resourceWirelessControllerBleProfileRead,
		UpdateContext: resourceWirelessControllerBleProfileUpdate,
		DeleteContext: resourceWirelessControllerBleProfileDelete,

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
			"advertising": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Advertising type.",
				Optional:         true,
				Computed:         true,
			},
			"beacon_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(40, 3500),

				Description: "Beacon interval (default = 100 msec).",
				Optional:    true,
				Computed:    true,
			},
			"ble_scanning": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Bluetooth Low Energy (BLE) scanning.",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"eddystone_instance": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 12),

				Description: "Eddystone instance ID.",
				Optional:    true,
				Computed:    true,
			},
			"eddystone_namespace": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 20),

				Description: "Eddystone namespace ID.",
				Optional:    true,
				Computed:    true,
			},
			"eddystone_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Eddystone URL.",
				Optional:    true,
				Computed:    true,
			},
			"ibeacon_uuid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Optional:    true,
				Computed:    true,
			},
			"major_id": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Major ID.",
				Optional:    true,
				Computed:    true,
			},
			"minor_id": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Minor ID.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Bluetooth Low Energy profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"txpower": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}, false),

				Description: "Transmit power level (default = 0).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerBleProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerBleProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerBleProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerBleProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerBleProfile")
	}

	return resourceWirelessControllerBleProfileRead(ctx, d, meta)
}

func resourceWirelessControllerBleProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerBleProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerBleProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerBleProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerBleProfile")
	}

	return resourceWirelessControllerBleProfileRead(ctx, d, meta)
}

func resourceWirelessControllerBleProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerBleProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerBleProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerBleProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerBleProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerBleProfile resource: %v", err)
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

	diags := refreshObjectWirelessControllerBleProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWirelessControllerBleProfile(d *schema.ResourceData, o *models.WirelessControllerBleProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Advertising != nil {
		v := *o.Advertising

		if err = d.Set("advertising", v); err != nil {
			return diag.Errorf("error reading advertising: %v", err)
		}
	}

	if o.BeaconInterval != nil {
		v := *o.BeaconInterval

		if err = d.Set("beacon_interval", v); err != nil {
			return diag.Errorf("error reading beacon_interval: %v", err)
		}
	}

	if o.BleScanning != nil {
		v := *o.BleScanning

		if err = d.Set("ble_scanning", v); err != nil {
			return diag.Errorf("error reading ble_scanning: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.EddystoneInstance != nil {
		v := *o.EddystoneInstance

		if err = d.Set("eddystone_instance", v); err != nil {
			return diag.Errorf("error reading eddystone_instance: %v", err)
		}
	}

	if o.EddystoneNamespace != nil {
		v := *o.EddystoneNamespace

		if err = d.Set("eddystone_namespace", v); err != nil {
			return diag.Errorf("error reading eddystone_namespace: %v", err)
		}
	}

	if o.EddystoneUrl != nil {
		v := *o.EddystoneUrl

		if err = d.Set("eddystone_url", v); err != nil {
			return diag.Errorf("error reading eddystone_url: %v", err)
		}
	}

	if o.IbeaconUuid != nil {
		v := *o.IbeaconUuid

		if err = d.Set("ibeacon_uuid", v); err != nil {
			return diag.Errorf("error reading ibeacon_uuid: %v", err)
		}
	}

	if o.MajorId != nil {
		v := *o.MajorId

		if err = d.Set("major_id", v); err != nil {
			return diag.Errorf("error reading major_id: %v", err)
		}
	}

	if o.MinorId != nil {
		v := *o.MinorId

		if err = d.Set("minor_id", v); err != nil {
			return diag.Errorf("error reading minor_id: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Txpower != nil {
		v := *o.Txpower

		if err = d.Set("txpower", v); err != nil {
			return diag.Errorf("error reading txpower: %v", err)
		}
	}

	return nil
}

func getObjectWirelessControllerBleProfile(d *schema.ResourceData, sv string) (*models.WirelessControllerBleProfile, diag.Diagnostics) {
	obj := models.WirelessControllerBleProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("advertising"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("advertising", sv)
				diags = append(diags, e)
			}
			obj.Advertising = &v2
		}
	}
	if v1, ok := d.GetOk("beacon_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("beacon_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BeaconInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("ble_scanning"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ble_scanning", sv)
				diags = append(diags, e)
			}
			obj.BleScanning = &v2
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("eddystone_instance"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eddystone_instance", sv)
				diags = append(diags, e)
			}
			obj.EddystoneInstance = &v2
		}
	}
	if v1, ok := d.GetOk("eddystone_namespace"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eddystone_namespace", sv)
				diags = append(diags, e)
			}
			obj.EddystoneNamespace = &v2
		}
	}
	if v1, ok := d.GetOk("eddystone_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eddystone_url", sv)
				diags = append(diags, e)
			}
			obj.EddystoneUrl = &v2
		}
	}
	if v1, ok := d.GetOk("ibeacon_uuid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ibeacon_uuid", sv)
				diags = append(diags, e)
			}
			obj.IbeaconUuid = &v2
		}
	}
	if v1, ok := d.GetOk("major_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("major_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MajorId = &tmp
		}
	}
	if v1, ok := d.GetOk("minor_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("minor_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MinorId = &tmp
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
	if v1, ok := d.GetOk("txpower"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("txpower", sv)
				diags = append(diags, e)
			}
			obj.Txpower = &v2
		}
	}
	return &obj, diags
}
