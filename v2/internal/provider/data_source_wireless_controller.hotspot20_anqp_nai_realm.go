// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure network access identifier (NAI) realm.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWirelessControllerHotspot20AnqpNaiRealm() *schema.Resource {
	return &schema.Resource{
		Description: "Configure network access identifier (NAI) realm.",

		ReadContext: dataSourceWirelessControllerHotspot20AnqpNaiRealmRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"nai_list": {
				Type:        schema.TypeList,
				Description: "NAI list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"eap_method": {
							Type:        schema.TypeList,
							Description: "EAP Methods.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_param": {
										Type:        schema.TypeList,
										Description: "EAP auth param.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeString,
													Description: "ID of authentication parameter.",
													Computed:    true,
												},
												"index": {
													Type:        schema.TypeInt,
													Description: "Param index.",
													Computed:    true,
												},
												"val": {
													Type:        schema.TypeString,
													Description: "Value of authentication parameter.",
													Computed:    true,
												},
											},
										},
									},
									"index": {
										Type:        schema.TypeInt,
										Description: "EAP method index.",
										Computed:    true,
									},
									"method": {
										Type:        schema.TypeString,
										Description: "EAP method type.",
										Computed:    true,
									},
								},
							},
						},
						"encoding": {
							Type:        schema.TypeString,
							Description: "Enable/disable format in accordance with IETF RFC 4282.",
							Computed:    true,
						},
						"nai_realm": {
							Type:        schema.TypeString,
							Description: "Configure NAI realms (delimited by a semi-colon character).",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "NAI realm name.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "NAI realm list name.",
				Required:    true,
			},
		},
	}
}

func dataSourceWirelessControllerHotspot20AnqpNaiRealmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20AnqpNaiRealm(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20AnqpNaiRealm dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20AnqpNaiRealm(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
