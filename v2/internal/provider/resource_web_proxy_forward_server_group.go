// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.

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

func resourceWebProxyForwardServerGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.",

		CreateContext: resourceWebProxyForwardServerGroupCreate,
		ReadContext:   resourceWebProxyForwardServerGroupRead,
		UpdateContext: resourceWebProxyForwardServerGroupUpdate,
		DeleteContext: resourceWebProxyForwardServerGroupDelete,

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
			"affinity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global).",
				Optional:    true,
				Computed:    true,
			},
			"group_down_option": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"block", "pass"}, false),

				Description: "Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination.",
				Optional:    true,
				Computed:    true,
			},
			"ldb_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"weighted", "least-session", "active-passive"}, false),

				Description: "Load balance method: weighted or least-session.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.",
				ForceNew:    true,
				Required:    true,
			},
			"server_list": {
				Type:        schema.TypeList,
				Description: "Add web forward servers to a list to form a server group. Optionally assign weights to each server.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Forward server name.",
							Optional:    true,
							Computed:    true,
						},
						"weight": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceWebProxyForwardServerGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WebProxyForwardServerGroup resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWebProxyForwardServerGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWebProxyForwardServerGroup(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebProxyForwardServerGroup")
	}

	return resourceWebProxyForwardServerGroupRead(ctx, d, meta)
}

func resourceWebProxyForwardServerGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebProxyForwardServerGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWebProxyForwardServerGroup(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WebProxyForwardServerGroup resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebProxyForwardServerGroup")
	}

	return resourceWebProxyForwardServerGroupRead(ctx, d, meta)
}

func resourceWebProxyForwardServerGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWebProxyForwardServerGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WebProxyForwardServerGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyForwardServerGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebProxyForwardServerGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebProxyForwardServerGroup resource: %v", err)
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

	diags := refreshObjectWebProxyForwardServerGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWebProxyForwardServerGroupServerList(v *[]models.WebProxyForwardServerGroupServerList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Weight; tmp != nil {
				v["weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectWebProxyForwardServerGroup(d *schema.ResourceData, o *models.WebProxyForwardServerGroup, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Affinity != nil {
		v := *o.Affinity

		if err = d.Set("affinity", v); err != nil {
			return diag.Errorf("error reading affinity: %v", err)
		}
	}

	if o.GroupDownOption != nil {
		v := *o.GroupDownOption

		if err = d.Set("group_down_option", v); err != nil {
			return diag.Errorf("error reading group_down_option: %v", err)
		}
	}

	if o.LdbMethod != nil {
		v := *o.LdbMethod

		if err = d.Set("ldb_method", v); err != nil {
			return diag.Errorf("error reading ldb_method: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ServerList != nil {
		if err = d.Set("server_list", flattenWebProxyForwardServerGroupServerList(o.ServerList, sort)); err != nil {
			return diag.Errorf("error reading server_list: %v", err)
		}
	}

	return nil
}

func expandWebProxyForwardServerGroupServerList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebProxyForwardServerGroupServerList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebProxyForwardServerGroupServerList

	for i := range l {
		tmp := models.WebProxyForwardServerGroupServerList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Weight = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWebProxyForwardServerGroup(d *schema.ResourceData, sv string) (*models.WebProxyForwardServerGroup, diag.Diagnostics) {
	obj := models.WebProxyForwardServerGroup{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("affinity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("affinity", sv)
				diags = append(diags, e)
			}
			obj.Affinity = &v2
		}
	}
	if v1, ok := d.GetOk("group_down_option"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_down_option", sv)
				diags = append(diags, e)
			}
			obj.GroupDownOption = &v2
		}
	}
	if v1, ok := d.GetOk("ldb_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldb_method", sv)
				diags = append(diags, e)
			}
			obj.LdbMethod = &v2
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
	if v, ok := d.GetOk("server_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("server_list", sv)
			diags = append(diags, e)
		}
		t, err := expandWebProxyForwardServerGroupServerList(d, v, "server_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ServerList = t
		}
	} else if d.HasChange("server_list") {
		old, new := d.GetChange("server_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ServerList = &[]models.WebProxyForwardServerGroupServerList{}
		}
	}
	return &obj, diags
}
