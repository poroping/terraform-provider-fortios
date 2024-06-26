// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DNS servers for a non-management VDOM.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemVdomDns() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DNS servers for a non-management VDOM.",

		ReadContext: dataSourceSystemVdomDnsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"alt_primary": {
				Type:        schema.TypeString,
				Description: "Alternate primary DNS server. This is not used as a failover DNS server.",
				Computed:    true,
			},
			"alt_secondary": {
				Type:        schema.TypeString,
				Description: "Alternate secondary DNS server. This is not used as a failover DNS server.",
				Computed:    true,
			},
			"dns_over_tls": {
				Type:        schema.TypeString,
				Description: "Enable/disable/enforce DNS over TLS.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Specify outgoing interface to reach server.",
				Computed:    true,
			},
			"interface_select_method": {
				Type:        schema.TypeString,
				Description: "Specify how to select outgoing interface to reach server.",
				Computed:    true,
			},
			"ip6_primary": {
				Type:        schema.TypeString,
				Description: "Primary IPv6 DNS server IP address for the VDOM.",
				Computed:    true,
			},
			"ip6_secondary": {
				Type:        schema.TypeString,
				Description: "Secondary IPv6 DNS server IP address for the VDOM.",
				Computed:    true,
			},
			"primary": {
				Type:        schema.TypeString,
				Description: "Primary DNS server IP address for the VDOM.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "DNS transport protocols.",
				Computed:    true,
			},
			"secondary": {
				Type:        schema.TypeString,
				Description: "Secondary DNS server IP address for the VDOM.",
				Computed:    true,
			},
			"server_hostname": {
				Type:        schema.TypeList,
				Description: "DNS server host name list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": {
							Type:        schema.TypeString,
							Description: "DNS server host name list separated by space (maximum 4 domains).",
							Computed:    true,
						},
					},
				},
			},
			"server_select_method": {
				Type:        schema.TypeString,
				Description: "Specify how configured servers are prioritized.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP for communications with the DNS server.",
				Computed:    true,
			},
			"ssl_certificate": {
				Type:        schema.TypeString,
				Description: "Name of local certificate for SSL connections.",
				Computed:    true,
			},
			"vdom_dns": {
				Type:        schema.TypeString,
				Description: "Enable/disable configuring DNS servers for the current VDOM.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemVdomDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemVdomDns"

	o, err := c.Cmdb.ReadSystemVdomDns(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVdomDns dataSource: %v", err)
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

	diags := refreshObjectSystemVdomDns(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
