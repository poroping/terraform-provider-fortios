// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: BGP IPv6 neighbor range table.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceRouterBgpNeighborRange6() *schema.Resource {
	return &schema.Resource{
		Description: "BGP IPv6 neighbor range table.",

		CreateContext: resourceRouterBgpNeighborRange6Create,
		ReadContext:   resourceRouterBgpNeighborRange6Read,
		UpdateContext: resourceRouterBgpNeighborRange6Update,
		DeleteContext: resourceRouterBgpNeighborRange6Delete,

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
			"fosid": {
				Type: schema.TypeInt,

				Description: "IPv6 neighbor range ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"max_neighbor_num": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1000),

				Description: "Maximum number of neighbors.",
				Optional:    true,
				Computed:    true,
			},
			"neighbor_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Neighbor group name.",
				Optional:    true,
				Computed:    true,
			},
			"prefix6": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Network,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "IPv6 prefix.",
				Optional:         true,
				Computed:         true,
			},
		},
	}
}

func resourceRouterBgpNeighborRange6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterBgpNeighborRange6 resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterBgpNeighborRange6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterBgpNeighborRange6(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpNeighborRange6")
	}

	return resourceRouterBgpNeighborRange6Read(ctx, d, meta)
}

func resourceRouterBgpNeighborRange6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterBgpNeighborRange6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterBgpNeighborRange6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterBgpNeighborRange6 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpNeighborRange6")
	}

	return resourceRouterBgpNeighborRange6Read(ctx, d, meta)
}

func resourceRouterBgpNeighborRange6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterBgpNeighborRange6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterBgpNeighborRange6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpNeighborRange6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterBgpNeighborRange6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBgpNeighborRange6 resource: %v", err)
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

	diags := refreshObjectRouterBgpNeighborRange6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterBgpNeighborRange6(d *schema.ResourceData, o *models.RouterBgpNeighborRange6, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.MaxNeighborNum != nil {
		v := *o.MaxNeighborNum

		if err = d.Set("max_neighbor_num", v); err != nil {
			return diag.Errorf("error reading max_neighbor_num: %v", err)
		}
	}

	if o.NeighborGroup != nil {
		v := *o.NeighborGroup

		if err = d.Set("neighbor_group", v); err != nil {
			return diag.Errorf("error reading neighbor_group: %v", err)
		}
	}

	if o.Prefix6 != nil {
		v := *o.Prefix6

		if err = d.Set("prefix6", v); err != nil {
			return diag.Errorf("error reading prefix6: %v", err)
		}
	}

	return nil
}

func getObjectRouterBgpNeighborRange6(d *schema.ResourceData, sv string) (*models.RouterBgpNeighborRange6, diag.Diagnostics) {
	obj := models.RouterBgpNeighborRange6{}
	diags := diag.Diagnostics{}

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
	if v1, ok := d.GetOk("max_neighbor_num"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_neighbor_num", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxNeighborNum = &tmp
		}
	}
	if v1, ok := d.GetOk("neighbor_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("neighbor_group", sv)
				diags = append(diags, e)
			}
			obj.NeighborGroup = &v2
		}
	}
	if v1, ok := d.GetOk("prefix6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("prefix6", sv)
				diags = append(diags, e)
			}
			obj.Prefix6 = &v2
		}
	}
	return &obj, diags
}
