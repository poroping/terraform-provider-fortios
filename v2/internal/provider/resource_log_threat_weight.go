// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure threat weight settings.

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

func resourceLogThreatWeight() *schema.Resource {
	return &schema.Resource{
		Description: "Configure threat weight settings.",

		CreateContext: resourceLogThreatWeightCreate,
		ReadContext:   resourceLogThreatWeightRead,
		UpdateContext: resourceLogThreatWeightUpdate,
		DeleteContext: resourceLogThreatWeightDelete,

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
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"application": {
				Type:        schema.TypeList,
				Description: "Application-control threat weight settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Application category.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for Application events.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"blocked_connection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

				Description: "Threat weight score for blocked connections.",
				Optional:    true,
				Computed:    true,
			},
			"botnet_connection_detected": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

				Description: "Threat weight score for detected botnet connections.",
				Optional:    true,
				Computed:    true,
			},
			"failed_connection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

				Description: "Threat weight score for failed connections.",
				Optional:    true,
				Computed:    true,
			},
			"geolocation": {
				Type:        schema.TypeList,
				Description: "Geolocation-based threat weight settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"country": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 2),

							Description: "Country code.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for Geolocation-based events.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ips": {
				Type:        schema.TypeList,
				Description: "IPS threat weight settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"critical_severity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for IPS critical severity events.",
							Optional:    true,
							Computed:    true,
						},
						"high_severity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for IPS high severity events.",
							Optional:    true,
							Computed:    true,
						},
						"info_severity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for IPS info severity events.",
							Optional:    true,
							Computed:    true,
						},
						"low_severity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for IPS low severity events.",
							Optional:    true,
							Computed:    true,
						},
						"medium_severity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for IPS medium severity events.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"level": {
				Type:        schema.TypeList,
				Description: "Score mapping for threat weight levels.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"critical": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Critical level score value (1 - 100).",
							Optional:    true,
							Computed:    true,
						},
						"high": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "High level score value (1 - 100).",
							Optional:    true,
							Computed:    true,
						},
						"low": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Low level score value (1 - 100).",
							Optional:    true,
							Computed:    true,
						},
						"medium": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Medium level score value (1 - 100).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"malware": {
				Type:        schema.TypeList,
				Description: "Anti-virus malware threat weight settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_blocked": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for blocked command detected.",
							Optional:    true,
							Computed:    true,
						},
						"content_disarm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for virus (content disarm) detected.",
							Optional:    true,
							Computed:    true,
						},
						"ems_threat_feed": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for virus (EMS threat feed) detected.",
							Optional:    true,
							Computed:    true,
						},
						"file_blocked": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for blocked file detected.",
							Optional:    true,
							Computed:    true,
						},
						"fortiai": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for FortiAI-detected virus.",
							Optional:    true,
							Computed:    true,
						},
						"fsa_high_risk": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for FortiSandbox high risk malware detected.",
							Optional:    true,
							Computed:    true,
						},
						"fsa_malicious": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for FortiSandbox malicious malware detected.",
							Optional:    true,
							Computed:    true,
						},
						"fsa_medium_risk": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for FortiSandbox medium risk malware detected.",
							Optional:    true,
							Computed:    true,
						},
						"malware_list": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for virus (malware list) detected.",
							Optional:    true,
							Computed:    true,
						},
						"mimefragmented": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for mimefragmented detected.",
							Optional:    true,
							Computed:    true,
						},
						"oversized": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for oversized file detected.",
							Optional:    true,
							Computed:    true,
						},
						"switch_proto": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for switch proto detected.",
							Optional:    true,
							Computed:    true,
						},
						"virus_file_type_executable": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for virus (filetype executable) detected.",
							Optional:    true,
							Computed:    true,
						},
						"virus_infected": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for virus (infected) detected.",
							Optional:    true,
							Computed:    true,
						},
						"virus_outbreak_prevention": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for virus (outbreak prevention) event.",
							Optional:    true,
							Computed:    true,
						},
						"virus_scan_error": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for virus (scan error) detected.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the threat weight feature.",
				Optional:    true,
				Computed:    true,
			},
			"url_block_detected": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

				Description: "Threat weight score for URL blocking.",
				Optional:    true,
				Computed:    true,
			},
			"web": {
				Type:        schema.TypeList,
				Description: "Web filtering threat weight settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Threat weight score for web category filtering matches.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "low", "medium", "high", "critical"}, false),

							Description: "Threat weight score for web category filtering matches.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceLogThreatWeightCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogThreatWeight(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogThreatWeight(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogThreatWeight")
	}

	return resourceLogThreatWeightRead(ctx, d, meta)
}

func resourceLogThreatWeightUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogThreatWeight(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogThreatWeight(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogThreatWeight resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogThreatWeight")
	}

	return resourceLogThreatWeightRead(ctx, d, meta)
}

func resourceLogThreatWeightDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectLogThreatWeight(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateLogThreatWeight(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogThreatWeight resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogThreatWeightRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogThreatWeight(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogThreatWeight resource: %v", err)
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

	diags := refreshObjectLogThreatWeight(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenLogThreatWeightApplication(v *[]models.LogThreatWeightApplication, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Level; tmp != nil {
				v["level"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenLogThreatWeightGeolocation(v *[]models.LogThreatWeightGeolocation, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Country; tmp != nil {
				v["country"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Level; tmp != nil {
				v["level"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenLogThreatWeightIps(v *[]models.LogThreatWeightIps, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CriticalSeverity; tmp != nil {
				v["critical_severity"] = *tmp
			}

			if tmp := cfg.HighSeverity; tmp != nil {
				v["high_severity"] = *tmp
			}

			if tmp := cfg.InfoSeverity; tmp != nil {
				v["info_severity"] = *tmp
			}

			if tmp := cfg.LowSeverity; tmp != nil {
				v["low_severity"] = *tmp
			}

			if tmp := cfg.MediumSeverity; tmp != nil {
				v["medium_severity"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenLogThreatWeightLevel(v *[]models.LogThreatWeightLevel, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Critical; tmp != nil {
				v["critical"] = *tmp
			}

			if tmp := cfg.High; tmp != nil {
				v["high"] = *tmp
			}

			if tmp := cfg.Low; tmp != nil {
				v["low"] = *tmp
			}

			if tmp := cfg.Medium; tmp != nil {
				v["medium"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenLogThreatWeightMalware(v *[]models.LogThreatWeightMalware, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CommandBlocked; tmp != nil {
				v["command_blocked"] = *tmp
			}

			if tmp := cfg.ContentDisarm; tmp != nil {
				v["content_disarm"] = *tmp
			}

			if tmp := cfg.EmsThreatFeed; tmp != nil {
				v["ems_threat_feed"] = *tmp
			}

			if tmp := cfg.FileBlocked; tmp != nil {
				v["file_blocked"] = *tmp
			}

			if tmp := cfg.Fortiai; tmp != nil {
				v["fortiai"] = *tmp
			}

			if tmp := cfg.FsaHighRisk; tmp != nil {
				v["fsa_high_risk"] = *tmp
			}

			if tmp := cfg.FsaMalicious; tmp != nil {
				v["fsa_malicious"] = *tmp
			}

			if tmp := cfg.FsaMediumRisk; tmp != nil {
				v["fsa_medium_risk"] = *tmp
			}

			if tmp := cfg.MalwareList; tmp != nil {
				v["malware_list"] = *tmp
			}

			if tmp := cfg.Mimefragmented; tmp != nil {
				v["mimefragmented"] = *tmp
			}

			if tmp := cfg.Oversized; tmp != nil {
				v["oversized"] = *tmp
			}

			if tmp := cfg.SwitchProto; tmp != nil {
				v["switch_proto"] = *tmp
			}

			if tmp := cfg.VirusFileTypeExecutable; tmp != nil {
				v["virus_file_type_executable"] = *tmp
			}

			if tmp := cfg.VirusInfected; tmp != nil {
				v["virus_infected"] = *tmp
			}

			if tmp := cfg.VirusOutbreakPrevention; tmp != nil {
				v["virus_outbreak_prevention"] = *tmp
			}

			if tmp := cfg.VirusScanError; tmp != nil {
				v["virus_scan_error"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenLogThreatWeightWeb(v *[]models.LogThreatWeightWeb, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Level; tmp != nil {
				v["level"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectLogThreatWeight(d *schema.ResourceData, o *models.LogThreatWeight, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Application != nil {
		if err = d.Set("application", flattenLogThreatWeightApplication(o.Application, sort)); err != nil {
			return diag.Errorf("error reading application: %v", err)
		}
	}

	if o.BlockedConnection != nil {
		v := *o.BlockedConnection

		if err = d.Set("blocked_connection", v); err != nil {
			return diag.Errorf("error reading blocked_connection: %v", err)
		}
	}

	if o.BotnetConnectionDetected != nil {
		v := *o.BotnetConnectionDetected

		if err = d.Set("botnet_connection_detected", v); err != nil {
			return diag.Errorf("error reading botnet_connection_detected: %v", err)
		}
	}

	if o.FailedConnection != nil {
		v := *o.FailedConnection

		if err = d.Set("failed_connection", v); err != nil {
			return diag.Errorf("error reading failed_connection: %v", err)
		}
	}

	if o.Geolocation != nil {
		if err = d.Set("geolocation", flattenLogThreatWeightGeolocation(o.Geolocation, sort)); err != nil {
			return diag.Errorf("error reading geolocation: %v", err)
		}
	}

	if o.Ips != nil {
		if err = d.Set("ips", flattenLogThreatWeightIps(o.Ips, sort)); err != nil {
			return diag.Errorf("error reading ips: %v", err)
		}
	}

	if o.Level != nil {
		if err = d.Set("level", flattenLogThreatWeightLevel(o.Level, sort)); err != nil {
			return diag.Errorf("error reading level: %v", err)
		}
	}

	if o.Malware != nil {
		if err = d.Set("malware", flattenLogThreatWeightMalware(o.Malware, sort)); err != nil {
			return diag.Errorf("error reading malware: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.UrlBlockDetected != nil {
		v := *o.UrlBlockDetected

		if err = d.Set("url_block_detected", v); err != nil {
			return diag.Errorf("error reading url_block_detected: %v", err)
		}
	}

	if o.Web != nil {
		if err = d.Set("web", flattenLogThreatWeightWeb(o.Web, sort)); err != nil {
			return diag.Errorf("error reading web: %v", err)
		}
	}

	return nil
}

func expandLogThreatWeightApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.LogThreatWeightApplication, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.LogThreatWeightApplication

	for i := range l {
		tmp := models.LogThreatWeightApplication{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Level = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandLogThreatWeightGeolocation(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.LogThreatWeightGeolocation, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.LogThreatWeightGeolocation

	for i := range l {
		tmp := models.LogThreatWeightGeolocation{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.country", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Country = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Level = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandLogThreatWeightIps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.LogThreatWeightIps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.LogThreatWeightIps

	for i := range l {
		tmp := models.LogThreatWeightIps{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.critical_severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CriticalSeverity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.high_severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HighSeverity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.info_severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InfoSeverity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.low_severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LowSeverity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.medium_severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MediumSeverity = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandLogThreatWeightLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.LogThreatWeightLevel, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.LogThreatWeightLevel

	for i := range l {
		tmp := models.LogThreatWeightLevel{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.critical", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Critical = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.high", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.High = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.low", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Low = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.medium", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Medium = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandLogThreatWeightMalware(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.LogThreatWeightMalware, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.LogThreatWeightMalware

	for i := range l {
		tmp := models.LogThreatWeightMalware{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.command_blocked", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CommandBlocked = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.content_disarm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ContentDisarm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ems_threat_feed", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EmsThreatFeed = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.file_blocked", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FileBlocked = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortiai", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortiai = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fsa_high_risk", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FsaHighRisk = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fsa_malicious", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FsaMalicious = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fsa_medium_risk", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FsaMediumRisk = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malware_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalwareList = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mimefragmented", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mimefragmented = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.oversized", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Oversized = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.switch_proto", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SwitchProto = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.virus_file_type_executable", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VirusFileTypeExecutable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.virus_infected", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VirusInfected = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.virus_outbreak_prevention", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VirusOutbreakPrevention = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.virus_scan_error", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VirusScanError = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandLogThreatWeightWeb(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.LogThreatWeightWeb, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.LogThreatWeightWeb

	for i := range l {
		tmp := models.LogThreatWeightWeb{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Level = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectLogThreatWeight(d *schema.ResourceData, sv string) (*models.LogThreatWeight, diag.Diagnostics) {
	obj := models.LogThreatWeight{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("application"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("application", sv)
			diags = append(diags, e)
		}
		t, err := expandLogThreatWeightApplication(d, v, "application", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Application = t
		}
	} else if d.HasChange("application") {
		old, new := d.GetChange("application")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Application = &[]models.LogThreatWeightApplication{}
		}
	}
	if v1, ok := d.GetOk("blocked_connection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("blocked_connection", sv)
				diags = append(diags, e)
			}
			obj.BlockedConnection = &v2
		}
	}
	if v1, ok := d.GetOk("botnet_connection_detected"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("botnet_connection_detected", sv)
				diags = append(diags, e)
			}
			obj.BotnetConnectionDetected = &v2
		}
	}
	if v1, ok := d.GetOk("failed_connection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("failed_connection", sv)
				diags = append(diags, e)
			}
			obj.FailedConnection = &v2
		}
	}
	if v, ok := d.GetOk("geolocation"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("geolocation", sv)
			diags = append(diags, e)
		}
		t, err := expandLogThreatWeightGeolocation(d, v, "geolocation", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Geolocation = t
		}
	} else if d.HasChange("geolocation") {
		old, new := d.GetChange("geolocation")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Geolocation = &[]models.LogThreatWeightGeolocation{}
		}
	}
	if v, ok := d.GetOk("ips"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ips", sv)
			diags = append(diags, e)
		}
		t, err := expandLogThreatWeightIps(d, v, "ips", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ips = t
		}
	} else if d.HasChange("ips") {
		old, new := d.GetChange("ips")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ips = &[]models.LogThreatWeightIps{}
		}
	}
	if v, ok := d.GetOk("level"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("level", sv)
			diags = append(diags, e)
		}
		t, err := expandLogThreatWeightLevel(d, v, "level", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Level = t
		}
	} else if d.HasChange("level") {
		old, new := d.GetChange("level")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Level = &[]models.LogThreatWeightLevel{}
		}
	}
	if v, ok := d.GetOk("malware"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("malware", sv)
			diags = append(diags, e)
		}
		t, err := expandLogThreatWeightMalware(d, v, "malware", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Malware = t
		}
	} else if d.HasChange("malware") {
		old, new := d.GetChange("malware")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Malware = &[]models.LogThreatWeightMalware{}
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
	if v1, ok := d.GetOk("url_block_detected"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("url_block_detected", sv)
				diags = append(diags, e)
			}
			obj.UrlBlockDetected = &v2
		}
	}
	if v, ok := d.GetOk("web"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("web", sv)
			diags = append(diags, e)
		}
		t, err := expandLogThreatWeightWeb(d, v, "web", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Web = t
		}
	} else if d.HasChange("web") {
		old, new := d.GetChange("web")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Web = &[]models.LogThreatWeightWeb{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectLogThreatWeight(d *schema.ResourceData, sv string) (*models.LogThreatWeight, diag.Diagnostics) {
	obj := models.LogThreatWeight{}
	diags := diag.Diagnostics{}

	obj.Application = &[]models.LogThreatWeightApplication{}
	obj.Geolocation = &[]models.LogThreatWeightGeolocation{}
	obj.Ips = &[]models.LogThreatWeightIps{}
	obj.Level = &[]models.LogThreatWeightLevel{}
	obj.Malware = &[]models.LogThreatWeightMalware{}
	obj.Web = &[]models.LogThreatWeightWeb{}

	return &obj, diags
}
