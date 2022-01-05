// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure profile groups.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallProfileGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure profile groups.",

		ReadContext: dataSourceFirewallProfileGroupRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"application_list": {
				Type:        schema.TypeString,
				Description: "Name of an existing Application list.",
				Computed:    true,
			},
			"av_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing Antivirus profile.",
				Computed:    true,
			},
			"cifs_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing CIFS profile.",
				Computed:    true,
			},
			"dlp_sensor": {
				Type:        schema.TypeString,
				Description: "Name of an existing DLP sensor.",
				Computed:    true,
			},
			"dnsfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing DNS filter profile.",
				Computed:    true,
			},
			"emailfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing email filter profile.",
				Computed:    true,
			},
			"file_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing file-filter profile.",
				Computed:    true,
			},
			"icap_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing ICAP profile.",
				Computed:    true,
			},
			"ips_sensor": {
				Type:        schema.TypeString,
				Description: "Name of an existing IPS sensor.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Profile group name.",
				Required:    true,
			},
			"profile_protocol_options": {
				Type:        schema.TypeString,
				Description: "Name of an existing Protocol options profile.",
				Computed:    true,
			},
			"sctp_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing SCTP filter profile.",
				Computed:    true,
			},
			"ssh_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing SSH filter profile.",
				Computed:    true,
			},
			"ssl_ssh_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing SSL SSH profile.",
				Computed:    true,
			},
			"videofilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing VideoFilter profile.",
				Computed:    true,
			},
			"voip_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing VoIP profile.",
				Computed:    true,
			},
			"waf_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing Web application firewall profile.",
				Computed:    true,
			},
			"webfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing Web filter profile.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallProfileGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallProfileGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallProfileGroup dataSource: %v", err)
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

	diags := refreshObjectFirewallProfileGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
