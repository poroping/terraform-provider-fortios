// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Administrative distance modifications.

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

func resourceRouterBgpAdminDistance() *schema.Resource {
	return &schema.Resource{
		Description: "Administrative distance modifications.",

		CreateContext: resourceRouterBgpAdminDistanceCreate,
		ReadContext:   resourceRouterBgpAdminDistanceRead,
		UpdateContext: resourceRouterBgpAdminDistanceUpdate,
		DeleteContext: resourceRouterBgpAdminDistanceDelete,

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
			"distance": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Administrative distance to apply (1 - 255).",
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
			"neighbour_prefix": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Neighbor address prefix.",
				Optional:    true,
				Computed:    true,
			},
			"route_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Access list of routes to apply new distance to.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterBgpAdminDistanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterBgpAdminDistance resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterBgpAdminDistance(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterBgpAdminDistance(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpAdminDistance")
	}

	return resourceRouterBgpAdminDistanceRead(ctx, d, meta)
}

func resourceRouterBgpAdminDistanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterBgpAdminDistance(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterBgpAdminDistance(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterBgpAdminDistance resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpAdminDistance")
	}

	return resourceRouterBgpAdminDistanceRead(ctx, d, meta)
}

func resourceRouterBgpAdminDistanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterBgpAdminDistance(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterBgpAdminDistance resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpAdminDistanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterBgpAdminDistance(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBgpAdminDistance resource: %v", err)
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

	diags := refreshObjectRouterBgpAdminDistance(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterBgpAdminDistance(d *schema.ResourceData, o *models.RouterBgpAdminDistance, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Distance != nil {
		v := *o.Distance

		if err = d.Set("distance", v); err != nil {
			return diag.Errorf("error reading distance: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.NeighbourPrefix != nil {
		v := *o.NeighbourPrefix
		if current, ok := d.GetOk("neighbour_prefix"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("neighbour_prefix", v); err != nil {
			return diag.Errorf("error reading neighbour_prefix: %v", err)
		}
	}

	if o.RouteList != nil {
		v := *o.RouteList

		if err = d.Set("route_list", v); err != nil {
			return diag.Errorf("error reading route_list: %v", err)
		}
	}

	return nil
}

func getObjectRouterBgpAdminDistance(d *schema.ResourceData, sv string) (*models.RouterBgpAdminDistance, diag.Diagnostics) {
	obj := models.RouterBgpAdminDistance{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("distance"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distance", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Distance = &tmp
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
	if v1, ok := d.GetOk("neighbour_prefix"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("neighbour_prefix", sv)
				diags = append(diags, e)
			}
			obj.NeighbourPrefix = &v2
		}
	}
	if v1, ok := d.GetOk("route_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_list", sv)
				diags = append(diags, e)
			}
			obj.RouteList = &v2
		}
	}
	return &obj, diags
}
