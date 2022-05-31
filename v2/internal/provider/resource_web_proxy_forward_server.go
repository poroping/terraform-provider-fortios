// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure forward-server addresses.

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

func resourceWebProxyForwardServer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure forward-server addresses.",

		CreateContext: resourceWebProxyForwardServerCreate,
		ReadContext:   resourceWebProxyForwardServerRead,
		UpdateContext: resourceWebProxyForwardServerUpdate,
		DeleteContext: resourceWebProxyForwardServerDelete,

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
			"addr_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ip", "fqdn"}, false),

				Description: "Address type of the forwarding proxy server: IP or FQDN.",
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
			"fqdn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Forward server Fully Qualified Domain Name (FQDN).",
				Optional:    true,
				Computed:    true,
			},
			"healthcheck": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally.",
				Optional:    true,
				Computed:    true,
			},
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Forward proxy server IP address.",
				Optional:    true,
				Computed:    true,
			},
			"monitor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "URL for forward server health check monitoring (default = http://www.google.com).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Server name.",
				ForceNew:    true,
				Required:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "HTTP authentication password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).",
				Optional:    true,
				Computed:    true,
			},
			"server_down_option": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"block", "pass"}, false),

				Description: "Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination.",
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "HTTP authentication user name.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWebProxyForwardServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WebProxyForwardServer resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWebProxyForwardServer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWebProxyForwardServer(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebProxyForwardServer")
	}

	return resourceWebProxyForwardServerRead(ctx, d, meta)
}

func resourceWebProxyForwardServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebProxyForwardServer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWebProxyForwardServer(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WebProxyForwardServer resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebProxyForwardServer")
	}

	return resourceWebProxyForwardServerRead(ctx, d, meta)
}

func resourceWebProxyForwardServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWebProxyForwardServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WebProxyForwardServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyForwardServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebProxyForwardServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebProxyForwardServer resource: %v", err)
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

	diags := refreshObjectWebProxyForwardServer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWebProxyForwardServer(d *schema.ResourceData, o *models.WebProxyForwardServer, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddrType != nil {
		v := *o.AddrType

		if err = d.Set("addr_type", v); err != nil {
			return diag.Errorf("error reading addr_type: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Fqdn != nil {
		v := *o.Fqdn

		if err = d.Set("fqdn", v); err != nil {
			return diag.Errorf("error reading fqdn: %v", err)
		}
	}

	if o.Healthcheck != nil {
		v := *o.Healthcheck

		if err = d.Set("healthcheck", v); err != nil {
			return diag.Errorf("error reading healthcheck: %v", err)
		}
	}

	if o.Ip != nil {
		v := *o.Ip

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.Monitor != nil {
		v := *o.Monitor

		if err = d.Set("monitor", v); err != nil {
			return diag.Errorf("error reading monitor: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.ServerDownOption != nil {
		v := *o.ServerDownOption

		if err = d.Set("server_down_option", v); err != nil {
			return diag.Errorf("error reading server_down_option: %v", err)
		}
	}

	if o.Username != nil {
		v := *o.Username

		if err = d.Set("username", v); err != nil {
			return diag.Errorf("error reading username: %v", err)
		}
	}

	return nil
}

func getObjectWebProxyForwardServer(d *schema.ResourceData, sv string) (*models.WebProxyForwardServer, diag.Diagnostics) {
	obj := models.WebProxyForwardServer{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("addr_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("addr_type", sv)
				diags = append(diags, e)
			}
			obj.AddrType = &v2
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
	if v1, ok := d.GetOk("fqdn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fqdn", sv)
				diags = append(diags, e)
			}
			obj.Fqdn = &v2
		}
	}
	if v1, ok := d.GetOk("healthcheck"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("healthcheck", sv)
				diags = append(diags, e)
			}
			obj.Healthcheck = &v2
		}
	}
	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
		}
	}
	if v1, ok := d.GetOk("monitor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("monitor", sv)
				diags = append(diags, e)
			}
			obj.Monitor = &v2
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
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port = &tmp
		}
	}
	if v1, ok := d.GetOk("server_down_option"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_down_option", sv)
				diags = append(diags, e)
			}
			obj.ServerDownOption = &v2
		}
	}
	if v1, ok := d.GetOk("username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("username", sv)
				diags = append(diags, e)
			}
			obj.Username = &v2
		}
	}
	return &obj, diags
}
