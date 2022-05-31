// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Access Proxy virtual hosts.

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

func resourceFirewallAccessProxyVirtualHost() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Access Proxy virtual hosts.",

		CreateContext: resourceFirewallAccessProxyVirtualHostCreate,
		ReadContext:   resourceFirewallAccessProxyVirtualHostRead,
		UpdateContext: resourceFirewallAccessProxyVirtualHostUpdate,
		DeleteContext: resourceFirewallAccessProxyVirtualHostDelete,

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
			"host": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "The host name.",
				Optional:    true,
				Computed:    true,
			},
			"host_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"sub-string", "wildcard"}, false),

				Description: "Type of host pattern.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Virtual host name.",
				ForceNew:    true,
				Required:    true,
			},
			"ssl_certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SSL certificate for this host.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallAccessProxyVirtualHostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallAccessProxyVirtualHost resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallAccessProxyVirtualHost(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallAccessProxyVirtualHost(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAccessProxyVirtualHost")
	}

	return resourceFirewallAccessProxyVirtualHostRead(ctx, d, meta)
}

func resourceFirewallAccessProxyVirtualHostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallAccessProxyVirtualHost(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallAccessProxyVirtualHost(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallAccessProxyVirtualHost resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAccessProxyVirtualHost")
	}

	return resourceFirewallAccessProxyVirtualHostRead(ctx, d, meta)
}

func resourceFirewallAccessProxyVirtualHostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallAccessProxyVirtualHost(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallAccessProxyVirtualHost resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxyVirtualHostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallAccessProxyVirtualHost(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAccessProxyVirtualHost resource: %v", err)
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

	diags := refreshObjectFirewallAccessProxyVirtualHost(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallAccessProxyVirtualHost(d *schema.ResourceData, o *models.FirewallAccessProxyVirtualHost, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Host != nil {
		v := *o.Host

		if err = d.Set("host", v); err != nil {
			return diag.Errorf("error reading host: %v", err)
		}
	}

	if o.HostType != nil {
		v := *o.HostType

		if err = d.Set("host_type", v); err != nil {
			return diag.Errorf("error reading host_type: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.SslCertificate != nil {
		v := *o.SslCertificate

		if err = d.Set("ssl_certificate", v); err != nil {
			return diag.Errorf("error reading ssl_certificate: %v", err)
		}
	}

	return nil
}

func getObjectFirewallAccessProxyVirtualHost(d *schema.ResourceData, sv string) (*models.FirewallAccessProxyVirtualHost, diag.Diagnostics) {
	obj := models.FirewallAccessProxyVirtualHost{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("host"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("host", sv)
				diags = append(diags, e)
			}
			obj.Host = &v2
		}
	}
	if v1, ok := d.GetOk("host_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("host_type", sv)
				diags = append(diags, e)
			}
			obj.HostType = &v2
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
	if v1, ok := d.GetOk("ssl_certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_certificate", sv)
				diags = append(diags, e)
			}
			obj.SslCertificate = &v2
		}
	}
	return &obj, diags
}
