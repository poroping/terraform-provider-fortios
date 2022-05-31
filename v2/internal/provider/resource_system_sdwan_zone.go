// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SD-WAN zones.

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

func resourceSystemSdwanZone() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SD-WAN zones.",

		CreateContext: resourceSystemSdwanZoneCreate,
		ReadContext:   resourceSystemSdwanZoneRead,
		UpdateContext: resourceSystemSdwanZoneUpdate,
		DeleteContext: resourceSystemSdwanZoneDelete,

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
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Zone name.",
				ForceNew:    true,
				Required:    true,
			},
			"service_sla_tie_break": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"cfg-order", "fib-best-match", "input-device"}, false),

				Description: "Method of selecting member if more than one meets the SLA.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemSdwanZoneCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemSdwanZone resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemSdwanZone(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSdwanZone(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSdwanZone")
	}

	return resourceSystemSdwanZoneRead(ctx, d, meta)
}

func resourceSystemSdwanZoneUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSdwanZone(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSdwanZone(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSdwanZone resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSdwanZone")
	}

	return resourceSystemSdwanZoneRead(ctx, d, meta)
}

func resourceSystemSdwanZoneDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSdwanZone(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSdwanZone resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanZoneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSdwanZone(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSdwanZone resource: %v", err)
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

	diags := refreshObjectSystemSdwanZone(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemSdwanZone(d *schema.ResourceData, o *models.SystemSdwanZone, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ServiceSlaTieBreak != nil {
		v := *o.ServiceSlaTieBreak

		if err = d.Set("service_sla_tie_break", v); err != nil {
			return diag.Errorf("error reading service_sla_tie_break: %v", err)
		}
	}

	return nil
}

func getObjectSystemSdwanZone(d *schema.ResourceData, sv string) (*models.SystemSdwanZone, diag.Diagnostics) {
	obj := models.SystemSdwanZone{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("service_sla_tie_break"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("service_sla_tie_break", sv)
				diags = append(diags, e)
			}
			obj.ServiceSlaTieBreak = &v2
		}
	}
	return &obj, diags
}
