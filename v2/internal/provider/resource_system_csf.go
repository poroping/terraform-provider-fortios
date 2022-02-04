// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.

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

func resourceSystemCsf() *schema.Resource {
	return &schema.Resource{
		Description: "Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.",

		CreateContext: resourceSystemCsfCreate,
		ReadContext:   resourceSystemCsfRead,
		UpdateContext: resourceSystemCsfUpdate,
		DeleteContext: resourceSystemCsfDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"accept_auth_by_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Accept connections with unknown certificates and ask admin for approval.",
				Optional:    true,
				Computed:    true,
			},
			"authorization_request_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"serial", "certificate"}, false),

				Description: "Authorization request type.",
				Optional:    true,
				Computed:    true,
			},
			"certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Certificate.",
				Optional:    true,
				Computed:    true,
			},
			"configuration_sync": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "local"}, false),

				Description: "Configuration sync mode.",
				Optional:    true,
				Computed:    true,
			},
			"downstream_access": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable downstream device access to this device's configuration and data.",
				Optional:    true,
				Computed:    true,
			},
			"downstream_accprofile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Default access profile for requests from downstream devices.",
				Optional:    true,
				Computed:    true,
			},
			"fabric_connector": {
				Type:        schema.TypeList,
				Description: "Fabric connector configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accprofile": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Override access profile.",
							Optional:    true,
							Computed:    true,
						},
						"configuration_write_access": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable downstream device write access to configuration.",
							Optional:    true,
							Computed:    true,
						},
						"serial": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 19),

							Description: "Serial.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fabric_device": {
				Type:        schema.TypeList,
				Description: "Fabric device configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_token": {
							Type: schema.TypeString,

							Description: "Device access token.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"device_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Device IP.",
							Optional:    true,
							Computed:    true,
						},
						"https_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "HTTPS port for fabric device.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Device name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fabric_object_unification": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "local"}, false),

				Description: "Fabric CMDB Object Unification.",
				Optional:    true,
				Computed:    true,
			},
			"fabric_workers": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4),

				Description: "Number of worker processes for Security Fabric daemon.",
				Optional:    true,
				Computed:    true,
			},
			"forticloud_account_enforcement": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Fabric FortiCloud account unification.",
				Optional:    true,
				Computed:    true,
			},
			"group_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.",
				Optional:    true,
				Computed:    true,
			},
			"group_password": {
				Type: schema.TypeString,

				Description: "Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"log_unification": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable broadcast of discovery messages for log unification.",
				Optional:    true,
				Computed:    true,
			},
			"management_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.",
				Optional:    true,
				Computed:    true,
			},
			"management_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Overriding port for management connection (Overrides admin port).",
				Optional:    true,
				Computed:    true,
			},
			"saml_configuration_sync": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "local"}, false),

				Description: "SAML setting configuration synchronization.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Security Fabric.",
				Optional:    true,
				Computed:    true,
			},
			"trusted_list": {
				Type:        schema.TypeList,
				Description: "Pre-authorized and blocked security fabric nodes.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"accept", "deny"}, false),

							Description: "Security fabric authorization action.",
							Optional:    true,
							Computed:    true,
						},
						"authorization_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"serial", "certificate"}, false),

							Description: "Authorization type.",
							Optional:    true,
							Computed:    true,
						},
						"certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32767),

							Description: "Certificate.",
							Optional:    true,
							Computed:    true,
						},
						"downstream_authorization": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Trust authorizations by this node's administrator.",
							Optional:    true,
							Computed:    true,
						},
						"ha_members": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.StringLenBetween(0, 19),
							DiffSuppressFunc: suppressors.DiffMultiStringEqual,
							Description:      "HA members.",
							Optional:         true,
							Computed:         true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Name.",
							Optional:    true,
							Computed:    true,
						},
						"serial": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 19),

							Description: "Serial.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"upstream": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.",
				Optional:    true,
				Computed:    true,
			},
			"upstream_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of the FortiGate upstream from this FortiGate in the Security Fabric.",
				Optional:    true,
				Computed:    true,
			},
			"upstream_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemCsfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemCsf(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemCsf(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemCsf")
	}

	return resourceSystemCsfRead(ctx, d, meta)
}

func resourceSystemCsfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemCsf(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemCsf(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemCsf resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemCsf")
	}

	return resourceSystemCsfRead(ctx, d, meta)
}

func resourceSystemCsfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemCsf(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemCsf(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemCsf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCsfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemCsf(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemCsf resource: %v", err)
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

	diags := refreshObjectSystemCsf(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemCsfFabricConnector(v *[]models.SystemCsfFabricConnector, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Accprofile; tmp != nil {
				v["accprofile"] = *tmp
			}

			if tmp := cfg.ConfigurationWriteAccess; tmp != nil {
				v["configuration_write_access"] = *tmp
			}

			if tmp := cfg.Serial; tmp != nil {
				v["serial"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "serial")
	}

	return flat
}

func flattenSystemCsfFabricDevice(v *[]models.SystemCsfFabricDevice, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AccessToken; tmp != nil {
				v["access_token"] = *tmp
			}

			if tmp := cfg.DeviceIp; tmp != nil {
				v["device_ip"] = *tmp
			}

			if tmp := cfg.HttpsPort; tmp != nil {
				v["https_port"] = *tmp
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

func flattenSystemCsfTrustedList(v *[]models.SystemCsfTrustedList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.AuthorizationType; tmp != nil {
				v["authorization_type"] = *tmp
			}

			if tmp := cfg.Certificate; tmp != nil {
				v["certificate"] = *tmp
			}

			if tmp := cfg.DownstreamAuthorization; tmp != nil {
				v["downstream_authorization"] = *tmp
			}

			if tmp := cfg.HaMembers; tmp != nil {
				v["ha_members"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Serial; tmp != nil {
				v["serial"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectSystemCsf(d *schema.ResourceData, o *models.SystemCsf, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AcceptAuthByCert != nil {
		v := *o.AcceptAuthByCert

		if err = d.Set("accept_auth_by_cert", v); err != nil {
			return diag.Errorf("error reading accept_auth_by_cert: %v", err)
		}
	}

	if o.AuthorizationRequestType != nil {
		v := *o.AuthorizationRequestType

		if err = d.Set("authorization_request_type", v); err != nil {
			return diag.Errorf("error reading authorization_request_type: %v", err)
		}
	}

	if o.Certificate != nil {
		v := *o.Certificate

		if err = d.Set("certificate", v); err != nil {
			return diag.Errorf("error reading certificate: %v", err)
		}
	}

	if o.ConfigurationSync != nil {
		v := *o.ConfigurationSync

		if err = d.Set("configuration_sync", v); err != nil {
			return diag.Errorf("error reading configuration_sync: %v", err)
		}
	}

	if o.DownstreamAccess != nil {
		v := *o.DownstreamAccess

		if err = d.Set("downstream_access", v); err != nil {
			return diag.Errorf("error reading downstream_access: %v", err)
		}
	}

	if o.DownstreamAccprofile != nil {
		v := *o.DownstreamAccprofile

		if err = d.Set("downstream_accprofile", v); err != nil {
			return diag.Errorf("error reading downstream_accprofile: %v", err)
		}
	}

	if o.FabricConnector != nil {
		if err = d.Set("fabric_connector", flattenSystemCsfFabricConnector(o.FabricConnector, sort)); err != nil {
			return diag.Errorf("error reading fabric_connector: %v", err)
		}
	}

	if o.FabricDevice != nil {
		if err = d.Set("fabric_device", flattenSystemCsfFabricDevice(o.FabricDevice, sort)); err != nil {
			return diag.Errorf("error reading fabric_device: %v", err)
		}
	}

	if o.FabricObjectUnification != nil {
		v := *o.FabricObjectUnification

		if err = d.Set("fabric_object_unification", v); err != nil {
			return diag.Errorf("error reading fabric_object_unification: %v", err)
		}
	}

	if o.FabricWorkers != nil {
		v := *o.FabricWorkers

		if err = d.Set("fabric_workers", v); err != nil {
			return diag.Errorf("error reading fabric_workers: %v", err)
		}
	}

	if o.ForticloudAccountEnforcement != nil {
		v := *o.ForticloudAccountEnforcement

		if err = d.Set("forticloud_account_enforcement", v); err != nil {
			return diag.Errorf("error reading forticloud_account_enforcement: %v", err)
		}
	}

	if o.GroupName != nil {
		v := *o.GroupName

		if err = d.Set("group_name", v); err != nil {
			return diag.Errorf("error reading group_name: %v", err)
		}
	}

	if o.GroupPassword != nil {
		v := *o.GroupPassword

		if v == "ENC XXXX" {
		} else if err = d.Set("group_password", v); err != nil {
			return diag.Errorf("error reading group_password: %v", err)
		}
	}

	if o.LogUnification != nil {
		v := *o.LogUnification

		if err = d.Set("log_unification", v); err != nil {
			return diag.Errorf("error reading log_unification: %v", err)
		}
	}

	if o.ManagementIp != nil {
		v := *o.ManagementIp

		if err = d.Set("management_ip", v); err != nil {
			return diag.Errorf("error reading management_ip: %v", err)
		}
	}

	if o.ManagementPort != nil {
		v := *o.ManagementPort

		if err = d.Set("management_port", v); err != nil {
			return diag.Errorf("error reading management_port: %v", err)
		}
	}

	if o.SamlConfigurationSync != nil {
		v := *o.SamlConfigurationSync

		if err = d.Set("saml_configuration_sync", v); err != nil {
			return diag.Errorf("error reading saml_configuration_sync: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TrustedList != nil {
		if err = d.Set("trusted_list", flattenSystemCsfTrustedList(o.TrustedList, sort)); err != nil {
			return diag.Errorf("error reading trusted_list: %v", err)
		}
	}

	if o.Upstream != nil {
		v := *o.Upstream

		if err = d.Set("upstream", v); err != nil {
			return diag.Errorf("error reading upstream: %v", err)
		}
	}

	if o.UpstreamIp != nil {
		v := *o.UpstreamIp

		if err = d.Set("upstream_ip", v); err != nil {
			return diag.Errorf("error reading upstream_ip: %v", err)
		}
	}

	if o.UpstreamPort != nil {
		v := *o.UpstreamPort

		if err = d.Set("upstream_port", v); err != nil {
			return diag.Errorf("error reading upstream_port: %v", err)
		}
	}

	return nil
}

func expandSystemCsfFabricConnector(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemCsfFabricConnector, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemCsfFabricConnector

	for i := range l {
		tmp := models.SystemCsfFabricConnector{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.accprofile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Accprofile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.configuration_write_access", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ConfigurationWriteAccess = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.serial", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Serial = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemCsfFabricDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemCsfFabricDevice, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemCsfFabricDevice

	for i := range l {
		tmp := models.SystemCsfFabricDevice{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.access_token", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AccessToken = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.device_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DeviceIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.https_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HttpsPort = &v3
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

func expandSystemCsfTrustedList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemCsfTrustedList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemCsfTrustedList

	for i := range l {
		tmp := models.SystemCsfTrustedList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.authorization_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthorizationType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Certificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.downstream_authorization", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DownstreamAuthorization = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ha_members", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HaMembers = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.serial", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Serial = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemCsf(d *schema.ResourceData, sv string) (*models.SystemCsf, diag.Diagnostics) {
	obj := models.SystemCsf{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("accept_auth_by_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("accept_auth_by_cert", sv)
				diags = append(diags, e)
			}
			obj.AcceptAuthByCert = &v2
		}
	}
	if v1, ok := d.GetOk("authorization_request_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("authorization_request_type", sv)
				diags = append(diags, e)
			}
			obj.AuthorizationRequestType = &v2
		}
	}
	if v1, ok := d.GetOk("certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("certificate", sv)
				diags = append(diags, e)
			}
			obj.Certificate = &v2
		}
	}
	if v1, ok := d.GetOk("configuration_sync"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("configuration_sync", sv)
				diags = append(diags, e)
			}
			obj.ConfigurationSync = &v2
		}
	}
	if v1, ok := d.GetOk("downstream_access"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("downstream_access", sv)
				diags = append(diags, e)
			}
			obj.DownstreamAccess = &v2
		}
	}
	if v1, ok := d.GetOk("downstream_accprofile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("downstream_accprofile", sv)
				diags = append(diags, e)
			}
			obj.DownstreamAccprofile = &v2
		}
	}
	if v, ok := d.GetOk("fabric_connector"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("fabric_connector", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemCsfFabricConnector(d, v, "fabric_connector", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FabricConnector = t
		}
	} else if d.HasChange("fabric_connector") {
		old, new := d.GetChange("fabric_connector")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FabricConnector = &[]models.SystemCsfFabricConnector{}
		}
	}
	if v, ok := d.GetOk("fabric_device"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("fabric_device", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemCsfFabricDevice(d, v, "fabric_device", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FabricDevice = t
		}
	} else if d.HasChange("fabric_device") {
		old, new := d.GetChange("fabric_device")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FabricDevice = &[]models.SystemCsfFabricDevice{}
		}
	}
	if v1, ok := d.GetOk("fabric_object_unification"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("fabric_object_unification", sv)
				diags = append(diags, e)
			}
			obj.FabricObjectUnification = &v2
		}
	}
	if v1, ok := d.GetOk("fabric_workers"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.5", "") {
				e := utils.AttributeVersionWarning("fabric_workers", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FabricWorkers = &tmp
		}
	}
	if v1, ok := d.GetOk("forticloud_account_enforcement"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("forticloud_account_enforcement", sv)
				diags = append(diags, e)
			}
			obj.ForticloudAccountEnforcement = &v2
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
	if v1, ok := d.GetOk("group_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_password", sv)
				diags = append(diags, e)
			}
			obj.GroupPassword = &v2
		}
	}
	if v1, ok := d.GetOk("log_unification"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("log_unification", sv)
				diags = append(diags, e)
			}
			obj.LogUnification = &v2
		}
	}
	if v1, ok := d.GetOk("management_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("management_ip", sv)
				diags = append(diags, e)
			}
			obj.ManagementIp = &v2
		}
	}
	if v1, ok := d.GetOk("management_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("management_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ManagementPort = &tmp
		}
	}
	if v1, ok := d.GetOk("saml_configuration_sync"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("saml_configuration_sync", sv)
				diags = append(diags, e)
			}
			obj.SamlConfigurationSync = &v2
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
	if v, ok := d.GetOk("trusted_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("trusted_list", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemCsfTrustedList(d, v, "trusted_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.TrustedList = t
		}
	} else if d.HasChange("trusted_list") {
		old, new := d.GetChange("trusted_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.TrustedList = &[]models.SystemCsfTrustedList{}
		}
	}
	if v1, ok := d.GetOk("upstream"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("upstream", sv)
				diags = append(diags, e)
			}
			obj.Upstream = &v2
		}
	}
	if v1, ok := d.GetOk("upstream_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.2") {
				e := utils.AttributeVersionWarning("upstream_ip", sv)
				diags = append(diags, e)
			}
			obj.UpstreamIp = &v2
		}
	}
	if v1, ok := d.GetOk("upstream_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("upstream_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UpstreamPort = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemCsf(d *schema.ResourceData, sv string) (*models.SystemCsf, diag.Diagnostics) {
	obj := models.SystemCsf{}
	diags := diag.Diagnostics{}

	obj.FabricConnector = &[]models.SystemCsfFabricConnector{}
	obj.FabricDevice = &[]models.SystemCsfFabricDevice{}
	obj.TrustedList = &[]models.SystemCsfTrustedList{}

	return &obj, diags
}
