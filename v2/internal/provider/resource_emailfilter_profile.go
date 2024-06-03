// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Email Filter profiles.

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

func resourceEmailfilterProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Email Filter profiles.",

		CreateContext: resourceEmailfilterProfileCreate,
		ReadContext:   resourceEmailfilterProfileRead,
		UpdateContext: resourceEmailfilterProfileUpdate,
		DeleteContext: resourceEmailfilterProfileDelete,

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
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"external": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable external Email inspection.",
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
				Optional:    true, MaxItems: 1,
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
			"gmail": {
				Type:        schema.TypeList,
				Description: "Gmail.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging.",
							Optional:    true,
							Computed:    true,
						},
						"log_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of all email traffic.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"imap": {
				Type:        schema.TypeList,
				Description: "IMAP.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"pass", "tag"}, false),

							Description: "Action for spam email.",
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
						"log_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of all email traffic.",
							Optional:    true,
							Computed:    true,
						},
						"tag_msg": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Subject text or header added to spam email.",
							Optional:    true,
							Computed:    true,
						},
						"tag_type": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Tag subject or header for spam email.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"mapi": {
				Type:        schema.TypeList,
				Description: "MAPI.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"pass", "discard"}, false),

							Description: "Action for spam email.",
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
						"log_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of all email traffic.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"msn_hotmail": {
				Type:        schema.TypeList,
				Description: "MSN Hotmail.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging.",
							Optional:    true,
							Computed:    true,
						},
						"log_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of all email traffic.",
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
			"options": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Options.",
				Optional:         true,
				Computed:         true,
			},
			"other_webmails": {
				Type:        schema.TypeList,
				Description: "Other supported webmails.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of all email traffic.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"pop3": {
				Type:        schema.TypeList,
				Description: "POP3.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"pass", "tag"}, false),

							Description: "Action for spam email.",
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
						"log_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of all email traffic.",
							Optional:    true,
							Computed:    true,
						},
						"tag_msg": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Subject text or header added to spam email.",
							Optional:    true,
							Computed:    true,
						},
						"tag_type": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Tag subject or header for spam email.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"replacemsg_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Replacement message group.",
				Optional:    true,
				Computed:    true,
			},
			"smtp": {
				Type:        schema.TypeList,
				Description: "SMTP.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"pass", "tag", "discard"}, false),

							Description: "Action for spam email.",
							Optional:    true,
							Computed:    true,
						},
						"hdrip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable SMTP email header IP checks for spamfsip, spamrbl, and spambal filters.",
							Optional:    true,
							Computed:    true,
						},
						"local_override": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable local filter to override SMTP remote check result.",
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
						"log_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of all email traffic.",
							Optional:    true,
							Computed:    true,
						},
						"tag_msg": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Subject text or header added to spam email.",
							Optional:    true,
							Computed:    true,
						},
						"tag_type": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Tag subject or header for spam email.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"spam_bal_table": {
				Type: schema.TypeInt,

				Description: "Anti-spam block/allow list table ID.",
				Optional:    true,
				Computed:    true,
			},
			"spam_bwl_table": {
				Type: schema.TypeInt,

				Description: "Anti-spam black/white list table ID.",
				Optional:    true,
				Computed:    true,
			},
			"spam_bword_table": {
				Type: schema.TypeInt,

				Description: "Anti-spam banned word table ID.",
				Optional:    true,
				Computed:    true,
			},
			"spam_bword_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2147483647),

				Description: "Spam banned word threshold.",
				Optional:    true,
				Computed:    true,
			},
			"spam_filtering": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable spam filtering.",
				Optional:    true,
				Computed:    true,
			},
			"spam_iptrust_table": {
				Type: schema.TypeInt,

				Description: "Anti-spam IP trust table ID.",
				Optional:    true,
				Computed:    true,
			},
			"spam_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable spam logging for email filtering.",
				Optional:    true,
				Computed:    true,
			},
			"spam_log_fortiguard_response": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging FortiGuard spam response.",
				Optional:    true,
				Computed:    true,
			},
			"spam_mheader_table": {
				Type: schema.TypeInt,

				Description: "Anti-spam MIME header table ID.",
				Optional:    true,
				Computed:    true,
			},
			"spam_rbl_table": {
				Type: schema.TypeInt,

				Description: "Anti-spam DNSBL table ID.",
				Optional:    true,
				Computed:    true,
			},
			"yahoo_mail": {
				Type:        schema.TypeList,
				Description: "Yahoo! Mail.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging.",
							Optional:    true,
							Computed:    true,
						},
						"log_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of all email traffic.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceEmailfilterProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating EmailfilterProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectEmailfilterProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateEmailfilterProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EmailfilterProfile")
	}

	return resourceEmailfilterProfileRead(ctx, d, meta)
}

func resourceEmailfilterProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectEmailfilterProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateEmailfilterProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating EmailfilterProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EmailfilterProfile")
	}

	return resourceEmailfilterProfileRead(ctx, d, meta)
}

func resourceEmailfilterProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteEmailfilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting EmailfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadEmailfilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading EmailfilterProfile resource: %v", err)
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

	diags := refreshObjectEmailfilterProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenEmailfilterProfileFileFilter(d *schema.ResourceData, v *models.EmailfilterProfileFileFilter, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.EmailfilterProfileFileFilter{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Entries; tmp != nil {
				v["entries"] = flattenEmailfilterProfileFileFilterEntries(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "entries"), sort)
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

func flattenEmailfilterProfileFileFilterEntries(d *schema.ResourceData, v *[]models.EmailfilterProfileFileFilterEntries, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Comment; tmp != nil {
				v["comment"] = *tmp
			}

			if tmp := cfg.FileType; tmp != nil {
				v["file_type"] = flattenEmailfilterProfileFileFilterEntriesFileType(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "file_type"), sort)
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

func flattenEmailfilterProfileFileFilterEntriesFileType(d *schema.ResourceData, v *[]models.EmailfilterProfileFileFilterEntriesFileType, prefix string, sort bool) interface{} {
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

func flattenEmailfilterProfileGmail(d *schema.ResourceData, v *models.EmailfilterProfileGmail, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.EmailfilterProfileGmail{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.LogAll; tmp != nil {
				v["log_all"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenEmailfilterProfileImap(d *schema.ResourceData, v *models.EmailfilterProfileImap, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.EmailfilterProfileImap{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.LogAll; tmp != nil {
				v["log_all"] = *tmp
			}

			if tmp := cfg.TagMsg; tmp != nil {
				v["tag_msg"] = *tmp
			}

			if tmp := cfg.TagType; tmp != nil {
				v["tag_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenEmailfilterProfileMapi(d *schema.ResourceData, v *models.EmailfilterProfileMapi, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.EmailfilterProfileMapi{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.LogAll; tmp != nil {
				v["log_all"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenEmailfilterProfileMsnHotmail(d *schema.ResourceData, v *models.EmailfilterProfileMsnHotmail, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.EmailfilterProfileMsnHotmail{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.LogAll; tmp != nil {
				v["log_all"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenEmailfilterProfileOtherWebmails(d *schema.ResourceData, v *models.EmailfilterProfileOtherWebmails, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.EmailfilterProfileOtherWebmails{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.LogAll; tmp != nil {
				v["log_all"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenEmailfilterProfilePop3(d *schema.ResourceData, v *models.EmailfilterProfilePop3, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.EmailfilterProfilePop3{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.LogAll; tmp != nil {
				v["log_all"] = *tmp
			}

			if tmp := cfg.TagMsg; tmp != nil {
				v["tag_msg"] = *tmp
			}

			if tmp := cfg.TagType; tmp != nil {
				v["tag_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenEmailfilterProfileSmtp(d *schema.ResourceData, v *models.EmailfilterProfileSmtp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.EmailfilterProfileSmtp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Hdrip; tmp != nil {
				v["hdrip"] = *tmp
			}

			if tmp := cfg.LocalOverride; tmp != nil {
				v["local_override"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.LogAll; tmp != nil {
				v["log_all"] = *tmp
			}

			if tmp := cfg.TagMsg; tmp != nil {
				v["tag_msg"] = *tmp
			}

			if tmp := cfg.TagType; tmp != nil {
				v["tag_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenEmailfilterProfileYahooMail(d *schema.ResourceData, v *models.EmailfilterProfileYahooMail, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.EmailfilterProfileYahooMail{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.LogAll; tmp != nil {
				v["log_all"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectEmailfilterProfile(d *schema.ResourceData, o *models.EmailfilterProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.External != nil {
		v := *o.External

		if err = d.Set("external", v); err != nil {
			return diag.Errorf("error reading external: %v", err)
		}
	}

	if o.FeatureSet != nil {
		v := *o.FeatureSet

		if err = d.Set("feature_set", v); err != nil {
			return diag.Errorf("error reading feature_set: %v", err)
		}
	}

	if _, ok := d.GetOk("file_filter"); ok {
		if o.FileFilter != nil {
			if err = d.Set("file_filter", flattenEmailfilterProfileFileFilter(d, o.FileFilter, "file_filter", sort)); err != nil {
				return diag.Errorf("error reading file_filter: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("gmail"); ok {
		if o.Gmail != nil {
			if err = d.Set("gmail", flattenEmailfilterProfileGmail(d, o.Gmail, "gmail", sort)); err != nil {
				return diag.Errorf("error reading gmail: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("imap"); ok {
		if o.Imap != nil {
			if err = d.Set("imap", flattenEmailfilterProfileImap(d, o.Imap, "imap", sort)); err != nil {
				return diag.Errorf("error reading imap: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("mapi"); ok {
		if o.Mapi != nil {
			if err = d.Set("mapi", flattenEmailfilterProfileMapi(d, o.Mapi, "mapi", sort)); err != nil {
				return diag.Errorf("error reading mapi: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("msn_hotmail"); ok {
		if o.MsnHotmail != nil {
			if err = d.Set("msn_hotmail", flattenEmailfilterProfileMsnHotmail(d, o.MsnHotmail, "msn_hotmail", sort)); err != nil {
				return diag.Errorf("error reading msn_hotmail: %v", err)
			}
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

	if _, ok := d.GetOk("other_webmails"); ok {
		if o.OtherWebmails != nil {
			if err = d.Set("other_webmails", flattenEmailfilterProfileOtherWebmails(d, o.OtherWebmails, "other_webmails", sort)); err != nil {
				return diag.Errorf("error reading other_webmails: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("pop3"); ok {
		if o.Pop3 != nil {
			if err = d.Set("pop3", flattenEmailfilterProfilePop3(d, o.Pop3, "pop3", sort)); err != nil {
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

	if _, ok := d.GetOk("smtp"); ok {
		if o.Smtp != nil {
			if err = d.Set("smtp", flattenEmailfilterProfileSmtp(d, o.Smtp, "smtp", sort)); err != nil {
				return diag.Errorf("error reading smtp: %v", err)
			}
		}
	}

	if o.SpamBalTable != nil {
		v := *o.SpamBalTable

		if err = d.Set("spam_bal_table", v); err != nil {
			return diag.Errorf("error reading spam_bal_table: %v", err)
		}
	}

	if o.SpamBwlTable != nil {
		v := *o.SpamBwlTable

		if err = d.Set("spam_bwl_table", v); err != nil {
			return diag.Errorf("error reading spam_bwl_table: %v", err)
		}
	}

	if o.SpamBwordTable != nil {
		v := *o.SpamBwordTable

		if err = d.Set("spam_bword_table", v); err != nil {
			return diag.Errorf("error reading spam_bword_table: %v", err)
		}
	}

	if o.SpamBwordThreshold != nil {
		v := *o.SpamBwordThreshold

		if err = d.Set("spam_bword_threshold", v); err != nil {
			return diag.Errorf("error reading spam_bword_threshold: %v", err)
		}
	}

	if o.SpamFiltering != nil {
		v := *o.SpamFiltering

		if err = d.Set("spam_filtering", v); err != nil {
			return diag.Errorf("error reading spam_filtering: %v", err)
		}
	}

	if o.SpamIptrustTable != nil {
		v := *o.SpamIptrustTable

		if err = d.Set("spam_iptrust_table", v); err != nil {
			return diag.Errorf("error reading spam_iptrust_table: %v", err)
		}
	}

	if o.SpamLog != nil {
		v := *o.SpamLog

		if err = d.Set("spam_log", v); err != nil {
			return diag.Errorf("error reading spam_log: %v", err)
		}
	}

	if o.SpamLogFortiguardResponse != nil {
		v := *o.SpamLogFortiguardResponse

		if err = d.Set("spam_log_fortiguard_response", v); err != nil {
			return diag.Errorf("error reading spam_log_fortiguard_response: %v", err)
		}
	}

	if o.SpamMheaderTable != nil {
		v := *o.SpamMheaderTable

		if err = d.Set("spam_mheader_table", v); err != nil {
			return diag.Errorf("error reading spam_mheader_table: %v", err)
		}
	}

	if o.SpamRblTable != nil {
		v := *o.SpamRblTable

		if err = d.Set("spam_rbl_table", v); err != nil {
			return diag.Errorf("error reading spam_rbl_table: %v", err)
		}
	}

	if _, ok := d.GetOk("yahoo_mail"); ok {
		if o.YahooMail != nil {
			if err = d.Set("yahoo_mail", flattenEmailfilterProfileYahooMail(d, o.YahooMail, "yahoo_mail", sort)); err != nil {
				return diag.Errorf("error reading yahoo_mail: %v", err)
			}
		}
	}

	return nil
}

func expandEmailfilterProfileFileFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.EmailfilterProfileFileFilter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterProfileFileFilter

	for i := range l {
		tmp := models.EmailfilterProfileFileFilter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.entries", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandEmailfilterProfileFileFilterEntries(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.EmailfilterProfileFileFilterEntries
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
	return &result[0], nil
}

func expandEmailfilterProfileFileFilterEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.EmailfilterProfileFileFilterEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterProfileFileFilterEntries

	for i := range l {
		tmp := models.EmailfilterProfileFileFilterEntries{}
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

		pre_append = fmt.Sprintf("%s.%d.file_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandEmailfilterProfileFileFilterEntriesFileType(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.EmailfilterProfileFileFilterEntriesFileType
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

func expandEmailfilterProfileFileFilterEntriesFileType(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.EmailfilterProfileFileFilterEntriesFileType, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterProfileFileFilterEntriesFileType

	for i := range l {
		tmp := models.EmailfilterProfileFileFilterEntriesFileType{}
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

func expandEmailfilterProfileGmail(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.EmailfilterProfileGmail, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterProfileGmail

	for i := range l {
		tmp := models.EmailfilterProfileGmail{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogAll = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandEmailfilterProfileImap(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.EmailfilterProfileImap, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterProfileImap

	for i := range l {
		tmp := models.EmailfilterProfileImap{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag_msg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TagMsg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TagType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandEmailfilterProfileMapi(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.EmailfilterProfileMapi, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterProfileMapi

	for i := range l {
		tmp := models.EmailfilterProfileMapi{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogAll = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandEmailfilterProfileMsnHotmail(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.EmailfilterProfileMsnHotmail, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterProfileMsnHotmail

	for i := range l {
		tmp := models.EmailfilterProfileMsnHotmail{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogAll = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandEmailfilterProfileOtherWebmails(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.EmailfilterProfileOtherWebmails, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterProfileOtherWebmails

	for i := range l {
		tmp := models.EmailfilterProfileOtherWebmails{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.log_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogAll = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandEmailfilterProfilePop3(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.EmailfilterProfilePop3, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterProfilePop3

	for i := range l {
		tmp := models.EmailfilterProfilePop3{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag_msg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TagMsg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TagType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandEmailfilterProfileSmtp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.EmailfilterProfileSmtp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterProfileSmtp

	for i := range l {
		tmp := models.EmailfilterProfileSmtp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hdrip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Hdrip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.local_override", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LocalOverride = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag_msg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TagMsg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TagType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandEmailfilterProfileYahooMail(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.EmailfilterProfileYahooMail, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterProfileYahooMail

	for i := range l {
		tmp := models.EmailfilterProfileYahooMail{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogAll = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func getObjectEmailfilterProfile(d *schema.ResourceData, sv string) (*models.EmailfilterProfile, diag.Diagnostics) {
	obj := models.EmailfilterProfile{}
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
	if v1, ok := d.GetOk("external"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("external", sv)
				diags = append(diags, e)
			}
			obj.External = &v2
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
		t, err := expandEmailfilterProfileFileFilter(d, v, "file_filter", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FileFilter = t
		}
	} else if d.HasChange("file_filter") {
		old, new := d.GetChange("file_filter")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FileFilter = &models.EmailfilterProfileFileFilter{}
		}
	}
	if v, ok := d.GetOk("gmail"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("gmail", sv)
			diags = append(diags, e)
		}
		t, err := expandEmailfilterProfileGmail(d, v, "gmail", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Gmail = t
		}
	} else if d.HasChange("gmail") {
		old, new := d.GetChange("gmail")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Gmail = &models.EmailfilterProfileGmail{}
		}
	}
	if v, ok := d.GetOk("imap"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("imap", sv)
			diags = append(diags, e)
		}
		t, err := expandEmailfilterProfileImap(d, v, "imap", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Imap = t
		}
	} else if d.HasChange("imap") {
		old, new := d.GetChange("imap")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Imap = &models.EmailfilterProfileImap{}
		}
	}
	if v, ok := d.GetOk("mapi"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("mapi", sv)
			diags = append(diags, e)
		}
		t, err := expandEmailfilterProfileMapi(d, v, "mapi", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Mapi = t
		}
	} else if d.HasChange("mapi") {
		old, new := d.GetChange("mapi")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Mapi = &models.EmailfilterProfileMapi{}
		}
	}
	if v, ok := d.GetOk("msn_hotmail"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("msn_hotmail", sv)
			diags = append(diags, e)
		}
		t, err := expandEmailfilterProfileMsnHotmail(d, v, "msn_hotmail", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.MsnHotmail = t
		}
	} else if d.HasChange("msn_hotmail") {
		old, new := d.GetChange("msn_hotmail")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.MsnHotmail = &models.EmailfilterProfileMsnHotmail{}
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
	if v, ok := d.GetOk("other_webmails"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "") {
			e := utils.AttributeVersionWarning("other_webmails", sv)
			diags = append(diags, e)
		}
		t, err := expandEmailfilterProfileOtherWebmails(d, v, "other_webmails", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.OtherWebmails = t
		}
	} else if d.HasChange("other_webmails") {
		old, new := d.GetChange("other_webmails")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.OtherWebmails = &models.EmailfilterProfileOtherWebmails{}
		}
	}
	if v, ok := d.GetOk("pop3"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("pop3", sv)
			diags = append(diags, e)
		}
		t, err := expandEmailfilterProfilePop3(d, v, "pop3", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Pop3 = t
		}
	} else if d.HasChange("pop3") {
		old, new := d.GetChange("pop3")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Pop3 = &models.EmailfilterProfilePop3{}
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
	if v, ok := d.GetOk("smtp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("smtp", sv)
			diags = append(diags, e)
		}
		t, err := expandEmailfilterProfileSmtp(d, v, "smtp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Smtp = t
		}
	} else if d.HasChange("smtp") {
		old, new := d.GetChange("smtp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Smtp = &models.EmailfilterProfileSmtp{}
		}
	}
	if v1, ok := d.GetOk("spam_bal_table"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("spam_bal_table", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SpamBalTable = &tmp
		}
	}
	if v1, ok := d.GetOk("spam_bwl_table"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("spam_bwl_table", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SpamBwlTable = &tmp
		}
	}
	if v1, ok := d.GetOk("spam_bword_table"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spam_bword_table", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SpamBwordTable = &tmp
		}
	}
	if v1, ok := d.GetOk("spam_bword_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spam_bword_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SpamBwordThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("spam_filtering"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spam_filtering", sv)
				diags = append(diags, e)
			}
			obj.SpamFiltering = &v2
		}
	}
	if v1, ok := d.GetOk("spam_iptrust_table"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spam_iptrust_table", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SpamIptrustTable = &tmp
		}
	}
	if v1, ok := d.GetOk("spam_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spam_log", sv)
				diags = append(diags, e)
			}
			obj.SpamLog = &v2
		}
	}
	if v1, ok := d.GetOk("spam_log_fortiguard_response"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spam_log_fortiguard_response", sv)
				diags = append(diags, e)
			}
			obj.SpamLogFortiguardResponse = &v2
		}
	}
	if v1, ok := d.GetOk("spam_mheader_table"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spam_mheader_table", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SpamMheaderTable = &tmp
		}
	}
	if v1, ok := d.GetOk("spam_rbl_table"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spam_rbl_table", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SpamRblTable = &tmp
		}
	}
	if v, ok := d.GetOk("yahoo_mail"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("yahoo_mail", sv)
			diags = append(diags, e)
		}
		t, err := expandEmailfilterProfileYahooMail(d, v, "yahoo_mail", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.YahooMail = t
		}
	} else if d.HasChange("yahoo_mail") {
		old, new := d.GetChange("yahoo_mail")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.YahooMail = &models.EmailfilterProfileYahooMail{}
		}
	}
	return &obj, diags
}
