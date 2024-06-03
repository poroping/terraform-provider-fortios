// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: OSPF6 area configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceRouterOspf6Area() *schema.Resource {
	return &schema.Resource{
		Description: "OSPF6 area configuration.",

		ReadContext: dataSourceRouterOspf6AreaRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"authentication": {
				Type:        schema.TypeString,
				Description: "Authentication mode.",
				Computed:    true,
			},
			"default_cost": {
				Type:        schema.TypeInt,
				Description: "Summary default cost of stub or NSSA area.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeString,
				Description: "Area entry IP address.",
				Computed:    true,
			},
			"ipsec_auth_alg": {
				Type:        schema.TypeString,
				Description: "Authentication algorithm.",
				Computed:    true,
			},
			"ipsec_enc_alg": {
				Type:        schema.TypeString,
				Description: "Encryption algorithm.",
				Computed:    true,
			},
			"ipsec_keys": {
				Type:        schema.TypeList,
				Description: "IPsec authentication and encryption keys.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_key": {
							Type:        schema.TypeString,
							Description: "Authentication key.",
							Computed:    true,
							Sensitive:   true,
						},
						"enc_key": {
							Type:        schema.TypeString,
							Description: "Encryption key.",
							Computed:    true,
							Sensitive:   true,
						},
						"spi": {
							Type:        schema.TypeInt,
							Description: "Security Parameters Index.",
							Computed:    true,
						},
					},
				},
			},
			"key_rollover_interval": {
				Type:        schema.TypeInt,
				Description: "Key roll-over interval.",
				Computed:    true,
			},
			"nssa_default_information_originate": {
				Type:        schema.TypeString,
				Description: "Enable/disable originate type 7 default into NSSA area.",
				Computed:    true,
			},
			"nssa_default_information_originate_metric": {
				Type:        schema.TypeInt,
				Description: "OSPFv3 default metric.",
				Computed:    true,
			},
			"nssa_default_information_originate_metric_type": {
				Type:        schema.TypeString,
				Description: "OSPFv3 metric type for default routes.",
				Computed:    true,
			},
			"nssa_redistribution": {
				Type:        schema.TypeString,
				Description: "Enable/disable redistribute into NSSA area.",
				Computed:    true,
			},
			"nssa_translator_role": {
				Type:        schema.TypeString,
				Description: "NSSA translator role type.",
				Computed:    true,
			},
			"range": {
				Type:        schema.TypeList,
				Description: "OSPF6 area range configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise status.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Range entry ID.",
							Computed:    true,
						},
						"prefix6": {
							Type:        schema.TypeString,
							Description: "IPv6 prefix.",
							Computed:    true,
						},
					},
				},
			},
			"stub_type": {
				Type:        schema.TypeString,
				Description: "Stub summary setting.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Area type setting.",
				Computed:    true,
			},
			"virtual_link": {
				Type:        schema.TypeList,
				Description: "OSPF6 virtual link configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeString,
							Description: "Authentication mode.",
							Computed:    true,
						},
						"dead_interval": {
							Type:        schema.TypeInt,
							Description: "Dead interval.",
							Computed:    true,
						},
						"hello_interval": {
							Type:        schema.TypeInt,
							Description: "Hello interval.",
							Computed:    true,
						},
						"ipsec_auth_alg": {
							Type:        schema.TypeString,
							Description: "Authentication algorithm.",
							Computed:    true,
						},
						"ipsec_enc_alg": {
							Type:        schema.TypeString,
							Description: "Encryption algorithm.",
							Computed:    true,
						},
						"ipsec_keys": {
							Type:        schema.TypeList,
							Description: "IPsec authentication and encryption keys.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_key": {
										Type:        schema.TypeString,
										Description: "Authentication key.",
										Computed:    true,
										Sensitive:   true,
									},
									"enc_key": {
										Type:        schema.TypeString,
										Description: "Encryption key.",
										Computed:    true,
										Sensitive:   true,
									},
									"spi": {
										Type:        schema.TypeInt,
										Description: "Security Parameters Index.",
										Computed:    true,
									},
								},
							},
						},
						"key_rollover_interval": {
							Type:        schema.TypeInt,
							Description: "Key roll-over interval.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Virtual link entry name.",
							Computed:    true,
						},
						"peer": {
							Type:        schema.TypeString,
							Description: "A.B.C.D, peer router ID.",
							Computed:    true,
						},
						"retransmit_interval": {
							Type:        schema.TypeInt,
							Description: "Retransmit interval.",
							Computed:    true,
						},
						"transmit_delay": {
							Type:        schema.TypeInt,
							Description: "Transmit delay.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterOspf6AreaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterOspf6Area(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspf6Area dataSource: %v", err)
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

	diags := refreshObjectRouterOspf6Area(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
