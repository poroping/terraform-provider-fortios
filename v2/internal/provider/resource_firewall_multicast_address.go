// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure multicast addresses.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceFirewallMulticastAddress() *schema.Resource {
	return &schema.Resource{
		Description: "Configure multicast addresses.",

		CreateContext: resourceFirewallMulticastAddressCreate,
		ReadContext:   resourceFirewallMulticastAddressRead,
		UpdateContext: resourceFirewallMulticastAddressUpdate,
		DeleteContext: resourceFirewallMulticastAddressDelete,

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
			"associated_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.",
				Optional:    true,
				Computed:    true,
			},
			"color": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"end_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Final IPv4 address (inclusive) in the range for the address.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Multicast address name.",
				ForceNew:    true,
				Required:    true,
			},
			"start_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "First IPv4 address (inclusive) in the range for the address.",
				Optional:    true,
				Computed:    true,
			},
			"subnet": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

				Description: "Broadcast address and subnet.",
				Optional:    true,
				Computed:    true,
			},
			"tagging": {
				Type:        schema.TypeList,
				Description: "Config object tagging.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Tag category.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Tagging entry name.",
							Optional:    true,
							Computed:    true,
						},
						"tags": {
							Type:        schema.TypeList,
							Description: "Tags.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Tag name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"multicastrange", "broadcastmask"}, false),

				Description: "Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address.",
				Optional:    true,
				Computed:    true,
			},
			"visibility": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable visibility of the multicast address on the GUI.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallMulticastAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallMulticastAddress resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallMulticastAddress(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallMulticastAddress(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallMulticastAddress")
	}

	return resourceFirewallMulticastAddressRead(ctx, d, meta)
}

func resourceFirewallMulticastAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallMulticastAddress(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallMulticastAddress(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallMulticastAddress resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallMulticastAddress")
	}

	return resourceFirewallMulticastAddressRead(ctx, d, meta)
}

func resourceFirewallMulticastAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallMulticastAddress(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallMulticastAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallMulticastAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallMulticastAddress(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallMulticastAddress resource: %v", err)
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

	diags := refreshObjectFirewallMulticastAddress(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallMulticastAddressTagging(v *[]models.FirewallMulticastAddressTagging, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Tags; tmp != nil {
				v["tags"] = flattenFirewallMulticastAddressTaggingTags(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenFirewallMulticastAddressTaggingTags(v *[]models.FirewallMulticastAddressTaggingTags, sort bool) interface{} {
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

func refreshObjectFirewallMulticastAddress(d *schema.ResourceData, o *models.FirewallMulticastAddress, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AssociatedInterface != nil {
		v := *o.AssociatedInterface

		if err = d.Set("associated_interface", v); err != nil {
			return diag.Errorf("error reading associated_interface: %v", err)
		}
	}

	if o.Color != nil {
		v := *o.Color

		if err = d.Set("color", v); err != nil {
			return diag.Errorf("error reading color: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.EndIp != nil {
		v := *o.EndIp

		if err = d.Set("end_ip", v); err != nil {
			return diag.Errorf("error reading end_ip: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.StartIp != nil {
		v := *o.StartIp

		if err = d.Set("start_ip", v); err != nil {
			return diag.Errorf("error reading start_ip: %v", err)
		}
	}

	if o.Subnet != nil {
		v := *o.Subnet
		if current, ok := d.GetOk("subnet"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("subnet", v); err != nil {
			return diag.Errorf("error reading subnet: %v", err)
		}
	}

	if o.Tagging != nil {
		if err = d.Set("tagging", flattenFirewallMulticastAddressTagging(o.Tagging, sort)); err != nil {
			return diag.Errorf("error reading tagging: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.Visibility != nil {
		v := *o.Visibility

		if err = d.Set("visibility", v); err != nil {
			return diag.Errorf("error reading visibility: %v", err)
		}
	}

	return nil
}

func expandFirewallMulticastAddressTagging(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallMulticastAddressTagging, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallMulticastAddressTagging

	for i := range l {
		tmp := models.FirewallMulticastAddressTagging{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallMulticastAddressTaggingTags(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallMulticastAddressTaggingTags
			// 	}
			tmp.Tags = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallMulticastAddressTaggingTags(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallMulticastAddressTaggingTags, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallMulticastAddressTaggingTags

	for i := range l {
		tmp := models.FirewallMulticastAddressTaggingTags{}
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

func getObjectFirewallMulticastAddress(d *schema.ResourceData, sv string) (*models.FirewallMulticastAddress, diag.Diagnostics) {
	obj := models.FirewallMulticastAddress{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("associated_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("associated_interface", sv)
				diags = append(diags, e)
			}
			obj.AssociatedInterface = &v2
		}
	}
	if v1, ok := d.GetOk("color"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("color", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Color = &tmp
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("end_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("end_ip", sv)
				diags = append(diags, e)
			}
			obj.EndIp = &v2
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
	if v1, ok := d.GetOk("start_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("start_ip", sv)
				diags = append(diags, e)
			}
			obj.StartIp = &v2
		}
	}
	if v1, ok := d.GetOk("subnet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("subnet", sv)
				diags = append(diags, e)
			}
			obj.Subnet = &v2
		}
	}
	if v, ok := d.GetOk("tagging"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("tagging", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallMulticastAddressTagging(d, v, "tagging", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Tagging = t
		}
	} else if d.HasChange("tagging") {
		old, new := d.GetChange("tagging")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Tagging = &[]models.FirewallMulticastAddressTagging{}
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("visibility"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("visibility", sv)
				diags = append(diags, e)
			}
			obj.Visibility = &v2
		}
	}
	return &obj, diags
}
