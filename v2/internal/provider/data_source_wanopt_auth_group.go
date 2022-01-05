// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WAN optimization authentication groups.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWanoptAuthGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WAN optimization authentication groups.",

		ReadContext: dataSourceWanoptAuthGroupRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auth_method": {
				Type:        schema.TypeString,
				Description: "Select certificate or pre-shared key authentication for this authentication group.",
				Computed:    true,
			},
			"cert": {
				Type:        schema.TypeString,
				Description: "Name of certificate to identify this peer.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Auth-group name.",
				Required:    true,
			},
			"peer": {
				Type:        schema.TypeString,
				Description: "If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.",
				Computed:    true,
			},
			"peer_accept": {
				Type:        schema.TypeString,
				Description: "Determine if this auth group accepts, any peer, a list of defined peers, or just one peer.",
				Computed:    true,
			},
			"psk": {
				Type:        schema.TypeString,
				Description: "Pre-shared key used by the peers in this authentication group.",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func dataSourceWanoptAuthGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWanoptAuthGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WanoptAuthGroup dataSource: %v", err)
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

	diags := refreshObjectWanoptAuthGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
