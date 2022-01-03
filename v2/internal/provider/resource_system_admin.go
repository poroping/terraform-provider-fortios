// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure admin users.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceSystemAdmin() *schema.Resource {
	return &schema.Resource{
		Description: "Configure admin users.",

		CreateContext: resourceSystemAdminCreate,
		ReadContext:   resourceSystemAdminRead,
		UpdateContext: resourceSystemAdminUpdate,
		DeleteContext: resourceSystemAdminDelete,

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
			"accprofile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Access profile for this administrator. Access profiles control administrator access to FortiGate features.",
				Optional:    true,
				Computed:    true,
			},
			"accprofile_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to use the name of an access profile provided by the remote authentication server to control the FortiGate features that this administrator can access.",
				Optional:    true,
				Computed:    true,
			},
			"allow_remove_admin_session": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allow admin session to be removed by privileged admin users.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"email_to": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "This administrator's email address.",
				Optional:    true,
				Computed:    true,
			},
			"force_password_change": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable force password change on next login.",
				Optional:    true,
				Computed:    true,
			},
			"fortitoken": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),

				Description: "This administrator's FortiToken serial number.",
				Optional:    true,
				Computed:    true,
			},
			"guest_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable guest authentication.",
				Optional:    true,
				Computed:    true,
			},
			"guest_lang": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Guest management portal language.",
				Optional:    true,
				Computed:    true,
			},
			"guest_usergroups": {
				Type:        schema.TypeList,
				Description: "Select guest user groups.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Select guest user groups.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ip6_trusthost1": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"ip6_trusthost10": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"ip6_trusthost2": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"ip6_trusthost3": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"ip6_trusthost4": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"ip6_trusthost5": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"ip6_trusthost6": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"ip6_trusthost7": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"ip6_trusthost8": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"ip6_trusthost9": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "User name.",
				ForceNew:    true,
				Required:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Admin user password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"password_expire": {
				Type: schema.TypeString,

				Description: "Password expire time.",
				Optional:    true,
				Computed:    true,
			},
			"peer_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Set to enable peer certificate authentication (for HTTPS admin access).",
				Optional:    true,
				Computed:    true,
			},
			"peer_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of peer group defined under config user group which has PKI members. Used for peer certificate authentication (for HTTPS admin access).",
				Optional:    true,
				Computed:    true,
			},
			"radius_vdom_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to use the names of VDOMs provided by the remote authentication server to control the VDOMs that this administrator can access.",
				Optional:    true,
				Computed:    true,
			},
			"remote_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable authentication using a remote RADIUS, LDAP, or TACACS+ server.",
				Optional:    true,
				Computed:    true,
			},
			"remote_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "User group name used for remote auth.",
				Optional:    true,
				Computed:    true,
			},
			"schedule": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Firewall schedule used to restrict when the administrator can log in. No schedule means no restrictions.",
				Optional:    true,
				Computed:    true,
			},
			"sms_custom_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Custom SMS server to send SMS messages to.",
				Optional:    true,
				Computed:    true,
			},
			"sms_phone": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Phone number on which the administrator receives SMS messages.",
				Optional:    true,
				Computed:    true,
			},
			"sms_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fortiguard", "custom"}, false),

				Description: "Send SMS messages using the FortiGuard SMS server or a custom server.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Select the certificate to be used by the FortiGate for authentication with an SSH client.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_public_key1": {
				Type: schema.TypeString,

				Description: "Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_public_key2": {
				Type: schema.TypeString,

				Description: "Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_public_key3": {
				Type: schema.TypeString,

				Description: "Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.",
				Optional:    true,
				Computed:    true,
			},
			"trusthost1": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"trusthost10": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"trusthost2": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"trusthost3": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"trusthost4": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"trusthost5": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"trusthost6": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"trusthost7": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"trusthost8": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"trusthost9": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"two_factor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "fortitoken", "fortitoken-cloud", "email", "sms"}, false),

				Description: "Enable/disable two-factor authentication.",
				Optional:    true,
				Computed:    true,
			},
			"two_factor_authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fortitoken", "email", "sms"}, false),

				Description: "Authentication method by FortiToken Cloud.",
				Optional:    true,
				Computed:    true,
			},
			"two_factor_notification": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"email", "sms"}, false),

				Description: "Notification method for user activation by FortiToken Cloud.",
				Optional:    true,
				Computed:    true,
			},
			"vdom": {
				Type:        schema.TypeList,
				Description: "Virtual domain(s) that the administrator can access.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Virtual domain name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"wildcard": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable wildcard RADIUS authentication.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemAdminCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemAdmin resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemAdmin(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemAdmin(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAdmin")
	}

	return resourceSystemAdminRead(ctx, d, meta)
}

func resourceSystemAdminUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAdmin(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemAdmin(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemAdmin resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAdmin")
	}

	return resourceSystemAdminRead(ctx, d, meta)
}

func resourceSystemAdminDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemAdmin(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAdmin(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAdmin resource: %v", err)
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

	diags := refreshObjectSystemAdmin(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemAdminGuestUsergroups(v *[]models.SystemAdminGuestUsergroups, sort bool) interface{} {
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

func flattenSystemAdminVdom(v *[]models.SystemAdminVdom, sort bool) interface{} {
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

func refreshObjectSystemAdmin(d *schema.ResourceData, o *models.SystemAdmin, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Accprofile != nil {
		v := *o.Accprofile

		if err = d.Set("accprofile", v); err != nil {
			return diag.Errorf("error reading accprofile: %v", err)
		}
	}

	if o.AccprofileOverride != nil {
		v := *o.AccprofileOverride

		if err = d.Set("accprofile_override", v); err != nil {
			return diag.Errorf("error reading accprofile_override: %v", err)
		}
	}

	if o.AllowRemoveAdminSession != nil {
		v := *o.AllowRemoveAdminSession

		if err = d.Set("allow_remove_admin_session", v); err != nil {
			return diag.Errorf("error reading allow_remove_admin_session: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.EmailTo != nil {
		v := *o.EmailTo

		if err = d.Set("email_to", v); err != nil {
			return diag.Errorf("error reading email_to: %v", err)
		}
	}

	if o.ForcePasswordChange != nil {
		v := *o.ForcePasswordChange

		if err = d.Set("force_password_change", v); err != nil {
			return diag.Errorf("error reading force_password_change: %v", err)
		}
	}

	if o.Fortitoken != nil {
		v := *o.Fortitoken

		if err = d.Set("fortitoken", v); err != nil {
			return diag.Errorf("error reading fortitoken: %v", err)
		}
	}

	if o.GuestAuth != nil {
		v := *o.GuestAuth

		if err = d.Set("guest_auth", v); err != nil {
			return diag.Errorf("error reading guest_auth: %v", err)
		}
	}

	if o.GuestLang != nil {
		v := *o.GuestLang

		if err = d.Set("guest_lang", v); err != nil {
			return diag.Errorf("error reading guest_lang: %v", err)
		}
	}

	if o.GuestUsergroups != nil {
		if err = d.Set("guest_usergroups", flattenSystemAdminGuestUsergroups(o.GuestUsergroups, sort)); err != nil {
			return diag.Errorf("error reading guest_usergroups: %v", err)
		}
	}

	if o.Ip6Trusthost1 != nil {
		v := *o.Ip6Trusthost1

		if err = d.Set("ip6_trusthost1", v); err != nil {
			return diag.Errorf("error reading ip6_trusthost1: %v", err)
		}
	}

	if o.Ip6Trusthost10 != nil {
		v := *o.Ip6Trusthost10

		if err = d.Set("ip6_trusthost10", v); err != nil {
			return diag.Errorf("error reading ip6_trusthost10: %v", err)
		}
	}

	if o.Ip6Trusthost2 != nil {
		v := *o.Ip6Trusthost2

		if err = d.Set("ip6_trusthost2", v); err != nil {
			return diag.Errorf("error reading ip6_trusthost2: %v", err)
		}
	}

	if o.Ip6Trusthost3 != nil {
		v := *o.Ip6Trusthost3

		if err = d.Set("ip6_trusthost3", v); err != nil {
			return diag.Errorf("error reading ip6_trusthost3: %v", err)
		}
	}

	if o.Ip6Trusthost4 != nil {
		v := *o.Ip6Trusthost4

		if err = d.Set("ip6_trusthost4", v); err != nil {
			return diag.Errorf("error reading ip6_trusthost4: %v", err)
		}
	}

	if o.Ip6Trusthost5 != nil {
		v := *o.Ip6Trusthost5

		if err = d.Set("ip6_trusthost5", v); err != nil {
			return diag.Errorf("error reading ip6_trusthost5: %v", err)
		}
	}

	if o.Ip6Trusthost6 != nil {
		v := *o.Ip6Trusthost6

		if err = d.Set("ip6_trusthost6", v); err != nil {
			return diag.Errorf("error reading ip6_trusthost6: %v", err)
		}
	}

	if o.Ip6Trusthost7 != nil {
		v := *o.Ip6Trusthost7

		if err = d.Set("ip6_trusthost7", v); err != nil {
			return diag.Errorf("error reading ip6_trusthost7: %v", err)
		}
	}

	if o.Ip6Trusthost8 != nil {
		v := *o.Ip6Trusthost8

		if err = d.Set("ip6_trusthost8", v); err != nil {
			return diag.Errorf("error reading ip6_trusthost8: %v", err)
		}
	}

	if o.Ip6Trusthost9 != nil {
		v := *o.Ip6Trusthost9

		if err = d.Set("ip6_trusthost9", v); err != nil {
			return diag.Errorf("error reading ip6_trusthost9: %v", err)
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

	if o.PasswordExpire != nil {
		v := *o.PasswordExpire

		if err = d.Set("password_expire", v); err != nil {
			return diag.Errorf("error reading password_expire: %v", err)
		}
	}

	if o.PeerAuth != nil {
		v := *o.PeerAuth

		if err = d.Set("peer_auth", v); err != nil {
			return diag.Errorf("error reading peer_auth: %v", err)
		}
	}

	if o.PeerGroup != nil {
		v := *o.PeerGroup

		if err = d.Set("peer_group", v); err != nil {
			return diag.Errorf("error reading peer_group: %v", err)
		}
	}

	if o.RadiusVdomOverride != nil {
		v := *o.RadiusVdomOverride

		if err = d.Set("radius_vdom_override", v); err != nil {
			return diag.Errorf("error reading radius_vdom_override: %v", err)
		}
	}

	if o.RemoteAuth != nil {
		v := *o.RemoteAuth

		if err = d.Set("remote_auth", v); err != nil {
			return diag.Errorf("error reading remote_auth: %v", err)
		}
	}

	if o.RemoteGroup != nil {
		v := *o.RemoteGroup

		if err = d.Set("remote_group", v); err != nil {
			return diag.Errorf("error reading remote_group: %v", err)
		}
	}

	if o.Schedule != nil {
		v := *o.Schedule

		if err = d.Set("schedule", v); err != nil {
			return diag.Errorf("error reading schedule: %v", err)
		}
	}

	if o.SmsCustomServer != nil {
		v := *o.SmsCustomServer

		if err = d.Set("sms_custom_server", v); err != nil {
			return diag.Errorf("error reading sms_custom_server: %v", err)
		}
	}

	if o.SmsPhone != nil {
		v := *o.SmsPhone

		if err = d.Set("sms_phone", v); err != nil {
			return diag.Errorf("error reading sms_phone: %v", err)
		}
	}

	if o.SmsServer != nil {
		v := *o.SmsServer

		if err = d.Set("sms_server", v); err != nil {
			return diag.Errorf("error reading sms_server: %v", err)
		}
	}

	if o.SshCertificate != nil {
		v := *o.SshCertificate

		if err = d.Set("ssh_certificate", v); err != nil {
			return diag.Errorf("error reading ssh_certificate: %v", err)
		}
	}

	if o.SshPublicKey1 != nil {
		v := *o.SshPublicKey1

		if err = d.Set("ssh_public_key1", v); err != nil {
			return diag.Errorf("error reading ssh_public_key1: %v", err)
		}
	}

	if o.SshPublicKey2 != nil {
		v := *o.SshPublicKey2

		if err = d.Set("ssh_public_key2", v); err != nil {
			return diag.Errorf("error reading ssh_public_key2: %v", err)
		}
	}

	if o.SshPublicKey3 != nil {
		v := *o.SshPublicKey3

		if err = d.Set("ssh_public_key3", v); err != nil {
			return diag.Errorf("error reading ssh_public_key3: %v", err)
		}
	}

	if o.Trusthost1 != nil {
		v := *o.Trusthost1
		if current, ok := d.GetOk("trusthost1"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trusthost1", v); err != nil {
			return diag.Errorf("error reading trusthost1: %v", err)
		}
	}

	if o.Trusthost10 != nil {
		v := *o.Trusthost10
		if current, ok := d.GetOk("trusthost10"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trusthost10", v); err != nil {
			return diag.Errorf("error reading trusthost10: %v", err)
		}
	}

	if o.Trusthost2 != nil {
		v := *o.Trusthost2
		if current, ok := d.GetOk("trusthost2"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trusthost2", v); err != nil {
			return diag.Errorf("error reading trusthost2: %v", err)
		}
	}

	if o.Trusthost3 != nil {
		v := *o.Trusthost3
		if current, ok := d.GetOk("trusthost3"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trusthost3", v); err != nil {
			return diag.Errorf("error reading trusthost3: %v", err)
		}
	}

	if o.Trusthost4 != nil {
		v := *o.Trusthost4
		if current, ok := d.GetOk("trusthost4"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trusthost4", v); err != nil {
			return diag.Errorf("error reading trusthost4: %v", err)
		}
	}

	if o.Trusthost5 != nil {
		v := *o.Trusthost5
		if current, ok := d.GetOk("trusthost5"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trusthost5", v); err != nil {
			return diag.Errorf("error reading trusthost5: %v", err)
		}
	}

	if o.Trusthost6 != nil {
		v := *o.Trusthost6
		if current, ok := d.GetOk("trusthost6"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trusthost6", v); err != nil {
			return diag.Errorf("error reading trusthost6: %v", err)
		}
	}

	if o.Trusthost7 != nil {
		v := *o.Trusthost7
		if current, ok := d.GetOk("trusthost7"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trusthost7", v); err != nil {
			return diag.Errorf("error reading trusthost7: %v", err)
		}
	}

	if o.Trusthost8 != nil {
		v := *o.Trusthost8
		if current, ok := d.GetOk("trusthost8"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trusthost8", v); err != nil {
			return diag.Errorf("error reading trusthost8: %v", err)
		}
	}

	if o.Trusthost9 != nil {
		v := *o.Trusthost9
		if current, ok := d.GetOk("trusthost9"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trusthost9", v); err != nil {
			return diag.Errorf("error reading trusthost9: %v", err)
		}
	}

	if o.TwoFactor != nil {
		v := *o.TwoFactor

		if err = d.Set("two_factor", v); err != nil {
			return diag.Errorf("error reading two_factor: %v", err)
		}
	}

	if o.TwoFactorAuthentication != nil {
		v := *o.TwoFactorAuthentication

		if err = d.Set("two_factor_authentication", v); err != nil {
			return diag.Errorf("error reading two_factor_authentication: %v", err)
		}
	}

	if o.TwoFactorNotification != nil {
		v := *o.TwoFactorNotification

		if err = d.Set("two_factor_notification", v); err != nil {
			return diag.Errorf("error reading two_factor_notification: %v", err)
		}
	}

	if o.Vdom != nil {
		if err = d.Set("vdom", flattenSystemAdminVdom(o.Vdom, sort)); err != nil {
			return diag.Errorf("error reading vdom: %v", err)
		}
	}

	if o.Wildcard != nil {
		v := *o.Wildcard

		if err = d.Set("wildcard", v); err != nil {
			return diag.Errorf("error reading wildcard: %v", err)
		}
	}

	return nil
}

func expandSystemAdminGuestUsergroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAdminGuestUsergroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAdminGuestUsergroups

	for i := range l {
		tmp := models.SystemAdminGuestUsergroups{}
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

func expandSystemAdminVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAdminVdom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAdminVdom

	for i := range l {
		tmp := models.SystemAdminVdom{}
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

func getObjectSystemAdmin(d *schema.ResourceData, sv string) (*models.SystemAdmin, diag.Diagnostics) {
	obj := models.SystemAdmin{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("accprofile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("accprofile", sv)
				diags = append(diags, e)
			}
			obj.Accprofile = &v2
		}
	}
	if v1, ok := d.GetOk("accprofile_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("accprofile_override", sv)
				diags = append(diags, e)
			}
			obj.AccprofileOverride = &v2
		}
	}
	if v1, ok := d.GetOk("allow_remove_admin_session"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allow_remove_admin_session", sv)
				diags = append(diags, e)
			}
			obj.AllowRemoveAdminSession = &v2
		}
	}
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v1, ok := d.GetOk("email_to"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("email_to", sv)
				diags = append(diags, e)
			}
			obj.EmailTo = &v2
		}
	}
	if v1, ok := d.GetOk("force_password_change"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("force_password_change", sv)
				diags = append(diags, e)
			}
			obj.ForcePasswordChange = &v2
		}
	}
	if v1, ok := d.GetOk("fortitoken"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortitoken", sv)
				diags = append(diags, e)
			}
			obj.Fortitoken = &v2
		}
	}
	if v1, ok := d.GetOk("guest_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("guest_auth", sv)
				diags = append(diags, e)
			}
			obj.GuestAuth = &v2
		}
	}
	if v1, ok := d.GetOk("guest_lang"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("guest_lang", sv)
				diags = append(diags, e)
			}
			obj.GuestLang = &v2
		}
	}
	if v, ok := d.GetOk("guest_usergroups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("guest_usergroups", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAdminGuestUsergroups(d, v, "guest_usergroups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.GuestUsergroups = t
		}
	} else if d.HasChange("guest_usergroups") {
		old, new := d.GetChange("guest_usergroups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.GuestUsergroups = &[]models.SystemAdminGuestUsergroups{}
		}
	}
	if v1, ok := d.GetOk("ip6_trusthost1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_trusthost1", sv)
				diags = append(diags, e)
			}
			obj.Ip6Trusthost1 = &v2
		}
	}
	if v1, ok := d.GetOk("ip6_trusthost10"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_trusthost10", sv)
				diags = append(diags, e)
			}
			obj.Ip6Trusthost10 = &v2
		}
	}
	if v1, ok := d.GetOk("ip6_trusthost2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_trusthost2", sv)
				diags = append(diags, e)
			}
			obj.Ip6Trusthost2 = &v2
		}
	}
	if v1, ok := d.GetOk("ip6_trusthost3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_trusthost3", sv)
				diags = append(diags, e)
			}
			obj.Ip6Trusthost3 = &v2
		}
	}
	if v1, ok := d.GetOk("ip6_trusthost4"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_trusthost4", sv)
				diags = append(diags, e)
			}
			obj.Ip6Trusthost4 = &v2
		}
	}
	if v1, ok := d.GetOk("ip6_trusthost5"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_trusthost5", sv)
				diags = append(diags, e)
			}
			obj.Ip6Trusthost5 = &v2
		}
	}
	if v1, ok := d.GetOk("ip6_trusthost6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_trusthost6", sv)
				diags = append(diags, e)
			}
			obj.Ip6Trusthost6 = &v2
		}
	}
	if v1, ok := d.GetOk("ip6_trusthost7"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_trusthost7", sv)
				diags = append(diags, e)
			}
			obj.Ip6Trusthost7 = &v2
		}
	}
	if v1, ok := d.GetOk("ip6_trusthost8"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_trusthost8", sv)
				diags = append(diags, e)
			}
			obj.Ip6Trusthost8 = &v2
		}
	}
	if v1, ok := d.GetOk("ip6_trusthost9"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_trusthost9", sv)
				diags = append(diags, e)
			}
			obj.Ip6Trusthost9 = &v2
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
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("password_expire"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password_expire", sv)
				diags = append(diags, e)
			}
			obj.PasswordExpire = &v2
		}
	}
	if v1, ok := d.GetOk("peer_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peer_auth", sv)
				diags = append(diags, e)
			}
			obj.PeerAuth = &v2
		}
	}
	if v1, ok := d.GetOk("peer_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peer_group", sv)
				diags = append(diags, e)
			}
			obj.PeerGroup = &v2
		}
	}
	if v1, ok := d.GetOk("radius_vdom_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radius_vdom_override", sv)
				diags = append(diags, e)
			}
			obj.RadiusVdomOverride = &v2
		}
	}
	if v1, ok := d.GetOk("remote_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remote_auth", sv)
				diags = append(diags, e)
			}
			obj.RemoteAuth = &v2
		}
	}
	if v1, ok := d.GetOk("remote_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remote_group", sv)
				diags = append(diags, e)
			}
			obj.RemoteGroup = &v2
		}
	}
	if v1, ok := d.GetOk("schedule"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("schedule", sv)
				diags = append(diags, e)
			}
			obj.Schedule = &v2
		}
	}
	if v1, ok := d.GetOk("sms_custom_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sms_custom_server", sv)
				diags = append(diags, e)
			}
			obj.SmsCustomServer = &v2
		}
	}
	if v1, ok := d.GetOk("sms_phone"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sms_phone", sv)
				diags = append(diags, e)
			}
			obj.SmsPhone = &v2
		}
	}
	if v1, ok := d.GetOk("sms_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sms_server", sv)
				diags = append(diags, e)
			}
			obj.SmsServer = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssh_certificate", sv)
				diags = append(diags, e)
			}
			obj.SshCertificate = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_public_key1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssh_public_key1", sv)
				diags = append(diags, e)
			}
			obj.SshPublicKey1 = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_public_key2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssh_public_key2", sv)
				diags = append(diags, e)
			}
			obj.SshPublicKey2 = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_public_key3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssh_public_key3", sv)
				diags = append(diags, e)
			}
			obj.SshPublicKey3 = &v2
		}
	}
	if v1, ok := d.GetOk("trusthost1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trusthost1", sv)
				diags = append(diags, e)
			}
			obj.Trusthost1 = &v2
		}
	}
	if v1, ok := d.GetOk("trusthost10"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trusthost10", sv)
				diags = append(diags, e)
			}
			obj.Trusthost10 = &v2
		}
	}
	if v1, ok := d.GetOk("trusthost2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trusthost2", sv)
				diags = append(diags, e)
			}
			obj.Trusthost2 = &v2
		}
	}
	if v1, ok := d.GetOk("trusthost3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trusthost3", sv)
				diags = append(diags, e)
			}
			obj.Trusthost3 = &v2
		}
	}
	if v1, ok := d.GetOk("trusthost4"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trusthost4", sv)
				diags = append(diags, e)
			}
			obj.Trusthost4 = &v2
		}
	}
	if v1, ok := d.GetOk("trusthost5"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trusthost5", sv)
				diags = append(diags, e)
			}
			obj.Trusthost5 = &v2
		}
	}
	if v1, ok := d.GetOk("trusthost6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trusthost6", sv)
				diags = append(diags, e)
			}
			obj.Trusthost6 = &v2
		}
	}
	if v1, ok := d.GetOk("trusthost7"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trusthost7", sv)
				diags = append(diags, e)
			}
			obj.Trusthost7 = &v2
		}
	}
	if v1, ok := d.GetOk("trusthost8"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trusthost8", sv)
				diags = append(diags, e)
			}
			obj.Trusthost8 = &v2
		}
	}
	if v1, ok := d.GetOk("trusthost9"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trusthost9", sv)
				diags = append(diags, e)
			}
			obj.Trusthost9 = &v2
		}
	}
	if v1, ok := d.GetOk("two_factor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor", sv)
				diags = append(diags, e)
			}
			obj.TwoFactor = &v2
		}
	}
	if v1, ok := d.GetOk("two_factor_authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor_authentication", sv)
				diags = append(diags, e)
			}
			obj.TwoFactorAuthentication = &v2
		}
	}
	if v1, ok := d.GetOk("two_factor_notification"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor_notification", sv)
				diags = append(diags, e)
			}
			obj.TwoFactorNotification = &v2
		}
	}
	if v, ok := d.GetOk("vdom"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("vdom", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAdminVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Vdom = t
		}
	} else if d.HasChange("vdom") {
		old, new := d.GetChange("vdom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Vdom = &[]models.SystemAdminVdom{}
		}
	}
	if v1, ok := d.GetOk("wildcard"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wildcard", sv)
				diags = append(diags, e)
			}
			obj.Wildcard = &v2
		}
	}
	return &obj, diags
}
