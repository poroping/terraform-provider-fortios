// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 routing policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceRouterPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 routing policies.",

		ReadContext: dataSourceRouterPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"action": {
				Type:        schema.TypeString,
				Description: "Action of the policy route.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"dst": {
				Type:        schema.TypeList,
				Description: "Destination IP and mask (x.x.x.x/x).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": {
							Type:        schema.TypeString,
							Description: "IP and mask.",
							Computed:    true,
						},
					},
				},
			},
			"dst_negate": {
				Type:        schema.TypeString,
				Description: "Enable/disable negating destination address match.",
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination address name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address/group name.",
							Computed:    true,
						},
					},
				},
			},
			"end_port": {
				Type:        schema.TypeInt,
				Description: "End destination port number (0 - 65535).",
				Computed:    true,
			},
			"end_source_port": {
				Type:        schema.TypeInt,
				Description: "End source port number (0 - 65535).",
				Computed:    true,
			},
			"gateway": {
				Type:        schema.TypeString,
				Description: "IP address of the gateway.",
				Computed:    true,
			},
			"input_device": {
				Type:        schema.TypeList,
				Description: "Incoming interface name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"input_device_negate": {
				Type:        schema.TypeString,
				Description: "Enable/disable negation of input device match.",
				Computed:    true,
			},
			"internet_service_custom": {
				Type:        schema.TypeList,
				Description: "Custom Destination Internet Service name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Destination Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_id": {
				Type:        schema.TypeList,
				Description: "Destination Internet Service ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Destination Internet Service ID.",
							Computed:    true,
						},
					},
				},
			},
			"output_device": {
				Type:        schema.TypeString,
				Description: "Outgoing interface name.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeInt,
				Description: "Protocol number (0 - 255).",
				Computed:    true,
			},
			"seq_num": {
				Type:        schema.TypeInt,
				Description: "Sequence number(1-65535).",
				Computed:    true,
			},
			"src": {
				Type:        schema.TypeList,
				Description: "Source IP and mask (x.x.x.x/x).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": {
							Type:        schema.TypeString,
							Description: "IP and mask.",
							Computed:    true,
						},
					},
				},
			},
			"src_negate": {
				Type:        schema.TypeString,
				Description: "Enable/disable negating source address match.",
				Computed:    true,
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address/group name.",
							Computed:    true,
						},
					},
				},
			},
			"start_port": {
				Type:        schema.TypeInt,
				Description: "Start destination port number (0 - 65535).",
				Computed:    true,
			},
			"start_source_port": {
				Type:        schema.TypeInt,
				Description: "Start source port number (0 - 65535).",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this policy route.",
				Computed:    true,
			},
			"tos": {
				Type:        schema.TypeString,
				Description: "Type of service bit pattern.",
				Computed:    true,
			},
			"tos_mask": {
				Type:        schema.TypeString,
				Description: "Type of service evaluated bits.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterPolicy dataSource: %v", err)
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

	diags := refreshObjectRouterPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
