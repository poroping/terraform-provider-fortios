// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Internet Service definition.

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

func resourceFirewallInternetServiceDefinition() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Internet Service definition.",

		CreateContext: resourceFirewallInternetServiceDefinitionCreate,
		ReadContext:   resourceFirewallInternetServiceDefinitionRead,
		UpdateContext: resourceFirewallInternetServiceDefinitionUpdate,
		DeleteContext: resourceFirewallInternetServiceDefinitionDelete,

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
			"entry": {
				Type:        schema.TypeList,
				Description: "Protocol and port information in an Internet Service entry.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category_id": {
							Type: schema.TypeInt,

							Description: "Internet Service category ID.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Internet Service name.",
							Optional:    true,
							Computed:    true,
						},
						"port_range": {
							Type:        schema.TypeList,
							Description: "Port ranges in the definition entry.",
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
						"seq_num": {
							Type: schema.TypeInt,

							Description: "Entry sequence number.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Internet Service application list ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallInternetServiceDefinitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallInternetServiceDefinition resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallInternetServiceDefinition(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallInternetServiceDefinition(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallInternetServiceDefinition")
	}

	return resourceFirewallInternetServiceDefinitionRead(ctx, d, meta)
}

func resourceFirewallInternetServiceDefinitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallInternetServiceDefinition(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallInternetServiceDefinition(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallInternetServiceDefinition resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallInternetServiceDefinition")
	}

	return resourceFirewallInternetServiceDefinitionRead(ctx, d, meta)
}

func resourceFirewallInternetServiceDefinitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallInternetServiceDefinition(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallInternetServiceDefinition resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceDefinitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallInternetServiceDefinition(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallInternetServiceDefinition resource: %v", err)
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

	diags := refreshObjectFirewallInternetServiceDefinition(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallInternetServiceDefinitionEntry(d *schema.ResourceData, v *[]models.FirewallInternetServiceDefinitionEntry, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.CategoryId; tmp != nil {
				v["category_id"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.PortRange; tmp != nil {
				v["port_range"] = flattenFirewallInternetServiceDefinitionEntryPortRange(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "port_range"), sort)
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.SeqNum; tmp != nil {
				v["seq_num"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "seq_num")
	}

	return flat
}

func flattenFirewallInternetServiceDefinitionEntryPortRange(d *schema.ResourceData, v *[]models.FirewallInternetServiceDefinitionEntryPortRange, prefix string, sort bool) interface{} {
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

func refreshObjectFirewallInternetServiceDefinition(d *schema.ResourceData, o *models.FirewallInternetServiceDefinition, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Entry != nil {
		if err = d.Set("entry", flattenFirewallInternetServiceDefinitionEntry(d, o.Entry, "entry", sort)); err != nil {
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

func expandFirewallInternetServiceDefinitionEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallInternetServiceDefinitionEntry, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallInternetServiceDefinitionEntry

	for i := range l {
		tmp := models.FirewallInternetServiceDefinitionEntry{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.category_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.CategoryId = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallInternetServiceDefinitionEntryPortRange(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallInternetServiceDefinitionEntryPortRange
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

		pre_append = fmt.Sprintf("%s.%d.seq_num", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SeqNum = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallInternetServiceDefinitionEntryPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallInternetServiceDefinitionEntryPortRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallInternetServiceDefinitionEntryPortRange

	for i := range l {
		tmp := models.FirewallInternetServiceDefinitionEntryPortRange{}
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

func getObjectFirewallInternetServiceDefinition(d *schema.ResourceData, sv string) (*models.FirewallInternetServiceDefinition, diag.Diagnostics) {
	obj := models.FirewallInternetServiceDefinition{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("entry"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("entry", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallInternetServiceDefinitionEntry(d, v, "entry", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Entry = t
		}
	} else if d.HasChange("entry") {
		old, new := d.GetChange("entry")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Entry = &[]models.FirewallInternetServiceDefinitionEntry{}
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
