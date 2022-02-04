// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Wireless Termination Points (WTP) system log server profile.

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

func resourceWirelessControllerSyslogProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Wireless Termination Points (WTP) system log server profile.",

		CreateContext: resourceWirelessControllerSyslogProfileCreate,
		ReadContext:   resourceWirelessControllerSyslogProfileRead,
		UpdateContext: resourceWirelessControllerSyslogProfileUpdate,
		DeleteContext: resourceWirelessControllerSyslogProfileDelete,

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
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"log_level": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debugging"}, false),

				Description: "Lowest level of log messages that FortiAP units send to this server (default = information).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WTP system log server profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"server_addr_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fqdn", "ip"}, false),

				Description: "Syslog server address type (default = ip).",
				Optional:    true,
				Computed:    true,
			},
			"server_fqdn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "FQDN of syslog server that FortiAP units send log messages to.",
				Optional:    true,
				Computed:    true,
			},
			"server_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of syslog server that FortiAP units send log messages to.",
				Optional:    true,
				Computed:    true,
			},
			"server_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port number of syslog server that FortiAP units send log messages to (default = 514).",
				Optional:    true,
				Computed:    true,
			},
			"server_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiAP units to send log messages to a syslog server (default = enable).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerSyslogProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerSyslogProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerSyslogProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerSyslogProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerSyslogProfile")
	}

	return resourceWirelessControllerSyslogProfileRead(ctx, d, meta)
}

func resourceWirelessControllerSyslogProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerSyslogProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerSyslogProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerSyslogProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerSyslogProfile")
	}

	return resourceWirelessControllerSyslogProfileRead(ctx, d, meta)
}

func resourceWirelessControllerSyslogProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerSyslogProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerSyslogProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSyslogProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerSyslogProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerSyslogProfile resource: %v", err)
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

	diags := refreshObjectWirelessControllerSyslogProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWirelessControllerSyslogProfile(d *schema.ResourceData, o *models.WirelessControllerSyslogProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.LogLevel != nil {
		v := *o.LogLevel

		if err = d.Set("log_level", v); err != nil {
			return diag.Errorf("error reading log_level: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ServerAddrType != nil {
		v := *o.ServerAddrType

		if err = d.Set("server_addr_type", v); err != nil {
			return diag.Errorf("error reading server_addr_type: %v", err)
		}
	}

	if o.ServerFqdn != nil {
		v := *o.ServerFqdn

		if err = d.Set("server_fqdn", v); err != nil {
			return diag.Errorf("error reading server_fqdn: %v", err)
		}
	}

	if o.ServerIp != nil {
		v := *o.ServerIp

		if err = d.Set("server_ip", v); err != nil {
			return diag.Errorf("error reading server_ip: %v", err)
		}
	}

	if o.ServerPort != nil {
		v := *o.ServerPort

		if err = d.Set("server_port", v); err != nil {
			return diag.Errorf("error reading server_port: %v", err)
		}
	}

	if o.ServerStatus != nil {
		v := *o.ServerStatus

		if err = d.Set("server_status", v); err != nil {
			return diag.Errorf("error reading server_status: %v", err)
		}
	}

	return nil
}

func getObjectWirelessControllerSyslogProfile(d *schema.ResourceData, sv string) (*models.WirelessControllerSyslogProfile, diag.Diagnostics) {
	obj := models.WirelessControllerSyslogProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("log_level"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_level", sv)
				diags = append(diags, e)
			}
			obj.LogLevel = &v2
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
	if v1, ok := d.GetOk("server_addr_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_addr_type", sv)
				diags = append(diags, e)
			}
			obj.ServerAddrType = &v2
		}
	}
	if v1, ok := d.GetOk("server_fqdn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_fqdn", sv)
				diags = append(diags, e)
			}
			obj.ServerFqdn = &v2
		}
	}
	if v1, ok := d.GetOk("server_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_ip", sv)
				diags = append(diags, e)
			}
			obj.ServerIp = &v2
		}
	}
	if v1, ok := d.GetOk("server_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ServerPort = &tmp
		}
	}
	if v1, ok := d.GetOk("server_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_status", sv)
				diags = append(diags, e)
			}
			obj.ServerStatus = &v2
		}
	}
	return &obj, diags
}
