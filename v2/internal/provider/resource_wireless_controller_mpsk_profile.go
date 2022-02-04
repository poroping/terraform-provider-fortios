// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure MPSK profile.

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

func resourceWirelessControllerMpskProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure MPSK profile.",

		CreateContext: resourceWirelessControllerMpskProfileCreate,
		ReadContext:   resourceWirelessControllerMpskProfileRead,
		UpdateContext: resourceWirelessControllerMpskProfileUpdate,
		DeleteContext: resourceWirelessControllerMpskProfileDelete,

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
			"mpsk_concurrent_clients": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).",
				Optional:    true,
				Computed:    true,
			},
			"mpsk_group": {
				Type:        schema.TypeList,
				Description: "List of multiple PSK groups.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mpsk_key": {
							Type:        schema.TypeList,
							Description: "List of multiple PSK entries.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"comment": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),

										Description: "Comment.",
										Optional:    true,
										Computed:    true,
									},
									"concurrent_client_limit_type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"default", "unlimited", "specified"}, false),

										Description: "MPSK client limit type options.",
										Optional:    true,
										Computed:    true,
									},
									"concurrent_clients": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Number of clients that can connect using this pre-shared key (1 - 65535, default is 256).",
										Optional:    true,
										Computed:    true,
									},
									"mac": {
										Type: schema.TypeString,

										Description: "MAC address.",
										Optional:    true,
										Computed:    true,
									},
									"mpsk_schedules": {
										Type:        schema.TypeList,
										Description: "Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 35),

													Description: "Schedule name.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Pre-shared key name.",
										Optional:    true,
										Computed:    true,
									},
									"passphrase": {
										Type: schema.TypeString,

										Description: "WPA Pre-shared key.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
								},
							},
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "MPSK group name.",
							Optional:    true,
							Computed:    true,
						},
						"vlan_id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),

							Description: "Optional VLAN ID.",
							Optional:    true,
							Computed:    true,
						},
						"vlan_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"no-vlan", "fixed-vlan"}, false),

							Description: "MPSK group VLAN options.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "MPSK profile name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceWirelessControllerMpskProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerMpskProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerMpskProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerMpskProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerMpskProfile")
	}

	return resourceWirelessControllerMpskProfileRead(ctx, d, meta)
}

func resourceWirelessControllerMpskProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerMpskProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerMpskProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerMpskProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerMpskProfile")
	}

	return resourceWirelessControllerMpskProfileRead(ctx, d, meta)
}

func resourceWirelessControllerMpskProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerMpskProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerMpskProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerMpskProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerMpskProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerMpskProfile resource: %v", err)
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

	diags := refreshObjectWirelessControllerMpskProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerMpskProfileMpskGroup(v *[]models.WirelessControllerMpskProfileMpskGroup, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.MpskKey; tmp != nil {
				v["mpsk_key"] = flattenWirelessControllerMpskProfileMpskGroupMpskKey(tmp, sort)
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.VlanId; tmp != nil {
				v["vlan_id"] = *tmp
			}

			if tmp := cfg.VlanType; tmp != nil {
				v["vlan_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKey(v *[]models.WirelessControllerMpskProfileMpskGroupMpskKey, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Comment; tmp != nil {
				v["comment"] = *tmp
			}

			if tmp := cfg.ConcurrentClientLimitType; tmp != nil {
				v["concurrent_client_limit_type"] = *tmp
			}

			if tmp := cfg.ConcurrentClients; tmp != nil {
				v["concurrent_clients"] = *tmp
			}

			if tmp := cfg.Mac; tmp != nil {
				v["mac"] = *tmp
			}

			if tmp := cfg.MpskSchedules; tmp != nil {
				v["mpsk_schedules"] = flattenWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(tmp, sort)
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Passphrase; tmp != nil {
				v["passphrase"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(v *[]models.WirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules, sort bool) interface{} {
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

func refreshObjectWirelessControllerMpskProfile(d *schema.ResourceData, o *models.WirelessControllerMpskProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.MpskConcurrentClients != nil {
		v := *o.MpskConcurrentClients

		if err = d.Set("mpsk_concurrent_clients", v); err != nil {
			return diag.Errorf("error reading mpsk_concurrent_clients: %v", err)
		}
	}

	if o.MpskGroup != nil {
		if err = d.Set("mpsk_group", flattenWirelessControllerMpskProfileMpskGroup(o.MpskGroup, sort)); err != nil {
			return diag.Errorf("error reading mpsk_group: %v", err)
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

func expandWirelessControllerMpskProfileMpskGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerMpskProfileMpskGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerMpskProfileMpskGroup

	for i := range l {
		tmp := models.WirelessControllerMpskProfileMpskGroup{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.mpsk_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerMpskProfileMpskGroupMpskKey(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerMpskProfileMpskGroupMpskKey
			// 	}
			tmp.MpskKey = v2

		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.VlanId = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VlanType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKey(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerMpskProfileMpskGroupMpskKey, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerMpskProfileMpskGroupMpskKey

	for i := range l {
		tmp := models.WirelessControllerMpskProfileMpskGroupMpskKey{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.comment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comment = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.concurrent_client_limit_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ConcurrentClientLimitType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.concurrent_clients", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ConcurrentClients = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mac", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mac = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mpsk_schedules", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules
			// 	}
			tmp.MpskSchedules = v2

		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.passphrase", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Passphrase = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules

	for i := range l {
		tmp := models.WirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules{}
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

func getObjectWirelessControllerMpskProfile(d *schema.ResourceData, sv string) (*models.WirelessControllerMpskProfile, diag.Diagnostics) {
	obj := models.WirelessControllerMpskProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("mpsk_concurrent_clients"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mpsk_concurrent_clients", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MpskConcurrentClients = &tmp
		}
	}
	if v, ok := d.GetOk("mpsk_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("mpsk_group", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerMpskProfileMpskGroup(d, v, "mpsk_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.MpskGroup = t
		}
	} else if d.HasChange("mpsk_group") {
		old, new := d.GetChange("mpsk_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.MpskGroup = &[]models.WirelessControllerMpskProfileMpskGroup{}
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
