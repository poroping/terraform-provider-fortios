// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure connection to SDN Connector.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemSdnConnector() *schema.Resource {
	return &schema.Resource{
		Description: "Configure connection to SDN Connector.",

		ReadContext: dataSourceSystemSdnConnectorRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"access_key": {
				Type:        schema.TypeString,
				Description: "AWS / ACS access key ID.",
				Computed:    true,
			},
			"api_key": {
				Type:        schema.TypeString,
				Description: "IBM cloud API key or service ID API key.",
				Computed:    true,
				Sensitive:   true,
			},
			"azure_region": {
				Type:        schema.TypeString,
				Description: "Azure server region.",
				Computed:    true,
			},
			"client_id": {
				Type:        schema.TypeString,
				Description: "Azure client ID (application ID).",
				Computed:    true,
			},
			"client_secret": {
				Type:        schema.TypeString,
				Description: "Azure client secret (application key).",
				Computed:    true,
				Sensitive:   true,
			},
			"compartment_id": {
				Type:        schema.TypeString,
				Description: "Compartment ID.",
				Computed:    true,
			},
			"compute_generation": {
				Type:        schema.TypeInt,
				Description: "Compute generation for IBM cloud infrastructure.",
				Computed:    true,
			},
			"domain": {
				Type:        schema.TypeString,
				Description: "Domain name.",
				Computed:    true,
			},
			"external_account_list": {
				Type:        schema.TypeList,
				Description: "Configure AWS external account list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region_list": {
							Type:        schema.TypeList,
							Description: "AWS region name list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": {
										Type:        schema.TypeString,
										Description: "AWS region name.",
										Computed:    true,
									},
								},
							},
						},
						"role_arn": {
							Type:        schema.TypeString,
							Description: "AWS role ARN to assume.",
							Computed:    true,
						},
					},
				},
			},
			"external_ip": {
				Type:        schema.TypeList,
				Description: "Configure GCP external IP.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "External IP name.",
							Computed:    true,
						},
					},
				},
			},
			"forwarding_rule": {
				Type:        schema.TypeList,
				Description: "Configure GCP forwarding rule.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_name": {
							Type:        schema.TypeString,
							Description: "Forwarding rule name.",
							Computed:    true,
						},
						"target": {
							Type:        schema.TypeString,
							Description: "Target instance name.",
							Computed:    true,
						},
					},
				},
			},
			"gcp_project": {
				Type:        schema.TypeString,
				Description: "GCP project name.",
				Computed:    true,
			},
			"gcp_project_list": {
				Type:        schema.TypeList,
				Description: "Configure GCP project list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gcp_zone_list": {
							Type:        schema.TypeList,
							Description: "Configure GCP zone list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "GCP zone name.",
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeString,
							Description: "GCP project ID.",
							Computed:    true,
						},
					},
				},
			},
			"group_name": {
				Type:        schema.TypeString,
				Description: "Group name of computers.",
				Computed:    true,
			},
			"ha_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable use for FortiGate HA service.",
				Computed:    true,
			},
			"ibm_region": {
				Type:        schema.TypeString,
				Description: "IBM cloud region name.",
				Computed:    true,
			},
			"login_endpoint": {
				Type:        schema.TypeString,
				Description: "Azure Stack login endpoint.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "SDN connector name.",
				Required:    true,
			},
			"nic": {
				Type:        schema.TypeList,
				Description: "Configure Azure network interface.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeList,
							Description: "Configure IP configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "IP configuration name.",
										Computed:    true,
									},
									"public_ip": {
										Type:        schema.TypeString,
										Description: "Public IP name.",
										Computed:    true,
									},
									"resource_group": {
										Type:        schema.TypeString,
										Description: "Resource group of Azure public IP.",
										Computed:    true,
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Network interface name.",
							Computed:    true,
						},
					},
				},
			},
			"oci_cert": {
				Type:        schema.TypeString,
				Description: "OCI certificate.",
				Computed:    true,
			},
			"oci_fingerprint": {
				Type:        schema.TypeString,
				Description: "OCI pubkey fingerprint.",
				Computed:    true,
			},
			"oci_region": {
				Type:        schema.TypeString,
				Description: "OCI server region.",
				Computed:    true,
			},
			"oci_region_type": {
				Type:        schema.TypeString,
				Description: "OCI region type.",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password of the remote SDN connector as login credentials.",
				Computed:    true,
				Sensitive:   true,
			},
			"private_key": {
				Type:        schema.TypeString,
				Description: "Private key of GCP service account.",
				Computed:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "AWS / ACS region name.",
				Computed:    true,
			},
			"resource_group": {
				Type:        schema.TypeString,
				Description: "Azure resource group.",
				Computed:    true,
			},
			"resource_url": {
				Type:        schema.TypeString,
				Description: "Azure Stack resource URL.",
				Computed:    true,
			},
			"route": {
				Type:        schema.TypeList,
				Description: "Configure GCP route.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Route name.",
							Computed:    true,
						},
					},
				},
			},
			"route_table": {
				Type:        schema.TypeList,
				Description: "Configure Azure route table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Route table name.",
							Computed:    true,
						},
						"resource_group": {
							Type:        schema.TypeString,
							Description: "Resource group of Azure route table.",
							Computed:    true,
						},
						"route": {
							Type:        schema.TypeList,
							Description: "Configure Azure route.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Route name.",
										Computed:    true,
									},
									"next_hop": {
										Type:        schema.TypeString,
										Description: "Next hop address.",
										Computed:    true,
									},
								},
							},
						},
						"subscription_id": {
							Type:        schema.TypeString,
							Description: "Subscription ID of Azure route table.",
							Computed:    true,
						},
					},
				},
			},
			"secret_key": {
				Type:        schema.TypeString,
				Description: "AWS / ACS secret access key.",
				Computed:    true,
				Sensitive:   true,
			},
			"secret_token": {
				Type:        schema.TypeString,
				Description: "Secret token of Kubernetes service account.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "Server address of the remote SDN connector.",
				Computed:    true,
			},
			"server_list": {
				Type:        schema.TypeList,
				Description: "Server address list of the remote SDN connector.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeString,
							Description: "IPv4 address.",
							Computed:    true,
						},
					},
				},
			},
			"server_port": {
				Type:        schema.TypeInt,
				Description: "Port number of the remote SDN connector.",
				Computed:    true,
			},
			"service_account": {
				Type:        schema.TypeString,
				Description: "GCP service account email.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable connection to the remote SDN connector.",
				Computed:    true,
			},
			"subscription_id": {
				Type:        schema.TypeString,
				Description: "Azure subscription ID.",
				Computed:    true,
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Description: "Tenant ID (directory ID).",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of SDN connector.",
				Computed:    true,
			},
			"update_interval": {
				Type:        schema.TypeInt,
				Description: "Dynamic object update interval (30 - 3600 sec, default = 60, 0 = disabled).",
				Computed:    true,
			},
			"use_metadata_iam": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of IAM role from metadata to call API.",
				Computed:    true,
			},
			"user_id": {
				Type:        schema.TypeString,
				Description: "User ID.",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "Username of the remote SDN connector as login credentials.",
				Computed:    true,
			},
			"vcenter_password": {
				Type:        schema.TypeString,
				Description: "vCenter server password for NSX quarantine.",
				Computed:    true,
				Sensitive:   true,
			},
			"vcenter_server": {
				Type:        schema.TypeString,
				Description: "vCenter server address for NSX quarantine.",
				Computed:    true,
			},
			"vcenter_username": {
				Type:        schema.TypeString,
				Description: "vCenter server username for NSX quarantine.",
				Computed:    true,
			},
			"verify_certificate": {
				Type:        schema.TypeString,
				Description: "Enable/disable server certificate verification.",
				Computed:    true,
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Description: "AWS VPC ID.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemSdnConnectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSdnConnector(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSdnConnector dataSource: %v", err)
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

	diags := refreshObjectSystemSdnConnector(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
