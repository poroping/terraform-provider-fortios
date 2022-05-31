// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure threat weight settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogThreatWeight() *schema.Resource {
	return &schema.Resource{
		Description: "Configure threat weight settings.",

		ReadContext: dataSourceLogThreatWeightRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"application": {
				Type:        schema.TypeList,
				Description: "Application-control threat weight settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:        schema.TypeInt,
							Description: "Application category.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Entry ID.",
							Computed:    true,
						},
						"level": {
							Type:        schema.TypeString,
							Description: "Threat weight score for Application events.",
							Computed:    true,
						},
					},
				},
			},
			"blocked_connection": {
				Type:        schema.TypeString,
				Description: "Threat weight score for blocked connections.",
				Computed:    true,
			},
			"botnet_connection_detected": {
				Type:        schema.TypeString,
				Description: "Threat weight score for detected botnet connections.",
				Computed:    true,
			},
			"failed_connection": {
				Type:        schema.TypeString,
				Description: "Threat weight score for failed connections.",
				Computed:    true,
			},
			"geolocation": {
				Type:        schema.TypeList,
				Description: "Geolocation-based threat weight settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"country": {
							Type:        schema.TypeString,
							Description: "Country code.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Entry ID.",
							Computed:    true,
						},
						"level": {
							Type:        schema.TypeString,
							Description: "Threat weight score for Geolocation-based events.",
							Computed:    true,
						},
					},
				},
			},
			"ips": {
				Type:        schema.TypeList,
				Description: "IPS threat weight settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"critical_severity": {
							Type:        schema.TypeString,
							Description: "Threat weight score for IPS critical severity events.",
							Computed:    true,
						},
						"high_severity": {
							Type:        schema.TypeString,
							Description: "Threat weight score for IPS high severity events.",
							Computed:    true,
						},
						"info_severity": {
							Type:        schema.TypeString,
							Description: "Threat weight score for IPS info severity events.",
							Computed:    true,
						},
						"low_severity": {
							Type:        schema.TypeString,
							Description: "Threat weight score for IPS low severity events.",
							Computed:    true,
						},
						"medium_severity": {
							Type:        schema.TypeString,
							Description: "Threat weight score for IPS medium severity events.",
							Computed:    true,
						},
					},
				},
			},
			"level": {
				Type:        schema.TypeList,
				Description: "Score mapping for threat weight levels.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"critical": {
							Type:        schema.TypeInt,
							Description: "Critical level score value (1 - 100).",
							Computed:    true,
						},
						"high": {
							Type:        schema.TypeInt,
							Description: "High level score value (1 - 100).",
							Computed:    true,
						},
						"low": {
							Type:        schema.TypeInt,
							Description: "Low level score value (1 - 100).",
							Computed:    true,
						},
						"medium": {
							Type:        schema.TypeInt,
							Description: "Medium level score value (1 - 100).",
							Computed:    true,
						},
					},
				},
			},
			"malware": {
				Type:        schema.TypeList,
				Description: "Anti-virus malware threat weight settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_blocked": {
							Type:        schema.TypeString,
							Description: "Threat weight score for blocked command detected.",
							Computed:    true,
						},
						"content_disarm": {
							Type:        schema.TypeString,
							Description: "Threat weight score for virus (content disarm) detected.",
							Computed:    true,
						},
						"ems_threat_feed": {
							Type:        schema.TypeString,
							Description: "Threat weight score for virus (EMS threat feed) detected.",
							Computed:    true,
						},
						"file_blocked": {
							Type:        schema.TypeString,
							Description: "Threat weight score for blocked file detected.",
							Computed:    true,
						},
						"fortiai": {
							Type:        schema.TypeString,
							Description: "Threat weight score for FortiAI-detected virus.",
							Computed:    true,
						},
						"fortindr": {
							Type:        schema.TypeString,
							Description: "Threat weight score for FortiNDR-detected virus.",
							Computed:    true,
						},
						"fortisandbox": {
							Type:        schema.TypeString,
							Description: "Threat weight score for FortiSandbox-detected virus.",
							Computed:    true,
						},
						"fsa_high_risk": {
							Type:        schema.TypeString,
							Description: "Threat weight score for FortiSandbox high risk malware detected.",
							Computed:    true,
						},
						"fsa_malicious": {
							Type:        schema.TypeString,
							Description: "Threat weight score for FortiSandbox malicious malware detected.",
							Computed:    true,
						},
						"fsa_medium_risk": {
							Type:        schema.TypeString,
							Description: "Threat weight score for FortiSandbox medium risk malware detected.",
							Computed:    true,
						},
						"malware_list": {
							Type:        schema.TypeString,
							Description: "Threat weight score for virus (malware list) detected.",
							Computed:    true,
						},
						"mimefragmented": {
							Type:        schema.TypeString,
							Description: "Threat weight score for mimefragmented detected.",
							Computed:    true,
						},
						"oversized": {
							Type:        schema.TypeString,
							Description: "Threat weight score for oversized file detected.",
							Computed:    true,
						},
						"switch_proto": {
							Type:        schema.TypeString,
							Description: "Threat weight score for switch proto detected.",
							Computed:    true,
						},
						"virus_file_type_executable": {
							Type:        schema.TypeString,
							Description: "Threat weight score for virus (file type executable) detected.",
							Computed:    true,
						},
						"virus_infected": {
							Type:        schema.TypeString,
							Description: "Threat weight score for virus (infected) detected.",
							Computed:    true,
						},
						"virus_outbreak_prevention": {
							Type:        schema.TypeString,
							Description: "Threat weight score for virus (outbreak prevention) event.",
							Computed:    true,
						},
						"virus_scan_error": {
							Type:        schema.TypeString,
							Description: "Threat weight score for virus (scan error) detected.",
							Computed:    true,
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable the threat weight feature.",
				Computed:    true,
			},
			"url_block_detected": {
				Type:        schema.TypeString,
				Description: "Threat weight score for URL blocking.",
				Computed:    true,
			},
			"web": {
				Type:        schema.TypeList,
				Description: "Web filtering threat weight settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:        schema.TypeInt,
							Description: "Threat weight score for web category filtering matches.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Entry ID.",
							Computed:    true,
						},
						"level": {
							Type:        schema.TypeString,
							Description: "Threat weight score for web category filtering matches.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLogThreatWeightRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogThreatWeight"

	o, err := c.Cmdb.ReadLogThreatWeight(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogThreatWeight dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
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

	d.SetId(mkey)

	return nil
}
