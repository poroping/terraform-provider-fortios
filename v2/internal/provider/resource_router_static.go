// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 static routing tables.

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

func resourceRouterStatic() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 static routing tables.",

		CreateContext: resourceRouterStaticCreate,
		ReadContext:   resourceRouterStaticRead,
		UpdateContext: resourceRouterStaticUpdate,
		DeleteContext: resourceRouterStaticDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"bfd": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Bidirectional Forwarding Detection (BFD).",
				Optional:    true,
				Computed:    true,
			},
			"blackhole": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable black hole.",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Gateway out interface or tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"distance": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Administrative distance (1 - 255).",
				Optional:    true,
				Computed:    true,
			},
			"dst": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Destination IP and mask for this route.",
				Optional:    true,
				Computed:    true,
			},
			"dstaddr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Name of firewall address or address group.",
				Optional:    true,
				Computed:    true,
			},
			"dynamic_gateway": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable use of dynamic gateway retrieved from a DHCP or PPP server.",
				Optional:    true,
				Computed:    true,
			},
			"gateway": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Gateway IP for this route.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service": {
				Type: schema.TypeInt,

				Description: "Application ID in the Internet service database.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_custom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "Application name in the Internet service custom database.",
				Optional:    true,
				Computed:    true,
			},
			"link_monitor_exempt": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable withdrawal of this static route when link monitor or health check is down.",
				Optional:    true,
				Computed:    true,
			},
			"priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Administrative priority (1 - 65535).",
				Optional:    true,
				Computed:    true,
			},
			"sdwan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable egress through SD-WAN.",
				Optional:    true,
				Computed:    true,
			},
			"sdwan_zone": {
				Type:        schema.TypeList,
				Description: "Choose SD-WAN Zone.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "SD-WAN zone name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"seq_num": {
				Type: schema.TypeInt,

				Description: "Sequence number.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"src": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Source prefix for this route.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this static route.",
				Optional:    true,
				Computed:    true,
			},
			"tag": {
				Type: schema.TypeInt,

				Description: "Route tag.",
				Optional:    true,
				Computed:    true,
			},
			"virtual_wan_link": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable egress through the virtual-wan-link.",
				Optional:    true,
				Computed:    true,
			},
			"vrf": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 251),

				Description: "Virtual Routing Forwarding ID.",
				Optional:    true,
				Computed:    true,
			},
			"weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Administrative weight (0 - 255).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterStaticCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterStatic resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterStatic(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterStatic(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterStatic")
	}

	return resourceRouterStaticRead(ctx, d, meta)
}

func resourceRouterStaticUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterStatic(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterStatic(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterStatic resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterStatic")
	}

	return resourceRouterStaticRead(ctx, d, meta)
}

func resourceRouterStaticDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterStatic(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterStatic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterStaticRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterStatic(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterStatic resource: %v", err)
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

	diags := refreshObjectRouterStatic(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterStaticSdwanZone(d *schema.ResourceData, v *[]models.RouterStaticSdwanZone, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func refreshObjectRouterStatic(d *schema.ResourceData, o *models.RouterStatic, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Bfd != nil {
		v := *o.Bfd

		if err = d.Set("bfd", v); err != nil {
			return diag.Errorf("error reading bfd: %v", err)
		}
	}

	if o.Blackhole != nil {
		v := *o.Blackhole

		if err = d.Set("blackhole", v); err != nil {
			return diag.Errorf("error reading blackhole: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Device != nil {
		v := *o.Device

		if err = d.Set("device", v); err != nil {
			return diag.Errorf("error reading device: %v", err)
		}
	}

	if o.Distance != nil {
		v := *o.Distance

		if err = d.Set("distance", v); err != nil {
			return diag.Errorf("error reading distance: %v", err)
		}
	}

	if o.Dst != nil {
		v := *o.Dst
		if current, ok := d.GetOk("dst"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("dst", v); err != nil {
			return diag.Errorf("error reading dst: %v", err)
		}
	}

	if o.Dstaddr != nil {
		v := *o.Dstaddr

		if err = d.Set("dstaddr", v); err != nil {
			return diag.Errorf("error reading dstaddr: %v", err)
		}
	}

	if o.DynamicGateway != nil {
		v := *o.DynamicGateway

		if err = d.Set("dynamic_gateway", v); err != nil {
			return diag.Errorf("error reading dynamic_gateway: %v", err)
		}
	}

	if o.Gateway != nil {
		v := *o.Gateway

		if err = d.Set("gateway", v); err != nil {
			return diag.Errorf("error reading gateway: %v", err)
		}
	}

	if o.InternetService != nil {
		v := *o.InternetService

		if err = d.Set("internet_service", v); err != nil {
			return diag.Errorf("error reading internet_service: %v", err)
		}
	}

	if o.InternetServiceCustom != nil {
		v := *o.InternetServiceCustom

		if err = d.Set("internet_service_custom", v); err != nil {
			return diag.Errorf("error reading internet_service_custom: %v", err)
		}
	}

	if o.LinkMonitorExempt != nil {
		v := *o.LinkMonitorExempt

		if err = d.Set("link_monitor_exempt", v); err != nil {
			return diag.Errorf("error reading link_monitor_exempt: %v", err)
		}
	}

	if o.Priority != nil {
		v := *o.Priority

		if err = d.Set("priority", v); err != nil {
			return diag.Errorf("error reading priority: %v", err)
		}
	}

	if o.Sdwan != nil {
		v := *o.Sdwan

		if err = d.Set("sdwan", v); err != nil {
			return diag.Errorf("error reading sdwan: %v", err)
		}
	}

	if o.SdwanZone != nil {
		if err = d.Set("sdwan_zone", flattenRouterStaticSdwanZone(d, o.SdwanZone, "sdwan_zone", sort)); err != nil {
			return diag.Errorf("error reading sdwan_zone: %v", err)
		}
	}

	if o.SeqNum != nil {
		v := *o.SeqNum

		if err = d.Set("seq_num", v); err != nil {
			return diag.Errorf("error reading seq_num: %v", err)
		}
	}

	if o.Src != nil {
		v := *o.Src
		if current, ok := d.GetOk("src"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("src", v); err != nil {
			return diag.Errorf("error reading src: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Tag != nil {
		v := *o.Tag

		if err = d.Set("tag", v); err != nil {
			return diag.Errorf("error reading tag: %v", err)
		}
	}

	if o.VirtualWanLink != nil {
		v := *o.VirtualWanLink

		if err = d.Set("virtual_wan_link", v); err != nil {
			return diag.Errorf("error reading virtual_wan_link: %v", err)
		}
	}

	if o.Vrf != nil {
		v := *o.Vrf

		if err = d.Set("vrf", v); err != nil {
			return diag.Errorf("error reading vrf: %v", err)
		}
	}

	if o.Weight != nil {
		v := *o.Weight

		if err = d.Set("weight", v); err != nil {
			return diag.Errorf("error reading weight: %v", err)
		}
	}

	return nil
}

func expandRouterStaticSdwanZone(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterStaticSdwanZone, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterStaticSdwanZone

	for i := range l {
		tmp := models.RouterStaticSdwanZone{}
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

func getObjectRouterStatic(d *schema.ResourceData, sv string) (*models.RouterStatic, diag.Diagnostics) {
	obj := models.RouterStatic{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("bfd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd", sv)
				diags = append(diags, e)
			}
			obj.Bfd = &v2
		}
	}
	if v1, ok := d.GetOk("blackhole"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("blackhole", sv)
				diags = append(diags, e)
			}
			obj.Blackhole = &v2
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
	if v1, ok := d.GetOk("device"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("device", sv)
				diags = append(diags, e)
			}
			obj.Device = &v2
		}
	}
	if v1, ok := d.GetOk("distance"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distance", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Distance = &tmp
		}
	}
	if v1, ok := d.GetOk("dst"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst", sv)
				diags = append(diags, e)
			}
			obj.Dst = &v2
		}
	}
	if v1, ok := d.GetOk("dstaddr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dstaddr", sv)
				diags = append(diags, e)
			}
			obj.Dstaddr = &v2
		}
	}
	if v1, ok := d.GetOk("dynamic_gateway"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dynamic_gateway", sv)
				diags = append(diags, e)
			}
			obj.DynamicGateway = &v2
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
	if v1, ok := d.GetOk("internet_service"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.InternetService = &tmp
		}
	}
	if v1, ok := d.GetOk("internet_service_custom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service_custom", sv)
				diags = append(diags, e)
			}
			obj.InternetServiceCustom = &v2
		}
	}
	if v1, ok := d.GetOk("link_monitor_exempt"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("link_monitor_exempt", sv)
				diags = append(diags, e)
			}
			obj.LinkMonitorExempt = &v2
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
	if v1, ok := d.GetOk("sdwan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "v7.0.1") {
				e := utils.AttributeVersionWarning("sdwan", sv)
				diags = append(diags, e)
			}
			obj.Sdwan = &v2
		}
	}
	if v, ok := d.GetOk("sdwan_zone"); ok {
		if !utils.CheckVer(sv, "v7.0.1", "") {
			e := utils.AttributeVersionWarning("sdwan_zone", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterStaticSdwanZone(d, v, "sdwan_zone", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SdwanZone = t
		}
	} else if d.HasChange("sdwan_zone") {
		old, new := d.GetChange("sdwan_zone")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SdwanZone = &[]models.RouterStaticSdwanZone{}
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
	if v1, ok := d.GetOk("src"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src", sv)
				diags = append(diags, e)
			}
			obj.Src = &v2
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
	if v1, ok := d.GetOk("tag"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("tag", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Tag = &tmp
		}
	}
	if v1, ok := d.GetOk("virtual_wan_link"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("virtual_wan_link", sv)
				diags = append(diags, e)
			}
			obj.VirtualWanLink = &v2
		}
	}
	if v1, ok := d.GetOk("vrf"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vrf", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Vrf = &tmp
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
	return &obj, diags
}
