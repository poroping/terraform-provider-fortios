// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure file-filter profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFileFilterProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure file-filter profiles.",

		ReadContext: dataSourceFileFilterProfileRead,

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
			"extended_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable file-filter extended logging.",
				Computed:    true,
			},
			"feature_set": {
				Type:        schema.TypeString,
				Description: "Flow/proxy feature set.",
				Computed:    true,
			},
			"log": {
				Type:        schema.TypeString,
				Description: "Enable/disable file-filter logging.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Profile name.",
				Required:    true,
			},
			"replacemsg_group": {
				Type:        schema.TypeString,
				Description: "Replacement message group.",
				Computed:    true,
			},
			"rules": {
				Type:        schema.TypeList,
				Description: "File filter rules.",
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
							Description: "Traffic direction (HTTP, FTP, SSH, CIFS only).",
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
						"name": {
							Type:        schema.TypeString,
							Description: "File-filter rule name.",
							Computed:    true,
						},
						"password_protected": {
							Type:        schema.TypeString,
							Description: "Match password-protected files.",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeString,
							Description: "Protocols to apply rule to.",
							Computed:    true,
						},
					},
				},
			},
			"scan_archive_contents": {
				Type:        schema.TypeString,
				Description: "Enable/disable archive contents scan.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFileFilterProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFileFilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FileFilterProfile dataSource: %v", err)
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

	diags := refreshObjectFileFilterProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
