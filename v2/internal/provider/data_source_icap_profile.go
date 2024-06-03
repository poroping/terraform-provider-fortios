// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure ICAP profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceIcapProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure ICAP profiles.",

		ReadContext: dataSourceIcapProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"204_response": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowance of 204 response from ICAP server.",
				Computed:    true,
			},
			"204_size_limit": {
				Type:        schema.TypeInt,
				Description: "204 response size limit to be saved by ICAP client in megabytes (1 - 10, default = 1 MB).",
				Computed:    true,
			},
			"chunk_encap": {
				Type:        schema.TypeString,
				Description: "Enable/disable chunked encapsulation (default = disable).",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"extension_feature": {
				Type:        schema.TypeString,
				Description: "Enable/disable ICAP extension features.",
				Computed:    true,
			},
			"file_transfer": {
				Type:        schema.TypeString,
				Description: "Configure the file transfer protocols to pass transferred files to an ICAP server as REQMOD.",
				Computed:    true,
			},
			"file_transfer_failure": {
				Type:        schema.TypeString,
				Description: "Action to take if the ICAP server cannot be contacted when processing a file transfer.",
				Computed:    true,
			},
			"file_transfer_path": {
				Type:        schema.TypeString,
				Description: "Path component of the ICAP URI that identifies the file transfer processing service.",
				Computed:    true,
			},
			"file_transfer_server": {
				Type:        schema.TypeString,
				Description: "ICAP server to use for a file transfer.",
				Computed:    true,
			},
			"icap_block_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable UTM log when infection found (default = disable).",
				Computed:    true,
			},
			"icap_headers": {
				Type:        schema.TypeList,
				Description: "Configure ICAP forwarded request headers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base64_encoding": {
							Type:        schema.TypeString,
							Description: "Enable/disable use of base64 encoding of HTTP content.",
							Computed:    true,
						},
						"content": {
							Type:        schema.TypeString,
							Description: "HTTP header content.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "HTTP forwarded header ID.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "HTTP forwarded header name.",
							Computed:    true,
						},
					},
				},
			},
			"methods": {
				Type:        schema.TypeString,
				Description: "The allowed HTTP methods that will be sent to ICAP server for further processing.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "ICAP profile name.",
				Required:    true,
			},
			"preview": {
				Type:        schema.TypeString,
				Description: "Enable/disable preview of data to ICAP server.",
				Computed:    true,
			},
			"preview_data_length": {
				Type:        schema.TypeInt,
				Description: "Preview data length to be sent to ICAP server.",
				Computed:    true,
			},
			"replacemsg_group": {
				Type:        schema.TypeString,
				Description: "Replacement message group.",
				Computed:    true,
			},
			"request": {
				Type:        schema.TypeString,
				Description: "Enable/disable whether an HTTP request is passed to an ICAP server.",
				Computed:    true,
			},
			"request_failure": {
				Type:        schema.TypeString,
				Description: "Action to take if the ICAP server cannot be contacted when processing an HTTP request.",
				Computed:    true,
			},
			"request_path": {
				Type:        schema.TypeString,
				Description: "Path component of the ICAP URI that identifies the HTTP request processing service.",
				Computed:    true,
			},
			"request_server": {
				Type:        schema.TypeString,
				Description: "ICAP server to use for an HTTP request.",
				Computed:    true,
			},
			"respmod_default_action": {
				Type:        schema.TypeString,
				Description: "Default action to ICAP response modification (respmod) processing.",
				Computed:    true,
			},
			"respmod_forward_rules": {
				Type:        schema.TypeList,
				Description: "ICAP response mode forward rules.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action to be taken for ICAP server.",
							Computed:    true,
						},
						"header_group": {
							Type:        schema.TypeList,
							Description: "HTTP header group.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"case_sensitivity": {
										Type:        schema.TypeString,
										Description: "Enable/disable case sensitivity when matching header.",
										Computed:    true,
									},
									"header": {
										Type:        schema.TypeString,
										Description: "HTTP header regular expression.",
										Computed:    true,
									},
									"header_name": {
										Type:        schema.TypeString,
										Description: "HTTP header.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "ID.",
										Computed:    true,
									},
								},
							},
						},
						"host": {
							Type:        schema.TypeString,
							Description: "Address object for the host.",
							Computed:    true,
						},
						"http_resp_status_code": {
							Type:        schema.TypeList,
							Description: "HTTP response status code.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": {
										Type:        schema.TypeInt,
										Description: "HTTP response status code.",
										Computed:    true,
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"response": {
				Type:        schema.TypeString,
				Description: "Enable/disable whether an HTTP response is passed to an ICAP server.",
				Computed:    true,
			},
			"response_failure": {
				Type:        schema.TypeString,
				Description: "Action to take if the ICAP server cannot be contacted when processing an HTTP response.",
				Computed:    true,
			},
			"response_path": {
				Type:        schema.TypeString,
				Description: "Path component of the ICAP URI that identifies the HTTP response processing service.",
				Computed:    true,
			},
			"response_req_hdr": {
				Type:        schema.TypeString,
				Description: "Enable/disable addition of req-hdr for ICAP response modification (respmod) processing.",
				Computed:    true,
			},
			"response_server": {
				Type:        schema.TypeString,
				Description: "ICAP server to use for an HTTP response.",
				Computed:    true,
			},
			"scan_progress_interval": {
				Type:        schema.TypeInt,
				Description: "Scan progress interval value.",
				Computed:    true,
			},
			"streaming_content_bypass": {
				Type:        schema.TypeString,
				Description: "Enable/disable bypassing of ICAP server for streaming content.",
				Computed:    true,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "Time (in seconds) that ICAP client waits for the response from ICAP server.",
				Computed:    true,
			},
		},
	}
}

func dataSourceIcapProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadIcapProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IcapProfile dataSource: %v", err)
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

	diags := refreshObjectIcapProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
