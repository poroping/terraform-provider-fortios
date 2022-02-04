// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DLP sensors.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceDlpSensor() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DLP sensors.",

		ReadContext: dataSourceDlpSensorRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"dlp_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable DLP logging.",
				Computed:    true,
			},
			"extended_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable extended logging for data leak prevention.",
				Computed:    true,
			},
			"feature_set": {
				Type:        schema.TypeString,
				Description: "Flow/proxy feature set.",
				Computed:    true,
			},
			"filter": {
				Type:        schema.TypeList,
				Description: "Set up DLP filters for this sensor.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action to take with content that this DLP sensor matches.",
							Computed:    true,
						},
						"archive": {
							Type:        schema.TypeString,
							Description: "Enable/disable DLP archiving.",
							Computed:    true,
						},
						"company_identifier": {
							Type:        schema.TypeString,
							Description: "Enter a company identifier watermark to match. Only watermarks that your company has placed on the files are matched.",
							Computed:    true,
						},
						"expiry": {
							Type:        schema.TypeString,
							Description: "Quarantine duration in days, hours, minutes (format = dddhhmm).",
							Computed:    true,
						},
						"file_size": {
							Type:        schema.TypeInt,
							Description: "Match files this size or larger (0 - 4294967295 kbytes).",
							Computed:    true,
						},
						"file_type": {
							Type:        schema.TypeInt,
							Description: "Select the number of a DLP file pattern table to match.",
							Computed:    true,
						},
						"filter_by": {
							Type:        schema.TypeString,
							Description: "Select the type of content to match.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"match_percentage": {
							Type:        schema.TypeInt,
							Description: "Percentage of fingerprints in the fingerprint databases designated with the selected sensitivity to match.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Filter name.",
							Computed:    true,
						},
						"proto": {
							Type:        schema.TypeString,
							Description: "Check messages or files over one or more of these protocols.",
							Computed:    true,
						},
						"regexp": {
							Type:        schema.TypeString,
							Description: "Enter a regular expression to match (max. 255 characters).",
							Computed:    true,
						},
						"sensitivity": {
							Type:        schema.TypeList,
							Description: "Select a DLP file pattern sensitivity to match.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Select a DLP sensitivity.",
										Computed:    true,
									},
								},
							},
						},
						"severity": {
							Type:        schema.TypeString,
							Description: "Select the severity or threat level that matches this filter.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Select whether to check the content of messages (an email message) or files (downloaded files or email attachments).",
							Computed:    true,
						},
					},
				},
			},
			"full_archive_proto": {
				Type:        schema.TypeString,
				Description: "Protocols to always content archive.",
				Computed:    true,
			},
			"nac_quar_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAC quarantine logging.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the DLP sensor.",
				Required:    true,
			},
			"options": {
				Type:        schema.TypeString,
				Description: "Configure DLP options.",
				Computed:    true,
			},
			"replacemsg_group": {
				Type:        schema.TypeString,
				Description: "Replacement message group used by this DLP sensor.",
				Computed:    true,
			},
			"summary_proto": {
				Type:        schema.TypeString,
				Description: "Protocols to always log summary.",
				Computed:    true,
			},
		},
	}
}

func dataSourceDlpSensorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadDlpSensor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading DlpSensor dataSource: %v", err)
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

	diags := refreshObjectDlpSensor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
