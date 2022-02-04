// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Internet Service definition.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallInternetServiceDefinition() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Internet Service definition.",

		ReadContext: dataSourceFirewallInternetServiceDefinitionRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"entry": {
				Type:        schema.TypeList,
				Description: "Protocol and port information in an Internet Service entry.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category_id": {
							Type:        schema.TypeInt,
							Description: "Internet Service category ID.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service name.",
							Computed:    true,
						},
						"port_range": {
							Type:        schema.TypeList,
							Description: "Port ranges in the definition entry.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_port": {
										Type:        schema.TypeInt,
										Description: "Ending TCP/UDP/SCTP destination port (1 to 65535).",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Custom entry port range ID.",
										Computed:    true,
									},
									"start_port": {
										Type:        schema.TypeInt,
										Description: "Starting TCP/UDP/SCTP destination port (1 to 65535).",
										Computed:    true,
									},
								},
							},
						},
						"protocol": {
							Type:        schema.TypeInt,
							Description: "Integer value for the protocol type as defined by IANA (0 - 255).",
							Computed:    true,
						},
						"seq_num": {
							Type:        schema.TypeInt,
							Description: "Entry sequence number.",
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Internet Service application list ID.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallInternetServiceDefinitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("fosid")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadFirewallInternetServiceDefinition(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallInternetServiceDefinition dataSource: %v", err)
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

	diags := refreshObjectFirewallInternetServiceDefinition(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
