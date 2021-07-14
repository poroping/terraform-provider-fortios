// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure connection to SDN Connector.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemSdnConnector() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemSdnConnectorRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"vcenter_server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_password": {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"access_key": {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"secret_key": {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"login_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_secret": {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"azure_region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nic": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"public_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_group": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"route_table": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subscription_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"route": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"next_hop": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_region_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_cert": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_ip": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"route": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"use_metadata_iam": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gcp_project": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_account": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_passwd": {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"private_key": {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"secret_token": {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"api_key": {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"compute_generation": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ibm_region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemSdnConnectorRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing SystemSdnConnector: type error")
	}

	o, err := c.ReadSystemSdnConnector(mkey, vdomparam, make(map[string][]string), 0)
	if err != nil {
		return fmt.Errorf("error describing SystemSdnConnector: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemSdnConnector(d, o)
	if err != nil {
		return fmt.Errorf("error describing SystemSdnConnector from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemSdnConnectorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorHaStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorVcenterServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorVcenterUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorVcenterPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorAccessKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorSecretKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorVpcId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorTenantId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorSubscriptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorLoginEndpoint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorResourceUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorClientId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorClientSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorResourceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorAzureRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorNic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorNicName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemSdnConnectorNicIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorNicName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorNicIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorNicIpName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "public_ip"
		if _, ok := i["public-ip"]; ok {
			tmp["public_ip"] = dataSourceFlattenSystemSdnConnectorNicIpPublicIp(i["public-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := i["resource-group"]; ok {
			tmp["resource_group"] = dataSourceFlattenSystemSdnConnectorNicIpResourceGroup(i["resource-group"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorNicIpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorNicIpPublicIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorNicIpResourceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRouteTable(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorRouteTableName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subscription_id"
		if _, ok := i["subscription-id"]; ok {
			tmp["subscription_id"] = dataSourceFlattenSystemSdnConnectorRouteTableSubscriptionId(i["subscription-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := i["resource-group"]; ok {
			tmp["resource_group"] = dataSourceFlattenSystemSdnConnectorRouteTableResourceGroup(i["resource-group"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route"
		if _, ok := i["route"]; ok {
			tmp["route"] = dataSourceFlattenSystemSdnConnectorRouteTableRoute(i["route"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorRouteTableName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRouteTableSubscriptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRouteTableResourceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRouteTableRoute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorRouteTableRouteName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop"
		if _, ok := i["next-hop"]; ok {
			tmp["next_hop"] = dataSourceFlattenSystemSdnConnectorRouteTableRouteNextHop(i["next-hop"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorRouteTableRouteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRouteTableRouteNextHop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorCompartmentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciRegionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciFingerprint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorExternalIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorExternalIpName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorExternalIpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRoute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorRouteName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorRouteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorUseMetadataIam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorGcpProject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorServiceAccount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorKeyPasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorSecretToken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorApiKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorComputeGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorIbmRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemSdnConnector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemSdnConnectorName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemSdnConnectorStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemSdnConnectorType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("error reading type: %v", err)
		}
	}

	if err = d.Set("ha_status", dataSourceFlattenSystemSdnConnectorHaStatus(o["ha-status"], d, "ha_status")); err != nil {
		if !fortiAPIPatch(o["ha-status"]) {
			return fmt.Errorf("error reading ha_status: %v", err)
		}
	}

	if err = d.Set("server", dataSourceFlattenSystemSdnConnectorServer(o["server"], d, "server")); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("error reading server: %v", err)
		}
	}

	if err = d.Set("server_port", dataSourceFlattenSystemSdnConnectorServerPort(o["server-port"], d, "server_port")); err != nil {
		if !fortiAPIPatch(o["server-port"]) {
			return fmt.Errorf("error reading server_port: %v", err)
		}
	}

	if err = d.Set("username", dataSourceFlattenSystemSdnConnectorUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("error reading username: %v", err)
		}
	}

	if err = d.Set("vcenter_server", dataSourceFlattenSystemSdnConnectorVcenterServer(o["vcenter-server"], d, "vcenter_server")); err != nil {
		if !fortiAPIPatch(o["vcenter-server"]) {
			return fmt.Errorf("error reading vcenter_server: %v", err)
		}
	}

	if err = d.Set("vcenter_username", dataSourceFlattenSystemSdnConnectorVcenterUsername(o["vcenter-username"], d, "vcenter_username")); err != nil {
		if !fortiAPIPatch(o["vcenter-username"]) {
			return fmt.Errorf("error reading vcenter_username: %v", err)
		}
	}

	if err = d.Set("region", dataSourceFlattenSystemSdnConnectorRegion(o["region"], d, "region")); err != nil {
		if !fortiAPIPatch(o["region"]) {
			return fmt.Errorf("error reading region: %v", err)
		}
	}

	if err = d.Set("vpc_id", dataSourceFlattenSystemSdnConnectorVpcId(o["vpc-id"], d, "vpc_id")); err != nil {
		if !fortiAPIPatch(o["vpc-id"]) {
			return fmt.Errorf("error reading vpc_id: %v", err)
		}
	}

	if err = d.Set("tenant_id", dataSourceFlattenSystemSdnConnectorTenantId(o["tenant-id"], d, "tenant_id")); err != nil {
		if !fortiAPIPatch(o["tenant-id"]) {
			return fmt.Errorf("error reading tenant_id: %v", err)
		}
	}

	if err = d.Set("subscription_id", dataSourceFlattenSystemSdnConnectorSubscriptionId(o["subscription-id"], d, "subscription_id")); err != nil {
		if !fortiAPIPatch(o["subscription-id"]) {
			return fmt.Errorf("error reading subscription_id: %v", err)
		}
	}

	if err = d.Set("login_endpoint", dataSourceFlattenSystemSdnConnectorLoginEndpoint(o["login-endpoint"], d, "login_endpoint")); err != nil {
		if !fortiAPIPatch(o["login-endpoint"]) {
			return fmt.Errorf("error reading login_endpoint: %v", err)
		}
	}

	if err = d.Set("resource_url", dataSourceFlattenSystemSdnConnectorResourceUrl(o["resource-url"], d, "resource_url")); err != nil {
		if !fortiAPIPatch(o["resource-url"]) {
			return fmt.Errorf("error reading resource_url: %v", err)
		}
	}

	if err = d.Set("client_id", dataSourceFlattenSystemSdnConnectorClientId(o["client-id"], d, "client_id")); err != nil {
		if !fortiAPIPatch(o["client-id"]) {
			return fmt.Errorf("error reading client_id: %v", err)
		}
	}

	if err = d.Set("resource_group", dataSourceFlattenSystemSdnConnectorResourceGroup(o["resource-group"], d, "resource_group")); err != nil {
		if !fortiAPIPatch(o["resource-group"]) {
			return fmt.Errorf("error reading resource_group: %v", err)
		}
	}

	if err = d.Set("azure_region", dataSourceFlattenSystemSdnConnectorAzureRegion(o["azure-region"], d, "azure_region")); err != nil {
		if !fortiAPIPatch(o["azure-region"]) {
			return fmt.Errorf("error reading azure_region: %v", err)
		}
	}

	if err = d.Set("nic", dataSourceFlattenSystemSdnConnectorNic(o["nic"], d, "nic")); err != nil {
		if !fortiAPIPatch(o["nic"]) {
			return fmt.Errorf("error reading nic: %v", err)
		}
	}

	if err = d.Set("route_table", dataSourceFlattenSystemSdnConnectorRouteTable(o["route-table"], d, "route_table")); err != nil {
		if !fortiAPIPatch(o["route-table"]) {
			return fmt.Errorf("error reading route_table: %v", err)
		}
	}

	if err = d.Set("user_id", dataSourceFlattenSystemSdnConnectorUserId(o["user-id"], d, "user_id")); err != nil {
		if !fortiAPIPatch(o["user-id"]) {
			return fmt.Errorf("error reading user_id: %v", err)
		}
	}

	if err = d.Set("compartment_id", dataSourceFlattenSystemSdnConnectorCompartmentId(o["compartment-id"], d, "compartment_id")); err != nil {
		if !fortiAPIPatch(o["compartment-id"]) {
			return fmt.Errorf("error reading compartment_id: %v", err)
		}
	}

	if err = d.Set("oci_region", dataSourceFlattenSystemSdnConnectorOciRegion(o["oci-region"], d, "oci_region")); err != nil {
		if !fortiAPIPatch(o["oci-region"]) {
			return fmt.Errorf("error reading oci_region: %v", err)
		}
	}

	if err = d.Set("oci_region_type", dataSourceFlattenSystemSdnConnectorOciRegionType(o["oci-region-type"], d, "oci_region_type")); err != nil {
		if !fortiAPIPatch(o["oci-region-type"]) {
			return fmt.Errorf("error reading oci_region_type: %v", err)
		}
	}

	if err = d.Set("oci_cert", dataSourceFlattenSystemSdnConnectorOciCert(o["oci-cert"], d, "oci_cert")); err != nil {
		if !fortiAPIPatch(o["oci-cert"]) {
			return fmt.Errorf("error reading oci_cert: %v", err)
		}
	}

	if err = d.Set("oci_fingerprint", dataSourceFlattenSystemSdnConnectorOciFingerprint(o["oci-fingerprint"], d, "oci_fingerprint")); err != nil {
		if !fortiAPIPatch(o["oci-fingerprint"]) {
			return fmt.Errorf("error reading oci_fingerprint: %v", err)
		}
	}

	if err = d.Set("external_ip", dataSourceFlattenSystemSdnConnectorExternalIp(o["external-ip"], d, "external_ip")); err != nil {
		if !fortiAPIPatch(o["external-ip"]) {
			return fmt.Errorf("error reading external_ip: %v", err)
		}
	}

	if err = d.Set("route", dataSourceFlattenSystemSdnConnectorRoute(o["route"], d, "route")); err != nil {
		if !fortiAPIPatch(o["route"]) {
			return fmt.Errorf("error reading route: %v", err)
		}
	}

	if err = d.Set("use_metadata_iam", dataSourceFlattenSystemSdnConnectorUseMetadataIam(o["use-metadata-iam"], d, "use_metadata_iam")); err != nil {
		if !fortiAPIPatch(o["use-metadata-iam"]) {
			return fmt.Errorf("error reading use_metadata_iam: %v", err)
		}
	}

	if err = d.Set("gcp_project", dataSourceFlattenSystemSdnConnectorGcpProject(o["gcp-project"], d, "gcp_project")); err != nil {
		if !fortiAPIPatch(o["gcp-project"]) {
			return fmt.Errorf("error reading gcp_project: %v", err)
		}
	}

	if err = d.Set("service_account", dataSourceFlattenSystemSdnConnectorServiceAccount(o["service-account"], d, "service_account")); err != nil {
		if !fortiAPIPatch(o["service-account"]) {
			return fmt.Errorf("error reading service_account: %v", err)
		}
	}

	if err = d.Set("domain", dataSourceFlattenSystemSdnConnectorDomain(o["domain"], d, "domain")); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("error reading domain: %v", err)
		}
	}

	if err = d.Set("group_name", dataSourceFlattenSystemSdnConnectorGroupName(o["group-name"], d, "group_name")); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("error reading group_name: %v", err)
		}
	}

	if err = d.Set("compute_generation", dataSourceFlattenSystemSdnConnectorComputeGeneration(o["compute-generation"], d, "compute_generation")); err != nil {
		if !fortiAPIPatch(o["compute-generation"]) {
			return fmt.Errorf("error reading compute_generation: %v", err)
		}
	}

	if err = d.Set("ibm_region", dataSourceFlattenSystemSdnConnectorIbmRegion(o["ibm-region"], d, "ibm_region")); err != nil {
		if !fortiAPIPatch(o["ibm-region"]) {
			return fmt.Errorf("error reading ibm_region: %v", err)
		}
	}

	if err = d.Set("update_interval", dataSourceFlattenSystemSdnConnectorUpdateInterval(o["update-interval"], d, "update_interval")); err != nil {
		if !fortiAPIPatch(o["update-interval"]) {
			return fmt.Errorf("error reading update_interval: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemSdnConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
