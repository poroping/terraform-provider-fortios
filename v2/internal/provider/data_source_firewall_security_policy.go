// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure NGFW IPv4/IPv6 application policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure NGFW IPv4/IPv6 application policies.",

		ReadContext: dataSourceFirewallSecurityPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"action": {
				Type:        schema.TypeString,
				Description: "Policy action (accept/deny).",
				Computed:    true,
			},
			"app_category": {
				Type:        schema.TypeList,
				Description: "Application category ID list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Category IDs.",
							Computed:    true,
						},
					},
				},
			},
			"app_group": {
				Type:        schema.TypeList,
				Description: "Application group names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Application group names.",
							Computed:    true,
						},
					},
				},
			},
			"application": {
				Type:        schema.TypeList,
				Description: "Application ID list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Application IDs.",
							Computed:    true,
						},
					},
				},
			},
			"application_list": {
				Type:        schema.TypeString,
				Description: "Name of an existing Application list.",
				Computed:    true,
			},
			"av_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing Antivirus profile.",
				Computed:    true,
			},
			"cifs_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing CIFS profile.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"dlp_sensor": {
				Type:        schema.TypeString,
				Description: "Name of an existing DLP sensor.",
				Computed:    true,
			},
			"dnsfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing DNS filter profile.",
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination IPv4 address name and address group names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"dstaddr_negate": {
				Type:        schema.TypeString,
				Description: "When enabled dstaddr/dstaddr6 specifies what the destination address must NOT be.",
				Computed:    true,
			},
			"dstaddr4": {
				Type:        schema.TypeList,
				Description: "Destination IPv4 address name and address group names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"dstaddr6": {
				Type:        schema.TypeList,
				Description: "Destination IPv6 address name and address group names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"dstintf": {
				Type:        schema.TypeList,
				Description: "Outgoing (egress) interface.",
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
			"emailfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing email filter profile.",
				Computed:    true,
			},
			"enforce_default_app_port": {
				Type:        schema.TypeString,
				Description: "Enable/disable default application port enforcement for allowed applications.",
				Computed:    true,
			},
			"file_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing file-filter profile.",
				Computed:    true,
			},
			"fsso_groups": {
				Type:        schema.TypeList,
				Description: "Names of FSSO groups.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Names of FSSO groups.",
							Computed:    true,
						},
					},
				},
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "Names of user groups that can authenticate with this policy.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "User group name.",
							Computed:    true,
						},
					},
				},
			},
			"icap_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing ICAP profile.",
				Computed:    true,
			},
			"internet_service": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.",
				Computed:    true,
			},
			"internet_service_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_group": {
				Type:        schema.TypeList,
				Description: "Internet Service group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_id": {
				Type:        schema.TypeList,
				Description: "Internet Service ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Internet Service ID.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_name": {
				Type:        schema.TypeList,
				Description: "Internet Service name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_negate": {
				Type:        schema.TypeString,
				Description: "When enabled internet-service specifies what the service must NOT be.",
				Computed:    true,
			},
			"internet_service_src": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.",
				Computed:    true,
			},
			"internet_service_src_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service source name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service source group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_group": {
				Type:        schema.TypeList,
				Description: "Internet Service source group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_id": {
				Type:        schema.TypeList,
				Description: "Internet Service source ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Internet Service ID.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_name": {
				Type:        schema.TypeList,
				Description: "Internet Service source name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_negate": {
				Type:        schema.TypeString,
				Description: "When enabled internet-service-src specifies what the service must NOT be.",
				Computed:    true,
			},
			"ips_sensor": {
				Type:        schema.TypeString,
				Description: "Name of an existing IPS sensor.",
				Computed:    true,
			},
			"learning_mode": {
				Type:        schema.TypeString,
				Description: "Enable to allow everything, but log all of the meaningful data for security information gathering. A learning report will be generated.",
				Computed:    true,
			},
			"logtraffic": {
				Type:        schema.TypeString,
				Description: "Enable or disable logging. Log all sessions or security profile sessions.",
				Computed:    true,
			},
			"logtraffic_start": {
				Type:        schema.TypeString,
				Description: "Record logs when a session starts.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Policy name.",
				Computed:    true,
			},
			"nat46": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAT46.",
				Computed:    true,
			},
			"nat64": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAT64.",
				Computed:    true,
			},
			"policyid": {
				Type:        schema.TypeInt,
				Description: "Policy ID.",
				Required:    true,
			},
			"profile_group": {
				Type:        schema.TypeString,
				Description: "Name of profile group.",
				Computed:    true,
			},
			"profile_protocol_options": {
				Type:        schema.TypeString,
				Description: "Name of an existing Protocol options profile.",
				Computed:    true,
			},
			"profile_type": {
				Type:        schema.TypeString,
				Description: "Determine whether the firewall policy allows security profile groups or single profiles only.",
				Computed:    true,
			},
			"schedule": {
				Type:        schema.TypeString,
				Description: "Schedule name.",
				Computed:    true,
			},
			"sctp_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing SCTP filter profile.",
				Computed:    true,
			},
			"send_deny_packet": {
				Type:        schema.TypeString,
				Description: "Enable to send a reply when a session is denied or blocked by a firewall policy.",
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Service and service group names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Service name.",
							Computed:    true,
						},
					},
				},
			},
			"service_negate": {
				Type:        schema.TypeString,
				Description: "When enabled service specifies what the service must NOT be.",
				Computed:    true,
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source IPv4 address name and address group names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"srcaddr_negate": {
				Type:        schema.TypeString,
				Description: "When enabled srcaddr/srcaddr6 specifies what the source address must NOT be.",
				Computed:    true,
			},
			"srcaddr4": {
				Type:        schema.TypeList,
				Description: "Source IPv4 address name and address group names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"srcaddr6": {
				Type:        schema.TypeList,
				Description: "Source IPv6 address name and address group names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"srcintf": {
				Type:        schema.TypeList,
				Description: "Incoming (ingress) interface.",
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
			"ssh_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing SSH filter profile.",
				Computed:    true,
			},
			"ssl_ssh_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing SSL SSH profile.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable or disable this policy.",
				Computed:    true,
			},
			"url_category": {
				Type:        schema.TypeList,
				Description: "URL category ID list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "URL category ID.",
							Computed:    true,
						},
					},
				},
			},
			"users": {
				Type:        schema.TypeList,
				Description: "Names of individual users that can authenticate with this policy.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "User name.",
							Computed:    true,
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
			"videofilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing VideoFilter profile.",
				Computed:    true,
			},
			"voip_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing VoIP profile.",
				Computed:    true,
			},
			"webfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing Web filter profile.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallSecurityPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallSecurityPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSecurityPolicy dataSource: %v", err)
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

	diags := refreshObjectFirewallSecurityPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
