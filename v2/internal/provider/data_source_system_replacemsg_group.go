// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure replacement message groups.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemReplacemsgGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure replacement message groups.",

		ReadContext: dataSourceSystemReplacemsgGroupRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"admin": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"alertmail": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"auth": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"automation": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"custom_message": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"device_detection_portal": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"fortiguard_wf": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"ftp": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"group_type": {
				Type:        schema.TypeString,
				Description: "Group type.",
				Computed:    true,
			},
			"http": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"icap": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"mail": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"nac_quar": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Group name.",
				Required:    true,
			},
			"nntp": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"spam": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"sslvpn": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"traffic_quota": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"utm": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
			"webproxy": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:        schema.TypeString,
							Description: "Message string.",
							Computed:    true,
						},
						"format": {
							Type:        schema.TypeString,
							Description: "Format flag.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "Header flag.",
							Computed:    true,
						},
						"msg_type": {
							Type:        schema.TypeString,
							Description: "Message type.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemReplacemsgGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemReplacemsgGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemReplacemsgGroup dataSource: %v", err)
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

	diags := refreshObjectSystemReplacemsgGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
