// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Email Filter profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceEmailfilterProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Email Filter profiles.",

		ReadContext: dataSourceEmailfilterProfileRead,

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
			"external": {
				Type:        schema.TypeString,
				Description: "Enable/disable external Email inspection.",
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
			"gmail": {
				Type:        schema.TypeList,
				Description: "Gmail.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"log_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of all email traffic.",
							Computed:    true,
						},
					},
				},
			},
			"imap": {
				Type:        schema.TypeList,
				Description: "IMAP.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action for spam email.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"log_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of all email traffic.",
							Computed:    true,
						},
						"tag_msg": {
							Type:        schema.TypeString,
							Description: "Subject text or header added to spam email.",
							Computed:    true,
						},
						"tag_type": {
							Type:        schema.TypeString,
							Description: "Tag subject or header for spam email.",
							Computed:    true,
						},
					},
				},
			},
			"mapi": {
				Type:        schema.TypeList,
				Description: "MAPI.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action for spam email.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"log_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of all email traffic.",
							Computed:    true,
						},
					},
				},
			},
			"msn_hotmail": {
				Type:        schema.TypeList,
				Description: "MSN Hotmail.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"log_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of all email traffic.",
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
			"options": {
				Type:        schema.TypeString,
				Description: "Options.",
				Computed:    true,
			},
			"other_webmails": {
				Type:        schema.TypeList,
				Description: "Other supported webmails.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of all email traffic.",
							Computed:    true,
						},
					},
				},
			},
			"pop3": {
				Type:        schema.TypeList,
				Description: "POP3.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action for spam email.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"log_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of all email traffic.",
							Computed:    true,
						},
						"tag_msg": {
							Type:        schema.TypeString,
							Description: "Subject text or header added to spam email.",
							Computed:    true,
						},
						"tag_type": {
							Type:        schema.TypeString,
							Description: "Tag subject or header for spam email.",
							Computed:    true,
						},
					},
				},
			},
			"replacemsg_group": {
				Type:        schema.TypeString,
				Description: "Replacement message group.",
				Computed:    true,
			},
			"smtp": {
				Type:        schema.TypeList,
				Description: "SMTP.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action for spam email.",
							Computed:    true,
						},
						"hdrip": {
							Type:        schema.TypeString,
							Description: "Enable/disable SMTP email header IP checks for spamfsip, spamrbl, and spambal filters.",
							Computed:    true,
						},
						"local_override": {
							Type:        schema.TypeString,
							Description: "Enable/disable local filter to override SMTP remote check result.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"log_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of all email traffic.",
							Computed:    true,
						},
						"tag_msg": {
							Type:        schema.TypeString,
							Description: "Subject text or header added to spam email.",
							Computed:    true,
						},
						"tag_type": {
							Type:        schema.TypeString,
							Description: "Tag subject or header for spam email.",
							Computed:    true,
						},
					},
				},
			},
			"spam_bal_table": {
				Type:        schema.TypeInt,
				Description: "Anti-spam block/allow list table ID.",
				Computed:    true,
			},
			"spam_bwl_table": {
				Type:        schema.TypeInt,
				Description: "Anti-spam black/white list table ID.",
				Computed:    true,
			},
			"spam_bword_table": {
				Type:        schema.TypeInt,
				Description: "Anti-spam banned word table ID.",
				Computed:    true,
			},
			"spam_bword_threshold": {
				Type:        schema.TypeInt,
				Description: "Spam banned word threshold.",
				Computed:    true,
			},
			"spam_filtering": {
				Type:        schema.TypeString,
				Description: "Enable/disable spam filtering.",
				Computed:    true,
			},
			"spam_iptrust_table": {
				Type:        schema.TypeInt,
				Description: "Anti-spam IP trust table ID.",
				Computed:    true,
			},
			"spam_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable spam logging for email filtering.",
				Computed:    true,
			},
			"spam_log_fortiguard_response": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging FortiGuard spam response.",
				Computed:    true,
			},
			"spam_mheader_table": {
				Type:        schema.TypeInt,
				Description: "Anti-spam MIME header table ID.",
				Computed:    true,
			},
			"spam_rbl_table": {
				Type:        schema.TypeInt,
				Description: "Anti-spam DNSBL table ID.",
				Computed:    true,
			},
			"yahoo_mail": {
				Type:        schema.TypeList,
				Description: "Yahoo! Mail.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"log_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of all email traffic.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEmailfilterProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadEmailfilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading EmailfilterProfile dataSource: %v", err)
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

	diags := refreshObjectEmailfilterProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
