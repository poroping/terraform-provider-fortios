// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Overlay Controller VPN settings.

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

func resourceVpnOcvpn() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Overlay Controller VPN settings.",

		CreateContext: resourceVpnOcvpnCreate,
		ReadContext:   resourceVpnOcvpnRead,
		UpdateContext: resourceVpnOcvpnUpdate,
		DeleteContext: resourceVpnOcvpnDelete,

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
			"auto_discovery": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable auto-discovery shortcuts.",
				Optional:    true,
				Computed:    true,
			},
			"auto_discovery_shortcut_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"independent", "dependent"}, false),

				Description: "Control deletion of child short-cut tunnels when the parent tunnel goes down.",
				Optional:    true,
				Computed:    true,
			},
			"eap": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable EAP client authentication.",
				Optional:    true,
				Computed:    true,
			},
			"eap_users": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "EAP authentication user group.",
				Optional:    true,
				Computed:    true,
			},
			"forticlient_access": {
				Type:        schema.TypeList,
				Description: "Configure FortiClient settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_groups": {
							Type:        schema.TypeList,
							Description: "FortiClient user authentication groups.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_group": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Authentication user group for FortiClient access.",
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Group name.",
										Optional:    true,
										Computed:    true,
									},
									"overlays": {
										Type:        schema.TypeList,
										Description: "OCVPN overlays to allow access to.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"overlay_name": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),

													Description: "Overlay name.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"psksecret": {
							Type: schema.TypeString,

							Description: "Pre-shared secret for FortiClient PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable FortiClient to access OCVPN networks.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ip_allocation_block": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

				Description: "Class B subnet reserved for private IP address assignment.",
				Optional:    true,
				Computed:    true,
			},
			"multipath": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable multipath redundancy.",
				Optional:    true,
				Computed:    true,
			},
			"nat": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NAT support.",
				Optional:    true,
				Computed:    true,
			},
			"overlays": {
				Type:        schema.TypeList,
				Description: "Network overlays to register with Overlay Controller VPN service.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assign_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable mode-cfg address assignment.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"inter_overlay": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),

							Description: "Allow or deny traffic from other overlays.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_end_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "End of IPv4 range.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_start_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Start of IPv4 range.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Overlay name.",
							Optional:    true,
							Computed:    true,
						},
						"overlay_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Overlay name.",
							Optional:    true,
							Computed:    true,
						},
						"subnets": {
							Type:        schema.TypeList,
							Description: "Internal subnets to register with OCVPN service.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "ID.",
										Optional:    true,
										Computed:    true,
									},
									"interface": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),

										Description: "LAN interface.",
										Optional:    true,
										Computed:    true,
									},
									"subnet": {
										Type:         schema.TypeString,
										ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

										Description: "IPv4 address and subnet mask.",
										Optional:    true,
										Computed:    true,
									},
									"type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"subnet", "interface"}, false),

										Description: "Subnet type.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"poll_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 120),

				Description: "Overlay Controller VPN polling interval.",
				Optional:    true,
				Computed:    true,
			},
			"role": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"spoke", "primary-hub", "secondary-hub"}, false),

				Description: "Set device role.",
				Optional:    true,
				Computed:    true,
			},
			"sdwan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable adding OCVPN tunnels to SD-WAN.",
				Optional:    true,
				Computed:    true,
			},
			"sdwan_zone": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Set SD-WAN zone.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Overlay Controller cloud assisted VPN.",
				Optional:    true,
				Computed:    true,
			},
			"wan_interface": {
				Type:        schema.TypeList,
				Description: "FortiGate WAN interfaces to use with OCVPN.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceVpnOcvpnCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnOcvpn(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnOcvpn(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnOcvpn")
	}

	return resourceVpnOcvpnRead(ctx, d, meta)
}

func resourceVpnOcvpnUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnOcvpn(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnOcvpn(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnOcvpn resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnOcvpn")
	}

	return resourceVpnOcvpnRead(ctx, d, meta)
}

func resourceVpnOcvpnDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectVpnOcvpn(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateVpnOcvpn(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnOcvpn resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnOcvpnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnOcvpn(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnOcvpn resource: %v", err)
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

	diags := refreshObjectVpnOcvpn(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenVpnOcvpnForticlientAccess(v *[]models.VpnOcvpnForticlientAccess, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AuthGroups; tmp != nil {
				v["auth_groups"] = flattenVpnOcvpnForticlientAccessAuthGroups(tmp, sort)
			}

			if tmp := cfg.Psksecret; tmp != nil {
				v["psksecret"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenVpnOcvpnForticlientAccessAuthGroups(v *[]models.VpnOcvpnForticlientAccessAuthGroups, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AuthGroup; tmp != nil {
				v["auth_group"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Overlays; tmp != nil {
				v["overlays"] = flattenVpnOcvpnForticlientAccessAuthGroupsOverlays(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenVpnOcvpnForticlientAccessAuthGroupsOverlays(v *[]models.VpnOcvpnForticlientAccessAuthGroupsOverlays, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.OverlayName; tmp != nil {
				v["overlay_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "overlay_name")
	}

	return flat
}

func flattenVpnOcvpnOverlays(v *[]models.VpnOcvpnOverlays, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AssignIp; tmp != nil {
				v["assign_ip"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.InterOverlay; tmp != nil {
				v["inter_overlay"] = *tmp
			}

			if tmp := cfg.Ipv4EndIp; tmp != nil {
				v["ipv4_end_ip"] = *tmp
			}

			if tmp := cfg.Ipv4StartIp; tmp != nil {
				v["ipv4_start_ip"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.OverlayName; tmp != nil {
				v["overlay_name"] = *tmp
			}

			if tmp := cfg.Subnets; tmp != nil {
				v["subnets"] = flattenVpnOcvpnOverlaysSubnets(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "overlay_name")
	}

	return flat
}

func flattenVpnOcvpnOverlaysSubnets(v *[]models.VpnOcvpnOverlaysSubnets, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.Subnet; tmp != nil {
				v["subnet"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenVpnOcvpnWanInterface(v *[]models.VpnOcvpnWanInterface, sort bool) interface{} {
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

func refreshObjectVpnOcvpn(d *schema.ResourceData, o *models.VpnOcvpn, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AutoDiscovery != nil {
		v := *o.AutoDiscovery

		if err = d.Set("auto_discovery", v); err != nil {
			return diag.Errorf("error reading auto_discovery: %v", err)
		}
	}

	if o.AutoDiscoveryShortcutMode != nil {
		v := *o.AutoDiscoveryShortcutMode

		if err = d.Set("auto_discovery_shortcut_mode", v); err != nil {
			return diag.Errorf("error reading auto_discovery_shortcut_mode: %v", err)
		}
	}

	if o.Eap != nil {
		v := *o.Eap

		if err = d.Set("eap", v); err != nil {
			return diag.Errorf("error reading eap: %v", err)
		}
	}

	if o.EapUsers != nil {
		v := *o.EapUsers

		if err = d.Set("eap_users", v); err != nil {
			return diag.Errorf("error reading eap_users: %v", err)
		}
	}

	if o.ForticlientAccess != nil {
		if err = d.Set("forticlient_access", flattenVpnOcvpnForticlientAccess(o.ForticlientAccess, sort)); err != nil {
			return diag.Errorf("error reading forticlient_access: %v", err)
		}
	}

	if o.IpAllocationBlock != nil {
		v := *o.IpAllocationBlock
		if current, ok := d.GetOk("ip_allocation_block"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("ip_allocation_block", v); err != nil {
			return diag.Errorf("error reading ip_allocation_block: %v", err)
		}
	}

	if o.Multipath != nil {
		v := *o.Multipath

		if err = d.Set("multipath", v); err != nil {
			return diag.Errorf("error reading multipath: %v", err)
		}
	}

	if o.Nat != nil {
		v := *o.Nat

		if err = d.Set("nat", v); err != nil {
			return diag.Errorf("error reading nat: %v", err)
		}
	}

	if o.Overlays != nil {
		if err = d.Set("overlays", flattenVpnOcvpnOverlays(o.Overlays, sort)); err != nil {
			return diag.Errorf("error reading overlays: %v", err)
		}
	}

	if o.PollInterval != nil {
		v := *o.PollInterval

		if err = d.Set("poll_interval", v); err != nil {
			return diag.Errorf("error reading poll_interval: %v", err)
		}
	}

	if o.Role != nil {
		v := *o.Role

		if err = d.Set("role", v); err != nil {
			return diag.Errorf("error reading role: %v", err)
		}
	}

	if o.Sdwan != nil {
		v := *o.Sdwan

		if err = d.Set("sdwan", v); err != nil {
			return diag.Errorf("error reading sdwan: %v", err)
		}
	}

	if o.SdwanZone != nil {
		v := *o.SdwanZone

		if err = d.Set("sdwan_zone", v); err != nil {
			return diag.Errorf("error reading sdwan_zone: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.WanInterface != nil {
		if err = d.Set("wan_interface", flattenVpnOcvpnWanInterface(o.WanInterface, sort)); err != nil {
			return diag.Errorf("error reading wan_interface: %v", err)
		}
	}

	return nil
}

func expandVpnOcvpnForticlientAccess(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnOcvpnForticlientAccess, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnOcvpnForticlientAccess

	for i := range l {
		tmp := models.VpnOcvpnForticlientAccess{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_groups", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnOcvpnForticlientAccessAuthGroups(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnOcvpnForticlientAccessAuthGroups
			// 	}
			tmp.AuthGroups = v2

		}

		pre_append = fmt.Sprintf("%s.%d.psksecret", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Psksecret = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnOcvpnForticlientAccessAuthGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnOcvpnForticlientAccessAuthGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnOcvpnForticlientAccessAuthGroups

	for i := range l {
		tmp := models.VpnOcvpnForticlientAccessAuthGroups{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthGroup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.overlays", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnOcvpnForticlientAccessAuthGroupsOverlays(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnOcvpnForticlientAccessAuthGroupsOverlays
			// 	}
			tmp.Overlays = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnOcvpnForticlientAccessAuthGroupsOverlays(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnOcvpnForticlientAccessAuthGroupsOverlays, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnOcvpnForticlientAccessAuthGroupsOverlays

	for i := range l {
		tmp := models.VpnOcvpnForticlientAccessAuthGroupsOverlays{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.overlay_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverlayName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnOcvpnOverlays(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnOcvpnOverlays, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnOcvpnOverlays

	for i := range l {
		tmp := models.VpnOcvpnOverlays{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.assign_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AssignIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.inter_overlay", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterOverlay = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_end_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4EndIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_start_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4StartIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.overlay_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverlayName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.subnets", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnOcvpnOverlaysSubnets(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnOcvpnOverlaysSubnets
			// 	}
			tmp.Subnets = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnOcvpnOverlaysSubnets(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnOcvpnOverlaysSubnets, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnOcvpnOverlaysSubnets

	for i := range l {
		tmp := models.VpnOcvpnOverlaysSubnets{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.subnet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Subnet = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnOcvpnWanInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnOcvpnWanInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnOcvpnWanInterface

	for i := range l {
		tmp := models.VpnOcvpnWanInterface{}
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

func getObjectVpnOcvpn(d *schema.ResourceData, sv string) (*models.VpnOcvpn, diag.Diagnostics) {
	obj := models.VpnOcvpn{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auto_discovery"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_discovery", sv)
				diags = append(diags, e)
			}
			obj.AutoDiscovery = &v2
		}
	}
	if v1, ok := d.GetOk("auto_discovery_shortcut_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("auto_discovery_shortcut_mode", sv)
				diags = append(diags, e)
			}
			obj.AutoDiscoveryShortcutMode = &v2
		}
	}
	if v1, ok := d.GetOk("eap"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eap", sv)
				diags = append(diags, e)
			}
			obj.Eap = &v2
		}
	}
	if v1, ok := d.GetOk("eap_users"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eap_users", sv)
				diags = append(diags, e)
			}
			obj.EapUsers = &v2
		}
	}
	if v, ok := d.GetOk("forticlient_access"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("forticlient_access", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnOcvpnForticlientAccess(d, v, "forticlient_access", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ForticlientAccess = t
		}
	} else if d.HasChange("forticlient_access") {
		old, new := d.GetChange("forticlient_access")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ForticlientAccess = &[]models.VpnOcvpnForticlientAccess{}
		}
	}
	if v1, ok := d.GetOk("ip_allocation_block"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("ip_allocation_block", sv)
				diags = append(diags, e)
			}
			obj.IpAllocationBlock = &v2
		}
	}
	if v1, ok := d.GetOk("multipath"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("multipath", sv)
				diags = append(diags, e)
			}
			obj.Multipath = &v2
		}
	}
	if v1, ok := d.GetOk("nat"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nat", sv)
				diags = append(diags, e)
			}
			obj.Nat = &v2
		}
	}
	if v, ok := d.GetOk("overlays"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("overlays", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnOcvpnOverlays(d, v, "overlays", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Overlays = t
		}
	} else if d.HasChange("overlays") {
		old, new := d.GetChange("overlays")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Overlays = &[]models.VpnOcvpnOverlays{}
		}
	}
	if v1, ok := d.GetOk("poll_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("poll_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PollInterval = &tmp
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
	if v1, ok := d.GetOk("sdwan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("sdwan", sv)
				diags = append(diags, e)
			}
			obj.Sdwan = &v2
		}
	}
	if v1, ok := d.GetOk("sdwan_zone"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("sdwan_zone", sv)
				diags = append(diags, e)
			}
			obj.SdwanZone = &v2
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
	if v, ok := d.GetOk("wan_interface"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("wan_interface", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnOcvpnWanInterface(d, v, "wan_interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.WanInterface = t
		}
	} else if d.HasChange("wan_interface") {
		old, new := d.GetChange("wan_interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.WanInterface = &[]models.VpnOcvpnWanInterface{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectVpnOcvpn(d *schema.ResourceData, sv string) (*models.VpnOcvpn, diag.Diagnostics) {
	obj := models.VpnOcvpn{}
	diags := diag.Diagnostics{}

	obj.ForticlientAccess = &[]models.VpnOcvpnForticlientAccess{}
	obj.Overlays = &[]models.VpnOcvpnOverlays{}
	obj.WanInterface = &[]models.VpnOcvpnWanInterface{}

	return &obj, diags
}
