// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiClient Enterprise Management Server (EMS) entries.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceEndpointControlFctems() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiClient Enterprise Management Server (EMS) entries.",

		CreateContext: resourceEndpointControlFctemsCreate,
		ReadContext:   resourceEndpointControlFctemsRead,
		UpdateContext: resourceEndpointControlFctemsUpdate,
		DeleteContext: resourceEndpointControlFctemsDelete,

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
			"admin_password": {
				Type: schema.TypeString,

				Description: "FortiClient EMS admin password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"admin_username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),

				Description: "FortiClient EMS admin username.",
				Optional:    true,
				Computed:    true,
			},
			"call_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 180),

				Description: "FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"capabilities": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "List of EMS capabilities.",
				Optional:         true,
				Computed:         true,
			},
			"certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "FortiClient EMS certificate.",
				Optional:    true,
				Computed:    true,
			},
			"cloud_server_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"production", "alpha", "beta"}, false),

				Description: "Cloud server type.",
				Optional:    true,
				Computed:    true,
			},
			"fortinetone_cloud_authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account.",
				Optional:    true,
				Computed:    true,
			},
			"https_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "FortiClient Enterprise Management Server (EMS) name.",
				ForceNew:    true,
				Required:    true,
			},
			"preserve_ssl_session": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting.",
				Optional:    true,
				Computed:    true,
			},
			"pull_avatars": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable pulling avatars from EMS.",
				Optional:    true,
				Computed:    true,
			},
			"pull_malware_hash": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable pulling FortiClient malware hash from EMS.",
				Optional:    true,
				Computed:    true,
			},
			"pull_sysinfo": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable pulling SysInfo from EMS.",
				Optional:    true,
				Computed:    true,
			},
			"pull_tags": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable pulling FortiClient user tags from EMS.",
				Optional:    true,
				Computed:    true,
			},
			"pull_vulnerabilities": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable pulling vulnerabilities from EMS.",
				Optional:    true,
				Computed:    true,
			},
			"serial_number": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),

				Description: "FortiClient EMS Serial Number.",
				Optional:    true,
				Computed:    true,
			},
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "FortiClient EMS FQDN or IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "REST API call source IP.",
				Optional:    true,
				Computed:    true,
			},
			"status_check_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 180),

				Description: "FortiClient EMS call timeout in seconds (1 - 120 seconds, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"websocket_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceEndpointControlFctemsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating EndpointControlFctems resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectEndpointControlFctems(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateEndpointControlFctems(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EndpointControlFctems")
	}

	return resourceEndpointControlFctemsRead(ctx, d, meta)
}

func resourceEndpointControlFctemsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectEndpointControlFctems(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateEndpointControlFctems(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating EndpointControlFctems resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EndpointControlFctems")
	}

	return resourceEndpointControlFctemsRead(ctx, d, meta)
}

func resourceEndpointControlFctemsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteEndpointControlFctems(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting EndpointControlFctems resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEndpointControlFctemsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadEndpointControlFctems(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading EndpointControlFctems resource: %v", err)
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

	diags := refreshObjectEndpointControlFctems(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectEndpointControlFctems(d *schema.ResourceData, o *models.EndpointControlFctems, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AdminPassword != nil {
		v := *o.AdminPassword

		if v == "ENC XXXX" {
		} else if err = d.Set("admin_password", v); err != nil {
			return diag.Errorf("error reading admin_password: %v", err)
		}
	}

	if o.AdminUsername != nil {
		v := *o.AdminUsername

		if err = d.Set("admin_username", v); err != nil {
			return diag.Errorf("error reading admin_username: %v", err)
		}
	}

	if o.CallTimeout != nil {
		v := *o.CallTimeout

		if err = d.Set("call_timeout", v); err != nil {
			return diag.Errorf("error reading call_timeout: %v", err)
		}
	}

	if o.Capabilities != nil {
		v := *o.Capabilities

		if err = d.Set("capabilities", v); err != nil {
			return diag.Errorf("error reading capabilities: %v", err)
		}
	}

	if o.Certificate != nil {
		v := *o.Certificate

		if err = d.Set("certificate", v); err != nil {
			return diag.Errorf("error reading certificate: %v", err)
		}
	}

	if o.CloudServerType != nil {
		v := *o.CloudServerType

		if err = d.Set("cloud_server_type", v); err != nil {
			return diag.Errorf("error reading cloud_server_type: %v", err)
		}
	}

	if o.FortinetoneCloudAuthentication != nil {
		v := *o.FortinetoneCloudAuthentication

		if err = d.Set("fortinetone_cloud_authentication", v); err != nil {
			return diag.Errorf("error reading fortinetone_cloud_authentication: %v", err)
		}
	}

	if o.HttpsPort != nil {
		v := *o.HttpsPort

		if err = d.Set("https_port", v); err != nil {
			return diag.Errorf("error reading https_port: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PreserveSslSession != nil {
		v := *o.PreserveSslSession

		if err = d.Set("preserve_ssl_session", v); err != nil {
			return diag.Errorf("error reading preserve_ssl_session: %v", err)
		}
	}

	if o.PullAvatars != nil {
		v := *o.PullAvatars

		if err = d.Set("pull_avatars", v); err != nil {
			return diag.Errorf("error reading pull_avatars: %v", err)
		}
	}

	if o.PullMalwareHash != nil {
		v := *o.PullMalwareHash

		if err = d.Set("pull_malware_hash", v); err != nil {
			return diag.Errorf("error reading pull_malware_hash: %v", err)
		}
	}

	if o.PullSysinfo != nil {
		v := *o.PullSysinfo

		if err = d.Set("pull_sysinfo", v); err != nil {
			return diag.Errorf("error reading pull_sysinfo: %v", err)
		}
	}

	if o.PullTags != nil {
		v := *o.PullTags

		if err = d.Set("pull_tags", v); err != nil {
			return diag.Errorf("error reading pull_tags: %v", err)
		}
	}

	if o.PullVulnerabilities != nil {
		v := *o.PullVulnerabilities

		if err = d.Set("pull_vulnerabilities", v); err != nil {
			return diag.Errorf("error reading pull_vulnerabilities: %v", err)
		}
	}

	if o.SerialNumber != nil {
		v := *o.SerialNumber

		if err = d.Set("serial_number", v); err != nil {
			return diag.Errorf("error reading serial_number: %v", err)
		}
	}

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.StatusCheckInterval != nil {
		v := *o.StatusCheckInterval

		if err = d.Set("status_check_interval", v); err != nil {
			return diag.Errorf("error reading status_check_interval: %v", err)
		}
	}

	if o.WebsocketOverride != nil {
		v := *o.WebsocketOverride

		if err = d.Set("websocket_override", v); err != nil {
			return diag.Errorf("error reading websocket_override: %v", err)
		}
	}

	return nil
}

func getObjectEndpointControlFctems(d *schema.ResourceData, sv string) (*models.EndpointControlFctems, diag.Diagnostics) {
	obj := models.EndpointControlFctems{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("admin_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("admin_password", sv)
				diags = append(diags, e)
			}
			obj.AdminPassword = &v2
		}
	}
	if v1, ok := d.GetOk("admin_username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("admin_username", sv)
				diags = append(diags, e)
			}
			obj.AdminUsername = &v2
		}
	}
	if v1, ok := d.GetOk("call_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("call_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CallTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("capabilities"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("capabilities", sv)
				diags = append(diags, e)
			}
			obj.Capabilities = &v2
		}
	}
	if v1, ok := d.GetOk("certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "v7.0.0") {
				e := utils.AttributeVersionWarning("certificate", sv)
				diags = append(diags, e)
			}
			obj.Certificate = &v2
		}
	}
	if v1, ok := d.GetOk("cloud_server_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("cloud_server_type", sv)
				diags = append(diags, e)
			}
			obj.CloudServerType = &v2
		}
	}
	if v1, ok := d.GetOk("fortinetone_cloud_authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortinetone_cloud_authentication", sv)
				diags = append(diags, e)
			}
			obj.FortinetoneCloudAuthentication = &v2
		}
	}
	if v1, ok := d.GetOk("https_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("https_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HttpsPort = &tmp
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
	if v1, ok := d.GetOk("preserve_ssl_session"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("preserve_ssl_session", sv)
				diags = append(diags, e)
			}
			obj.PreserveSslSession = &v2
		}
	}
	if v1, ok := d.GetOk("pull_avatars"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("pull_avatars", sv)
				diags = append(diags, e)
			}
			obj.PullAvatars = &v2
		}
	}
	if v1, ok := d.GetOk("pull_malware_hash"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("pull_malware_hash", sv)
				diags = append(diags, e)
			}
			obj.PullMalwareHash = &v2
		}
	}
	if v1, ok := d.GetOk("pull_sysinfo"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("pull_sysinfo", sv)
				diags = append(diags, e)
			}
			obj.PullSysinfo = &v2
		}
	}
	if v1, ok := d.GetOk("pull_tags"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("pull_tags", sv)
				diags = append(diags, e)
			}
			obj.PullTags = &v2
		}
	}
	if v1, ok := d.GetOk("pull_vulnerabilities"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("pull_vulnerabilities", sv)
				diags = append(diags, e)
			}
			obj.PullVulnerabilities = &v2
		}
	}
	if v1, ok := d.GetOk("serial_number"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("serial_number", sv)
				diags = append(diags, e)
			}
			obj.SerialNumber = &v2
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
	if v1, ok := d.GetOk("source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip", sv)
				diags = append(diags, e)
			}
			obj.SourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("status_check_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.5", "v7.0.0") {
				e := utils.AttributeVersionWarning("status_check_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.StatusCheckInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("websocket_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("websocket_override", sv)
				diags = append(diags, e)
			}
			obj.WebsocketOverride = &v2
		}
	}
	return &obj, diags
}
