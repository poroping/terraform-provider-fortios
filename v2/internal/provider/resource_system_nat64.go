// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure NAT64.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceSystemNat64() *schema.Resource {
	return &schema.Resource{
		Description: "Configure NAT64.",

		CreateContext: resourceSystemNat64Create,
		ReadContext:   resourceSystemNat64Read,
		UpdateContext: resourceSystemNat64Update,
		DeleteContext: resourceSystemNat64Delete,

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
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"always_synthesize_aaaa_record": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable AAAA record synthesis (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"generate_ipv6_fragment_header": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv6 fragment header generation.",
				Optional:    true,
				Computed:    true,
			},
			"nat46_force_ipv4_packet_forwarding": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable mandatory IPv4 packet forwarding in nat46.",
				Optional:    true,
				Computed:    true,
			},
			"nat64_prefix": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "NAT64 prefix must be ::/96 (default = 64:ff9b::/96).",
				Optional:         true,
				Computed:         true,
			},
			"secondary_prefix": {
				Type:        schema.TypeList,
				Description: "Secondary NAT64 prefix.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "NAT64 prefix name.",
							Optional:    true,
							Computed:    true,
						},
						"nat64_prefix": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Prefix,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "NAT64 prefix.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"secondary_prefix_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable secondary NAT64 prefix.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NAT64 (default = disable).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemNat64Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemNat64(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemNat64(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemNat64")
	}

	return resourceSystemNat64Read(ctx, d, meta)
}

func resourceSystemNat64Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemNat64(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemNat64(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemNat64 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemNat64")
	}

	return resourceSystemNat64Read(ctx, d, meta)
}

func resourceSystemNat64Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemNat64(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemNat64 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNat64Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemNat64(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemNat64 resource: %v", err)
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

	diags := refreshObjectSystemNat64(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemNat64SecondaryPrefix(v *[]models.SystemNat64SecondaryPrefix, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Nat64Prefix; tmp != nil {
				v["nat64_prefix"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectSystemNat64(d *schema.ResourceData, o *models.SystemNat64, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AlwaysSynthesizeAaaaRecord != nil {
		v := *o.AlwaysSynthesizeAaaaRecord

		if err = d.Set("always_synthesize_aaaa_record", v); err != nil {
			return diag.Errorf("error reading always_synthesize_aaaa_record: %v", err)
		}
	}

	if o.GenerateIpv6FragmentHeader != nil {
		v := *o.GenerateIpv6FragmentHeader

		if err = d.Set("generate_ipv6_fragment_header", v); err != nil {
			return diag.Errorf("error reading generate_ipv6_fragment_header: %v", err)
		}
	}

	if o.Nat46ForceIpv4PacketForwarding != nil {
		v := *o.Nat46ForceIpv4PacketForwarding

		if err = d.Set("nat46_force_ipv4_packet_forwarding", v); err != nil {
			return diag.Errorf("error reading nat46_force_ipv4_packet_forwarding: %v", err)
		}
	}

	if o.Nat64Prefix != nil {
		v := *o.Nat64Prefix

		if err = d.Set("nat64_prefix", v); err != nil {
			return diag.Errorf("error reading nat64_prefix: %v", err)
		}
	}

	if o.SecondaryPrefix != nil {
		if err = d.Set("secondary_prefix", flattenSystemNat64SecondaryPrefix(o.SecondaryPrefix, sort)); err != nil {
			return diag.Errorf("error reading secondary_prefix: %v", err)
		}
	}

	if o.SecondaryPrefixStatus != nil {
		v := *o.SecondaryPrefixStatus

		if err = d.Set("secondary_prefix_status", v); err != nil {
			return diag.Errorf("error reading secondary_prefix_status: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func expandSystemNat64SecondaryPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNat64SecondaryPrefix, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNat64SecondaryPrefix

	for i := range l {
		tmp := models.SystemNat64SecondaryPrefix{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nat64_prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Nat64Prefix = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemNat64(d *schema.ResourceData, sv string) (*models.SystemNat64, diag.Diagnostics) {
	obj := models.SystemNat64{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("always_synthesize_aaaa_record"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("always_synthesize_aaaa_record", sv)
				diags = append(diags, e)
			}
			obj.AlwaysSynthesizeAaaaRecord = &v2
		}
	}
	if v1, ok := d.GetOk("generate_ipv6_fragment_header"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("generate_ipv6_fragment_header", sv)
				diags = append(diags, e)
			}
			obj.GenerateIpv6FragmentHeader = &v2
		}
	}
	if v1, ok := d.GetOk("nat46_force_ipv4_packet_forwarding"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nat46_force_ipv4_packet_forwarding", sv)
				diags = append(diags, e)
			}
			obj.Nat46ForceIpv4PacketForwarding = &v2
		}
	}
	if v1, ok := d.GetOk("nat64_prefix"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nat64_prefix", sv)
				diags = append(diags, e)
			}
			obj.Nat64Prefix = &v2
		}
	}
	if v, ok := d.GetOk("secondary_prefix"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("secondary_prefix", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNat64SecondaryPrefix(d, v, "secondary_prefix", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SecondaryPrefix = t
		}
	} else if d.HasChange("secondary_prefix") {
		old, new := d.GetChange("secondary_prefix")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SecondaryPrefix = &[]models.SystemNat64SecondaryPrefix{}
		}
	}
	if v1, ok := d.GetOk("secondary_prefix_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secondary_prefix_status", sv)
				diags = append(diags, e)
			}
			obj.SecondaryPrefixStatus = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	return &obj, diags
}
