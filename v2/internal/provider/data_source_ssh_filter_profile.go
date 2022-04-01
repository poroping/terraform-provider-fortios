// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SSH filter profile.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSshFilterProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SSH filter profile.",

		ReadContext: dataSourceSshFilterProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"block": {
				Type:        schema.TypeString,
				Description: "SSH blocking options.",
				Computed:    true,
			},
			"default_command_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging unmatched shell commands.",
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
			"log": {
				Type:        schema.TypeString,
				Description: "SSH logging options.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "SSH filter profile name.",
				Required:    true,
			},
			"shell_commands": {
				Type:        schema.TypeList,
				Description: "SSH command filter.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action to take for SSH shell command matches.",
							Computed:    true,
						},
						"alert": {
							Type:        schema.TypeString,
							Description: "Enable/disable alert.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Id.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"pattern": {
							Type:        schema.TypeString,
							Description: "SSH shell command pattern.",
							Computed:    true,
						},
						"severity": {
							Type:        schema.TypeString,
							Description: "Log severity.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Matching type.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSshFilterProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSshFilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SshFilterProfile dataSource: %v", err)
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

	diags := refreshObjectSshFilterProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
