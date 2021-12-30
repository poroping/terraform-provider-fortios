// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure anti-spam block/allow list.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceEmailfilterBlockAllowList() *schema.Resource {
	return &schema.Resource{
		Description: "Configure anti-spam block/allow list.",

		CreateContext: resourceEmailfilterBlockAllowListCreate,
		ReadContext:   resourceEmailfilterBlockAllowListRead,
		UpdateContext: resourceEmailfilterBlockAllowListUpdate,
		DeleteContext: resourceEmailfilterBlockAllowListDelete,

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
				Description:  "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"entries": {
				Type:        schema.TypeList,
				Description: "Anti-spam block/allow entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"reject", "spam", "clear"}, false),

							Description: "Reject, mark as spam or good email.",
							Optional:    true,
							Computed:    true,
						},
						"addr_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),

							Description: "IP address type.",
							Optional:    true,
							Computed:    true,
						},
						"email_pattern": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Email address pattern.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"ip4_subnet": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Classnet,

							Description: "IPv4 network address/subnet mask bits.",
							Optional:    true,
							Computed:    true,
						},
						"ip6_subnet": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Network,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "IPv6 network address/subnet mask bits.",
							Optional:         true,
							Computed:         true,
						},
						"pattern_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"wildcard", "regexp"}, false),

							Description: "Wildcard pattern or regular expression.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable status.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ip", "email"}, false),

							Description: "Entry type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "ID.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Name of table.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceEmailfilterBlockAllowListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating EmailfilterBlockAllowList resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectEmailfilterBlockAllowList(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateEmailfilterBlockAllowList(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EmailfilterBlockAllowList")
	}

	return resourceEmailfilterBlockAllowListRead(ctx, d, meta)
}

func resourceEmailfilterBlockAllowListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectEmailfilterBlockAllowList(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateEmailfilterBlockAllowList(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating EmailfilterBlockAllowList resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EmailfilterBlockAllowList")
	}

	return resourceEmailfilterBlockAllowListRead(ctx, d, meta)
}

func resourceEmailfilterBlockAllowListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteEmailfilterBlockAllowList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting EmailfilterBlockAllowList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterBlockAllowListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadEmailfilterBlockAllowList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading EmailfilterBlockAllowList resource: %v", err)
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

	diags := refreshObjectEmailfilterBlockAllowList(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenEmailfilterBlockAllowListEntries(v *[]models.EmailfilterBlockAllowListEntries, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.AddrType; tmp != nil {
				v["addr_type"] = *tmp
			}

			if tmp := cfg.EmailPattern; tmp != nil {
				v["email_pattern"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip4Subnet; tmp != nil {
				v["ip4_subnet"] = *tmp
			}

			if tmp := cfg.Ip6Subnet; tmp != nil {
				v["ip6_subnet"] = *tmp
			}

			if tmp := cfg.PatternType; tmp != nil {
				v["pattern_type"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectEmailfilterBlockAllowList(d *schema.ResourceData, o *models.EmailfilterBlockAllowList, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Entries != nil {
		if err = d.Set("entries", flattenEmailfilterBlockAllowListEntries(o.Entries, sort)); err != nil {
			return diag.Errorf("error reading entries: %v", err)
		}
	}

	if o.Fosid != nil {
		v := *o.Fosid

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	return nil
}

func expandEmailfilterBlockAllowListEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.EmailfilterBlockAllowListEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterBlockAllowListEntries

	for i := range l {
		tmp := models.EmailfilterBlockAllowListEntries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.addr_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AddrType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.email_pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EmailPattern = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip4_subnet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip4Subnet = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_subnet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6Subnet = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pattern_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PatternType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectEmailfilterBlockAllowList(d *schema.ResourceData, sv string) (*models.EmailfilterBlockAllowList, diag.Diagnostics) {
	obj := models.EmailfilterBlockAllowList{}
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
	if v, ok := d.GetOk("entries"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("entries", sv)
			diags = append(diags, e)
		}
		t, err := expandEmailfilterBlockAllowListEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Entries = t
		}
	} else if d.HasChange("entries") {
		old, new := d.GetChange("entries")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Entries = &[]models.EmailfilterBlockAllowListEntries{}
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
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	return &obj, diags
}
