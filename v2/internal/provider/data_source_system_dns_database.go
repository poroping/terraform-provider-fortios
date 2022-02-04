// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DNS databases.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemDnsDatabase() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DNS databases.",

		ReadContext: dataSourceSystemDnsDatabaseRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_transfer": {
				Type:        schema.TypeString,
				Description: "DNS zone transfer IP address list.",
				Computed:    true,
			},
			"authoritative": {
				Type:        schema.TypeString,
				Description: "Enable/disable authoritative zone.",
				Computed:    true,
			},
			"contact": {
				Type:        schema.TypeString,
				Description: "Email address of the administrator for this zone. You can specify only the username, such as admin or the full email address, such as admin@test.com When using only a username, the domain of the email will be this zone.",
				Computed:    true,
			},
			"dns_entry": {
				Type:        schema.TypeList,
				Description: "DNS entry.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"canonical_name": {
							Type:        schema.TypeString,
							Description: "Canonical name of the host.",
							Computed:    true,
						},
						"hostname": {
							Type:        schema.TypeString,
							Description: "Name of the host.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "DNS entry ID.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "IPv4 address of the host.",
							Computed:    true,
						},
						"ipv6": {
							Type:        schema.TypeString,
							Description: "IPv6 address of the host.",
							Computed:    true,
						},
						"preference": {
							Type:        schema.TypeInt,
							Description: "DNS entry preference (0 - 65535, highest preference = 0, default = 10).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable resource record status.",
							Computed:    true,
						},
						"ttl": {
							Type:        schema.TypeInt,
							Description: "Time-to-live for this entry (0 to 2147483647 sec, default = 0).",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Resource record type.",
							Computed:    true,
						},
					},
				},
			},
			"domain": {
				Type:        schema.TypeString,
				Description: "Domain name.",
				Computed:    true,
			},
			"forwarder": {
				Type:        schema.TypeString,
				Description: "DNS zone forwarder IP address list.",
				Computed:    true,
			},
			"ip_master": {
				Type:        schema.TypeString,
				Description: "IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.",
				Computed:    true,
			},
			"ip_primary": {
				Type:        schema.TypeString,
				Description: "IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Zone name.",
				Required:    true,
			},
			"primary_name": {
				Type:        schema.TypeString,
				Description: "Domain name of the default DNS server for this zone.",
				Computed:    true,
			},
			"rr_max": {
				Type:        schema.TypeInt,
				Description: "Maximum number of resource records (10 - 65536, 0 means infinite).",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP for forwarding to DNS server.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this DNS zone.",
				Computed:    true,
			},
			"ttl": {
				Type:        schema.TypeInt,
				Description: "Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Zone type (primary to manage entries directly, secondary to import entries from other zones).",
				Computed:    true,
			},
			"view": {
				Type:        schema.TypeString,
				Description: "Zone view (public to serve public clients, shadow to serve internal clients).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemDnsDatabaseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemDnsDatabase(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemDnsDatabase dataSource: %v", err)
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

	diags := refreshObjectSystemDnsDatabase(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
