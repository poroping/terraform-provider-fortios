// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure custom services.

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

func resourceFirewallServiceCustom() *schema.Resource {
	return &schema.Resource{
		Description: "Configure custom services.",

		CreateContext: resourceFirewallServiceCustomCreate,
		ReadContext:   resourceFirewallServiceCustomRead,
		UpdateContext: resourceFirewallServiceCustomUpdate,
		DeleteContext: resourceFirewallServiceCustomDelete,

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
			"app_category": {
				Type:        schema.TypeList,
				Description: "Application category ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Application category id.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"app_service_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "app-id", "app-category"}, false),

				Description: "Application service type.",
				Optional:    true,
				Computed:    true,
			},
			"application": {
				Type:        schema.TypeList,
				Description: "Application ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Application id.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"category": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Service category.",
				Optional:    true,
				Computed:    true,
			},
			"check_reset_range": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "strict", "default"}, false),

				Description: "Configure the type of ICMP error message verification.",
				Optional:    true,
				Computed:    true,
			},
			"color": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Color of icon on the GUI.",
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
			"fabric_object": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Security Fabric global object setting.",
				Optional:    true,
				Computed:    true,
			},
			"fqdn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Fully qualified domain name.",
				Optional:    true,
				Computed:    true,
			},
			"helper": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "disable", "ftp", "tftp", "ras", "h323", "tns", "mms", "sip", "pptp", "rtsp", "dns-udp", "dns-tcp", "pmap", "rsh", "dcerpc", "mgcp"}, false),

				Description: "Helper name.",
				Optional:    true,
				Computed:    true,
			},
			"icmpcode": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "ICMP code.",
				Optional:    true,
				Computed:    true,
			},
			"icmptype": {
				Type: schema.TypeInt,

				Description: "ICMP type.",
				Optional:    true,
				Computed:    true,
			},
			"iprange": {
				Type: schema.TypeString,

				Description: "Start and end of the IP range associated with service.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Custom service name.",
				ForceNew:    true,
				Required:    true,
			},
			"protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"TCP/UDP/SCTP", "ICMP", "ICMP6", "IP", "HTTP", "FTP", "CONNECT", "SOCKS-TCP", "SOCKS-UDP", "ALL"}, false),

				Description: "Protocol type based on IANA numbers.",
				Optional:    true,
				Computed:    true,
			},
			"protocol_number": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 254),

				Description: "IP protocol number.",
				Optional:    true,
				Computed:    true,
			},
			"proxy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable web proxy service.",
				Optional:    true,
				Computed:    true,
			},
			"sctp_portrange": {
				Type: schema.TypeString,

				Description: "Multiple SCTP port ranges.",
				Optional:    true,
				Computed:    true,
			},
			"session_ttl": {
				Type: schema.TypeString,

				Description: "Session TTL (300 - 2764800, 0 = default).",
				Optional:    true,
				Computed:    true,
			},
			"tcp_halfclose_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),

				Description: "Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).",
				Optional:    true,
				Computed:    true,
			},
			"tcp_halfopen_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),

				Description: "Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).",
				Optional:    true,
				Computed:    true,
			},
			"tcp_portrange": {
				Type: schema.TypeString,

				Description: "Multiple TCP port ranges.",
				Optional:    true,
				Computed:    true,
			},
			"tcp_rst_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 300),

				Description: "Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).",
				Optional:    true,
				Computed:    true,
			},
			"tcp_timewait_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),

				Description: "Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).",
				Optional:    true,
				Computed:    true,
			},
			"udp_idle_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),

				Description: "UDP half close timeout (0 - 86400 sec, 0 = default).",
				Optional:    true,
				Computed:    true,
			},
			"udp_portrange": {
				Type: schema.TypeString,

				Description: "Multiple UDP port ranges.",
				Optional:    true,
				Computed:    true,
			},
			"visibility": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the visibility of the service on the GUI.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallServiceCustomCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallServiceCustom resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallServiceCustom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallServiceCustom(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallServiceCustom")
	}

	return resourceFirewallServiceCustomRead(ctx, d, meta)
}

func resourceFirewallServiceCustomUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallServiceCustom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallServiceCustom(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallServiceCustom resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallServiceCustom")
	}

	return resourceFirewallServiceCustomRead(ctx, d, meta)
}

func resourceFirewallServiceCustomDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallServiceCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallServiceCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallServiceCustomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallServiceCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallServiceCustom resource: %v", err)
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

	diags := refreshObjectFirewallServiceCustom(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallServiceCustomAppCategory(v *[]models.FirewallServiceCustomAppCategory, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallServiceCustomApplication(v *[]models.FirewallServiceCustomApplication, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectFirewallServiceCustom(d *schema.ResourceData, o *models.FirewallServiceCustom, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AppCategory != nil {
		if err = d.Set("app_category", flattenFirewallServiceCustomAppCategory(o.AppCategory, sort)); err != nil {
			return diag.Errorf("error reading app_category: %v", err)
		}
	}

	if o.AppServiceType != nil {
		v := *o.AppServiceType

		if err = d.Set("app_service_type", v); err != nil {
			return diag.Errorf("error reading app_service_type: %v", err)
		}
	}

	if o.Application != nil {
		if err = d.Set("application", flattenFirewallServiceCustomApplication(o.Application, sort)); err != nil {
			return diag.Errorf("error reading application: %v", err)
		}
	}

	if o.Category != nil {
		v := *o.Category

		if err = d.Set("category", v); err != nil {
			return diag.Errorf("error reading category: %v", err)
		}
	}

	if o.CheckResetRange != nil {
		v := *o.CheckResetRange

		if err = d.Set("check_reset_range", v); err != nil {
			return diag.Errorf("error reading check_reset_range: %v", err)
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

	if o.FabricObject != nil {
		v := *o.FabricObject

		if err = d.Set("fabric_object", v); err != nil {
			return diag.Errorf("error reading fabric_object: %v", err)
		}
	}

	if o.Fqdn != nil {
		v := *o.Fqdn

		if err = d.Set("fqdn", v); err != nil {
			return diag.Errorf("error reading fqdn: %v", err)
		}
	}

	if o.Helper != nil {
		v := *o.Helper

		if err = d.Set("helper", v); err != nil {
			return diag.Errorf("error reading helper: %v", err)
		}
	}

	if o.Icmpcode != nil {
		v := *o.Icmpcode

		if err = d.Set("icmpcode", v); err != nil {
			return diag.Errorf("error reading icmpcode: %v", err)
		}
	}

	if o.Icmptype != nil {
		v := *o.Icmptype

		if err = d.Set("icmptype", v); err != nil {
			return diag.Errorf("error reading icmptype: %v", err)
		}
	}

	if o.Iprange != nil {
		v := *o.Iprange

		if err = d.Set("iprange", v); err != nil {
			return diag.Errorf("error reading iprange: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.ProtocolNumber != nil {
		v := *o.ProtocolNumber

		if err = d.Set("protocol_number", v); err != nil {
			return diag.Errorf("error reading protocol_number: %v", err)
		}
	}

	if o.Proxy != nil {
		v := *o.Proxy

		if err = d.Set("proxy", v); err != nil {
			return diag.Errorf("error reading proxy: %v", err)
		}
	}

	if o.SctpPortrange != nil {
		v := *o.SctpPortrange

		if err = d.Set("sctp_portrange", v); err != nil {
			return diag.Errorf("error reading sctp_portrange: %v", err)
		}
	}

	if o.SessionTtl != nil {
		v := *o.SessionTtl

		if err = d.Set("session_ttl", v); err != nil {
			return diag.Errorf("error reading session_ttl: %v", err)
		}
	}

	if o.TcpHalfcloseTimer != nil {
		v := *o.TcpHalfcloseTimer

		if err = d.Set("tcp_halfclose_timer", v); err != nil {
			return diag.Errorf("error reading tcp_halfclose_timer: %v", err)
		}
	}

	if o.TcpHalfopenTimer != nil {
		v := *o.TcpHalfopenTimer

		if err = d.Set("tcp_halfopen_timer", v); err != nil {
			return diag.Errorf("error reading tcp_halfopen_timer: %v", err)
		}
	}

	if o.TcpPortrange != nil {
		v := *o.TcpPortrange

		if err = d.Set("tcp_portrange", v); err != nil {
			return diag.Errorf("error reading tcp_portrange: %v", err)
		}
	}

	if o.TcpRstTimer != nil {
		v := *o.TcpRstTimer

		if err = d.Set("tcp_rst_timer", v); err != nil {
			return diag.Errorf("error reading tcp_rst_timer: %v", err)
		}
	}

	if o.TcpTimewaitTimer != nil {
		v := *o.TcpTimewaitTimer

		if err = d.Set("tcp_timewait_timer", v); err != nil {
			return diag.Errorf("error reading tcp_timewait_timer: %v", err)
		}
	}

	if o.UdpIdleTimer != nil {
		v := *o.UdpIdleTimer

		if err = d.Set("udp_idle_timer", v); err != nil {
			return diag.Errorf("error reading udp_idle_timer: %v", err)
		}
	}

	if o.UdpPortrange != nil {
		v := *o.UdpPortrange

		if err = d.Set("udp_portrange", v); err != nil {
			return diag.Errorf("error reading udp_portrange: %v", err)
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

func expandFirewallServiceCustomAppCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallServiceCustomAppCategory, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallServiceCustomAppCategory

	for i := range l {
		tmp := models.FirewallServiceCustomAppCategory{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallServiceCustomApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallServiceCustomApplication, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallServiceCustomApplication

	for i := range l {
		tmp := models.FirewallServiceCustomApplication{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectFirewallServiceCustom(d *schema.ResourceData, sv string) (*models.FirewallServiceCustom, diag.Diagnostics) {
	obj := models.FirewallServiceCustom{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("app_category"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("app_category", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallServiceCustomAppCategory(d, v, "app_category", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AppCategory = t
		}
	} else if d.HasChange("app_category") {
		old, new := d.GetChange("app_category")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AppCategory = &[]models.FirewallServiceCustomAppCategory{}
		}
	}
	if v1, ok := d.GetOk("app_service_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("app_service_type", sv)
				diags = append(diags, e)
			}
			obj.AppServiceType = &v2
		}
	}
	if v, ok := d.GetOk("application"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("application", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallServiceCustomApplication(d, v, "application", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Application = t
		}
	} else if d.HasChange("application") {
		old, new := d.GetChange("application")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Application = &[]models.FirewallServiceCustomApplication{}
		}
	}
	if v1, ok := d.GetOk("category"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("category", sv)
				diags = append(diags, e)
			}
			obj.Category = &v2
		}
	}
	if v1, ok := d.GetOk("check_reset_range"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("check_reset_range", sv)
				diags = append(diags, e)
			}
			obj.CheckResetRange = &v2
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
	if v1, ok := d.GetOk("fabric_object"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.5", "") {
				e := utils.AttributeVersionWarning("fabric_object", sv)
				diags = append(diags, e)
			}
			obj.FabricObject = &v2
		}
	}
	if v1, ok := d.GetOk("fqdn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fqdn", sv)
				diags = append(diags, e)
			}
			obj.Fqdn = &v2
		}
	}
	if v1, ok := d.GetOk("helper"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("helper", sv)
				diags = append(diags, e)
			}
			obj.Helper = &v2
		}
	}
	if v1, ok := d.GetOk("icmpcode"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("icmpcode", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Icmpcode = &tmp
		}
	}
	if v1, ok := d.GetOk("icmptype"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("icmptype", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Icmptype = &tmp
		}
	}
	if v1, ok := d.GetOk("iprange"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("iprange", sv)
				diags = append(diags, e)
			}
			obj.Iprange = &v2
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
	if v1, ok := d.GetOk("protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("protocol", sv)
				diags = append(diags, e)
			}
			obj.Protocol = &v2
		}
	}
	if v1, ok := d.GetOk("protocol_number"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("protocol_number", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ProtocolNumber = &tmp
		}
	}
	if v1, ok := d.GetOk("proxy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy", sv)
				diags = append(diags, e)
			}
			obj.Proxy = &v2
		}
	}
	if v1, ok := d.GetOk("sctp_portrange"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sctp_portrange", sv)
				diags = append(diags, e)
			}
			obj.SctpPortrange = &v2
		}
	}
	if v1, ok := d.GetOk("session_ttl"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_ttl", sv)
				diags = append(diags, e)
			}
			obj.SessionTtl = &v2
		}
	}
	if v1, ok := d.GetOk("tcp_halfclose_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_halfclose_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpHalfcloseTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("tcp_halfopen_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_halfopen_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpHalfopenTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("tcp_portrange"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_portrange", sv)
				diags = append(diags, e)
			}
			obj.TcpPortrange = &v2
		}
	}
	if v1, ok := d.GetOk("tcp_rst_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("tcp_rst_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpRstTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("tcp_timewait_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_timewait_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpTimewaitTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("udp_idle_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("udp_idle_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UdpIdleTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("udp_portrange"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("udp_portrange", sv)
				diags = append(diags, e)
			}
			obj.UdpPortrange = &v2
		}
	}
	if v1, ok := d.GetOk("visibility"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("visibility", sv)
				diags = append(diags, e)
			}
			obj.Visibility = &v2
		}
	}
	return &obj, diags
}
