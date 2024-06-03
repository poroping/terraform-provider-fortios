// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure connection to SDN Connector.

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

func resourceSystemSdnConnector() *schema.Resource {
	return &schema.Resource{
		Description: "Configure connection to SDN Connector.",

		CreateContext: resourceSystemSdnConnectorCreate,
		ReadContext:   resourceSystemSdnConnectorRead,
		UpdateContext: resourceSystemSdnConnectorUpdate,
		DeleteContext: resourceSystemSdnConnectorDelete,

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
			"access_key": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "AWS / ACS access key ID.",
				Optional:    true,
				Computed:    true,
			},
			"alt_resource_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable AWS alternative resource IP.",
				Optional:    true,
				Computed:    true,
			},
			"api_key": {
				Type: schema.TypeString,

				Description: "IBM cloud API key or service ID API key.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"azure_region": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"global", "china", "germany", "usgov", "local"}, false),

				Description: "Azure server region.",
				Optional:    true,
				Computed:    true,
			},
			"client_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Azure client ID (application ID).",
				Optional:    true,
				Computed:    true,
			},
			"client_secret": {
				Type: schema.TypeString,

				Description: "Azure client secret (application key).",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"compartment_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Compartment ID.",
				Optional:    true,
				Computed:    true,
			},
			"compute_generation": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2),

				Description: "Compute generation for IBM cloud infrastructure.",
				Optional:    true,
				Computed:    true,
			},
			"domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Domain name.",
				Optional:    true,
				Computed:    true,
			},
			"external_account_list": {
				Type:        schema.TypeList,
				Description: "Configure AWS external account list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1399),

							Description: "AWS external ID.",
							Optional:    true,
							Computed:    true,
						},
						"region_list": {
							Type:        schema.TypeList,
							Description: "AWS region name list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),

										Description: "AWS region name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"role_arn": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 2047),

							Description: "AWS role ARN to assume.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"external_ip": {
				Type:        schema.TypeList,
				Description: "Configure GCP external IP.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "External IP name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"forwarding_rule": {
				Type:        schema.TypeList,
				Description: "Configure GCP forwarding rule.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Forwarding rule name.",
							Optional:    true,
							Computed:    true,
						},
						"target": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Target instance name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"gcp_project": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "GCP project name.",
				Optional:    true,
				Computed:    true,
			},
			"gcp_project_list": {
				Type:        schema.TypeList,
				Description: "Configure GCP project list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gcp_zone_list": {
							Type:        schema.TypeList,
							Description: "Configure GCP zone list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),

										Description: "GCP zone name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "GCP project ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"group_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Full path group name of computers.",
				Optional:    true,
				Computed:    true,
			},
			"ha_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable use for FortiGate HA service.",
				Optional:    true,
				Computed:    true,
			},
			"ibm_region": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"dallas", "washington-dc", "london", "frankfurt", "sydney", "tokyo", "osaka", "toronto", "sao-paulo"}, false),

				Description: "IBM cloud region name.",
				Optional:    true,
				Computed:    true,
			},
			"login_endpoint": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Azure Stack login endpoint.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SDN connector name.",
				ForceNew:    true,
				Required:    true,
			},
			"nic": {
				Type:        schema.TypeList,
				Description: "Configure Azure network interface.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeList,
							Description: "Configure IP configuration.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "IP configuration name.",
										Optional:    true,
										Computed:    true,
									},
									"public_ip": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "Public IP name.",
										Optional:    true,
										Computed:    true,
									},
									"resource_group": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "Resource group of Azure public IP.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Network interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"oci_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "OCI certificate.",
				Optional:    true,
				Computed:    true,
			},
			"oci_fingerprint": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "OCI pubkey fingerprint.",
				Optional:    true,
				Computed:    true,
			},
			"oci_region": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "OCI server region.",
				Optional:    true,
				Computed:    true,
			},
			"oci_region_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"commercial", "government"}, false),

				Description: "OCI region type.",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Password of the remote SDN connector as login credentials.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"private_key": {
				Type: schema.TypeString,

				Description: "Private key of GCP service account.",
				Optional:    true,
				Computed:    true,
			},
			"region": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "AWS / ACS region name.",
				Optional:    true,
				Computed:    true,
			},
			"resource_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Azure resource group.",
				Optional:    true,
				Computed:    true,
			},
			"resource_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Azure Stack resource URL.",
				Optional:    true,
				Computed:    true,
			},
			"route": {
				Type:        schema.TypeList,
				Description: "Configure GCP route.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Route name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"route_table": {
				Type:        schema.TypeList,
				Description: "Configure Azure route table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Route table name.",
							Optional:    true,
							Computed:    true,
						},
						"resource_group": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Resource group of Azure route table.",
							Optional:    true,
							Computed:    true,
						},
						"route": {
							Type:        schema.TypeList,
							Description: "Configure Azure route.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "Route name.",
										Optional:    true,
										Computed:    true,
									},
									"next_hop": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),

										Description: "Next hop address.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"subscription_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Subscription ID of Azure route table.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"secret_key": {
				Type: schema.TypeString,

				Description: "AWS / ACS secret access key.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"secret_token": {
				Type: schema.TypeString,

				Description: "Secret token of Kubernetes service account.",
				Optional:    true,
				Computed:    true,
			},
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Server address of the remote SDN connector.",
				Optional:    true,
				Computed:    true,
			},
			"server_ca_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Trust only those servers whose certificate is directly/indirectly signed by this certificate.",
				Optional:    true,
				Computed:    true,
			},
			"server_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Trust servers that contain this certificate only.",
				Optional:    true,
				Computed:    true,
			},
			"server_list": {
				Type:        schema.TypeList,
				Description: "Server address list of the remote SDN connector.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "IPv4 address.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"server_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port number of the remote SDN connector.",
				Optional:    true,
				Computed:    true,
			},
			"service_account": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "GCP service account email.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable connection to the remote SDN connector.",
				Optional:    true,
				Computed:    true,
			},
			"subscription_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Azure subscription ID.",
				Optional:    true,
				Computed:    true,
			},
			"tenant_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Tenant ID (directory ID).",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"aci", "alicloud", "aws", "azure", "gcp", "nsx", "nuage", "oci", "openstack", "kubernetes", "vmware", "sepm", "aci-direct", "ibm", "nutanix", "sap"}, false),

				Description: "Type of SDN connector.",
				Optional:    true,
				Computed:    true,
			},
			"update_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),

				Description: "Dynamic object update interval (30 - 3600 sec, default = 60, 0 = disabled).",
				Optional:    true,
				Computed:    true,
			},
			"use_metadata_iam": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable use of IAM role from metadata to call API.",
				Optional:    true,
				Computed:    true,
			},
			"user_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "User ID.",
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "Username of the remote SDN connector as login credentials.",
				Optional:    true,
				Computed:    true,
			},
			"vcenter_password": {
				Type: schema.TypeString,

				Description: "vCenter server password for NSX quarantine.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"vcenter_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "vCenter server address for NSX quarantine.",
				Optional:    true,
				Computed:    true,
			},
			"vcenter_username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "vCenter server username for NSX quarantine.",
				Optional:    true,
				Computed:    true,
			},
			"verify_certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable server certificate verification.",
				Optional:    true,
				Computed:    true,
			},
			"vpc_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "AWS VPC ID.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemSdnConnectorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemSdnConnector resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemSdnConnector(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSdnConnector(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSdnConnector")
	}

	return resourceSystemSdnConnectorRead(ctx, d, meta)
}

func resourceSystemSdnConnectorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSdnConnector(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSdnConnector(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSdnConnector resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSdnConnector")
	}

	return resourceSystemSdnConnectorRead(ctx, d, meta)
}

func resourceSystemSdnConnectorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSdnConnector(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSdnConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdnConnectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSdnConnector(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSdnConnector resource: %v", err)
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

	diags := refreshObjectSystemSdnConnector(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemSdnConnectorExternalAccountList(d *schema.ResourceData, v *[]models.SystemSdnConnectorExternalAccountList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ExternalId; tmp != nil {
				v["external_id"] = *tmp
			}

			if tmp := cfg.RegionList; tmp != nil {
				v["region_list"] = flattenSystemSdnConnectorExternalAccountListRegionList(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "region_list"), sort)
			}

			if tmp := cfg.RoleArn; tmp != nil {
				v["role_arn"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "role_arn")
	}

	return flat
}

func flattenSystemSdnConnectorExternalAccountListRegionList(d *schema.ResourceData, v *[]models.SystemSdnConnectorExternalAccountListRegionList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Region; tmp != nil {
				v["region"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "region")
	}

	return flat
}

func flattenSystemSdnConnectorExternalIp(d *schema.ResourceData, v *[]models.SystemSdnConnectorExternalIp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func flattenSystemSdnConnectorForwardingRule(d *schema.ResourceData, v *[]models.SystemSdnConnectorForwardingRule, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.RuleName; tmp != nil {
				v["rule_name"] = *tmp
			}

			if tmp := cfg.Target; tmp != nil {
				v["target"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "rule_name")
	}

	return flat
}

func flattenSystemSdnConnectorGcpProjectList(d *schema.ResourceData, v *[]models.SystemSdnConnectorGcpProjectList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.GcpZoneList; tmp != nil {
				v["gcp_zone_list"] = flattenSystemSdnConnectorGcpProjectListGcpZoneList(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "gcp_zone_list"), sort)
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

func flattenSystemSdnConnectorGcpProjectListGcpZoneList(d *schema.ResourceData, v *[]models.SystemSdnConnectorGcpProjectListGcpZoneList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func flattenSystemSdnConnectorNic(d *schema.ResourceData, v *[]models.SystemSdnConnectorNic, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = flattenSystemSdnConnectorNicIp(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "ip"), sort)
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

func flattenSystemSdnConnectorNicIp(d *schema.ResourceData, v *[]models.SystemSdnConnectorNicIp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.PublicIp; tmp != nil {
				v["public_ip"] = *tmp
			}

			if tmp := cfg.ResourceGroup; tmp != nil {
				v["resource_group"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemSdnConnectorRoute(d *schema.ResourceData, v *[]models.SystemSdnConnectorRoute, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func flattenSystemSdnConnectorRouteTable(d *schema.ResourceData, v *[]models.SystemSdnConnectorRouteTable, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.ResourceGroup; tmp != nil {
				v["resource_group"] = *tmp
			}

			if tmp := cfg.Route; tmp != nil {
				v["route"] = flattenSystemSdnConnectorRouteTableRoute(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "route"), sort)
			}

			if tmp := cfg.SubscriptionId; tmp != nil {
				v["subscription_id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemSdnConnectorRouteTableRoute(d *schema.ResourceData, v *[]models.SystemSdnConnectorRouteTableRoute, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.NextHop; tmp != nil {
				v["next_hop"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemSdnConnectorServerList(d *schema.ResourceData, v *[]models.SystemSdnConnectorServerList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "ip")
	}

	return flat
}

func refreshObjectSystemSdnConnector(d *schema.ResourceData, o *models.SystemSdnConnector, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AccessKey != nil {
		v := *o.AccessKey

		if err = d.Set("access_key", v); err != nil {
			return diag.Errorf("error reading access_key: %v", err)
		}
	}

	if o.AltResourceIp != nil {
		v := *o.AltResourceIp

		if err = d.Set("alt_resource_ip", v); err != nil {
			return diag.Errorf("error reading alt_resource_ip: %v", err)
		}
	}

	if o.ApiKey != nil {
		v := *o.ApiKey

		if v == "ENC XXXX" {
		} else if err = d.Set("api_key", v); err != nil {
			return diag.Errorf("error reading api_key: %v", err)
		}
	}

	if o.AzureRegion != nil {
		v := *o.AzureRegion

		if err = d.Set("azure_region", v); err != nil {
			return diag.Errorf("error reading azure_region: %v", err)
		}
	}

	if o.ClientId != nil {
		v := *o.ClientId

		if err = d.Set("client_id", v); err != nil {
			return diag.Errorf("error reading client_id: %v", err)
		}
	}

	if o.ClientSecret != nil {
		v := *o.ClientSecret

		if v == "ENC XXXX" {
		} else if err = d.Set("client_secret", v); err != nil {
			return diag.Errorf("error reading client_secret: %v", err)
		}
	}

	if o.CompartmentId != nil {
		v := *o.CompartmentId

		if err = d.Set("compartment_id", v); err != nil {
			return diag.Errorf("error reading compartment_id: %v", err)
		}
	}

	if o.ComputeGeneration != nil {
		v := *o.ComputeGeneration

		if err = d.Set("compute_generation", v); err != nil {
			return diag.Errorf("error reading compute_generation: %v", err)
		}
	}

	if o.Domain != nil {
		v := *o.Domain

		if err = d.Set("domain", v); err != nil {
			return diag.Errorf("error reading domain: %v", err)
		}
	}

	if o.ExternalAccountList != nil {
		if err = d.Set("external_account_list", flattenSystemSdnConnectorExternalAccountList(d, o.ExternalAccountList, "external_account_list", sort)); err != nil {
			return diag.Errorf("error reading external_account_list: %v", err)
		}
	}

	if o.ExternalIp != nil {
		if err = d.Set("external_ip", flattenSystemSdnConnectorExternalIp(d, o.ExternalIp, "external_ip", sort)); err != nil {
			return diag.Errorf("error reading external_ip: %v", err)
		}
	}

	if o.ForwardingRule != nil {
		if err = d.Set("forwarding_rule", flattenSystemSdnConnectorForwardingRule(d, o.ForwardingRule, "forwarding_rule", sort)); err != nil {
			return diag.Errorf("error reading forwarding_rule: %v", err)
		}
	}

	if o.GcpProject != nil {
		v := *o.GcpProject

		if err = d.Set("gcp_project", v); err != nil {
			return diag.Errorf("error reading gcp_project: %v", err)
		}
	}

	if o.GcpProjectList != nil {
		if err = d.Set("gcp_project_list", flattenSystemSdnConnectorGcpProjectList(d, o.GcpProjectList, "gcp_project_list", sort)); err != nil {
			return diag.Errorf("error reading gcp_project_list: %v", err)
		}
	}

	if o.GroupName != nil {
		v := *o.GroupName

		if err = d.Set("group_name", v); err != nil {
			return diag.Errorf("error reading group_name: %v", err)
		}
	}

	if o.HaStatus != nil {
		v := *o.HaStatus

		if err = d.Set("ha_status", v); err != nil {
			return diag.Errorf("error reading ha_status: %v", err)
		}
	}

	if o.IbmRegion != nil {
		v := *o.IbmRegion

		if err = d.Set("ibm_region", v); err != nil {
			return diag.Errorf("error reading ibm_region: %v", err)
		}
	}

	if o.LoginEndpoint != nil {
		v := *o.LoginEndpoint

		if err = d.Set("login_endpoint", v); err != nil {
			return diag.Errorf("error reading login_endpoint: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Nic != nil {
		if err = d.Set("nic", flattenSystemSdnConnectorNic(d, o.Nic, "nic", sort)); err != nil {
			return diag.Errorf("error reading nic: %v", err)
		}
	}

	if o.OciCert != nil {
		v := *o.OciCert

		if err = d.Set("oci_cert", v); err != nil {
			return diag.Errorf("error reading oci_cert: %v", err)
		}
	}

	if o.OciFingerprint != nil {
		v := *o.OciFingerprint

		if err = d.Set("oci_fingerprint", v); err != nil {
			return diag.Errorf("error reading oci_fingerprint: %v", err)
		}
	}

	if o.OciRegion != nil {
		v := *o.OciRegion

		if err = d.Set("oci_region", v); err != nil {
			return diag.Errorf("error reading oci_region: %v", err)
		}
	}

	if o.OciRegionType != nil {
		v := *o.OciRegionType

		if err = d.Set("oci_region_type", v); err != nil {
			return diag.Errorf("error reading oci_region_type: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.PrivateKey != nil {
		v := *o.PrivateKey

		if err = d.Set("private_key", v); err != nil {
			return diag.Errorf("error reading private_key: %v", err)
		}
	}

	if o.Region != nil {
		v := *o.Region

		if err = d.Set("region", v); err != nil {
			return diag.Errorf("error reading region: %v", err)
		}
	}

	if o.ResourceGroup != nil {
		v := *o.ResourceGroup

		if err = d.Set("resource_group", v); err != nil {
			return diag.Errorf("error reading resource_group: %v", err)
		}
	}

	if o.ResourceUrl != nil {
		v := *o.ResourceUrl

		if err = d.Set("resource_url", v); err != nil {
			return diag.Errorf("error reading resource_url: %v", err)
		}
	}

	if o.Route != nil {
		if err = d.Set("route", flattenSystemSdnConnectorRoute(d, o.Route, "route", sort)); err != nil {
			return diag.Errorf("error reading route: %v", err)
		}
	}

	if o.RouteTable != nil {
		if err = d.Set("route_table", flattenSystemSdnConnectorRouteTable(d, o.RouteTable, "route_table", sort)); err != nil {
			return diag.Errorf("error reading route_table: %v", err)
		}
	}

	if o.SecretKey != nil {
		v := *o.SecretKey

		if v == "ENC XXXX" {
		} else if err = d.Set("secret_key", v); err != nil {
			return diag.Errorf("error reading secret_key: %v", err)
		}
	}

	if o.SecretToken != nil {
		v := *o.SecretToken

		if err = d.Set("secret_token", v); err != nil {
			return diag.Errorf("error reading secret_token: %v", err)
		}
	}

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.ServerCaCert != nil {
		v := *o.ServerCaCert

		if err = d.Set("server_ca_cert", v); err != nil {
			return diag.Errorf("error reading server_ca_cert: %v", err)
		}
	}

	if o.ServerCert != nil {
		v := *o.ServerCert

		if err = d.Set("server_cert", v); err != nil {
			return diag.Errorf("error reading server_cert: %v", err)
		}
	}

	if o.ServerList != nil {
		if err = d.Set("server_list", flattenSystemSdnConnectorServerList(d, o.ServerList, "server_list", sort)); err != nil {
			return diag.Errorf("error reading server_list: %v", err)
		}
	}

	if o.ServerPort != nil {
		v := *o.ServerPort

		if err = d.Set("server_port", v); err != nil {
			return diag.Errorf("error reading server_port: %v", err)
		}
	}

	if o.ServiceAccount != nil {
		v := *o.ServiceAccount

		if err = d.Set("service_account", v); err != nil {
			return diag.Errorf("error reading service_account: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.SubscriptionId != nil {
		v := *o.SubscriptionId

		if err = d.Set("subscription_id", v); err != nil {
			return diag.Errorf("error reading subscription_id: %v", err)
		}
	}

	if o.TenantId != nil {
		v := *o.TenantId

		if err = d.Set("tenant_id", v); err != nil {
			return diag.Errorf("error reading tenant_id: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.UpdateInterval != nil {
		v := *o.UpdateInterval

		if err = d.Set("update_interval", v); err != nil {
			return diag.Errorf("error reading update_interval: %v", err)
		}
	}

	if o.UseMetadataIam != nil {
		v := *o.UseMetadataIam

		if err = d.Set("use_metadata_iam", v); err != nil {
			return diag.Errorf("error reading use_metadata_iam: %v", err)
		}
	}

	if o.UserId != nil {
		v := *o.UserId

		if err = d.Set("user_id", v); err != nil {
			return diag.Errorf("error reading user_id: %v", err)
		}
	}

	if o.Username != nil {
		v := *o.Username

		if err = d.Set("username", v); err != nil {
			return diag.Errorf("error reading username: %v", err)
		}
	}

	if o.VcenterPassword != nil {
		v := *o.VcenterPassword

		if v == "ENC XXXX" {
		} else if err = d.Set("vcenter_password", v); err != nil {
			return diag.Errorf("error reading vcenter_password: %v", err)
		}
	}

	if o.VcenterServer != nil {
		v := *o.VcenterServer

		if err = d.Set("vcenter_server", v); err != nil {
			return diag.Errorf("error reading vcenter_server: %v", err)
		}
	}

	if o.VcenterUsername != nil {
		v := *o.VcenterUsername

		if err = d.Set("vcenter_username", v); err != nil {
			return diag.Errorf("error reading vcenter_username: %v", err)
		}
	}

	if o.VerifyCertificate != nil {
		v := *o.VerifyCertificate

		if err = d.Set("verify_certificate", v); err != nil {
			return diag.Errorf("error reading verify_certificate: %v", err)
		}
	}

	if o.VpcId != nil {
		v := *o.VpcId

		if err = d.Set("vpc_id", v); err != nil {
			return diag.Errorf("error reading vpc_id: %v", err)
		}
	}

	return nil
}

func expandSystemSdnConnectorExternalAccountList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdnConnectorExternalAccountList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdnConnectorExternalAccountList

	for i := range l {
		tmp := models.SystemSdnConnectorExternalAccountList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.external_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExternalId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.region_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdnConnectorExternalAccountListRegionList(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdnConnectorExternalAccountListRegionList
			// 	}
			tmp.RegionList = v2

		}

		pre_append = fmt.Sprintf("%s.%d.role_arn", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RoleArn = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSdnConnectorExternalAccountListRegionList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdnConnectorExternalAccountListRegionList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdnConnectorExternalAccountListRegionList

	for i := range l {
		tmp := models.SystemSdnConnectorExternalAccountListRegionList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.region", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Region = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSdnConnectorExternalIp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdnConnectorExternalIp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdnConnectorExternalIp

	for i := range l {
		tmp := models.SystemSdnConnectorExternalIp{}
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

func expandSystemSdnConnectorForwardingRule(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdnConnectorForwardingRule, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdnConnectorForwardingRule

	for i := range l {
		tmp := models.SystemSdnConnectorForwardingRule{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.rule_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RuleName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Target = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSdnConnectorGcpProjectList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdnConnectorGcpProjectList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdnConnectorGcpProjectList

	for i := range l {
		tmp := models.SystemSdnConnectorGcpProjectList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.gcp_zone_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdnConnectorGcpProjectListGcpZoneList(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdnConnectorGcpProjectListGcpZoneList
			// 	}
			tmp.GcpZoneList = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSdnConnectorGcpProjectListGcpZoneList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdnConnectorGcpProjectListGcpZoneList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdnConnectorGcpProjectListGcpZoneList

	for i := range l {
		tmp := models.SystemSdnConnectorGcpProjectListGcpZoneList{}
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

func expandSystemSdnConnectorNic(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdnConnectorNic, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdnConnectorNic

	for i := range l {
		tmp := models.SystemSdnConnectorNic{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdnConnectorNicIp(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdnConnectorNicIp
			// 	}
			tmp.Ip = v2

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

func expandSystemSdnConnectorNicIp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdnConnectorNicIp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdnConnectorNicIp

	for i := range l {
		tmp := models.SystemSdnConnectorNicIp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.public_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PublicIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.resource_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ResourceGroup = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSdnConnectorRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdnConnectorRoute, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdnConnectorRoute

	for i := range l {
		tmp := models.SystemSdnConnectorRoute{}
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

func expandSystemSdnConnectorRouteTable(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdnConnectorRouteTable, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdnConnectorRouteTable

	for i := range l {
		tmp := models.SystemSdnConnectorRouteTable{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.resource_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ResourceGroup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdnConnectorRouteTableRoute(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdnConnectorRouteTableRoute
			// 	}
			tmp.Route = v2

		}

		pre_append = fmt.Sprintf("%s.%d.subscription_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SubscriptionId = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSdnConnectorRouteTableRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdnConnectorRouteTableRoute, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdnConnectorRouteTableRoute

	for i := range l {
		tmp := models.SystemSdnConnectorRouteTableRoute{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.next_hop", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NextHop = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSdnConnectorServerList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdnConnectorServerList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdnConnectorServerList

	for i := range l {
		tmp := models.SystemSdnConnectorServerList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemSdnConnector(d *schema.ResourceData, sv string) (*models.SystemSdnConnector, diag.Diagnostics) {
	obj := models.SystemSdnConnector{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("access_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("access_key", sv)
				diags = append(diags, e)
			}
			obj.AccessKey = &v2
		}
	}
	if v1, ok := d.GetOk("alt_resource_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("alt_resource_ip", sv)
				diags = append(diags, e)
			}
			obj.AltResourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("api_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("api_key", sv)
				diags = append(diags, e)
			}
			obj.ApiKey = &v2
		}
	}
	if v1, ok := d.GetOk("azure_region"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("azure_region", sv)
				diags = append(diags, e)
			}
			obj.AzureRegion = &v2
		}
	}
	if v1, ok := d.GetOk("client_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("client_id", sv)
				diags = append(diags, e)
			}
			obj.ClientId = &v2
		}
	}
	if v1, ok := d.GetOk("client_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("client_secret", sv)
				diags = append(diags, e)
			}
			obj.ClientSecret = &v2
		}
	}
	if v1, ok := d.GetOk("compartment_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("compartment_id", sv)
				diags = append(diags, e)
			}
			obj.CompartmentId = &v2
		}
	}
	if v1, ok := d.GetOk("compute_generation"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("compute_generation", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ComputeGeneration = &tmp
		}
	}
	if v1, ok := d.GetOk("domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("domain", sv)
				diags = append(diags, e)
			}
			obj.Domain = &v2
		}
	}
	if v, ok := d.GetOk("external_account_list"); ok {
		if !utils.CheckVer(sv, "v7.0.4", "") {
			e := utils.AttributeVersionWarning("external_account_list", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdnConnectorExternalAccountList(d, v, "external_account_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ExternalAccountList = t
		}
	} else if d.HasChange("external_account_list") {
		old, new := d.GetChange("external_account_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ExternalAccountList = &[]models.SystemSdnConnectorExternalAccountList{}
		}
	}
	if v, ok := d.GetOk("external_ip"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("external_ip", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdnConnectorExternalIp(d, v, "external_ip", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ExternalIp = t
		}
	} else if d.HasChange("external_ip") {
		old, new := d.GetChange("external_ip")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ExternalIp = &[]models.SystemSdnConnectorExternalIp{}
		}
	}
	if v, ok := d.GetOk("forwarding_rule"); ok {
		if !utils.CheckVer(sv, "v7.0.2", "") {
			e := utils.AttributeVersionWarning("forwarding_rule", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdnConnectorForwardingRule(d, v, "forwarding_rule", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ForwardingRule = t
		}
	} else if d.HasChange("forwarding_rule") {
		old, new := d.GetChange("forwarding_rule")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ForwardingRule = &[]models.SystemSdnConnectorForwardingRule{}
		}
	}
	if v1, ok := d.GetOk("gcp_project"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.7") {
				e := utils.AttributeVersionWarning("gcp_project", sv)
				diags = append(diags, e)
			}
			obj.GcpProject = &v2
		}
	}
	if v, ok := d.GetOk("gcp_project_list"); ok {
		if !utils.CheckVer(sv, "v6.4.7", "v7.0.0") {
			e := utils.AttributeVersionWarning("gcp_project_list", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdnConnectorGcpProjectList(d, v, "gcp_project_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.GcpProjectList = t
		}
	} else if d.HasChange("gcp_project_list") {
		old, new := d.GetChange("gcp_project_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.GcpProjectList = &[]models.SystemSdnConnectorGcpProjectList{}
		}
	}
	if v1, ok := d.GetOk("group_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_name", sv)
				diags = append(diags, e)
			}
			obj.GroupName = &v2
		}
	}
	if v1, ok := d.GetOk("ha_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ha_status", sv)
				diags = append(diags, e)
			}
			obj.HaStatus = &v2
		}
	}
	if v1, ok := d.GetOk("ibm_region"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("ibm_region", sv)
				diags = append(diags, e)
			}
			obj.IbmRegion = &v2
		}
	}
	if v1, ok := d.GetOk("login_endpoint"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("login_endpoint", sv)
				diags = append(diags, e)
			}
			obj.LoginEndpoint = &v2
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
	if v, ok := d.GetOk("nic"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("nic", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdnConnectorNic(d, v, "nic", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Nic = t
		}
	} else if d.HasChange("nic") {
		old, new := d.GetChange("nic")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Nic = &[]models.SystemSdnConnectorNic{}
		}
	}
	if v1, ok := d.GetOk("oci_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("oci_cert", sv)
				diags = append(diags, e)
			}
			obj.OciCert = &v2
		}
	}
	if v1, ok := d.GetOk("oci_fingerprint"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("oci_fingerprint", sv)
				diags = append(diags, e)
			}
			obj.OciFingerprint = &v2
		}
	}
	if v1, ok := d.GetOk("oci_region"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("oci_region", sv)
				diags = append(diags, e)
			}
			obj.OciRegion = &v2
		}
	}
	if v1, ok := d.GetOk("oci_region_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("oci_region_type", sv)
				diags = append(diags, e)
			}
			obj.OciRegionType = &v2
		}
	}
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("private_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("private_key", sv)
				diags = append(diags, e)
			}
			obj.PrivateKey = &v2
		}
	}
	if v1, ok := d.GetOk("region"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("region", sv)
				diags = append(diags, e)
			}
			obj.Region = &v2
		}
	}
	if v1, ok := d.GetOk("resource_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("resource_group", sv)
				diags = append(diags, e)
			}
			obj.ResourceGroup = &v2
		}
	}
	if v1, ok := d.GetOk("resource_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("resource_url", sv)
				diags = append(diags, e)
			}
			obj.ResourceUrl = &v2
		}
	}
	if v, ok := d.GetOk("route"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("route", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdnConnectorRoute(d, v, "route", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Route = t
		}
	} else if d.HasChange("route") {
		old, new := d.GetChange("route")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Route = &[]models.SystemSdnConnectorRoute{}
		}
	}
	if v, ok := d.GetOk("route_table"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("route_table", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdnConnectorRouteTable(d, v, "route_table", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.RouteTable = t
		}
	} else if d.HasChange("route_table") {
		old, new := d.GetChange("route_table")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.RouteTable = &[]models.SystemSdnConnectorRouteTable{}
		}
	}
	if v1, ok := d.GetOk("secret_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secret_key", sv)
				diags = append(diags, e)
			}
			obj.SecretKey = &v2
		}
	}
	if v1, ok := d.GetOk("secret_token"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secret_token", sv)
				diags = append(diags, e)
			}
			obj.SecretToken = &v2
		}
	}
	if v1, ok := d.GetOk("server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server", sv)
				diags = append(diags, e)
			}
			obj.Server = &v2
		}
	}
	if v1, ok := d.GetOk("server_ca_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("server_ca_cert", sv)
				diags = append(diags, e)
			}
			obj.ServerCaCert = &v2
		}
	}
	if v1, ok := d.GetOk("server_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("server_cert", sv)
				diags = append(diags, e)
			}
			obj.ServerCert = &v2
		}
	}
	if v, ok := d.GetOk("server_list"); ok {
		if !utils.CheckVer(sv, "v6.4.5", "") {
			e := utils.AttributeVersionWarning("server_list", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdnConnectorServerList(d, v, "server_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ServerList = t
		}
	} else if d.HasChange("server_list") {
		old, new := d.GetChange("server_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ServerList = &[]models.SystemSdnConnectorServerList{}
		}
	}
	if v1, ok := d.GetOk("server_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ServerPort = &tmp
		}
	}
	if v1, ok := d.GetOk("service_account"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("service_account", sv)
				diags = append(diags, e)
			}
			obj.ServiceAccount = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("subscription_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("subscription_id", sv)
				diags = append(diags, e)
			}
			obj.SubscriptionId = &v2
		}
	}
	if v1, ok := d.GetOk("tenant_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tenant_id", sv)
				diags = append(diags, e)
			}
			obj.TenantId = &v2
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("update_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UpdateInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("use_metadata_iam"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("use_metadata_iam", sv)
				diags = append(diags, e)
			}
			obj.UseMetadataIam = &v2
		}
	}
	if v1, ok := d.GetOk("user_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_id", sv)
				diags = append(diags, e)
			}
			obj.UserId = &v2
		}
	}
	if v1, ok := d.GetOk("username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("username", sv)
				diags = append(diags, e)
			}
			obj.Username = &v2
		}
	}
	if v1, ok := d.GetOk("vcenter_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("vcenter_password", sv)
				diags = append(diags, e)
			}
			obj.VcenterPassword = &v2
		}
	}
	if v1, ok := d.GetOk("vcenter_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("vcenter_server", sv)
				diags = append(diags, e)
			}
			obj.VcenterServer = &v2
		}
	}
	if v1, ok := d.GetOk("vcenter_username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("vcenter_username", sv)
				diags = append(diags, e)
			}
			obj.VcenterUsername = &v2
		}
	}
	if v1, ok := d.GetOk("verify_certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("verify_certificate", sv)
				diags = append(diags, e)
			}
			obj.VerifyCertificate = &v2
		}
	}
	if v1, ok := d.GetOk("vpc_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vpc_id", sv)
				diags = append(diags, e)
			}
			obj.VpcId = &v2
		}
	}
	return &obj, diags
}
