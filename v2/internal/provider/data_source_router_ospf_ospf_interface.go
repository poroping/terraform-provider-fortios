// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: OSPF interface configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceRouterOspfOspfInterface() *schema.Resource {
	return &schema.Resource{
		Description: "OSPF interface configuration.",

		ReadContext: dataSourceRouterOspfOspfInterfaceRead,

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
			"authentication_key": {
				Type:        schema.TypeString,
				Description: "Authentication key.",
				Computed:    true,
				Sensitive:   true,
			},
			"bfd": {
				Type:        schema.TypeString,
				Description: "Bidirectional Forwarding Detection (BFD).",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"cost": {
				Type:        schema.TypeInt,
				Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
				Computed:    true,
			},
			"database_filter_out": {
				Type:        schema.TypeString,
				Description: "Enable/disable control of flooding out LSAs.",
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
			"hello_multiplier": {
				Type:        schema.TypeInt,
				Description: "Number of hello packets within dead interval.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Configuration interface name.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "IP address.",
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
			"mtu": {
				Type:        schema.TypeInt,
				Description: "MTU for database description packets.",
				Computed:    true,
			},
			"mtu_ignore": {
				Type:        schema.TypeString,
				Description: "Enable/disable ignore MTU.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Interface entry name.",
				Required:    true,
			},
			"network_type": {
				Type:        schema.TypeString,
				Description: "Network type.",
				Computed:    true,
			},
			"prefix_length": {
				Type:        schema.TypeInt,
				Description: "Prefix length.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Priority.",
				Computed:    true,
			},
			"resync_timeout": {
				Type:        schema.TypeInt,
				Description: "Graceful restart neighbor resynchronization timeout.",
				Computed:    true,
			},
			"retransmit_interval": {
				Type:        schema.TypeInt,
				Description: "Retransmit interval.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable status.",
				Computed:    true,
			},
			"transmit_delay": {
				Type:        schema.TypeInt,
				Description: "Transmit delay.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterOspfOspfInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterOspfOspfInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspfOspfInterface dataSource: %v", err)
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

	diags := refreshObjectRouterOspfOspfInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
