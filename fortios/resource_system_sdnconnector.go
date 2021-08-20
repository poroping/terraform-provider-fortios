// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Configure connection to SDN Connector.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSdnConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdnConnectorCreate,
		Read:   resourceSystemSdnConnectorRead,
		Update: resourceSystemSdnConnectorUpdate,
		Delete: resourceSystemSdnConnectorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"batchid": {
				Type:        schema.TypeInt,
				Description: "Associate with batch. From 6.4.x+. Currently a WIP and broken.",
				Optional:    true,
				Default:     0,
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
			"access_key": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "AWS / ACS access key ID.",
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
				ValidateFunc: fortiValidateEnum([]string{"global", "china", "germany", "usgov", "local"}),

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
			"gcp_project": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "GCP project name.",
				Optional:    true,
				Computed:    true,
			},
			"group_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Group name of computers.",
				Optional:    true,
				Computed:    true,
			},
			"ha_status": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable use for FortiGate HA service.",
				Optional:    true,
				Computed:    true,
			},
			"ibm_region": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"us-south", "us-east", "germany", "great-britain", "japan", "australia"}),

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
				ValidateFunc: fortiValidateEnum([]string{"commercial", "government"}),

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
				ValidateFunc: fortiValidateEnableDisable(),

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
				ValidateFunc: fortiValidateEnum([]string{"aci", "alicloud", "aws", "azure", "gcp", "nsx", "nuage", "oci", "openstack", "kubernetes", "vmware", "sepm", "aci-direct", "ibm", "nutanix"}),

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
				ValidateFunc: fortiValidateEnableDisable(),

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
				ValidateFunc: fortiValidateEnableDisable(),

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

func resourceSystemSdnConnectorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	allow_append := false

	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}

	urlparams["allow_append"] = []string{strconv.FormatBool(allow_append)}

	key := "name"
	mkey := ""
	if v, ok := d.GetOk(key); ok {
		if s, ok := v.(string); ok {
			mkey = s
		}
	}

	obj, err := getObjectSystemSdnConnector(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating SystemSdnConnector resource while getting object: %v", err)
	}

	if mkey == "" && allow_append {
		return fmt.Errorf("error creating SystemSdnConnector resource: %q must be set if \"allow_append\" is true", key)
	}

	o := make(map[string]interface{})
	if mkey != "" && allow_append {
		o, err = c.UpdateSystemSdnConnector(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.CreateSystemSdnConnector(obj, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error creating SystemSdnConnector resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSdnConnector")
	}

	return resourceSystemSdnConnectorRead(d, m)
}

func resourceSystemSdnConnectorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	obj, err := getObjectSystemSdnConnector(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating SystemSdnConnector resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSdnConnector(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating SystemSdnConnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSdnConnector")
	}

	return resourceSystemSdnConnectorRead(d, m)
}

func resourceSystemSdnConnectorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	err := c.DeleteSystemSdnConnector(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting SystemSdnConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdnConnectorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	o, err := c.ReadSystemSdnConnector(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading SystemSdnConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdnConnector(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading SystemSdnConnector resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdnConnectorAccessKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorAzureRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorClientId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorCompartmentId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorComputeGeneration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorExternalIp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemSdnConnectorExternalIpName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorExternalIpName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorGcpProject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorHaStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorIbmRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorLoginEndpoint(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorNic(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenSystemSdnConnectorNicIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemSdnConnectorNicName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorNicIp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemSdnConnectorNicIpName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "public_ip"
		if _, ok := i["public-ip"]; ok {

			tmp["public_ip"] = flattenSystemSdnConnectorNicIpPublicIp(i["public-ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := i["resource-group"]; ok {

			tmp["resource_group"] = flattenSystemSdnConnectorNicIpResourceGroup(i["resource-group"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorNicIpName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicIpPublicIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicIpResourceGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciFingerprint(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciRegionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorPrivateKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorResourceGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorResourceUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRoute(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemSdnConnectorRouteName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorRouteName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTable(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemSdnConnectorRouteTableName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := i["resource-group"]; ok {

			tmp["resource_group"] = flattenSystemSdnConnectorRouteTableResourceGroup(i["resource-group"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route"
		if _, ok := i["route"]; ok {

			tmp["route"] = flattenSystemSdnConnectorRouteTableRoute(i["route"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subscription_id"
		if _, ok := i["subscription-id"]; ok {

			tmp["subscription_id"] = flattenSystemSdnConnectorRouteTableSubscriptionId(i["subscription-id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorRouteTableName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableResourceGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableRoute(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemSdnConnectorRouteTableRouteName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop"
		if _, ok := i["next-hop"]; ok {

			tmp["next_hop"] = flattenSystemSdnConnectorRouteTableRouteNextHop(i["next-hop"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorRouteTableRouteName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableRouteNextHop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableSubscriptionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorSecretToken(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorServerList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenSystemSdnConnectorServerListIp(i["ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip", d)
	return result
}

func flattenSystemSdnConnectorServerListIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorServiceAccount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorSubscriptionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorTenantId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorUpdateInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorUseMetadataIam(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorUserId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorVcenterServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorVcenterUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorVerifyCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorVpcId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSdnConnector(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("access_key", flattenSystemSdnConnectorAccessKey(o["access-key"], d, "access_key", sv)); err != nil {
		if !fortiAPIPatch(o["access-key"]) {
			return fmt.Errorf("error reading access_key: %v", err)
		}
	}

	if err = d.Set("azure_region", flattenSystemSdnConnectorAzureRegion(o["azure-region"], d, "azure_region", sv)); err != nil {
		if !fortiAPIPatch(o["azure-region"]) {
			return fmt.Errorf("error reading azure_region: %v", err)
		}
	}

	if err = d.Set("client_id", flattenSystemSdnConnectorClientId(o["client-id"], d, "client_id", sv)); err != nil {
		if !fortiAPIPatch(o["client-id"]) {
			return fmt.Errorf("error reading client_id: %v", err)
		}
	}

	if err = d.Set("compartment_id", flattenSystemSdnConnectorCompartmentId(o["compartment-id"], d, "compartment_id", sv)); err != nil {
		if !fortiAPIPatch(o["compartment-id"]) {
			return fmt.Errorf("error reading compartment_id: %v", err)
		}
	}

	if err = d.Set("compute_generation", flattenSystemSdnConnectorComputeGeneration(o["compute-generation"], d, "compute_generation", sv)); err != nil {
		if !fortiAPIPatch(o["compute-generation"]) {
			return fmt.Errorf("error reading compute_generation: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemSdnConnectorDomain(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("error reading domain: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("external_ip", flattenSystemSdnConnectorExternalIp(o["external-ip"], d, "external_ip", sv)); err != nil {
			if !fortiAPIPatch(o["external-ip"]) {
				return fmt.Errorf("error reading external_ip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("external_ip"); ok {
			if err = d.Set("external_ip", flattenSystemSdnConnectorExternalIp(o["external-ip"], d, "external_ip", sv)); err != nil {
				if !fortiAPIPatch(o["external-ip"]) {
					return fmt.Errorf("error reading external_ip: %v", err)
				}
			}
		}
	}

	if err = d.Set("gcp_project", flattenSystemSdnConnectorGcpProject(o["gcp-project"], d, "gcp_project", sv)); err != nil {
		if !fortiAPIPatch(o["gcp-project"]) {
			return fmt.Errorf("error reading gcp_project: %v", err)
		}
	}

	if err = d.Set("group_name", flattenSystemSdnConnectorGroupName(o["group-name"], d, "group_name", sv)); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("error reading group_name: %v", err)
		}
	}

	if err = d.Set("ha_status", flattenSystemSdnConnectorHaStatus(o["ha-status"], d, "ha_status", sv)); err != nil {
		if !fortiAPIPatch(o["ha-status"]) {
			return fmt.Errorf("error reading ha_status: %v", err)
		}
	}

	if err = d.Set("ibm_region", flattenSystemSdnConnectorIbmRegion(o["ibm-region"], d, "ibm_region", sv)); err != nil {
		if !fortiAPIPatch(o["ibm-region"]) {
			return fmt.Errorf("error reading ibm_region: %v", err)
		}
	}

	if err = d.Set("login_endpoint", flattenSystemSdnConnectorLoginEndpoint(o["login-endpoint"], d, "login_endpoint", sv)); err != nil {
		if !fortiAPIPatch(o["login-endpoint"]) {
			return fmt.Errorf("error reading login_endpoint: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSdnConnectorName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nic", flattenSystemSdnConnectorNic(o["nic"], d, "nic", sv)); err != nil {
			if !fortiAPIPatch(o["nic"]) {
				return fmt.Errorf("error reading nic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nic"); ok {
			if err = d.Set("nic", flattenSystemSdnConnectorNic(o["nic"], d, "nic", sv)); err != nil {
				if !fortiAPIPatch(o["nic"]) {
					return fmt.Errorf("error reading nic: %v", err)
				}
			}
		}
	}

	if err = d.Set("oci_cert", flattenSystemSdnConnectorOciCert(o["oci-cert"], d, "oci_cert", sv)); err != nil {
		if !fortiAPIPatch(o["oci-cert"]) {
			return fmt.Errorf("error reading oci_cert: %v", err)
		}
	}

	if err = d.Set("oci_fingerprint", flattenSystemSdnConnectorOciFingerprint(o["oci-fingerprint"], d, "oci_fingerprint", sv)); err != nil {
		if !fortiAPIPatch(o["oci-fingerprint"]) {
			return fmt.Errorf("error reading oci_fingerprint: %v", err)
		}
	}

	if err = d.Set("oci_region", flattenSystemSdnConnectorOciRegion(o["oci-region"], d, "oci_region", sv)); err != nil {
		if !fortiAPIPatch(o["oci-region"]) {
			return fmt.Errorf("error reading oci_region: %v", err)
		}
	}

	if err = d.Set("oci_region_type", flattenSystemSdnConnectorOciRegionType(o["oci-region-type"], d, "oci_region_type", sv)); err != nil {
		if !fortiAPIPatch(o["oci-region-type"]) {
			return fmt.Errorf("error reading oci_region_type: %v", err)
		}
	}

	if err = d.Set("private_key", flattenSystemSdnConnectorPrivateKey(o["private-key"], d, "private_key", sv)); err != nil {
		if !fortiAPIPatch(o["private-key"]) {
			return fmt.Errorf("error reading private_key: %v", err)
		}
	}

	if err = d.Set("region", flattenSystemSdnConnectorRegion(o["region"], d, "region", sv)); err != nil {
		if !fortiAPIPatch(o["region"]) {
			return fmt.Errorf("error reading region: %v", err)
		}
	}

	if err = d.Set("resource_group", flattenSystemSdnConnectorResourceGroup(o["resource-group"], d, "resource_group", sv)); err != nil {
		if !fortiAPIPatch(o["resource-group"]) {
			return fmt.Errorf("error reading resource_group: %v", err)
		}
	}

	if err = d.Set("resource_url", flattenSystemSdnConnectorResourceUrl(o["resource-url"], d, "resource_url", sv)); err != nil {
		if !fortiAPIPatch(o["resource-url"]) {
			return fmt.Errorf("error reading resource_url: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("route", flattenSystemSdnConnectorRoute(o["route"], d, "route", sv)); err != nil {
			if !fortiAPIPatch(o["route"]) {
				return fmt.Errorf("error reading route: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("route"); ok {
			if err = d.Set("route", flattenSystemSdnConnectorRoute(o["route"], d, "route", sv)); err != nil {
				if !fortiAPIPatch(o["route"]) {
					return fmt.Errorf("error reading route: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("route_table", flattenSystemSdnConnectorRouteTable(o["route-table"], d, "route_table", sv)); err != nil {
			if !fortiAPIPatch(o["route-table"]) {
				return fmt.Errorf("error reading route_table: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("route_table"); ok {
			if err = d.Set("route_table", flattenSystemSdnConnectorRouteTable(o["route-table"], d, "route_table", sv)); err != nil {
				if !fortiAPIPatch(o["route-table"]) {
					return fmt.Errorf("error reading route_table: %v", err)
				}
			}
		}
	}

	if err = d.Set("secret_token", flattenSystemSdnConnectorSecretToken(o["secret-token"], d, "secret_token", sv)); err != nil {
		if !fortiAPIPatch(o["secret-token"]) {
			return fmt.Errorf("error reading secret_token: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemSdnConnectorServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("error reading server: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_list", flattenSystemSdnConnectorServerList(o["server-list"], d, "server_list", sv)); err != nil {
			if !fortiAPIPatch(o["server-list"]) {
				return fmt.Errorf("error reading server_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_list"); ok {
			if err = d.Set("server_list", flattenSystemSdnConnectorServerList(o["server-list"], d, "server_list", sv)); err != nil {
				if !fortiAPIPatch(o["server-list"]) {
					return fmt.Errorf("error reading server_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_port", flattenSystemSdnConnectorServerPort(o["server-port"], d, "server_port", sv)); err != nil {
		if !fortiAPIPatch(o["server-port"]) {
			return fmt.Errorf("error reading server_port: %v", err)
		}
	}

	if err = d.Set("service_account", flattenSystemSdnConnectorServiceAccount(o["service-account"], d, "service_account", sv)); err != nil {
		if !fortiAPIPatch(o["service-account"]) {
			return fmt.Errorf("error reading service_account: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSdnConnectorStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("subscription_id", flattenSystemSdnConnectorSubscriptionId(o["subscription-id"], d, "subscription_id", sv)); err != nil {
		if !fortiAPIPatch(o["subscription-id"]) {
			return fmt.Errorf("error reading subscription_id: %v", err)
		}
	}

	if err = d.Set("tenant_id", flattenSystemSdnConnectorTenantId(o["tenant-id"], d, "tenant_id", sv)); err != nil {
		if !fortiAPIPatch(o["tenant-id"]) {
			return fmt.Errorf("error reading tenant_id: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemSdnConnectorType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("error reading type: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenSystemSdnConnectorUpdateInterval(o["update-interval"], d, "update_interval", sv)); err != nil {
		if !fortiAPIPatch(o["update-interval"]) {
			return fmt.Errorf("error reading update_interval: %v", err)
		}
	}

	if err = d.Set("use_metadata_iam", flattenSystemSdnConnectorUseMetadataIam(o["use-metadata-iam"], d, "use_metadata_iam", sv)); err != nil {
		if !fortiAPIPatch(o["use-metadata-iam"]) {
			return fmt.Errorf("error reading use_metadata_iam: %v", err)
		}
	}

	if err = d.Set("user_id", flattenSystemSdnConnectorUserId(o["user-id"], d, "user_id", sv)); err != nil {
		if !fortiAPIPatch(o["user-id"]) {
			return fmt.Errorf("error reading user_id: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemSdnConnectorUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("error reading username: %v", err)
		}
	}

	if err = d.Set("vcenter_server", flattenSystemSdnConnectorVcenterServer(o["vcenter-server"], d, "vcenter_server", sv)); err != nil {
		if !fortiAPIPatch(o["vcenter-server"]) {
			return fmt.Errorf("error reading vcenter_server: %v", err)
		}
	}

	if err = d.Set("vcenter_username", flattenSystemSdnConnectorVcenterUsername(o["vcenter-username"], d, "vcenter_username", sv)); err != nil {
		if !fortiAPIPatch(o["vcenter-username"]) {
			return fmt.Errorf("error reading vcenter_username: %v", err)
		}
	}

	if err = d.Set("verify_certificate", flattenSystemSdnConnectorVerifyCertificate(o["verify-certificate"], d, "verify_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["verify-certificate"]) {
			return fmt.Errorf("error reading verify_certificate: %v", err)
		}
	}

	if err = d.Set("vpc_id", flattenSystemSdnConnectorVpcId(o["vpc-id"], d, "vpc_id", sv)); err != nil {
		if !fortiAPIPatch(o["vpc-id"]) {
			return fmt.Errorf("error reading vpc_id: %v", err)
		}
	}

	return nil
}

func expandSystemSdnConnectorAccessKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorApiKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorAzureRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorClientId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorClientSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorCompartmentId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorComputeGeneration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorExternalIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemSdnConnectorExternalIpName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorExternalIpName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorGcpProject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorHaStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorIbmRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorLoginEndpoint(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandSystemSdnConnectorNicIp(d, i["ip"], pre_append, sv)
		} else {
			tmp["ip"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemSdnConnectorNicName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorNicIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemSdnConnectorNicIpName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "public_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["public-ip"], _ = expandSystemSdnConnectorNicIpPublicIp(d, i["public_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["resource-group"], _ = expandSystemSdnConnectorNicIpResourceGroup(d, i["resource_group"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorNicIpName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicIpPublicIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicIpResourceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciFingerprint(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciRegionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorPrivateKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorResourceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorResourceUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemSdnConnectorRouteName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorRouteName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemSdnConnectorRouteTableName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["resource-group"], _ = expandSystemSdnConnectorRouteTableResourceGroup(d, i["resource_group"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route"], _ = expandSystemSdnConnectorRouteTableRoute(d, i["route"], pre_append, sv)
		} else {
			tmp["route"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subscription_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["subscription-id"], _ = expandSystemSdnConnectorRouteTableSubscriptionId(d, i["subscription_id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorRouteTableName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableResourceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemSdnConnectorRouteTableRouteName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["next-hop"], _ = expandSystemSdnConnectorRouteTableRouteNextHop(d, i["next_hop"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorRouteTableRouteName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableRouteNextHop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableSubscriptionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSecretKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSecretToken(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServerList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandSystemSdnConnectorServerListIp(d, i["ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorServerListIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServiceAccount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSubscriptionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorTenantId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUpdateInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUseMetadataIam(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUserId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVcenterPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVcenterServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVcenterUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVerifyCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVpcId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdnConnector(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_key"); ok {
		t, err := expandSystemSdnConnectorAccessKey(d, v, "access_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-key"] = t
		}
	}

	if v, ok := d.GetOk("api_key"); ok {
		t, err := expandSystemSdnConnectorApiKey(d, v, "api_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-key"] = t
		}
	}

	if v, ok := d.GetOk("azure_region"); ok {
		t, err := expandSystemSdnConnectorAzureRegion(d, v, "azure_region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-region"] = t
		}
	}

	if v, ok := d.GetOk("client_id"); ok {
		t, err := expandSystemSdnConnectorClientId(d, v, "client_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-id"] = t
		}
	}

	if v, ok := d.GetOk("client_secret"); ok {
		t, err := expandSystemSdnConnectorClientSecret(d, v, "client_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-secret"] = t
		}
	}

	if v, ok := d.GetOk("compartment_id"); ok {
		t, err := expandSystemSdnConnectorCompartmentId(d, v, "compartment_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compartment-id"] = t
		}
	}

	if v, ok := d.GetOk("compute_generation"); ok {
		t, err := expandSystemSdnConnectorComputeGeneration(d, v, "compute_generation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compute-generation"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandSystemSdnConnectorDomain(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("external_ip"); ok {
		t, err := expandSystemSdnConnectorExternalIp(d, v, "external_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-ip"] = t
		}
	} else if d.HasChange("external_ip") {
		old, new := d.GetChange("external_ip")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["external-ip"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("gcp_project"); ok {
		t, err := expandSystemSdnConnectorGcpProject(d, v, "gcp_project", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-project"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok {
		t, err := expandSystemSdnConnectorGroupName(d, v, "group_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("ha_status"); ok {
		t, err := expandSystemSdnConnectorHaStatus(d, v, "ha_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-status"] = t
		}
	}

	if v, ok := d.GetOk("ibm_region"); ok {
		t, err := expandSystemSdnConnectorIbmRegion(d, v, "ibm_region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibm-region"] = t
		}
	}

	if v, ok := d.GetOk("login_endpoint"); ok {
		t, err := expandSystemSdnConnectorLoginEndpoint(d, v, "login_endpoint", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-endpoint"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSdnConnectorName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nic"); ok {
		t, err := expandSystemSdnConnectorNic(d, v, "nic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nic"] = t
		}
	} else if d.HasChange("nic") {
		old, new := d.GetChange("nic")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["nic"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("oci_cert"); ok {
		t, err := expandSystemSdnConnectorOciCert(d, v, "oci_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-cert"] = t
		}
	}

	if v, ok := d.GetOk("oci_fingerprint"); ok {
		t, err := expandSystemSdnConnectorOciFingerprint(d, v, "oci_fingerprint", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-fingerprint"] = t
		}
	}

	if v, ok := d.GetOk("oci_region"); ok {
		t, err := expandSystemSdnConnectorOciRegion(d, v, "oci_region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-region"] = t
		}
	}

	if v, ok := d.GetOk("oci_region_type"); ok {
		t, err := expandSystemSdnConnectorOciRegionType(d, v, "oci_region_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-region-type"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemSdnConnectorPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok {
		t, err := expandSystemSdnConnectorPrivateKey(d, v, "private_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	}

	if v, ok := d.GetOk("region"); ok {
		t, err := expandSystemSdnConnectorRegion(d, v, "region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region"] = t
		}
	}

	if v, ok := d.GetOk("resource_group"); ok {
		t, err := expandSystemSdnConnectorResourceGroup(d, v, "resource_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-group"] = t
		}
	}

	if v, ok := d.GetOk("resource_url"); ok {
		t, err := expandSystemSdnConnectorResourceUrl(d, v, "resource_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-url"] = t
		}
	}

	if v, ok := d.GetOk("route"); ok {
		t, err := expandSystemSdnConnectorRoute(d, v, "route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route"] = t
		}
	} else if d.HasChange("route") {
		old, new := d.GetChange("route")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["route"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("route_table"); ok {
		t, err := expandSystemSdnConnectorRouteTable(d, v, "route_table", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-table"] = t
		}
	} else if d.HasChange("route_table") {
		old, new := d.GetChange("route_table")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["route-table"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("secret_key"); ok {
		t, err := expandSystemSdnConnectorSecretKey(d, v, "secret_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret-key"] = t
		}
	}

	if v, ok := d.GetOk("secret_token"); ok {
		t, err := expandSystemSdnConnectorSecretToken(d, v, "secret_token", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret-token"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandSystemSdnConnectorServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_list"); ok {
		t, err := expandSystemSdnConnectorServerList(d, v, "server_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-list"] = t
		}
	} else if d.HasChange("server_list") {
		old, new := d.GetChange("server_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["server-list"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("server_port"); ok {
		t, err := expandSystemSdnConnectorServerPort(d, v, "server_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-port"] = t
		}
	}

	if v, ok := d.GetOk("service_account"); ok {
		t, err := expandSystemSdnConnectorServiceAccount(d, v, "service_account", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-account"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemSdnConnectorStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("subscription_id"); ok {
		t, err := expandSystemSdnConnectorSubscriptionId(d, v, "subscription_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subscription-id"] = t
		}
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		t, err := expandSystemSdnConnectorTenantId(d, v, "tenant_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemSdnConnectorType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok {
		t, err := expandSystemSdnConnectorUpdateInterval(d, v, "update_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	if v, ok := d.GetOk("use_metadata_iam"); ok {
		t, err := expandSystemSdnConnectorUseMetadataIam(d, v, "use_metadata_iam", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-metadata-iam"] = t
		}
	}

	if v, ok := d.GetOk("user_id"); ok {
		t, err := expandSystemSdnConnectorUserId(d, v, "user_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-id"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandSystemSdnConnectorUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("vcenter_password"); ok {
		t, err := expandSystemSdnConnectorVcenterPassword(d, v, "vcenter_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcenter-password"] = t
		}
	}

	if v, ok := d.GetOk("vcenter_server"); ok {
		t, err := expandSystemSdnConnectorVcenterServer(d, v, "vcenter_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcenter-server"] = t
		}
	}

	if v, ok := d.GetOk("vcenter_username"); ok {
		t, err := expandSystemSdnConnectorVcenterUsername(d, v, "vcenter_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcenter-username"] = t
		}
	}

	if v, ok := d.GetOk("verify_certificate"); ok {
		t, err := expandSystemSdnConnectorVerifyCertificate(d, v, "verify_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-certificate"] = t
		}
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		t, err := expandSystemSdnConnectorVpcId(d, v, "vpc_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpc-id"] = t
		}
	}

	return &obj, nil
}
