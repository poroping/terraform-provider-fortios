// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiGuard services.

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

func resourceSystemFortiguard() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiGuard services.",

		CreateContext: resourceSystemFortiguardCreate,
		ReadContext:   resourceSystemFortiguardRead,
		UpdateContext: resourceSystemFortiguardUpdate,
		DeleteContext: resourceSystemFortiguardDelete,

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
			"antispam_cache": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance.",
				Optional:    true,
				Computed:    true,
			},
			"antispam_cache_mpercent": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),

				Description: "Maximum percentage of FortiGate memory the antispam cache is allowed to use (1 - 15).",
				Optional:    true,
				Computed:    true,
			},
			"antispam_cache_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 86400),

				Description: "Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.",
				Optional:    true,
				Computed:    true,
			},
			"antispam_expiration": {
				Type: schema.TypeInt,

				Description: "Expiration date of the FortiGuard antispam contract.",
				Optional:    true,
				Computed:    true,
			},
			"antispam_force_off": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable turning off the FortiGuard antispam service.",
				Optional:    true,
				Computed:    true,
			},
			"antispam_license": {
				Type: schema.TypeInt,

				Description: "Interval of time between license checks for the FortiGuard antispam contract.",
				Optional:    true,
				Computed:    true,
			},
			"antispam_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),

				Description: "Antispam query time out (1 - 30 sec, default = 7).",
				Optional:    true,
				Computed:    true,
			},
			"anycast_sdns_server_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of the FortiGuard anycast DNS rating server.",
				Optional:    true,
				Computed:    true,
			},
			"anycast_sdns_server_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port to connect to on the FortiGuard anycast DNS rating server.",
				Optional:    true,
				Computed:    true,
			},
			"auto_join_forticloud": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Automatically connect to and login to FortiCloud.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_server_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of the FortiDDNS server.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_server_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 address of the FortiDDNS server.",
				Optional:         true,
				Computed:         true,
			},
			"ddns_server_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port used to communicate with FortiDDNS servers.",
				Optional:    true,
				Computed:    true,
			},
			"fortiguard_anycast": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of FortiGuard's Anycast network.",
				Optional:    true,
				Computed:    true,
			},
			"fortiguard_anycast_source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fortinet", "aws", "debug"}, false),

				Description: "Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Specify outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"interface_select_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

				Description: "Specify how to select outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"load_balance_servers": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 266),

				Description: "Number of servers to alternate between as first FortiGuard option.",
				Optional:    true,
				Computed:    true,
			},
			"outbreak_prevention_cache": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiGuard Virus Outbreak Prevention cache.",
				Optional:    true,
				Computed:    true,
			},
			"outbreak_prevention_cache_mpercent": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),

				Description: "Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).",
				Optional:    true,
				Computed:    true,
			},
			"outbreak_prevention_cache_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 86400),

				Description: "Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).",
				Optional:    true,
				Computed:    true,
			},
			"outbreak_prevention_expiration": {
				Type: schema.TypeInt,

				Description: "Expiration date of FortiGuard Virus Outbreak Prevention contract.",
				Optional:    true,
				Computed:    true,
			},
			"outbreak_prevention_force_off": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Turn off FortiGuard Virus Outbreak Prevention service.",
				Optional:    true,
				Computed:    true,
			},
			"outbreak_prevention_license": {
				Type: schema.TypeInt,

				Description: "Interval of time between license checks for FortiGuard Virus Outbreak Prevention contract.",
				Optional:    true,
				Computed:    true,
			},
			"outbreak_prevention_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),

				Description: "FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).",
				Optional:    true,
				Computed:    true,
			},
			"persistent_connection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of persistent connection to receive update notification from FortiGuard.",
				Optional:    true,
				Computed:    true,
			},
			"port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"8888", "53", "80", "443"}, false),

				Description: "Port used to communicate with the FortiGuard servers.",
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"udp", "http", "https"}, false),

				Description: "Protocol used to communicate with the FortiGuard servers.",
				Optional:    true,
				Computed:    true,
			},
			"proxy_password": {
				Type: schema.TypeString,

				Description: "Proxy user password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"proxy_server_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of the proxy server.",
				Optional:    true,
				Computed:    true,
			},
			"proxy_server_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port used to communicate with the proxy server.",
				Optional:    true,
				Computed:    true,
			},
			"proxy_username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "Proxy user name.",
				Optional:    true,
				Computed:    true,
			},
			"sandbox_region": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Cloud sandbox region.",
				Optional:    true,
				Computed:    true,
			},
			"sdns_options": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Customization options for the FortiGuard DNS service.",
				Optional:         true,
				Computed:         true,
			},
			"sdns_server_ip": {
				Type: schema.TypeString,

				Description: "IP address of the FortiGuard DNS rating server.",
				Optional:    true,
				Computed:    true,
			},
			"sdns_server_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port to connect to on the FortiGuard DNS rating server.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IPv4 address used to communicate with FortiGuard.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Source IPv6 address used to communicate with FortiGuard.",
				Optional:         true,
				Computed:         true,
			},
			"update_build_proxy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable proxy dictionary rebuild.",
				Optional:    true,
				Computed:    true,
			},
			"update_extdb": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable external resource update.",
				Optional:    true,
				Computed:    true,
			},
			"update_ffdb": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Internet Service Database update.",
				Optional:    true,
				Computed:    true,
			},
			"update_server_location": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"automatic", "usa", "eu"}, false),

				Description: "Location from which to receive FortiGuard updates.",
				Optional:    true,
				Computed:    true,
			},
			"update_uwdb": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowlist update.",
				Optional:    true,
				Computed:    true,
			},
			"videofilter_expiration": {
				Type: schema.TypeInt,

				Description: "Expiration date of the FortiGuard video filter contract.",
				Optional:    true,
				Computed:    true,
			},
			"videofilter_license": {
				Type: schema.TypeInt,

				Description: "Interval of time between license checks for the FortiGuard video filter contract.",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_cache": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiGuard web filter caching.",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_cache_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 86400),

				Description: "Time-to-live for web filter cache entries in seconds (300 - 86400).",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_expiration": {
				Type: schema.TypeInt,

				Description: "Expiration date of the FortiGuard web filter contract.",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_force_off": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable turning off the FortiGuard web filtering service.",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_license": {
				Type: schema.TypeInt,

				Description: "Interval of time between license checks for the FortiGuard web filter contract.",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),

				Description: "Web filter query time out (1 - 30 sec, default = 7).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemFortiguardCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemFortiguard(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemFortiguard(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemFortiguard")
	}

	return resourceSystemFortiguardRead(ctx, d, meta)
}

func resourceSystemFortiguardUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemFortiguard(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemFortiguard(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemFortiguard resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemFortiguard")
	}

	return resourceSystemFortiguardRead(ctx, d, meta)
}

func resourceSystemFortiguardDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemFortiguard(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemFortiguard(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortiguardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemFortiguard(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemFortiguard resource: %v", err)
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

	diags := refreshObjectSystemFortiguard(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemFortiguard(d *schema.ResourceData, o *models.SystemFortiguard, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AntispamCache != nil {
		v := *o.AntispamCache

		if err = d.Set("antispam_cache", v); err != nil {
			return diag.Errorf("error reading antispam_cache: %v", err)
		}
	}

	if o.AntispamCacheMpercent != nil {
		v := *o.AntispamCacheMpercent

		if err = d.Set("antispam_cache_mpercent", v); err != nil {
			return diag.Errorf("error reading antispam_cache_mpercent: %v", err)
		}
	}

	if o.AntispamCacheTtl != nil {
		v := *o.AntispamCacheTtl

		if err = d.Set("antispam_cache_ttl", v); err != nil {
			return diag.Errorf("error reading antispam_cache_ttl: %v", err)
		}
	}

	if o.AntispamExpiration != nil {
		v := *o.AntispamExpiration

		if err = d.Set("antispam_expiration", v); err != nil {
			return diag.Errorf("error reading antispam_expiration: %v", err)
		}
	}

	if o.AntispamForceOff != nil {
		v := *o.AntispamForceOff

		if err = d.Set("antispam_force_off", v); err != nil {
			return diag.Errorf("error reading antispam_force_off: %v", err)
		}
	}

	if o.AntispamLicense != nil {
		v := *o.AntispamLicense

		if err = d.Set("antispam_license", v); err != nil {
			return diag.Errorf("error reading antispam_license: %v", err)
		}
	}

	if o.AntispamTimeout != nil {
		v := *o.AntispamTimeout

		if err = d.Set("antispam_timeout", v); err != nil {
			return diag.Errorf("error reading antispam_timeout: %v", err)
		}
	}

	if o.AnycastSdnsServerIp != nil {
		v := *o.AnycastSdnsServerIp

		if err = d.Set("anycast_sdns_server_ip", v); err != nil {
			return diag.Errorf("error reading anycast_sdns_server_ip: %v", err)
		}
	}

	if o.AnycastSdnsServerPort != nil {
		v := *o.AnycastSdnsServerPort

		if err = d.Set("anycast_sdns_server_port", v); err != nil {
			return diag.Errorf("error reading anycast_sdns_server_port: %v", err)
		}
	}

	if o.AutoJoinForticloud != nil {
		v := *o.AutoJoinForticloud

		if err = d.Set("auto_join_forticloud", v); err != nil {
			return diag.Errorf("error reading auto_join_forticloud: %v", err)
		}
	}

	if o.DdnsServerIp != nil {
		v := *o.DdnsServerIp

		if err = d.Set("ddns_server_ip", v); err != nil {
			return diag.Errorf("error reading ddns_server_ip: %v", err)
		}
	}

	if o.DdnsServerIp6 != nil {
		v := *o.DdnsServerIp6

		if err = d.Set("ddns_server_ip6", v); err != nil {
			return diag.Errorf("error reading ddns_server_ip6: %v", err)
		}
	}

	if o.DdnsServerPort != nil {
		v := *o.DdnsServerPort

		if err = d.Set("ddns_server_port", v); err != nil {
			return diag.Errorf("error reading ddns_server_port: %v", err)
		}
	}

	if o.FortiguardAnycast != nil {
		v := *o.FortiguardAnycast

		if err = d.Set("fortiguard_anycast", v); err != nil {
			return diag.Errorf("error reading fortiguard_anycast: %v", err)
		}
	}

	if o.FortiguardAnycastSource != nil {
		v := *o.FortiguardAnycastSource

		if err = d.Set("fortiguard_anycast_source", v); err != nil {
			return diag.Errorf("error reading fortiguard_anycast_source: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.InterfaceSelectMethod != nil {
		v := *o.InterfaceSelectMethod

		if err = d.Set("interface_select_method", v); err != nil {
			return diag.Errorf("error reading interface_select_method: %v", err)
		}
	}

	if o.LoadBalanceServers != nil {
		v := *o.LoadBalanceServers

		if err = d.Set("load_balance_servers", v); err != nil {
			return diag.Errorf("error reading load_balance_servers: %v", err)
		}
	}

	if o.OutbreakPreventionCache != nil {
		v := *o.OutbreakPreventionCache

		if err = d.Set("outbreak_prevention_cache", v); err != nil {
			return diag.Errorf("error reading outbreak_prevention_cache: %v", err)
		}
	}

	if o.OutbreakPreventionCacheMpercent != nil {
		v := *o.OutbreakPreventionCacheMpercent

		if err = d.Set("outbreak_prevention_cache_mpercent", v); err != nil {
			return diag.Errorf("error reading outbreak_prevention_cache_mpercent: %v", err)
		}
	}

	if o.OutbreakPreventionCacheTtl != nil {
		v := *o.OutbreakPreventionCacheTtl

		if err = d.Set("outbreak_prevention_cache_ttl", v); err != nil {
			return diag.Errorf("error reading outbreak_prevention_cache_ttl: %v", err)
		}
	}

	if o.OutbreakPreventionExpiration != nil {
		v := *o.OutbreakPreventionExpiration

		if err = d.Set("outbreak_prevention_expiration", v); err != nil {
			return diag.Errorf("error reading outbreak_prevention_expiration: %v", err)
		}
	}

	if o.OutbreakPreventionForceOff != nil {
		v := *o.OutbreakPreventionForceOff

		if err = d.Set("outbreak_prevention_force_off", v); err != nil {
			return diag.Errorf("error reading outbreak_prevention_force_off: %v", err)
		}
	}

	if o.OutbreakPreventionLicense != nil {
		v := *o.OutbreakPreventionLicense

		if err = d.Set("outbreak_prevention_license", v); err != nil {
			return diag.Errorf("error reading outbreak_prevention_license: %v", err)
		}
	}

	if o.OutbreakPreventionTimeout != nil {
		v := *o.OutbreakPreventionTimeout

		if err = d.Set("outbreak_prevention_timeout", v); err != nil {
			return diag.Errorf("error reading outbreak_prevention_timeout: %v", err)
		}
	}

	if o.PersistentConnection != nil {
		v := *o.PersistentConnection

		if err = d.Set("persistent_connection", v); err != nil {
			return diag.Errorf("error reading persistent_connection: %v", err)
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

	if o.ProxyPassword != nil {
		v := *o.ProxyPassword

		if v == "ENC XXXX" {
		} else if err = d.Set("proxy_password", v); err != nil {
			return diag.Errorf("error reading proxy_password: %v", err)
		}
	}

	if o.ProxyServerIp != nil {
		v := *o.ProxyServerIp

		if err = d.Set("proxy_server_ip", v); err != nil {
			return diag.Errorf("error reading proxy_server_ip: %v", err)
		}
	}

	if o.ProxyServerPort != nil {
		v := *o.ProxyServerPort

		if err = d.Set("proxy_server_port", v); err != nil {
			return diag.Errorf("error reading proxy_server_port: %v", err)
		}
	}

	if o.ProxyUsername != nil {
		v := *o.ProxyUsername

		if err = d.Set("proxy_username", v); err != nil {
			return diag.Errorf("error reading proxy_username: %v", err)
		}
	}

	if o.SandboxRegion != nil {
		v := *o.SandboxRegion

		if err = d.Set("sandbox_region", v); err != nil {
			return diag.Errorf("error reading sandbox_region: %v", err)
		}
	}

	if o.SdnsOptions != nil {
		v := *o.SdnsOptions

		if err = d.Set("sdns_options", v); err != nil {
			return diag.Errorf("error reading sdns_options: %v", err)
		}
	}

	if o.SdnsServerIp != nil {
		v := *o.SdnsServerIp

		if err = d.Set("sdns_server_ip", v); err != nil {
			return diag.Errorf("error reading sdns_server_ip: %v", err)
		}
	}

	if o.SdnsServerPort != nil {
		v := *o.SdnsServerPort

		if err = d.Set("sdns_server_port", v); err != nil {
			return diag.Errorf("error reading sdns_server_port: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.SourceIp6 != nil {
		v := *o.SourceIp6

		if err = d.Set("source_ip6", v); err != nil {
			return diag.Errorf("error reading source_ip6: %v", err)
		}
	}

	if o.UpdateBuildProxy != nil {
		v := *o.UpdateBuildProxy

		if err = d.Set("update_build_proxy", v); err != nil {
			return diag.Errorf("error reading update_build_proxy: %v", err)
		}
	}

	if o.UpdateExtdb != nil {
		v := *o.UpdateExtdb

		if err = d.Set("update_extdb", v); err != nil {
			return diag.Errorf("error reading update_extdb: %v", err)
		}
	}

	if o.UpdateFfdb != nil {
		v := *o.UpdateFfdb

		if err = d.Set("update_ffdb", v); err != nil {
			return diag.Errorf("error reading update_ffdb: %v", err)
		}
	}

	if o.UpdateServerLocation != nil {
		v := *o.UpdateServerLocation

		if err = d.Set("update_server_location", v); err != nil {
			return diag.Errorf("error reading update_server_location: %v", err)
		}
	}

	if o.UpdateUwdb != nil {
		v := *o.UpdateUwdb

		if err = d.Set("update_uwdb", v); err != nil {
			return diag.Errorf("error reading update_uwdb: %v", err)
		}
	}

	if o.VideofilterExpiration != nil {
		v := *o.VideofilterExpiration

		if err = d.Set("videofilter_expiration", v); err != nil {
			return diag.Errorf("error reading videofilter_expiration: %v", err)
		}
	}

	if o.VideofilterLicense != nil {
		v := *o.VideofilterLicense

		if err = d.Set("videofilter_license", v); err != nil {
			return diag.Errorf("error reading videofilter_license: %v", err)
		}
	}

	if o.WebfilterCache != nil {
		v := *o.WebfilterCache

		if err = d.Set("webfilter_cache", v); err != nil {
			return diag.Errorf("error reading webfilter_cache: %v", err)
		}
	}

	if o.WebfilterCacheTtl != nil {
		v := *o.WebfilterCacheTtl

		if err = d.Set("webfilter_cache_ttl", v); err != nil {
			return diag.Errorf("error reading webfilter_cache_ttl: %v", err)
		}
	}

	if o.WebfilterExpiration != nil {
		v := *o.WebfilterExpiration

		if err = d.Set("webfilter_expiration", v); err != nil {
			return diag.Errorf("error reading webfilter_expiration: %v", err)
		}
	}

	if o.WebfilterForceOff != nil {
		v := *o.WebfilterForceOff

		if err = d.Set("webfilter_force_off", v); err != nil {
			return diag.Errorf("error reading webfilter_force_off: %v", err)
		}
	}

	if o.WebfilterLicense != nil {
		v := *o.WebfilterLicense

		if err = d.Set("webfilter_license", v); err != nil {
			return diag.Errorf("error reading webfilter_license: %v", err)
		}
	}

	if o.WebfilterTimeout != nil {
		v := *o.WebfilterTimeout

		if err = d.Set("webfilter_timeout", v); err != nil {
			return diag.Errorf("error reading webfilter_timeout: %v", err)
		}
	}

	return nil
}

func getObjectSystemFortiguard(d *schema.ResourceData, sv string) (*models.SystemFortiguard, diag.Diagnostics) {
	obj := models.SystemFortiguard{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("antispam_cache"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("antispam_cache", sv)
				diags = append(diags, e)
			}
			obj.AntispamCache = &v2
		}
	}
	if v1, ok := d.GetOk("antispam_cache_mpercent"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("antispam_cache_mpercent", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AntispamCacheMpercent = &tmp
		}
	}
	if v1, ok := d.GetOk("antispam_cache_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("antispam_cache_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AntispamCacheTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("antispam_expiration"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("antispam_expiration", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AntispamExpiration = &tmp
		}
	}
	if v1, ok := d.GetOk("antispam_force_off"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("antispam_force_off", sv)
				diags = append(diags, e)
			}
			obj.AntispamForceOff = &v2
		}
	}
	if v1, ok := d.GetOk("antispam_license"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("antispam_license", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AntispamLicense = &tmp
		}
	}
	if v1, ok := d.GetOk("antispam_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("antispam_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AntispamTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("anycast_sdns_server_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("anycast_sdns_server_ip", sv)
				diags = append(diags, e)
			}
			obj.AnycastSdnsServerIp = &v2
		}
	}
	if v1, ok := d.GetOk("anycast_sdns_server_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("anycast_sdns_server_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AnycastSdnsServerPort = &tmp
		}
	}
	if v1, ok := d.GetOk("auto_join_forticloud"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("auto_join_forticloud", sv)
				diags = append(diags, e)
			}
			obj.AutoJoinForticloud = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_server_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_server_ip", sv)
				diags = append(diags, e)
			}
			obj.DdnsServerIp = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_server_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("ddns_server_ip6", sv)
				diags = append(diags, e)
			}
			obj.DdnsServerIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_server_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_server_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DdnsServerPort = &tmp
		}
	}
	if v1, ok := d.GetOk("fortiguard_anycast"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortiguard_anycast", sv)
				diags = append(diags, e)
			}
			obj.FortiguardAnycast = &v2
		}
	}
	if v1, ok := d.GetOk("fortiguard_anycast_source"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortiguard_anycast_source", sv)
				diags = append(diags, e)
			}
			obj.FortiguardAnycastSource = &v2
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface_select_method", sv)
				diags = append(diags, e)
			}
			obj.InterfaceSelectMethod = &v2
		}
	}
	if v1, ok := d.GetOk("load_balance_servers"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("load_balance_servers", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LoadBalanceServers = &tmp
		}
	}
	if v1, ok := d.GetOk("outbreak_prevention_cache"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outbreak_prevention_cache", sv)
				diags = append(diags, e)
			}
			obj.OutbreakPreventionCache = &v2
		}
	}
	if v1, ok := d.GetOk("outbreak_prevention_cache_mpercent"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outbreak_prevention_cache_mpercent", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OutbreakPreventionCacheMpercent = &tmp
		}
	}
	if v1, ok := d.GetOk("outbreak_prevention_cache_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outbreak_prevention_cache_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OutbreakPreventionCacheTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("outbreak_prevention_expiration"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outbreak_prevention_expiration", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OutbreakPreventionExpiration = &tmp
		}
	}
	if v1, ok := d.GetOk("outbreak_prevention_force_off"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outbreak_prevention_force_off", sv)
				diags = append(diags, e)
			}
			obj.OutbreakPreventionForceOff = &v2
		}
	}
	if v1, ok := d.GetOk("outbreak_prevention_license"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outbreak_prevention_license", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OutbreakPreventionLicense = &tmp
		}
	}
	if v1, ok := d.GetOk("outbreak_prevention_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outbreak_prevention_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OutbreakPreventionTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("persistent_connection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("persistent_connection", sv)
				diags = append(diags, e)
			}
			obj.PersistentConnection = &v2
		}
	}
	if v1, ok := d.GetOk("port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port", sv)
				diags = append(diags, e)
			}
			obj.Port = &v2
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
	if v1, ok := d.GetOk("proxy_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_password", sv)
				diags = append(diags, e)
			}
			obj.ProxyPassword = &v2
		}
	}
	if v1, ok := d.GetOk("proxy_server_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_server_ip", sv)
				diags = append(diags, e)
			}
			obj.ProxyServerIp = &v2
		}
	}
	if v1, ok := d.GetOk("proxy_server_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_server_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ProxyServerPort = &tmp
		}
	}
	if v1, ok := d.GetOk("proxy_username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_username", sv)
				diags = append(diags, e)
			}
			obj.ProxyUsername = &v2
		}
	}
	if v1, ok := d.GetOk("sandbox_region"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sandbox_region", sv)
				diags = append(diags, e)
			}
			obj.SandboxRegion = &v2
		}
	}
	if v1, ok := d.GetOk("sdns_options"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("sdns_options", sv)
				diags = append(diags, e)
			}
			obj.SdnsOptions = &v2
		}
	}
	if v1, ok := d.GetOk("sdns_server_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sdns_server_ip", sv)
				diags = append(diags, e)
			}
			obj.SdnsServerIp = &v2
		}
	}
	if v1, ok := d.GetOk("sdns_server_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sdns_server_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SdnsServerPort = &tmp
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
	if v1, ok := d.GetOk("source_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip6", sv)
				diags = append(diags, e)
			}
			obj.SourceIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("update_build_proxy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("update_build_proxy", sv)
				diags = append(diags, e)
			}
			obj.UpdateBuildProxy = &v2
		}
	}
	if v1, ok := d.GetOk("update_extdb"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("update_extdb", sv)
				diags = append(diags, e)
			}
			obj.UpdateExtdb = &v2
		}
	}
	if v1, ok := d.GetOk("update_ffdb"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("update_ffdb", sv)
				diags = append(diags, e)
			}
			obj.UpdateFfdb = &v2
		}
	}
	if v1, ok := d.GetOk("update_server_location"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_server_location", sv)
				diags = append(diags, e)
			}
			obj.UpdateServerLocation = &v2
		}
	}
	if v1, ok := d.GetOk("update_uwdb"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("update_uwdb", sv)
				diags = append(diags, e)
			}
			obj.UpdateUwdb = &v2
		}
	}
	if v1, ok := d.GetOk("videofilter_expiration"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("videofilter_expiration", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.VideofilterExpiration = &tmp
		}
	}
	if v1, ok := d.GetOk("videofilter_license"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("videofilter_license", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.VideofilterLicense = &tmp
		}
	}
	if v1, ok := d.GetOk("webfilter_cache"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webfilter_cache", sv)
				diags = append(diags, e)
			}
			obj.WebfilterCache = &v2
		}
	}
	if v1, ok := d.GetOk("webfilter_cache_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webfilter_cache_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WebfilterCacheTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("webfilter_expiration"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webfilter_expiration", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WebfilterExpiration = &tmp
		}
	}
	if v1, ok := d.GetOk("webfilter_force_off"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webfilter_force_off", sv)
				diags = append(diags, e)
			}
			obj.WebfilterForceOff = &v2
		}
	}
	if v1, ok := d.GetOk("webfilter_license"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webfilter_license", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WebfilterLicense = &tmp
		}
	}
	if v1, ok := d.GetOk("webfilter_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webfilter_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WebfilterTimeout = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemFortiguard(d *schema.ResourceData, sv string) (*models.SystemFortiguard, diag.Diagnostics) {
	obj := models.SystemFortiguard{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
