// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: configure ips view-map

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

func resourceIpsViewMap() *schema.Resource {
	return &schema.Resource{
		Description: "configure ips view-map",

		CreateContext: resourceIpsViewMapCreate,
		ReadContext:   resourceIpsViewMapRead,
		UpdateContext: resourceIpsViewMapUpdate,
		DeleteContext: resourceIpsViewMapDelete,

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

				Description: "View ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"id_policy_id": {
				Type: schema.TypeInt,

				Description: "ID-based policy ID.",
				Optional:    true,
				Computed:    true,
			},
			"policy_id": {
				Type: schema.TypeInt,

				Description: "Policy ID.",
				Optional:    true,
				Computed:    true,
			},
			"vdom_id": {
				Type: schema.TypeInt,

				Description: "VDOM ID.",
				Optional:    true,
				Computed:    true,
			},
			"which": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"firewall", "interface", "interface6", "sniffer", "sniffer6", "explicit"}, false),

				Description: "Policy.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceIpsViewMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating IpsViewMap resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectIpsViewMap(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateIpsViewMap(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IpsViewMap")
	}

	return resourceIpsViewMapRead(ctx, d, meta)
}

func resourceIpsViewMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectIpsViewMap(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateIpsViewMap(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating IpsViewMap resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IpsViewMap")
	}

	return resourceIpsViewMapRead(ctx, d, meta)
}

func resourceIpsViewMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteIpsViewMap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting IpsViewMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsViewMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadIpsViewMap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IpsViewMap resource: %v", err)
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

	diags := refreshObjectIpsViewMap(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectIpsViewMap(d *schema.ResourceData, o *models.IpsViewMap, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.IdPolicyId != nil {
		v := *o.IdPolicyId

		if err = d.Set("id_policy_id", v); err != nil {
			return diag.Errorf("error reading id_policy_id: %v", err)
		}
	}

	if o.PolicyId != nil {
		v := *o.PolicyId

		if err = d.Set("policy_id", v); err != nil {
			return diag.Errorf("error reading policy_id: %v", err)
		}
	}

	if o.VdomId != nil {
		v := *o.VdomId

		if err = d.Set("vdom_id", v); err != nil {
			return diag.Errorf("error reading vdom_id: %v", err)
		}
	}

	if o.Which != nil {
		v := *o.Which

		if err = d.Set("which", v); err != nil {
			return diag.Errorf("error reading which: %v", err)
		}
	}

	return nil
}

func getObjectIpsViewMap(d *schema.ResourceData, sv string) (*models.IpsViewMap, diag.Diagnostics) {
	obj := models.IpsViewMap{}
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
	if v1, ok := d.GetOk("id_policy_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("id_policy_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IdPolicyId = &tmp
		}
	}
	if v1, ok := d.GetOk("policy_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("policy_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PolicyId = &tmp
		}
	}
	if v1, ok := d.GetOk("vdom_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vdom_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.VdomId = &tmp
		}
	}
	if v1, ok := d.GetOk("which"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("which", sv)
				diags = append(diags, e)
			}
			obj.Which = &v2
		}
	}
	return &obj, diags
}
