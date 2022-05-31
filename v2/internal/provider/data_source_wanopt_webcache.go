// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure global Web cache settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWanoptWebcache() *schema.Resource {
	return &schema.Resource{
		Description: "Configure global Web cache settings.",

		ReadContext: dataSourceWanoptWebcacheRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"always_revalidate": {
				Type:        schema.TypeString,
				Description: "Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client.",
				Computed:    true,
			},
			"cache_by_default": {
				Type:        schema.TypeString,
				Description: "Enable/disable caching content that lacks explicit caching policies from the server.",
				Computed:    true,
			},
			"cache_cookie": {
				Type:        schema.TypeString,
				Description: "Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached.",
				Computed:    true,
			},
			"cache_expired": {
				Type:        schema.TypeString,
				Description: "Enable/disable caching type-1 objects that are already expired on arrival.",
				Computed:    true,
			},
			"default_ttl": {
				Type:        schema.TypeInt,
				Description: "Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.",
				Computed:    true,
			},
			"external": {
				Type:        schema.TypeString,
				Description: "Enable/disable external Web caching.",
				Computed:    true,
			},
			"fresh_factor": {
				Type:        schema.TypeInt,
				Description: "Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.",
				Computed:    true,
			},
			"host_validate": {
				Type:        schema.TypeString,
				Description: "Enable/disable validating \"Host:\" with original server IP.",
				Computed:    true,
			},
			"ignore_conditional": {
				Type:        schema.TypeString,
				Description: "Enable/disable controlling the behavior of cache-control HTTP 1.1 header values.",
				Computed:    true,
			},
			"ignore_ie_reload": {
				Type:        schema.TypeString,
				Description: "Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header.",
				Computed:    true,
			},
			"ignore_ims": {
				Type:        schema.TypeString,
				Description: "Enable/disable ignoring the if-modified-since (IMS) header.",
				Computed:    true,
			},
			"ignore_pnc": {
				Type:        schema.TypeString,
				Description: "Enable/disable ignoring the pragma no-cache (PNC) header.",
				Computed:    true,
			},
			"max_object_size": {
				Type:        schema.TypeInt,
				Description: "Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.",
				Computed:    true,
			},
			"max_ttl": {
				Type:        schema.TypeInt,
				Description: "Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).",
				Computed:    true,
			},
			"min_ttl": {
				Type:        schema.TypeInt,
				Description: "Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).",
				Computed:    true,
			},
			"neg_resp_time": {
				Type:        schema.TypeInt,
				Description: "Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).",
				Computed:    true,
			},
			"reval_pnc": {
				Type:        schema.TypeString,
				Description: "Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWanoptWebcacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "WanoptWebcache"

	o, err := c.Cmdb.ReadWanoptWebcache(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WanoptWebcache dataSource: %v", err)
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

	diags := refreshObjectWanoptWebcache(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
