// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: FortiGate interfaces added to the SD-WAN.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemSdwanMembers() *schema.Resource {
	return &schema.Resource{
		Description: "FortiGate interfaces added to the SD-WAN.",

		CreateContext: resourceSystemSdwanMembersCreate,
		ReadContext:   resourceSystemSdwanMembersRead,
		UpdateContext: resourceSystemSdwanMembersUpdate,
		DeleteContext: resourceSystemSdwanMembersDelete,

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
				Type:         schema.TypeBool,
				Description:  "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"seq_num"},
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comments.",
				Optional:    true,
				Computed:    true,
			},
			"cost": {
				Type: schema.TypeInt,

				Description: "Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"gateway": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.",
				Optional:    true,
				Computed:    true,
			},
			"gateway6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 gateway.",
				Optional:         true,
				Computed:         true,
			},
			"ingress_spillover_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Interface name.",
				Optional:    true,
				Computed:    true,
			},
			"priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Priority of the interface for IPv4 (0 - 65535, default = 0). Used for SD-WAN rules or priority rules.",
				Optional:    true,
				Computed:    true,
			},
			"priority6": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Priority of the interface for IPv6 (1 - 65535, default = 1024). Used for SD-WAN rules or priority rules.",
				Optional:    true,
				Computed:    true,
			},
			"seq_num": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 512),

				Description: "Sequence number(1-512).",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP address used in the health-check packet to the server.",
				Optional:    true,
				Computed:    true,
			},
			"source6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Source IPv6 address used in the health-check packet to the server.",
				Optional:         true,
				Computed:         true,
			},
			"spillover_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable this interface in the SD-WAN.",
				Optional:    true,
				Computed:    true,
			},
			"volume_ratio": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).",
				Optional:    true,
				Computed:    true,
			},
			"weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.",
				Optional:    true,
				Computed:    true,
			},
			"zone": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Zone name.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemSdwanMembersCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "seq_num"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemSdwanMembers resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemSdwanMembers(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSdwanMembers(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSdwanMembers")
	}

	return resourceSystemSdwanMembersRead(ctx, d, meta)
}

func resourceSystemSdwanMembersUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSdwanMembers(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSdwanMembers(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSdwanMembers resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSdwanMembers")
	}

	return resourceSystemSdwanMembersRead(ctx, d, meta)
}

func resourceSystemSdwanMembersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSdwanMembers(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSdwanMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanMembersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSdwanMembers(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSdwanMembers resource: %v", err)
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

	diags := refreshObjectSystemSdwanMembers(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemSdwanMembers(d *schema.ResourceData, o *models.SystemSdwanMembers, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Cost != nil {
		v := *o.Cost

		if err = d.Set("cost", v); err != nil {
			return diag.Errorf("error reading cost: %v", err)
		}
	}

	if o.Gateway != nil {
		v := *o.Gateway

		if err = d.Set("gateway", v); err != nil {
			return diag.Errorf("error reading gateway: %v", err)
		}
	}

	if o.Gateway6 != nil {
		v := *o.Gateway6

		if err = d.Set("gateway6", v); err != nil {
			return diag.Errorf("error reading gateway6: %v", err)
		}
	}

	if o.IngressSpilloverThreshold != nil {
		v := *o.IngressSpilloverThreshold

		if err = d.Set("ingress_spillover_threshold", v); err != nil {
			return diag.Errorf("error reading ingress_spillover_threshold: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Priority != nil {
		v := *o.Priority

		if err = d.Set("priority", v); err != nil {
			return diag.Errorf("error reading priority: %v", err)
		}
	}

	if o.Priority6 != nil {
		v := *o.Priority6

		if err = d.Set("priority6", v); err != nil {
			return diag.Errorf("error reading priority6: %v", err)
		}
	}

	if o.SeqNum != nil {
		v := *o.SeqNum

		if err = d.Set("seq_num", v); err != nil {
			return diag.Errorf("error reading seq_num: %v", err)
		}
	}

	if o.Source != nil {
		v := *o.Source

		if err = d.Set("source", v); err != nil {
			return diag.Errorf("error reading source: %v", err)
		}
	}

	if o.Source6 != nil {
		v := *o.Source6

		if err = d.Set("source6", v); err != nil {
			return diag.Errorf("error reading source6: %v", err)
		}
	}

	if o.SpilloverThreshold != nil {
		v := *o.SpilloverThreshold

		if err = d.Set("spillover_threshold", v); err != nil {
			return diag.Errorf("error reading spillover_threshold: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.VolumeRatio != nil {
		v := *o.VolumeRatio

		if err = d.Set("volume_ratio", v); err != nil {
			return diag.Errorf("error reading volume_ratio: %v", err)
		}
	}

	if o.Weight != nil {
		v := *o.Weight

		if err = d.Set("weight", v); err != nil {
			return diag.Errorf("error reading weight: %v", err)
		}
	}

	if o.Zone != nil {
		v := *o.Zone

		if err = d.Set("zone", v); err != nil {
			return diag.Errorf("error reading zone: %v", err)
		}
	}

	return nil
}

func getObjectSystemSdwanMembers(d *schema.ResourceData, sv string) (*models.SystemSdwanMembers, diag.Diagnostics) {
	obj := models.SystemSdwanMembers{}
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
	if v1, ok := d.GetOk("cost"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cost", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Cost = &tmp
		}
	}
	if v1, ok := d.GetOk("gateway"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gateway", sv)
				diags = append(diags, e)
			}
			obj.Gateway = &v2
		}
	}
	if v1, ok := d.GetOk("gateway6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gateway6", sv)
				diags = append(diags, e)
			}
			obj.Gateway6 = &v2
		}
	}
	if v1, ok := d.GetOk("ingress_spillover_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ingress_spillover_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IngressSpilloverThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("priority"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Priority = &tmp
		}
	}
	if v1, ok := d.GetOk("priority6"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("priority6", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Priority6 = &tmp
		}
	}
	if v1, ok := d.GetOk("seq_num"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("seq_num", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SeqNum = &tmp
		}
	}
	if v1, ok := d.GetOk("source"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source", sv)
				diags = append(diags, e)
			}
			obj.Source = &v2
		}
	}
	if v1, ok := d.GetOk("source6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source6", sv)
				diags = append(diags, e)
			}
			obj.Source6 = &v2
		}
	}
	if v1, ok := d.GetOk("spillover_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spillover_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SpilloverThreshold = &tmp
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
	if v1, ok := d.GetOk("volume_ratio"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("volume_ratio", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.VolumeRatio = &tmp
		}
	}
	if v1, ok := d.GetOk("weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Weight = &tmp
		}
	}
	if v1, ok := d.GetOk("zone"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("zone", sv)
				diags = append(diags, e)
			}
			obj.Zone = &v2
		}
	}
	return &obj, diags
}
