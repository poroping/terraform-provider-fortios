// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: OSPF area configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceRouterOspfArea() *schema.Resource {
	return &schema.Resource{
		Description: "OSPF area configuration.",

		ReadContext: dataSourceRouterOspfAreaRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"authentication": {
				Type:        schema.TypeString,
				Description: "Authentication type.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"default_cost": {
				Type:        schema.TypeInt,
				Description: "Summary default cost of stub or NSSA area.",
				Computed:    true,
			},
			"filter_list": {
				Type:        schema.TypeList,
				Description: "OSPF area filter-list configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": {
							Type:        schema.TypeString,
							Description: "Direction.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Filter list entry ID.",
							Computed:    true,
						},
						"list": {
							Type:        schema.TypeString,
							Description: "Access-list or prefix-list name.",
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:        schema.TypeString,
				Description: "Area entry IP address.",
				Computed:    true,
			},
			"nssa_default_information_originate": {
				Type:        schema.TypeString,
				Description: "Redistribute, advertise, or do not originate Type-7 default route into NSSA area.",
				Computed:    true,
			},
			"nssa_default_information_originate_metric": {
				Type:        schema.TypeInt,
				Description: "OSPF default metric.",
				Computed:    true,
			},
			"nssa_default_information_originate_metric_type": {
				Type:        schema.TypeString,
				Description: "OSPF metric type for default routes.",
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
				Description: "OSPF area range configuration.",
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
						"prefix": {
							Type:        schema.TypeString,
							Description: "Prefix.",
							Computed:    true,
						},
						"substitute": {
							Type:        schema.TypeString,
							Description: "Substitute prefix.",
							Computed:    true,
						},
						"substitute_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable substitute status.",
							Computed:    true,
						},
					},
				},
			},
			"shortcut": {
				Type:        schema.TypeString,
				Description: "Enable/disable shortcut option.",
				Computed:    true,
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
				Description: "OSPF virtual link configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeString,
							Description: "Authentication type.",
							Computed:    true,
						},
						"authentication_key": {
							Type:        schema.TypeString,
							Description: "Authentication key.",
							Computed:    true,
							Sensitive:   true,
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
						"keychain": {
							Type:        schema.TypeString,
							Description: "Message-digest key-chain name.",
							Computed:    true,
						},
						"md5_keychain": {
							Type:        schema.TypeString,
							Description: "Authentication MD5 key-chain name.",
							Computed:    true,
						},
						"md5_keys": {
							Type:        schema.TypeList,
							Description: "MD5 key.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Key ID (1 - 255).",
										Computed:    true,
									},
									"key_string": {
										Type:        schema.TypeString,
										Description: "Password for the key.",
										Computed:    true,
										Sensitive:   true,
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Virtual link entry name.",
							Computed:    true,
						},
						"peer": {
							Type:        schema.TypeString,
							Description: "Peer IP.",
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

func dataSourceRouterOspfAreaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	o, err := c.Cmdb.ReadRouterOspfArea(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspfArea dataSource: %v", err)
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

	diags := refreshObjectRouterOspfArea(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
