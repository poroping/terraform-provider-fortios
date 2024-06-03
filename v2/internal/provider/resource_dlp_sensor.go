// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure sensors used by DLP blocking.

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

func resourceDlpSensor() *schema.Resource {
	return &schema.Resource{
		Description: "Configure sensors used by DLP blocking.",

		CreateContext: resourceDlpSensorCreate,
		ReadContext:   resourceDlpSensorRead,
		UpdateContext: resourceDlpSensorUpdate,
		DeleteContext: resourceDlpSensorDelete,

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

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"dlp_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DLP logging.",
				Optional:    true,
				Computed:    true,
			},
			"entries": {
				Type:        schema.TypeList,
				Description: "DLP sensor entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Count of dictionary matches to trigger sensor entry match (Dictionary might not be able to trigger more than once based on its 'repeat' option, 1 - 255, default = 1).",
							Optional:    true,
							Computed:    true,
						},
						"dictionary": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Select a DLP dictionary.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable this entry.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"eval": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Expression to evaluate.",
				Optional:    true,
				Computed:    true,
			},
			"extended_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable extended logging for data leak prevention.",
				Optional:    true,
				Computed:    true,
			},
			"feature_set": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"flow", "proxy"}, false),

				Description: "Flow/proxy feature set.",
				Optional:    true,
				Computed:    true,
			},
			"filter": {
				Type:        schema.TypeList,
				Description: "Set up DLP filters for this sensor.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "log-only", "block", "quarantine-ip"}, false),

							Description: "Action to take with content that this DLP sensor matches.",
							Optional:    true,
							Computed:    true,
						},
						"archive": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable DLP archiving.",
							Optional:    true,
							Computed:    true,
						},
						"company_identifier": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Enter a company identifier watermark to match. Only watermarks that your company has placed on the files are matched.",
							Optional:    true,
							Computed:    true,
						},
						"expiry": {
							Type: schema.TypeString,

							Description: "Quarantine duration in days, hours, minutes (format = dddhhmm).",
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Type: schema.TypeInt,

							Description: "Match files this size or larger (0 - 4294967295 kbytes).",
							Optional:    true,
							Computed:    true,
						},
						"file_type": {
							Type: schema.TypeInt,

							Description: "Select the number of a DLP file pattern table to match.",
							Optional:    true,
							Computed:    true,
						},
						"filter_by": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"credit-card", "ssn", "regexp", "file-type", "file-size", "fingerprint", "watermark", "encrypted"}, false),

							Description: "Select the type of content to match.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"match_percentage": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Percentage of fingerprints in the fingerprint databases designated with the selected sensitivity to match.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter name.",
							Optional:    true,
							Computed:    true,
						},
						"proto": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Check messages or files over one or more of these protocols.",
							Optional:         true,
							Computed:         true,
						},
						"regexp": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Enter a regular expression to match (max. 255 characters).",
							Optional:    true,
							Computed:    true,
						},
						"sensitivity": {
							Type:        schema.TypeList,
							Description: "Select a DLP file pattern sensitivity to match.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Select a DLP sensitivity.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"severity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"info", "low", "medium", "high", "critical"}, false),

							Description: "Select the severity or threat level that matches this filter.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"file", "message"}, false),

							Description: "Select whether to check the content of messages (an email message) or files (downloaded files or email attachments).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"full_archive_proto": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Protocols to always content archive.",
				Optional:         true,
				Computed:         true,
			},
			"match_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"match-all", "match-any", "match-eval"}, false),

				Description: "Logical relation between entries (default = match-any).",
				Optional:    true,
				Computed:    true,
			},
			"nac_quar_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NAC quarantine logging.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of table containing the sensor.",
				ForceNew:    true,
				Required:    true,
			},
			"options": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{}, false),

				Description: "Configure DLP options.",
				Optional:    true,
				Computed:    true,
			},
			"replacemsg_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Replacement message group used by this DLP sensor.",
				Optional:    true,
				Computed:    true,
			},
			"summary_proto": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Protocols to always log summary.",
				Optional:         true,
				Computed:         true,
			},
		},
	}
}

func resourceDlpSensorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating DlpSensor resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectDlpSensor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateDlpSensor(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("DlpSensor")
	}

	return resourceDlpSensorRead(ctx, d, meta)
}

func resourceDlpSensorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectDlpSensor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateDlpSensor(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating DlpSensor resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("DlpSensor")
	}

	return resourceDlpSensorRead(ctx, d, meta)
}

func resourceDlpSensorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteDlpSensor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting DlpSensor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpSensorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadDlpSensor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading DlpSensor resource: %v", err)
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

	diags := refreshObjectDlpSensor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenDlpSensorEntries(d *schema.ResourceData, v *[]models.DlpSensorEntries, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Count; tmp != nil {
				v["count"] = *tmp
			}

			if tmp := cfg.Dictionary; tmp != nil {
				v["dictionary"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenDlpSensorFilter(d *schema.ResourceData, v *[]models.DlpSensorFilter, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Archive; tmp != nil {
				v["archive"] = *tmp
			}

			if tmp := cfg.CompanyIdentifier; tmp != nil {
				v["company_identifier"] = *tmp
			}

			if tmp := cfg.Expiry; tmp != nil {
				v["expiry"] = *tmp
			}

			if tmp := cfg.FileSize; tmp != nil {
				v["file_size"] = *tmp
			}

			if tmp := cfg.FileType; tmp != nil {
				v["file_type"] = *tmp
			}

			if tmp := cfg.FilterBy; tmp != nil {
				v["filter_by"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.MatchPercentage; tmp != nil {
				v["match_percentage"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Proto; tmp != nil {
				v["proto"] = *tmp
			}

			if tmp := cfg.Regexp; tmp != nil {
				v["regexp"] = *tmp
			}

			if tmp := cfg.Sensitivity; tmp != nil {
				v["sensitivity"] = flattenDlpSensorFilterSensitivity(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "sensitivity"), sort)
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
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

func flattenDlpSensorFilterSensitivity(d *schema.ResourceData, v *[]models.DlpSensorFilterSensitivity, prefix string, sort bool) interface{} {
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

func refreshObjectDlpSensor(d *schema.ResourceData, o *models.DlpSensor, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.DlpLog != nil {
		v := *o.DlpLog

		if err = d.Set("dlp_log", v); err != nil {
			return diag.Errorf("error reading dlp_log: %v", err)
		}
	}

	if o.Entries != nil {
		if err = d.Set("entries", flattenDlpSensorEntries(d, o.Entries, "entries", sort)); err != nil {
			return diag.Errorf("error reading entries: %v", err)
		}
	}

	if o.Eval != nil {
		v := *o.Eval

		if err = d.Set("eval", v); err != nil {
			return diag.Errorf("error reading eval: %v", err)
		}
	}

	if o.ExtendedLog != nil {
		v := *o.ExtendedLog

		if err = d.Set("extended_log", v); err != nil {
			return diag.Errorf("error reading extended_log: %v", err)
		}
	}

	if o.FeatureSet != nil {
		v := *o.FeatureSet

		if err = d.Set("feature_set", v); err != nil {
			return diag.Errorf("error reading feature_set: %v", err)
		}
	}

	if o.Filter != nil {
		if err = d.Set("filter", flattenDlpSensorFilter(d, o.Filter, "filter", sort)); err != nil {
			return diag.Errorf("error reading filter: %v", err)
		}
	}

	if o.FullArchiveProto != nil {
		v := *o.FullArchiveProto

		if err = d.Set("full_archive_proto", v); err != nil {
			return diag.Errorf("error reading full_archive_proto: %v", err)
		}
	}

	if o.MatchType != nil {
		v := *o.MatchType

		if err = d.Set("match_type", v); err != nil {
			return diag.Errorf("error reading match_type: %v", err)
		}
	}

	if o.NacQuarLog != nil {
		v := *o.NacQuarLog

		if err = d.Set("nac_quar_log", v); err != nil {
			return diag.Errorf("error reading nac_quar_log: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Options != nil {
		v := *o.Options

		if err = d.Set("options", v); err != nil {
			return diag.Errorf("error reading options: %v", err)
		}
	}

	if o.ReplacemsgGroup != nil {
		v := *o.ReplacemsgGroup

		if err = d.Set("replacemsg_group", v); err != nil {
			return diag.Errorf("error reading replacemsg_group: %v", err)
		}
	}

	if o.SummaryProto != nil {
		v := *o.SummaryProto

		if err = d.Set("summary_proto", v); err != nil {
			return diag.Errorf("error reading summary_proto: %v", err)
		}
	}

	return nil
}

func expandDlpSensorEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.DlpSensorEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.DlpSensorEntries

	for i := range l {
		tmp := models.DlpSensorEntries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Count = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dictionary", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dictionary = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandDlpSensorFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.DlpSensorFilter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.DlpSensorFilter

	for i := range l {
		tmp := models.DlpSensorFilter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.archive", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Archive = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.company_identifier", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CompanyIdentifier = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.expiry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Expiry = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.file_size", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FileSize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.file_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FileType = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter_by", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FilterBy = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_percentage", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MatchPercentage = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.proto", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Proto = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.regexp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Regexp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sensitivity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandDlpSensorFilterSensitivity(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.DlpSensorFilterSensitivity
			// 	}
			tmp.Sensitivity = v2

		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
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

func expandDlpSensorFilterSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.DlpSensorFilterSensitivity, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.DlpSensorFilterSensitivity

	for i := range l {
		tmp := models.DlpSensorFilterSensitivity{}
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

func getObjectDlpSensor(d *schema.ResourceData, sv string) (*models.DlpSensor, diag.Diagnostics) {
	obj := models.DlpSensor{}
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
	if v1, ok := d.GetOk("dlp_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.0") {
				e := utils.AttributeVersionWarning("dlp_log", sv)
				diags = append(diags, e)
			}
			obj.DlpLog = &v2
		}
	}
	if v, ok := d.GetOk("entries"); ok {
		if !utils.CheckVer(sv, "v7.2.0", "") {
			e := utils.AttributeVersionWarning("entries", sv)
			diags = append(diags, e)
		}
		t, err := expandDlpSensorEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Entries = t
		}
	} else if d.HasChange("entries") {
		old, new := d.GetChange("entries")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Entries = &[]models.DlpSensorEntries{}
		}
	}
	if v1, ok := d.GetOk("eval"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("eval", sv)
				diags = append(diags, e)
			}
			obj.Eval = &v2
		}
	}
	if v1, ok := d.GetOk("extended_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.0") {
				e := utils.AttributeVersionWarning("extended_log", sv)
				diags = append(diags, e)
			}
			obj.ExtendedLog = &v2
		}
	}
	if v1, ok := d.GetOk("feature_set"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "v7.2.0") {
				e := utils.AttributeVersionWarning("feature_set", sv)
				diags = append(diags, e)
			}
			obj.FeatureSet = &v2
		}
	}
	if v, ok := d.GetOk("filter"); ok {
		if !utils.CheckVer(sv, "", "v7.2.0") {
			e := utils.AttributeVersionWarning("filter", sv)
			diags = append(diags, e)
		}
		t, err := expandDlpSensorFilter(d, v, "filter", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Filter = t
		}
	} else if d.HasChange("filter") {
		old, new := d.GetChange("filter")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Filter = &[]models.DlpSensorFilter{}
		}
	}
	if v1, ok := d.GetOk("full_archive_proto"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.0") {
				e := utils.AttributeVersionWarning("full_archive_proto", sv)
				diags = append(diags, e)
			}
			obj.FullArchiveProto = &v2
		}
	}
	if v1, ok := d.GetOk("match_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("match_type", sv)
				diags = append(diags, e)
			}
			obj.MatchType = &v2
		}
	}
	if v1, ok := d.GetOk("nac_quar_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.0") {
				e := utils.AttributeVersionWarning("nac_quar_log", sv)
				diags = append(diags, e)
			}
			obj.NacQuarLog = &v2
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
	if v1, ok := d.GetOk("options"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("options", sv)
				diags = append(diags, e)
			}
			obj.Options = &v2
		}
	}
	if v1, ok := d.GetOk("replacemsg_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.0") {
				e := utils.AttributeVersionWarning("replacemsg_group", sv)
				diags = append(diags, e)
			}
			obj.ReplacemsgGroup = &v2
		}
	}
	if v1, ok := d.GetOk("summary_proto"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.0") {
				e := utils.AttributeVersionWarning("summary_proto", sv)
				diags = append(diags, e)
			}
			obj.SummaryProto = &v2
		}
	}
	return &obj, diags
}
