// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure quarantine options.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceAntivirusQuarantine() *schema.Resource {
	return &schema.Resource{
		Description: "Configure quarantine options.",

		ReadContext: dataSourceAntivirusQuarantineRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"agelimit": {
				Type:        schema.TypeInt,
				Description: "Age limit for quarantined files (0 - 479 hours, 0 means forever).",
				Computed:    true,
			},
			"destination": {
				Type:        schema.TypeString,
				Description: "Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them.",
				Computed:    true,
			},
			"drop_blocked": {
				Type:        schema.TypeString,
				Description: "Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.",
				Computed:    true,
			},
			"drop_heuristic": {
				Type:        schema.TypeString,
				Description: "Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.",
				Computed:    true,
			},
			"drop_infected": {
				Type:        schema.TypeString,
				Description: "Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.",
				Computed:    true,
			},
			"drop_machine_learning": {
				Type:        schema.TypeString,
				Description: "Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.",
				Computed:    true,
			},
			"lowspace": {
				Type:        schema.TypeString,
				Description: "Select the method for handling additional files when running low on disk space.",
				Computed:    true,
			},
			"maxfilesize": {
				Type:        schema.TypeInt,
				Description: "Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).",
				Computed:    true,
			},
			"quarantine_quota": {
				Type:        schema.TypeInt,
				Description: "The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).",
				Computed:    true,
			},
			"store_blocked": {
				Type:        schema.TypeString,
				Description: "Quarantine blocked files found in sessions using the selected protocols.",
				Computed:    true,
			},
			"store_heuristic": {
				Type:        schema.TypeString,
				Description: "Quarantine files detected by heuristics found in sessions using the selected protocols.",
				Computed:    true,
			},
			"store_infected": {
				Type:        schema.TypeString,
				Description: "Quarantine infected files found in sessions using the selected protocols.",
				Computed:    true,
			},
			"store_machine_learning": {
				Type:        schema.TypeString,
				Description: "Quarantine files detected by machine learning found in sessions using the selected protocols.",
				Computed:    true,
			},
		},
	}
}

func dataSourceAntivirusQuarantineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "AntivirusQuarantine"

	o, err := c.Cmdb.ReadAntivirusQuarantine(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AntivirusQuarantine dataSource: %v", err)
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

	diags := refreshObjectAntivirusQuarantine(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
