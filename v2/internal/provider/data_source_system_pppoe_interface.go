// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure the PPPoE interfaces.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemPppoeInterface() *schema.Resource {
	return &schema.Resource{
		Description: "Configure the PPPoE interfaces.",

		ReadContext: dataSourceSystemPppoeInterfaceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"ac_name": {
				Type:        schema.TypeString,
				Description: "PPPoE AC name.",
				Computed:    true,
			},
			"auth_type": {
				Type:        schema.TypeString,
				Description: "PPP authentication type to use.",
				Computed:    true,
			},
			"device": {
				Type:        schema.TypeString,
				Description: "Name for the physical interface.",
				Computed:    true,
			},
			"dial_on_demand": {
				Type:        schema.TypeString,
				Description: "Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface.",
				Computed:    true,
			},
			"disc_retry_timeout": {
				Type:        schema.TypeInt,
				Description: "PPPoE discovery init timeout value in (0-4294967295 sec).",
				Computed:    true,
			},
			"idle_timeout": {
				Type:        schema.TypeInt,
				Description: "PPPoE auto disconnect after idle timeout (0-4294967295 sec).",
				Computed:    true,
			},
			"ipunnumbered": {
				Type:        schema.TypeString,
				Description: "PPPoE unnumbered IP.",
				Computed:    true,
			},
			"ipv6": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 Control Protocol (IPv6CP).",
				Computed:    true,
			},
			"lcp_echo_interval": {
				Type:        schema.TypeInt,
				Description: "Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.",
				Computed:    true,
			},
			"lcp_max_echo_fails": {
				Type:        schema.TypeInt,
				Description: "Maximum missed LCP echo messages before disconnect.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the PPPoE interface.",
				Required:    true,
			},
			"padt_retry_timeout": {
				Type:        schema.TypeInt,
				Description: "PPPoE terminate timeout value in (0-4294967295 sec).",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Enter the password.",
				Computed:    true,
				Sensitive:   true,
			},
			"pppoe_unnumbered_negotiate": {
				Type:        schema.TypeString,
				Description: "Enable/disable PPPoE unnumbered negotiation.",
				Computed:    true,
			},
			"service_name": {
				Type:        schema.TypeString,
				Description: "PPPoE service name.",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "User name.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemPppoeInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemPppoeInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemPppoeInterface dataSource: %v", err)
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

	diags := refreshObjectSystemPppoeInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
