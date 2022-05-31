// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure AntiVirus profiles.

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

func resourceAntivirusProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure AntiVirus profiles.",

		CreateContext: resourceAntivirusProfileCreate,
		ReadContext:   resourceAntivirusProfileRead,
		UpdateContext: resourceAntivirusProfileUpdate,
		DeleteContext: resourceAntivirusProfileDelete,

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
			"analytics_accept_filetype": {
				Type: schema.TypeInt,

				Description: "Only submit files matching this DLP file-pattern to FortiSandbox (post-transfer scan only).",
				Optional:    true,
				Computed:    true,
			},
			"analytics_bl_filetype": {
				Type: schema.TypeInt,

				Description: "Only submit files matching this DLP file-pattern to FortiSandbox.",
				Optional:    true,
				Computed:    true,
			},
			"analytics_db": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable using the FortiSandbox signature database to supplement the AV signature databases.",
				Optional:    true,
				Computed:    true,
			},
			"analytics_ignore_filetype": {
				Type: schema.TypeInt,

				Description: "Do not submit files matching this DLP file-pattern to FortiSandbox (post-transfer scan only).",
				Optional:    true,
				Computed:    true,
			},
			"analytics_max_upload": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 26214),

				Description: "Maximum size of files that can be uploaded to FortiSandbox.",
				Optional:    true,
				Computed:    true,
			},
			"analytics_wl_filetype": {
				Type: schema.TypeInt,

				Description: "Do not submit files matching this DLP file-pattern to FortiSandbox.",
				Optional:    true,
				Computed:    true,
			},
			"av_block_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging for AntiVirus file blocking.",
				Optional:    true,
				Computed:    true,
			},
			"av_virus_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable AntiVirus logging.",
				Optional:    true,
				Computed:    true,
			},
			"cifs": {
				Type:        schema.TypeList,
				Description: "Configure CIFS AntiVirus options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to block.",
							Optional:         true,
							Computed:         true,
						},
						"archive_log": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to log.",
							Optional:         true,
							Computed:         true,
						},
						"av_scan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable AntiVirus scan service.",
							Optional:    true,
							Computed:    true,
						},
						"emulator": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the virus emulator.",
							Optional:    true,
							Computed:    true,
						},
						"external_blocklist": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Optional:    true,
							Computed:    true,
						},
						"fortiai": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable/disable scanning of files by FortiAI.",
							Optional:    true,
							Computed:    true,
						},
						"fortindr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiNDR.",
							Optional:    true,
							Computed:    true,
						},
						"fortisandbox": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiSandbox.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable/disable CIFS AntiVirus scanning, monitoring, and quarantine.",
							Optional:         true,
							Computed:         true,
						},
						"outbreak_prevention": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable virus outbreak prevention service.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable quarantine for infected files.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"content_disarm": {
				Type:        schema.TypeList,
				Description: "AV Content Disarm and Reconstruction settings.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cover_page": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable inserting a cover page into the disarmed document.",
							Optional:    true,
							Computed:    true,
						},
						"detect_only": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable only detect disarmable files, do not alter content.",
							Optional:    true,
							Computed:    true,
						},
						"error_action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"block", "log-only", "ignore"}, false),

							Description: "Action to be taken if CDR engine encounters an unrecoverable error.",
							Optional:    true,
							Computed:    true,
						},
						"office_action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of PowerPoint action events in Microsoft Office documents.",
							Optional:    true,
							Computed:    true,
						},
						"office_dde": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of Dynamic Data Exchange events in Microsoft Office documents.",
							Optional:    true,
							Computed:    true,
						},
						"office_embed": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of embedded objects in Microsoft Office documents.",
							Optional:    true,
							Computed:    true,
						},
						"office_hylink": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of hyperlinks in Microsoft Office documents.",
							Optional:    true,
							Computed:    true,
						},
						"office_linked": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of linked objects in Microsoft Office documents.",
							Optional:    true,
							Computed:    true,
						},
						"office_macro": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of macros in Microsoft Office documents.",
							Optional:    true,
							Computed:    true,
						},
						"original_file_destination": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"fortisandbox", "quarantine", "discard"}, false),

							Description: "Destination to send original file if active content is removed.",
							Optional:    true,
							Computed:    true,
						},
						"pdf_act_form": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of PDF document actions that submit data to other targets.",
							Optional:    true,
							Computed:    true,
						},
						"pdf_act_gotor": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of PDF document actions that access other PDF documents.",
							Optional:    true,
							Computed:    true,
						},
						"pdf_act_java": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of PDF document actions that execute JavaScript code.",
							Optional:    true,
							Computed:    true,
						},
						"pdf_act_launch": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of PDF document actions that launch other applications.",
							Optional:    true,
							Computed:    true,
						},
						"pdf_act_movie": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of PDF document actions that play a movie.",
							Optional:    true,
							Computed:    true,
						},
						"pdf_act_sound": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of PDF document actions that play a sound.",
							Optional:    true,
							Computed:    true,
						},
						"pdf_embedfile": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of embedded files in PDF documents.",
							Optional:    true,
							Computed:    true,
						},
						"pdf_hyperlink": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of hyperlinks from PDF documents.",
							Optional:    true,
							Computed:    true,
						},
						"pdf_javacode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of JavaScript code in PDF documents.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ems_threat_feed": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable use of EMS threat feed when performing AntiVirus scan. Analyzes files including the content of archives.",
				Optional:    true,
				Computed:    true,
			},
			"extended_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable extended logging for antivirus.",
				Optional:    true,
				Computed:    true,
			},
			"external_blocklist": {
				Type:        schema.TypeList,
				Description: "One or more external malware block lists.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "External blocklist.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"external_blocklist_archive_scan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable external-blocklist archive scanning.",
				Optional:    true,
				Computed:    true,
			},
			"external_blocklist_enable_all": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable all external blocklists.",
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
			"fortiai_error_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"log-only", "block", "ignore"}, false),

				Description: "Action to take if FortiAI encounters an error.",
				Optional:    true,
				Computed:    true,
			},
			"fortiai_timeout_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"log-only", "block", "ignore"}, false),

				Description: "Action to take if FortiAI encounters a scan timeout.",
				Optional:    true,
				Computed:    true,
			},
			"fortindr_error_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"log-only", "block", "ignore"}, false),

				Description: "Action to take if FortiNDR encounters an error.",
				Optional:    true,
				Computed:    true,
			},
			"fortindr_timeout_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"log-only", "block", "ignore"}, false),

				Description: "Action to take if FortiNDR encounters a scan timeout.",
				Optional:    true,
				Computed:    true,
			},
			"fortisandbox_error_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"log-only", "block", "ignore"}, false),

				Description: "Action to take if FortiSandbox inline scan encounters an error.",
				Optional:    true,
				Computed:    true,
			},
			"fortisandbox_max_upload": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 26214),

				Description: "Maximum size of files that can be uploaded to FortiSandbox.",
				Optional:    true,
				Computed:    true,
			},
			"fortisandbox_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"inline", "analytics-suspicious", "analytics-everything"}, false),

				Description: "FortiSandbox scan modes.",
				Optional:    true,
				Computed:    true,
			},
			"fortisandbox_timeout_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"log-only", "block", "ignore"}, false),

				Description: "Action to take if FortiSandbox inline scan encounters a scan timeout.",
				Optional:    true,
				Computed:    true,
			},
			"ftgd_analytics": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "suspicious", "everything"}, false),

				Description: "Settings to control which files are uploaded to FortiSandbox.",
				Optional:    true,
				Computed:    true,
			},
			"ftp": {
				Type:        schema.TypeList,
				Description: "Configure FTP AntiVirus options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to block.",
							Optional:         true,
							Computed:         true,
						},
						"archive_log": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to log.",
							Optional:         true,
							Computed:         true,
						},
						"av_scan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable AntiVirus scan service.",
							Optional:    true,
							Computed:    true,
						},
						"emulator": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the virus emulator.",
							Optional:    true,
							Computed:    true,
						},
						"external_blocklist": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Optional:    true,
							Computed:    true,
						},
						"fortiai": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable/disable scanning of files by FortiAI.",
							Optional:    true,
							Computed:    true,
						},
						"fortindr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiNDR.",
							Optional:    true,
							Computed:    true,
						},
						"fortisandbox": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiSandbox.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable/disable FTP AntiVirus scanning, monitoring, and quarantine.",
							Optional:         true,
							Computed:         true,
						},
						"outbreak_prevention": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable virus outbreak prevention service.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable quarantine for infected files.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"http": {
				Type:        schema.TypeList,
				Description: "Configure HTTP AntiVirus options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to block.",
							Optional:         true,
							Computed:         true,
						},
						"archive_log": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to log.",
							Optional:         true,
							Computed:         true,
						},
						"av_scan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable AntiVirus scan service.",
							Optional:    true,
							Computed:    true,
						},
						"content_disarm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan.",
							Optional:    true,
							Computed:    true,
						},
						"emulator": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the virus emulator.",
							Optional:    true,
							Computed:    true,
						},
						"external_blocklist": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Optional:    true,
							Computed:    true,
						},
						"fortiai": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable/disable scanning of files by FortiAI.",
							Optional:    true,
							Computed:    true,
						},
						"fortindr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiNDR.",
							Optional:    true,
							Computed:    true,
						},
						"fortisandbox": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiSandbox.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable/disable HTTP AntiVirus scanning, monitoring, and quarantine.",
							Optional:         true,
							Computed:         true,
						},
						"outbreak_prevention": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable virus outbreak prevention service.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable quarantine for infected files.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"imap": {
				Type:        schema.TypeList,
				Description: "Configure IMAP AntiVirus options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to block.",
							Optional:         true,
							Computed:         true,
						},
						"archive_log": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to log.",
							Optional:         true,
							Computed:         true,
						},
						"av_scan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable AntiVirus scan service.",
							Optional:    true,
							Computed:    true,
						},
						"content_disarm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan.",
							Optional:    true,
							Computed:    true,
						},
						"emulator": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the virus emulator.",
							Optional:    true,
							Computed:    true,
						},
						"executables": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"default", "virus"}, false),

							Description: "Treat Windows executable files as viruses for the purpose of blocking or monitoring.",
							Optional:    true,
							Computed:    true,
						},
						"external_blocklist": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Optional:    true,
							Computed:    true,
						},
						"fortiai": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable/disable scanning of files by FortiAI.",
							Optional:    true,
							Computed:    true,
						},
						"fortindr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiNDR.",
							Optional:    true,
							Computed:    true,
						},
						"fortisandbox": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiSandbox.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable/disable IMAP AntiVirus scanning, monitoring, and quarantine.",
							Optional:         true,
							Computed:         true,
						},
						"outbreak_prevention": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable virus outbreak prevention service.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable quarantine for infected files.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"mapi": {
				Type:        schema.TypeList,
				Description: "Configure MAPI AntiVirus options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to block.",
							Optional:         true,
							Computed:         true,
						},
						"archive_log": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to log.",
							Optional:         true,
							Computed:         true,
						},
						"av_scan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable AntiVirus scan service.",
							Optional:    true,
							Computed:    true,
						},
						"emulator": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the virus emulator.",
							Optional:    true,
							Computed:    true,
						},
						"executables": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"default", "virus"}, false),

							Description: "Treat Windows executable files as viruses for the purpose of blocking or monitoring.",
							Optional:    true,
							Computed:    true,
						},
						"external_blocklist": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Optional:    true,
							Computed:    true,
						},
						"fortiai": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable/disable scanning of files by FortiAI.",
							Optional:    true,
							Computed:    true,
						},
						"fortindr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiNDR.",
							Optional:    true,
							Computed:    true,
						},
						"fortisandbox": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiSandbox.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable/disable MAPI AntiVirus scanning, monitoring, and quarantine.",
							Optional:         true,
							Computed:         true,
						},
						"outbreak_prevention": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable virus outbreak prevention service.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable quarantine for infected files.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"mobile_malware_db": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable using the mobile malware signature database.",
				Optional:    true,
				Computed:    true,
			},
			"nac_quar": {
				Type:        schema.TypeList,
				Description: "Configure AntiVirus quarantine settings.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expiry": {
							Type: schema.TypeString,

							Description: "Duration of quarantine.",
							Optional:    true,
							Computed:    true,
						},
						"infected": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "quar-src-ip"}, false),

							Description: "Enable/Disable quarantining infected hosts to the banned user list.",
							Optional:    true,
							Computed:    true,
						},
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable AntiVirus quarantine logging.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"nntp": {
				Type:        schema.TypeList,
				Description: "Configure NNTP AntiVirus options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to block.",
							Optional:         true,
							Computed:         true,
						},
						"archive_log": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to log.",
							Optional:         true,
							Computed:         true,
						},
						"av_scan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable AntiVirus scan service.",
							Optional:    true,
							Computed:    true,
						},
						"emulator": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the virus emulator.",
							Optional:    true,
							Computed:    true,
						},
						"external_blocklist": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Optional:    true,
							Computed:    true,
						},
						"fortiai": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable/disable scanning of files by FortiAI.",
							Optional:    true,
							Computed:    true,
						},
						"fortindr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiNDR.",
							Optional:    true,
							Computed:    true,
						},
						"fortisandbox": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiSandbox.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable/disable NNTP AntiVirus scanning, monitoring, and quarantine.",
							Optional:         true,
							Computed:         true,
						},
						"outbreak_prevention": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable virus outbreak prevention service.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable quarantine for infected files.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"outbreak_prevention": {
				Type:        schema.TypeList,
				Description: "Configure Virus Outbreak Prevention settings.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_blocklist": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable external malware blocklist.",
							Optional:    true,
							Computed:    true,
						},
						"ftgd_service": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable FortiGuard Virus outbreak prevention service.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"outbreak_prevention_archive_scan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable outbreak-prevention archive scanning.",
				Optional:    true,
				Computed:    true,
			},
			"pop3": {
				Type:        schema.TypeList,
				Description: "Configure POP3 AntiVirus options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to block.",
							Optional:         true,
							Computed:         true,
						},
						"archive_log": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to log.",
							Optional:         true,
							Computed:         true,
						},
						"av_scan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable AntiVirus scan service.",
							Optional:    true,
							Computed:    true,
						},
						"content_disarm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan.",
							Optional:    true,
							Computed:    true,
						},
						"emulator": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the virus emulator.",
							Optional:    true,
							Computed:    true,
						},
						"executables": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"default", "virus"}, false),

							Description: "Treat Windows executable files as viruses for the purpose of blocking or monitoring.",
							Optional:    true,
							Computed:    true,
						},
						"external_blocklist": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Optional:    true,
							Computed:    true,
						},
						"fortiai": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable/disable scanning of files by FortiAI.",
							Optional:    true,
							Computed:    true,
						},
						"fortindr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiNDR.",
							Optional:    true,
							Computed:    true,
						},
						"fortisandbox": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiSandbox.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable/disable POP3 AntiVirus scanning, monitoring, and quarantine.",
							Optional:         true,
							Computed:         true,
						},
						"outbreak_prevention": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable virus outbreak prevention service.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable quarantine for infected files.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"replacemsg_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Replacement message group customized for this profile.",
				Optional:    true,
				Computed:    true,
			},
			"scan_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "legacy"}, false),

				Description: "Configure scan mode (default or legacy).",
				Optional:    true,
				Computed:    true,
			},
			"smtp": {
				Type:        schema.TypeList,
				Description: "Configure SMTP AntiVirus options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to block.",
							Optional:         true,
							Computed:         true,
						},
						"archive_log": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to log.",
							Optional:         true,
							Computed:         true,
						},
						"av_scan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable AntiVirus scan service.",
							Optional:    true,
							Computed:    true,
						},
						"content_disarm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan.",
							Optional:    true,
							Computed:    true,
						},
						"emulator": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the virus emulator.",
							Optional:    true,
							Computed:    true,
						},
						"executables": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"default", "virus"}, false),

							Description: "Treat Windows executable files as viruses for the purpose of blocking or monitoring.",
							Optional:    true,
							Computed:    true,
						},
						"external_blocklist": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Optional:    true,
							Computed:    true,
						},
						"fortiai": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable/disable scanning of files by FortiAI.",
							Optional:    true,
							Computed:    true,
						},
						"fortindr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiNDR.",
							Optional:    true,
							Computed:    true,
						},
						"fortisandbox": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiSandbox.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable/disable SMTP AntiVirus scanning, monitoring, and quarantine.",
							Optional:         true,
							Computed:         true,
						},
						"outbreak_prevention": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable virus outbreak prevention service.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable quarantine for infected files.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ssh": {
				Type:        schema.TypeList,
				Description: "Configure SFTP and SCP AntiVirus options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to block.",
							Optional:         true,
							Computed:         true,
						},
						"archive_log": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Select the archive types to log.",
							Optional:         true,
							Computed:         true,
						},
						"av_scan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable AntiVirus scan service.",
							Optional:    true,
							Computed:    true,
						},
						"emulator": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the virus emulator.",
							Optional:    true,
							Computed:    true,
						},
						"external_blocklist": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Optional:    true,
							Computed:    true,
						},
						"fortiai": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable/disable scanning of files by FortiAI.",
							Optional:    true,
							Computed:    true,
						},
						"fortindr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiNDR.",
							Optional:    true,
							Computed:    true,
						},
						"fortisandbox": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable scanning of files by FortiSandbox.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable/disable SFTP and SCP AntiVirus scanning, monitoring, and quarantine.",
							Optional:         true,
							Computed:         true,
						},
						"outbreak_prevention": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

							Description: "Enable virus outbreak prevention service.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable quarantine for infected files.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceAntivirusProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating AntivirusProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectAntivirusProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateAntivirusProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AntivirusProfile")
	}

	return resourceAntivirusProfileRead(ctx, d, meta)
}

func resourceAntivirusProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectAntivirusProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateAntivirusProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating AntivirusProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AntivirusProfile")
	}

	return resourceAntivirusProfileRead(ctx, d, meta)
}

func resourceAntivirusProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteAntivirusProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting AntivirusProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadAntivirusProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AntivirusProfile resource: %v", err)
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

	diags := refreshObjectAntivirusProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenAntivirusProfileCifs(d *schema.ResourceData, v *models.AntivirusProfileCifs, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.AntivirusProfileCifs{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ArchiveBlock; tmp != nil {
				v["archive_block"] = *tmp
			}

			if tmp := cfg.ArchiveLog; tmp != nil {
				v["archive_log"] = *tmp
			}

			if tmp := cfg.AvScan; tmp != nil {
				v["av_scan"] = *tmp
			}

			if tmp := cfg.Emulator; tmp != nil {
				v["emulator"] = *tmp
			}

			if tmp := cfg.ExternalBlocklist; tmp != nil {
				v["external_blocklist"] = *tmp
			}

			if tmp := cfg.Fortiai; tmp != nil {
				v["fortiai"] = *tmp
			}

			if tmp := cfg.Fortindr; tmp != nil {
				v["fortindr"] = *tmp
			}

			if tmp := cfg.Fortisandbox; tmp != nil {
				v["fortisandbox"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OutbreakPrevention; tmp != nil {
				v["outbreak_prevention"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenAntivirusProfileContentDisarm(d *schema.ResourceData, v *models.AntivirusProfileContentDisarm, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.AntivirusProfileContentDisarm{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.CoverPage; tmp != nil {
				v["cover_page"] = *tmp
			}

			if tmp := cfg.DetectOnly; tmp != nil {
				v["detect_only"] = *tmp
			}

			if tmp := cfg.ErrorAction; tmp != nil {
				v["error_action"] = *tmp
			}

			if tmp := cfg.OfficeAction; tmp != nil {
				v["office_action"] = *tmp
			}

			if tmp := cfg.OfficeDde; tmp != nil {
				v["office_dde"] = *tmp
			}

			if tmp := cfg.OfficeEmbed; tmp != nil {
				v["office_embed"] = *tmp
			}

			if tmp := cfg.OfficeHylink; tmp != nil {
				v["office_hylink"] = *tmp
			}

			if tmp := cfg.OfficeLinked; tmp != nil {
				v["office_linked"] = *tmp
			}

			if tmp := cfg.OfficeMacro; tmp != nil {
				v["office_macro"] = *tmp
			}

			if tmp := cfg.OriginalFileDestination; tmp != nil {
				v["original_file_destination"] = *tmp
			}

			if tmp := cfg.PdfActForm; tmp != nil {
				v["pdf_act_form"] = *tmp
			}

			if tmp := cfg.PdfActGotor; tmp != nil {
				v["pdf_act_gotor"] = *tmp
			}

			if tmp := cfg.PdfActJava; tmp != nil {
				v["pdf_act_java"] = *tmp
			}

			if tmp := cfg.PdfActLaunch; tmp != nil {
				v["pdf_act_launch"] = *tmp
			}

			if tmp := cfg.PdfActMovie; tmp != nil {
				v["pdf_act_movie"] = *tmp
			}

			if tmp := cfg.PdfActSound; tmp != nil {
				v["pdf_act_sound"] = *tmp
			}

			if tmp := cfg.PdfEmbedfile; tmp != nil {
				v["pdf_embedfile"] = *tmp
			}

			if tmp := cfg.PdfHyperlink; tmp != nil {
				v["pdf_hyperlink"] = *tmp
			}

			if tmp := cfg.PdfJavacode; tmp != nil {
				v["pdf_javacode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenAntivirusProfileExternalBlocklist(d *schema.ResourceData, v *[]models.AntivirusProfileExternalBlocklist, prefix string, sort bool) interface{} {
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

func flattenAntivirusProfileFtp(d *schema.ResourceData, v *models.AntivirusProfileFtp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.AntivirusProfileFtp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ArchiveBlock; tmp != nil {
				v["archive_block"] = *tmp
			}

			if tmp := cfg.ArchiveLog; tmp != nil {
				v["archive_log"] = *tmp
			}

			if tmp := cfg.AvScan; tmp != nil {
				v["av_scan"] = *tmp
			}

			if tmp := cfg.Emulator; tmp != nil {
				v["emulator"] = *tmp
			}

			if tmp := cfg.ExternalBlocklist; tmp != nil {
				v["external_blocklist"] = *tmp
			}

			if tmp := cfg.Fortiai; tmp != nil {
				v["fortiai"] = *tmp
			}

			if tmp := cfg.Fortindr; tmp != nil {
				v["fortindr"] = *tmp
			}

			if tmp := cfg.Fortisandbox; tmp != nil {
				v["fortisandbox"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OutbreakPrevention; tmp != nil {
				v["outbreak_prevention"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenAntivirusProfileHttp(d *schema.ResourceData, v *models.AntivirusProfileHttp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.AntivirusProfileHttp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ArchiveBlock; tmp != nil {
				v["archive_block"] = *tmp
			}

			if tmp := cfg.ArchiveLog; tmp != nil {
				v["archive_log"] = *tmp
			}

			if tmp := cfg.AvScan; tmp != nil {
				v["av_scan"] = *tmp
			}

			if tmp := cfg.ContentDisarm; tmp != nil {
				v["content_disarm"] = *tmp
			}

			if tmp := cfg.Emulator; tmp != nil {
				v["emulator"] = *tmp
			}

			if tmp := cfg.ExternalBlocklist; tmp != nil {
				v["external_blocklist"] = *tmp
			}

			if tmp := cfg.Fortiai; tmp != nil {
				v["fortiai"] = *tmp
			}

			if tmp := cfg.Fortindr; tmp != nil {
				v["fortindr"] = *tmp
			}

			if tmp := cfg.Fortisandbox; tmp != nil {
				v["fortisandbox"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OutbreakPrevention; tmp != nil {
				v["outbreak_prevention"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenAntivirusProfileImap(d *schema.ResourceData, v *models.AntivirusProfileImap, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.AntivirusProfileImap{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ArchiveBlock; tmp != nil {
				v["archive_block"] = *tmp
			}

			if tmp := cfg.ArchiveLog; tmp != nil {
				v["archive_log"] = *tmp
			}

			if tmp := cfg.AvScan; tmp != nil {
				v["av_scan"] = *tmp
			}

			if tmp := cfg.ContentDisarm; tmp != nil {
				v["content_disarm"] = *tmp
			}

			if tmp := cfg.Emulator; tmp != nil {
				v["emulator"] = *tmp
			}

			if tmp := cfg.Executables; tmp != nil {
				v["executables"] = *tmp
			}

			if tmp := cfg.ExternalBlocklist; tmp != nil {
				v["external_blocklist"] = *tmp
			}

			if tmp := cfg.Fortiai; tmp != nil {
				v["fortiai"] = *tmp
			}

			if tmp := cfg.Fortindr; tmp != nil {
				v["fortindr"] = *tmp
			}

			if tmp := cfg.Fortisandbox; tmp != nil {
				v["fortisandbox"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OutbreakPrevention; tmp != nil {
				v["outbreak_prevention"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenAntivirusProfileMapi(d *schema.ResourceData, v *models.AntivirusProfileMapi, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.AntivirusProfileMapi{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ArchiveBlock; tmp != nil {
				v["archive_block"] = *tmp
			}

			if tmp := cfg.ArchiveLog; tmp != nil {
				v["archive_log"] = *tmp
			}

			if tmp := cfg.AvScan; tmp != nil {
				v["av_scan"] = *tmp
			}

			if tmp := cfg.Emulator; tmp != nil {
				v["emulator"] = *tmp
			}

			if tmp := cfg.Executables; tmp != nil {
				v["executables"] = *tmp
			}

			if tmp := cfg.ExternalBlocklist; tmp != nil {
				v["external_blocklist"] = *tmp
			}

			if tmp := cfg.Fortiai; tmp != nil {
				v["fortiai"] = *tmp
			}

			if tmp := cfg.Fortindr; tmp != nil {
				v["fortindr"] = *tmp
			}

			if tmp := cfg.Fortisandbox; tmp != nil {
				v["fortisandbox"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OutbreakPrevention; tmp != nil {
				v["outbreak_prevention"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenAntivirusProfileNacQuar(d *schema.ResourceData, v *models.AntivirusProfileNacQuar, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.AntivirusProfileNacQuar{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Expiry; tmp != nil {
				v["expiry"] = *tmp
			}

			if tmp := cfg.Infected; tmp != nil {
				v["infected"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenAntivirusProfileNntp(d *schema.ResourceData, v *models.AntivirusProfileNntp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.AntivirusProfileNntp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ArchiveBlock; tmp != nil {
				v["archive_block"] = *tmp
			}

			if tmp := cfg.ArchiveLog; tmp != nil {
				v["archive_log"] = *tmp
			}

			if tmp := cfg.AvScan; tmp != nil {
				v["av_scan"] = *tmp
			}

			if tmp := cfg.Emulator; tmp != nil {
				v["emulator"] = *tmp
			}

			if tmp := cfg.ExternalBlocklist; tmp != nil {
				v["external_blocklist"] = *tmp
			}

			if tmp := cfg.Fortiai; tmp != nil {
				v["fortiai"] = *tmp
			}

			if tmp := cfg.Fortindr; tmp != nil {
				v["fortindr"] = *tmp
			}

			if tmp := cfg.Fortisandbox; tmp != nil {
				v["fortisandbox"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OutbreakPrevention; tmp != nil {
				v["outbreak_prevention"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenAntivirusProfileOutbreakPrevention(d *schema.ResourceData, v *models.AntivirusProfileOutbreakPrevention, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.AntivirusProfileOutbreakPrevention{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ExternalBlocklist; tmp != nil {
				v["external_blocklist"] = *tmp
			}

			if tmp := cfg.FtgdService; tmp != nil {
				v["ftgd_service"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenAntivirusProfilePop3(d *schema.ResourceData, v *models.AntivirusProfilePop3, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.AntivirusProfilePop3{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ArchiveBlock; tmp != nil {
				v["archive_block"] = *tmp
			}

			if tmp := cfg.ArchiveLog; tmp != nil {
				v["archive_log"] = *tmp
			}

			if tmp := cfg.AvScan; tmp != nil {
				v["av_scan"] = *tmp
			}

			if tmp := cfg.ContentDisarm; tmp != nil {
				v["content_disarm"] = *tmp
			}

			if tmp := cfg.Emulator; tmp != nil {
				v["emulator"] = *tmp
			}

			if tmp := cfg.Executables; tmp != nil {
				v["executables"] = *tmp
			}

			if tmp := cfg.ExternalBlocklist; tmp != nil {
				v["external_blocklist"] = *tmp
			}

			if tmp := cfg.Fortiai; tmp != nil {
				v["fortiai"] = *tmp
			}

			if tmp := cfg.Fortindr; tmp != nil {
				v["fortindr"] = *tmp
			}

			if tmp := cfg.Fortisandbox; tmp != nil {
				v["fortisandbox"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OutbreakPrevention; tmp != nil {
				v["outbreak_prevention"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenAntivirusProfileSmtp(d *schema.ResourceData, v *models.AntivirusProfileSmtp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.AntivirusProfileSmtp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ArchiveBlock; tmp != nil {
				v["archive_block"] = *tmp
			}

			if tmp := cfg.ArchiveLog; tmp != nil {
				v["archive_log"] = *tmp
			}

			if tmp := cfg.AvScan; tmp != nil {
				v["av_scan"] = *tmp
			}

			if tmp := cfg.ContentDisarm; tmp != nil {
				v["content_disarm"] = *tmp
			}

			if tmp := cfg.Emulator; tmp != nil {
				v["emulator"] = *tmp
			}

			if tmp := cfg.Executables; tmp != nil {
				v["executables"] = *tmp
			}

			if tmp := cfg.ExternalBlocklist; tmp != nil {
				v["external_blocklist"] = *tmp
			}

			if tmp := cfg.Fortiai; tmp != nil {
				v["fortiai"] = *tmp
			}

			if tmp := cfg.Fortindr; tmp != nil {
				v["fortindr"] = *tmp
			}

			if tmp := cfg.Fortisandbox; tmp != nil {
				v["fortisandbox"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OutbreakPrevention; tmp != nil {
				v["outbreak_prevention"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenAntivirusProfileSsh(d *schema.ResourceData, v *models.AntivirusProfileSsh, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.AntivirusProfileSsh{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ArchiveBlock; tmp != nil {
				v["archive_block"] = *tmp
			}

			if tmp := cfg.ArchiveLog; tmp != nil {
				v["archive_log"] = *tmp
			}

			if tmp := cfg.AvScan; tmp != nil {
				v["av_scan"] = *tmp
			}

			if tmp := cfg.Emulator; tmp != nil {
				v["emulator"] = *tmp
			}

			if tmp := cfg.ExternalBlocklist; tmp != nil {
				v["external_blocklist"] = *tmp
			}

			if tmp := cfg.Fortiai; tmp != nil {
				v["fortiai"] = *tmp
			}

			if tmp := cfg.Fortindr; tmp != nil {
				v["fortindr"] = *tmp
			}

			if tmp := cfg.Fortisandbox; tmp != nil {
				v["fortisandbox"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OutbreakPrevention; tmp != nil {
				v["outbreak_prevention"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectAntivirusProfile(d *schema.ResourceData, o *models.AntivirusProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AnalyticsAcceptFiletype != nil {
		v := *o.AnalyticsAcceptFiletype

		if err = d.Set("analytics_accept_filetype", v); err != nil {
			return diag.Errorf("error reading analytics_accept_filetype: %v", err)
		}
	}

	if o.AnalyticsBlFiletype != nil {
		v := *o.AnalyticsBlFiletype

		if err = d.Set("analytics_bl_filetype", v); err != nil {
			return diag.Errorf("error reading analytics_bl_filetype: %v", err)
		}
	}

	if o.AnalyticsDb != nil {
		v := *o.AnalyticsDb

		if err = d.Set("analytics_db", v); err != nil {
			return diag.Errorf("error reading analytics_db: %v", err)
		}
	}

	if o.AnalyticsIgnoreFiletype != nil {
		v := *o.AnalyticsIgnoreFiletype

		if err = d.Set("analytics_ignore_filetype", v); err != nil {
			return diag.Errorf("error reading analytics_ignore_filetype: %v", err)
		}
	}

	if o.AnalyticsMaxUpload != nil {
		v := *o.AnalyticsMaxUpload

		if err = d.Set("analytics_max_upload", v); err != nil {
			return diag.Errorf("error reading analytics_max_upload: %v", err)
		}
	}

	if o.AnalyticsWlFiletype != nil {
		v := *o.AnalyticsWlFiletype

		if err = d.Set("analytics_wl_filetype", v); err != nil {
			return diag.Errorf("error reading analytics_wl_filetype: %v", err)
		}
	}

	if o.AvBlockLog != nil {
		v := *o.AvBlockLog

		if err = d.Set("av_block_log", v); err != nil {
			return diag.Errorf("error reading av_block_log: %v", err)
		}
	}

	if o.AvVirusLog != nil {
		v := *o.AvVirusLog

		if err = d.Set("av_virus_log", v); err != nil {
			return diag.Errorf("error reading av_virus_log: %v", err)
		}
	}

	if _, ok := d.GetOk("cifs"); ok {
		if o.Cifs != nil {
			if err = d.Set("cifs", flattenAntivirusProfileCifs(d, o.Cifs, "cifs", sort)); err != nil {
				return diag.Errorf("error reading cifs: %v", err)
			}
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if _, ok := d.GetOk("content_disarm"); ok {
		if o.ContentDisarm != nil {
			if err = d.Set("content_disarm", flattenAntivirusProfileContentDisarm(d, o.ContentDisarm, "content_disarm", sort)); err != nil {
				return diag.Errorf("error reading content_disarm: %v", err)
			}
		}
	}

	if o.EmsThreatFeed != nil {
		v := *o.EmsThreatFeed

		if err = d.Set("ems_threat_feed", v); err != nil {
			return diag.Errorf("error reading ems_threat_feed: %v", err)
		}
	}

	if o.ExtendedLog != nil {
		v := *o.ExtendedLog

		if err = d.Set("extended_log", v); err != nil {
			return diag.Errorf("error reading extended_log: %v", err)
		}
	}

	if o.ExternalBlocklist != nil {
		if err = d.Set("external_blocklist", flattenAntivirusProfileExternalBlocklist(d, o.ExternalBlocklist, "external_blocklist", sort)); err != nil {
			return diag.Errorf("error reading external_blocklist: %v", err)
		}
	}

	if o.ExternalBlocklistArchiveScan != nil {
		v := *o.ExternalBlocklistArchiveScan

		if err = d.Set("external_blocklist_archive_scan", v); err != nil {
			return diag.Errorf("error reading external_blocklist_archive_scan: %v", err)
		}
	}

	if o.ExternalBlocklistEnableAll != nil {
		v := *o.ExternalBlocklistEnableAll

		if err = d.Set("external_blocklist_enable_all", v); err != nil {
			return diag.Errorf("error reading external_blocklist_enable_all: %v", err)
		}
	}

	if o.FeatureSet != nil {
		v := *o.FeatureSet

		if err = d.Set("feature_set", v); err != nil {
			return diag.Errorf("error reading feature_set: %v", err)
		}
	}

	if o.FortiaiErrorAction != nil {
		v := *o.FortiaiErrorAction

		if err = d.Set("fortiai_error_action", v); err != nil {
			return diag.Errorf("error reading fortiai_error_action: %v", err)
		}
	}

	if o.FortiaiTimeoutAction != nil {
		v := *o.FortiaiTimeoutAction

		if err = d.Set("fortiai_timeout_action", v); err != nil {
			return diag.Errorf("error reading fortiai_timeout_action: %v", err)
		}
	}

	if o.FortindrErrorAction != nil {
		v := *o.FortindrErrorAction

		if err = d.Set("fortindr_error_action", v); err != nil {
			return diag.Errorf("error reading fortindr_error_action: %v", err)
		}
	}

	if o.FortindrTimeoutAction != nil {
		v := *o.FortindrTimeoutAction

		if err = d.Set("fortindr_timeout_action", v); err != nil {
			return diag.Errorf("error reading fortindr_timeout_action: %v", err)
		}
	}

	if o.FortisandboxErrorAction != nil {
		v := *o.FortisandboxErrorAction

		if err = d.Set("fortisandbox_error_action", v); err != nil {
			return diag.Errorf("error reading fortisandbox_error_action: %v", err)
		}
	}

	if o.FortisandboxMaxUpload != nil {
		v := *o.FortisandboxMaxUpload

		if err = d.Set("fortisandbox_max_upload", v); err != nil {
			return diag.Errorf("error reading fortisandbox_max_upload: %v", err)
		}
	}

	if o.FortisandboxMode != nil {
		v := *o.FortisandboxMode

		if err = d.Set("fortisandbox_mode", v); err != nil {
			return diag.Errorf("error reading fortisandbox_mode: %v", err)
		}
	}

	if o.FortisandboxTimeoutAction != nil {
		v := *o.FortisandboxTimeoutAction

		if err = d.Set("fortisandbox_timeout_action", v); err != nil {
			return diag.Errorf("error reading fortisandbox_timeout_action: %v", err)
		}
	}

	if o.FtgdAnalytics != nil {
		v := *o.FtgdAnalytics

		if err = d.Set("ftgd_analytics", v); err != nil {
			return diag.Errorf("error reading ftgd_analytics: %v", err)
		}
	}

	if _, ok := d.GetOk("ftp"); ok {
		if o.Ftp != nil {
			if err = d.Set("ftp", flattenAntivirusProfileFtp(d, o.Ftp, "ftp", sort)); err != nil {
				return diag.Errorf("error reading ftp: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("http"); ok {
		if o.Http != nil {
			if err = d.Set("http", flattenAntivirusProfileHttp(d, o.Http, "http", sort)); err != nil {
				return diag.Errorf("error reading http: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("imap"); ok {
		if o.Imap != nil {
			if err = d.Set("imap", flattenAntivirusProfileImap(d, o.Imap, "imap", sort)); err != nil {
				return diag.Errorf("error reading imap: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("mapi"); ok {
		if o.Mapi != nil {
			if err = d.Set("mapi", flattenAntivirusProfileMapi(d, o.Mapi, "mapi", sort)); err != nil {
				return diag.Errorf("error reading mapi: %v", err)
			}
		}
	}

	if o.MobileMalwareDb != nil {
		v := *o.MobileMalwareDb

		if err = d.Set("mobile_malware_db", v); err != nil {
			return diag.Errorf("error reading mobile_malware_db: %v", err)
		}
	}

	if _, ok := d.GetOk("nac_quar"); ok {
		if o.NacQuar != nil {
			if err = d.Set("nac_quar", flattenAntivirusProfileNacQuar(d, o.NacQuar, "nac_quar", sort)); err != nil {
				return diag.Errorf("error reading nac_quar: %v", err)
			}
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if _, ok := d.GetOk("nntp"); ok {
		if o.Nntp != nil {
			if err = d.Set("nntp", flattenAntivirusProfileNntp(d, o.Nntp, "nntp", sort)); err != nil {
				return diag.Errorf("error reading nntp: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("outbreak_prevention"); ok {
		if o.OutbreakPrevention != nil {
			if err = d.Set("outbreak_prevention", flattenAntivirusProfileOutbreakPrevention(d, o.OutbreakPrevention, "outbreak_prevention", sort)); err != nil {
				return diag.Errorf("error reading outbreak_prevention: %v", err)
			}
		}
	}

	if o.OutbreakPreventionArchiveScan != nil {
		v := *o.OutbreakPreventionArchiveScan

		if err = d.Set("outbreak_prevention_archive_scan", v); err != nil {
			return diag.Errorf("error reading outbreak_prevention_archive_scan: %v", err)
		}
	}

	if _, ok := d.GetOk("pop3"); ok {
		if o.Pop3 != nil {
			if err = d.Set("pop3", flattenAntivirusProfilePop3(d, o.Pop3, "pop3", sort)); err != nil {
				return diag.Errorf("error reading pop3: %v", err)
			}
		}
	}

	if o.ReplacemsgGroup != nil {
		v := *o.ReplacemsgGroup

		if err = d.Set("replacemsg_group", v); err != nil {
			return diag.Errorf("error reading replacemsg_group: %v", err)
		}
	}

	if o.ScanMode != nil {
		v := *o.ScanMode

		if err = d.Set("scan_mode", v); err != nil {
			return diag.Errorf("error reading scan_mode: %v", err)
		}
	}

	if _, ok := d.GetOk("smtp"); ok {
		if o.Smtp != nil {
			if err = d.Set("smtp", flattenAntivirusProfileSmtp(d, o.Smtp, "smtp", sort)); err != nil {
				return diag.Errorf("error reading smtp: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("ssh"); ok {
		if o.Ssh != nil {
			if err = d.Set("ssh", flattenAntivirusProfileSsh(d, o.Ssh, "ssh", sort)); err != nil {
				return diag.Errorf("error reading ssh: %v", err)
			}
		}
	}

	return nil
}

func expandAntivirusProfileCifs(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.AntivirusProfileCifs, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfileCifs

	for i := range l {
		tmp := models.AntivirusProfileCifs{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.archive_block", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveBlock = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.archive_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.av_scan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AvScan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.emulator", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Emulator = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.external_blocklist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExternalBlocklist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortiai", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortiai = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortindr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortindr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortisandbox", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortisandbox = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.outbreak_prevention", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OutbreakPrevention = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandAntivirusProfileContentDisarm(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.AntivirusProfileContentDisarm, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfileContentDisarm

	for i := range l {
		tmp := models.AntivirusProfileContentDisarm{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cover_page", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CoverPage = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.detect_only", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DetectOnly = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.error_action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ErrorAction = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.office_action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OfficeAction = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.office_dde", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OfficeDde = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.office_embed", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OfficeEmbed = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.office_hylink", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OfficeHylink = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.office_linked", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OfficeLinked = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.office_macro", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OfficeMacro = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.original_file_destination", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OriginalFileDestination = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pdf_act_form", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PdfActForm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pdf_act_gotor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PdfActGotor = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pdf_act_java", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PdfActJava = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pdf_act_launch", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PdfActLaunch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pdf_act_movie", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PdfActMovie = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pdf_act_sound", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PdfActSound = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pdf_embedfile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PdfEmbedfile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pdf_hyperlink", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PdfHyperlink = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pdf_javacode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PdfJavacode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandAntivirusProfileExternalBlocklist(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.AntivirusProfileExternalBlocklist, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfileExternalBlocklist

	for i := range l {
		tmp := models.AntivirusProfileExternalBlocklist{}
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

func expandAntivirusProfileFtp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.AntivirusProfileFtp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfileFtp

	for i := range l {
		tmp := models.AntivirusProfileFtp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.archive_block", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveBlock = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.archive_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.av_scan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AvScan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.emulator", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Emulator = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.external_blocklist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExternalBlocklist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortiai", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortiai = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortindr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortindr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortisandbox", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortisandbox = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.outbreak_prevention", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OutbreakPrevention = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandAntivirusProfileHttp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.AntivirusProfileHttp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfileHttp

	for i := range l {
		tmp := models.AntivirusProfileHttp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.archive_block", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveBlock = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.archive_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.av_scan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AvScan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.content_disarm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ContentDisarm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.emulator", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Emulator = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.external_blocklist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExternalBlocklist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortiai", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortiai = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortindr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortindr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortisandbox", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortisandbox = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.outbreak_prevention", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OutbreakPrevention = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandAntivirusProfileImap(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.AntivirusProfileImap, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfileImap

	for i := range l {
		tmp := models.AntivirusProfileImap{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.archive_block", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveBlock = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.archive_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.av_scan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AvScan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.content_disarm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ContentDisarm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.emulator", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Emulator = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.executables", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Executables = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.external_blocklist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExternalBlocklist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortiai", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortiai = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortindr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortindr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortisandbox", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortisandbox = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.outbreak_prevention", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OutbreakPrevention = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandAntivirusProfileMapi(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.AntivirusProfileMapi, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfileMapi

	for i := range l {
		tmp := models.AntivirusProfileMapi{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.archive_block", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveBlock = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.archive_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.av_scan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AvScan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.emulator", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Emulator = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.executables", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Executables = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.external_blocklist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExternalBlocklist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortiai", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortiai = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortindr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortindr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortisandbox", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortisandbox = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.outbreak_prevention", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OutbreakPrevention = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandAntivirusProfileNacQuar(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.AntivirusProfileNacQuar, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfileNacQuar

	for i := range l {
		tmp := models.AntivirusProfileNacQuar{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.expiry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Expiry = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.infected", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Infected = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandAntivirusProfileNntp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.AntivirusProfileNntp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfileNntp

	for i := range l {
		tmp := models.AntivirusProfileNntp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.archive_block", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveBlock = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.archive_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.av_scan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AvScan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.emulator", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Emulator = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.external_blocklist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExternalBlocklist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortiai", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortiai = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortindr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortindr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortisandbox", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortisandbox = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.outbreak_prevention", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OutbreakPrevention = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandAntivirusProfileOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.AntivirusProfileOutbreakPrevention, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfileOutbreakPrevention

	for i := range l {
		tmp := models.AntivirusProfileOutbreakPrevention{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.external_blocklist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExternalBlocklist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ftgd_service", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FtgdService = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandAntivirusProfilePop3(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.AntivirusProfilePop3, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfilePop3

	for i := range l {
		tmp := models.AntivirusProfilePop3{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.archive_block", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveBlock = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.archive_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.av_scan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AvScan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.content_disarm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ContentDisarm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.emulator", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Emulator = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.executables", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Executables = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.external_blocklist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExternalBlocklist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortiai", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortiai = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortindr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortindr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortisandbox", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortisandbox = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.outbreak_prevention", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OutbreakPrevention = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandAntivirusProfileSmtp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.AntivirusProfileSmtp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfileSmtp

	for i := range l {
		tmp := models.AntivirusProfileSmtp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.archive_block", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveBlock = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.archive_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.av_scan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AvScan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.content_disarm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ContentDisarm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.emulator", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Emulator = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.executables", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Executables = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.external_blocklist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExternalBlocklist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortiai", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortiai = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortindr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortindr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortisandbox", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortisandbox = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.outbreak_prevention", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OutbreakPrevention = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandAntivirusProfileSsh(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.AntivirusProfileSsh, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AntivirusProfileSsh

	for i := range l {
		tmp := models.AntivirusProfileSsh{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.archive_block", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveBlock = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.archive_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArchiveLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.av_scan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AvScan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.emulator", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Emulator = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.external_blocklist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExternalBlocklist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortiai", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortiai = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortindr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortindr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortisandbox", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortisandbox = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.outbreak_prevention", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OutbreakPrevention = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func getObjectAntivirusProfile(d *schema.ResourceData, sv string) (*models.AntivirusProfile, diag.Diagnostics) {
	obj := models.AntivirusProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("analytics_accept_filetype"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("analytics_accept_filetype", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AnalyticsAcceptFiletype = &tmp
		}
	}
	if v1, ok := d.GetOk("analytics_bl_filetype"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("analytics_bl_filetype", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AnalyticsBlFiletype = &tmp
		}
	}
	if v1, ok := d.GetOk("analytics_db"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("analytics_db", sv)
				diags = append(diags, e)
			}
			obj.AnalyticsDb = &v2
		}
	}
	if v1, ok := d.GetOk("analytics_ignore_filetype"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("analytics_ignore_filetype", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AnalyticsIgnoreFiletype = &tmp
		}
	}
	if v1, ok := d.GetOk("analytics_max_upload"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.2.0") {
				e := utils.AttributeVersionWarning("analytics_max_upload", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AnalyticsMaxUpload = &tmp
		}
	}
	if v1, ok := d.GetOk("analytics_wl_filetype"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("analytics_wl_filetype", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AnalyticsWlFiletype = &tmp
		}
	}
	if v1, ok := d.GetOk("av_block_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("av_block_log", sv)
				diags = append(diags, e)
			}
			obj.AvBlockLog = &v2
		}
	}
	if v1, ok := d.GetOk("av_virus_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("av_virus_log", sv)
				diags = append(diags, e)
			}
			obj.AvVirusLog = &v2
		}
	}
	if v, ok := d.GetOk("cifs"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("cifs", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfileCifs(d, v, "cifs", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Cifs = t
		}
	} else if d.HasChange("cifs") {
		old, new := d.GetChange("cifs")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Cifs = &models.AntivirusProfileCifs{}
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
	if v, ok := d.GetOk("content_disarm"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("content_disarm", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfileContentDisarm(d, v, "content_disarm", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ContentDisarm = t
		}
	} else if d.HasChange("content_disarm") {
		old, new := d.GetChange("content_disarm")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ContentDisarm = &models.AntivirusProfileContentDisarm{}
		}
	}
	if v1, ok := d.GetOk("ems_threat_feed"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ems_threat_feed", sv)
				diags = append(diags, e)
			}
			obj.EmsThreatFeed = &v2
		}
	}
	if v1, ok := d.GetOk("extended_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extended_log", sv)
				diags = append(diags, e)
			}
			obj.ExtendedLog = &v2
		}
	}
	if v, ok := d.GetOk("external_blocklist"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("external_blocklist", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfileExternalBlocklist(d, v, "external_blocklist", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ExternalBlocklist = t
		}
	} else if d.HasChange("external_blocklist") {
		old, new := d.GetChange("external_blocklist")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ExternalBlocklist = &[]models.AntivirusProfileExternalBlocklist{}
		}
	}
	if v1, ok := d.GetOk("external_blocklist_archive_scan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "v7.0.1") {
				e := utils.AttributeVersionWarning("external_blocklist_archive_scan", sv)
				diags = append(diags, e)
			}
			obj.ExternalBlocklistArchiveScan = &v2
		}
	}
	if v1, ok := d.GetOk("external_blocklist_enable_all"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("external_blocklist_enable_all", sv)
				diags = append(diags, e)
			}
			obj.ExternalBlocklistEnableAll = &v2
		}
	}
	if v1, ok := d.GetOk("feature_set"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("feature_set", sv)
				diags = append(diags, e)
			}
			obj.FeatureSet = &v2
		}
	}
	if v1, ok := d.GetOk("fortiai_error_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "v7.2.0") {
				e := utils.AttributeVersionWarning("fortiai_error_action", sv)
				diags = append(diags, e)
			}
			obj.FortiaiErrorAction = &v2
		}
	}
	if v1, ok := d.GetOk("fortiai_timeout_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "v7.2.0") {
				e := utils.AttributeVersionWarning("fortiai_timeout_action", sv)
				diags = append(diags, e)
			}
			obj.FortiaiTimeoutAction = &v2
		}
	}
	if v1, ok := d.GetOk("fortindr_error_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("fortindr_error_action", sv)
				diags = append(diags, e)
			}
			obj.FortindrErrorAction = &v2
		}
	}
	if v1, ok := d.GetOk("fortindr_timeout_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("fortindr_timeout_action", sv)
				diags = append(diags, e)
			}
			obj.FortindrTimeoutAction = &v2
		}
	}
	if v1, ok := d.GetOk("fortisandbox_error_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("fortisandbox_error_action", sv)
				diags = append(diags, e)
			}
			obj.FortisandboxErrorAction = &v2
		}
	}
	if v1, ok := d.GetOk("fortisandbox_max_upload"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("fortisandbox_max_upload", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FortisandboxMaxUpload = &tmp
		}
	}
	if v1, ok := d.GetOk("fortisandbox_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("fortisandbox_mode", sv)
				diags = append(diags, e)
			}
			obj.FortisandboxMode = &v2
		}
	}
	if v1, ok := d.GetOk("fortisandbox_timeout_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("fortisandbox_timeout_action", sv)
				diags = append(diags, e)
			}
			obj.FortisandboxTimeoutAction = &v2
		}
	}
	if v1, ok := d.GetOk("ftgd_analytics"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.0") {
				e := utils.AttributeVersionWarning("ftgd_analytics", sv)
				diags = append(diags, e)
			}
			obj.FtgdAnalytics = &v2
		}
	}
	if v, ok := d.GetOk("ftp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ftp", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfileFtp(d, v, "ftp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ftp = t
		}
	} else if d.HasChange("ftp") {
		old, new := d.GetChange("ftp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ftp = &models.AntivirusProfileFtp{}
		}
	}
	if v, ok := d.GetOk("http"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("http", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfileHttp(d, v, "http", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Http = t
		}
	} else if d.HasChange("http") {
		old, new := d.GetChange("http")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Http = &models.AntivirusProfileHttp{}
		}
	}
	if v, ok := d.GetOk("imap"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("imap", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfileImap(d, v, "imap", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Imap = t
		}
	} else if d.HasChange("imap") {
		old, new := d.GetChange("imap")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Imap = &models.AntivirusProfileImap{}
		}
	}
	if v, ok := d.GetOk("mapi"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("mapi", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfileMapi(d, v, "mapi", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Mapi = t
		}
	} else if d.HasChange("mapi") {
		old, new := d.GetChange("mapi")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Mapi = &models.AntivirusProfileMapi{}
		}
	}
	if v1, ok := d.GetOk("mobile_malware_db"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mobile_malware_db", sv)
				diags = append(diags, e)
			}
			obj.MobileMalwareDb = &v2
		}
	}
	if v, ok := d.GetOk("nac_quar"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("nac_quar", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfileNacQuar(d, v, "nac_quar", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NacQuar = t
		}
	} else if d.HasChange("nac_quar") {
		old, new := d.GetChange("nac_quar")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NacQuar = &models.AntivirusProfileNacQuar{}
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
	if v, ok := d.GetOk("nntp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("nntp", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfileNntp(d, v, "nntp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Nntp = t
		}
	} else if d.HasChange("nntp") {
		old, new := d.GetChange("nntp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Nntp = &models.AntivirusProfileNntp{}
		}
	}
	if v, ok := d.GetOk("outbreak_prevention"); ok {
		if !utils.CheckVer(sv, "", "v7.0.0") {
			e := utils.AttributeVersionWarning("outbreak_prevention", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfileOutbreakPrevention(d, v, "outbreak_prevention", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.OutbreakPrevention = t
		}
	} else if d.HasChange("outbreak_prevention") {
		old, new := d.GetChange("outbreak_prevention")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.OutbreakPrevention = &models.AntivirusProfileOutbreakPrevention{}
		}
	}
	if v1, ok := d.GetOk("outbreak_prevention_archive_scan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("outbreak_prevention_archive_scan", sv)
				diags = append(diags, e)
			}
			obj.OutbreakPreventionArchiveScan = &v2
		}
	}
	if v, ok := d.GetOk("pop3"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("pop3", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfilePop3(d, v, "pop3", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Pop3 = t
		}
	} else if d.HasChange("pop3") {
		old, new := d.GetChange("pop3")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Pop3 = &models.AntivirusProfilePop3{}
		}
	}
	if v1, ok := d.GetOk("replacemsg_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("replacemsg_group", sv)
				diags = append(diags, e)
			}
			obj.ReplacemsgGroup = &v2
		}
	}
	if v1, ok := d.GetOk("scan_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scan_mode", sv)
				diags = append(diags, e)
			}
			obj.ScanMode = &v2
		}
	}
	if v, ok := d.GetOk("smtp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("smtp", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfileSmtp(d, v, "smtp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Smtp = t
		}
	} else if d.HasChange("smtp") {
		old, new := d.GetChange("smtp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Smtp = &models.AntivirusProfileSmtp{}
		}
	}
	if v, ok := d.GetOk("ssh"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ssh", sv)
			diags = append(diags, e)
		}
		t, err := expandAntivirusProfileSsh(d, v, "ssh", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ssh = t
		}
	} else if d.HasChange("ssh") {
		old, new := d.GetChange("ssh")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ssh = &models.AntivirusProfileSsh{}
		}
	}
	return &obj, diags
}
