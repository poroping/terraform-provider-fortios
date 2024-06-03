// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WCCP.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemWccp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WCCP.",

		ReadContext: dataSourceSystemWccpRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"assignment_bucket_format": {
				Type:        schema.TypeString,
				Description: "Assignment bucket format for the WCCP cache engine.",
				Computed:    true,
			},
			"assignment_dstaddr_mask": {
				Type:        schema.TypeString,
				Description: "Assignment destination address mask.",
				Computed:    true,
			},
			"assignment_method": {
				Type:        schema.TypeString,
				Description: "Hash key assignment preference.",
				Computed:    true,
			},
			"assignment_srcaddr_mask": {
				Type:        schema.TypeString,
				Description: "Assignment source address mask.",
				Computed:    true,
			},
			"assignment_weight": {
				Type:        schema.TypeInt,
				Description: "Assignment of hash weight/ratio for the WCCP cache engine.",
				Computed:    true,
			},
			"authentication": {
				Type:        schema.TypeString,
				Description: "Enable/disable MD5 authentication.",
				Computed:    true,
			},
			"cache_engine_method": {
				Type:        schema.TypeString,
				Description: "Method used to forward traffic to the routers or to return to the cache engine.",
				Computed:    true,
			},
			"cache_id": {
				Type:        schema.TypeString,
				Description: "IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.",
				Computed:    true,
			},
			"forward_method": {
				Type:        schema.TypeString,
				Description: "Method used to forward traffic to the cache servers.",
				Computed:    true,
			},
			"group_address": {
				Type:        schema.TypeString,
				Description: "IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password for MD5 authentication.",
				Computed:    true,
				Sensitive:   true,
			},
			"ports": {
				Type:        schema.TypeString,
				Description: "Service ports.",
				Computed:    true,
			},
			"ports_defined": {
				Type:        schema.TypeString,
				Description: "Match method.",
				Computed:    true,
			},
			"primary_hash": {
				Type:        schema.TypeString,
				Description: "Hash method.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Service priority.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeInt,
				Description: "Service protocol.",
				Computed:    true,
			},
			"return_method": {
				Type:        schema.TypeString,
				Description: "Method used to decline a redirected packet and return it to the FortiGate unit.",
				Computed:    true,
			},
			"router_id": {
				Type:        schema.TypeString,
				Description: "IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.",
				Computed:    true,
			},
			"router_list": {
				Type:        schema.TypeString,
				Description: "IP addresses of one or more WCCP routers.",
				Computed:    true,
			},
			"server_list": {
				Type:        schema.TypeString,
				Description: "IP addresses and netmasks for up to four cache servers.",
				Computed:    true,
			},
			"server_type": {
				Type:        schema.TypeString,
				Description: "Cache server type.",
				Computed:    true,
			},
			"service_id": {
				Type:        schema.TypeString,
				Description: "Service ID.",
				Required:    true,
			},
			"service_type": {
				Type:        schema.TypeString,
				Description: "WCCP service type used by the cache server for logical interception and redirection of traffic.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemWccpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("service_id")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadSystemWccp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemWccp dataSource: %v", err)
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

	diags := refreshObjectSystemWccp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
