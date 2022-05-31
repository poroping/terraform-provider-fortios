// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SSH proxy host public keys.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallSshHostKey() *schema.Resource {
	return &schema.Resource{
		Description: "SSH proxy host public keys.",

		ReadContext: dataSourceFirewallSshHostKeyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"hostname": {
				Type:        schema.TypeString,
				Description: "Hostname of the SSH server to match SSH certificate principals.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "IP address of the SSH server.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "SSH public key name.",
				Required:    true,
			},
			"nid": {
				Type:        schema.TypeString,
				Description: "Set the nid of the ECDSA key.",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Port of the SSH server.",
				Computed:    true,
			},
			"public_key": {
				Type:        schema.TypeString,
				Description: "SSH public key.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Set the trust status of the public key.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Set the type of the public key.",
				Computed:    true,
			},
			"usage": {
				Type:        schema.TypeString,
				Description: "Usage for this public key.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallSshHostKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallSshHostKey(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSshHostKey dataSource: %v", err)
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

	diags := refreshObjectFirewallSshHostKey(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
