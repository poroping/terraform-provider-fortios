// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Link Health Monitor.

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

func resourceSystemLinkMonitor() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Link Health Monitor.",

		CreateContext: resourceSystemLinkMonitorCreate,
		ReadContext:   resourceSystemLinkMonitorRead,
		UpdateContext: resourceSystemLinkMonitorUpdate,
		DeleteContext: resourceSystemLinkMonitorDelete,

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
			"addr_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),

				Description: "Address mode (IPv4 or IPv6).",
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Type: schema.TypeInt,

				Description: "Traffic class ID.",
				Optional:    true,
				Computed:    true,
			},
			"diffservcode": {
				Type: schema.TypeString,

				Description: "Differentiated services code point (DSCP) in the IP header of the probe packet.",
				Optional:    true,
				Computed:    true,
			},
			"fail_weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Threshold weight to trigger link failure alert.",
				Optional:    true,
				Computed:    true,
			},
			"failtime": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "Number of retry attempts before the server is considered down (1 - 10, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"gateway_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Gateway IP address used to probe the server.",
				Optional:    true,
				Computed:    true,
			},
			"gateway_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Gateway IPv6 address used to probe the server.",
				Optional:         true,
				Computed:         true,
			},
			"ha_priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 50),

				Description: "HA election priority (1 - 50).",
				Optional:    true,
				Computed:    true,
			},
			"http_agent": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),

				Description: "String in the http-agent field in the HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"http_get": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),

				Description: "If you are monitoring an HTML server you can send an HTTP-GET request with a custom string. Use this option to define the string.",
				Optional:    true,
				Computed:    true,
			},
			"http_match": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),

				Description: "String that you expect to see in the HTTP-GET requests of the traffic to be monitored.",
				Optional:    true,
				Computed:    true,
			},
			"interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(500, 3600000),

				Description: "Detection interval in milliseconds (500 - 3600 * 1000 msec, default = 500).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Link monitor name.",
				ForceNew:    true,
				Required:    true,
			},
			"packet_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(64, 1024),

				Description: "Packet size of a TWAMP test session.",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "TWAMP controller password in authentication mode.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port number of the traffic to be used to monitor the server.",
				Optional:    true,
				Computed:    true,
			},
			"probe_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 30),

				Description: "Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"probe_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(500, 5000),

				Description: "Time to wait before a probe packet is considered lost (500 - 5000 msec, default = 500).",
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Protocols used to monitor the server.",
				Optional:         true,
				Computed:         true,
			},
			"recoverytime": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "Number of successful responses received before server is considered recovered (1 - 10, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"route": {
				Type:        schema.TypeList,
				Description: "Subnet to monitor.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IP and netmask (x.x.x.x/y).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"security_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "authentication"}, false),

				Description: "Twamp controller security mode.",
				Optional:    true,
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeList,
				Description: "IP address of the server(s) to be monitored.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Server address.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"server_config": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "individual"}, false),

				Description: "Mode of server configuration.",
				Optional:    true,
				Computed:    true,
			},
			"server_list": {
				Type:        schema.TypeList,
				Description: "Servers for link-monitor to monitor.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "IP address of the server to be monitored.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),

							Description: "Server ID.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Port number of the traffic to be used to monitor the server.",
							Optional:    true,
							Computed:    true,
						},
						"protocol": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Protocols used to monitor the server.",
							Optional:         true,
							Computed:         true,
						},
						"weight": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Weight of the monitor to this dst (0 - 255).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"service_detection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Only use monitor to read quality values. If enabled, static routes and cascade interfaces will not be updated.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP address used in packet to the server.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Source IPv6 address used in packet to the server.",
				Optional:         true,
				Computed:         true,
			},
			"srcintf": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Interface that receives the traffic to be monitored.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this link monitor.",
				Optional:    true,
				Computed:    true,
			},
			"update_cascade_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable update cascade interface.",
				Optional:    true,
				Computed:    true,
			},
			"update_policy_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable updating the policy route.",
				Optional:    true,
				Computed:    true,
			},
			"update_static_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable updating the static route.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemLinkMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemLinkMonitor resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemLinkMonitor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemLinkMonitor(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemLinkMonitor")
	}

	return resourceSystemLinkMonitorRead(ctx, d, meta)
}

func resourceSystemLinkMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemLinkMonitor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemLinkMonitor(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemLinkMonitor resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemLinkMonitor")
	}

	return resourceSystemLinkMonitorRead(ctx, d, meta)
}

func resourceSystemLinkMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemLinkMonitor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemLinkMonitor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLinkMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemLinkMonitor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemLinkMonitor resource: %v", err)
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

	diags := refreshObjectSystemLinkMonitor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemLinkMonitorRoute(d *schema.ResourceData, v *[]models.SystemLinkMonitorRoute, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Subnet; tmp != nil {
				v["subnet"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "subnet")
	}

	return flat
}

func flattenSystemLinkMonitorServer(d *schema.ResourceData, v *[]models.SystemLinkMonitorServer, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Address; tmp != nil {
				v["address"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "address")
	}

	return flat
}

func flattenSystemLinkMonitorServerList(d *schema.ResourceData, v *[]models.SystemLinkMonitorServerList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dst; tmp != nil {
				v["dst"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.Weight; tmp != nil {
				v["weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemLinkMonitor(d *schema.ResourceData, o *models.SystemLinkMonitor, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddrMode != nil {
		v := *o.AddrMode

		if err = d.Set("addr_mode", v); err != nil {
			return diag.Errorf("error reading addr_mode: %v", err)
		}
	}

	if o.ClassId != nil {
		v := *o.ClassId

		if err = d.Set("class_id", v); err != nil {
			return diag.Errorf("error reading class_id: %v", err)
		}
	}

	if o.Diffservcode != nil {
		v := *o.Diffservcode

		if err = d.Set("diffservcode", v); err != nil {
			return diag.Errorf("error reading diffservcode: %v", err)
		}
	}

	if o.FailWeight != nil {
		v := *o.FailWeight

		if err = d.Set("fail_weight", v); err != nil {
			return diag.Errorf("error reading fail_weight: %v", err)
		}
	}

	if o.Failtime != nil {
		v := *o.Failtime

		if err = d.Set("failtime", v); err != nil {
			return diag.Errorf("error reading failtime: %v", err)
		}
	}

	if o.GatewayIp != nil {
		v := *o.GatewayIp

		if err = d.Set("gateway_ip", v); err != nil {
			return diag.Errorf("error reading gateway_ip: %v", err)
		}
	}

	if o.GatewayIp6 != nil {
		v := *o.GatewayIp6

		if err = d.Set("gateway_ip6", v); err != nil {
			return diag.Errorf("error reading gateway_ip6: %v", err)
		}
	}

	if o.HaPriority != nil {
		v := *o.HaPriority

		if err = d.Set("ha_priority", v); err != nil {
			return diag.Errorf("error reading ha_priority: %v", err)
		}
	}

	if o.HttpAgent != nil {
		v := *o.HttpAgent

		if err = d.Set("http_agent", v); err != nil {
			return diag.Errorf("error reading http_agent: %v", err)
		}
	}

	if o.HttpGet != nil {
		v := *o.HttpGet

		if err = d.Set("http_get", v); err != nil {
			return diag.Errorf("error reading http_get: %v", err)
		}
	}

	if o.HttpMatch != nil {
		v := *o.HttpMatch

		if err = d.Set("http_match", v); err != nil {
			return diag.Errorf("error reading http_match: %v", err)
		}
	}

	if o.Interval != nil {
		v := *o.Interval

		if err = d.Set("interval", v); err != nil {
			return diag.Errorf("error reading interval: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PacketSize != nil {
		v := *o.PacketSize

		if err = d.Set("packet_size", v); err != nil {
			return diag.Errorf("error reading packet_size: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.ProbeCount != nil {
		v := *o.ProbeCount

		if err = d.Set("probe_count", v); err != nil {
			return diag.Errorf("error reading probe_count: %v", err)
		}
	}

	if o.ProbeTimeout != nil {
		v := *o.ProbeTimeout

		if err = d.Set("probe_timeout", v); err != nil {
			return diag.Errorf("error reading probe_timeout: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.Recoverytime != nil {
		v := *o.Recoverytime

		if err = d.Set("recoverytime", v); err != nil {
			return diag.Errorf("error reading recoverytime: %v", err)
		}
	}

	if o.Route != nil {
		if err = d.Set("route", flattenSystemLinkMonitorRoute(d, o.Route, "route", sort)); err != nil {
			return diag.Errorf("error reading route: %v", err)
		}
	}

	if o.SecurityMode != nil {
		v := *o.SecurityMode

		if err = d.Set("security_mode", v); err != nil {
			return diag.Errorf("error reading security_mode: %v", err)
		}
	}

	if o.Server != nil {
		if err = d.Set("server", flattenSystemLinkMonitorServer(d, o.Server, "server", sort)); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.ServerConfig != nil {
		v := *o.ServerConfig

		if err = d.Set("server_config", v); err != nil {
			return diag.Errorf("error reading server_config: %v", err)
		}
	}

	if o.ServerList != nil {
		if err = d.Set("server_list", flattenSystemLinkMonitorServerList(d, o.ServerList, "server_list", sort)); err != nil {
			return diag.Errorf("error reading server_list: %v", err)
		}
	}

	if o.ServiceDetection != nil {
		v := *o.ServiceDetection

		if err = d.Set("service_detection", v); err != nil {
			return diag.Errorf("error reading service_detection: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.SourceIp6 != nil {
		v := *o.SourceIp6

		if err = d.Set("source_ip6", v); err != nil {
			return diag.Errorf("error reading source_ip6: %v", err)
		}
	}

	if o.Srcintf != nil {
		v := *o.Srcintf

		if err = d.Set("srcintf", v); err != nil {
			return diag.Errorf("error reading srcintf: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.UpdateCascadeInterface != nil {
		v := *o.UpdateCascadeInterface

		if err = d.Set("update_cascade_interface", v); err != nil {
			return diag.Errorf("error reading update_cascade_interface: %v", err)
		}
	}

	if o.UpdatePolicyRoute != nil {
		v := *o.UpdatePolicyRoute

		if err = d.Set("update_policy_route", v); err != nil {
			return diag.Errorf("error reading update_policy_route: %v", err)
		}
	}

	if o.UpdateStaticRoute != nil {
		v := *o.UpdateStaticRoute

		if err = d.Set("update_static_route", v); err != nil {
			return diag.Errorf("error reading update_static_route: %v", err)
		}
	}

	return nil
}

func expandSystemLinkMonitorRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemLinkMonitorRoute, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemLinkMonitorRoute

	for i := range l {
		tmp := models.SystemLinkMonitorRoute{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.subnet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Subnet = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemLinkMonitorServer(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemLinkMonitorServer, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemLinkMonitorServer

	for i := range l {
		tmp := models.SystemLinkMonitorServer{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Address = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemLinkMonitorServerList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemLinkMonitorServerList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemLinkMonitorServerList

	for i := range l {
		tmp := models.SystemLinkMonitorServerList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dst", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dst = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Port = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Protocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Weight = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemLinkMonitor(d *schema.ResourceData, sv string) (*models.SystemLinkMonitor, diag.Diagnostics) {
	obj := models.SystemLinkMonitor{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("addr_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("addr_mode", sv)
				diags = append(diags, e)
			}
			obj.AddrMode = &v2
		}
	}
	if v1, ok := d.GetOk("class_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("class_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ClassId = &tmp
		}
	}
	if v1, ok := d.GetOk("diffservcode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("diffservcode", sv)
				diags = append(diags, e)
			}
			obj.Diffservcode = &v2
		}
	}
	if v1, ok := d.GetOk("fail_weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("fail_weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FailWeight = &tmp
		}
	}
	if v1, ok := d.GetOk("failtime"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("failtime", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Failtime = &tmp
		}
	}
	if v1, ok := d.GetOk("gateway_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gateway_ip", sv)
				diags = append(diags, e)
			}
			obj.GatewayIp = &v2
		}
	}
	if v1, ok := d.GetOk("gateway_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gateway_ip6", sv)
				diags = append(diags, e)
			}
			obj.GatewayIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("ha_priority"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ha_priority", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HaPriority = &tmp
		}
	}
	if v1, ok := d.GetOk("http_agent"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_agent", sv)
				diags = append(diags, e)
			}
			obj.HttpAgent = &v2
		}
	}
	if v1, ok := d.GetOk("http_get"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_get", sv)
				diags = append(diags, e)
			}
			obj.HttpGet = &v2
		}
	}
	if v1, ok := d.GetOk("http_match"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_match", sv)
				diags = append(diags, e)
			}
			obj.HttpMatch = &v2
		}
	}
	if v1, ok := d.GetOk("interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Interval = &tmp
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
	if v1, ok := d.GetOk("packet_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("packet_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PacketSize = &tmp
		}
	}
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port = &tmp
		}
	}
	if v1, ok := d.GetOk("probe_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("probe_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ProbeCount = &tmp
		}
	}
	if v1, ok := d.GetOk("probe_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("probe_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ProbeTimeout = &tmp
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
	if v1, ok := d.GetOk("recoverytime"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("recoverytime", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Recoverytime = &tmp
		}
	}
	if v, ok := d.GetOk("route"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("route", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemLinkMonitorRoute(d, v, "route", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Route = t
		}
	} else if d.HasChange("route") {
		old, new := d.GetChange("route")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Route = &[]models.SystemLinkMonitorRoute{}
		}
	}
	if v1, ok := d.GetOk("security_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_mode", sv)
				diags = append(diags, e)
			}
			obj.SecurityMode = &v2
		}
	}
	if v, ok := d.GetOk("server"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("server", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemLinkMonitorServer(d, v, "server", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Server = t
		}
	} else if d.HasChange("server") {
		old, new := d.GetChange("server")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Server = &[]models.SystemLinkMonitorServer{}
		}
	}
	if v1, ok := d.GetOk("server_config"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("server_config", sv)
				diags = append(diags, e)
			}
			obj.ServerConfig = &v2
		}
	}
	if v, ok := d.GetOk("server_list"); ok {
		if !utils.CheckVer(sv, "v7.0.1", "") {
			e := utils.AttributeVersionWarning("server_list", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemLinkMonitorServerList(d, v, "server_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ServerList = t
		}
	} else if d.HasChange("server_list") {
		old, new := d.GetChange("server_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ServerList = &[]models.SystemLinkMonitorServerList{}
		}
	}
	if v1, ok := d.GetOk("service_detection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("service_detection", sv)
				diags = append(diags, e)
			}
			obj.ServiceDetection = &v2
		}
	}
	if v1, ok := d.GetOk("source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip", sv)
				diags = append(diags, e)
			}
			obj.SourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("source_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip6", sv)
				diags = append(diags, e)
			}
			obj.SourceIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("srcintf"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("srcintf", sv)
				diags = append(diags, e)
			}
			obj.Srcintf = &v2
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
	if v1, ok := d.GetOk("update_cascade_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_cascade_interface", sv)
				diags = append(diags, e)
			}
			obj.UpdateCascadeInterface = &v2
		}
	}
	if v1, ok := d.GetOk("update_policy_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("update_policy_route", sv)
				diags = append(diags, e)
			}
			obj.UpdatePolicyRoute = &v2
		}
	}
	if v1, ok := d.GetOk("update_static_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_static_route", sv)
				diags = append(diags, e)
			}
			obj.UpdateStaticRoute = &v2
		}
	}
	return &obj, diags
}
