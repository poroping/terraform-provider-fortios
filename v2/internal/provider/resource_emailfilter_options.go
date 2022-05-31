// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure AntiSpam options.

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

func resourceEmailfilterOptions() *schema.Resource {
	return &schema.Resource{
		Description: "Configure AntiSpam options.",

		CreateContext: resourceEmailfilterOptionsCreate,
		ReadContext:   resourceEmailfilterOptionsRead,
		UpdateContext: resourceEmailfilterOptionsUpdate,
		DeleteContext: resourceEmailfilterOptionsDelete,

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
			"dns_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),

				Description: "DNS query time out (1 - 30 sec).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceEmailfilterOptionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectEmailfilterOptions(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateEmailfilterOptions(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EmailfilterOptions")
	}

	return resourceEmailfilterOptionsRead(ctx, d, meta)
}

func resourceEmailfilterOptionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectEmailfilterOptions(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateEmailfilterOptions(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating EmailfilterOptions resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EmailfilterOptions")
	}

	return resourceEmailfilterOptionsRead(ctx, d, meta)
}

func resourceEmailfilterOptionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectEmailfilterOptions(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateEmailfilterOptions(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting EmailfilterOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadEmailfilterOptions(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading EmailfilterOptions resource: %v", err)
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

	diags := refreshObjectEmailfilterOptions(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectEmailfilterOptions(d *schema.ResourceData, o *models.EmailfilterOptions, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DnsTimeout != nil {
		v := *o.DnsTimeout

		if err = d.Set("dns_timeout", v); err != nil {
			return diag.Errorf("error reading dns_timeout: %v", err)
		}
	}

	return nil
}

func getObjectEmailfilterOptions(d *schema.ResourceData, sv string) (*models.EmailfilterOptions, diag.Diagnostics) {
	obj := models.EmailfilterOptions{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("dns_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DnsTimeout = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectEmailfilterOptions(d *schema.ResourceData, sv string) (*models.EmailfilterOptions, diag.Diagnostics) {
	obj := models.EmailfilterOptions{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
