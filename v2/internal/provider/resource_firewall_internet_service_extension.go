// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Internet Services Extension.

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

func resourceFirewallInternetServiceExtension() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Internet Services Extension.",

		CreateContext: resourceFirewallInternetServiceExtensionCreate,
		ReadContext:   resourceFirewallInternetServiceExtensionRead,
		UpdateContext: resourceFirewallInternetServiceExtensionUpdate,
		DeleteContext: resourceFirewallInternetServiceExtensionDelete,

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
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"disable_entry": {
				Type:        schema.TypeList,
				Description: "Disable entries in the Internet Service database.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Disable entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"ip_range": {
							Type:        schema.TypeList,
							Description: "IP ranges in the disable entry.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_ip": {
										Type:         schema.TypeString,
										ValidateFunc: validation.IsIPv4Address,

										Description: "End IP address.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Disable entry range ID.",
										Optional:    true,
										Computed:    true,
									},
									"start_ip": {
										Type:         schema.TypeString,
										ValidateFunc: validation.IsIPv4Address,

										Description: "Start IP address.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"port_range": {
							Type:        schema.TypeList,
							Description: "Port ranges in the disable entry.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Ending TCP/UDP/SCTP destination port (1 to 65535).",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Custom entry port range ID.",
										Optional:    true,
										Computed:    true,
									},
									"start_port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Starting TCP/UDP/SCTP destination port (1 to 65535).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"protocol": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Integer value for the protocol type as defined by IANA (0 - 255).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"entry": {
				Type:        schema.TypeList,
				Description: "Entries added to the Internet Service extension database.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst": {
							Type:        schema.TypeList,
							Description: "Destination address or address group name.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Select the destination address or address group object from available options.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Entry ID(1-255).",
							Optional:    true,
							Computed:    true,
						},
						"port_range": {
							Type:        schema.TypeList,
							Description: "Port ranges in the custom entry.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Custom entry port range ID.",
										Optional:    true,
										Computed:    true,
									},
									"start_port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"protocol": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Integer value for the protocol type as defined by IANA (0 - 255).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Internet Service ID in the Internet Service database.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallInternetServiceExtensionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallInternetServiceExtension resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallInternetServiceExtension(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallInternetServiceExtension(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallInternetServiceExtension")
	}

	return resourceFirewallInternetServiceExtensionRead(ctx, d, meta)
}

func resourceFirewallInternetServiceExtensionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallInternetServiceExtension(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallInternetServiceExtension(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallInternetServiceExtension resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallInternetServiceExtension")
	}

	return resourceFirewallInternetServiceExtensionRead(ctx, d, meta)
}

func resourceFirewallInternetServiceExtensionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallInternetServiceExtension(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallInternetServiceExtension resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceExtensionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallInternetServiceExtension(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallInternetServiceExtension resource: %v", err)
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

	diags := refreshObjectFirewallInternetServiceExtension(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallInternetServiceExtensionDisableEntry(d *schema.ResourceData, v *[]models.FirewallInternetServiceExtensionDisableEntry, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.IpRange; tmp != nil {
				v["ip_range"] = flattenFirewallInternetServiceExtensionDisableEntryIpRange(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "ip_range"), sort)
			}

			if tmp := cfg.PortRange; tmp != nil {
				v["port_range"] = flattenFirewallInternetServiceExtensionDisableEntryPortRange(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "port_range"), sort)
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallInternetServiceExtensionDisableEntryIpRange(d *schema.ResourceData, v *[]models.FirewallInternetServiceExtensionDisableEntryIpRange, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.EndIp; tmp != nil {
				v["end_ip"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.StartIp; tmp != nil {
				v["start_ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallInternetServiceExtensionDisableEntryPortRange(d *schema.ResourceData, v *[]models.FirewallInternetServiceExtensionDisableEntryPortRange, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.EndPort; tmp != nil {
				v["end_port"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.StartPort; tmp != nil {
				v["start_port"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallInternetServiceExtensionEntry(d *schema.ResourceData, v *[]models.FirewallInternetServiceExtensionEntry, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dst; tmp != nil {
				v["dst"] = flattenFirewallInternetServiceExtensionEntryDst(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "dst"), sort)
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.PortRange; tmp != nil {
				v["port_range"] = flattenFirewallInternetServiceExtensionEntryPortRange(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "port_range"), sort)
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallInternetServiceExtensionEntryDst(d *schema.ResourceData, v *[]models.FirewallInternetServiceExtensionEntryDst, prefix string, sort bool) interface{} {
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

func flattenFirewallInternetServiceExtensionEntryPortRange(d *schema.ResourceData, v *[]models.FirewallInternetServiceExtensionEntryPortRange, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.EndPort; tmp != nil {
				v["end_port"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.StartPort; tmp != nil {
				v["start_port"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectFirewallInternetServiceExtension(d *schema.ResourceData, o *models.FirewallInternetServiceExtension, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.DisableEntry != nil {
		if err = d.Set("disable_entry", flattenFirewallInternetServiceExtensionDisableEntry(d, o.DisableEntry, "disable_entry", sort)); err != nil {
			return diag.Errorf("error reading disable_entry: %v", err)
		}
	}

	if o.Entry != nil {
		if err = d.Set("entry", flattenFirewallInternetServiceExtensionEntry(d, o.Entry, "entry", sort)); err != nil {
			return diag.Errorf("error reading entry: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	return nil
}

func expandFirewallInternetServiceExtensionDisableEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallInternetServiceExtensionDisableEntry, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallInternetServiceExtensionDisableEntry

	for i := range l {
		tmp := models.FirewallInternetServiceExtensionDisableEntry{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip_range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallInternetServiceExtensionDisableEntryIpRange(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallInternetServiceExtensionDisableEntryIpRange
			// 	}
			tmp.IpRange = v2

		}

		pre_append = fmt.Sprintf("%s.%d.port_range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallInternetServiceExtensionDisableEntryPortRange(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallInternetServiceExtensionDisableEntryPortRange
			// 	}
			tmp.PortRange = v2

		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Protocol = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallInternetServiceExtensionDisableEntryIpRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallInternetServiceExtensionDisableEntryIpRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallInternetServiceExtensionDisableEntryIpRange

	for i := range l {
		tmp := models.FirewallInternetServiceExtensionDisableEntryIpRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EndIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StartIp = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallInternetServiceExtensionDisableEntryPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallInternetServiceExtensionDisableEntryPortRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallInternetServiceExtensionDisableEntryPortRange

	for i := range l {
		tmp := models.FirewallInternetServiceExtensionDisableEntryPortRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.EndPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.StartPort = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallInternetServiceExtensionEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallInternetServiceExtensionEntry, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallInternetServiceExtensionEntry

	for i := range l {
		tmp := models.FirewallInternetServiceExtensionEntry{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dst", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallInternetServiceExtensionEntryDst(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallInternetServiceExtensionEntryDst
			// 	}
			tmp.Dst = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallInternetServiceExtensionEntryPortRange(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallInternetServiceExtensionEntryPortRange
			// 	}
			tmp.PortRange = v2

		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Protocol = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallInternetServiceExtensionEntryDst(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallInternetServiceExtensionEntryDst, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallInternetServiceExtensionEntryDst

	for i := range l {
		tmp := models.FirewallInternetServiceExtensionEntryDst{}
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

func expandFirewallInternetServiceExtensionEntryPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallInternetServiceExtensionEntryPortRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallInternetServiceExtensionEntryPortRange

	for i := range l {
		tmp := models.FirewallInternetServiceExtensionEntryPortRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.EndPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.StartPort = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectFirewallInternetServiceExtension(d *schema.ResourceData, sv string) (*models.FirewallInternetServiceExtension, diag.Diagnostics) {
	obj := models.FirewallInternetServiceExtension{}
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
	if v, ok := d.GetOk("disable_entry"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("disable_entry", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallInternetServiceExtensionDisableEntry(d, v, "disable_entry", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DisableEntry = t
		}
	} else if d.HasChange("disable_entry") {
		old, new := d.GetChange("disable_entry")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DisableEntry = &[]models.FirewallInternetServiceExtensionDisableEntry{}
		}
	}
	if v, ok := d.GetOk("entry"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("entry", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallInternetServiceExtensionEntry(d, v, "entry", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Entry = t
		}
	} else if d.HasChange("entry") {
		old, new := d.GetChange("entry")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Entry = &[]models.FirewallInternetServiceExtensionEntry{}
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	return &obj, diags
}
