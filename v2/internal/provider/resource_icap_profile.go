// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure ICAP profiles.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceIcapProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure ICAP profiles.",

		CreateContext: resourceIcapProfileCreate,
		ReadContext:   resourceIcapProfileRead,
		UpdateContext: resourceIcapProfileUpdate,
		DeleteContext: resourceIcapProfileDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"204_response": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable allowance of 204 response from ICAP server.",
				Optional:    true,
				Computed:    true,
			},
			"204_size_limit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),

				Description: "204 response size limit to be saved by ICAP client in megabytes (1 - 10, default = 1 MB).",
				Optional:    true,
				Computed:    true,
			},
			"chunk_encap": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable chunked encapsulation (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"extension_feature": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Enable/disable ICAP extension features.",
				Optional:         true,
				Computed:         true,
			},
			"file_transfer": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Configure the file transfer protocols to pass transferred files to an ICAP server as REQMOD.",
				Optional:         true,
				Computed:         true,
			},
			"file_transfer_failure": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"error", "bypass"}, false),

				Description: "Action to take if the ICAP server cannot be contacted when processing a file transfer.",
				Optional:    true,
				Computed:    true,
			},
			"file_transfer_path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Path component of the ICAP URI that identifies the file transfer processing service.",
				Optional:    true,
				Computed:    true,
			},
			"file_transfer_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "ICAP server to use for a file transfer.",
				Optional:    true,
				Computed:    true,
			},
			"icap_block_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable UTM log when infection found (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"icap_headers": {
				Type:        schema.TypeList,
				Description: "Configure ICAP forwarded request headers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base64_encoding": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable use of base64 encoding of HTTP content.",
							Optional:    true,
							Computed:    true,
						},
						"content": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "HTTP header content.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "HTTP forwarded header ID.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "HTTP forwarded header name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"methods": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "The allowed HTTP methods that will be sent to ICAP server for further processing.",
				Optional:         true,
				Computed:         true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "ICAP profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"preview": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable preview of data to ICAP server.",
				Optional:    true,
				Computed:    true,
			},
			"preview_data_length": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4096),

				Description: "Preview data length to be sent to ICAP server.",
				Optional:    true,
				Computed:    true,
			},
			"replacemsg_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Replacement message group.",
				Optional:    true,
				Computed:    true,
			},
			"request": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable whether an HTTP request is passed to an ICAP server.",
				Optional:    true,
				Computed:    true,
			},
			"request_failure": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"error", "bypass"}, false),

				Description: "Action to take if the ICAP server cannot be contacted when processing an HTTP request.",
				Optional:    true,
				Computed:    true,
			},
			"request_path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Path component of the ICAP URI that identifies the HTTP request processing service.",
				Optional:    true,
				Computed:    true,
			},
			"request_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "ICAP server to use for an HTTP request.",
				Optional:    true,
				Computed:    true,
			},
			"respmod_default_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"forward", "bypass"}, false),

				Description: "Default action to ICAP response modification (respmod) processing.",
				Optional:    true,
				Computed:    true,
			},
			"respmod_forward_rules": {
				Type:        schema.TypeList,
				Description: "ICAP response mode forward rules.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"forward", "bypass"}, false),

							Description: "Action to be taken for ICAP server.",
							Optional:    true,
							Computed:    true,
						},
						"header_group": {
							Type:        schema.TypeList,
							Description: "HTTP header group.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"case_sensitivity": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "Enable/disable case sensitivity when matching header.",
										Optional:    true,
										Computed:    true,
									},
									"header": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),

										Description: "HTTP header regular expression.",
										Optional:    true,
										Computed:    true,
									},
									"header_name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "HTTP header.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "ID.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"host": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address object for the host.",
							Optional:    true,
							Computed:    true,
						},
						"http_resp_status_code": {
							Type:        schema.TypeList,
							Description: "HTTP response status code.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(100, 599),

										Description: "HTTP response status code.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"response": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable whether an HTTP response is passed to an ICAP server.",
				Optional:    true,
				Computed:    true,
			},
			"response_failure": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"error", "bypass"}, false),

				Description: "Action to take if the ICAP server cannot be contacted when processing an HTTP response.",
				Optional:    true,
				Computed:    true,
			},
			"response_path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Path component of the ICAP URI that identifies the HTTP response processing service.",
				Optional:    true,
				Computed:    true,
			},
			"response_req_hdr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable addition of req-hdr for ICAP response modification (respmod) processing.",
				Optional:    true,
				Computed:    true,
			},
			"response_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "ICAP server to use for an HTTP response.",
				Optional:    true,
				Computed:    true,
			},
			"scan_progress_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 30),

				Description: "Scan progress interval value.",
				Optional:    true,
				Computed:    true,
			},
			"streaming_content_bypass": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable bypassing of ICAP server for streaming content.",
				Optional:    true,
				Computed:    true,
			},
			"timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 3600),

				Description: "Time (in seconds) that ICAP client waits for the response from ICAP server.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceIcapProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating IcapProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectIcapProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateIcapProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IcapProfile")
	}

	return resourceIcapProfileRead(ctx, d, meta)
}

func resourceIcapProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectIcapProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateIcapProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating IcapProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IcapProfile")
	}

	return resourceIcapProfileRead(ctx, d, meta)
}

func resourceIcapProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteIcapProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting IcapProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIcapProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadIcapProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IcapProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
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
	return nil
}

func flattenIcapProfileIcapHeaders(d *schema.ResourceData, v *[]models.IcapProfileIcapHeaders, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Base64Encoding; tmp != nil {
				v["base64_encoding"] = *tmp
			}

			if tmp := cfg.Content; tmp != nil {
				v["content"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenIcapProfileRespmodForwardRules(d *schema.ResourceData, v *[]models.IcapProfileRespmodForwardRules, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.HeaderGroup; tmp != nil {
				v["header_group"] = flattenIcapProfileRespmodForwardRulesHeaderGroup(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "header_group"), sort)
			}

			if tmp := cfg.Host; tmp != nil {
				v["host"] = *tmp
			}

			if tmp := cfg.HttpRespStatusCode; tmp != nil {
				v["http_resp_status_code"] = flattenIcapProfileRespmodForwardRulesHttpRespStatusCode(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "http_resp_status_code"), sort)
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenIcapProfileRespmodForwardRulesHeaderGroup(d *schema.ResourceData, v *[]models.IcapProfileRespmodForwardRulesHeaderGroup, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.CaseSensitivity; tmp != nil {
				v["case_sensitivity"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.HeaderName; tmp != nil {
				v["header_name"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenIcapProfileRespmodForwardRulesHttpRespStatusCode(d *schema.ResourceData, v *[]models.IcapProfileRespmodForwardRulesHttpRespStatusCode, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Code; tmp != nil {
				v["code"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "code")
	}

	return flat
}

func refreshObjectIcapProfile(d *schema.ResourceData, o *models.IcapProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.The204Response != nil {
		v := *o.The204Response

		if err = d.Set("204_response", v); err != nil {
			return diag.Errorf("error reading 204_response: %v", err)
		}
	}

	if o.The204SizeLimit != nil {
		v := *o.The204SizeLimit

		if err = d.Set("204_size_limit", v); err != nil {
			return diag.Errorf("error reading 204_size_limit: %v", err)
		}
	}

	if o.ChunkEncap != nil {
		v := *o.ChunkEncap

		if err = d.Set("chunk_encap", v); err != nil {
			return diag.Errorf("error reading chunk_encap: %v", err)
		}
	}

	if o.ExtensionFeature != nil {
		v := *o.ExtensionFeature

		if err = d.Set("extension_feature", v); err != nil {
			return diag.Errorf("error reading extension_feature: %v", err)
		}
	}

	if o.FileTransfer != nil {
		v := *o.FileTransfer

		if err = d.Set("file_transfer", v); err != nil {
			return diag.Errorf("error reading file_transfer: %v", err)
		}
	}

	if o.FileTransferFailure != nil {
		v := *o.FileTransferFailure

		if err = d.Set("file_transfer_failure", v); err != nil {
			return diag.Errorf("error reading file_transfer_failure: %v", err)
		}
	}

	if o.FileTransferPath != nil {
		v := *o.FileTransferPath

		if err = d.Set("file_transfer_path", v); err != nil {
			return diag.Errorf("error reading file_transfer_path: %v", err)
		}
	}

	if o.FileTransferServer != nil {
		v := *o.FileTransferServer

		if err = d.Set("file_transfer_server", v); err != nil {
			return diag.Errorf("error reading file_transfer_server: %v", err)
		}
	}

	if o.IcapBlockLog != nil {
		v := *o.IcapBlockLog

		if err = d.Set("icap_block_log", v); err != nil {
			return diag.Errorf("error reading icap_block_log: %v", err)
		}
	}

	if o.IcapHeaders != nil {
		if err = d.Set("icap_headers", flattenIcapProfileIcapHeaders(d, o.IcapHeaders, "icap_headers", sort)); err != nil {
			return diag.Errorf("error reading icap_headers: %v", err)
		}
	}

	if o.Methods != nil {
		v := *o.Methods

		if err = d.Set("methods", v); err != nil {
			return diag.Errorf("error reading methods: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Preview != nil {
		v := *o.Preview

		if err = d.Set("preview", v); err != nil {
			return diag.Errorf("error reading preview: %v", err)
		}
	}

	if o.PreviewDataLength != nil {
		v := *o.PreviewDataLength

		if err = d.Set("preview_data_length", v); err != nil {
			return diag.Errorf("error reading preview_data_length: %v", err)
		}
	}

	if o.ReplacemsgGroup != nil {
		v := *o.ReplacemsgGroup

		if err = d.Set("replacemsg_group", v); err != nil {
			return diag.Errorf("error reading replacemsg_group: %v", err)
		}
	}

	if o.Request != nil {
		v := *o.Request

		if err = d.Set("request", v); err != nil {
			return diag.Errorf("error reading request: %v", err)
		}
	}

	if o.RequestFailure != nil {
		v := *o.RequestFailure

		if err = d.Set("request_failure", v); err != nil {
			return diag.Errorf("error reading request_failure: %v", err)
		}
	}

	if o.RequestPath != nil {
		v := *o.RequestPath

		if err = d.Set("request_path", v); err != nil {
			return diag.Errorf("error reading request_path: %v", err)
		}
	}

	if o.RequestServer != nil {
		v := *o.RequestServer

		if err = d.Set("request_server", v); err != nil {
			return diag.Errorf("error reading request_server: %v", err)
		}
	}

	if o.RespmodDefaultAction != nil {
		v := *o.RespmodDefaultAction

		if err = d.Set("respmod_default_action", v); err != nil {
			return diag.Errorf("error reading respmod_default_action: %v", err)
		}
	}

	if o.RespmodForwardRules != nil {
		if err = d.Set("respmod_forward_rules", flattenIcapProfileRespmodForwardRules(d, o.RespmodForwardRules, "respmod_forward_rules", sort)); err != nil {
			return diag.Errorf("error reading respmod_forward_rules: %v", err)
		}
	}

	if o.Response != nil {
		v := *o.Response

		if err = d.Set("response", v); err != nil {
			return diag.Errorf("error reading response: %v", err)
		}
	}

	if o.ResponseFailure != nil {
		v := *o.ResponseFailure

		if err = d.Set("response_failure", v); err != nil {
			return diag.Errorf("error reading response_failure: %v", err)
		}
	}

	if o.ResponsePath != nil {
		v := *o.ResponsePath

		if err = d.Set("response_path", v); err != nil {
			return diag.Errorf("error reading response_path: %v", err)
		}
	}

	if o.ResponseReqHdr != nil {
		v := *o.ResponseReqHdr

		if err = d.Set("response_req_hdr", v); err != nil {
			return diag.Errorf("error reading response_req_hdr: %v", err)
		}
	}

	if o.ResponseServer != nil {
		v := *o.ResponseServer

		if err = d.Set("response_server", v); err != nil {
			return diag.Errorf("error reading response_server: %v", err)
		}
	}

	if o.ScanProgressInterval != nil {
		v := *o.ScanProgressInterval

		if err = d.Set("scan_progress_interval", v); err != nil {
			return diag.Errorf("error reading scan_progress_interval: %v", err)
		}
	}

	if o.StreamingContentBypass != nil {
		v := *o.StreamingContentBypass

		if err = d.Set("streaming_content_bypass", v); err != nil {
			return diag.Errorf("error reading streaming_content_bypass: %v", err)
		}
	}

	if o.Timeout != nil {
		v := *o.Timeout

		if err = d.Set("timeout", v); err != nil {
			return diag.Errorf("error reading timeout: %v", err)
		}
	}

	return nil
}

func expandIcapProfileIcapHeaders(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.IcapProfileIcapHeaders, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IcapProfileIcapHeaders

	for i := range l {
		tmp := models.IcapProfileIcapHeaders{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.base64_encoding", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Base64Encoding = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.content", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Content = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandIcapProfileRespmodForwardRules(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.IcapProfileRespmodForwardRules, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IcapProfileRespmodForwardRules

	for i := range l {
		tmp := models.IcapProfileRespmodForwardRules{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandIcapProfileRespmodForwardRulesHeaderGroup(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.IcapProfileRespmodForwardRulesHeaderGroup
			// 	}
			tmp.HeaderGroup = v2

		}

		pre_append = fmt.Sprintf("%s.%d.host", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Host = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_resp_status_code", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandIcapProfileRespmodForwardRulesHttpRespStatusCode(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.IcapProfileRespmodForwardRulesHttpRespStatusCode
			// 	}
			tmp.HttpRespStatusCode = v2

		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandIcapProfileRespmodForwardRulesHeaderGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.IcapProfileRespmodForwardRulesHeaderGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IcapProfileRespmodForwardRulesHeaderGroup

	for i := range l {
		tmp := models.IcapProfileRespmodForwardRulesHeaderGroup{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.case_sensitivity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CaseSensitivity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HeaderName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandIcapProfileRespmodForwardRulesHttpRespStatusCode(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.IcapProfileRespmodForwardRulesHttpRespStatusCode, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IcapProfileRespmodForwardRulesHttpRespStatusCode

	for i := range l {
		tmp := models.IcapProfileRespmodForwardRulesHttpRespStatusCode{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.code", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Code = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectIcapProfile(d *schema.ResourceData, sv string) (*models.IcapProfile, diag.Diagnostics) {
	obj := models.IcapProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("204_response"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("204_response", sv)
				diags = append(diags, e)
			}
			obj.The204Response = &v2
		}
	}
	if v1, ok := d.GetOk("204_size_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("204_size_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.The204SizeLimit = &tmp
		}
	}
	if v1, ok := d.GetOk("chunk_encap"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("chunk_encap", sv)
				diags = append(diags, e)
			}
			obj.ChunkEncap = &v2
		}
	}
	if v1, ok := d.GetOk("extension_feature"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("extension_feature", sv)
				diags = append(diags, e)
			}
			obj.ExtensionFeature = &v2
		}
	}
	if v1, ok := d.GetOk("file_transfer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("file_transfer", sv)
				diags = append(diags, e)
			}
			obj.FileTransfer = &v2
		}
	}
	if v1, ok := d.GetOk("file_transfer_failure"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("file_transfer_failure", sv)
				diags = append(diags, e)
			}
			obj.FileTransferFailure = &v2
		}
	}
	if v1, ok := d.GetOk("file_transfer_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("file_transfer_path", sv)
				diags = append(diags, e)
			}
			obj.FileTransferPath = &v2
		}
	}
	if v1, ok := d.GetOk("file_transfer_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("file_transfer_server", sv)
				diags = append(diags, e)
			}
			obj.FileTransferServer = &v2
		}
	}
	if v1, ok := d.GetOk("icap_block_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("icap_block_log", sv)
				diags = append(diags, e)
			}
			obj.IcapBlockLog = &v2
		}
	}
	if v, ok := d.GetOk("icap_headers"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("icap_headers", sv)
			diags = append(diags, e)
		}
		t, err := expandIcapProfileIcapHeaders(d, v, "icap_headers", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IcapHeaders = t
		}
	} else if d.HasChange("icap_headers") {
		old, new := d.GetChange("icap_headers")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IcapHeaders = &[]models.IcapProfileIcapHeaders{}
		}
	}
	if v1, ok := d.GetOk("methods"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("methods", sv)
				diags = append(diags, e)
			}
			obj.Methods = &v2
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("preview"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("preview", sv)
				diags = append(diags, e)
			}
			obj.Preview = &v2
		}
	}
	if v1, ok := d.GetOk("preview_data_length"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("preview_data_length", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PreviewDataLength = &tmp
		}
	}
	if v1, ok := d.GetOk("replacemsg_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("replacemsg_group", sv)
				diags = append(diags, e)
			}
			obj.ReplacemsgGroup = &v2
		}
	}
	if v1, ok := d.GetOk("request"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("request", sv)
				diags = append(diags, e)
			}
			obj.Request = &v2
		}
	}
	if v1, ok := d.GetOk("request_failure"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("request_failure", sv)
				diags = append(diags, e)
			}
			obj.RequestFailure = &v2
		}
	}
	if v1, ok := d.GetOk("request_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("request_path", sv)
				diags = append(diags, e)
			}
			obj.RequestPath = &v2
		}
	}
	if v1, ok := d.GetOk("request_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("request_server", sv)
				diags = append(diags, e)
			}
			obj.RequestServer = &v2
		}
	}
	if v1, ok := d.GetOk("respmod_default_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("respmod_default_action", sv)
				diags = append(diags, e)
			}
			obj.RespmodDefaultAction = &v2
		}
	}
	if v, ok := d.GetOk("respmod_forward_rules"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("respmod_forward_rules", sv)
			diags = append(diags, e)
		}
		t, err := expandIcapProfileRespmodForwardRules(d, v, "respmod_forward_rules", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.RespmodForwardRules = t
		}
	} else if d.HasChange("respmod_forward_rules") {
		old, new := d.GetChange("respmod_forward_rules")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.RespmodForwardRules = &[]models.IcapProfileRespmodForwardRules{}
		}
	}
	if v1, ok := d.GetOk("response"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("response", sv)
				diags = append(diags, e)
			}
			obj.Response = &v2
		}
	}
	if v1, ok := d.GetOk("response_failure"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("response_failure", sv)
				diags = append(diags, e)
			}
			obj.ResponseFailure = &v2
		}
	}
	if v1, ok := d.GetOk("response_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("response_path", sv)
				diags = append(diags, e)
			}
			obj.ResponsePath = &v2
		}
	}
	if v1, ok := d.GetOk("response_req_hdr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("response_req_hdr", sv)
				diags = append(diags, e)
			}
			obj.ResponseReqHdr = &v2
		}
	}
	if v1, ok := d.GetOk("response_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("response_server", sv)
				diags = append(diags, e)
			}
			obj.ResponseServer = &v2
		}
	}
	if v1, ok := d.GetOk("scan_progress_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("scan_progress_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ScanProgressInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("streaming_content_bypass"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("streaming_content_bypass", sv)
				diags = append(diags, e)
			}
			obj.StreamingContentBypass = &v2
		}
	}
	if v1, ok := d.GetOk("timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Timeout = &tmp
		}
	}
	return &obj, diags
}
