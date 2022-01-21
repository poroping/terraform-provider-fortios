// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Define internet service names.

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

func resourceFirewallInternetServiceName() *schema.Resource {
	return &schema.Resource{
		Description: "Define internet service names.",

		CreateContext: resourceFirewallInternetServiceNameCreate,
		ReadContext:   resourceFirewallInternetServiceNameRead,
		UpdateContext: resourceFirewallInternetServiceNameUpdate,
		DeleteContext: resourceFirewallInternetServiceNameDelete,

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
			"city_id": {
				Type: schema.TypeInt,

				Description: "City ID.",
				Optional:    true,
				Computed:    true,
			},
			"country_id": {
				Type: schema.TypeInt,

				Description: "Country or Area ID.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_id": {
				Type: schema.TypeInt,

				Description: "Internet Service ID.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Internet Service name.",
				ForceNew:    true,
				Required:    true,
			},
			"region_id": {
				Type: schema.TypeInt,

				Description: "Region ID.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "location"}, false),

				Description: "Internet Service name type.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallInternetServiceNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallInternetServiceName resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallInternetServiceName(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallInternetServiceName(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallInternetServiceName")
	}

	return resourceFirewallInternetServiceNameRead(ctx, d, meta)
}

func resourceFirewallInternetServiceNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallInternetServiceName(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallInternetServiceName(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallInternetServiceName resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallInternetServiceName")
	}

	return resourceFirewallInternetServiceNameRead(ctx, d, meta)
}

func resourceFirewallInternetServiceNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallInternetServiceName(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallInternetServiceName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallInternetServiceName(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallInternetServiceName resource: %v", err)
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

	diags := refreshObjectFirewallInternetServiceName(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallInternetServiceName(d *schema.ResourceData, o *models.FirewallInternetServiceName, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CityId != nil {
		v := *o.CityId

		if err = d.Set("city_id", v); err != nil {
			return diag.Errorf("error reading city_id: %v", err)
		}
	}

	if o.CountryId != nil {
		v := *o.CountryId

		if err = d.Set("country_id", v); err != nil {
			return diag.Errorf("error reading country_id: %v", err)
		}
	}

	if o.InternetServiceId != nil {
		v := *o.InternetServiceId

		if err = d.Set("internet_service_id", v); err != nil {
			return diag.Errorf("error reading internet_service_id: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.RegionId != nil {
		v := *o.RegionId

		if err = d.Set("region_id", v); err != nil {
			return diag.Errorf("error reading region_id: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	return nil
}

func getObjectFirewallInternetServiceName(d *schema.ResourceData, sv string) (*models.FirewallInternetServiceName, diag.Diagnostics) {
	obj := models.FirewallInternetServiceName{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("city_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("city_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CityId = &tmp
		}
	}
	if v1, ok := d.GetOk("country_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("country_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CountryId = &tmp
		}
	}
	if v1, ok := d.GetOk("internet_service_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.InternetServiceId = &tmp
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
	if v1, ok := d.GetOk("region_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("region_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RegionId = &tmp
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	return &obj, diags
}
