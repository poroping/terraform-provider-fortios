// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VDOM property.

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

func resourceSystemVdomProperty() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VDOM property.",

		CreateContext: resourceSystemVdomPropertyCreate,
		ReadContext:   resourceSystemVdomPropertyRead,
		UpdateContext: resourceSystemVdomPropertyUpdate,
		DeleteContext: resourceSystemVdomPropertyDelete,

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
			"custom_service": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of firewall custom services.",
				Optional:         true,
				Computed:         true,
			},
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Description.",
				Optional:    true,
				Computed:    true,
			},
			"dialup_tunnel": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of dial-up tunnels.",
				Optional:         true,
				Computed:         true,
			},
			"firewall_address": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).",
				Optional:         true,
				Computed:         true,
			},
			"firewall_addrgrp": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of firewall address groups (IPv4, IPv6).",
				Optional:         true,
				Computed:         true,
			},
			"firewall_policy": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).",
				Optional:         true,
				Computed:         true,
			},
			"ipsec_phase1": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of VPN IPsec phase 1 tunnels.",
				Optional:         true,
				Computed:         true,
			},
			"ipsec_phase1_interface": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of VPN IPsec phase1 interface tunnels.",
				Optional:         true,
				Computed:         true,
			},
			"ipsec_phase2": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of VPN IPsec phase 2 tunnels.",
				Optional:         true,
				Computed:         true,
			},
			"ipsec_phase2_interface": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of VPN IPsec phase2 interface tunnels.",
				Optional:         true,
				Computed:         true,
			},
			"log_disk_quota": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Log disk quota in megabytes (MB). Range depends on how much disk space is available.",
				Optional:         true,
				Computed:         true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "VDOM name.",
				ForceNew:    true,
				Required:    true,
			},
			"onetime_schedule": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of firewall one-time schedules.",
				Optional:         true,
				Computed:         true,
			},
			"proxy": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of concurrent proxy users.",
				Optional:         true,
				Computed:         true,
			},
			"recurring_schedule": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of firewall recurring schedules.",
				Optional:         true,
				Computed:         true,
			},
			"service_group": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of firewall service groups.",
				Optional:         true,
				Computed:         true,
			},
			"session": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of sessions.",
				Optional:         true,
				Computed:         true,
			},
			"snmp_index": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2147483647),

				Description: "Permanent SNMP Index of the virtual domain (1 - 2147483647).",
				Optional:    true,
				Computed:    true,
			},
			"sslvpn": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of SSL-VPNs.",
				Optional:         true,
				Computed:         true,
			},
			"user": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of local users.",
				Optional:         true,
				Computed:         true,
			},
			"user_group": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Maximum guaranteed number of user groups.",
				Optional:         true,
				Computed:         true,
			},
		},
	}
}

func resourceSystemVdomPropertyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemVdomProperty resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemVdomProperty(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemVdomProperty(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVdomProperty")
	}

	return resourceSystemVdomPropertyRead(ctx, d, meta)
}

func resourceSystemVdomPropertyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVdomProperty(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemVdomProperty(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemVdomProperty resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVdomProperty")
	}

	return resourceSystemVdomPropertyRead(ctx, d, meta)
}

func resourceSystemVdomPropertyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemVdomProperty(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemVdomProperty resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomPropertyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVdomProperty(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVdomProperty resource: %v", err)
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

	diags := refreshObjectSystemVdomProperty(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemVdomProperty(d *schema.ResourceData, o *models.SystemVdomProperty, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CustomService != nil {
		v := *o.CustomService

		if err = d.Set("custom_service", v); err != nil {
			return diag.Errorf("error reading custom_service: %v", err)
		}
	}

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.DialupTunnel != nil {
		v := *o.DialupTunnel

		if err = d.Set("dialup_tunnel", v); err != nil {
			return diag.Errorf("error reading dialup_tunnel: %v", err)
		}
	}

	if o.FirewallAddress != nil {
		v := *o.FirewallAddress

		if err = d.Set("firewall_address", v); err != nil {
			return diag.Errorf("error reading firewall_address: %v", err)
		}
	}

	if o.FirewallAddrgrp != nil {
		v := *o.FirewallAddrgrp

		if err = d.Set("firewall_addrgrp", v); err != nil {
			return diag.Errorf("error reading firewall_addrgrp: %v", err)
		}
	}

	if o.FirewallPolicy != nil {
		v := *o.FirewallPolicy

		if err = d.Set("firewall_policy", v); err != nil {
			return diag.Errorf("error reading firewall_policy: %v", err)
		}
	}

	if o.IpsecPhase1 != nil {
		v := *o.IpsecPhase1

		if err = d.Set("ipsec_phase1", v); err != nil {
			return diag.Errorf("error reading ipsec_phase1: %v", err)
		}
	}

	if o.IpsecPhase1Interface != nil {
		v := *o.IpsecPhase1Interface

		if err = d.Set("ipsec_phase1_interface", v); err != nil {
			return diag.Errorf("error reading ipsec_phase1_interface: %v", err)
		}
	}

	if o.IpsecPhase2 != nil {
		v := *o.IpsecPhase2

		if err = d.Set("ipsec_phase2", v); err != nil {
			return diag.Errorf("error reading ipsec_phase2: %v", err)
		}
	}

	if o.IpsecPhase2Interface != nil {
		v := *o.IpsecPhase2Interface

		if err = d.Set("ipsec_phase2_interface", v); err != nil {
			return diag.Errorf("error reading ipsec_phase2_interface: %v", err)
		}
	}

	if o.LogDiskQuota != nil {
		v := *o.LogDiskQuota

		if err = d.Set("log_disk_quota", v); err != nil {
			return diag.Errorf("error reading log_disk_quota: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.OnetimeSchedule != nil {
		v := *o.OnetimeSchedule

		if err = d.Set("onetime_schedule", v); err != nil {
			return diag.Errorf("error reading onetime_schedule: %v", err)
		}
	}

	if o.Proxy != nil {
		v := *o.Proxy

		if err = d.Set("proxy", v); err != nil {
			return diag.Errorf("error reading proxy: %v", err)
		}
	}

	if o.RecurringSchedule != nil {
		v := *o.RecurringSchedule

		if err = d.Set("recurring_schedule", v); err != nil {
			return diag.Errorf("error reading recurring_schedule: %v", err)
		}
	}

	if o.ServiceGroup != nil {
		v := *o.ServiceGroup

		if err = d.Set("service_group", v); err != nil {
			return diag.Errorf("error reading service_group: %v", err)
		}
	}

	if o.Session != nil {
		v := *o.Session

		if err = d.Set("session", v); err != nil {
			return diag.Errorf("error reading session: %v", err)
		}
	}

	if o.SnmpIndex != nil {
		v := *o.SnmpIndex

		if err = d.Set("snmp_index", v); err != nil {
			return diag.Errorf("error reading snmp_index: %v", err)
		}
	}

	if o.Sslvpn != nil {
		v := *o.Sslvpn

		if err = d.Set("sslvpn", v); err != nil {
			return diag.Errorf("error reading sslvpn: %v", err)
		}
	}

	if o.User != nil {
		v := *o.User

		if err = d.Set("user", v); err != nil {
			return diag.Errorf("error reading user: %v", err)
		}
	}

	if o.UserGroup != nil {
		v := *o.UserGroup

		if err = d.Set("user_group", v); err != nil {
			return diag.Errorf("error reading user_group: %v", err)
		}
	}

	return nil
}

func getObjectSystemVdomProperty(d *schema.ResourceData, sv string) (*models.SystemVdomProperty, diag.Diagnostics) {
	obj := models.SystemVdomProperty{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("custom_service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("custom_service", sv)
				diags = append(diags, e)
			}
			obj.CustomService = &v2
		}
	}
	if v1, ok := d.GetOk("description"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("description", sv)
				diags = append(diags, e)
			}
			obj.Description = &v2
		}
	}
	if v1, ok := d.GetOk("dialup_tunnel"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dialup_tunnel", sv)
				diags = append(diags, e)
			}
			obj.DialupTunnel = &v2
		}
	}
	if v1, ok := d.GetOk("firewall_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("firewall_address", sv)
				diags = append(diags, e)
			}
			obj.FirewallAddress = &v2
		}
	}
	if v1, ok := d.GetOk("firewall_addrgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("firewall_addrgrp", sv)
				diags = append(diags, e)
			}
			obj.FirewallAddrgrp = &v2
		}
	}
	if v1, ok := d.GetOk("firewall_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("firewall_policy", sv)
				diags = append(diags, e)
			}
			obj.FirewallPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_phase1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_phase1", sv)
				diags = append(diags, e)
			}
			obj.IpsecPhase1 = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_phase1_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_phase1_interface", sv)
				diags = append(diags, e)
			}
			obj.IpsecPhase1Interface = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_phase2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_phase2", sv)
				diags = append(diags, e)
			}
			obj.IpsecPhase2 = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_phase2_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_phase2_interface", sv)
				diags = append(diags, e)
			}
			obj.IpsecPhase2Interface = &v2
		}
	}
	if v1, ok := d.GetOk("log_disk_quota"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_disk_quota", sv)
				diags = append(diags, e)
			}
			obj.LogDiskQuota = &v2
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
	if v1, ok := d.GetOk("onetime_schedule"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("onetime_schedule", sv)
				diags = append(diags, e)
			}
			obj.OnetimeSchedule = &v2
		}
	}
	if v1, ok := d.GetOk("proxy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy", sv)
				diags = append(diags, e)
			}
			obj.Proxy = &v2
		}
	}
	if v1, ok := d.GetOk("recurring_schedule"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("recurring_schedule", sv)
				diags = append(diags, e)
			}
			obj.RecurringSchedule = &v2
		}
	}
	if v1, ok := d.GetOk("service_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("service_group", sv)
				diags = append(diags, e)
			}
			obj.ServiceGroup = &v2
		}
	}
	if v1, ok := d.GetOk("session"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session", sv)
				diags = append(diags, e)
			}
			obj.Session = &v2
		}
	}
	if v1, ok := d.GetOk("snmp_index"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("snmp_index", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SnmpIndex = &tmp
		}
	}
	if v1, ok := d.GetOk("sslvpn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sslvpn", sv)
				diags = append(diags, e)
			}
			obj.Sslvpn = &v2
		}
	}
	if v1, ok := d.GetOk("user"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user", sv)
				diags = append(diags, e)
			}
			obj.User = &v2
		}
	}
	if v1, ok := d.GetOk("user_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_group", sv)
				diags = append(diags, e)
			}
			obj.UserGroup = &v2
		}
	}
	return &obj, diags
}
