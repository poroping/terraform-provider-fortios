// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure an aggregate of IPsec tunnels.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemIpsecAggregate() *schema.Resource {
	return &schema.Resource{
		Description: "Configure an aggregate of IPsec tunnels.",

		CreateContext: resourceSystemIpsecAggregateCreate,
		ReadContext:   resourceSystemIpsecAggregateRead,
		UpdateContext: resourceSystemIpsecAggregateUpdate,
		DeleteContext: resourceSystemIpsecAggregateDelete,

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
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"L3", "L4", "round-robin", "redundant", "weighted-round-robin"}, false),

				Description: "Frame distribution algorithm.",
				Optional:    true,
				Computed:    true,
			},
			"member": {
				Type:        schema.TypeList,
				Description: "Member tunnels of the aggregate.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tunnel_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Tunnel name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "IPsec aggregate name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceSystemIpsecAggregateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemIpsecAggregate resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemIpsecAggregate(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemIpsecAggregate(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIpsecAggregate")
	}

	return resourceSystemIpsecAggregateRead(ctx, d, meta)
}

func resourceSystemIpsecAggregateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemIpsecAggregate(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemIpsecAggregate(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemIpsecAggregate resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIpsecAggregate")
	}

	return resourceSystemIpsecAggregateRead(ctx, d, meta)
}

func resourceSystemIpsecAggregateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemIpsecAggregate(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemIpsecAggregate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpsecAggregateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemIpsecAggregate(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemIpsecAggregate resource: %v", err)
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

	diags := refreshObjectSystemIpsecAggregate(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemIpsecAggregateMember(v *[]models.SystemIpsecAggregateMember, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.TunnelName; tmp != nil {
				v["tunnel_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "tunnel_name")
	}

	return flat
}

func refreshObjectSystemIpsecAggregate(d *schema.ResourceData, o *models.SystemIpsecAggregate, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Algorithm != nil {
		v := *o.Algorithm

		if err = d.Set("algorithm", v); err != nil {
			return diag.Errorf("error reading algorithm: %v", err)
		}
	}

	if o.Member != nil {
		if err = d.Set("member", flattenSystemIpsecAggregateMember(o.Member, sort)); err != nil {
			return diag.Errorf("error reading member: %v", err)
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

func expandSystemIpsecAggregateMember(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemIpsecAggregateMember, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIpsecAggregateMember

	for i := range l {
		tmp := models.SystemIpsecAggregateMember{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.tunnel_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TunnelName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemIpsecAggregate(d *schema.ResourceData, sv string) (*models.SystemIpsecAggregate, diag.Diagnostics) {
	obj := models.SystemIpsecAggregate{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("algorithm", sv)
				diags = append(diags, e)
			}
			obj.Algorithm = &v2
		}
	}
	if v, ok := d.GetOk("member"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("member", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIpsecAggregateMember(d, v, "member", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Member = t
		}
	} else if d.HasChange("member") {
		old, new := d.GetChange("member")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Member = &[]models.SystemIpsecAggregateMember{}
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
