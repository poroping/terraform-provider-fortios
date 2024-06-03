// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure names for shaping classes.

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

func resourceFirewallTrafficClass() *schema.Resource {
	return &schema.Resource{
		Description: "Configure names for shaping classes.",

		CreateContext: resourceFirewallTrafficClassCreate,
		ReadContext:   resourceFirewallTrafficClassRead,
		UpdateContext: resourceFirewallTrafficClassUpdate,
		DeleteContext: resourceFirewallTrafficClassDelete,

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
			"class_id": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 31),

				Description: "Class ID to be named.",
				ForceNew:    true,
				Required:    true,
			},
			"class_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Define the name for this class-id.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallTrafficClassCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "class_id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallTrafficClass resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallTrafficClass(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallTrafficClass(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallTrafficClass")
	}

	return resourceFirewallTrafficClassRead(ctx, d, meta)
}

func resourceFirewallTrafficClassUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallTrafficClass(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallTrafficClass(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallTrafficClass resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallTrafficClass")
	}

	return resourceFirewallTrafficClassRead(ctx, d, meta)
}

func resourceFirewallTrafficClassDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallTrafficClass(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallTrafficClass resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallTrafficClassRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallTrafficClass(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallTrafficClass resource: %v", err)
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

	diags := refreshObjectFirewallTrafficClass(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallTrafficClass(d *schema.ResourceData, o *models.FirewallTrafficClass, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ClassId != nil {
		v := *o.ClassId

		if err = d.Set("class_id", v); err != nil {
			return diag.Errorf("error reading class_id: %v", err)
		}
	}

	if o.ClassName != nil {
		v := *o.ClassName

		if err = d.Set("class_name", v); err != nil {
			return diag.Errorf("error reading class_name: %v", err)
		}
	}

	return nil
}

func getObjectFirewallTrafficClass(d *schema.ResourceData, sv string) (*models.FirewallTrafficClass, diag.Diagnostics) {
	obj := models.FirewallTrafficClass{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("class_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("class_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ClassId = &tmp
		}
	}
	if v1, ok := d.GetOk("class_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("class_name", sv)
				diags = append(diags, e)
			}
			obj.ClassName = &v2
		}
	}
	return &obj, diags
}
