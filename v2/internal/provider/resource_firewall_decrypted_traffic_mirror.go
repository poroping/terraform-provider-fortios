// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure decrypted traffic mirror.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceFirewallDecryptedTrafficMirror() *schema.Resource {
	return &schema.Resource{
		Description: "Configure decrypted traffic mirror.",

		CreateContext: resourceFirewallDecryptedTrafficMirrorCreate,
		ReadContext:   resourceFirewallDecryptedTrafficMirrorRead,
		UpdateContext: resourceFirewallDecryptedTrafficMirrorUpdate,
		DeleteContext: resourceFirewallDecryptedTrafficMirrorDelete,

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
			"dstmac": {
				Type: schema.TypeString,

				Description: "Set destination MAC address for mirrored traffic.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeList,
				Description: "Decrypted traffic mirror interface.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Decrypted traffic mirror interface.",
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
			"traffic_source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"client", "server", "both"}, false),

				Description: "Source of decrypted traffic to be mirrored.",
				Optional:    true,
				Computed:    true,
			},
			"traffic_type": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Types of decrypted traffic to be mirrored.",
				Optional:         true,
				Computed:         true,
			},
		},
	}
}

func resourceFirewallDecryptedTrafficMirrorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallDecryptedTrafficMirror resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallDecryptedTrafficMirror(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallDecryptedTrafficMirror(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallDecryptedTrafficMirror")
	}

	return resourceFirewallDecryptedTrafficMirrorRead(ctx, d, meta)
}

func resourceFirewallDecryptedTrafficMirrorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallDecryptedTrafficMirror(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallDecryptedTrafficMirror(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallDecryptedTrafficMirror resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallDecryptedTrafficMirror")
	}

	return resourceFirewallDecryptedTrafficMirrorRead(ctx, d, meta)
}

func resourceFirewallDecryptedTrafficMirrorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallDecryptedTrafficMirror(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallDecryptedTrafficMirror resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallDecryptedTrafficMirrorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallDecryptedTrafficMirror(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallDecryptedTrafficMirror resource: %v", err)
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

	diags := refreshObjectFirewallDecryptedTrafficMirror(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallDecryptedTrafficMirrorInterface(v *[]models.FirewallDecryptedTrafficMirrorInterface, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func refreshObjectFirewallDecryptedTrafficMirror(d *schema.ResourceData, o *models.FirewallDecryptedTrafficMirror, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Dstmac != nil {
		v := *o.Dstmac

		if err = d.Set("dstmac", v); err != nil {
			return diag.Errorf("error reading dstmac: %v", err)
		}
	}

	if o.Interface != nil {
		if err = d.Set("interface", flattenFirewallDecryptedTrafficMirrorInterface(o.Interface, sort)); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.TrafficSource != nil {
		v := *o.TrafficSource

		if err = d.Set("traffic_source", v); err != nil {
			return diag.Errorf("error reading traffic_source: %v", err)
		}
	}

	if o.TrafficType != nil {
		v := *o.TrafficType

		if err = d.Set("traffic_type", v); err != nil {
			return diag.Errorf("error reading traffic_type: %v", err)
		}
	}

	return nil
}

func expandFirewallDecryptedTrafficMirrorInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallDecryptedTrafficMirrorInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallDecryptedTrafficMirrorInterface

	for i := range l {
		tmp := models.FirewallDecryptedTrafficMirrorInterface{}
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

func getObjectFirewallDecryptedTrafficMirror(d *schema.ResourceData, sv string) (*models.FirewallDecryptedTrafficMirror, diag.Diagnostics) {
	obj := models.FirewallDecryptedTrafficMirror{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("dstmac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dstmac", sv)
				diags = append(diags, e)
			}
			obj.Dstmac = &v2
		}
	}
	if v, ok := d.GetOk("interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("interface", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallDecryptedTrafficMirrorInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Interface = t
		}
	} else if d.HasChange("interface") {
		old, new := d.GetChange("interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Interface = &[]models.FirewallDecryptedTrafficMirrorInterface{}
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
	if v1, ok := d.GetOk("traffic_source"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_source", sv)
				diags = append(diags, e)
			}
			obj.TrafficSource = &v2
		}
	}
	if v1, ok := d.GetOk("traffic_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_type", sv)
				diags = append(diags, e)
			}
			obj.TrafficType = &v2
		}
	}
	return &obj, diags
}
