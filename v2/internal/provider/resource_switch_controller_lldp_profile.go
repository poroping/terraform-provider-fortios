// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch LLDP profiles.

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

func resourceSwitchControllerLldpProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch LLDP profiles.",

		CreateContext: resourceSwitchControllerLldpProfileCreate,
		ReadContext:   resourceSwitchControllerLldpProfileRead,
		UpdateContext: resourceSwitchControllerLldpProfileUpdate,
		DeleteContext: resourceSwitchControllerLldpProfileDelete,

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
			"8021_tlvs": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Transmitted IEEE 802.1 TLVs.",
				Optional:         true,
				Computed:         true,
			},
			"8023_tlvs": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Transmitted IEEE 802.3 TLVs.",
				Optional:         true,
				Computed:         true,
			},
			"auto_isl": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable auto inter-switch LAG.",
				Optional:    true,
				Computed:    true,
			},
			"auto_isl_hello_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),

				Description: "Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).",
				Optional:    true,
				Computed:    true,
			},
			"auto_isl_port_group": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 9),

				Description: "Auto inter-switch LAG port group ID (0 - 9).",
				Optional:    true,
				Computed:    true,
			},
			"auto_isl_receive_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 90),

				Description: "Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).",
				Optional:    true,
				Computed:    true,
			},
			"auto_mclag_icl": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable MCLAG inter chassis link.",
				Optional:    true,
				Computed:    true,
			},
			"custom_tlvs": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit custom TLV entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"information_string": {
							Type: schema.TypeString,

							Description: "Organizationally defined information string (0 - 507 hexadecimal bytes).",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "TLV name (not sent).",
							Optional:    true,
							Computed:    true,
						},
						"oui": {
							Type: schema.TypeString,

							Description: "Organizationally unique identifier (OUI), a 3-byte hexadecimal number, for this TLV.",
							Optional:    true,
							Computed:    true,
						},
						"subtype": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Organizationally defined subtype (0 - 255).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"med_location_service": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Location service type name.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable or disable this TLV.",
							Optional:    true,
							Computed:    true,
						},
						"sys_location_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Location service ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"med_network_policy": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assign_vlan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable VLAN assignment when this profile is applied on managed FortiSwitch port.",
							Optional:    true,
							Computed:    true,
						},
						"dscp": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "Advertised Differentiated Services Code Point (DSCP) value, a packet header value indicating the level of service requested for traffic, such as high priority or best effort delivery.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Policy type name.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),

							Description: "Advertised Layer 2 priority (0 - 7; from lowest to highest priority).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable or disable this TLV.",
							Optional:    true,
							Computed:    true,
						},
						"vlan_intf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "VLAN interface to advertise; if configured on port.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"med_tlvs": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Transmitted LLDP-MED TLVs (type-length-value descriptions).",
				Optional:         true,
				Computed:         true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Profile name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceSwitchControllerLldpProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SwitchControllerLldpProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSwitchControllerLldpProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerLldpProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerLldpProfile")
	}

	return resourceSwitchControllerLldpProfileRead(ctx, d, meta)
}

func resourceSwitchControllerLldpProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerLldpProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerLldpProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerLldpProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerLldpProfile")
	}

	return resourceSwitchControllerLldpProfileRead(ctx, d, meta)
}

func resourceSwitchControllerLldpProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSwitchControllerLldpProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerLldpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLldpProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerLldpProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerLldpProfile resource: %v", err)
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

	diags := refreshObjectSwitchControllerLldpProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSwitchControllerLldpProfileCustomTlvs(v *[]models.SwitchControllerLldpProfileCustomTlvs, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.InformationString; tmp != nil {
				v["information_string"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Oui; tmp != nil {
				v["oui"] = *tmp
			}

			if tmp := cfg.Subtype; tmp != nil {
				v["subtype"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSwitchControllerLldpProfileMedLocationService(v *[]models.SwitchControllerLldpProfileMedLocationService, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.SysLocationId; tmp != nil {
				v["sys_location_id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSwitchControllerLldpProfileMedNetworkPolicy(v *[]models.SwitchControllerLldpProfileMedNetworkPolicy, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AssignVlan; tmp != nil {
				v["assign_vlan"] = *tmp
			}

			if tmp := cfg.Dscp; tmp != nil {
				v["dscp"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.VlanIntf; tmp != nil {
				v["vlan_intf"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectSwitchControllerLldpProfile(d *schema.ResourceData, o *models.SwitchControllerLldpProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.The8021Tlvs != nil {
		v := *o.The8021Tlvs

		if err = d.Set("8021_tlvs", v); err != nil {
			return diag.Errorf("error reading 8021_tlvs: %v", err)
		}
	}

	if o.The8023Tlvs != nil {
		v := *o.The8023Tlvs

		if err = d.Set("8023_tlvs", v); err != nil {
			return diag.Errorf("error reading 8023_tlvs: %v", err)
		}
	}

	if o.AutoIsl != nil {
		v := *o.AutoIsl

		if err = d.Set("auto_isl", v); err != nil {
			return diag.Errorf("error reading auto_isl: %v", err)
		}
	}

	if o.AutoIslHelloTimer != nil {
		v := *o.AutoIslHelloTimer

		if err = d.Set("auto_isl_hello_timer", v); err != nil {
			return diag.Errorf("error reading auto_isl_hello_timer: %v", err)
		}
	}

	if o.AutoIslPortGroup != nil {
		v := *o.AutoIslPortGroup

		if err = d.Set("auto_isl_port_group", v); err != nil {
			return diag.Errorf("error reading auto_isl_port_group: %v", err)
		}
	}

	if o.AutoIslReceiveTimeout != nil {
		v := *o.AutoIslReceiveTimeout

		if err = d.Set("auto_isl_receive_timeout", v); err != nil {
			return diag.Errorf("error reading auto_isl_receive_timeout: %v", err)
		}
	}

	if o.AutoMclagIcl != nil {
		v := *o.AutoMclagIcl

		if err = d.Set("auto_mclag_icl", v); err != nil {
			return diag.Errorf("error reading auto_mclag_icl: %v", err)
		}
	}

	if o.CustomTlvs != nil {
		if err = d.Set("custom_tlvs", flattenSwitchControllerLldpProfileCustomTlvs(o.CustomTlvs, sort)); err != nil {
			return diag.Errorf("error reading custom_tlvs: %v", err)
		}
	}

	if o.MedLocationService != nil {
		if err = d.Set("med_location_service", flattenSwitchControllerLldpProfileMedLocationService(o.MedLocationService, sort)); err != nil {
			return diag.Errorf("error reading med_location_service: %v", err)
		}
	}

	if o.MedNetworkPolicy != nil {
		if err = d.Set("med_network_policy", flattenSwitchControllerLldpProfileMedNetworkPolicy(o.MedNetworkPolicy, sort)); err != nil {
			return diag.Errorf("error reading med_network_policy: %v", err)
		}
	}

	if o.MedTlvs != nil {
		v := *o.MedTlvs

		if err = d.Set("med_tlvs", v); err != nil {
			return diag.Errorf("error reading med_tlvs: %v", err)
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

func expandSwitchControllerLldpProfileCustomTlvs(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerLldpProfileCustomTlvs, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerLldpProfileCustomTlvs

	for i := range l {
		tmp := models.SwitchControllerLldpProfileCustomTlvs{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.information_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InformationString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.oui", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Oui = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.subtype", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Subtype = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerLldpProfileMedLocationService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerLldpProfileMedLocationService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerLldpProfileMedLocationService

	for i := range l {
		tmp := models.SwitchControllerLldpProfileMedLocationService{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sys_location_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SysLocationId = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerLldpProfileMedNetworkPolicy, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerLldpProfileMedNetworkPolicy

	for i := range l {
		tmp := models.SwitchControllerLldpProfileMedNetworkPolicy{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.assign_vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AssignVlan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Dscp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Priority = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan_intf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VlanIntf = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSwitchControllerLldpProfile(d *schema.ResourceData, sv string) (*models.SwitchControllerLldpProfile, diag.Diagnostics) {
	obj := models.SwitchControllerLldpProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("8021_tlvs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("8021_tlvs", sv)
				diags = append(diags, e)
			}
			obj.The8021Tlvs = &v2
		}
	}
	if v1, ok := d.GetOk("8023_tlvs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("8023_tlvs", sv)
				diags = append(diags, e)
			}
			obj.The8023Tlvs = &v2
		}
	}
	if v1, ok := d.GetOk("auto_isl"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_isl", sv)
				diags = append(diags, e)
			}
			obj.AutoIsl = &v2
		}
	}
	if v1, ok := d.GetOk("auto_isl_hello_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_isl_hello_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AutoIslHelloTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("auto_isl_port_group"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_isl_port_group", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AutoIslPortGroup = &tmp
		}
	}
	if v1, ok := d.GetOk("auto_isl_receive_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_isl_receive_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AutoIslReceiveTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("auto_mclag_icl"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("auto_mclag_icl", sv)
				diags = append(diags, e)
			}
			obj.AutoMclagIcl = &v2
		}
	}
	if v, ok := d.GetOk("custom_tlvs"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("custom_tlvs", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerLldpProfileCustomTlvs(d, v, "custom_tlvs", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.CustomTlvs = t
		}
	} else if d.HasChange("custom_tlvs") {
		old, new := d.GetChange("custom_tlvs")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.CustomTlvs = &[]models.SwitchControllerLldpProfileCustomTlvs{}
		}
	}
	if v, ok := d.GetOk("med_location_service"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("med_location_service", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerLldpProfileMedLocationService(d, v, "med_location_service", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.MedLocationService = t
		}
	} else if d.HasChange("med_location_service") {
		old, new := d.GetChange("med_location_service")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.MedLocationService = &[]models.SwitchControllerLldpProfileMedLocationService{}
		}
	}
	if v, ok := d.GetOk("med_network_policy"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("med_network_policy", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerLldpProfileMedNetworkPolicy(d, v, "med_network_policy", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.MedNetworkPolicy = t
		}
	} else if d.HasChange("med_network_policy") {
		old, new := d.GetChange("med_network_policy")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.MedNetworkPolicy = &[]models.SwitchControllerLldpProfileMedNetworkPolicy{}
		}
	}
	if v1, ok := d.GetOk("med_tlvs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("med_tlvs", sv)
				diags = append(diags, e)
			}
			obj.MedTlvs = &v2
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
