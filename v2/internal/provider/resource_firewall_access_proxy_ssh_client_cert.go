// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Access Proxy SSH client certificate.

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

func resourceFirewallAccessProxySshClientCert() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Access Proxy SSH client certificate.",

		CreateContext: resourceFirewallAccessProxySshClientCertCreate,
		ReadContext:   resourceFirewallAccessProxySshClientCertRead,
		UpdateContext: resourceFirewallAccessProxySshClientCertUpdate,
		DeleteContext: resourceFirewallAccessProxySshClientCertDelete,

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
			"auth_ca": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Name of the SSH server public key authentication CA.",
				Optional:    true,
				Computed:    true,
			},
			"cert_extension": {
				Type:        schema.TypeList,
				Description: "Configure certificate extension for user certificate.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"critical": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"no", "yes"}, false),

							Description: "Critical option.",
							Optional:    true,
							Computed:    true,
						},
						"data": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Data of certificate extension.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Name of certificate extension.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"fixed", "user"}, false),

							Description: "Type of certificate extension.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "SSH client certificate name.",
				ForceNew:    true,
				Required:    true,
			},
			"permit_agent_forwarding": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable appending permit-agent-forwarding certificate extension.",
				Optional:    true,
				Computed:    true,
			},
			"permit_port_forwarding": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable appending permit-port-forwarding certificate extension.",
				Optional:    true,
				Computed:    true,
			},
			"permit_pty": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable appending permit-pty certificate extension.",
				Optional:    true,
				Computed:    true,
			},
			"permit_user_rc": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable appending permit-user-rc certificate extension.",
				Optional:    true,
				Computed:    true,
			},
			"permit_x11_forwarding": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable appending permit-x11-forwarding certificate extension.",
				Optional:    true,
				Computed:    true,
			},
			"source_address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallAccessProxySshClientCertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallAccessProxySshClientCert resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallAccessProxySshClientCert(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallAccessProxySshClientCert(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAccessProxySshClientCert")
	}

	return resourceFirewallAccessProxySshClientCertRead(ctx, d, meta)
}

func resourceFirewallAccessProxySshClientCertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallAccessProxySshClientCert(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallAccessProxySshClientCert(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallAccessProxySshClientCert resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAccessProxySshClientCert")
	}

	return resourceFirewallAccessProxySshClientCertRead(ctx, d, meta)
}

func resourceFirewallAccessProxySshClientCertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallAccessProxySshClientCert(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallAccessProxySshClientCert resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxySshClientCertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallAccessProxySshClientCert(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAccessProxySshClientCert resource: %v", err)
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

	diags := refreshObjectFirewallAccessProxySshClientCert(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallAccessProxySshClientCertCertExtension(d *schema.ResourceData, v *[]models.FirewallAccessProxySshClientCertCertExtension, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Critical; tmp != nil {
				v["critical"] = *tmp
			}

			if tmp := cfg.Data; tmp != nil {
				v["data"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectFirewallAccessProxySshClientCert(d *schema.ResourceData, o *models.FirewallAccessProxySshClientCert, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthCa != nil {
		v := *o.AuthCa

		if err = d.Set("auth_ca", v); err != nil {
			return diag.Errorf("error reading auth_ca: %v", err)
		}
	}

	if o.CertExtension != nil {
		if err = d.Set("cert_extension", flattenFirewallAccessProxySshClientCertCertExtension(d, o.CertExtension, "cert_extension", sort)); err != nil {
			return diag.Errorf("error reading cert_extension: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PermitAgentForwarding != nil {
		v := *o.PermitAgentForwarding

		if err = d.Set("permit_agent_forwarding", v); err != nil {
			return diag.Errorf("error reading permit_agent_forwarding: %v", err)
		}
	}

	if o.PermitPortForwarding != nil {
		v := *o.PermitPortForwarding

		if err = d.Set("permit_port_forwarding", v); err != nil {
			return diag.Errorf("error reading permit_port_forwarding: %v", err)
		}
	}

	if o.PermitPty != nil {
		v := *o.PermitPty

		if err = d.Set("permit_pty", v); err != nil {
			return diag.Errorf("error reading permit_pty: %v", err)
		}
	}

	if o.PermitUserRc != nil {
		v := *o.PermitUserRc

		if err = d.Set("permit_user_rc", v); err != nil {
			return diag.Errorf("error reading permit_user_rc: %v", err)
		}
	}

	if o.PermitX11Forwarding != nil {
		v := *o.PermitX11Forwarding

		if err = d.Set("permit_x11_forwarding", v); err != nil {
			return diag.Errorf("error reading permit_x11_forwarding: %v", err)
		}
	}

	if o.SourceAddress != nil {
		v := *o.SourceAddress

		if err = d.Set("source_address", v); err != nil {
			return diag.Errorf("error reading source_address: %v", err)
		}
	}

	return nil
}

func expandFirewallAccessProxySshClientCertCertExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAccessProxySshClientCertCertExtension, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAccessProxySshClientCertCertExtension

	for i := range l {
		tmp := models.FirewallAccessProxySshClientCertCertExtension{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.critical", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Critical = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.data", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Data = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectFirewallAccessProxySshClientCert(d *schema.ResourceData, sv string) (*models.FirewallAccessProxySshClientCert, diag.Diagnostics) {
	obj := models.FirewallAccessProxySshClientCert{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth_ca"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_ca", sv)
				diags = append(diags, e)
			}
			obj.AuthCa = &v2
		}
	}
	if v, ok := d.GetOk("cert_extension"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("cert_extension", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAccessProxySshClientCertCertExtension(d, v, "cert_extension", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.CertExtension = t
		}
	} else if d.HasChange("cert_extension") {
		old, new := d.GetChange("cert_extension")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.CertExtension = &[]models.FirewallAccessProxySshClientCertCertExtension{}
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
	if v1, ok := d.GetOk("permit_agent_forwarding"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("permit_agent_forwarding", sv)
				diags = append(diags, e)
			}
			obj.PermitAgentForwarding = &v2
		}
	}
	if v1, ok := d.GetOk("permit_port_forwarding"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("permit_port_forwarding", sv)
				diags = append(diags, e)
			}
			obj.PermitPortForwarding = &v2
		}
	}
	if v1, ok := d.GetOk("permit_pty"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("permit_pty", sv)
				diags = append(diags, e)
			}
			obj.PermitPty = &v2
		}
	}
	if v1, ok := d.GetOk("permit_user_rc"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("permit_user_rc", sv)
				diags = append(diags, e)
			}
			obj.PermitUserRc = &v2
		}
	}
	if v1, ok := d.GetOk("permit_x11_forwarding"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("permit_x11_forwarding", sv)
				diags = append(diags, e)
			}
			obj.PermitX11Forwarding = &v2
		}
	}
	if v1, ok := d.GetOk("source_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_address", sv)
				diags = append(diags, e)
			}
			obj.SourceAddress = &v2
		}
	}
	return &obj, diags
}
