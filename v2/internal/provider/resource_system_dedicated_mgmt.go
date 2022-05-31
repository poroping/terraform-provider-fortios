// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure dedicated management.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceSystemDedicatedMgmt() *schema.Resource {
	return &schema.Resource{
		Description: "Configure dedicated management.",

		CreateContext: resourceSystemDedicatedMgmtCreate,
		ReadContext:   resourceSystemDedicatedMgmtRead,
		UpdateContext: resourceSystemDedicatedMgmtUpdate,
		DeleteContext: resourceSystemDedicatedMgmtDelete,

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
			"default_gateway": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Default gateway for dedicated management interface.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_end_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "DHCP end IP for dedicated management.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_netmask": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Netmask,

				Description: "DHCP netmask.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DHCP server on management interface.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_start_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "DHCP start IP for dedicated management.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Dedicated management interface.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable dedicated management.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemDedicatedMgmtCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemDedicatedMgmt(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemDedicatedMgmt(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemDedicatedMgmt")
	}

	return resourceSystemDedicatedMgmtRead(ctx, d, meta)
}

func resourceSystemDedicatedMgmtUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemDedicatedMgmt(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemDedicatedMgmt(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemDedicatedMgmt resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemDedicatedMgmt")
	}

	return resourceSystemDedicatedMgmtRead(ctx, d, meta)
}

func resourceSystemDedicatedMgmtDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemDedicatedMgmt(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemDedicatedMgmt(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemDedicatedMgmt resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDedicatedMgmtRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemDedicatedMgmt(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemDedicatedMgmt resource: %v", err)
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

	diags := refreshObjectSystemDedicatedMgmt(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemDedicatedMgmt(d *schema.ResourceData, o *models.SystemDedicatedMgmt, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DefaultGateway != nil {
		v := *o.DefaultGateway

		if err = d.Set("default_gateway", v); err != nil {
			return diag.Errorf("error reading default_gateway: %v", err)
		}
	}

	if o.DhcpEndIp != nil {
		v := *o.DhcpEndIp

		if err = d.Set("dhcp_end_ip", v); err != nil {
			return diag.Errorf("error reading dhcp_end_ip: %v", err)
		}
	}

	if o.DhcpNetmask != nil {
		v := *o.DhcpNetmask

		if err = d.Set("dhcp_netmask", v); err != nil {
			return diag.Errorf("error reading dhcp_netmask: %v", err)
		}
	}

	if o.DhcpServer != nil {
		v := *o.DhcpServer

		if err = d.Set("dhcp_server", v); err != nil {
			return diag.Errorf("error reading dhcp_server: %v", err)
		}
	}

	if o.DhcpStartIp != nil {
		v := *o.DhcpStartIp

		if err = d.Set("dhcp_start_ip", v); err != nil {
			return diag.Errorf("error reading dhcp_start_ip: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
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

func getObjectSystemDedicatedMgmt(d *schema.ResourceData, sv string) (*models.SystemDedicatedMgmt, diag.Diagnostics) {
	obj := models.SystemDedicatedMgmt{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("default_gateway"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_gateway", sv)
				diags = append(diags, e)
			}
			obj.DefaultGateway = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_end_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_end_ip", sv)
				diags = append(diags, e)
			}
			obj.DhcpEndIp = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_netmask"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_netmask", sv)
				diags = append(diags, e)
			}
			obj.DhcpNetmask = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_server", sv)
				diags = append(diags, e)
			}
			obj.DhcpServer = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_start_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_start_ip", sv)
				diags = append(diags, e)
			}
			obj.DhcpStartIp = &v2
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
func getEmptyObjectSystemDedicatedMgmt(d *schema.ResourceData, sv string) (*models.SystemDedicatedMgmt, diag.Diagnostics) {
	obj := models.SystemDedicatedMgmt{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
