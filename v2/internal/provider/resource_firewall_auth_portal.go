// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure firewall authentication portals.

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

func resourceFirewallAuthPortal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure firewall authentication portals.",

		CreateContext: resourceFirewallAuthPortalCreate,
		ReadContext:   resourceFirewallAuthPortalRead,
		UpdateContext: resourceFirewallAuthPortalUpdate,
		DeleteContext: resourceFirewallAuthPortalDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "Firewall user groups permitted to authenticate through this portal. Separate group names with spaces.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"identity_based_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of the identity-based route that applies to this portal.",
				Optional:    true,
				Computed:    true,
			},
			"portal_addr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Address (or FQDN) of the authentication portal.",
				Optional:    true,
				Computed:    true,
			},
			"portal_addr6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "IPv6 address (or FQDN) of authentication portal.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallAuthPortalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallAuthPortal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallAuthPortal(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAuthPortal")
	}

	return resourceFirewallAuthPortalRead(ctx, d, meta)
}

func resourceFirewallAuthPortalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallAuthPortal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallAuthPortal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallAuthPortal resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAuthPortal")
	}

	return resourceFirewallAuthPortalRead(ctx, d, meta)
}

func resourceFirewallAuthPortalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectFirewallAuthPortal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateFirewallAuthPortal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallAuthPortal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAuthPortalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallAuthPortal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAuthPortal resource: %v", err)
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

	diags := refreshObjectFirewallAuthPortal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallAuthPortalGroups(d *schema.ResourceData, v *[]models.FirewallAuthPortalGroups, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectFirewallAuthPortal(d *schema.ResourceData, o *models.FirewallAuthPortal, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Groups != nil {
		if err = d.Set("groups", flattenFirewallAuthPortalGroups(d, o.Groups, "groups", sort)); err != nil {
			return diag.Errorf("error reading groups: %v", err)
		}
	}

	if o.IdentityBasedRoute != nil {
		v := *o.IdentityBasedRoute

		if err = d.Set("identity_based_route", v); err != nil {
			return diag.Errorf("error reading identity_based_route: %v", err)
		}
	}

	if o.PortalAddr != nil {
		v := *o.PortalAddr

		if err = d.Set("portal_addr", v); err != nil {
			return diag.Errorf("error reading portal_addr: %v", err)
		}
	}

	if o.PortalAddr6 != nil {
		v := *o.PortalAddr6

		if err = d.Set("portal_addr6", v); err != nil {
			return diag.Errorf("error reading portal_addr6: %v", err)
		}
	}

	return nil
}

func expandFirewallAuthPortalGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAuthPortalGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAuthPortalGroups

	for i := range l {
		tmp := models.FirewallAuthPortalGroups{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectFirewallAuthPortal(d *schema.ResourceData, sv string) (*models.FirewallAuthPortal, diag.Diagnostics) {
	obj := models.FirewallAuthPortal{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("groups", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAuthPortalGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Groups = t
		}
	} else if d.HasChange("groups") {
		old, new := d.GetChange("groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Groups = &[]models.FirewallAuthPortalGroups{}
		}
	}
	if v1, ok := d.GetOk("identity_based_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("identity_based_route", sv)
				diags = append(diags, e)
			}
			obj.IdentityBasedRoute = &v2
		}
	}
	if v1, ok := d.GetOk("portal_addr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("portal_addr", sv)
				diags = append(diags, e)
			}
			obj.PortalAddr = &v2
		}
	}
	if v1, ok := d.GetOk("portal_addr6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("portal_addr6", sv)
				diags = append(diags, e)
			}
			obj.PortalAddr6 = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectFirewallAuthPortal(d *schema.ResourceData, sv string) (*models.FirewallAuthPortal, diag.Diagnostics) {
	obj := models.FirewallAuthPortal{}
	diags := diag.Diagnostics{}

	obj.Groups = &[]models.FirewallAuthPortalGroups{}

	return &obj, diags
}
