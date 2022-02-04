// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: BGP aggregate address table.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceRouterBgpAggregateAddress() *schema.Resource {
	return &schema.Resource{
		Description: "BGP aggregate address table.",

		CreateContext: resourceRouterBgpAggregateAddressCreate,
		ReadContext:   resourceRouterBgpAggregateAddressRead,
		UpdateContext: resourceRouterBgpAggregateAddressUpdate,
		DeleteContext: resourceRouterBgpAggregateAddressDelete,

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
				Description:  "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"as_set": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable generate AS set path information.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"prefix": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

				Description: "Aggregate prefix.",
				Optional:    true,
				Computed:    true,
			},
			"summary_only": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable filter more specific routes from updates.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterBgpAggregateAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterBgpAggregateAddress resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterBgpAggregateAddress(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterBgpAggregateAddress(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpAggregateAddress")
	}

	return resourceRouterBgpAggregateAddressRead(ctx, d, meta)
}

func resourceRouterBgpAggregateAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterBgpAggregateAddress(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterBgpAggregateAddress(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterBgpAggregateAddress resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpAggregateAddress")
	}

	return resourceRouterBgpAggregateAddressRead(ctx, d, meta)
}

func resourceRouterBgpAggregateAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterBgpAggregateAddress(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterBgpAggregateAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpAggregateAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterBgpAggregateAddress(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBgpAggregateAddress resource: %v", err)
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

	diags := refreshObjectRouterBgpAggregateAddress(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterBgpAggregateAddress(d *schema.ResourceData, o *models.RouterBgpAggregateAddress, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AsSet != nil {
		v := *o.AsSet

		if err = d.Set("as_set", v); err != nil {
			return diag.Errorf("error reading as_set: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Prefix != nil {
		v := *o.Prefix
		if current, ok := d.GetOk("prefix"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("prefix", v); err != nil {
			return diag.Errorf("error reading prefix: %v", err)
		}
	}

	if o.SummaryOnly != nil {
		v := *o.SummaryOnly

		if err = d.Set("summary_only", v); err != nil {
			return diag.Errorf("error reading summary_only: %v", err)
		}
	}

	return nil
}

func getObjectRouterBgpAggregateAddress(d *schema.ResourceData, sv string) (*models.RouterBgpAggregateAddress, diag.Diagnostics) {
	obj := models.RouterBgpAggregateAddress{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("as_set"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("as_set", sv)
				diags = append(diags, e)
			}
			obj.AsSet = &v2
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
	if v1, ok := d.GetOk("prefix"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("prefix", sv)
				diags = append(diags, e)
			}
			obj.Prefix = &v2
		}
	}
	if v1, ok := d.GetOk("summary_only"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("summary_only", sv)
				diags = append(diags, e)
			}
			obj.SummaryOnly = &v2
		}
	}
	return &obj, diags
}
