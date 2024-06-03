// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure LLDP network policy.

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

func resourceSystemLldpNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure LLDP network policy.",

		CreateContext: resourceSystemLldpNetworkPolicyCreate,
		ReadContext:   resourceSystemLldpNetworkPolicyRead,
		UpdateContext: resourceSystemLldpNetworkPolicyUpdate,
		DeleteContext: resourceSystemLldpNetworkPolicyDelete,

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
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"guest": {
				Type:        schema.TypeList,
				Description: "Guest.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),

							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable advertising this policy.",
							Optional:    true,
							Computed:    true,
						},
						"tag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "dot1q", "dot1p"}, false),

							Description: "Advertise tagged or untagged traffic.",
							Optional:    true,
							Computed:    true,
						},
						"vlan": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),

							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"guest_voice_signaling": {
				Type:        schema.TypeList,
				Description: "Guest Voice Signaling.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),

							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable advertising this policy.",
							Optional:    true,
							Computed:    true,
						},
						"tag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "dot1q", "dot1p"}, false),

							Description: "Advertise tagged or untagged traffic.",
							Optional:    true,
							Computed:    true,
						},
						"vlan": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),

							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "LLDP network policy name.",
				ForceNew:    true,
				Required:    true,
			},
			"softphone": {
				Type:        schema.TypeList,
				Description: "Softphone.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),

							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable advertising this policy.",
							Optional:    true,
							Computed:    true,
						},
						"tag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "dot1q", "dot1p"}, false),

							Description: "Advertise tagged or untagged traffic.",
							Optional:    true,
							Computed:    true,
						},
						"vlan": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),

							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"streaming_video": {
				Type:        schema.TypeList,
				Description: "Streaming Video.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),

							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable advertising this policy.",
							Optional:    true,
							Computed:    true,
						},
						"tag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "dot1q", "dot1p"}, false),

							Description: "Advertise tagged or untagged traffic.",
							Optional:    true,
							Computed:    true,
						},
						"vlan": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),

							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"video_conferencing": {
				Type:        schema.TypeList,
				Description: "Video Conferencing.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),

							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable advertising this policy.",
							Optional:    true,
							Computed:    true,
						},
						"tag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "dot1q", "dot1p"}, false),

							Description: "Advertise tagged or untagged traffic.",
							Optional:    true,
							Computed:    true,
						},
						"vlan": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),

							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"video_signaling": {
				Type:        schema.TypeList,
				Description: "Video Signaling.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),

							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable advertising this policy.",
							Optional:    true,
							Computed:    true,
						},
						"tag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "dot1q", "dot1p"}, false),

							Description: "Advertise tagged or untagged traffic.",
							Optional:    true,
							Computed:    true,
						},
						"vlan": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),

							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"voice": {
				Type:        schema.TypeList,
				Description: "Voice.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),

							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable advertising this policy.",
							Optional:    true,
							Computed:    true,
						},
						"tag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "dot1q", "dot1p"}, false),

							Description: "Advertise tagged or untagged traffic.",
							Optional:    true,
							Computed:    true,
						},
						"vlan": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),

							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"voice_signaling": {
				Type:        schema.TypeList,
				Description: "Voice signaling.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),

							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable advertising this policy.",
							Optional:    true,
							Computed:    true,
						},
						"tag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "dot1q", "dot1p"}, false),

							Description: "Advertise tagged or untagged traffic.",
							Optional:    true,
							Computed:    true,
						},
						"vlan": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),

							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemLldpNetworkPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemLldpNetworkPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemLldpNetworkPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemLldpNetworkPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemLldpNetworkPolicy")
	}

	return resourceSystemLldpNetworkPolicyRead(ctx, d, meta)
}

func resourceSystemLldpNetworkPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemLldpNetworkPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemLldpNetworkPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemLldpNetworkPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemLldpNetworkPolicy")
	}

	return resourceSystemLldpNetworkPolicyRead(ctx, d, meta)
}

func resourceSystemLldpNetworkPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemLldpNetworkPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemLldpNetworkPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLldpNetworkPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemLldpNetworkPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemLldpNetworkPolicy resource: %v", err)
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

	diags := refreshObjectSystemLldpNetworkPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemLldpNetworkPolicyGuest(d *schema.ResourceData, v *models.SystemLldpNetworkPolicyGuest, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemLldpNetworkPolicyGuest{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dscp; tmp != nil {
				v["dscp"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Tag; tmp != nil {
				v["tag"] = *tmp
			}

			if tmp := cfg.Vlan; tmp != nil {
				v["vlan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignaling(d *schema.ResourceData, v *models.SystemLldpNetworkPolicyGuestVoiceSignaling, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemLldpNetworkPolicyGuestVoiceSignaling{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dscp; tmp != nil {
				v["dscp"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Tag; tmp != nil {
				v["tag"] = *tmp
			}

			if tmp := cfg.Vlan; tmp != nil {
				v["vlan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemLldpNetworkPolicySoftphone(d *schema.ResourceData, v *models.SystemLldpNetworkPolicySoftphone, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemLldpNetworkPolicySoftphone{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dscp; tmp != nil {
				v["dscp"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Tag; tmp != nil {
				v["tag"] = *tmp
			}

			if tmp := cfg.Vlan; tmp != nil {
				v["vlan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemLldpNetworkPolicyStreamingVideo(d *schema.ResourceData, v *models.SystemLldpNetworkPolicyStreamingVideo, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemLldpNetworkPolicyStreamingVideo{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dscp; tmp != nil {
				v["dscp"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Tag; tmp != nil {
				v["tag"] = *tmp
			}

			if tmp := cfg.Vlan; tmp != nil {
				v["vlan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemLldpNetworkPolicyVideoConferencing(d *schema.ResourceData, v *models.SystemLldpNetworkPolicyVideoConferencing, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemLldpNetworkPolicyVideoConferencing{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dscp; tmp != nil {
				v["dscp"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Tag; tmp != nil {
				v["tag"] = *tmp
			}

			if tmp := cfg.Vlan; tmp != nil {
				v["vlan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemLldpNetworkPolicyVideoSignaling(d *schema.ResourceData, v *models.SystemLldpNetworkPolicyVideoSignaling, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemLldpNetworkPolicyVideoSignaling{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dscp; tmp != nil {
				v["dscp"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Tag; tmp != nil {
				v["tag"] = *tmp
			}

			if tmp := cfg.Vlan; tmp != nil {
				v["vlan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemLldpNetworkPolicyVoice(d *schema.ResourceData, v *models.SystemLldpNetworkPolicyVoice, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemLldpNetworkPolicyVoice{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dscp; tmp != nil {
				v["dscp"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Tag; tmp != nil {
				v["tag"] = *tmp
			}

			if tmp := cfg.Vlan; tmp != nil {
				v["vlan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemLldpNetworkPolicyVoiceSignaling(d *schema.ResourceData, v *models.SystemLldpNetworkPolicyVoiceSignaling, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemLldpNetworkPolicyVoiceSignaling{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dscp; tmp != nil {
				v["dscp"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Tag; tmp != nil {
				v["tag"] = *tmp
			}

			if tmp := cfg.Vlan; tmp != nil {
				v["vlan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectSystemLldpNetworkPolicy(d *schema.ResourceData, o *models.SystemLldpNetworkPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if _, ok := d.GetOk("guest"); ok {
		if o.Guest != nil {
			if err = d.Set("guest", flattenSystemLldpNetworkPolicyGuest(d, o.Guest, "guest", sort)); err != nil {
				return diag.Errorf("error reading guest: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("guest_voice_signaling"); ok {
		if o.GuestVoiceSignaling != nil {
			if err = d.Set("guest_voice_signaling", flattenSystemLldpNetworkPolicyGuestVoiceSignaling(d, o.GuestVoiceSignaling, "guest_voice_signaling", sort)); err != nil {
				return diag.Errorf("error reading guest_voice_signaling: %v", err)
			}
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if _, ok := d.GetOk("softphone"); ok {
		if o.Softphone != nil {
			if err = d.Set("softphone", flattenSystemLldpNetworkPolicySoftphone(d, o.Softphone, "softphone", sort)); err != nil {
				return diag.Errorf("error reading softphone: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("streaming_video"); ok {
		if o.StreamingVideo != nil {
			if err = d.Set("streaming_video", flattenSystemLldpNetworkPolicyStreamingVideo(d, o.StreamingVideo, "streaming_video", sort)); err != nil {
				return diag.Errorf("error reading streaming_video: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("video_conferencing"); ok {
		if o.VideoConferencing != nil {
			if err = d.Set("video_conferencing", flattenSystemLldpNetworkPolicyVideoConferencing(d, o.VideoConferencing, "video_conferencing", sort)); err != nil {
				return diag.Errorf("error reading video_conferencing: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("video_signaling"); ok {
		if o.VideoSignaling != nil {
			if err = d.Set("video_signaling", flattenSystemLldpNetworkPolicyVideoSignaling(d, o.VideoSignaling, "video_signaling", sort)); err != nil {
				return diag.Errorf("error reading video_signaling: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("voice"); ok {
		if o.Voice != nil {
			if err = d.Set("voice", flattenSystemLldpNetworkPolicyVoice(d, o.Voice, "voice", sort)); err != nil {
				return diag.Errorf("error reading voice: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("voice_signaling"); ok {
		if o.VoiceSignaling != nil {
			if err = d.Set("voice_signaling", flattenSystemLldpNetworkPolicyVoiceSignaling(d, o.VoiceSignaling, "voice_signaling", sort)); err != nil {
				return diag.Errorf("error reading voice_signaling: %v", err)
			}
		}
	}

	return nil
}

func expandSystemLldpNetworkPolicyGuest(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemLldpNetworkPolicyGuest, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemLldpNetworkPolicyGuest

	for i := range l {
		tmp := models.SystemLldpNetworkPolicyGuest{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dscp = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Priority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Vlan = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignaling(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemLldpNetworkPolicyGuestVoiceSignaling, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemLldpNetworkPolicyGuestVoiceSignaling

	for i := range l {
		tmp := models.SystemLldpNetworkPolicyGuestVoiceSignaling{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dscp = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Priority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Vlan = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemLldpNetworkPolicySoftphone(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemLldpNetworkPolicySoftphone, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemLldpNetworkPolicySoftphone

	for i := range l {
		tmp := models.SystemLldpNetworkPolicySoftphone{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dscp = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Priority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Vlan = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemLldpNetworkPolicyStreamingVideo(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemLldpNetworkPolicyStreamingVideo, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemLldpNetworkPolicyStreamingVideo

	for i := range l {
		tmp := models.SystemLldpNetworkPolicyStreamingVideo{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dscp = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Priority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Vlan = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemLldpNetworkPolicyVideoConferencing(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemLldpNetworkPolicyVideoConferencing, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemLldpNetworkPolicyVideoConferencing

	for i := range l {
		tmp := models.SystemLldpNetworkPolicyVideoConferencing{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dscp = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Priority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Vlan = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemLldpNetworkPolicyVideoSignaling(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemLldpNetworkPolicyVideoSignaling, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemLldpNetworkPolicyVideoSignaling

	for i := range l {
		tmp := models.SystemLldpNetworkPolicyVideoSignaling{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dscp = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Priority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Vlan = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemLldpNetworkPolicyVoice(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemLldpNetworkPolicyVoice, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemLldpNetworkPolicyVoice

	for i := range l {
		tmp := models.SystemLldpNetworkPolicyVoice{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dscp = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Priority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Vlan = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemLldpNetworkPolicyVoiceSignaling(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemLldpNetworkPolicyVoiceSignaling, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemLldpNetworkPolicyVoiceSignaling

	for i := range l {
		tmp := models.SystemLldpNetworkPolicyVoiceSignaling{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dscp = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Priority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Vlan = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func getObjectSystemLldpNetworkPolicy(d *schema.ResourceData, sv string) (*models.SystemLldpNetworkPolicy, diag.Diagnostics) {
	obj := models.SystemLldpNetworkPolicy{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v, ok := d.GetOk("guest"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("guest", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemLldpNetworkPolicyGuest(d, v, "guest", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Guest = t
		}
	} else if d.HasChange("guest") {
		old, new := d.GetChange("guest")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Guest = &models.SystemLldpNetworkPolicyGuest{}
		}
	}
	if v, ok := d.GetOk("guest_voice_signaling"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("guest_voice_signaling", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemLldpNetworkPolicyGuestVoiceSignaling(d, v, "guest_voice_signaling", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.GuestVoiceSignaling = t
		}
	} else if d.HasChange("guest_voice_signaling") {
		old, new := d.GetChange("guest_voice_signaling")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.GuestVoiceSignaling = &models.SystemLldpNetworkPolicyGuestVoiceSignaling{}
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
	if v, ok := d.GetOk("softphone"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("softphone", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemLldpNetworkPolicySoftphone(d, v, "softphone", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Softphone = t
		}
	} else if d.HasChange("softphone") {
		old, new := d.GetChange("softphone")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Softphone = &models.SystemLldpNetworkPolicySoftphone{}
		}
	}
	if v, ok := d.GetOk("streaming_video"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("streaming_video", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemLldpNetworkPolicyStreamingVideo(d, v, "streaming_video", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.StreamingVideo = t
		}
	} else if d.HasChange("streaming_video") {
		old, new := d.GetChange("streaming_video")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.StreamingVideo = &models.SystemLldpNetworkPolicyStreamingVideo{}
		}
	}
	if v, ok := d.GetOk("video_conferencing"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("video_conferencing", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemLldpNetworkPolicyVideoConferencing(d, v, "video_conferencing", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VideoConferencing = t
		}
	} else if d.HasChange("video_conferencing") {
		old, new := d.GetChange("video_conferencing")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VideoConferencing = &models.SystemLldpNetworkPolicyVideoConferencing{}
		}
	}
	if v, ok := d.GetOk("video_signaling"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("video_signaling", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemLldpNetworkPolicyVideoSignaling(d, v, "video_signaling", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VideoSignaling = t
		}
	} else if d.HasChange("video_signaling") {
		old, new := d.GetChange("video_signaling")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VideoSignaling = &models.SystemLldpNetworkPolicyVideoSignaling{}
		}
	}
	if v, ok := d.GetOk("voice"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("voice", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemLldpNetworkPolicyVoice(d, v, "voice", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Voice = t
		}
	} else if d.HasChange("voice") {
		old, new := d.GetChange("voice")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Voice = &models.SystemLldpNetworkPolicyVoice{}
		}
	}
	if v, ok := d.GetOk("voice_signaling"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("voice_signaling", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemLldpNetworkPolicyVoiceSignaling(d, v, "voice_signaling", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VoiceSignaling = t
		}
	} else if d.HasChange("voice_signaling") {
		old, new := d.GetChange("voice_signaling")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VoiceSignaling = &models.SystemLldpNetworkPolicyVoiceSignaling{}
		}
	}
	return &obj, diags
}
