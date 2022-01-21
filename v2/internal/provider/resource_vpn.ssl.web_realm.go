// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Realm.

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

func resourceVpnSslWebRealm() *schema.Resource {
	return &schema.Resource{
		Description: "Realm.",

		CreateContext: resourceVpnSslWebRealmCreate,
		ReadContext:   resourceVpnSslWebRealmRead,
		UpdateContext: resourceVpnSslWebRealmUpdate,
		DeleteContext: resourceVpnSslWebRealmDelete,

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
			"login_page": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32768),

				Description: "Replacement HTML for SSL-VPN login page.",
				Optional:    true,
				Computed:    true,
			},
			"max_concurrent_user": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Maximum concurrent users (0 - 65535, 0 means unlimited).",
				Optional:    true,
				Computed:    true,
			},
			"nas_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address used as a NAS-IP to communicate with the RADIUS server.",
				Optional:    true,
				Computed:    true,
			},
			"radius_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).",
				Optional:    true,
				Computed:    true,
			},
			"radius_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "RADIUS server associated with realm.",
				Optional:    true,
				Computed:    true,
			},
			"url_path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "URL path to access SSL-VPN login page.",
				ForceNew:    true,
				Required:    true,
			},
			"virtual_host": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Virtual host name for realm.",
				Optional:    true,
				Computed:    true,
			},
			"virtual_host_only": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable enforcement of virtual host method for SSL-VPN client access.",
				Optional:    true,
				Computed:    true,
			},
			"virtual_host_server_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of the server certificate to used for this realm.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpnSslWebRealmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "url_path"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating VpnSslWebRealm resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnSslWebRealm(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnSslWebRealm(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnSslWebRealm")
	}

	return resourceVpnSslWebRealmRead(ctx, d, meta)
}

func resourceVpnSslWebRealmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnSslWebRealm(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnSslWebRealm(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnSslWebRealm resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnSslWebRealm")
	}

	return resourceVpnSslWebRealmRead(ctx, d, meta)
}

func resourceVpnSslWebRealmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVpnSslWebRealm(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnSslWebRealm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebRealmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnSslWebRealm(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnSslWebRealm resource: %v", err)
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

	diags := refreshObjectVpnSslWebRealm(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectVpnSslWebRealm(d *schema.ResourceData, o *models.VpnSslWebRealm, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.LoginPage != nil {
		v := *o.LoginPage

		if err = d.Set("login_page", v); err != nil {
			return diag.Errorf("error reading login_page: %v", err)
		}
	}

	if o.MaxConcurrentUser != nil {
		v := *o.MaxConcurrentUser

		if err = d.Set("max_concurrent_user", v); err != nil {
			return diag.Errorf("error reading max_concurrent_user: %v", err)
		}
	}

	if o.NasIp != nil {
		v := *o.NasIp

		if err = d.Set("nas_ip", v); err != nil {
			return diag.Errorf("error reading nas_ip: %v", err)
		}
	}

	if o.RadiusPort != nil {
		v := *o.RadiusPort

		if err = d.Set("radius_port", v); err != nil {
			return diag.Errorf("error reading radius_port: %v", err)
		}
	}

	if o.RadiusServer != nil {
		v := *o.RadiusServer

		if err = d.Set("radius_server", v); err != nil {
			return diag.Errorf("error reading radius_server: %v", err)
		}
	}

	if o.UrlPath != nil {
		v := *o.UrlPath

		if err = d.Set("url_path", v); err != nil {
			return diag.Errorf("error reading url_path: %v", err)
		}
	}

	if o.VirtualHost != nil {
		v := *o.VirtualHost

		if err = d.Set("virtual_host", v); err != nil {
			return diag.Errorf("error reading virtual_host: %v", err)
		}
	}

	if o.VirtualHostOnly != nil {
		v := *o.VirtualHostOnly

		if err = d.Set("virtual_host_only", v); err != nil {
			return diag.Errorf("error reading virtual_host_only: %v", err)
		}
	}

	if o.VirtualHostServerCert != nil {
		v := *o.VirtualHostServerCert

		if err = d.Set("virtual_host_server_cert", v); err != nil {
			return diag.Errorf("error reading virtual_host_server_cert: %v", err)
		}
	}

	return nil
}

func getObjectVpnSslWebRealm(d *schema.ResourceData, sv string) (*models.VpnSslWebRealm, diag.Diagnostics) {
	obj := models.VpnSslWebRealm{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("login_page"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("login_page", sv)
				diags = append(diags, e)
			}
			obj.LoginPage = &v2
		}
	}
	if v1, ok := d.GetOk("max_concurrent_user"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_concurrent_user", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxConcurrentUser = &tmp
		}
	}
	if v1, ok := d.GetOk("nas_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("nas_ip", sv)
				diags = append(diags, e)
			}
			obj.NasIp = &v2
		}
	}
	if v1, ok := d.GetOk("radius_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("radius_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RadiusPort = &tmp
		}
	}
	if v1, ok := d.GetOk("radius_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("radius_server", sv)
				diags = append(diags, e)
			}
			obj.RadiusServer = &v2
		}
	}
	if v1, ok := d.GetOk("url_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("url_path", sv)
				diags = append(diags, e)
			}
			obj.UrlPath = &v2
		}
	}
	if v1, ok := d.GetOk("virtual_host"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("virtual_host", sv)
				diags = append(diags, e)
			}
			obj.VirtualHost = &v2
		}
	}
	if v1, ok := d.GetOk("virtual_host_only"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("virtual_host_only", sv)
				diags = append(diags, e)
			}
			obj.VirtualHostOnly = &v2
		}
	}
	if v1, ok := d.GetOk("virtual_host_server_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("virtual_host_server_cert", sv)
				diags = append(diags, e)
			}
			obj.VirtualHostServerCert = &v2
		}
	}
	return &obj, diags
}
