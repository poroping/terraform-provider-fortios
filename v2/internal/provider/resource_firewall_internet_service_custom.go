// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure custom Internet Services.

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

func resourceFirewallInternetServiceCustom() *schema.Resource {
	return &schema.Resource{
		Description: "Configure custom Internet Services.",

		CreateContext: resourceFirewallInternetServiceCustomCreate,
		ReadContext:   resourceFirewallInternetServiceCustomRead,
		UpdateContext: resourceFirewallInternetServiceCustomUpdate,
		DeleteContext: resourceFirewallInternetServiceCustomDelete,

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
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"entry": {
				Type:        schema.TypeList,
				Description: "Entries added to the Internet Service database and custom database.",
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
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Internet Service name.",
				ForceNew:    true,
				Required:    true,
			},
			"reputation": {
				Type: schema.TypeInt,

				Description: "Reputation level of the custom Internet Service.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallInternetServiceCustomCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallInternetServiceCustom resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallInternetServiceCustom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallInternetServiceCustom(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallInternetServiceCustom")
	}

	return resourceFirewallInternetServiceCustomRead(ctx, d, meta)
}

func resourceFirewallInternetServiceCustomUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallInternetServiceCustom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallInternetServiceCustom(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallInternetServiceCustom resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallInternetServiceCustom")
	}

	return resourceFirewallInternetServiceCustomRead(ctx, d, meta)
}

func resourceFirewallInternetServiceCustomDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallInternetServiceCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallInternetServiceCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceCustomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallInternetServiceCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallInternetServiceCustom resource: %v", err)
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

	diags := refreshObjectFirewallInternetServiceCustom(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallInternetServiceCustomEntry(v *[]models.FirewallInternetServiceCustomEntry, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Dst; tmp != nil {
				v["dst"] = flattenFirewallInternetServiceCustomEntryDst(tmp, sort)
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.PortRange; tmp != nil {
				v["port_range"] = flattenFirewallInternetServiceCustomEntryPortRange(tmp, sort)
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

func flattenFirewallInternetServiceCustomEntryDst(v *[]models.FirewallInternetServiceCustomEntryDst, sort bool) interface{} {
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

func flattenFirewallInternetServiceCustomEntryPortRange(v *[]models.FirewallInternetServiceCustomEntryPortRange, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func refreshObjectFirewallInternetServiceCustom(d *schema.ResourceData, o *models.FirewallInternetServiceCustom, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Entry != nil {
		if err = d.Set("entry", flattenFirewallInternetServiceCustomEntry(o.Entry, sort)); err != nil {
			return diag.Errorf("error reading entry: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Reputation != nil {
		v := *o.Reputation

		if err = d.Set("reputation", v); err != nil {
			return diag.Errorf("error reading reputation: %v", err)
		}
	}

	return nil
}

func expandFirewallInternetServiceCustomEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallInternetServiceCustomEntry, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallInternetServiceCustomEntry

	for i := range l {
		tmp := models.FirewallInternetServiceCustomEntry{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dst", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallInternetServiceCustomEntryDst(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallInternetServiceCustomEntryDst
			// 	}
			tmp.Dst = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallInternetServiceCustomEntryPortRange(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallInternetServiceCustomEntryPortRange
			// 	}
			tmp.PortRange = v2

		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Protocol = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallInternetServiceCustomEntryDst(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallInternetServiceCustomEntryDst, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallInternetServiceCustomEntryDst

	for i := range l {
		tmp := models.FirewallInternetServiceCustomEntryDst{}
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

func expandFirewallInternetServiceCustomEntryPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallInternetServiceCustomEntryPortRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallInternetServiceCustomEntryPortRange

	for i := range l {
		tmp := models.FirewallInternetServiceCustomEntryPortRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.EndPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.StartPort = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectFirewallInternetServiceCustom(d *schema.ResourceData, sv string) (*models.FirewallInternetServiceCustom, diag.Diagnostics) {
	obj := models.FirewallInternetServiceCustom{}
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
	if v, ok := d.GetOk("entry"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("entry", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallInternetServiceCustomEntry(d, v, "entry", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Entry = t
		}
	} else if d.HasChange("entry") {
		old, new := d.GetChange("entry")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Entry = &[]models.FirewallInternetServiceCustomEntry{}
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
	if v1, ok := d.GetOk("reputation"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reputation", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Reputation = &tmp
		}
	}
	return &obj, diags
}
