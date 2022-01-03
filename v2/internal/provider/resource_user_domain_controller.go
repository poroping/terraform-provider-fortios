// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure domain controller entries.

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

func resourceUserDomainController() *schema.Resource {
	return &schema.Resource{
		Description: "Configure domain controller entries.",

		CreateContext: resourceUserDomainControllerCreate,
		ReadContext:   resourceUserDomainControllerRead,
		UpdateContext: resourceUserDomainControllerUpdate,
		DeleteContext: resourceUserDomainControllerDelete,

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
			"ad_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "ds", "lds"}, false),

				Description: "Set Active Directory mode.",
				Optional:    true,
				Computed:    true,
			},
			"adlds_dn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "AD LDS distinguished name.",
				Optional:    true,
				Computed:    true,
			},
			"adlds_ip_address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "AD LDS IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"adlds_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "AD LDS IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"adlds_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port number of AD LDS service (default = 389).",
				Optional:    true,
				Computed:    true,
			},
			"dns_srv_lookup": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DNS service lookup.",
				Optional:    true,
				Computed:    true,
			},
			"domain_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Domain DNS name.",
				Optional:    true,
				Computed:    true,
			},
			"extra_server": {
				Type:        schema.TypeList,
				Description: "extra servers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Server ID.",
							Optional:    true,
							Computed:    true,
						},
						"ip_address": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Domain controller IP address.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Port to be used for communication with the domain controller (default = 445).",
							Optional:    true,
							Computed:    true,
						},
						"source_ip_address": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "FortiGate IPv4 address to be used for communication with the domain controller.",
							Optional:    true,
							Computed:    true,
						},
						"source_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Source port to be used for communication with the domain controller.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"hostname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Hostname of the server to connect to.",
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
			"ip_address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Domain controller IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Domain controller IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"ldap_server": {
				Type:        schema.TypeList,
				Description: "LDAP server name(s).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "LDAP server name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Domain controller entry name.",
				ForceNew:    true,
				Required:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Password for specified username.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port to be used for communication with the domain controller (default = 445).",
				Optional:    true,
				Computed:    true,
			},
			"replication_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip_address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "FortiGate IPv4 address to be used for communication with the domain controller.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "FortiGate IPv6 address to be used for communication with the domain controller.",
				Optional:         true,
				Computed:         true,
			},
			"source_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Source port to be used for communication with the domain controller.",
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "User name to sign in with. Must have proper permissions for service.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserDomainControllerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserDomainController resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserDomainController(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserDomainController(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserDomainController")
	}

	return resourceUserDomainControllerRead(ctx, d, meta)
}

func resourceUserDomainControllerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserDomainController(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserDomainController(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserDomainController resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserDomainController")
	}

	return resourceUserDomainControllerRead(ctx, d, meta)
}

func resourceUserDomainControllerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserDomainController(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserDomainController resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserDomainControllerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserDomainController(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserDomainController resource: %v", err)
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

	diags := refreshObjectUserDomainController(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenUserDomainControllerExtraServer(v *[]models.UserDomainControllerExtraServer, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.IpAddress; tmp != nil {
				v["ip_address"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.SourceIpAddress; tmp != nil {
				v["source_ip_address"] = *tmp
			}

			if tmp := cfg.SourcePort; tmp != nil {
				v["source_port"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenUserDomainControllerLdapServer(v *[]models.UserDomainControllerLdapServer, sort bool) interface{} {
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

func refreshObjectUserDomainController(d *schema.ResourceData, o *models.UserDomainController, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AdMode != nil {
		v := *o.AdMode

		if err = d.Set("ad_mode", v); err != nil {
			return diag.Errorf("error reading ad_mode: %v", err)
		}
	}

	if o.AdldsDn != nil {
		v := *o.AdldsDn

		if err = d.Set("adlds_dn", v); err != nil {
			return diag.Errorf("error reading adlds_dn: %v", err)
		}
	}

	if o.AdldsIpAddress != nil {
		v := *o.AdldsIpAddress

		if err = d.Set("adlds_ip_address", v); err != nil {
			return diag.Errorf("error reading adlds_ip_address: %v", err)
		}
	}

	if o.AdldsIp6 != nil {
		v := *o.AdldsIp6

		if err = d.Set("adlds_ip6", v); err != nil {
			return diag.Errorf("error reading adlds_ip6: %v", err)
		}
	}

	if o.AdldsPort != nil {
		v := *o.AdldsPort

		if err = d.Set("adlds_port", v); err != nil {
			return diag.Errorf("error reading adlds_port: %v", err)
		}
	}

	if o.DnsSrvLookup != nil {
		v := *o.DnsSrvLookup

		if err = d.Set("dns_srv_lookup", v); err != nil {
			return diag.Errorf("error reading dns_srv_lookup: %v", err)
		}
	}

	if o.DomainName != nil {
		v := *o.DomainName

		if err = d.Set("domain_name", v); err != nil {
			return diag.Errorf("error reading domain_name: %v", err)
		}
	}

	if o.ExtraServer != nil {
		if err = d.Set("extra_server", flattenUserDomainControllerExtraServer(o.ExtraServer, sort)); err != nil {
			return diag.Errorf("error reading extra_server: %v", err)
		}
	}

	if o.Hostname != nil {
		v := *o.Hostname

		if err = d.Set("hostname", v); err != nil {
			return diag.Errorf("error reading hostname: %v", err)
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

	if o.IpAddress != nil {
		v := *o.IpAddress

		if err = d.Set("ip_address", v); err != nil {
			return diag.Errorf("error reading ip_address: %v", err)
		}
	}

	if o.Ip6 != nil {
		v := *o.Ip6

		if err = d.Set("ip6", v); err != nil {
			return diag.Errorf("error reading ip6: %v", err)
		}
	}

	if o.LdapServer != nil {
		if err = d.Set("ldap_server", flattenUserDomainControllerLdapServer(o.LdapServer, sort)); err != nil {
			return diag.Errorf("error reading ldap_server: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.ReplicationPort != nil {
		v := *o.ReplicationPort

		if err = d.Set("replication_port", v); err != nil {
			return diag.Errorf("error reading replication_port: %v", err)
		}
	}

	if o.SourceIpAddress != nil {
		v := *o.SourceIpAddress

		if err = d.Set("source_ip_address", v); err != nil {
			return diag.Errorf("error reading source_ip_address: %v", err)
		}
	}

	if o.SourceIp6 != nil {
		v := *o.SourceIp6

		if err = d.Set("source_ip6", v); err != nil {
			return diag.Errorf("error reading source_ip6: %v", err)
		}
	}

	if o.SourcePort != nil {
		v := *o.SourcePort

		if err = d.Set("source_port", v); err != nil {
			return diag.Errorf("error reading source_port: %v", err)
		}
	}

	if o.Username != nil {
		v := *o.Username

		if err = d.Set("username", v); err != nil {
			return diag.Errorf("error reading username: %v", err)
		}
	}

	return nil
}

func expandUserDomainControllerExtraServer(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserDomainControllerExtraServer, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserDomainControllerExtraServer

	for i := range l {
		tmp := models.UserDomainControllerExtraServer{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IpAddress = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Port = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source_ip_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SourceIpAddress = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.SourcePort = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandUserDomainControllerLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserDomainControllerLdapServer, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserDomainControllerLdapServer

	for i := range l {
		tmp := models.UserDomainControllerLdapServer{}
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

func getObjectUserDomainController(d *schema.ResourceData, sv string) (*models.UserDomainController, diag.Diagnostics) {
	obj := models.UserDomainController{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("ad_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ad_mode", sv)
				diags = append(diags, e)
			}
			obj.AdMode = &v2
		}
	}
	if v1, ok := d.GetOk("adlds_dn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("adlds_dn", sv)
				diags = append(diags, e)
			}
			obj.AdldsDn = &v2
		}
	}
	if v1, ok := d.GetOk("adlds_ip_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("adlds_ip_address", sv)
				diags = append(diags, e)
			}
			obj.AdldsIpAddress = &v2
		}
	}
	if v1, ok := d.GetOk("adlds_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("adlds_ip6", sv)
				diags = append(diags, e)
			}
			obj.AdldsIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("adlds_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("adlds_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdldsPort = &tmp
		}
	}
	if v1, ok := d.GetOk("dns_srv_lookup"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("dns_srv_lookup", sv)
				diags = append(diags, e)
			}
			obj.DnsSrvLookup = &v2
		}
	}
	if v1, ok := d.GetOk("domain_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("domain_name", sv)
				diags = append(diags, e)
			}
			obj.DomainName = &v2
		}
	}
	if v, ok := d.GetOk("extra_server"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("extra_server", sv)
			diags = append(diags, e)
		}
		t, err := expandUserDomainControllerExtraServer(d, v, "extra_server", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ExtraServer = t
		}
	} else if d.HasChange("extra_server") {
		old, new := d.GetChange("extra_server")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ExtraServer = &[]models.UserDomainControllerExtraServer{}
		}
	}
	if v1, ok := d.GetOk("hostname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("hostname", sv)
				diags = append(diags, e)
			}
			obj.Hostname = &v2
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("interface_select_method", sv)
				diags = append(diags, e)
			}
			obj.InterfaceSelectMethod = &v2
		}
	}
	if v1, ok := d.GetOk("ip_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_address", sv)
				diags = append(diags, e)
			}
			obj.IpAddress = &v2
		}
	}
	if v1, ok := d.GetOk("ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ip6", sv)
				diags = append(diags, e)
			}
			obj.Ip6 = &v2
		}
	}
	if v, ok := d.GetOk("ldap_server"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ldap_server", sv)
			diags = append(diags, e)
		}
		t, err := expandUserDomainControllerLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.LdapServer = t
		}
	} else if d.HasChange("ldap_server") {
		old, new := d.GetChange("ldap_server")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.LdapServer = &[]models.UserDomainControllerLdapServer{}
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
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
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
	if v1, ok := d.GetOk("replication_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("replication_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ReplicationPort = &tmp
		}
	}
	if v1, ok := d.GetOk("source_ip_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("source_ip_address", sv)
				diags = append(diags, e)
			}
			obj.SourceIpAddress = &v2
		}
	}
	if v1, ok := d.GetOk("source_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("source_ip6", sv)
				diags = append(diags, e)
			}
			obj.SourceIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("source_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("source_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SourcePort = &tmp
		}
	}
	if v1, ok := d.GetOk("username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("username", sv)
				diags = append(diags, e)
			}
			obj.Username = &v2
		}
	}
	return &obj, diags
}
