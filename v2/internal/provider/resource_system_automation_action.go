// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Action for automation stitches.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemAutomationAction() *schema.Resource {
	return &schema.Resource{
		Description: "Action for automation stitches.",

		CreateContext: resourceSystemAutomationActionCreate,
		ReadContext:   resourceSystemAutomationActionRead,
		UpdateContext: resourceSystemAutomationActionUpdate,
		DeleteContext: resourceSystemAutomationActionDelete,

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
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"accprofile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Access profile for CLI script action to access FortiGate features.",
				Optional:    true,
				Computed:    true,
			},
			"action_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"email", "fortiexplorer-notification", "alert", "disable-ssid", "quarantine", "quarantine-forticlient", "quarantine-nsx", "quarantine-fortinac", "ban-ip", "aws-lambda", "azure-function", "google-cloud-function", "alicloud-function", "webhook", "cli-script", "slack-notification", "microsoft-teams-notification"}, false),

				Description: "Action type.",
				Optional:    true,
				Computed:    true,
			},
			"alicloud_access_key_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "AliCloud AccessKey ID.",
				Optional:    true,
				Computed:    true,
			},
			"alicloud_access_key_secret": {
				Type: schema.TypeString,

				Description: "AliCloud AccessKey secret.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"alicloud_account_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "AliCloud account ID.",
				Optional:    true,
				Computed:    true,
			},
			"alicloud_function": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),

				Description: "AliCloud function name.",
				Optional:    true,
				Computed:    true,
			},
			"alicloud_function_authorization": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"anonymous", "function"}, false),

				Description: "AliCloud function authorization type.",
				Optional:    true,
				Computed:    true,
			},
			"alicloud_function_domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "AliCloud function domain.",
				Optional:    true,
				Computed:    true,
			},
			"alicloud_region": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "AliCloud region.",
				Optional:    true,
				Computed:    true,
			},
			"alicloud_service": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),

				Description: "AliCloud service name.",
				Optional:    true,
				Computed:    true,
			},
			"alicloud_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "AliCloud version.",
				Optional:    true,
				Computed:    true,
			},
			"aws_api_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "AWS API Gateway ID.",
				Optional:    true,
				Computed:    true,
			},
			"aws_api_key": {
				Type: schema.TypeString,

				Description: "AWS API Gateway API key.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"aws_api_path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "AWS API Gateway path.",
				Optional:    true,
				Computed:    true,
			},
			"aws_api_stage": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "AWS API Gateway deployment stage name.",
				Optional:    true,
				Computed:    true,
			},
			"aws_domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "AWS domain.",
				Optional:    true,
				Computed:    true,
			},
			"aws_region": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "AWS region.",
				Optional:    true,
				Computed:    true,
			},
			"azure_api_key": {
				Type: schema.TypeString,

				Description: "Azure function API key.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"azure_app": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Azure function application name.",
				Optional:    true,
				Computed:    true,
			},
			"azure_domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Azure function domain.",
				Optional:    true,
				Computed:    true,
			},
			"azure_function": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Azure function name.",
				Optional:    true,
				Computed:    true,
			},
			"azure_function_authorization": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"anonymous", "function", "admin"}, false),

				Description: "Azure function authorization level.",
				Optional:    true,
				Computed:    true,
			},
			"delay": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),

				Description: "Delay before execution (in seconds).",
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Description.",
				Optional:    true,
				Computed:    true,
			},
			"email_body": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Email body.",
				Optional:    true,
				Computed:    true,
			},
			"email_from": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Email sender name.",
				Optional:    true,
				Computed:    true,
			},
			"email_subject": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),

				Description: "Email subject.",
				Optional:    true,
				Computed:    true,
			},
			"email_to": {
				Type:        schema.TypeList,
				Description: "Email addresses.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Email address.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"execute_security_fabric": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric.",
				Optional:    true,
				Computed:    true,
			},
			"gcp_function": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Google Cloud function name.",
				Optional:    true,
				Computed:    true,
			},
			"gcp_function_domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Google Cloud function domain.",
				Optional:    true,
				Computed:    true,
			},
			"gcp_function_region": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Google Cloud function region.",
				Optional:    true,
				Computed:    true,
			},
			"gcp_project": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Google Cloud Platform project name.",
				Optional:    true,
				Computed:    true,
			},
			"headers": {
				Type:        schema.TypeList,
				Description: "Request headers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Request header.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"http_body": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4095),

				Description: "Request body (if necessary). Should be serialized json string.",
				Optional:    true,
				Computed:    true,
			},
			"message": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4095),

				Description: "Message content.",
				Optional:    true,
				Computed:    true,
			},
			"message_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"text", "json"}, false),

				Description: "Message type.",
				Optional:    true,
				Computed:    true,
			},
			"method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"post", "put", "get", "patch", "delete"}, false),

				Description: "Request method (POST, PUT, GET, PATCH or DELETE).",
				Optional:    true,
				Computed:    true,
			},
			"minimum_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2592000),

				Description: "Limit execution to no more than once in this interval (in seconds).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Protocol port.",
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"http", "https"}, false),

				Description: "Request protocol.",
				Optional:    true,
				Computed:    true,
			},
			"replacement_message": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable replacement message.",
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
			"required": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Required in action chain.",
				Optional:    true,
				Computed:    true,
			},
			"script": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "CLI script.",
				Optional:    true,
				Computed:    true,
			},
			"sdn_connector": {
				Type:        schema.TypeList,
				Description: "NSX SDN connector names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "SDN connector name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"security_tag": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "NSX security tag.",
				Optional:    true,
				Computed:    true,
			},
			"tls_certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Custom TLS certificate for API request.",
				Optional:    true,
				Computed:    true,
			},
			"uri": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Request API URI.",
				Optional:    true,
				Computed:    true,
			},
			"verify_host_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable verification of the remote host certificate.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemAutomationActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemAutomationAction resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemAutomationAction(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemAutomationAction(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAutomationAction")
	}

	return resourceSystemAutomationActionRead(ctx, d, meta)
}

func resourceSystemAutomationActionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAutomationAction(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemAutomationAction(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemAutomationAction resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAutomationAction")
	}

	return resourceSystemAutomationActionRead(ctx, d, meta)
}

func resourceSystemAutomationActionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemAutomationAction(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemAutomationAction resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAutomationAction(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAutomationAction resource: %v", err)
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

	diags := refreshObjectSystemAutomationAction(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemAutomationActionEmailTo(v *[]models.SystemAutomationActionEmailTo, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
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

func flattenSystemAutomationActionHeaders(v *[]models.SystemAutomationActionHeaders, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "header")
	}

	return flat
}

func flattenSystemAutomationActionSdnConnector(v *[]models.SystemAutomationActionSdnConnector, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
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

func refreshObjectSystemAutomationAction(d *schema.ResourceData, o *models.SystemAutomationAction, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Accprofile != nil {
		v := *o.Accprofile

		if err = d.Set("accprofile", v); err != nil {
			return diag.Errorf("error reading accprofile: %v", err)
		}
	}

	if o.ActionType != nil {
		v := *o.ActionType

		if err = d.Set("action_type", v); err != nil {
			return diag.Errorf("error reading action_type: %v", err)
		}
	}

	if o.AlicloudAccessKeyId != nil {
		v := *o.AlicloudAccessKeyId

		if err = d.Set("alicloud_access_key_id", v); err != nil {
			return diag.Errorf("error reading alicloud_access_key_id: %v", err)
		}
	}

	if o.AlicloudAccessKeySecret != nil {
		v := *o.AlicloudAccessKeySecret

		if err = d.Set("alicloud_access_key_secret", v); err != nil {
			return diag.Errorf("error reading alicloud_access_key_secret: %v", err)
		}
	}

	if o.AlicloudAccountId != nil {
		v := *o.AlicloudAccountId

		if err = d.Set("alicloud_account_id", v); err != nil {
			return diag.Errorf("error reading alicloud_account_id: %v", err)
		}
	}

	if o.AlicloudFunction != nil {
		v := *o.AlicloudFunction

		if err = d.Set("alicloud_function", v); err != nil {
			return diag.Errorf("error reading alicloud_function: %v", err)
		}
	}

	if o.AlicloudFunctionAuthorization != nil {
		v := *o.AlicloudFunctionAuthorization

		if err = d.Set("alicloud_function_authorization", v); err != nil {
			return diag.Errorf("error reading alicloud_function_authorization: %v", err)
		}
	}

	if o.AlicloudFunctionDomain != nil {
		v := *o.AlicloudFunctionDomain

		if err = d.Set("alicloud_function_domain", v); err != nil {
			return diag.Errorf("error reading alicloud_function_domain: %v", err)
		}
	}

	if o.AlicloudRegion != nil {
		v := *o.AlicloudRegion

		if err = d.Set("alicloud_region", v); err != nil {
			return diag.Errorf("error reading alicloud_region: %v", err)
		}
	}

	if o.AlicloudService != nil {
		v := *o.AlicloudService

		if err = d.Set("alicloud_service", v); err != nil {
			return diag.Errorf("error reading alicloud_service: %v", err)
		}
	}

	if o.AlicloudVersion != nil {
		v := *o.AlicloudVersion

		if err = d.Set("alicloud_version", v); err != nil {
			return diag.Errorf("error reading alicloud_version: %v", err)
		}
	}

	if o.AwsApiId != nil {
		v := *o.AwsApiId

		if err = d.Set("aws_api_id", v); err != nil {
			return diag.Errorf("error reading aws_api_id: %v", err)
		}
	}

	if o.AwsApiKey != nil {
		v := *o.AwsApiKey

		if err = d.Set("aws_api_key", v); err != nil {
			return diag.Errorf("error reading aws_api_key: %v", err)
		}
	}

	if o.AwsApiPath != nil {
		v := *o.AwsApiPath

		if err = d.Set("aws_api_path", v); err != nil {
			return diag.Errorf("error reading aws_api_path: %v", err)
		}
	}

	if o.AwsApiStage != nil {
		v := *o.AwsApiStage

		if err = d.Set("aws_api_stage", v); err != nil {
			return diag.Errorf("error reading aws_api_stage: %v", err)
		}
	}

	if o.AwsDomain != nil {
		v := *o.AwsDomain

		if err = d.Set("aws_domain", v); err != nil {
			return diag.Errorf("error reading aws_domain: %v", err)
		}
	}

	if o.AwsRegion != nil {
		v := *o.AwsRegion

		if err = d.Set("aws_region", v); err != nil {
			return diag.Errorf("error reading aws_region: %v", err)
		}
	}

	if o.AzureApiKey != nil {
		v := *o.AzureApiKey

		if err = d.Set("azure_api_key", v); err != nil {
			return diag.Errorf("error reading azure_api_key: %v", err)
		}
	}

	if o.AzureApp != nil {
		v := *o.AzureApp

		if err = d.Set("azure_app", v); err != nil {
			return diag.Errorf("error reading azure_app: %v", err)
		}
	}

	if o.AzureDomain != nil {
		v := *o.AzureDomain

		if err = d.Set("azure_domain", v); err != nil {
			return diag.Errorf("error reading azure_domain: %v", err)
		}
	}

	if o.AzureFunction != nil {
		v := *o.AzureFunction

		if err = d.Set("azure_function", v); err != nil {
			return diag.Errorf("error reading azure_function: %v", err)
		}
	}

	if o.AzureFunctionAuthorization != nil {
		v := *o.AzureFunctionAuthorization

		if err = d.Set("azure_function_authorization", v); err != nil {
			return diag.Errorf("error reading azure_function_authorization: %v", err)
		}
	}

	if o.Delay != nil {
		v := *o.Delay

		if err = d.Set("delay", v); err != nil {
			return diag.Errorf("error reading delay: %v", err)
		}
	}

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.EmailBody != nil {
		v := *o.EmailBody

		if err = d.Set("email_body", v); err != nil {
			return diag.Errorf("error reading email_body: %v", err)
		}
	}

	if o.EmailFrom != nil {
		v := *o.EmailFrom

		if err = d.Set("email_from", v); err != nil {
			return diag.Errorf("error reading email_from: %v", err)
		}
	}

	if o.EmailSubject != nil {
		v := *o.EmailSubject

		if err = d.Set("email_subject", v); err != nil {
			return diag.Errorf("error reading email_subject: %v", err)
		}
	}

	if o.EmailTo != nil {
		if err = d.Set("email_to", flattenSystemAutomationActionEmailTo(o.EmailTo, sort)); err != nil {
			return diag.Errorf("error reading email_to: %v", err)
		}
	}

	if o.ExecuteSecurityFabric != nil {
		v := *o.ExecuteSecurityFabric

		if err = d.Set("execute_security_fabric", v); err != nil {
			return diag.Errorf("error reading execute_security_fabric: %v", err)
		}
	}

	if o.GcpFunction != nil {
		v := *o.GcpFunction

		if err = d.Set("gcp_function", v); err != nil {
			return diag.Errorf("error reading gcp_function: %v", err)
		}
	}

	if o.GcpFunctionDomain != nil {
		v := *o.GcpFunctionDomain

		if err = d.Set("gcp_function_domain", v); err != nil {
			return diag.Errorf("error reading gcp_function_domain: %v", err)
		}
	}

	if o.GcpFunctionRegion != nil {
		v := *o.GcpFunctionRegion

		if err = d.Set("gcp_function_region", v); err != nil {
			return diag.Errorf("error reading gcp_function_region: %v", err)
		}
	}

	if o.GcpProject != nil {
		v := *o.GcpProject

		if err = d.Set("gcp_project", v); err != nil {
			return diag.Errorf("error reading gcp_project: %v", err)
		}
	}

	if o.Headers != nil {
		if err = d.Set("headers", flattenSystemAutomationActionHeaders(o.Headers, sort)); err != nil {
			return diag.Errorf("error reading headers: %v", err)
		}
	}

	if o.HttpBody != nil {
		v := *o.HttpBody

		if err = d.Set("http_body", v); err != nil {
			return diag.Errorf("error reading http_body: %v", err)
		}
	}

	if o.Message != nil {
		v := *o.Message

		if err = d.Set("message", v); err != nil {
			return diag.Errorf("error reading message: %v", err)
		}
	}

	if o.MessageType != nil {
		v := *o.MessageType

		if err = d.Set("message_type", v); err != nil {
			return diag.Errorf("error reading message_type: %v", err)
		}
	}

	if o.Method != nil {
		v := *o.Method

		if err = d.Set("method", v); err != nil {
			return diag.Errorf("error reading method: %v", err)
		}
	}

	if o.MinimumInterval != nil {
		v := *o.MinimumInterval

		if err = d.Set("minimum_interval", v); err != nil {
			return diag.Errorf("error reading minimum_interval: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.ReplacementMessage != nil {
		v := *o.ReplacementMessage

		if err = d.Set("replacement_message", v); err != nil {
			return diag.Errorf("error reading replacement_message: %v", err)
		}
	}

	if o.ReplacemsgGroup != nil {
		v := *o.ReplacemsgGroup

		if err = d.Set("replacemsg_group", v); err != nil {
			return diag.Errorf("error reading replacemsg_group: %v", err)
		}
	}

	if o.Required != nil {
		v := *o.Required

		if err = d.Set("required", v); err != nil {
			return diag.Errorf("error reading required: %v", err)
		}
	}

	if o.Script != nil {
		v := *o.Script

		if err = d.Set("script", v); err != nil {
			return diag.Errorf("error reading script: %v", err)
		}
	}

	if o.SdnConnector != nil {
		if err = d.Set("sdn_connector", flattenSystemAutomationActionSdnConnector(o.SdnConnector, sort)); err != nil {
			return diag.Errorf("error reading sdn_connector: %v", err)
		}
	}

	if o.SecurityTag != nil {
		v := *o.SecurityTag

		if err = d.Set("security_tag", v); err != nil {
			return diag.Errorf("error reading security_tag: %v", err)
		}
	}

	if o.TlsCertificate != nil {
		v := *o.TlsCertificate

		if err = d.Set("tls_certificate", v); err != nil {
			return diag.Errorf("error reading tls_certificate: %v", err)
		}
	}

	if o.Uri != nil {
		v := *o.Uri

		if err = d.Set("uri", v); err != nil {
			return diag.Errorf("error reading uri: %v", err)
		}
	}

	if o.VerifyHostCert != nil {
		v := *o.VerifyHostCert

		if err = d.Set("verify_host_cert", v); err != nil {
			return diag.Errorf("error reading verify_host_cert: %v", err)
		}
	}

	return nil
}

func expandSystemAutomationActionEmailTo(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAutomationActionEmailTo, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAutomationActionEmailTo

	for i := range l {
		tmp := models.SystemAutomationActionEmailTo{}
		var pre_append string

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

func expandSystemAutomationActionHeaders(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAutomationActionHeaders, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAutomationActionHeaders

	for i := range l {
		tmp := models.SystemAutomationActionHeaders{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemAutomationActionSdnConnector(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAutomationActionSdnConnector, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAutomationActionSdnConnector

	for i := range l {
		tmp := models.SystemAutomationActionSdnConnector{}
		var pre_append string

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

func getObjectSystemAutomationAction(d *schema.ResourceData, sv string) (*models.SystemAutomationAction, diag.Diagnostics) {
	obj := models.SystemAutomationAction{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("accprofile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "") {
				e := utils.AttributeVersionWarning("accprofile", sv)
				diags = append(diags, e)
			}
			obj.Accprofile = &v2
		}
	}
	if v1, ok := d.GetOk("action_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("action_type", sv)
				diags = append(diags, e)
			}
			obj.ActionType = &v2
		}
	}
	if v1, ok := d.GetOk("alicloud_access_key_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("alicloud_access_key_id", sv)
				diags = append(diags, e)
			}
			obj.AlicloudAccessKeyId = &v2
		}
	}
	if v1, ok := d.GetOk("alicloud_access_key_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("alicloud_access_key_secret", sv)
				diags = append(diags, e)
			}
			obj.AlicloudAccessKeySecret = &v2
		}
	}
	if v1, ok := d.GetOk("alicloud_account_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("alicloud_account_id", sv)
				diags = append(diags, e)
			}
			obj.AlicloudAccountId = &v2
		}
	}
	if v1, ok := d.GetOk("alicloud_function"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("alicloud_function", sv)
				diags = append(diags, e)
			}
			obj.AlicloudFunction = &v2
		}
	}
	if v1, ok := d.GetOk("alicloud_function_authorization"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("alicloud_function_authorization", sv)
				diags = append(diags, e)
			}
			obj.AlicloudFunctionAuthorization = &v2
		}
	}
	if v1, ok := d.GetOk("alicloud_function_domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("alicloud_function_domain", sv)
				diags = append(diags, e)
			}
			obj.AlicloudFunctionDomain = &v2
		}
	}
	if v1, ok := d.GetOk("alicloud_region"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("alicloud_region", sv)
				diags = append(diags, e)
			}
			obj.AlicloudRegion = &v2
		}
	}
	if v1, ok := d.GetOk("alicloud_service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("alicloud_service", sv)
				diags = append(diags, e)
			}
			obj.AlicloudService = &v2
		}
	}
	if v1, ok := d.GetOk("alicloud_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("alicloud_version", sv)
				diags = append(diags, e)
			}
			obj.AlicloudVersion = &v2
		}
	}
	if v1, ok := d.GetOk("aws_api_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("aws_api_id", sv)
				diags = append(diags, e)
			}
			obj.AwsApiId = &v2
		}
	}
	if v1, ok := d.GetOk("aws_api_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("aws_api_key", sv)
				diags = append(diags, e)
			}
			obj.AwsApiKey = &v2
		}
	}
	if v1, ok := d.GetOk("aws_api_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("aws_api_path", sv)
				diags = append(diags, e)
			}
			obj.AwsApiPath = &v2
		}
	}
	if v1, ok := d.GetOk("aws_api_stage"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("aws_api_stage", sv)
				diags = append(diags, e)
			}
			obj.AwsApiStage = &v2
		}
	}
	if v1, ok := d.GetOk("aws_domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("aws_domain", sv)
				diags = append(diags, e)
			}
			obj.AwsDomain = &v2
		}
	}
	if v1, ok := d.GetOk("aws_region"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("aws_region", sv)
				diags = append(diags, e)
			}
			obj.AwsRegion = &v2
		}
	}
	if v1, ok := d.GetOk("azure_api_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("azure_api_key", sv)
				diags = append(diags, e)
			}
			obj.AzureApiKey = &v2
		}
	}
	if v1, ok := d.GetOk("azure_app"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("azure_app", sv)
				diags = append(diags, e)
			}
			obj.AzureApp = &v2
		}
	}
	if v1, ok := d.GetOk("azure_domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("azure_domain", sv)
				diags = append(diags, e)
			}
			obj.AzureDomain = &v2
		}
	}
	if v1, ok := d.GetOk("azure_function"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("azure_function", sv)
				diags = append(diags, e)
			}
			obj.AzureFunction = &v2
		}
	}
	if v1, ok := d.GetOk("azure_function_authorization"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("azure_function_authorization", sv)
				diags = append(diags, e)
			}
			obj.AzureFunctionAuthorization = &v2
		}
	}
	if v1, ok := d.GetOk("delay"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.0.1") {
				e := utils.AttributeVersionWarning("delay", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Delay = &tmp
		}
	}
	if v1, ok := d.GetOk("description"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("description", sv)
				diags = append(diags, e)
			}
			obj.Description = &v2
		}
	}
	if v1, ok := d.GetOk("email_body"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("email_body", sv)
				diags = append(diags, e)
			}
			obj.EmailBody = &v2
		}
	}
	if v1, ok := d.GetOk("email_from"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("email_from", sv)
				diags = append(diags, e)
			}
			obj.EmailFrom = &v2
		}
	}
	if v1, ok := d.GetOk("email_subject"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("email_subject", sv)
				diags = append(diags, e)
			}
			obj.EmailSubject = &v2
		}
	}
	if v, ok := d.GetOk("email_to"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("email_to", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAutomationActionEmailTo(d, v, "email_to", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.EmailTo = t
		}
	} else if d.HasChange("email_to") {
		old, new := d.GetChange("email_to")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.EmailTo = &[]models.SystemAutomationActionEmailTo{}
		}
	}
	if v1, ok := d.GetOk("execute_security_fabric"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("execute_security_fabric", sv)
				diags = append(diags, e)
			}
			obj.ExecuteSecurityFabric = &v2
		}
	}
	if v1, ok := d.GetOk("gcp_function"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("gcp_function", sv)
				diags = append(diags, e)
			}
			obj.GcpFunction = &v2
		}
	}
	if v1, ok := d.GetOk("gcp_function_domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("gcp_function_domain", sv)
				diags = append(diags, e)
			}
			obj.GcpFunctionDomain = &v2
		}
	}
	if v1, ok := d.GetOk("gcp_function_region"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("gcp_function_region", sv)
				diags = append(diags, e)
			}
			obj.GcpFunctionRegion = &v2
		}
	}
	if v1, ok := d.GetOk("gcp_project"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("gcp_project", sv)
				diags = append(diags, e)
			}
			obj.GcpProject = &v2
		}
	}
	if v, ok := d.GetOk("headers"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("headers", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAutomationActionHeaders(d, v, "headers", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Headers = t
		}
	} else if d.HasChange("headers") {
		old, new := d.GetChange("headers")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Headers = &[]models.SystemAutomationActionHeaders{}
		}
	}
	if v1, ok := d.GetOk("http_body"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_body", sv)
				diags = append(diags, e)
			}
			obj.HttpBody = &v2
		}
	}
	if v1, ok := d.GetOk("message"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("message", sv)
				diags = append(diags, e)
			}
			obj.Message = &v2
		}
	}
	if v1, ok := d.GetOk("message_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("message_type", sv)
				diags = append(diags, e)
			}
			obj.MessageType = &v2
		}
	}
	if v1, ok := d.GetOk("method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("method", sv)
				diags = append(diags, e)
			}
			obj.Method = &v2
		}
	}
	if v1, ok := d.GetOk("minimum_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("minimum_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MinimumInterval = &tmp
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
	if v1, ok := d.GetOk("port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port = &tmp
		}
	}
	if v1, ok := d.GetOk("protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("protocol", sv)
				diags = append(diags, e)
			}
			obj.Protocol = &v2
		}
	}
	if v1, ok := d.GetOk("replacement_message"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("replacement_message", sv)
				diags = append(diags, e)
			}
			obj.ReplacementMessage = &v2
		}
	}
	if v1, ok := d.GetOk("replacemsg_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("replacemsg_group", sv)
				diags = append(diags, e)
			}
			obj.ReplacemsgGroup = &v2
		}
	}
	if v1, ok := d.GetOk("required"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.1") {
				e := utils.AttributeVersionWarning("required", sv)
				diags = append(diags, e)
			}
			obj.Required = &v2
		}
	}
	if v1, ok := d.GetOk("script"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("script", sv)
				diags = append(diags, e)
			}
			obj.Script = &v2
		}
	}
	if v, ok := d.GetOk("sdn_connector"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("sdn_connector", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAutomationActionSdnConnector(d, v, "sdn_connector", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SdnConnector = t
		}
	} else if d.HasChange("sdn_connector") {
		old, new := d.GetChange("sdn_connector")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SdnConnector = &[]models.SystemAutomationActionSdnConnector{}
		}
	}
	if v1, ok := d.GetOk("security_tag"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_tag", sv)
				diags = append(diags, e)
			}
			obj.SecurityTag = &v2
		}
	}
	if v1, ok := d.GetOk("tls_certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tls_certificate", sv)
				diags = append(diags, e)
			}
			obj.TlsCertificate = &v2
		}
	}
	if v1, ok := d.GetOk("uri"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uri", sv)
				diags = append(diags, e)
			}
			obj.Uri = &v2
		}
	}
	if v1, ok := d.GetOk("verify_host_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("verify_host_cert", sv)
				diags = append(diags, e)
			}
			obj.VerifyHostCert = &v2
		}
	}
	return &obj, diags
}
