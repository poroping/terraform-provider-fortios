// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FSSO active directory servers for polling mode.

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

func resourceUserFssoPolling() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FSSO active directory servers for polling mode.",

		CreateContext: resourceUserFssoPollingCreate,
		ReadContext:   resourceUserFssoPollingRead,
		UpdateContext: resourceUserFssoPollingUpdate,
		DeleteContext: resourceUserFssoPollingDelete,

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
				Type:         schema.TypeBool,
				Description:  "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"adgrp": {
				Type:        schema.TypeList,
				Description: "LDAP Group Info.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "Name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"default_domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Default domain managed by this Active Directory server.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Active Directory server ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"ldap_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "LDAP server name used in LDAP connection strings.",
				Optional:    true,
				Computed:    true,
			},
			"logon_history": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 48),

				Description: "Number of hours of logon history to keep, 0 means keep all history.",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Password required to log into this Active Directory server.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"polling_frequency": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),

				Description: "Polling frequency (every 1 to 30 seconds).",
				Optional:    true,
				Computed:    true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port to communicate with this Active Directory server.",
				Optional:    true,
				Computed:    true,
			},
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Host name or IP address of the Active Directory server.",
				Optional:    true,
				Computed:    true,
			},
			"smb_ntlmv1_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable support of NTLMv1 for Samba authentication.",
				Optional:    true,
				Computed:    true,
			},
			"smbv1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable support of SMBv1 for Samba.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable polling for the status of this Active Directory server.",
				Optional:    true,
				Computed:    true,
			},
			"user": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "User name required to log into this Active Directory server.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserFssoPollingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating UserFssoPolling resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserFssoPolling(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserFssoPolling(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserFssoPolling")
	}

	return resourceUserFssoPollingRead(ctx, d, meta)
}

func resourceUserFssoPollingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserFssoPolling(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserFssoPolling(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserFssoPolling resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserFssoPolling")
	}

	return resourceUserFssoPollingRead(ctx, d, meta)
}

func resourceUserFssoPollingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserFssoPolling(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserFssoPolling resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserFssoPollingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserFssoPolling(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserFssoPolling resource: %v", err)
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

	diags := refreshObjectUserFssoPolling(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenUserFssoPollingAdgrp(d *schema.ResourceData, v *[]models.UserFssoPollingAdgrp, prefix string, sort bool) interface{} {
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

func refreshObjectUserFssoPolling(d *schema.ResourceData, o *models.UserFssoPolling, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Adgrp != nil {
		if err = d.Set("adgrp", flattenUserFssoPollingAdgrp(d, o.Adgrp, "adgrp", sort)); err != nil {
			return diag.Errorf("error reading adgrp: %v", err)
		}
	}

	if o.DefaultDomain != nil {
		v := *o.DefaultDomain

		if err = d.Set("default_domain", v); err != nil {
			return diag.Errorf("error reading default_domain: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.LdapServer != nil {
		v := *o.LdapServer

		if err = d.Set("ldap_server", v); err != nil {
			return diag.Errorf("error reading ldap_server: %v", err)
		}
	}

	if o.LogonHistory != nil {
		v := *o.LogonHistory

		if err = d.Set("logon_history", v); err != nil {
			return diag.Errorf("error reading logon_history: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.PollingFrequency != nil {
		v := *o.PollingFrequency

		if err = d.Set("polling_frequency", v); err != nil {
			return diag.Errorf("error reading polling_frequency: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.SmbNtlmv1Auth != nil {
		v := *o.SmbNtlmv1Auth

		if err = d.Set("smb_ntlmv1_auth", v); err != nil {
			return diag.Errorf("error reading smb_ntlmv1_auth: %v", err)
		}
	}

	if o.Smbv1 != nil {
		v := *o.Smbv1

		if err = d.Set("smbv1", v); err != nil {
			return diag.Errorf("error reading smbv1: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.User != nil {
		v := *o.User

		if err = d.Set("user", v); err != nil {
			return diag.Errorf("error reading user: %v", err)
		}
	}

	return nil
}

func expandUserFssoPollingAdgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserFssoPollingAdgrp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserFssoPollingAdgrp

	for i := range l {
		tmp := models.UserFssoPollingAdgrp{}
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

func getObjectUserFssoPolling(d *schema.ResourceData, sv string) (*models.UserFssoPolling, diag.Diagnostics) {
	obj := models.UserFssoPolling{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("adgrp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("adgrp", sv)
			diags = append(diags, e)
		}
		t, err := expandUserFssoPollingAdgrp(d, v, "adgrp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Adgrp = t
		}
	} else if d.HasChange("adgrp") {
		old, new := d.GetChange("adgrp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Adgrp = &[]models.UserFssoPollingAdgrp{}
		}
	}
	if v1, ok := d.GetOk("default_domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_domain", sv)
				diags = append(diags, e)
			}
			obj.DefaultDomain = &v2
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	if v1, ok := d.GetOk("ldap_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldap_server", sv)
				diags = append(diags, e)
			}
			obj.LdapServer = &v2
		}
	}
	if v1, ok := d.GetOk("logon_history"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("logon_history", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LogonHistory = &tmp
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
	if v1, ok := d.GetOk("polling_frequency"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("polling_frequency", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PollingFrequency = &tmp
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
	if v1, ok := d.GetOk("server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server", sv)
				diags = append(diags, e)
			}
			obj.Server = &v2
		}
	}
	if v1, ok := d.GetOk("smb_ntlmv1_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("smb_ntlmv1_auth", sv)
				diags = append(diags, e)
			}
			obj.SmbNtlmv1Auth = &v2
		}
	}
	if v1, ok := d.GetOk("smbv1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("smbv1", sv)
				diags = append(diags, e)
			}
			obj.Smbv1 = &v2
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
	if v1, ok := d.GetOk("user"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user", sv)
				diags = append(diags, e)
			}
			obj.User = &v2
		}
	}
	return &obj, diags
}
