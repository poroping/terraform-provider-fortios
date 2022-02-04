// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Web filter profiles.

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

func resourceWebfilterProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Web filter profiles.",

		CreateContext: resourceWebfilterProfileCreate,
		ReadContext:   resourceWebfilterProfileRead,
		UpdateContext: resourceWebfilterProfileUpdate,
		DeleteContext: resourceWebfilterProfileDelete,

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
			"antiphish": {
				Type:        schema.TypeList,
				Description: "AntiPhishing profile.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"domain-controller", "ldap"}, false),

							Description: "Authentication methods.",
							Optional:    true,
							Computed:    true,
						},
						"check_basic_auth": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable checking of HTTP Basic Auth field for known credentials.",
							Optional:    true,
							Computed:    true,
						},
						"check_uri": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable checking of GET URI parameters for known credentials.",
							Optional:    true,
							Computed:    true,
						},
						"check_username_only": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable username only matching of credentials. Action will be taken for valid usernames regardless of password validity.",
							Optional:    true,
							Computed:    true,
						},
						"custom_patterns": {
							Type:        schema.TypeList,
							Description: "Custom username and password regex patterns.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"username", "password"}, false),

										Description: "Category that the pattern matches.",
										Optional:    true,
										Computed:    true,
									},
									"pattern": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),

										Description: "Target pattern.",
										Optional:    true,
										Computed:    true,
									},
									"type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"regex", "literal"}, false),

										Description: "Pattern will be treated either as a regex pattern or literal string.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"default_action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"exempt", "log", "block"}, false),

							Description: "Action to be taken when there is no matching rule.",
							Optional:    true,
							Computed:    true,
						},
						"domain_controller": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Domain for which to verify received credentials against.",
							Optional:    true,
							Computed:    true,
						},
						"inspection_entries": {
							Type:        schema.TypeList,
							Description: "AntiPhishing entries.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"exempt", "log", "block"}, false),

										Description: "Action to be taken upon an AntiPhishing match.",
										Optional:    true,
										Computed:    true,
									},
									"fortiguard_category": {
										Type: schema.TypeString,

										DiffSuppressFunc: suppressors.DiffMultiStringEqual,
										Description:      "FortiGuard category to match.",
										Optional:         true,
										Computed:         true,
									},
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "Inspection target name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"ldap": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "LDAP server for which to verify received credentials against.",
							Optional:    true,
							Computed:    true,
						},
						"max_body_len": {
							Type: schema.TypeInt,

							Description: "Maximum size of a POST body to check for credentials.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Toggle AntiPhishing functionality.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"extended_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable extended logging for web filtering.",
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
			"file_filter": {
				Type:        schema.TypeList,
				Description: "File filter.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entries": {
							Type:        schema.TypeList,
							Description: "File filter entries.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"log", "block"}, false),

										Description: "Action taken for matched file.",
										Optional:    true,
										Computed:    true,
									},
									"comment": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),

										Description: "Comment.",
										Optional:    true,
										Computed:    true,
									},
									"direction": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"incoming", "outgoing", "any"}, false),

										Description: "Match files transmitted in the session's originating or reply direction.",
										Optional:    true,
										Computed:    true,
									},
									"file_type": {
										Type:        schema.TypeList,
										Description: "Select file type.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 39),

													Description: "File type name.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"filter": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Add a file filter.",
										Optional:    true,
										Computed:    true,
									},
									"password_protected": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"yes", "any"}, false),

										Description: "Match password-protected files.",
										Optional:    true,
										Computed:    true,
									},
									"protocol": {
										Type: schema.TypeString,

										DiffSuppressFunc: suppressors.DiffFakeListEqual,
										Description:      "Protocols to apply with.",
										Optional:         true,
										Computed:         true,
									},
								},
							},
						},
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable file filter logging.",
							Optional:    true,
							Computed:    true,
						},
						"scan_archive_contents": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable file filter archive contents scan.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable file filter.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ftgd_wf": {
				Type:        schema.TypeList,
				Description: "FortiGuard Web Filter settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exempt_quota": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffMultiStringEqual,
							Description:      "Do not stop quota for these categories.",
							Optional:         true,
							Computed:         true,
						},
						"filters": {
							Type:        schema.TypeList,
							Description: "FortiGuard filters.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"block", "authenticate", "monitor", "warning"}, false),

										Description: "Action to take for matches.",
										Optional:    true,
										Computed:    true,
									},
									"auth_usr_grp": {
										Type:        schema.TypeList,
										Description: "Groups with permission to authenticate.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),

													Description: "User group name.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"category": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),

										Description: "Categories and groups the filter examines.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),

										Description: "ID number.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
									"override_replacemsg": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 28),

										Description: "Override replacement message.",
										Optional:    true,
										Computed:    true,
									},
									"warn_duration": {
										Type: schema.TypeString,

										Description: "Duration of warnings.",
										Optional:    true,
										Computed:    true,
									},
									"warning_duration_type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"session", "timeout"}, false),

										Description: "Re-display warning after closing browser or after a timeout.",
										Optional:    true,
										Computed:    true,
									},
									"warning_prompt": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"per-domain", "per-category"}, false),

										Description: "Warning prompts in each category or each domain.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"max_quota_timeout": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 86400),

							Description: "Maximum FortiGuard quota used by single page view in seconds (excludes streams).",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Options for FortiGuard Web Filter.",
							Optional:         true,
							Computed:         true,
						},
						"ovrd": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffMultiStringEqual,
							Description:      "Allow web filter profile overrides.",
							Optional:         true,
							Computed:         true,
						},
						"quota": {
							Type:        schema.TypeList,
							Description: "FortiGuard traffic quota settings.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category": {
										Type: schema.TypeString,

										DiffSuppressFunc: suppressors.DiffMultiStringEqual,
										Description:      "FortiGuard categories to apply quota to (category action must be set to monitor).",
										Optional:         true,
										Computed:         true,
									},
									"duration": {
										Type: schema.TypeString,

										Description: "Duration of quota.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "ID number.",
										Optional:    true,
										Computed:    true,
									},
									"override_replacemsg": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 28),

										Description: "Override replacement message.",
										Optional:    true,
										Computed:    true,
									},
									"type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"time", "traffic"}, false),

										Description: "Quota type.",
										Optional:    true,
										Computed:    true,
									},
									"unit": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"B", "KB", "MB", "GB"}, false),

										Description: "Traffic quota unit of measurement.",
										Optional:    true,
										Computed:    true,
									},
									"value": {
										Type: schema.TypeInt,

										Description: "Traffic quota value.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"rate_crl_urls": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable rating CRL by URL.",
							Optional:    true,
							Computed:    true,
						},
						"rate_css_urls": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable rating CSS by URL.",
							Optional:    true,
							Computed:    true,
						},
						"rate_image_urls": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable rating images by URL.",
							Optional:    true,
							Computed:    true,
						},
						"rate_javascript_urls": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable rating JavaScript by URL.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"https_replacemsg": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable replacement messages for HTTPS.",
				Optional:    true,
				Computed:    true,
			},
			"log_all_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging all URLs visited.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"options": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Options.",
				Optional:         true,
				Computed:         true,
			},
			"override": {
				Type:        schema.TypeList,
				Description: "Web Filter override settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ovrd_cookie": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),

							Description: "Allow/deny browser-based (cookie) overrides.",
							Optional:    true,
							Computed:    true,
						},
						"ovrd_dur": {
							Type: schema.TypeString,

							Description: "Override duration.",
							Optional:    true,
							Computed:    true,
						},
						"ovrd_dur_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"constant", "ask"}, false),

							Description: "Override duration mode.",
							Optional:    true,
							Computed:    true,
						},
						"ovrd_scope": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"user", "user-group", "ip", "browser", "ask"}, false),

							Description: "Override scope.",
							Optional:    true,
							Computed:    true,
						},
						"ovrd_user_group": {
							Type:        schema.TypeList,
							Description: "User groups with permission to use the override.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "User group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"profile": {
							Type:        schema.TypeList,
							Description: "Web filter profile with permission to create overrides.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Web profile.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"profile_attribute": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"User-Name", "NAS-IP-Address", "Framed-IP-Address", "Framed-IP-Netmask", "Filter-Id", "Login-IP-Host", "Reply-Message", "Callback-Number", "Callback-Id", "Framed-Route", "Framed-IPX-Network", "Class", "Called-Station-Id", "Calling-Station-Id", "NAS-Identifier", "Proxy-State", "Login-LAT-Service", "Login-LAT-Node", "Login-LAT-Group", "Framed-AppleTalk-Zone", "Acct-Session-Id", "Acct-Multi-Session-Id"}, false),

							Description: "Profile attribute to retrieve from the RADIUS server.",
							Optional:    true,
							Computed:    true,
						},
						"profile_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"list", "radius"}, false),

							Description: "Override profile type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ovrd_perm": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Permitted override types.",
				Optional:         true,
				Computed:         true,
			},
			"post_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"normal", "block"}, false),

				Description: "Action taken for HTTP POST traffic.",
				Optional:    true,
				Computed:    true,
			},
			"replacemsg_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Replacement message group.",
				Optional:    true,
				Computed:    true,
			},
			"web": {
				Type:        schema.TypeList,
				Description: "Web content filtering settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowlist": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "FortiGuard allowlist settings.",
							Optional:         true,
							Computed:         true,
						},
						"blacklist": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable automatic addition of URLs detected by FortiSandbox to blacklist.",
							Optional:    true,
							Computed:    true,
						},
						"blocklist": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable automatic addition of URLs detected by FortiSandbox to blocklist.",
							Optional:    true,
							Computed:    true,
						},
						"bword_table": {
							Type: schema.TypeInt,

							Description: "Banned word table ID.",
							Optional:    true,
							Computed:    true,
						},
						"bword_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 2147483647),

							Description: "Banned word score threshold.",
							Optional:    true,
							Computed:    true,
						},
						"content_header_list": {
							Type: schema.TypeInt,

							Description: "Content header list.",
							Optional:    true,
							Computed:    true,
						},
						"keyword_match": {
							Type:        schema.TypeList,
							Description: "Search keywords to log when match is found.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pattern": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Pattern/keyword to search for.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"log_search": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging all search phrases.",
							Optional:    true,
							Computed:    true,
						},
						"safe_search": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Safe search type.",
							Optional:         true,
							Computed:         true,
						},
						"urlfilter_table": {
							Type: schema.TypeInt,

							Description: "URL filter table ID.",
							Optional:    true,
							Computed:    true,
						},
						"vimeo_restrict": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Set Vimeo-restrict (\"7\" = don't show mature content, \"134\" = don't show unrated and mature content). A value of cookie \"content_rating\".",
							Optional:    true,
							Computed:    true,
						},
						"whitelist": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "FortiGuard whitelist settings.",
							Optional:         true,
							Computed:         true,
						},
						"youtube_restrict": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "strict", "moderate"}, false),

							Description: "YouTube EDU filter level.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"web_antiphishing_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging of AntiPhishing checks.",
				Optional:    true,
				Computed:    true,
			},
			"web_content_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging logging blocked web content.",
				Optional:    true,
				Computed:    true,
			},
			"web_extended_all_action_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable extended any filter action logging for web filtering.",
				Optional:    true,
				Computed:    true,
			},
			"web_filter_activex_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging ActiveX.",
				Optional:    true,
				Computed:    true,
			},
			"web_filter_applet_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging Java applets.",
				Optional:    true,
				Computed:    true,
			},
			"web_filter_command_block_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging blocked commands.",
				Optional:    true,
				Computed:    true,
			},
			"web_filter_cookie_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging cookie filtering.",
				Optional:    true,
				Computed:    true,
			},
			"web_filter_cookie_removal_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging blocked cookies.",
				Optional:    true,
				Computed:    true,
			},
			"web_filter_js_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging Java scripts.",
				Optional:    true,
				Computed:    true,
			},
			"web_filter_jscript_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging JScripts.",
				Optional:    true,
				Computed:    true,
			},
			"web_filter_referer_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging referrers.",
				Optional:    true,
				Computed:    true,
			},
			"web_filter_unknown_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging unknown scripts.",
				Optional:    true,
				Computed:    true,
			},
			"web_filter_vbs_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging VBS scripts.",
				Optional:    true,
				Computed:    true,
			},
			"web_ftgd_err_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging rating errors.",
				Optional:    true,
				Computed:    true,
			},
			"web_ftgd_quota_usage": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging daily quota usage.",
				Optional:    true,
				Computed:    true,
			},
			"web_invalid_domain_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging invalid domain names.",
				Optional:    true,
				Computed:    true,
			},
			"web_url_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging URL filtering.",
				Optional:    true,
				Computed:    true,
			},
			"wisp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable web proxy WISP.",
				Optional:    true,
				Computed:    true,
			},
			"wisp_algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"primary-secondary", "round-robin", "auto-learning"}, false),

				Description: "WISP server selection algorithm.",
				Optional:    true,
				Computed:    true,
			},
			"wisp_servers": {
				Type:        schema.TypeList,
				Description: "WISP servers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Server name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"youtube_channel_filter": {
				Type:        schema.TypeList,
				Description: "YouTube channel filter.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "YouTube channel ID to be filtered.",
							Optional:    true,
							Computed:    true,
						},
						"comment": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Comment.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"youtube_channel_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "blacklist", "whitelist"}, false),

				Description: "YouTube channel filter status.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWebfilterProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WebfilterProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWebfilterProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWebfilterProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterProfile")
	}

	return resourceWebfilterProfileRead(ctx, d, meta)
}

func resourceWebfilterProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebfilterProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWebfilterProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WebfilterProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterProfile")
	}

	return resourceWebfilterProfileRead(ctx, d, meta)
}

func resourceWebfilterProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWebfilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WebfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebfilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterProfile resource: %v", err)
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

	diags := refreshObjectWebfilterProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWebfilterProfileAntiphish(v *[]models.WebfilterProfileAntiphish, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Authentication; tmp != nil {
				v["authentication"] = *tmp
			}

			if tmp := cfg.CheckBasicAuth; tmp != nil {
				v["check_basic_auth"] = *tmp
			}

			if tmp := cfg.CheckUri; tmp != nil {
				v["check_uri"] = *tmp
			}

			if tmp := cfg.CheckUsernameOnly; tmp != nil {
				v["check_username_only"] = *tmp
			}

			if tmp := cfg.CustomPatterns; tmp != nil {
				v["custom_patterns"] = flattenWebfilterProfileAntiphishCustomPatterns(tmp, sort)
			}

			if tmp := cfg.DefaultAction; tmp != nil {
				v["default_action"] = *tmp
			}

			if tmp := cfg.DomainController; tmp != nil {
				v["domain_controller"] = *tmp
			}

			if tmp := cfg.InspectionEntries; tmp != nil {
				v["inspection_entries"] = flattenWebfilterProfileAntiphishInspectionEntries(tmp, sort)
			}

			if tmp := cfg.Ldap; tmp != nil {
				v["ldap"] = *tmp
			}

			if tmp := cfg.MaxBodyLen; tmp != nil {
				v["max_body_len"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWebfilterProfileAntiphishCustomPatterns(v *[]models.WebfilterProfileAntiphishCustomPatterns, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Pattern; tmp != nil {
				v["pattern"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "pattern")
	}

	return flat
}

func flattenWebfilterProfileAntiphishInspectionEntries(v *[]models.WebfilterProfileAntiphishInspectionEntries, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.FortiguardCategory; tmp != nil {
				v["fortiguard_category"] = *tmp
			}

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

func flattenWebfilterProfileFileFilter(v *[]models.WebfilterProfileFileFilter, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Entries; tmp != nil {
				v["entries"] = flattenWebfilterProfileFileFilterEntries(tmp, sort)
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.ScanArchiveContents; tmp != nil {
				v["scan_archive_contents"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWebfilterProfileFileFilterEntries(v *[]models.WebfilterProfileFileFilterEntries, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Comment; tmp != nil {
				v["comment"] = *tmp
			}

			if tmp := cfg.Direction; tmp != nil {
				v["direction"] = *tmp
			}

			if tmp := cfg.FileType; tmp != nil {
				v["file_type"] = flattenWebfilterProfileFileFilterEntriesFileType(tmp, sort)
			}

			if tmp := cfg.Filter; tmp != nil {
				v["filter"] = *tmp
			}

			if tmp := cfg.PasswordProtected; tmp != nil {
				v["password_protected"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "filter")
	}

	return flat
}

func flattenWebfilterProfileFileFilterEntriesFileType(v *[]models.WebfilterProfileFileFilterEntriesFileType, sort bool) interface{} {
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

func flattenWebfilterProfileFtgdWf(v *[]models.WebfilterProfileFtgdWf, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.ExemptQuota; tmp != nil {
				v["exempt_quota"] = *tmp
			}

			if tmp := cfg.Filters; tmp != nil {
				v["filters"] = flattenWebfilterProfileFtgdWfFilters(tmp, sort)
			}

			if tmp := cfg.MaxQuotaTimeout; tmp != nil {
				v["max_quota_timeout"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.Ovrd; tmp != nil {
				v["ovrd"] = *tmp
			}

			if tmp := cfg.Quota; tmp != nil {
				v["quota"] = flattenWebfilterProfileFtgdWfQuota(tmp, sort)
			}

			if tmp := cfg.RateCrlUrls; tmp != nil {
				v["rate_crl_urls"] = *tmp
			}

			if tmp := cfg.RateCssUrls; tmp != nil {
				v["rate_css_urls"] = *tmp
			}

			if tmp := cfg.RateImageUrls; tmp != nil {
				v["rate_image_urls"] = *tmp
			}

			if tmp := cfg.RateJavascriptUrls; tmp != nil {
				v["rate_javascript_urls"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWebfilterProfileFtgdWfFilters(v *[]models.WebfilterProfileFtgdWfFilters, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.AuthUsrGrp; tmp != nil {
				v["auth_usr_grp"] = flattenWebfilterProfileFtgdWfFiltersAuthUsrGrp(tmp, sort)
			}

			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.OverrideReplacemsg; tmp != nil {
				v["override_replacemsg"] = *tmp
			}

			if tmp := cfg.WarnDuration; tmp != nil {
				v["warn_duration"] = *tmp
			}

			if tmp := cfg.WarningDurationType; tmp != nil {
				v["warning_duration_type"] = *tmp
			}

			if tmp := cfg.WarningPrompt; tmp != nil {
				v["warning_prompt"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWebfilterProfileFtgdWfFiltersAuthUsrGrp(v *[]models.WebfilterProfileFtgdWfFiltersAuthUsrGrp, sort bool) interface{} {
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

func flattenWebfilterProfileFtgdWfQuota(v *[]models.WebfilterProfileFtgdWfQuota, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Duration; tmp != nil {
				v["duration"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.OverrideReplacemsg; tmp != nil {
				v["override_replacemsg"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.Unit; tmp != nil {
				v["unit"] = *tmp
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWebfilterProfileOverride(v *[]models.WebfilterProfileOverride, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.OvrdCookie; tmp != nil {
				v["ovrd_cookie"] = *tmp
			}

			if tmp := cfg.OvrdDur; tmp != nil {
				v["ovrd_dur"] = *tmp
			}

			if tmp := cfg.OvrdDurMode; tmp != nil {
				v["ovrd_dur_mode"] = *tmp
			}

			if tmp := cfg.OvrdScope; tmp != nil {
				v["ovrd_scope"] = *tmp
			}

			if tmp := cfg.OvrdUserGroup; tmp != nil {
				v["ovrd_user_group"] = flattenWebfilterProfileOverrideOvrdUserGroup(tmp, sort)
			}

			if tmp := cfg.Profile; tmp != nil {
				v["profile"] = flattenWebfilterProfileOverrideProfile(tmp, sort)
			}

			if tmp := cfg.ProfileAttribute; tmp != nil {
				v["profile_attribute"] = *tmp
			}

			if tmp := cfg.ProfileType; tmp != nil {
				v["profile_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWebfilterProfileOverrideOvrdUserGroup(v *[]models.WebfilterProfileOverrideOvrdUserGroup, sort bool) interface{} {
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

func flattenWebfilterProfileOverrideProfile(v *[]models.WebfilterProfileOverrideProfile, sort bool) interface{} {
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

func flattenWebfilterProfileWeb(v *[]models.WebfilterProfileWeb, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Allowlist; tmp != nil {
				v["allowlist"] = *tmp
			}

			if tmp := cfg.Blacklist; tmp != nil {
				v["blacklist"] = *tmp
			}

			if tmp := cfg.Blocklist; tmp != nil {
				v["blocklist"] = *tmp
			}

			if tmp := cfg.BwordTable; tmp != nil {
				v["bword_table"] = *tmp
			}

			if tmp := cfg.BwordThreshold; tmp != nil {
				v["bword_threshold"] = *tmp
			}

			if tmp := cfg.ContentHeaderList; tmp != nil {
				v["content_header_list"] = *tmp
			}

			if tmp := cfg.KeywordMatch; tmp != nil {
				v["keyword_match"] = flattenWebfilterProfileWebKeywordMatch(tmp, sort)
			}

			if tmp := cfg.LogSearch; tmp != nil {
				v["log_search"] = *tmp
			}

			if tmp := cfg.SafeSearch; tmp != nil {
				v["safe_search"] = *tmp
			}

			if tmp := cfg.UrlfilterTable; tmp != nil {
				v["urlfilter_table"] = *tmp
			}

			if tmp := cfg.VimeoRestrict; tmp != nil {
				v["vimeo_restrict"] = *tmp
			}

			if tmp := cfg.Whitelist; tmp != nil {
				v["whitelist"] = *tmp
			}

			if tmp := cfg.YoutubeRestrict; tmp != nil {
				v["youtube_restrict"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWebfilterProfileWebKeywordMatch(v *[]models.WebfilterProfileWebKeywordMatch, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Pattern; tmp != nil {
				v["pattern"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "pattern")
	}

	return flat
}

func flattenWebfilterProfileWispServers(v *[]models.WebfilterProfileWispServers, sort bool) interface{} {
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

func flattenWebfilterProfileYoutubeChannelFilter(v *[]models.WebfilterProfileYoutubeChannelFilter, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.ChannelId; tmp != nil {
				v["channel_id"] = *tmp
			}

			if tmp := cfg.Comment; tmp != nil {
				v["comment"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectWebfilterProfile(d *schema.ResourceData, o *models.WebfilterProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Antiphish != nil {
		if err = d.Set("antiphish", flattenWebfilterProfileAntiphish(o.Antiphish, sort)); err != nil {
			return diag.Errorf("error reading antiphish: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
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

	if o.FileFilter != nil {
		if err = d.Set("file_filter", flattenWebfilterProfileFileFilter(o.FileFilter, sort)); err != nil {
			return diag.Errorf("error reading file_filter: %v", err)
		}
	}

	if o.FtgdWf != nil {
		if err = d.Set("ftgd_wf", flattenWebfilterProfileFtgdWf(o.FtgdWf, sort)); err != nil {
			return diag.Errorf("error reading ftgd_wf: %v", err)
		}
	}

	if o.HttpsReplacemsg != nil {
		v := *o.HttpsReplacemsg

		if err = d.Set("https_replacemsg", v); err != nil {
			return diag.Errorf("error reading https_replacemsg: %v", err)
		}
	}

	if o.LogAllUrl != nil {
		v := *o.LogAllUrl

		if err = d.Set("log_all_url", v); err != nil {
			return diag.Errorf("error reading log_all_url: %v", err)
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

	if o.Override != nil {
		if err = d.Set("override", flattenWebfilterProfileOverride(o.Override, sort)); err != nil {
			return diag.Errorf("error reading override: %v", err)
		}
	}

	if o.OvrdPerm != nil {
		v := *o.OvrdPerm

		if err = d.Set("ovrd_perm", v); err != nil {
			return diag.Errorf("error reading ovrd_perm: %v", err)
		}
	}

	if o.PostAction != nil {
		v := *o.PostAction

		if err = d.Set("post_action", v); err != nil {
			return diag.Errorf("error reading post_action: %v", err)
		}
	}

	if o.ReplacemsgGroup != nil {
		v := *o.ReplacemsgGroup

		if err = d.Set("replacemsg_group", v); err != nil {
			return diag.Errorf("error reading replacemsg_group: %v", err)
		}
	}

	if o.Web != nil {
		if err = d.Set("web", flattenWebfilterProfileWeb(o.Web, sort)); err != nil {
			return diag.Errorf("error reading web: %v", err)
		}
	}

	if o.WebAntiphishingLog != nil {
		v := *o.WebAntiphishingLog

		if err = d.Set("web_antiphishing_log", v); err != nil {
			return diag.Errorf("error reading web_antiphishing_log: %v", err)
		}
	}

	if o.WebContentLog != nil {
		v := *o.WebContentLog

		if err = d.Set("web_content_log", v); err != nil {
			return diag.Errorf("error reading web_content_log: %v", err)
		}
	}

	if o.WebExtendedAllActionLog != nil {
		v := *o.WebExtendedAllActionLog

		if err = d.Set("web_extended_all_action_log", v); err != nil {
			return diag.Errorf("error reading web_extended_all_action_log: %v", err)
		}
	}

	if o.WebFilterActivexLog != nil {
		v := *o.WebFilterActivexLog

		if err = d.Set("web_filter_activex_log", v); err != nil {
			return diag.Errorf("error reading web_filter_activex_log: %v", err)
		}
	}

	if o.WebFilterAppletLog != nil {
		v := *o.WebFilterAppletLog

		if err = d.Set("web_filter_applet_log", v); err != nil {
			return diag.Errorf("error reading web_filter_applet_log: %v", err)
		}
	}

	if o.WebFilterCommandBlockLog != nil {
		v := *o.WebFilterCommandBlockLog

		if err = d.Set("web_filter_command_block_log", v); err != nil {
			return diag.Errorf("error reading web_filter_command_block_log: %v", err)
		}
	}

	if o.WebFilterCookieLog != nil {
		v := *o.WebFilterCookieLog

		if err = d.Set("web_filter_cookie_log", v); err != nil {
			return diag.Errorf("error reading web_filter_cookie_log: %v", err)
		}
	}

	if o.WebFilterCookieRemovalLog != nil {
		v := *o.WebFilterCookieRemovalLog

		if err = d.Set("web_filter_cookie_removal_log", v); err != nil {
			return diag.Errorf("error reading web_filter_cookie_removal_log: %v", err)
		}
	}

	if o.WebFilterJsLog != nil {
		v := *o.WebFilterJsLog

		if err = d.Set("web_filter_js_log", v); err != nil {
			return diag.Errorf("error reading web_filter_js_log: %v", err)
		}
	}

	if o.WebFilterJscriptLog != nil {
		v := *o.WebFilterJscriptLog

		if err = d.Set("web_filter_jscript_log", v); err != nil {
			return diag.Errorf("error reading web_filter_jscript_log: %v", err)
		}
	}

	if o.WebFilterRefererLog != nil {
		v := *o.WebFilterRefererLog

		if err = d.Set("web_filter_referer_log", v); err != nil {
			return diag.Errorf("error reading web_filter_referer_log: %v", err)
		}
	}

	if o.WebFilterUnknownLog != nil {
		v := *o.WebFilterUnknownLog

		if err = d.Set("web_filter_unknown_log", v); err != nil {
			return diag.Errorf("error reading web_filter_unknown_log: %v", err)
		}
	}

	if o.WebFilterVbsLog != nil {
		v := *o.WebFilterVbsLog

		if err = d.Set("web_filter_vbs_log", v); err != nil {
			return diag.Errorf("error reading web_filter_vbs_log: %v", err)
		}
	}

	if o.WebFtgdErrLog != nil {
		v := *o.WebFtgdErrLog

		if err = d.Set("web_ftgd_err_log", v); err != nil {
			return diag.Errorf("error reading web_ftgd_err_log: %v", err)
		}
	}

	if o.WebFtgdQuotaUsage != nil {
		v := *o.WebFtgdQuotaUsage

		if err = d.Set("web_ftgd_quota_usage", v); err != nil {
			return diag.Errorf("error reading web_ftgd_quota_usage: %v", err)
		}
	}

	if o.WebInvalidDomainLog != nil {
		v := *o.WebInvalidDomainLog

		if err = d.Set("web_invalid_domain_log", v); err != nil {
			return diag.Errorf("error reading web_invalid_domain_log: %v", err)
		}
	}

	if o.WebUrlLog != nil {
		v := *o.WebUrlLog

		if err = d.Set("web_url_log", v); err != nil {
			return diag.Errorf("error reading web_url_log: %v", err)
		}
	}

	if o.Wisp != nil {
		v := *o.Wisp

		if err = d.Set("wisp", v); err != nil {
			return diag.Errorf("error reading wisp: %v", err)
		}
	}

	if o.WispAlgorithm != nil {
		v := *o.WispAlgorithm

		if err = d.Set("wisp_algorithm", v); err != nil {
			return diag.Errorf("error reading wisp_algorithm: %v", err)
		}
	}

	if o.WispServers != nil {
		if err = d.Set("wisp_servers", flattenWebfilterProfileWispServers(o.WispServers, sort)); err != nil {
			return diag.Errorf("error reading wisp_servers: %v", err)
		}
	}

	if o.YoutubeChannelFilter != nil {
		if err = d.Set("youtube_channel_filter", flattenWebfilterProfileYoutubeChannelFilter(o.YoutubeChannelFilter, sort)); err != nil {
			return diag.Errorf("error reading youtube_channel_filter: %v", err)
		}
	}

	if o.YoutubeChannelStatus != nil {
		v := *o.YoutubeChannelStatus

		if err = d.Set("youtube_channel_status", v); err != nil {
			return diag.Errorf("error reading youtube_channel_status: %v", err)
		}
	}

	return nil
}

func expandWebfilterProfileAntiphish(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileAntiphish, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileAntiphish

	for i := range l {
		tmp := models.WebfilterProfileAntiphish{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.authentication", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Authentication = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.check_basic_auth", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CheckBasicAuth = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.check_uri", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CheckUri = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.check_username_only", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CheckUsernameOnly = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.custom_patterns", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebfilterProfileAntiphishCustomPatterns(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebfilterProfileAntiphishCustomPatterns
			// 	}
			tmp.CustomPatterns = v2

		}

		pre_append = fmt.Sprintf("%s.%d.default_action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DefaultAction = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.domain_controller", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DomainController = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.inspection_entries", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebfilterProfileAntiphishInspectionEntries(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebfilterProfileAntiphishInspectionEntries
			// 	}
			tmp.InspectionEntries = v2

		}

		pre_append = fmt.Sprintf("%s.%d.ldap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ldap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_body_len", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MaxBodyLen = &v2
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

func expandWebfilterProfileAntiphishCustomPatterns(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileAntiphishCustomPatterns, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileAntiphishCustomPatterns

	for i := range l {
		tmp := models.WebfilterProfileAntiphishCustomPatterns{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Pattern = &v2
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

func expandWebfilterProfileAntiphishInspectionEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileAntiphishInspectionEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileAntiphishInspectionEntries

	for i := range l {
		tmp := models.WebfilterProfileAntiphishInspectionEntries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortiguard_category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FortiguardCategory = &v2
			}
		}

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

func expandWebfilterProfileFileFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileFileFilter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileFileFilter

	for i := range l {
		tmp := models.WebfilterProfileFileFilter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.entries", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebfilterProfileFileFilterEntries(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebfilterProfileFileFilterEntries
			// 	}
			tmp.Entries = v2

		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.scan_archive_contents", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScanArchiveContents = &v2
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

func expandWebfilterProfileFileFilterEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileFileFilterEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileFileFilterEntries

	for i := range l {
		tmp := models.WebfilterProfileFileFilterEntries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.comment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comment = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.direction", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Direction = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.file_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebfilterProfileFileFilterEntriesFileType(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebfilterProfileFileFilterEntriesFileType
			// 	}
			tmp.FileType = v2

		}

		pre_append = fmt.Sprintf("%s.%d.filter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Filter = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.password_protected", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PasswordProtected = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Protocol = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWebfilterProfileFileFilterEntriesFileType(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileFileFilterEntriesFileType, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileFileFilterEntriesFileType

	for i := range l {
		tmp := models.WebfilterProfileFileFilterEntriesFileType{}
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

func expandWebfilterProfileFtgdWf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileFtgdWf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileFtgdWf

	for i := range l {
		tmp := models.WebfilterProfileFtgdWf{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.exempt_quota", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExemptQuota = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filters", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebfilterProfileFtgdWfFilters(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebfilterProfileFtgdWfFilters
			// 	}
			tmp.Filters = v2

		}

		pre_append = fmt.Sprintf("%s.%d.max_quota_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MaxQuotaTimeout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ovrd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ovrd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quota", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebfilterProfileFtgdWfQuota(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebfilterProfileFtgdWfQuota
			// 	}
			tmp.Quota = v2

		}

		pre_append = fmt.Sprintf("%s.%d.rate_crl_urls", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RateCrlUrls = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rate_css_urls", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RateCssUrls = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rate_image_urls", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RateImageUrls = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rate_javascript_urls", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RateJavascriptUrls = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWebfilterProfileFtgdWfFilters(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileFtgdWfFilters, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileFtgdWfFilters

	for i := range l {
		tmp := models.WebfilterProfileFtgdWfFilters{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_usr_grp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebfilterProfileFtgdWfFiltersAuthUsrGrp(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebfilterProfileFtgdWfFiltersAuthUsrGrp
			// 	}
			tmp.AuthUsrGrp = v2

		}

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

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_replacemsg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideReplacemsg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.warn_duration", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.WarnDuration = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.warning_duration_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.WarningDurationType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.warning_prompt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.WarningPrompt = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWebfilterProfileFtgdWfFiltersAuthUsrGrp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileFtgdWfFiltersAuthUsrGrp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileFtgdWfFiltersAuthUsrGrp

	for i := range l {
		tmp := models.WebfilterProfileFtgdWfFiltersAuthUsrGrp{}
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

func expandWebfilterProfileFtgdWfQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileFtgdWfQuota, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileFtgdWfQuota

	for i := range l {
		tmp := models.WebfilterProfileFtgdWfQuota{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.duration", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Duration = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_replacemsg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideReplacemsg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Unit = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Value = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWebfilterProfileOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileOverride, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileOverride

	for i := range l {
		tmp := models.WebfilterProfileOverride{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ovrd_cookie", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OvrdCookie = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ovrd_dur", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OvrdDur = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ovrd_dur_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OvrdDurMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ovrd_scope", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OvrdScope = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ovrd_user_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebfilterProfileOverrideOvrdUserGroup(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebfilterProfileOverrideOvrdUserGroup
			// 	}
			tmp.OvrdUserGroup = v2

		}

		pre_append = fmt.Sprintf("%s.%d.profile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebfilterProfileOverrideProfile(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebfilterProfileOverrideProfile
			// 	}
			tmp.Profile = v2

		}

		pre_append = fmt.Sprintf("%s.%d.profile_attribute", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProfileAttribute = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.profile_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProfileType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWebfilterProfileOverrideOvrdUserGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileOverrideOvrdUserGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileOverrideOvrdUserGroup

	for i := range l {
		tmp := models.WebfilterProfileOverrideOvrdUserGroup{}
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

func expandWebfilterProfileOverrideProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileOverrideProfile, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileOverrideProfile

	for i := range l {
		tmp := models.WebfilterProfileOverrideProfile{}
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

func expandWebfilterProfileWeb(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileWeb, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileWeb

	for i := range l {
		tmp := models.WebfilterProfileWeb{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.allowlist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Allowlist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.blacklist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Blacklist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.blocklist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Blocklist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bword_table", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.BwordTable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bword_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.BwordThreshold = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.content_header_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.ContentHeaderList = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keyword_match", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebfilterProfileWebKeywordMatch(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebfilterProfileWebKeywordMatch
			// 	}
			tmp.KeywordMatch = v2

		}

		pre_append = fmt.Sprintf("%s.%d.log_search", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogSearch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.safe_search", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SafeSearch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.urlfilter_table", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.UrlfilterTable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vimeo_restrict", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VimeoRestrict = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.whitelist", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Whitelist = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.youtube_restrict", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.YoutubeRestrict = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWebfilterProfileWebKeywordMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileWebKeywordMatch, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileWebKeywordMatch

	for i := range l {
		tmp := models.WebfilterProfileWebKeywordMatch{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Pattern = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWebfilterProfileWispServers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileWispServers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileWispServers

	for i := range l {
		tmp := models.WebfilterProfileWispServers{}
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

func expandWebfilterProfileYoutubeChannelFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterProfileYoutubeChannelFilter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterProfileYoutubeChannelFilter

	for i := range l {
		tmp := models.WebfilterProfileYoutubeChannelFilter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.channel_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ChannelId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.comment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comment = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWebfilterProfile(d *schema.ResourceData, sv string) (*models.WebfilterProfile, diag.Diagnostics) {
	obj := models.WebfilterProfile{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("antiphish"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("antiphish", sv)
			diags = append(diags, e)
		}
		t, err := expandWebfilterProfileAntiphish(d, v, "antiphish", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Antiphish = t
		}
	} else if d.HasChange("antiphish") {
		old, new := d.GetChange("antiphish")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Antiphish = &[]models.WebfilterProfileAntiphish{}
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
	if v1, ok := d.GetOk("extended_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extended_log", sv)
				diags = append(diags, e)
			}
			obj.ExtendedLog = &v2
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
	if v, ok := d.GetOk("file_filter"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("file_filter", sv)
			diags = append(diags, e)
		}
		t, err := expandWebfilterProfileFileFilter(d, v, "file_filter", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FileFilter = t
		}
	} else if d.HasChange("file_filter") {
		old, new := d.GetChange("file_filter")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FileFilter = &[]models.WebfilterProfileFileFilter{}
		}
	}
	if v, ok := d.GetOk("ftgd_wf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ftgd_wf", sv)
			diags = append(diags, e)
		}
		t, err := expandWebfilterProfileFtgdWf(d, v, "ftgd_wf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FtgdWf = t
		}
	} else if d.HasChange("ftgd_wf") {
		old, new := d.GetChange("ftgd_wf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FtgdWf = &[]models.WebfilterProfileFtgdWf{}
		}
	}
	if v1, ok := d.GetOk("https_replacemsg"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("https_replacemsg", sv)
				diags = append(diags, e)
			}
			obj.HttpsReplacemsg = &v2
		}
	}
	if v1, ok := d.GetOk("log_all_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_all_url", sv)
				diags = append(diags, e)
			}
			obj.LogAllUrl = &v2
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
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("options", sv)
				diags = append(diags, e)
			}
			obj.Options = &v2
		}
	}
	if v, ok := d.GetOk("override"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("override", sv)
			diags = append(diags, e)
		}
		t, err := expandWebfilterProfileOverride(d, v, "override", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Override = t
		}
	} else if d.HasChange("override") {
		old, new := d.GetChange("override")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Override = &[]models.WebfilterProfileOverride{}
		}
	}
	if v1, ok := d.GetOk("ovrd_perm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ovrd_perm", sv)
				diags = append(diags, e)
			}
			obj.OvrdPerm = &v2
		}
	}
	if v1, ok := d.GetOk("post_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("post_action", sv)
				diags = append(diags, e)
			}
			obj.PostAction = &v2
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
	if v, ok := d.GetOk("web"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("web", sv)
			diags = append(diags, e)
		}
		t, err := expandWebfilterProfileWeb(d, v, "web", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Web = t
		}
	} else if d.HasChange("web") {
		old, new := d.GetChange("web")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Web = &[]models.WebfilterProfileWeb{}
		}
	}
	if v1, ok := d.GetOk("web_antiphishing_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("web_antiphishing_log", sv)
				diags = append(diags, e)
			}
			obj.WebAntiphishingLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_content_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_content_log", sv)
				diags = append(diags, e)
			}
			obj.WebContentLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_extended_all_action_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_extended_all_action_log", sv)
				diags = append(diags, e)
			}
			obj.WebExtendedAllActionLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_filter_activex_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_filter_activex_log", sv)
				diags = append(diags, e)
			}
			obj.WebFilterActivexLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_filter_applet_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_filter_applet_log", sv)
				diags = append(diags, e)
			}
			obj.WebFilterAppletLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_filter_command_block_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_filter_command_block_log", sv)
				diags = append(diags, e)
			}
			obj.WebFilterCommandBlockLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_filter_cookie_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_filter_cookie_log", sv)
				diags = append(diags, e)
			}
			obj.WebFilterCookieLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_filter_cookie_removal_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_filter_cookie_removal_log", sv)
				diags = append(diags, e)
			}
			obj.WebFilterCookieRemovalLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_filter_js_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_filter_js_log", sv)
				diags = append(diags, e)
			}
			obj.WebFilterJsLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_filter_jscript_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_filter_jscript_log", sv)
				diags = append(diags, e)
			}
			obj.WebFilterJscriptLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_filter_referer_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_filter_referer_log", sv)
				diags = append(diags, e)
			}
			obj.WebFilterRefererLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_filter_unknown_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_filter_unknown_log", sv)
				diags = append(diags, e)
			}
			obj.WebFilterUnknownLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_filter_vbs_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_filter_vbs_log", sv)
				diags = append(diags, e)
			}
			obj.WebFilterVbsLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_ftgd_err_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_ftgd_err_log", sv)
				diags = append(diags, e)
			}
			obj.WebFtgdErrLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_ftgd_quota_usage"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_ftgd_quota_usage", sv)
				diags = append(diags, e)
			}
			obj.WebFtgdQuotaUsage = &v2
		}
	}
	if v1, ok := d.GetOk("web_invalid_domain_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_invalid_domain_log", sv)
				diags = append(diags, e)
			}
			obj.WebInvalidDomainLog = &v2
		}
	}
	if v1, ok := d.GetOk("web_url_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_url_log", sv)
				diags = append(diags, e)
			}
			obj.WebUrlLog = &v2
		}
	}
	if v1, ok := d.GetOk("wisp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wisp", sv)
				diags = append(diags, e)
			}
			obj.Wisp = &v2
		}
	}
	if v1, ok := d.GetOk("wisp_algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wisp_algorithm", sv)
				diags = append(diags, e)
			}
			obj.WispAlgorithm = &v2
		}
	}
	if v, ok := d.GetOk("wisp_servers"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("wisp_servers", sv)
			diags = append(diags, e)
		}
		t, err := expandWebfilterProfileWispServers(d, v, "wisp_servers", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.WispServers = t
		}
	} else if d.HasChange("wisp_servers") {
		old, new := d.GetChange("wisp_servers")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.WispServers = &[]models.WebfilterProfileWispServers{}
		}
	}
	if v, ok := d.GetOk("youtube_channel_filter"); ok {
		if !utils.CheckVer(sv, "", "v7.0.0") {
			e := utils.AttributeVersionWarning("youtube_channel_filter", sv)
			diags = append(diags, e)
		}
		t, err := expandWebfilterProfileYoutubeChannelFilter(d, v, "youtube_channel_filter", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.YoutubeChannelFilter = t
		}
	} else if d.HasChange("youtube_channel_filter") {
		old, new := d.GetChange("youtube_channel_filter")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.YoutubeChannelFilter = &[]models.WebfilterProfileYoutubeChannelFilter{}
		}
	}
	if v1, ok := d.GetOk("youtube_channel_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("youtube_channel_status", sv)
				diags = append(diags, e)
			}
			obj.YoutubeChannelStatus = &v2
		}
	}
	return &obj, diags
}
