// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 to IPv6 virtual IPs.

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

func resourceFirewallVip46() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 to IPv6 virtual IPs.",

		CreateContext: resourceFirewallVip46Create,
		ReadContext:   resourceFirewallVip46Read,
		UpdateContext: resourceFirewallVip46Update,
		DeleteContext: resourceFirewallVip46Delete,

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
			"arp_reply": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable ARP reply.",
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
			"extip": {
				Type: schema.TypeString,

				Description: "Start-external-IP [-end-external-IP].",
				Optional:    true,
				Computed:    true,
			},
			"extport": {
				Type: schema.TypeString,

				Description: "External service port.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Custom defined id.",
				Optional:    true,
				Computed:    true,
			},
			"ldb_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"static", "round-robin", "weighted", "least-session", "least-rtt", "first-alive"}, false),

				Description: "Load balance method.",
				Optional:    true,
				Computed:    true,
			},
			"mappedip": {
				Type: schema.TypeString,

				Description: "Start-mapped-IPv6-address [-end mapped-IPv6-address].",
				Optional:    true,
				Computed:    true,
			},
			"mappedport": {
				Type: schema.TypeString,

				Description: "Mapped service port.",
				Optional:    true,
				Computed:    true,
			},
			"monitor": {
				Type:        schema.TypeList,
				Description: "Health monitors.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Health monitor name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "VIP46 name.",
				ForceNew:    true,
				Required:    true,
			},
			"portforward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable port forwarding.",
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tcp", "udp"}, false),

				Description: "Mapped port protocol.",
				Optional:    true,
				Computed:    true,
			},
			"realservers": {
				Type:        schema.TypeList,
				Description: "Real servers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_ip": {
							Type: schema.TypeString,

							Description: "Restrict server to a client IP in this range.",
							Optional:    true,
							Computed:    true,
						},
						"healthcheck": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable", "vip"}, false),

							Description: "Per server health check.",
							Optional:    true,
							Computed:    true,
						},
						"holddown_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(30, 65535),

							Description: "Hold down interval.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Real server ID.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "Mapped server IPv6.",
							Optional:         true,
							Computed:         true,
						},
						"max_connections": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 2147483647),

							Description: "Maximum number of connections allowed to server.",
							Optional:    true,
							Computed:    true,
						},
						"monitor": {
							Type:        schema.TypeList,
							Description: "Health monitors.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Health monitor name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Mapped server port.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"active", "standby", "disable"}, false),

							Description: "Server administrative status.",
							Optional:    true,
							Computed:    true,
						},
						"weight": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "weight",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"server_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"http", "tcp", "udp", "ip"}, false),

				Description: "Server type.",
				Optional:    true,
				Computed:    true,
			},
			"src_filter": {
				Type:        schema.TypeList,
				Description: "Source IP filter (x.x.x.x/x).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Src-filter range.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"srcintf_filter": {
				Type:        schema.TypeList,
				Description: "Interfaces to which the VIP46 applies. Separate the names with spaces.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"static-nat", "server-load-balance"}, false),

				Description: "VIP type: static NAT or server load balance.",
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Type: schema.TypeString,

				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallVip46Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallVip46 resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallVip46(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallVip46(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallVip46")
	}

	return resourceFirewallVip46Read(ctx, d, meta)
}

func resourceFirewallVip46Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallVip46(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallVip46(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallVip46 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallVip46")
	}

	return resourceFirewallVip46Read(ctx, d, meta)
}

func resourceFirewallVip46Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallVip46(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallVip46 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVip46Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallVip46(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallVip46 resource: %v", err)
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

	diags := refreshObjectFirewallVip46(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallVip46Monitor(v *[]models.FirewallVip46Monitor, sort bool) interface{} {
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

func flattenFirewallVip46Realservers(v *[]models.FirewallVip46Realservers, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.ClientIp; tmp != nil {
				v["client_ip"] = *tmp
			}

			if tmp := cfg.Healthcheck; tmp != nil {
				v["healthcheck"] = *tmp
			}

			if tmp := cfg.HolddownInterval; tmp != nil {
				v["holddown_interval"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.MaxConnections; tmp != nil {
				v["max_connections"] = *tmp
			}

			if tmp := cfg.Monitor; tmp != nil {
				v["monitor"] = flattenFirewallVip46RealserversMonitor(tmp, sort)
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
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

func flattenFirewallVip46RealserversMonitor(v *[]models.FirewallVip46RealserversMonitor, sort bool) interface{} {
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

func flattenFirewallVip46SrcFilter(v *[]models.FirewallVip46SrcFilter, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Range; tmp != nil {
				v["range"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "range")
	}

	return flat
}

func flattenFirewallVip46SrcintfFilter(v *[]models.FirewallVip46SrcintfFilter, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.InterfaceName; tmp != nil {
				v["interface_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "interface_name")
	}

	return flat
}

func refreshObjectFirewallVip46(d *schema.ResourceData, o *models.FirewallVip46, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ArpReply != nil {
		v := *o.ArpReply

		if err = d.Set("arp_reply", v); err != nil {
			return diag.Errorf("error reading arp_reply: %v", err)
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

	if o.Extip != nil {
		v := *o.Extip

		if err = d.Set("extip", v); err != nil {
			return diag.Errorf("error reading extip: %v", err)
		}
	}

	if o.Extport != nil {
		v := *o.Extport

		if err = d.Set("extport", v); err != nil {
			return diag.Errorf("error reading extport: %v", err)
		}
	}

	if o.Fosid != nil {
		v := *o.Fosid

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.LdbMethod != nil {
		v := *o.LdbMethod

		if err = d.Set("ldb_method", v); err != nil {
			return diag.Errorf("error reading ldb_method: %v", err)
		}
	}

	if o.Mappedip != nil {
		v := *o.Mappedip

		if err = d.Set("mappedip", v); err != nil {
			return diag.Errorf("error reading mappedip: %v", err)
		}
	}

	if o.Mappedport != nil {
		v := *o.Mappedport

		if err = d.Set("mappedport", v); err != nil {
			return diag.Errorf("error reading mappedport: %v", err)
		}
	}

	if o.Monitor != nil {
		if err = d.Set("monitor", flattenFirewallVip46Monitor(o.Monitor, sort)); err != nil {
			return diag.Errorf("error reading monitor: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Portforward != nil {
		v := *o.Portforward

		if err = d.Set("portforward", v); err != nil {
			return diag.Errorf("error reading portforward: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.Realservers != nil {
		if err = d.Set("realservers", flattenFirewallVip46Realservers(o.Realservers, sort)); err != nil {
			return diag.Errorf("error reading realservers: %v", err)
		}
	}

	if o.ServerType != nil {
		v := *o.ServerType

		if err = d.Set("server_type", v); err != nil {
			return diag.Errorf("error reading server_type: %v", err)
		}
	}

	if o.SrcFilter != nil {
		if err = d.Set("src_filter", flattenFirewallVip46SrcFilter(o.SrcFilter, sort)); err != nil {
			return diag.Errorf("error reading src_filter: %v", err)
		}
	}

	if o.SrcintfFilter != nil {
		if err = d.Set("srcintf_filter", flattenFirewallVip46SrcintfFilter(o.SrcintfFilter, sort)); err != nil {
			return diag.Errorf("error reading srcintf_filter: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.Uuid != nil {
		v := *o.Uuid

		if err = d.Set("uuid", v); err != nil {
			return diag.Errorf("error reading uuid: %v", err)
		}
	}

	return nil
}

func expandFirewallVip46Monitor(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallVip46Monitor, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallVip46Monitor

	for i := range l {
		tmp := models.FirewallVip46Monitor{}
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

func expandFirewallVip46Realservers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallVip46Realservers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallVip46Realservers

	for i := range l {
		tmp := models.FirewallVip46Realservers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.client_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.healthcheck", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Healthcheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.holddown_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HolddownInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_connections", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MaxConnections = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.monitor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallVip46RealserversMonitor(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallVip46RealserversMonitor
			// 	}
			tmp.Monitor = v2

		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Port = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Weight = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallVip46RealserversMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallVip46RealserversMonitor, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallVip46RealserversMonitor

	for i := range l {
		tmp := models.FirewallVip46RealserversMonitor{}
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

func expandFirewallVip46SrcFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallVip46SrcFilter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallVip46SrcFilter

	for i := range l {
		tmp := models.FirewallVip46SrcFilter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Range = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallVip46SrcintfFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallVip46SrcintfFilter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallVip46SrcintfFilter

	for i := range l {
		tmp := models.FirewallVip46SrcintfFilter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterfaceName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectFirewallVip46(d *schema.ResourceData, sv string) (*models.FirewallVip46, diag.Diagnostics) {
	obj := models.FirewallVip46{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("arp_reply"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("arp_reply", sv)
				diags = append(diags, e)
			}
			obj.ArpReply = &v2
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
	if v1, ok := d.GetOk("extip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extip", sv)
				diags = append(diags, e)
			}
			obj.Extip = &v2
		}
	}
	if v1, ok := d.GetOk("extport"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extport", sv)
				diags = append(diags, e)
			}
			obj.Extport = &v2
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Fosid = &tmp
		}
	}
	if v1, ok := d.GetOk("ldb_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldb_method", sv)
				diags = append(diags, e)
			}
			obj.LdbMethod = &v2
		}
	}
	if v1, ok := d.GetOk("mappedip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mappedip", sv)
				diags = append(diags, e)
			}
			obj.Mappedip = &v2
		}
	}
	if v1, ok := d.GetOk("mappedport"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mappedport", sv)
				diags = append(diags, e)
			}
			obj.Mappedport = &v2
		}
	}
	if v, ok := d.GetOk("monitor"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("monitor", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallVip46Monitor(d, v, "monitor", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Monitor = t
		}
	} else if d.HasChange("monitor") {
		old, new := d.GetChange("monitor")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Monitor = &[]models.FirewallVip46Monitor{}
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
	if v1, ok := d.GetOk("portforward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("portforward", sv)
				diags = append(diags, e)
			}
			obj.Portforward = &v2
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
	if v, ok := d.GetOk("realservers"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("realservers", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallVip46Realservers(d, v, "realservers", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Realservers = t
		}
	} else if d.HasChange("realservers") {
		old, new := d.GetChange("realservers")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Realservers = &[]models.FirewallVip46Realservers{}
		}
	}
	if v1, ok := d.GetOk("server_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_type", sv)
				diags = append(diags, e)
			}
			obj.ServerType = &v2
		}
	}
	if v, ok := d.GetOk("src_filter"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("src_filter", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallVip46SrcFilter(d, v, "src_filter", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SrcFilter = t
		}
	} else if d.HasChange("src_filter") {
		old, new := d.GetChange("src_filter")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SrcFilter = &[]models.FirewallVip46SrcFilter{}
		}
	}
	if v, ok := d.GetOk("srcintf_filter"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("srcintf_filter", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallVip46SrcintfFilter(d, v, "srcintf_filter", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SrcintfFilter = t
		}
	} else if d.HasChange("srcintf_filter") {
		old, new := d.GetChange("srcintf_filter")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SrcintfFilter = &[]models.FirewallVip46SrcintfFilter{}
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
	if v1, ok := d.GetOk("uuid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uuid", sv)
				diags = append(diags, e)
			}
			obj.Uuid = &v2
		}
	}
	return &obj, diags
}
