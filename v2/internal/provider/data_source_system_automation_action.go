// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Action for automation stitches.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemAutomationAction() *schema.Resource {
	return &schema.Resource{
		Description: "Action for automation stitches.",

		ReadContext: dataSourceSystemAutomationActionRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"accprofile": {
				Type:        schema.TypeString,
				Description: "Access profile for CLI script action to access FortiGate features.",
				Computed:    true,
			},
			"action_type": {
				Type:        schema.TypeString,
				Description: "Action type.",
				Computed:    true,
			},
			"alicloud_access_key_id": {
				Type:        schema.TypeString,
				Description: "AliCloud AccessKey ID.",
				Computed:    true,
			},
			"alicloud_access_key_secret": {
				Type:        schema.TypeString,
				Description: "AliCloud AccessKey secret.",
				Computed:    true,
				Sensitive:   true,
			},
			"alicloud_account_id": {
				Type:        schema.TypeString,
				Description: "AliCloud account ID.",
				Computed:    true,
			},
			"alicloud_function": {
				Type:        schema.TypeString,
				Description: "AliCloud function name.",
				Computed:    true,
			},
			"alicloud_function_authorization": {
				Type:        schema.TypeString,
				Description: "AliCloud function authorization type.",
				Computed:    true,
			},
			"alicloud_function_domain": {
				Type:        schema.TypeString,
				Description: "AliCloud function domain.",
				Computed:    true,
			},
			"alicloud_region": {
				Type:        schema.TypeString,
				Description: "AliCloud region.",
				Computed:    true,
			},
			"alicloud_service": {
				Type:        schema.TypeString,
				Description: "AliCloud service name.",
				Computed:    true,
			},
			"alicloud_version": {
				Type:        schema.TypeString,
				Description: "AliCloud version.",
				Computed:    true,
			},
			"aws_api_id": {
				Type:        schema.TypeString,
				Description: "AWS API Gateway ID.",
				Computed:    true,
			},
			"aws_api_key": {
				Type:        schema.TypeString,
				Description: "AWS API Gateway API key.",
				Computed:    true,
				Sensitive:   true,
			},
			"aws_api_path": {
				Type:        schema.TypeString,
				Description: "AWS API Gateway path.",
				Computed:    true,
			},
			"aws_api_stage": {
				Type:        schema.TypeString,
				Description: "AWS API Gateway deployment stage name.",
				Computed:    true,
			},
			"aws_domain": {
				Type:        schema.TypeString,
				Description: "AWS domain.",
				Computed:    true,
			},
			"aws_region": {
				Type:        schema.TypeString,
				Description: "AWS region.",
				Computed:    true,
			},
			"azure_api_key": {
				Type:        schema.TypeString,
				Description: "Azure function API key.",
				Computed:    true,
				Sensitive:   true,
			},
			"azure_app": {
				Type:        schema.TypeString,
				Description: "Azure function application name.",
				Computed:    true,
			},
			"azure_domain": {
				Type:        schema.TypeString,
				Description: "Azure function domain.",
				Computed:    true,
			},
			"azure_function": {
				Type:        schema.TypeString,
				Description: "Azure function name.",
				Computed:    true,
			},
			"azure_function_authorization": {
				Type:        schema.TypeString,
				Description: "Azure function authorization level.",
				Computed:    true,
			},
			"delay": {
				Type:        schema.TypeInt,
				Description: "Delay before execution (in seconds).",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description.",
				Computed:    true,
			},
			"email_body": {
				Type:        schema.TypeString,
				Description: "Email body.",
				Computed:    true,
			},
			"email_from": {
				Type:        schema.TypeString,
				Description: "Email sender name.",
				Computed:    true,
			},
			"email_subject": {
				Type:        schema.TypeString,
				Description: "Email subject.",
				Computed:    true,
			},
			"email_to": {
				Type:        schema.TypeList,
				Description: "Email addresses.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Email address.",
							Computed:    true,
						},
					},
				},
			},
			"execute_security_fabric": {
				Type:        schema.TypeString,
				Description: "Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric.",
				Computed:    true,
			},
			"gcp_function": {
				Type:        schema.TypeString,
				Description: "Google Cloud function name.",
				Computed:    true,
			},
			"gcp_function_domain": {
				Type:        schema.TypeString,
				Description: "Google Cloud function domain.",
				Computed:    true,
			},
			"gcp_function_region": {
				Type:        schema.TypeString,
				Description: "Google Cloud function region.",
				Computed:    true,
			},
			"gcp_project": {
				Type:        schema.TypeString,
				Description: "Google Cloud Platform project name.",
				Computed:    true,
			},
			"headers": {
				Type:        schema.TypeList,
				Description: "Request headers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header": {
							Type:        schema.TypeString,
							Description: "Request header.",
							Computed:    true,
						},
					},
				},
			},
			"http_body": {
				Type:        schema.TypeString,
				Description: "Request body (if necessary). Should be serialized json string.",
				Computed:    true,
			},
			"message": {
				Type:        schema.TypeString,
				Description: "Message content.",
				Computed:    true,
			},
			"message_type": {
				Type:        schema.TypeString,
				Description: "Message type.",
				Computed:    true,
			},
			"method": {
				Type:        schema.TypeString,
				Description: "Request method (POST, PUT, GET, PATCH or DELETE).",
				Computed:    true,
			},
			"minimum_interval": {
				Type:        schema.TypeInt,
				Description: "Limit execution to no more than once in this interval (in seconds).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Protocol port.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Request protocol.",
				Computed:    true,
			},
			"replacement_message": {
				Type:        schema.TypeString,
				Description: "Enable/disable replacement message.",
				Computed:    true,
			},
			"replacemsg_group": {
				Type:        schema.TypeString,
				Description: "Replacement message group.",
				Computed:    true,
			},
			"required": {
				Type:        schema.TypeString,
				Description: "Required in action chain.",
				Computed:    true,
			},
			"script": {
				Type:        schema.TypeString,
				Description: "CLI script.",
				Computed:    true,
			},
			"sdn_connector": {
				Type:        schema.TypeList,
				Description: "NSX SDN connector names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "SDN connector name.",
							Computed:    true,
						},
					},
				},
			},
			"security_tag": {
				Type:        schema.TypeString,
				Description: "NSX security tag.",
				Computed:    true,
			},
			"tls_certificate": {
				Type:        schema.TypeString,
				Description: "Custom TLS certificate for API request.",
				Computed:    true,
			},
			"uri": {
				Type:        schema.TypeString,
				Description: "Request API URI.",
				Computed:    true,
			},
			"verify_host_cert": {
				Type:        schema.TypeString,
				Description: "Enable/disable verification of the remote host certificate.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemAutomationActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAutomationAction(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAutomationAction dataSource: %v", err)
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

	diags := refreshObjectSystemAutomationAction(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
