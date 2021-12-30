// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS system settings.

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

func resourceSystemIps() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS system settings.",

		CreateContext: resourceSystemIpsCreate,
		ReadContext:   resourceSystemIpsRead,
		UpdateContext: resourceSystemIpsUpdate,
		DeleteContext: resourceSystemIpsDelete,

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
			"override_signature_hold_by_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold.",
				Optional:    true,
				Computed:    true,
			},
			"signature_hold_time": {
				Type: schema.TypeString,

				Description: "Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).\n",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemIpsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemIps(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemIps(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIps")
	}

	return resourceSystemIpsRead(ctx, d, meta)
}

func resourceSystemIpsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemIps(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemIps(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemIps resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIps")
	}

	return resourceSystemIpsRead(ctx, d, meta)
}

func resourceSystemIpsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemIps(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemIps resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemIps(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemIps resource: %v", err)
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

	diags := refreshObjectSystemIps(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemIps(d *schema.ResourceData, o *models.SystemIps, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.OverrideSignatureHoldById != nil {
		v := *o.OverrideSignatureHoldById

		if err = d.Set("override_signature_hold_by_id", v); err != nil {
			return diag.Errorf("error reading override_signature_hold_by_id: %v", err)
		}
	}

	if o.SignatureHoldTime != nil {
		v := *o.SignatureHoldTime

		if err = d.Set("signature_hold_time", v); err != nil {
			return diag.Errorf("error reading signature_hold_time: %v", err)
		}
	}

	return nil
}

func getObjectSystemIps(d *schema.ResourceData, sv string) (*models.SystemIps, diag.Diagnostics) {
	obj := models.SystemIps{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("override_signature_hold_by_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_signature_hold_by_id", sv)
				diags = append(diags, e)
			}
			obj.OverrideSignatureHoldById = &v2
		}
	}
	if v1, ok := d.GetOk("signature_hold_time"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("signature_hold_time", sv)
				diags = append(diags, e)
			}
			obj.SignatureHoldTime = &v2
		}
	}
	return &obj, diags
}
