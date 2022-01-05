// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.6,v7.0.2 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure physical switches.

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

func resourceSystemPhysicalSwitch() *schema.Resource {
	return &schema.Resource{
		Description: "Configure physical switches.",

		CreateContext: resourceSystemPhysicalSwitchCreate,
		ReadContext:   resourceSystemPhysicalSwitchRead,
		UpdateContext: resourceSystemPhysicalSwitchUpdate,
		DeleteContext: resourceSystemPhysicalSwitchDelete,

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
			"age_enable": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable layer 2 age timer.",
				Optional:    true,
				Computed:    true,
			},
			"age_val": {
				Type: schema.TypeInt,

				Description: "Layer 2 table age timer value.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceSystemPhysicalSwitchCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemPhysicalSwitch resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemPhysicalSwitch(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemPhysicalSwitch(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemPhysicalSwitch")
	}

	return resourceSystemPhysicalSwitchRead(ctx, d, meta)
}

func resourceSystemPhysicalSwitchUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemPhysicalSwitch(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemPhysicalSwitch(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemPhysicalSwitch resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemPhysicalSwitch")
	}

	return resourceSystemPhysicalSwitchRead(ctx, d, meta)
}

func resourceSystemPhysicalSwitchDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemPhysicalSwitch(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemPhysicalSwitch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPhysicalSwitchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemPhysicalSwitch(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemPhysicalSwitch resource: %v", err)
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

	diags := refreshObjectSystemPhysicalSwitch(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemPhysicalSwitch(d *schema.ResourceData, o *models.SystemPhysicalSwitch, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AgeEnable != nil {
		v := *o.AgeEnable

		if err = d.Set("age_enable", v); err != nil {
			return diag.Errorf("error reading age_enable: %v", err)
		}
	}

	if o.AgeVal != nil {
		v := *o.AgeVal

		if err = d.Set("age_val", v); err != nil {
			return diag.Errorf("error reading age_val: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	return nil
}

func getObjectSystemPhysicalSwitch(d *schema.ResourceData, sv string) (*models.SystemPhysicalSwitch, diag.Diagnostics) {
	obj := models.SystemPhysicalSwitch{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("age_enable"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("age_enable", sv)
				diags = append(diags, e)
			}
			obj.AgeEnable = &v2
		}
	}
	if v1, ok := d.GetOk("age_val"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("age_val", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AgeVal = &tmp
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
	return &obj, diags
}
