// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Designate logical storage for DLP fingerprint database.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceDlpSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Designate logical storage for DLP fingerprint database.",

		ReadContext: dataSourceDlpSettingsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"cache_mem_percent": {
				Type:        schema.TypeInt,
				Description: "Maximum percentage of available memory allocated to caching (1 - 15).",
				Computed:    true,
			},
			"chunk_size": {
				Type:        schema.TypeInt,
				Description: "Maximum fingerprint chunk size. Caution, changing this setting will flush the entire database.",
				Computed:    true,
			},
			"db_mode": {
				Type:        schema.TypeString,
				Description: "Behavior when the maximum size is reached.",
				Computed:    true,
			},
			"size": {
				Type:        schema.TypeInt,
				Description: "Maximum total size of files within the storage (MB).",
				Computed:    true,
			},
			"storage_device": {
				Type:        schema.TypeString,
				Description: "Storage device name.",
				Computed:    true,
			},
		},
	}
}

func dataSourceDlpSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "DlpSettings"

	o, err := c.Cmdb.ReadDlpSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading DlpSettings dataSource: %v", err)
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

	diags := refreshObjectDlpSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
