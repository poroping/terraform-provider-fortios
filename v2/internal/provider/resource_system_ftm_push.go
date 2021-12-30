// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiToken Mobile push services.

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

func resourceSystemFtmPush() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiToken Mobile push services.",

		CreateContext: resourceSystemFtmPushCreate,
		ReadContext:   resourceSystemFtmPushRead,
		UpdateContext: resourceSystemFtmPushUpdate,
		DeleteContext: resourceSystemFtmPushDelete,

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
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "IPv4 address or domain name of FortiToken Mobile push services server.",
				Optional:    true,
				Computed:    true,
			},
			"server_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of the server certificate to be used for SSL (default = Fortinet_Factory).",
				Optional:    true,
				Computed:    true,
			},
			"server_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).",
				Optional:    true,
				Computed:    true,
			},
			"server_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the use of FortiToken Mobile push services.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemFtmPushCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemFtmPush(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemFtmPush(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemFtmPush")
	}

	return resourceSystemFtmPushRead(ctx, d, meta)
}

func resourceSystemFtmPushUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemFtmPush(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemFtmPush(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemFtmPush resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemFtmPush")
	}

	return resourceSystemFtmPushRead(ctx, d, meta)
}

func resourceSystemFtmPushDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemFtmPush(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemFtmPush resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFtmPushRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemFtmPush(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemFtmPush resource: %v", err)
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

	diags := refreshObjectSystemFtmPush(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemFtmPush(d *schema.ResourceData, o *models.SystemFtmPush, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.ServerCert != nil {
		v := *o.ServerCert

		if err = d.Set("server_cert", v); err != nil {
			return diag.Errorf("error reading server_cert: %v", err)
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

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func getObjectSystemFtmPush(d *schema.ResourceData, sv string) (*models.SystemFtmPush, diag.Diagnostics) {
	obj := models.SystemFtmPush{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("server", sv)
				diags = append(diags, e)
			}
			obj.Server = &v2
		}
	}
	if v1, ok := d.GetOk("server_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("server_cert", sv)
				diags = append(diags, e)
			}
			obj.ServerCert = &v2
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
