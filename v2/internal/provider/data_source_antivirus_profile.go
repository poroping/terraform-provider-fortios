// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure AntiVirus profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceAntivirusProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure AntiVirus profiles.",

		ReadContext: dataSourceAntivirusProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"analytics_accept_filetype": {
				Type:        schema.TypeInt,
				Description: "Only submit files matching this DLP file-pattern to FortiSandbox.",
				Computed:    true,
			},
			"analytics_bl_filetype": {
				Type:        schema.TypeInt,
				Description: "Only submit files matching this DLP file-pattern to FortiSandbox.",
				Computed:    true,
			},
			"analytics_db": {
				Type:        schema.TypeString,
				Description: "Enable/disable using the FortiSandbox signature database to supplement the AV signature databases.",
				Computed:    true,
			},
			"analytics_ignore_filetype": {
				Type:        schema.TypeInt,
				Description: "Do not submit files matching this DLP file-pattern to FortiSandbox.",
				Computed:    true,
			},
			"analytics_max_upload": {
				Type:        schema.TypeInt,
				Description: "Maximum size of files that can be uploaded to FortiSandbox.",
				Computed:    true,
			},
			"analytics_wl_filetype": {
				Type:        schema.TypeInt,
				Description: "Do not submit files matching this DLP file-pattern to FortiSandbox.",
				Computed:    true,
			},
			"av_block_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging for AntiVirus file blocking.",
				Computed:    true,
			},
			"av_virus_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable AntiVirus logging.",
				Computed:    true,
			},
			"cifs": {
				Type:        schema.TypeList,
				Description: "Configure CIFS AntiVirus options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type:        schema.TypeString,
							Description: "Select the archive types to block.",
							Computed:    true,
						},
						"archive_log": {
							Type:        schema.TypeString,
							Description: "Select the archive types to log.",
							Computed:    true,
						},
						"av_scan": {
							Type:        schema.TypeString,
							Description: "Enable AntiVirus scan service.",
							Computed:    true,
						},
						"emulator": {
							Type:        schema.TypeString,
							Description: "Enable/disable the virus emulator.",
							Computed:    true,
						},
						"external_blocklist": {
							Type:        schema.TypeString,
							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Computed:    true,
						},
						"fortiai": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of files by FortiAI server.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "Enable/disable CIFS AntiVirus scanning, monitoring, and quarantine.",
							Computed:    true,
						},
						"outbreak_prevention": {
							Type:        schema.TypeString,
							Description: "Enable virus outbreak prevention service.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Enable/disable quarantine for infected files.",
							Computed:    true,
						},
					},
				},
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"content_disarm": {
				Type:        schema.TypeList,
				Description: "AV Content Disarm and Reconstruction settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cover_page": {
							Type:        schema.TypeString,
							Description: "Enable/disable inserting a cover page into the disarmed document.",
							Computed:    true,
						},
						"detect_only": {
							Type:        schema.TypeString,
							Description: "Enable/disable only detect disarmable files, do not alter content.",
							Computed:    true,
						},
						"error_action": {
							Type:        schema.TypeString,
							Description: "Action to be taken if CDR engine encounters an unrecoverable error.",
							Computed:    true,
						},
						"office_action": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of PowerPoint action events in Microsoft Office documents.",
							Computed:    true,
						},
						"office_dde": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of Dynamic Data Exchange events in Microsoft Office documents.",
							Computed:    true,
						},
						"office_embed": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of embedded objects in Microsoft Office documents.",
							Computed:    true,
						},
						"office_hylink": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of hyperlinks in Microsoft Office documents.",
							Computed:    true,
						},
						"office_linked": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of linked objects in Microsoft Office documents.",
							Computed:    true,
						},
						"office_macro": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of macros in Microsoft Office documents.",
							Computed:    true,
						},
						"original_file_destination": {
							Type:        schema.TypeString,
							Description: "Destination to send original file if active content is removed.",
							Computed:    true,
						},
						"pdf_act_form": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of PDF document actions that submit data to other targets.",
							Computed:    true,
						},
						"pdf_act_gotor": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of PDF document actions that access other PDF documents.",
							Computed:    true,
						},
						"pdf_act_java": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of PDF document actions that execute JavaScript code.",
							Computed:    true,
						},
						"pdf_act_launch": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of PDF document actions that launch other applications.",
							Computed:    true,
						},
						"pdf_act_movie": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of PDF document actions that play a movie.",
							Computed:    true,
						},
						"pdf_act_sound": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of PDF document actions that play a sound.",
							Computed:    true,
						},
						"pdf_embedfile": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of embedded files in PDF documents.",
							Computed:    true,
						},
						"pdf_hyperlink": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of hyperlinks from PDF documents.",
							Computed:    true,
						},
						"pdf_javacode": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of JavaScript code in PDF documents.",
							Computed:    true,
						},
					},
				},
			},
			"ems_threat_feed": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of EMS threat feed when performing AntiVirus scan. Analyzes files including the content of archives.",
				Computed:    true,
			},
			"extended_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable extended logging for antivirus.",
				Computed:    true,
			},
			"external_blocklist": {
				Type:        schema.TypeList,
				Description: "One or more external malware block lists.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "External blocklist.",
							Computed:    true,
						},
					},
				},
			},
			"external_blocklist_archive_scan": {
				Type:        schema.TypeString,
				Description: "Enable/disable external-blocklist archive scanning.",
				Computed:    true,
			},
			"external_blocklist_enable_all": {
				Type:        schema.TypeString,
				Description: "Enable/disable all external blocklists.",
				Computed:    true,
			},
			"feature_set": {
				Type:        schema.TypeString,
				Description: "Flow/proxy feature set.",
				Computed:    true,
			},
			"fortiai_error_action": {
				Type:        schema.TypeString,
				Description: "Action to take if FortiAI encounters an error.",
				Computed:    true,
			},
			"fortiai_timeout_action": {
				Type:        schema.TypeString,
				Description: "Action to take if FortiAI encounters a scan timeout.",
				Computed:    true,
			},
			"ftgd_analytics": {
				Type:        schema.TypeString,
				Description: "Settings to control which files are uploaded to FortiSandbox.",
				Computed:    true,
			},
			"ftp": {
				Type:        schema.TypeList,
				Description: "Configure FTP AntiVirus options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type:        schema.TypeString,
							Description: "Select the archive types to block.",
							Computed:    true,
						},
						"archive_log": {
							Type:        schema.TypeString,
							Description: "Select the archive types to log.",
							Computed:    true,
						},
						"av_scan": {
							Type:        schema.TypeString,
							Description: "Enable AntiVirus scan service.",
							Computed:    true,
						},
						"emulator": {
							Type:        schema.TypeString,
							Description: "Enable/disable the virus emulator.",
							Computed:    true,
						},
						"external_blocklist": {
							Type:        schema.TypeString,
							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Computed:    true,
						},
						"fortiai": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of files by FortiAI server.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "Enable/disable FTP AntiVirus scanning, monitoring, and quarantine.",
							Computed:    true,
						},
						"outbreak_prevention": {
							Type:        schema.TypeString,
							Description: "Enable virus outbreak prevention service.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Enable/disable quarantine for infected files.",
							Computed:    true,
						},
					},
				},
			},
			"http": {
				Type:        schema.TypeList,
				Description: "Configure HTTP AntiVirus options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type:        schema.TypeString,
							Description: "Select the archive types to block.",
							Computed:    true,
						},
						"archive_log": {
							Type:        schema.TypeString,
							Description: "Select the archive types to log.",
							Computed:    true,
						},
						"av_scan": {
							Type:        schema.TypeString,
							Description: "Enable AntiVirus scan service.",
							Computed:    true,
						},
						"content_disarm": {
							Type:        schema.TypeString,
							Description: "Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan.",
							Computed:    true,
						},
						"emulator": {
							Type:        schema.TypeString,
							Description: "Enable/disable the virus emulator.",
							Computed:    true,
						},
						"external_blocklist": {
							Type:        schema.TypeString,
							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Computed:    true,
						},
						"fortiai": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of files by FortiAI server.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "Enable/disable HTTP AntiVirus scanning, monitoring, and quarantine.",
							Computed:    true,
						},
						"outbreak_prevention": {
							Type:        schema.TypeString,
							Description: "Enable virus outbreak prevention service.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Enable/disable quarantine for infected files.",
							Computed:    true,
						},
					},
				},
			},
			"imap": {
				Type:        schema.TypeList,
				Description: "Configure IMAP AntiVirus options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type:        schema.TypeString,
							Description: "Select the archive types to block.",
							Computed:    true,
						},
						"archive_log": {
							Type:        schema.TypeString,
							Description: "Select the archive types to log.",
							Computed:    true,
						},
						"av_scan": {
							Type:        schema.TypeString,
							Description: "Enable AntiVirus scan service.",
							Computed:    true,
						},
						"content_disarm": {
							Type:        schema.TypeString,
							Description: "Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan.",
							Computed:    true,
						},
						"emulator": {
							Type:        schema.TypeString,
							Description: "Enable/disable the virus emulator.",
							Computed:    true,
						},
						"executables": {
							Type:        schema.TypeString,
							Description: "Treat Windows executable files as viruses for the purpose of blocking or monitoring.",
							Computed:    true,
						},
						"external_blocklist": {
							Type:        schema.TypeString,
							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Computed:    true,
						},
						"fortiai": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of files by FortiAI server.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "Enable/disable IMAP AntiVirus scanning, monitoring, and quarantine.",
							Computed:    true,
						},
						"outbreak_prevention": {
							Type:        schema.TypeString,
							Description: "Enable virus outbreak prevention service.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Enable/disable quarantine for infected files.",
							Computed:    true,
						},
					},
				},
			},
			"mapi": {
				Type:        schema.TypeList,
				Description: "Configure MAPI AntiVirus options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type:        schema.TypeString,
							Description: "Select the archive types to block.",
							Computed:    true,
						},
						"archive_log": {
							Type:        schema.TypeString,
							Description: "Select the archive types to log.",
							Computed:    true,
						},
						"av_scan": {
							Type:        schema.TypeString,
							Description: "Enable AntiVirus scan service.",
							Computed:    true,
						},
						"emulator": {
							Type:        schema.TypeString,
							Description: "Enable/disable the virus emulator.",
							Computed:    true,
						},
						"executables": {
							Type:        schema.TypeString,
							Description: "Treat Windows executable files as viruses for the purpose of blocking or monitoring.",
							Computed:    true,
						},
						"external_blocklist": {
							Type:        schema.TypeString,
							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Computed:    true,
						},
						"fortiai": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of files by FortiAI server.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "Enable/disable MAPI AntiVirus scanning, monitoring, and quarantine.",
							Computed:    true,
						},
						"outbreak_prevention": {
							Type:        schema.TypeString,
							Description: "Enable virus outbreak prevention service.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Enable/disable quarantine for infected files.",
							Computed:    true,
						},
					},
				},
			},
			"mobile_malware_db": {
				Type:        schema.TypeString,
				Description: "Enable/disable using the mobile malware signature database.",
				Computed:    true,
			},
			"nac_quar": {
				Type:        schema.TypeList,
				Description: "Configure AntiVirus quarantine settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expiry": {
							Type:        schema.TypeString,
							Description: "Duration of quarantine.",
							Computed:    true,
						},
						"infected": {
							Type:        schema.TypeString,
							Description: "Enable/Disable quarantining infected hosts to the banned user list.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable AntiVirus quarantine logging.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Profile name.",
				Required:    true,
			},
			"nntp": {
				Type:        schema.TypeList,
				Description: "Configure NNTP AntiVirus options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type:        schema.TypeString,
							Description: "Select the archive types to block.",
							Computed:    true,
						},
						"archive_log": {
							Type:        schema.TypeString,
							Description: "Select the archive types to log.",
							Computed:    true,
						},
						"av_scan": {
							Type:        schema.TypeString,
							Description: "Enable AntiVirus scan service.",
							Computed:    true,
						},
						"emulator": {
							Type:        schema.TypeString,
							Description: "Enable/disable the virus emulator.",
							Computed:    true,
						},
						"external_blocklist": {
							Type:        schema.TypeString,
							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Computed:    true,
						},
						"fortiai": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of files by FortiAI server.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "Enable/disable NNTP AntiVirus scanning, monitoring, and quarantine.",
							Computed:    true,
						},
						"outbreak_prevention": {
							Type:        schema.TypeString,
							Description: "Enable virus outbreak prevention service.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Enable/disable quarantine for infected files.",
							Computed:    true,
						},
					},
				},
			},
			"outbreak_prevention": {
				Type:        schema.TypeList,
				Description: "Configure Virus Outbreak Prevention settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_blocklist": {
							Type:        schema.TypeString,
							Description: "Enable/disable external malware blocklist.",
							Computed:    true,
						},
						"ftgd_service": {
							Type:        schema.TypeString,
							Description: "Enable/disable FortiGuard Virus outbreak prevention service.",
							Computed:    true,
						},
					},
				},
			},
			"outbreak_prevention_archive_scan": {
				Type:        schema.TypeString,
				Description: "Enable/disable outbreak-prevention archive scanning.",
				Computed:    true,
			},
			"pop3": {
				Type:        schema.TypeList,
				Description: "Configure POP3 AntiVirus options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type:        schema.TypeString,
							Description: "Select the archive types to block.",
							Computed:    true,
						},
						"archive_log": {
							Type:        schema.TypeString,
							Description: "Select the archive types to log.",
							Computed:    true,
						},
						"av_scan": {
							Type:        schema.TypeString,
							Description: "Enable AntiVirus scan service.",
							Computed:    true,
						},
						"content_disarm": {
							Type:        schema.TypeString,
							Description: "Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan.",
							Computed:    true,
						},
						"emulator": {
							Type:        schema.TypeString,
							Description: "Enable/disable the virus emulator.",
							Computed:    true,
						},
						"executables": {
							Type:        schema.TypeString,
							Description: "Treat Windows executable files as viruses for the purpose of blocking or monitoring.",
							Computed:    true,
						},
						"external_blocklist": {
							Type:        schema.TypeString,
							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Computed:    true,
						},
						"fortiai": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of files by FortiAI server.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "Enable/disable POP3 AntiVirus scanning, monitoring, and quarantine.",
							Computed:    true,
						},
						"outbreak_prevention": {
							Type:        schema.TypeString,
							Description: "Enable virus outbreak prevention service.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Enable/disable quarantine for infected files.",
							Computed:    true,
						},
					},
				},
			},
			"replacemsg_group": {
				Type:        schema.TypeString,
				Description: "Replacement message group customized for this profile.",
				Computed:    true,
			},
			"scan_mode": {
				Type:        schema.TypeString,
				Description: "Choose between default scan mode and legacy scan mode. ",
				Computed:    true,
			},
			"smtp": {
				Type:        schema.TypeList,
				Description: "Configure SMTP AntiVirus options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type:        schema.TypeString,
							Description: "Select the archive types to block.",
							Computed:    true,
						},
						"archive_log": {
							Type:        schema.TypeString,
							Description: "Select the archive types to log.",
							Computed:    true,
						},
						"av_scan": {
							Type:        schema.TypeString,
							Description: "Enable AntiVirus scan service.",
							Computed:    true,
						},
						"content_disarm": {
							Type:        schema.TypeString,
							Description: "Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan.",
							Computed:    true,
						},
						"emulator": {
							Type:        schema.TypeString,
							Description: "Enable/disable the virus emulator.",
							Computed:    true,
						},
						"executables": {
							Type:        schema.TypeString,
							Description: "Treat Windows executable files as viruses for the purpose of blocking or monitoring.",
							Computed:    true,
						},
						"external_blocklist": {
							Type:        schema.TypeString,
							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Computed:    true,
						},
						"fortiai": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of files by FortiAI server.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "Enable/disable SMTP AntiVirus scanning, monitoring, and quarantine.",
							Computed:    true,
						},
						"outbreak_prevention": {
							Type:        schema.TypeString,
							Description: "Enable virus outbreak prevention service.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Enable/disable quarantine for infected files.",
							Computed:    true,
						},
					},
				},
			},
			"ssh": {
				Type:        schema.TypeList,
				Description: "Configure SFTP and SCP AntiVirus options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": {
							Type:        schema.TypeString,
							Description: "Select the archive types to block.",
							Computed:    true,
						},
						"archive_log": {
							Type:        schema.TypeString,
							Description: "Select the archive types to log.",
							Computed:    true,
						},
						"av_scan": {
							Type:        schema.TypeString,
							Description: "Enable AntiVirus scan service.",
							Computed:    true,
						},
						"emulator": {
							Type:        schema.TypeString,
							Description: "Enable/disable the virus emulator.",
							Computed:    true,
						},
						"external_blocklist": {
							Type:        schema.TypeString,
							Description: "Enable external-blocklist. Analyzes files including the content of archives.",
							Computed:    true,
						},
						"fortiai": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of files by FortiAI server.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "Enable/disable SFTP and SCP AntiVirus scanning, monitoring, and quarantine.",
							Computed:    true,
						},
						"outbreak_prevention": {
							Type:        schema.TypeString,
							Description: "Enable virus outbreak prevention service.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Enable/disable quarantine for infected files.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAntivirusProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadAntivirusProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AntivirusProfile dataSource: %v", err)
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

	diags := refreshObjectAntivirusProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
