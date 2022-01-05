// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure additional port mappings for Internet Services.

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

func resourceFirewallInternetServiceAppend() *schema.Resource {
	return &schema.Resource{
		Description: "Configure additional port mappings for Internet Services.",

		CreateContext: resourceFirewallInternetServiceAppendCreate,
		ReadContext:   resourceFirewallInternetServiceAppendRead,
		UpdateContext: resourceFirewallInternetServiceAppendUpdate,
		DeleteContext: resourceFirewallInternetServiceAppendDelete,

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
			"append_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Appending TCP/UDP/SCTP destination port (1 to 65535).",
				Optional:    true,
				Computed:    true,
			},
			"match_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Matching TCP/UDP/SCTP destination port (0 to 65535, 0 means any port).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallInternetServiceAppendCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallInternetServiceAppend(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallInternetServiceAppend(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallInternetServiceAppend")
	}

	return resourceFirewallInternetServiceAppendRead(ctx, d, meta)
}

func resourceFirewallInternetServiceAppendUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallInternetServiceAppend(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallInternetServiceAppend(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallInternetServiceAppend resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallInternetServiceAppend")
	}

	return resourceFirewallInternetServiceAppendRead(ctx, d, meta)
}

func resourceFirewallInternetServiceAppendDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectFirewallInternetServiceAppend(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateFirewallInternetServiceAppend(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallInternetServiceAppend resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceAppendRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallInternetServiceAppend(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallInternetServiceAppend resource: %v", err)
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

	diags := refreshObjectFirewallInternetServiceAppend(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallInternetServiceAppend(d *schema.ResourceData, o *models.FirewallInternetServiceAppend, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AppendPort != nil {
		v := *o.AppendPort

		if err = d.Set("append_port", v); err != nil {
			return diag.Errorf("error reading append_port: %v", err)
		}
	}

	if o.MatchPort != nil {
		v := *o.MatchPort

		if err = d.Set("match_port", v); err != nil {
			return diag.Errorf("error reading match_port: %v", err)
		}
	}

	return nil
}

func getObjectFirewallInternetServiceAppend(d *schema.ResourceData, sv string) (*models.FirewallInternetServiceAppend, diag.Diagnostics) {
	obj := models.FirewallInternetServiceAppend{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("append_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("append_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AppendPort = &tmp
		}
	}
	if v1, ok := d.GetOk("match_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MatchPort = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectFirewallInternetServiceAppend(d *schema.ResourceData, sv string) (*models.FirewallInternetServiceAppend, diag.Diagnostics) {
	obj := models.FirewallInternetServiceAppend{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
