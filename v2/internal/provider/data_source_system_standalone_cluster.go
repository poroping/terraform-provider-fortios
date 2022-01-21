// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemStandaloneCluster() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.",

		ReadContext: dataSourceSystemStandaloneClusterRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"encryption": {
				Type:        schema.TypeString,
				Description: "Enable/disable encryption when synchronizing sessions.",
				Computed:    true,
			},
			"group_member_id": {
				Type:        schema.TypeInt,
				Description: "Cluster member ID (0 - 15).",
				Computed:    true,
			},
			"layer2_connection": {
				Type:        schema.TypeString,
				Description: "Indicate whether layer 2 connections are present among FGSP members.",
				Computed:    true,
			},
			"psksecret": {
				Type:        schema.TypeString,
				Description: "Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).",
				Computed:    true,
				Sensitive:   true,
			},
			"session_sync_dev": {
				Type:        schema.TypeString,
				Description: "Offload session-sync process to kernel and sync sessions using connected interface(s) directly.",
				Computed:    true,
			},
			"standalone_group_id": {
				Type:        schema.TypeInt,
				Description: "Cluster group ID (0 - 255). Must be the same for all members.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemStandaloneClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemStandaloneCluster"

	o, err := c.Cmdb.ReadSystemStandaloneCluster(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemStandaloneCluster dataSource: %v", err)
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

	diags := refreshObjectSystemStandaloneCluster(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
