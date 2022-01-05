// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceDlpFpDocSource() *schema.Resource {
	return &schema.Resource{
		Description: "Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.",

		ReadContext: dataSourceDlpFpDocSourceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"date": {
				Type:        schema.TypeInt,
				Description: "Day of the month on which to scan the server (1 - 31).",
				Computed:    true,
			},
			"file_path": {
				Type:        schema.TypeString,
				Description: "Path on the server to the fingerprint files (max 119 characters).",
				Computed:    true,
			},
			"file_pattern": {
				Type:        schema.TypeString,
				Description: "Files matching this pattern on the server are fingerprinted. Optionally use the * and ? wildcards.",
				Computed:    true,
			},
			"keep_modified": {
				Type:        schema.TypeString,
				Description: "Enable so that when a file is changed on the server the FortiGate keeps the old fingerprint and adds a new fingerprint to the database.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the DLP fingerprint database.",
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password required to log into the file server.",
				Computed:    true,
				Sensitive:   true,
			},
			"period": {
				Type:        schema.TypeString,
				Description: "Frequency for which the FortiGate checks the server for new or changed files.",
				Computed:    true,
			},
			"remove_deleted": {
				Type:        schema.TypeString,
				Description: "Enable to keep the fingerprint database up to date when a file is deleted from the server.",
				Computed:    true,
			},
			"scan_on_creation": {
				Type:        schema.TypeString,
				Description: "Enable to keep the fingerprint database up to date when a file is added or changed on the server.",
				Computed:    true,
			},
			"scan_subdirectories": {
				Type:        schema.TypeString,
				Description: "Enable/disable scanning subdirectories to find files to create fingerprints from.",
				Computed:    true,
			},
			"sensitivity": {
				Type:        schema.TypeString,
				Description: "Select a sensitivity or threat level for matches with this fingerprint database. Add sensitivities using sensitivity.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "IPv4 or IPv6 address of the server.",
				Computed:    true,
			},
			"server_type": {
				Type:        schema.TypeString,
				Description: "Protocol used to communicate with the file server. Currently only Samba (SMB) servers are supported.",
				Computed:    true,
			},
			"tod_hour": {
				Type:        schema.TypeInt,
				Description: "Hour of the day on which to scan the server (0 - 23, default = 1).",
				Computed:    true,
			},
			"tod_min": {
				Type:        schema.TypeInt,
				Description: "Minute of the hour on which to scan the server (0 - 59).",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "User name required to log into the file server.",
				Computed:    true,
			},
			"vdom": {
				Type:        schema.TypeString,
				Description: "Select the VDOM that can communicate with the file server.",
				Computed:    true,
			},
			"weekday": {
				Type:        schema.TypeString,
				Description: "Day of the week on which to scan the server.",
				Computed:    true,
			},
		},
	}
}

func dataSourceDlpFpDocSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadDlpFpDocSource(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading DlpFpDocSource dataSource: %v", err)
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

	diags := refreshObjectDlpFpDocSource(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
