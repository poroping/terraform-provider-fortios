// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Show Internet Service application.

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

func resourceFirewallInternetService() *schema.Resource {
	return &schema.Resource{
		Description: "Show Internet Service application.",

		CreateContext: resourceFirewallInternetServiceCreate,
		ReadContext:   resourceFirewallInternetServiceRead,
		UpdateContext: resourceFirewallInternetServiceUpdate,
		DeleteContext: resourceFirewallInternetServiceDelete,

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
				Type:         schema.TypeBool,
				Description:  "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"database": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"isdb", "irdb"}, false),

				Description: "Database name this Internet Service belongs to.",
				Optional:    true,
				Computed:    true,
			},
			"direction": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"src", "dst", "both"}, false),

				Description: "How this service may be used in a firewall policy (source, destination or both).",
				Optional:    true,
				Computed:    true,
			},
			"extra_ip_range_number": {
				Type: schema.TypeInt,

				Description: "Extra number of IP ranges.",
				Optional:    true,
				Computed:    true,
			},
			"icon_id": {
				Type: schema.TypeInt,

				Description: "Icon ID of Internet Service.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Internet Service ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"ip_number": {
				Type: schema.TypeInt,

				Description: "Total number of IP addresses.",
				Optional:    true,
				Computed:    true,
			},
			"ip_range_number": {
				Type: schema.TypeInt,

				Description: "Number of IP ranges.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Internet Service name.",
				Optional:    true,
				Computed:    true,
			},
			"obsolete": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Indicates whether the Internet Service can be used.",
				Optional:    true,
				Computed:    true,
			},
			"reputation": {
				Type: schema.TypeInt,

				Description: "Reputation level of the Internet Service.",
				Optional:    true,
				Computed:    true,
			},
			"singularity": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Singular level of the Internet Service.",
				Optional:    true,
				Computed:    true,
			},
			"sld_id": {
				Type: schema.TypeInt,

				Description: "Second Level Domain.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallInternetServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallInternetService resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallInternetService(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallInternetService(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallInternetService")
	}

	return resourceFirewallInternetServiceRead(ctx, d, meta)
}

func resourceFirewallInternetServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallInternetService(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallInternetService(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallInternetService resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallInternetService")
	}

	return resourceFirewallInternetServiceRead(ctx, d, meta)
}

func resourceFirewallInternetServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallInternetService(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallInternetService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallInternetService(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallInternetService resource: %v", err)
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

	diags := refreshObjectFirewallInternetService(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallInternetService(d *schema.ResourceData, o *models.FirewallInternetService, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Database != nil {
		v := *o.Database

		if err = d.Set("database", v); err != nil {
			return diag.Errorf("error reading database: %v", err)
		}
	}

	if o.Direction != nil {
		v := *o.Direction

		if err = d.Set("direction", v); err != nil {
			return diag.Errorf("error reading direction: %v", err)
		}
	}

	if o.ExtraIpRangeNumber != nil {
		v := *o.ExtraIpRangeNumber

		if err = d.Set("extra_ip_range_number", v); err != nil {
			return diag.Errorf("error reading extra_ip_range_number: %v", err)
		}
	}

	if o.IconId != nil {
		v := *o.IconId

		if err = d.Set("icon_id", v); err != nil {
			return diag.Errorf("error reading icon_id: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.IpNumber != nil {
		v := *o.IpNumber

		if err = d.Set("ip_number", v); err != nil {
			return diag.Errorf("error reading ip_number: %v", err)
		}
	}

	if o.IpRangeNumber != nil {
		v := *o.IpRangeNumber

		if err = d.Set("ip_range_number", v); err != nil {
			return diag.Errorf("error reading ip_range_number: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Obsolete != nil {
		v := *o.Obsolete

		if err = d.Set("obsolete", v); err != nil {
			return diag.Errorf("error reading obsolete: %v", err)
		}
	}

	if o.Reputation != nil {
		v := *o.Reputation

		if err = d.Set("reputation", v); err != nil {
			return diag.Errorf("error reading reputation: %v", err)
		}
	}

	if o.Singularity != nil {
		v := *o.Singularity

		if err = d.Set("singularity", v); err != nil {
			return diag.Errorf("error reading singularity: %v", err)
		}
	}

	if o.SldId != nil {
		v := *o.SldId

		if err = d.Set("sld_id", v); err != nil {
			return diag.Errorf("error reading sld_id: %v", err)
		}
	}

	return nil
}

func getObjectFirewallInternetService(d *schema.ResourceData, sv string) (*models.FirewallInternetService, diag.Diagnostics) {
	obj := models.FirewallInternetService{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("database"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("database", sv)
				diags = append(diags, e)
			}
			obj.Database = &v2
		}
	}
	if v1, ok := d.GetOk("direction"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("direction", sv)
				diags = append(diags, e)
			}
			obj.Direction = &v2
		}
	}
	if v1, ok := d.GetOk("extra_ip_range_number"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extra_ip_range_number", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ExtraIpRangeNumber = &tmp
		}
	}
	if v1, ok := d.GetOk("icon_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("icon_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IconId = &tmp
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	if v1, ok := d.GetOk("ip_number"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_number", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IpNumber = &tmp
		}
	}
	if v1, ok := d.GetOk("ip_range_number"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_range_number", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IpRangeNumber = &tmp
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
	if v1, ok := d.GetOk("obsolete"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("obsolete", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Obsolete = &tmp
		}
	}
	if v1, ok := d.GetOk("reputation"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("reputation", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Reputation = &tmp
		}
	}
	if v1, ok := d.GetOk("singularity"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("singularity", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Singularity = &tmp
		}
	}
	if v1, ok := d.GetOk("sld_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("sld_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SldId = &tmp
		}
	}
	return &obj, diags
}
