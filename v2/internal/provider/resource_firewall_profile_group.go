// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure profile groups.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceFirewallProfileGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure profile groups.",

		CreateContext: resourceFirewallProfileGroupCreate,
		ReadContext:   resourceFirewallProfileGroupRead,
		UpdateContext: resourceFirewallProfileGroupUpdate,
		DeleteContext: resourceFirewallProfileGroupDelete,

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
			"application_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Application list.",
				Optional:    true,
				Computed:    true,
			},
			"av_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Antivirus profile.",
				Optional:    true,
				Computed:    true,
			},
			"cifs_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing CIFS profile.",
				Optional:    true,
				Computed:    true,
			},
			"dlp_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing DLP profile.",
				Optional:    true,
				Computed:    true,
			},
			"dlp_sensor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing DLP sensor.",
				Optional:    true,
				Computed:    true,
			},
			"dnsfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing DNS filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"emailfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing email filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"file_filter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing file-filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"icap_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing ICAP profile.",
				Optional:    true,
				Computed:    true,
			},
			"ips_sensor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing IPS sensor.",
				Optional:    true,
				Computed:    true,
			},
			"ips_voip_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing VoIP (ips) profile.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Profile group name.",
				ForceNew:    true,
				Required:    true,
			},
			"profile_protocol_options": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Protocol options profile.",
				Optional:    true,
				Computed:    true,
			},
			"sctp_filter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing SCTP filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_filter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing SSH filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_ssh_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing SSL SSH profile.",
				Optional:    true,
				Computed:    true,
			},
			"videofilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing VideoFilter profile.",
				Optional:    true,
				Computed:    true,
			},
			"voip_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing VoIP (voipd) profile.",
				Optional:    true,
				Computed:    true,
			},
			"waf_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Web application firewall profile.",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Web filter profile.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallProfileGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallProfileGroup resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallProfileGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallProfileGroup(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallProfileGroup")
	}

	return resourceFirewallProfileGroupRead(ctx, d, meta)
}

func resourceFirewallProfileGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallProfileGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallProfileGroup(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallProfileGroup resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallProfileGroup")
	}

	return resourceFirewallProfileGroupRead(ctx, d, meta)
}

func resourceFirewallProfileGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallProfileGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallProfileGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProfileGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallProfileGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallProfileGroup resource: %v", err)
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

	diags := refreshObjectFirewallProfileGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallProfileGroup(d *schema.ResourceData, o *models.FirewallProfileGroup, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ApplicationList != nil {
		v := *o.ApplicationList

		if err = d.Set("application_list", v); err != nil {
			return diag.Errorf("error reading application_list: %v", err)
		}
	}

	if o.AvProfile != nil {
		v := *o.AvProfile

		if err = d.Set("av_profile", v); err != nil {
			return diag.Errorf("error reading av_profile: %v", err)
		}
	}

	if o.CifsProfile != nil {
		v := *o.CifsProfile

		if err = d.Set("cifs_profile", v); err != nil {
			return diag.Errorf("error reading cifs_profile: %v", err)
		}
	}

	if o.DlpProfile != nil {
		v := *o.DlpProfile

		if err = d.Set("dlp_profile", v); err != nil {
			return diag.Errorf("error reading dlp_profile: %v", err)
		}
	}

	if o.DlpSensor != nil {
		v := *o.DlpSensor

		if err = d.Set("dlp_sensor", v); err != nil {
			return diag.Errorf("error reading dlp_sensor: %v", err)
		}
	}

	if o.DnsfilterProfile != nil {
		v := *o.DnsfilterProfile

		if err = d.Set("dnsfilter_profile", v); err != nil {
			return diag.Errorf("error reading dnsfilter_profile: %v", err)
		}
	}

	if o.EmailfilterProfile != nil {
		v := *o.EmailfilterProfile

		if err = d.Set("emailfilter_profile", v); err != nil {
			return diag.Errorf("error reading emailfilter_profile: %v", err)
		}
	}

	if o.FileFilterProfile != nil {
		v := *o.FileFilterProfile

		if err = d.Set("file_filter_profile", v); err != nil {
			return diag.Errorf("error reading file_filter_profile: %v", err)
		}
	}

	if o.IcapProfile != nil {
		v := *o.IcapProfile

		if err = d.Set("icap_profile", v); err != nil {
			return diag.Errorf("error reading icap_profile: %v", err)
		}
	}

	if o.IpsSensor != nil {
		v := *o.IpsSensor

		if err = d.Set("ips_sensor", v); err != nil {
			return diag.Errorf("error reading ips_sensor: %v", err)
		}
	}

	if o.IpsVoipFilter != nil {
		v := *o.IpsVoipFilter

		if err = d.Set("ips_voip_filter", v); err != nil {
			return diag.Errorf("error reading ips_voip_filter: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ProfileProtocolOptions != nil {
		v := *o.ProfileProtocolOptions

		if err = d.Set("profile_protocol_options", v); err != nil {
			return diag.Errorf("error reading profile_protocol_options: %v", err)
		}
	}

	if o.SctpFilterProfile != nil {
		v := *o.SctpFilterProfile

		if err = d.Set("sctp_filter_profile", v); err != nil {
			return diag.Errorf("error reading sctp_filter_profile: %v", err)
		}
	}

	if o.SshFilterProfile != nil {
		v := *o.SshFilterProfile

		if err = d.Set("ssh_filter_profile", v); err != nil {
			return diag.Errorf("error reading ssh_filter_profile: %v", err)
		}
	}

	if o.SslSshProfile != nil {
		v := *o.SslSshProfile

		if err = d.Set("ssl_ssh_profile", v); err != nil {
			return diag.Errorf("error reading ssl_ssh_profile: %v", err)
		}
	}

	if o.VideofilterProfile != nil {
		v := *o.VideofilterProfile

		if err = d.Set("videofilter_profile", v); err != nil {
			return diag.Errorf("error reading videofilter_profile: %v", err)
		}
	}

	if o.VoipProfile != nil {
		v := *o.VoipProfile

		if err = d.Set("voip_profile", v); err != nil {
			return diag.Errorf("error reading voip_profile: %v", err)
		}
	}

	if o.WafProfile != nil {
		v := *o.WafProfile

		if err = d.Set("waf_profile", v); err != nil {
			return diag.Errorf("error reading waf_profile: %v", err)
		}
	}

	if o.WebfilterProfile != nil {
		v := *o.WebfilterProfile

		if err = d.Set("webfilter_profile", v); err != nil {
			return diag.Errorf("error reading webfilter_profile: %v", err)
		}
	}

	return nil
}

func getObjectFirewallProfileGroup(d *schema.ResourceData, sv string) (*models.FirewallProfileGroup, diag.Diagnostics) {
	obj := models.FirewallProfileGroup{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("application_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("application_list", sv)
				diags = append(diags, e)
			}
			obj.ApplicationList = &v2
		}
	}
	if v1, ok := d.GetOk("av_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("av_profile", sv)
				diags = append(diags, e)
			}
			obj.AvProfile = &v2
		}
	}
	if v1, ok := d.GetOk("cifs_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cifs_profile", sv)
				diags = append(diags, e)
			}
			obj.CifsProfile = &v2
		}
	}
	if v1, ok := d.GetOk("dlp_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("dlp_profile", sv)
				diags = append(diags, e)
			}
			obj.DlpProfile = &v2
		}
	}
	if v1, ok := d.GetOk("dlp_sensor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.0") {
				e := utils.AttributeVersionWarning("dlp_sensor", sv)
				diags = append(diags, e)
			}
			obj.DlpSensor = &v2
		}
	}
	if v1, ok := d.GetOk("dnsfilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dnsfilter_profile", sv)
				diags = append(diags, e)
			}
			obj.DnsfilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("emailfilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("emailfilter_profile", sv)
				diags = append(diags, e)
			}
			obj.EmailfilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("file_filter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("file_filter_profile", sv)
				diags = append(diags, e)
			}
			obj.FileFilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("icap_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("icap_profile", sv)
				diags = append(diags, e)
			}
			obj.IcapProfile = &v2
		}
	}
	if v1, ok := d.GetOk("ips_sensor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ips_sensor", sv)
				diags = append(diags, e)
			}
			obj.IpsSensor = &v2
		}
	}
	if v1, ok := d.GetOk("ips_voip_filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("ips_voip_filter", sv)
				diags = append(diags, e)
			}
			obj.IpsVoipFilter = &v2
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
	if v1, ok := d.GetOk("profile_protocol_options"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("profile_protocol_options", sv)
				diags = append(diags, e)
			}
			obj.ProfileProtocolOptions = &v2
		}
	}
	if v1, ok := d.GetOk("sctp_filter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("sctp_filter_profile", sv)
				diags = append(diags, e)
			}
			obj.SctpFilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_filter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssh_filter_profile", sv)
				diags = append(diags, e)
			}
			obj.SshFilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_ssh_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_ssh_profile", sv)
				diags = append(diags, e)
			}
			obj.SslSshProfile = &v2
		}
	}
	if v1, ok := d.GetOk("videofilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("videofilter_profile", sv)
				diags = append(diags, e)
			}
			obj.VideofilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("voip_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("voip_profile", sv)
				diags = append(diags, e)
			}
			obj.VoipProfile = &v2
		}
	}
	if v1, ok := d.GetOk("waf_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("waf_profile", sv)
				diags = append(diags, e)
			}
			obj.WafProfile = &v2
		}
	}
	if v1, ok := d.GetOk("webfilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webfilter_profile", sv)
				diags = append(diags, e)
			}
			obj.WebfilterProfile = &v2
		}
	}
	return &obj, diags
}
