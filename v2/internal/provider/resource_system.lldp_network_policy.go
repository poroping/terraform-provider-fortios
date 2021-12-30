// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
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

func resourceSystemlldpNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure LLDP network policy.",

		CreateContext: resourceSystemlldpNetworkPolicyCreate,
		ReadContext:   resourceSystemlldpNetworkPolicyRead,
		UpdateContext: resourceSystemlldpNetworkPolicyUpdate,
		DeleteContext: resourceSystemlldpNetworkPolicyDelete,

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
				Optional:    true,
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
				Optional:    true,
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
				Optional:    true,
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
				Optional:    true,
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
				Optional:    true,
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
				Optional:    true,
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
				Optional:    true,
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
				Optional:    true,
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

func resourceSystemlldpNetworkPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemlldpNetworkPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemlldpNetworkPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemlldpNetworkPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemlldpNetworkPolicy")
	}

	return resourceSystemlldpNetworkPolicyRead(ctx, d, meta)
}

func resourceSystemlldpNetworkPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemlldpNetworkPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemlldpNetworkPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemlldpNetworkPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemlldpNetworkPolicy")
	}

	return resourceSystemlldpNetworkPolicyRead(ctx, d, meta)
}

func resourceSystemlldpNetworkPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemlldpNetworkPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemlldpNetworkPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemlldpNetworkPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemlldpNetworkPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemlldpNetworkPolicy resource: %v", err)
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

	diags := refreshObjectSystemlldpNetworkPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemlldpNetworkPolicyGuest(v *[]models.SystemlldpNetworkPolicyGuest, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenSystemlldpNetworkPolicyGuestVoiceSignaling(v *[]models.SystemlldpNetworkPolicyGuestVoiceSignaling, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenSystemlldpNetworkPolicySoftphone(v *[]models.SystemlldpNetworkPolicySoftphone, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenSystemlldpNetworkPolicyStreamingVideo(v *[]models.SystemlldpNetworkPolicyStreamingVideo, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenSystemlldpNetworkPolicyVideoConferencing(v *[]models.SystemlldpNetworkPolicyVideoConferencing, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenSystemlldpNetworkPolicyVideoSignaling(v *[]models.SystemlldpNetworkPolicyVideoSignaling, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenSystemlldpNetworkPolicyVoice(v *[]models.SystemlldpNetworkPolicyVoice, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenSystemlldpNetworkPolicyVoiceSignaling(v *[]models.SystemlldpNetworkPolicyVoiceSignaling, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func refreshObjectSystemlldpNetworkPolicy(d *schema.ResourceData, o *models.SystemlldpNetworkPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Guest != nil {
		if err = d.Set("guest", flattenSystemlldpNetworkPolicyGuest(o.Guest, sort)); err != nil {
			return diag.Errorf("error reading guest: %v", err)
		}
	}

	if o.GuestVoiceSignaling != nil {
		if err = d.Set("guest_voice_signaling", flattenSystemlldpNetworkPolicyGuestVoiceSignaling(o.GuestVoiceSignaling, sort)); err != nil {
			return diag.Errorf("error reading guest_voice_signaling: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Softphone != nil {
		if err = d.Set("softphone", flattenSystemlldpNetworkPolicySoftphone(o.Softphone, sort)); err != nil {
			return diag.Errorf("error reading softphone: %v", err)
		}
	}

	if o.StreamingVideo != nil {
		if err = d.Set("streaming_video", flattenSystemlldpNetworkPolicyStreamingVideo(o.StreamingVideo, sort)); err != nil {
			return diag.Errorf("error reading streaming_video: %v", err)
		}
	}

	if o.VideoConferencing != nil {
		if err = d.Set("video_conferencing", flattenSystemlldpNetworkPolicyVideoConferencing(o.VideoConferencing, sort)); err != nil {
			return diag.Errorf("error reading video_conferencing: %v", err)
		}
	}

	if o.VideoSignaling != nil {
		if err = d.Set("video_signaling", flattenSystemlldpNetworkPolicyVideoSignaling(o.VideoSignaling, sort)); err != nil {
			return diag.Errorf("error reading video_signaling: %v", err)
		}
	}

	if o.Voice != nil {
		if err = d.Set("voice", flattenSystemlldpNetworkPolicyVoice(o.Voice, sort)); err != nil {
			return diag.Errorf("error reading voice: %v", err)
		}
	}

	if o.VoiceSignaling != nil {
		if err = d.Set("voice_signaling", flattenSystemlldpNetworkPolicyVoiceSignaling(o.VoiceSignaling, sort)); err != nil {
			return diag.Errorf("error reading voice_signaling: %v", err)
		}
	}

	return nil
}

func expandSystemlldpNetworkPolicyGuest(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemlldpNetworkPolicyGuest, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemlldpNetworkPolicyGuest

	for i := range l {
		tmp := models.SystemlldpNetworkPolicyGuest{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Dscp = &v2
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

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Vlan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemlldpNetworkPolicyGuestVoiceSignaling(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemlldpNetworkPolicyGuestVoiceSignaling, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemlldpNetworkPolicyGuestVoiceSignaling

	for i := range l {
		tmp := models.SystemlldpNetworkPolicyGuestVoiceSignaling{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Dscp = &v2
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

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Vlan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemlldpNetworkPolicySoftphone(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemlldpNetworkPolicySoftphone, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemlldpNetworkPolicySoftphone

	for i := range l {
		tmp := models.SystemlldpNetworkPolicySoftphone{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Dscp = &v2
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

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Vlan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemlldpNetworkPolicyStreamingVideo(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemlldpNetworkPolicyStreamingVideo, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemlldpNetworkPolicyStreamingVideo

	for i := range l {
		tmp := models.SystemlldpNetworkPolicyStreamingVideo{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Dscp = &v2
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

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Vlan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemlldpNetworkPolicyVideoConferencing(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemlldpNetworkPolicyVideoConferencing, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemlldpNetworkPolicyVideoConferencing

	for i := range l {
		tmp := models.SystemlldpNetworkPolicyVideoConferencing{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Dscp = &v2
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

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Vlan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemlldpNetworkPolicyVideoSignaling(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemlldpNetworkPolicyVideoSignaling, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemlldpNetworkPolicyVideoSignaling

	for i := range l {
		tmp := models.SystemlldpNetworkPolicyVideoSignaling{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Dscp = &v2
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

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Vlan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemlldpNetworkPolicyVoice(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemlldpNetworkPolicyVoice, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemlldpNetworkPolicyVoice

	for i := range l {
		tmp := models.SystemlldpNetworkPolicyVoice{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Dscp = &v2
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

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Vlan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemlldpNetworkPolicyVoiceSignaling(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemlldpNetworkPolicyVoiceSignaling, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemlldpNetworkPolicyVoiceSignaling

	for i := range l {
		tmp := models.SystemlldpNetworkPolicyVoiceSignaling{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Dscp = &v2
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

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Vlan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemlldpNetworkPolicy(d *schema.ResourceData, sv string) (*models.SystemlldpNetworkPolicy, diag.Diagnostics) {
	obj := models.SystemlldpNetworkPolicy{}
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
		t, err := expandSystemlldpNetworkPolicyGuest(d, v, "guest", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Guest = t
		}
	} else if d.HasChange("guest") {
		old, new := d.GetChange("guest")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Guest = &[]models.SystemlldpNetworkPolicyGuest{}
		}
	}
	if v, ok := d.GetOk("guest_voice_signaling"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("guest_voice_signaling", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemlldpNetworkPolicyGuestVoiceSignaling(d, v, "guest_voice_signaling", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.GuestVoiceSignaling = t
		}
	} else if d.HasChange("guest_voice_signaling") {
		old, new := d.GetChange("guest_voice_signaling")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.GuestVoiceSignaling = &[]models.SystemlldpNetworkPolicyGuestVoiceSignaling{}
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
		t, err := expandSystemlldpNetworkPolicySoftphone(d, v, "softphone", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Softphone = t
		}
	} else if d.HasChange("softphone") {
		old, new := d.GetChange("softphone")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Softphone = &[]models.SystemlldpNetworkPolicySoftphone{}
		}
	}
	if v, ok := d.GetOk("streaming_video"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("streaming_video", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemlldpNetworkPolicyStreamingVideo(d, v, "streaming_video", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.StreamingVideo = t
		}
	} else if d.HasChange("streaming_video") {
		old, new := d.GetChange("streaming_video")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.StreamingVideo = &[]models.SystemlldpNetworkPolicyStreamingVideo{}
		}
	}
	if v, ok := d.GetOk("video_conferencing"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("video_conferencing", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemlldpNetworkPolicyVideoConferencing(d, v, "video_conferencing", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VideoConferencing = t
		}
	} else if d.HasChange("video_conferencing") {
		old, new := d.GetChange("video_conferencing")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VideoConferencing = &[]models.SystemlldpNetworkPolicyVideoConferencing{}
		}
	}
	if v, ok := d.GetOk("video_signaling"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("video_signaling", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemlldpNetworkPolicyVideoSignaling(d, v, "video_signaling", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VideoSignaling = t
		}
	} else if d.HasChange("video_signaling") {
		old, new := d.GetChange("video_signaling")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VideoSignaling = &[]models.SystemlldpNetworkPolicyVideoSignaling{}
		}
	}
	if v, ok := d.GetOk("voice"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("voice", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemlldpNetworkPolicyVoice(d, v, "voice", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Voice = t
		}
	} else if d.HasChange("voice") {
		old, new := d.GetChange("voice")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Voice = &[]models.SystemlldpNetworkPolicyVoice{}
		}
	}
	if v, ok := d.GetOk("voice_signaling"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("voice_signaling", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemlldpNetworkPolicyVoiceSignaling(d, v, "voice_signaling", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VoiceSignaling = t
		}
	} else if d.HasChange("voice_signaling") {
		old, new := d.GetChange("voice_signaling")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VoiceSignaling = &[]models.SystemlldpNetworkPolicyVoiceSignaling{}
		}
	}
	return &obj, diags
}
