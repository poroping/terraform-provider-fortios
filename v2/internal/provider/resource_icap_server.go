// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure ICAP servers.

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

func resourceIcapServer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure ICAP servers.",

		CreateContext: resourceIcapServerCreate,
		ReadContext:   resourceIcapServerRead,
		UpdateContext: resourceIcapServerUpdate,
		DeleteContext: resourceIcapServerDelete,

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
				ValidateFunc: validation.StringInSlice([]string{"ip4", "ip6", "fqdn"}, false),

				Description: "Address type of the remote ICAP server: IPv4, IPv6 or FQDN.",
				Optional:    true,
				Computed:    true,
			},
			"fqdn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "ICAP remote server Fully Qualified Domain Name (FQDN).",
				Optional:    true,
				Computed:    true,
			},
			"healthcheck": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable ICAP remote server health checking. Attempts to connect to the remote ICAP server to verify that the server is operating normally.",
				Optional:    true,
				Computed:    true,
			},
			"healthcheck_service": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "ICAP Service name to use for health checks.",
				Optional:    true,
				Computed:    true,
			},
			"ip_address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 address of the ICAP server.",
				Optional:    true,
				Computed:    true,
			},
			"ip_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"4", "6"}, false),

				Description: "IP version.",
				Optional:    true,
				Computed:    true,
			},
			"ip6_address": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 address of the ICAP server.",
				Optional:         true,
				Computed:         true,
			},
			"max_connections": {
				Type: schema.TypeInt,

				Description: "Maximum number of concurrent connections to ICAP server (unlimited = 0, default = 100). Must not be less than wad-worker-count.",
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
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "ICAP server port.",
				Optional:    true,
				Computed:    true,
			},
			"secure": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable secure connection to ICAP server.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "CA certificate name.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceIcapServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating IcapServer resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectIcapServer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateIcapServer(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IcapServer")
	}

	return resourceIcapServerRead(ctx, d, meta)
}

func resourceIcapServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectIcapServer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateIcapServer(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating IcapServer resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IcapServer")
	}

	return resourceIcapServerRead(ctx, d, meta)
}

func resourceIcapServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteIcapServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting IcapServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIcapServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadIcapServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IcapServer resource: %v", err)
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

	diags := refreshObjectIcapServer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectIcapServer(d *schema.ResourceData, o *models.IcapServer, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddrType != nil {
		v := *o.AddrType

		if err = d.Set("addr_type", v); err != nil {
			return diag.Errorf("error reading addr_type: %v", err)
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

	if o.HealthcheckService != nil {
		v := *o.HealthcheckService

		if err = d.Set("healthcheck_service", v); err != nil {
			return diag.Errorf("error reading healthcheck_service: %v", err)
		}
	}

	if o.IpAddress != nil {
		v := *o.IpAddress

		if err = d.Set("ip_address", v); err != nil {
			return diag.Errorf("error reading ip_address: %v", err)
		}
	}

	if o.IpVersion != nil {
		v := *o.IpVersion

		if err = d.Set("ip_version", v); err != nil {
			return diag.Errorf("error reading ip_version: %v", err)
		}
	}

	if o.Ip6Address != nil {
		v := *o.Ip6Address

		if err = d.Set("ip6_address", v); err != nil {
			return diag.Errorf("error reading ip6_address: %v", err)
		}
	}

	if o.MaxConnections != nil {
		v := *o.MaxConnections

		if err = d.Set("max_connections", v); err != nil {
			return diag.Errorf("error reading max_connections: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.Secure != nil {
		v := *o.Secure

		if err = d.Set("secure", v); err != nil {
			return diag.Errorf("error reading secure: %v", err)
		}
	}

	if o.SslCert != nil {
		v := *o.SslCert

		if err = d.Set("ssl_cert", v); err != nil {
			return diag.Errorf("error reading ssl_cert: %v", err)
		}
	}

	return nil
}

func getObjectIcapServer(d *schema.ResourceData, sv string) (*models.IcapServer, diag.Diagnostics) {
	obj := models.IcapServer{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("addr_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("addr_type", sv)
				diags = append(diags, e)
			}
			obj.AddrType = &v2
		}
	}
	if v1, ok := d.GetOk("fqdn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("fqdn", sv)
				diags = append(diags, e)
			}
			obj.Fqdn = &v2
		}
	}
	if v1, ok := d.GetOk("healthcheck"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("healthcheck", sv)
				diags = append(diags, e)
			}
			obj.Healthcheck = &v2
		}
	}
	if v1, ok := d.GetOk("healthcheck_service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("healthcheck_service", sv)
				diags = append(diags, e)
			}
			obj.HealthcheckService = &v2
		}
	}
	if v1, ok := d.GetOk("ip_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_address", sv)
				diags = append(diags, e)
			}
			obj.IpAddress = &v2
		}
	}
	if v1, ok := d.GetOk("ip_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.0") {
				e := utils.AttributeVersionWarning("ip_version", sv)
				diags = append(diags, e)
			}
			obj.IpVersion = &v2
		}
	}
	if v1, ok := d.GetOk("ip6_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_address", sv)
				diags = append(diags, e)
			}
			obj.Ip6Address = &v2
		}
	}
	if v1, ok := d.GetOk("max_connections"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_connections", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxConnections = &tmp
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
	if v1, ok := d.GetOk("secure"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("secure", sv)
				diags = append(diags, e)
			}
			obj.Secure = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ssl_cert", sv)
				diags = append(diags, e)
			}
			obj.SslCert = &v2
		}
	}
	return &obj, diags
}
