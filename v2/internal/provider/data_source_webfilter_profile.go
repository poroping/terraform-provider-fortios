// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Web filter profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWebfilterProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Web filter profiles.",

		ReadContext: dataSourceWebfilterProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"antiphish": {
				Type:        schema.TypeList,
				Description: "AntiPhishing profile.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeString,
							Description: "Authentication methods.",
							Computed:    true,
						},
						"check_basic_auth": {
							Type:        schema.TypeString,
							Description: "Enable/disable checking of HTTP Basic Auth field for known credentials.",
							Computed:    true,
						},
						"check_uri": {
							Type:        schema.TypeString,
							Description: "Enable/disable checking of GET URI parameters for known credentials.",
							Computed:    true,
						},
						"check_username_only": {
							Type:        schema.TypeString,
							Description: "Enable/disable username only matching of credentials. Action will be taken for valid usernames regardless of password validity.",
							Computed:    true,
						},
						"custom_patterns": {
							Type:        schema.TypeList,
							Description: "Custom username and password regex patterns.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category": {
										Type:        schema.TypeString,
										Description: "Category that the pattern matches.",
										Computed:    true,
									},
									"pattern": {
										Type:        schema.TypeString,
										Description: "Target pattern.",
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "Pattern will be treated either as a regex pattern or literal string.",
										Computed:    true,
									},
								},
							},
						},
						"default_action": {
							Type:        schema.TypeString,
							Description: "Action to be taken when there is no matching rule.",
							Computed:    true,
						},
						"domain_controller": {
							Type:        schema.TypeString,
							Description: "Domain for which to verify received credentials against.",
							Computed:    true,
						},
						"inspection_entries": {
							Type:        schema.TypeList,
							Description: "AntiPhishing entries.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action to be taken upon an AntiPhishing match.",
										Computed:    true,
									},
									"fortiguard_category": {
										Type:        schema.TypeString,
										Description: "FortiGuard category to match.",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Inspection target name.",
										Computed:    true,
									},
								},
							},
						},
						"ldap": {
							Type:        schema.TypeString,
							Description: "LDAP server for which to verify received credentials against.",
							Computed:    true,
						},
						"max_body_len": {
							Type:        schema.TypeInt,
							Description: "Maximum size of a POST body to check for credentials.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Toggle AntiPhishing functionality.",
							Computed:    true,
						},
					},
				},
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"extended_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable extended logging for web filtering.",
				Computed:    true,
			},
			"feature_set": {
				Type:        schema.TypeString,
				Description: "Flow/proxy feature set.",
				Computed:    true,
			},
			"file_filter": {
				Type:        schema.TypeList,
				Description: "File filter.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entries": {
							Type:        schema.TypeList,
							Description: "File filter entries.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action taken for matched file.",
										Computed:    true,
									},
									"comment": {
										Type:        schema.TypeString,
										Description: "Comment.",
										Computed:    true,
									},
									"direction": {
										Type:        schema.TypeString,
										Description: "Match files transmitted in the session's originating or reply direction.",
										Computed:    true,
									},
									"file_type": {
										Type:        schema.TypeList,
										Description: "Select file type.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Description: "File type name.",
													Computed:    true,
												},
											},
										},
									},
									"filter": {
										Type:        schema.TypeString,
										Description: "Add a file filter.",
										Computed:    true,
									},
									"password_protected": {
										Type:        schema.TypeString,
										Description: "Match password-protected files.",
										Computed:    true,
									},
									"protocol": {
										Type:        schema.TypeString,
										Description: "Protocols to apply with.",
										Computed:    true,
									},
								},
							},
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable file filter logging.",
							Computed:    true,
						},
						"scan_archive_contents": {
							Type:        schema.TypeString,
							Description: "Enable/disable file filter archive contents scan.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable file filter.",
							Computed:    true,
						},
					},
				},
			},
			"ftgd_wf": {
				Type:        schema.TypeList,
				Description: "FortiGuard Web Filter settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exempt_quota": {
							Type:        schema.TypeString,
							Description: "Do not stop quota for these categories.",
							Computed:    true,
						},
						"filters": {
							Type:        schema.TypeList,
							Description: "FortiGuard filters.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action to take for matches.",
										Computed:    true,
									},
									"auth_usr_grp": {
										Type:        schema.TypeList,
										Description: "Groups with permission to authenticate.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Description: "User group name.",
													Computed:    true,
												},
											},
										},
									},
									"category": {
										Type:        schema.TypeInt,
										Description: "Categories and groups the filter examines.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "ID number.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
									"override_replacemsg": {
										Type:        schema.TypeString,
										Description: "Override replacement message.",
										Computed:    true,
									},
									"warn_duration": {
										Type:        schema.TypeString,
										Description: "Duration of warnings.",
										Computed:    true,
									},
									"warning_duration_type": {
										Type:        schema.TypeString,
										Description: "Re-display warning after closing browser or after a timeout.",
										Computed:    true,
									},
									"warning_prompt": {
										Type:        schema.TypeString,
										Description: "Warning prompts in each category or each domain.",
										Computed:    true,
									},
								},
							},
						},
						"max_quota_timeout": {
							Type:        schema.TypeInt,
							Description: "Maximum FortiGuard quota used by single page view in seconds (excludes streams).",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "Options for FortiGuard Web Filter.",
							Computed:    true,
						},
						"ovrd": {
							Type:        schema.TypeString,
							Description: "Allow web filter profile overrides.",
							Computed:    true,
						},
						"quota": {
							Type:        schema.TypeList,
							Description: "FortiGuard traffic quota settings.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category": {
										Type:        schema.TypeString,
										Description: "FortiGuard categories to apply quota to (category action must be set to monitor).",
										Computed:    true,
									},
									"duration": {
										Type:        schema.TypeString,
										Description: "Duration of quota.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "ID number.",
										Computed:    true,
									},
									"override_replacemsg": {
										Type:        schema.TypeString,
										Description: "Override replacement message.",
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "Quota type.",
										Computed:    true,
									},
									"unit": {
										Type:        schema.TypeString,
										Description: "Traffic quota unit of measurement.",
										Computed:    true,
									},
									"value": {
										Type:        schema.TypeInt,
										Description: "Traffic quota value.",
										Computed:    true,
									},
								},
							},
						},
						"rate_crl_urls": {
							Type:        schema.TypeString,
							Description: "Enable/disable rating CRL by URL.",
							Computed:    true,
						},
						"rate_css_urls": {
							Type:        schema.TypeString,
							Description: "Enable/disable rating CSS by URL.",
							Computed:    true,
						},
						"rate_image_urls": {
							Type:        schema.TypeString,
							Description: "Enable/disable rating images by URL.",
							Computed:    true,
						},
						"rate_javascript_urls": {
							Type:        schema.TypeString,
							Description: "Enable/disable rating JavaScript by URL.",
							Computed:    true,
						},
					},
				},
			},
			"https_replacemsg": {
				Type:        schema.TypeString,
				Description: "Enable replacement messages for HTTPS.",
				Computed:    true,
			},
			"log_all_url": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging all URLs visited.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Profile name.",
				Required:    true,
			},
			"options": {
				Type:        schema.TypeString,
				Description: "Options.",
				Computed:    true,
			},
			"override": {
				Type:        schema.TypeList,
				Description: "Web Filter override settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ovrd_cookie": {
							Type:        schema.TypeString,
							Description: "Allow/deny browser-based (cookie) overrides.",
							Computed:    true,
						},
						"ovrd_dur": {
							Type:        schema.TypeString,
							Description: "Override duration.",
							Computed:    true,
						},
						"ovrd_dur_mode": {
							Type:        schema.TypeString,
							Description: "Override duration mode.",
							Computed:    true,
						},
						"ovrd_scope": {
							Type:        schema.TypeString,
							Description: "Override scope.",
							Computed:    true,
						},
						"ovrd_user_group": {
							Type:        schema.TypeList,
							Description: "User groups with permission to use the override.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "User group name.",
										Computed:    true,
									},
								},
							},
						},
						"profile": {
							Type:        schema.TypeList,
							Description: "Web filter profile with permission to create overrides.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Web profile.",
										Computed:    true,
									},
								},
							},
						},
						"profile_attribute": {
							Type:        schema.TypeString,
							Description: "Profile attribute to retrieve from the RADIUS server.",
							Computed:    true,
						},
						"profile_type": {
							Type:        schema.TypeString,
							Description: "Override profile type.",
							Computed:    true,
						},
					},
				},
			},
			"ovrd_perm": {
				Type:        schema.TypeString,
				Description: "Permitted override types.",
				Computed:    true,
			},
			"post_action": {
				Type:        schema.TypeString,
				Description: "Action taken for HTTP POST traffic.",
				Computed:    true,
			},
			"replacemsg_group": {
				Type:        schema.TypeString,
				Description: "Replacement message group.",
				Computed:    true,
			},
			"web": {
				Type:        schema.TypeList,
				Description: "Web content filtering settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowlist": {
							Type:        schema.TypeString,
							Description: "FortiGuard allowlist settings.",
							Computed:    true,
						},
						"blacklist": {
							Type:        schema.TypeString,
							Description: "Enable/disable automatic addition of URLs detected by FortiSandbox to blacklist.",
							Computed:    true,
						},
						"blocklist": {
							Type:        schema.TypeString,
							Description: "Enable/disable automatic addition of URLs detected by FortiSandbox to blocklist.",
							Computed:    true,
						},
						"bword_table": {
							Type:        schema.TypeInt,
							Description: "Banned word table ID.",
							Computed:    true,
						},
						"bword_threshold": {
							Type:        schema.TypeInt,
							Description: "Banned word score threshold.",
							Computed:    true,
						},
						"content_header_list": {
							Type:        schema.TypeInt,
							Description: "Content header list.",
							Computed:    true,
						},
						"keyword_match": {
							Type:        schema.TypeList,
							Description: "Search keywords to log when match is found.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pattern": {
										Type:        schema.TypeString,
										Description: "Pattern/keyword to search for.",
										Computed:    true,
									},
								},
							},
						},
						"log_search": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging all search phrases.",
							Computed:    true,
						},
						"safe_search": {
							Type:        schema.TypeString,
							Description: "Safe search type.",
							Computed:    true,
						},
						"urlfilter_table": {
							Type:        schema.TypeInt,
							Description: "URL filter table ID.",
							Computed:    true,
						},
						"vimeo_restrict": {
							Type:        schema.TypeString,
							Description: "Set Vimeo-restrict (\"7\" = don't show mature content, \"134\" = don't show unrated and mature content). A value of cookie \"content_rating\".",
							Computed:    true,
						},
						"whitelist": {
							Type:        schema.TypeString,
							Description: "FortiGuard whitelist settings.",
							Computed:    true,
						},
						"youtube_restrict": {
							Type:        schema.TypeString,
							Description: "YouTube EDU filter level.",
							Computed:    true,
						},
					},
				},
			},
			"web_antiphishing_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging of AntiPhishing checks.",
				Computed:    true,
			},
			"web_content_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging logging blocked web content.",
				Computed:    true,
			},
			"web_extended_all_action_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable extended any filter action logging for web filtering.",
				Computed:    true,
			},
			"web_filter_activex_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging ActiveX.",
				Computed:    true,
			},
			"web_filter_applet_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging Java applets.",
				Computed:    true,
			},
			"web_filter_command_block_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging blocked commands.",
				Computed:    true,
			},
			"web_filter_cookie_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging cookie filtering.",
				Computed:    true,
			},
			"web_filter_cookie_removal_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging blocked cookies.",
				Computed:    true,
			},
			"web_filter_js_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging Java scripts.",
				Computed:    true,
			},
			"web_filter_jscript_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging JScripts.",
				Computed:    true,
			},
			"web_filter_referer_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging referrers.",
				Computed:    true,
			},
			"web_filter_unknown_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging unknown scripts.",
				Computed:    true,
			},
			"web_filter_vbs_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging VBS scripts.",
				Computed:    true,
			},
			"web_ftgd_err_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging rating errors.",
				Computed:    true,
			},
			"web_ftgd_quota_usage": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging daily quota usage.",
				Computed:    true,
			},
			"web_invalid_domain_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging invalid domain names.",
				Computed:    true,
			},
			"web_url_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging URL filtering.",
				Computed:    true,
			},
			"wisp": {
				Type:        schema.TypeString,
				Description: "Enable/disable web proxy WISP.",
				Computed:    true,
			},
			"wisp_algorithm": {
				Type:        schema.TypeString,
				Description: "WISP server selection algorithm.",
				Computed:    true,
			},
			"wisp_servers": {
				Type:        schema.TypeList,
				Description: "WISP servers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Server name.",
							Computed:    true,
						},
					},
				},
			},
			"youtube_channel_filter": {
				Type:        schema.TypeList,
				Description: "YouTube channel filter.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel_id": {
							Type:        schema.TypeString,
							Description: "YouTube channel ID to be filtered.",
							Computed:    true,
						},
						"comment": {
							Type:        schema.TypeString,
							Description: "Comment.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
					},
				},
			},
			"youtube_channel_status": {
				Type:        schema.TypeString,
				Description: "YouTube channel filter status.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWebfilterProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebfilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterProfile dataSource: %v", err)
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

	diags := refreshObjectWebfilterProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
