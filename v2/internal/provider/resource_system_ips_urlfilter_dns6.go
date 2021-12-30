// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS URL filter IPv6 DNS servers.

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

func resourceSystemIpsUrlfilterDns6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS URL filter IPv6 DNS servers.",

		CreateContext: resourceSystemIpsUrlfilterDns6Create,
		ReadContext:   resourceSystemIpsUrlfilterDns6Read,
		UpdateContext: resourceSystemIpsUrlfilterDns6Update,
		DeleteContext: resourceSystemIpsUrlfilterDns6Delete,

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
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"address6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 address of DNS server.",
				ForceNew:         true,
				Required:         true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this server for IPv6 DNS queries.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemIpsUrlfilterDns6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "address6"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemIpsUrlfilterDns6 resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemIpsUrlfilterDns6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemIpsUrlfilterDns6(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIpsUrlfilterDns6")
	}

	return resourceSystemIpsUrlfilterDns6Read(ctx, d, meta)
}

func resourceSystemIpsUrlfilterDns6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemIpsUrlfilterDns6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemIpsUrlfilterDns6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemIpsUrlfilterDns6 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIpsUrlfilterDns6")
	}

	return resourceSystemIpsUrlfilterDns6Read(ctx, d, meta)
}

func resourceSystemIpsUrlfilterDns6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemIpsUrlfilterDns6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemIpsUrlfilterDns6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpsUrlfilterDns6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemIpsUrlfilterDns6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemIpsUrlfilterDns6 resource: %v", err)
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

	diags := refreshObjectSystemIpsUrlfilterDns6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemIpsUrlfilterDns6(d *schema.ResourceData, o *models.SystemIpsUrlfilterDns6, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Address6 != nil {
		v := *o.Address6

		if err = d.Set("address6", v); err != nil {
			return diag.Errorf("error reading address6: %v", err)
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

func getObjectSystemIpsUrlfilterDns6(d *schema.ResourceData, sv string) (*models.SystemIpsUrlfilterDns6, diag.Diagnostics) {
	obj := models.SystemIpsUrlfilterDns6{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("address6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("address6", sv)
				diags = append(diags, e)
			}
			obj.Address6 = &v2
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
