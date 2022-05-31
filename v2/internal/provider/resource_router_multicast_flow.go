// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure multicast-flow.

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

func resourceRouterMulticastFlow() *schema.Resource {
	return &schema.Resource{
		Description: "Configure multicast-flow.",

		CreateContext: resourceRouterMulticastFlowCreate,
		ReadContext:   resourceRouterMulticastFlowRead,
		UpdateContext: resourceRouterMulticastFlowUpdate,
		DeleteContext: resourceRouterMulticastFlowDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"flows": {
				Type:        schema.TypeList,
				Description: "Multicast-flow entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_addr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Multicast group IP address.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Flow ID.",
							Optional:    true,
							Computed:    true,
						},
						"source_addr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Multicast source IP address.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceRouterMulticastFlowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterMulticastFlow resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterMulticastFlow(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterMulticastFlow(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterMulticastFlow")
	}

	return resourceRouterMulticastFlowRead(ctx, d, meta)
}

func resourceRouterMulticastFlowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterMulticastFlow(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterMulticastFlow(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterMulticastFlow resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterMulticastFlow")
	}

	return resourceRouterMulticastFlowRead(ctx, d, meta)
}

func resourceRouterMulticastFlowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterMulticastFlow(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterMulticastFlow resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastFlowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterMulticastFlow(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterMulticastFlow resource: %v", err)
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

	diags := refreshObjectRouterMulticastFlow(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterMulticastFlowFlows(d *schema.ResourceData, v *[]models.RouterMulticastFlowFlows, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.GroupAddr; tmp != nil {
				v["group_addr"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.SourceAddr; tmp != nil {
				v["source_addr"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectRouterMulticastFlow(d *schema.ResourceData, o *models.RouterMulticastFlow, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Flows != nil {
		if err = d.Set("flows", flattenRouterMulticastFlowFlows(d, o.Flows, "flows", sort)); err != nil {
			return diag.Errorf("error reading flows: %v", err)
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

func expandRouterMulticastFlowFlows(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterMulticastFlowFlows, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterMulticastFlowFlows

	for i := range l {
		tmp := models.RouterMulticastFlowFlows{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.group_addr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.GroupAddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source_addr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SourceAddr = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectRouterMulticastFlow(d *schema.ResourceData, sv string) (*models.RouterMulticastFlow, diag.Diagnostics) {
	obj := models.RouterMulticastFlow{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v, ok := d.GetOk("flows"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("flows", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterMulticastFlowFlows(d, v, "flows", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Flows = t
		}
	} else if d.HasChange("flows") {
		old, new := d.GetChange("flows")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Flows = &[]models.RouterMulticastFlowFlows{}
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
