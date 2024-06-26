// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.

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

func resourceSystemSdwanNeighbor() *schema.Resource {
	return &schema.Resource{
		Description: "Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.",

		CreateContext: resourceSystemSdwanNeighborCreate,
		ReadContext:   resourceSystemSdwanNeighborRead,
		UpdateContext: resourceSystemSdwanNeighborUpdate,
		DeleteContext: resourceSystemSdwanNeighborDelete,

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
			"health_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SD-WAN health-check name.",
				Optional:    true,
				Computed:    true,
			},
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 45),

				Description: "IP/IPv6 address of neighbor.",
				ForceNew:    true,
				Required:    true,
			},
			"member": {
				Type:        schema.TypeList,
				Description: "Member sequence number list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type: schema.TypeInt,

							Description: "Member sequence number.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"minimum_sla_meet_members": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Minimum number of members which meet SLA when the neighbor is preferred.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"sla", "speedtest"}, false),

				Description: "What metric to select the neighbor.",
				Optional:    true,
				Computed:    true,
			},
			"role": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"standalone", "primary", "secondary"}, false),

				Description: "Role of neighbor.",
				Optional:    true,
				Computed:    true,
			},
			"sla_id": {
				Type: schema.TypeInt,

				Description: "SLA ID.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemSdwanNeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "ip"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemSdwanNeighbor resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemSdwanNeighbor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSdwanNeighbor(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSdwanNeighbor")
	}

	return resourceSystemSdwanNeighborRead(ctx, d, meta)
}

func resourceSystemSdwanNeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSdwanNeighbor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSdwanNeighbor(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSdwanNeighbor resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSdwanNeighbor")
	}

	return resourceSystemSdwanNeighborRead(ctx, d, meta)
}

func resourceSystemSdwanNeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSdwanNeighbor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSdwanNeighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSdwanNeighbor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSdwanNeighbor resource: %v", err)
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

	diags := refreshObjectSystemSdwanNeighbor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemSdwanNeighbor(d *schema.ResourceData, o *models.SystemSdwanNeighbor, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.HealthCheck != nil {
		v := *o.HealthCheck

		if err = d.Set("health_check", v); err != nil {
			return diag.Errorf("error reading health_check: %v", err)
		}
	}

	if o.Ip != nil {
		v := *o.Ip

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.Member != nil {
		if err = d.Set("member", flattenSystemSdwanNeighborMember(d, o.Member, "member", sort)); err != nil {
			return diag.Errorf("error reading member: %v", err)
		}
	}

	if o.MinimumSlaMeetMembers != nil {
		v := *o.MinimumSlaMeetMembers

		if err = d.Set("minimum_sla_meet_members", v); err != nil {
			return diag.Errorf("error reading minimum_sla_meet_members: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.Role != nil {
		v := *o.Role

		if err = d.Set("role", v); err != nil {
			return diag.Errorf("error reading role: %v", err)
		}
	}

	if o.SlaId != nil {
		v := *o.SlaId

		if err = d.Set("sla_id", v); err != nil {
			return diag.Errorf("error reading sla_id: %v", err)
		}
	}

	return nil
}

func getObjectSystemSdwanNeighbor(d *schema.ResourceData, sv string) (*models.SystemSdwanNeighbor, diag.Diagnostics) {
	obj := models.SystemSdwanNeighbor{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("health_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("health_check", sv)
				diags = append(diags, e)
			}
			obj.HealthCheck = &v2
		}
	}
	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
		}
	}
	if v, ok := d.GetOk("member"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("member", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdwanNeighborMember(d, v, "member", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Member = t
		}
	} else if d.HasChange("member") {
		old, new := d.GetChange("member")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Member = &[]models.SystemSdwanNeighborMember{}
		}
	}
	if v1, ok := d.GetOk("minimum_sla_meet_members"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("minimum_sla_meet_members", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MinimumSlaMeetMembers = &tmp
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
		}
	}
	if v1, ok := d.GetOk("role"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("role", sv)
				diags = append(diags, e)
			}
			obj.Role = &v2
		}
	}
	if v1, ok := d.GetOk("sla_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sla_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SlaId = &tmp
		}
	}
	return &obj, diags
}
