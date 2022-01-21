// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure system PTP information.

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

func resourceSystemPtp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure system PTP information.",

		CreateContext: resourceSystemPtpCreate,
		ReadContext:   resourceSystemPtpRead,
		UpdateContext: resourceSystemPtpUpdate,
		DeleteContext: resourceSystemPtpDelete,

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
			"delay_mechanism": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"E2E", "P2P"}, false),

				Description: "End to end delay detection or peer to peer delay detection.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "PTP client will reply through this interface.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"multicast", "hybrid"}, false),

				Description: "Multicast transmission or hybrid transmission.",
				Optional:    true,
				Computed:    true,
			},
			"request_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 6),

				Description: "The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.",
				Optional:    true,
				Computed:    true,
			},
			"server_interface": {
				Type:        schema.TypeList,
				Description: "FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delay_mechanism": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"E2E", "P2P"}, false),

							Description: "End to end delay detection or peer to peer delay detection.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"server_interface_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"server_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable setting the FortiGate system time by synchronizing with an PTP Server.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemPtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemPtp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemPtp(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemPtp")
	}

	return resourceSystemPtpRead(ctx, d, meta)
}

func resourceSystemPtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemPtp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemPtp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemPtp resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemPtp")
	}

	return resourceSystemPtpRead(ctx, d, meta)
}

func resourceSystemPtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemPtp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemPtp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemPtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemPtp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemPtp resource: %v", err)
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

	diags := refreshObjectSystemPtp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemPtpServerInterface(v *[]models.SystemPtpServerInterface, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.DelayMechanism; tmp != nil {
				v["delay_mechanism"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.ServerInterfaceName; tmp != nil {
				v["server_interface_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemPtp(d *schema.ResourceData, o *models.SystemPtp, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DelayMechanism != nil {
		v := *o.DelayMechanism

		if err = d.Set("delay_mechanism", v); err != nil {
			return diag.Errorf("error reading delay_mechanism: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.RequestInterval != nil {
		v := *o.RequestInterval

		if err = d.Set("request_interval", v); err != nil {
			return diag.Errorf("error reading request_interval: %v", err)
		}
	}

	if o.ServerInterface != nil {
		if err = d.Set("server_interface", flattenSystemPtpServerInterface(o.ServerInterface, sort)); err != nil {
			return diag.Errorf("error reading server_interface: %v", err)
		}
	}

	if o.ServerMode != nil {
		v := *o.ServerMode

		if err = d.Set("server_mode", v); err != nil {
			return diag.Errorf("error reading server_mode: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func expandSystemPtpServerInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemPtpServerInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemPtpServerInterface

	for i := range l {
		tmp := models.SystemPtpServerInterface{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.delay_mechanism", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DelayMechanism = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server_interface_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ServerInterfaceName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemPtp(d *schema.ResourceData, sv string) (*models.SystemPtp, diag.Diagnostics) {
	obj := models.SystemPtp{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("delay_mechanism"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("delay_mechanism", sv)
				diags = append(diags, e)
			}
			obj.DelayMechanism = &v2
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
		}
	}
	if v1, ok := d.GetOk("request_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("request_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RequestInterval = &tmp
		}
	}
	if v, ok := d.GetOk("server_interface"); ok {
		if !utils.CheckVer(sv, "v7.0.1", "") {
			e := utils.AttributeVersionWarning("server_interface", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemPtpServerInterface(d, v, "server_interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ServerInterface = t
		}
	} else if d.HasChange("server_interface") {
		old, new := d.GetChange("server_interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ServerInterface = &[]models.SystemPtpServerInterface{}
		}
	}
	if v1, ok := d.GetOk("server_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("server_mode", sv)
				diags = append(diags, e)
			}
			obj.ServerMode = &v2
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
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemPtp(d *schema.ResourceData, sv string) (*models.SystemPtp, diag.Diagnostics) {
	obj := models.SystemPtp{}
	diags := diag.Diagnostics{}

	obj.ServerInterface = &[]models.SystemPtpServerInterface{}

	return &obj, diags
}
