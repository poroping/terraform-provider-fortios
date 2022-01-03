// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure endpoint control settings.

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

func resourceEndpointControlSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure endpoint control settings.",

		CreateContext: resourceEndpointControlSettingsCreate,
		ReadContext:   resourceEndpointControlSettingsRead,
		UpdateContext: resourceEndpointControlSettingsUpdate,
		DeleteContext: resourceEndpointControlSettingsDelete,

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
			"forticlient_disconnect_unsupported_client": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable disconnecting of unsupported FortiClient endpoints.",
				Optional:    true,
				Computed:    true,
			},
			"forticlient_keepalive_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(20, 300),

				Description: "Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"forticlient_sys_update_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 1440),

				Description: "Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).",
				Optional:    true,
				Computed:    true,
			},
			"forticlient_user_avatar": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable uploading FortiClient user avatars.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceEndpointControlSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectEndpointControlSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateEndpointControlSettings(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EndpointControlSettings")
	}

	return resourceEndpointControlSettingsRead(ctx, d, meta)
}

func resourceEndpointControlSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectEndpointControlSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateEndpointControlSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating EndpointControlSettings resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EndpointControlSettings")
	}

	return resourceEndpointControlSettingsRead(ctx, d, meta)
}

func resourceEndpointControlSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectEndpointControlSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateEndpointControlSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting EndpointControlSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEndpointControlSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadEndpointControlSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading EndpointControlSettings resource: %v", err)
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

	diags := refreshObjectEndpointControlSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectEndpointControlSettings(d *schema.ResourceData, o *models.EndpointControlSettings, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ForticlientDisconnectUnsupportedClient != nil {
		v := *o.ForticlientDisconnectUnsupportedClient

		if err = d.Set("forticlient_disconnect_unsupported_client", v); err != nil {
			return diag.Errorf("error reading forticlient_disconnect_unsupported_client: %v", err)
		}
	}

	if o.ForticlientKeepaliveInterval != nil {
		v := *o.ForticlientKeepaliveInterval

		if err = d.Set("forticlient_keepalive_interval", v); err != nil {
			return diag.Errorf("error reading forticlient_keepalive_interval: %v", err)
		}
	}

	if o.ForticlientSysUpdateInterval != nil {
		v := *o.ForticlientSysUpdateInterval

		if err = d.Set("forticlient_sys_update_interval", v); err != nil {
			return diag.Errorf("error reading forticlient_sys_update_interval: %v", err)
		}
	}

	if o.ForticlientUserAvatar != nil {
		v := *o.ForticlientUserAvatar

		if err = d.Set("forticlient_user_avatar", v); err != nil {
			return diag.Errorf("error reading forticlient_user_avatar: %v", err)
		}
	}

	return nil
}

func getObjectEndpointControlSettings(d *schema.ResourceData, sv string) (*models.EndpointControlSettings, diag.Diagnostics) {
	obj := models.EndpointControlSettings{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("forticlient_disconnect_unsupported_client"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forticlient_disconnect_unsupported_client", sv)
				diags = append(diags, e)
			}
			obj.ForticlientDisconnectUnsupportedClient = &v2
		}
	}
	if v1, ok := d.GetOk("forticlient_keepalive_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forticlient_keepalive_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ForticlientKeepaliveInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("forticlient_sys_update_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forticlient_sys_update_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ForticlientSysUpdateInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("forticlient_user_avatar"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forticlient_user_avatar", sv)
				diags = append(diags, e)
			}
			obj.ForticlientUserAvatar = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectEndpointControlSettings(d *schema.ResourceData, sv string) (*models.EndpointControlSettings, diag.Diagnostics) {
	obj := models.EndpointControlSettings{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
