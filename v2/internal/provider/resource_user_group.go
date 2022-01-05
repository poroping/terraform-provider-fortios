// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure user groups.

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

func resourceUserGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure user groups.",

		CreateContext: resourceUserGroupCreate,
		ReadContext:   resourceUserGroupRead,
		UpdateContext: resourceUserGroupUpdate,
		DeleteContext: resourceUserGroupDelete,

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
			"auth_concurrent_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable overriding the global number of concurrent authentication sessions for this user group.",
				Optional:    true,
				Computed:    true,
			},
			"auth_concurrent_value": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),

				Description: "Maximum number of concurrent authenticated connections per user (0 - 100).",
				Optional:    true,
				Computed:    true,
			},
			"authtimeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 43200),

				Description: "Authentication timeout in minutes for this user group. 0 to use the global user setting auth-timeout.",
				Optional:    true,
				Computed:    true,
			},
			"company": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"optional", "mandatory", "disabled"}, false),

				Description: "Set the action for the company guest user field.",
				Optional:    true,
				Computed:    true,
			},
			"email": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable the guest user email address field.",
				Optional:    true,
				Computed:    true,
			},
			"expire": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 31536000),

				Description: "Time in seconds before guest user accounts expire. (1 - 31536000 sec)",
				Optional:    true,
				Computed:    true,
			},
			"expire_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"immediately", "first-successful-login"}, false),

				Description: "Determine when the expiration countdown begins.",
				Optional:    true,
				Computed:    true,
			},
			"group_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"firewall", "fsso-service", "rsso", "guest"}, false),

				Description: "Set the group to be for firewall authentication, FSSO, RSSO, or guest users.",
				Optional:    true,
				Computed:    true,
			},
			"guest": {
				Type:        schema.TypeList,
				Description: "Guest User.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Comment.",
							Optional:    true,
							Computed:    true,
						},
						"company": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Set the action for the company guest user field.",
							Optional:    true,
							Computed:    true,
						},
						"email": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Email.",
							Optional:    true,
							Computed:    true,
						},
						"expiration": {
							Type: schema.TypeString,

							Description: "Expire time.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Guest ID.",
							Optional:    true,
							Computed:    true,
						},
						"mobile_phone": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Mobile phone.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Guest name.",
							Optional:    true,
							Computed:    true,
						},
						"password": {
							Type: schema.TypeString,

							Description: "Guest password.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"sponsor": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Set the action for the sponsor guest user field.",
							Optional:    true,
							Computed:    true,
						},
						"user_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Guest ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"http_digest_realm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Realm attribute for MD5-digest authentication.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Group ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"match": {
				Type:        schema.TypeList,
				Description: "Group matches.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "Name of matching user or group on remote authentication server.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"server_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Name of remote auth server.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"max_accounts": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1024),

				Description: "Maximum number of guest accounts that can be created for this group (0 means unlimited).",
				Optional:    true,
				Computed:    true,
			},
			"member": {
				Type:        schema.TypeList,
				Description: "Names of users, peers, LDAP severs, or RADIUS servers to add to the user group.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "Group member name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"mobile_phone": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable the guest user mobile phone number field.",
				Optional:    true,
				Computed:    true,
			},
			"multiple_guest_add": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable addition of multiple guests.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Group name.",
				ForceNew:    true,
				Required:    true,
			},
			"password": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto-generate", "specify", "disable"}, false),

				Description: "Guest user password type.",
				Optional:    true,
				Computed:    true,
			},
			"sms_custom_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SMS server.",
				Optional:    true,
				Computed:    true,
			},
			"sms_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fortiguard", "custom"}, false),

				Description: "Send SMS through FortiGuard or other external server.",
				Optional:    true,
				Computed:    true,
			},
			"sponsor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"optional", "mandatory", "disabled"}, false),

				Description: "Set the action for the sponsor guest user field.",
				Optional:    true,
				Computed:    true,
			},
			"sso_attribute_value": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),

				Description: "Name of the RADIUS user group that this local user group represents.",
				Optional:    true,
				Computed:    true,
			},
			"user_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"email", "auto-generate", "specify"}, false),

				Description: "Guest user ID type.",
				Optional:    true,
				Computed:    true,
			},
			"user_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable the guest user name entry.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserGroup resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserGroup(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserGroup")
	}

	return resourceUserGroupRead(ctx, d, meta)
}

func resourceUserGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserGroup(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserGroup resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserGroup")
	}

	return resourceUserGroupRead(ctx, d, meta)
}

func resourceUserGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserGroup resource: %v", err)
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

	diags := refreshObjectUserGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenUserGroupGuest(v *[]models.UserGroupGuest, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Comment; tmp != nil {
				v["comment"] = *tmp
			}

			if tmp := cfg.Company; tmp != nil {
				v["company"] = *tmp
			}

			if tmp := cfg.Email; tmp != nil {
				v["email"] = *tmp
			}

			if tmp := cfg.Expiration; tmp != nil {
				v["expiration"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.MobilePhone; tmp != nil {
				v["mobile_phone"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Password; tmp != nil {
				v["password"] = *tmp
			}

			if tmp := cfg.Sponsor; tmp != nil {
				v["sponsor"] = *tmp
			}

			if tmp := cfg.UserId; tmp != nil {
				v["user_id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenUserGroupMatch(v *[]models.UserGroupMatch, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.GroupName; tmp != nil {
				v["group_name"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.ServerName; tmp != nil {
				v["server_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenUserGroupMember(v *[]models.UserGroupMember, sort bool) interface{} {
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

func refreshObjectUserGroup(d *schema.ResourceData, o *models.UserGroup, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthConcurrentOverride != nil {
		v := *o.AuthConcurrentOverride

		if err = d.Set("auth_concurrent_override", v); err != nil {
			return diag.Errorf("error reading auth_concurrent_override: %v", err)
		}
	}

	if o.AuthConcurrentValue != nil {
		v := *o.AuthConcurrentValue

		if err = d.Set("auth_concurrent_value", v); err != nil {
			return diag.Errorf("error reading auth_concurrent_value: %v", err)
		}
	}

	if o.Authtimeout != nil {
		v := *o.Authtimeout

		if err = d.Set("authtimeout", v); err != nil {
			return diag.Errorf("error reading authtimeout: %v", err)
		}
	}

	if o.Company != nil {
		v := *o.Company

		if err = d.Set("company", v); err != nil {
			return diag.Errorf("error reading company: %v", err)
		}
	}

	if o.Email != nil {
		v := *o.Email

		if err = d.Set("email", v); err != nil {
			return diag.Errorf("error reading email: %v", err)
		}
	}

	if o.Expire != nil {
		v := *o.Expire

		if err = d.Set("expire", v); err != nil {
			return diag.Errorf("error reading expire: %v", err)
		}
	}

	if o.ExpireType != nil {
		v := *o.ExpireType

		if err = d.Set("expire_type", v); err != nil {
			return diag.Errorf("error reading expire_type: %v", err)
		}
	}

	if o.GroupType != nil {
		v := *o.GroupType

		if err = d.Set("group_type", v); err != nil {
			return diag.Errorf("error reading group_type: %v", err)
		}
	}

	if o.Guest != nil {
		if err = d.Set("guest", flattenUserGroupGuest(o.Guest, sort)); err != nil {
			return diag.Errorf("error reading guest: %v", err)
		}
	}

	if o.HttpDigestRealm != nil {
		v := *o.HttpDigestRealm

		if err = d.Set("http_digest_realm", v); err != nil {
			return diag.Errorf("error reading http_digest_realm: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Match != nil {
		if err = d.Set("match", flattenUserGroupMatch(o.Match, sort)); err != nil {
			return diag.Errorf("error reading match: %v", err)
		}
	}

	if o.MaxAccounts != nil {
		v := *o.MaxAccounts

		if err = d.Set("max_accounts", v); err != nil {
			return diag.Errorf("error reading max_accounts: %v", err)
		}
	}

	if o.Member != nil {
		if err = d.Set("member", flattenUserGroupMember(o.Member, sort)); err != nil {
			return diag.Errorf("error reading member: %v", err)
		}
	}

	if o.MobilePhone != nil {
		v := *o.MobilePhone

		if err = d.Set("mobile_phone", v); err != nil {
			return diag.Errorf("error reading mobile_phone: %v", err)
		}
	}

	if o.MultipleGuestAdd != nil {
		v := *o.MultipleGuestAdd

		if err = d.Set("multiple_guest_add", v); err != nil {
			return diag.Errorf("error reading multiple_guest_add: %v", err)
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

	if o.SmsCustomServer != nil {
		v := *o.SmsCustomServer

		if err = d.Set("sms_custom_server", v); err != nil {
			return diag.Errorf("error reading sms_custom_server: %v", err)
		}
	}

	if o.SmsServer != nil {
		v := *o.SmsServer

		if err = d.Set("sms_server", v); err != nil {
			return diag.Errorf("error reading sms_server: %v", err)
		}
	}

	if o.Sponsor != nil {
		v := *o.Sponsor

		if err = d.Set("sponsor", v); err != nil {
			return diag.Errorf("error reading sponsor: %v", err)
		}
	}

	if o.SsoAttributeValue != nil {
		v := *o.SsoAttributeValue

		if err = d.Set("sso_attribute_value", v); err != nil {
			return diag.Errorf("error reading sso_attribute_value: %v", err)
		}
	}

	if o.UserId != nil {
		v := *o.UserId

		if err = d.Set("user_id", v); err != nil {
			return diag.Errorf("error reading user_id: %v", err)
		}
	}

	if o.UserName != nil {
		v := *o.UserName

		if err = d.Set("user_name", v); err != nil {
			return diag.Errorf("error reading user_name: %v", err)
		}
	}

	return nil
}

func expandUserGroupGuest(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserGroupGuest, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserGroupGuest

	for i := range l {
		tmp := models.UserGroupGuest{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.comment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comment = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.company", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Company = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.email", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Email = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.expiration", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Expiration = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mobile_phone", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MobilePhone = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Password = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sponsor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Sponsor = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.user_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UserId = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandUserGroupMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserGroupMatch, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserGroupMatch

	for i := range l {
		tmp := models.UserGroupMatch{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.group_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.GroupName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ServerName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandUserGroupMember(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserGroupMember, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserGroupMember

	for i := range l {
		tmp := models.UserGroupMember{}
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

func getObjectUserGroup(d *schema.ResourceData, sv string) (*models.UserGroup, diag.Diagnostics) {
	obj := models.UserGroup{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth_concurrent_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_concurrent_override", sv)
				diags = append(diags, e)
			}
			obj.AuthConcurrentOverride = &v2
		}
	}
	if v1, ok := d.GetOk("auth_concurrent_value"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_concurrent_value", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthConcurrentValue = &tmp
		}
	}
	if v1, ok := d.GetOk("authtimeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authtimeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Authtimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("company"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("company", sv)
				diags = append(diags, e)
			}
			obj.Company = &v2
		}
	}
	if v1, ok := d.GetOk("email"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("email", sv)
				diags = append(diags, e)
			}
			obj.Email = &v2
		}
	}
	if v1, ok := d.GetOk("expire"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("expire", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Expire = &tmp
		}
	}
	if v1, ok := d.GetOk("expire_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("expire_type", sv)
				diags = append(diags, e)
			}
			obj.ExpireType = &v2
		}
	}
	if v1, ok := d.GetOk("group_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_type", sv)
				diags = append(diags, e)
			}
			obj.GroupType = &v2
		}
	}
	if v, ok := d.GetOk("guest"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("guest", sv)
			diags = append(diags, e)
		}
		t, err := expandUserGroupGuest(d, v, "guest", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Guest = t
		}
	} else if d.HasChange("guest") {
		old, new := d.GetChange("guest")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Guest = &[]models.UserGroupGuest{}
		}
	}
	if v1, ok := d.GetOk("http_digest_realm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_digest_realm", sv)
				diags = append(diags, e)
			}
			obj.HttpDigestRealm = &v2
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
	if v, ok := d.GetOk("match"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("match", sv)
			diags = append(diags, e)
		}
		t, err := expandUserGroupMatch(d, v, "match", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Match = t
		}
	} else if d.HasChange("match") {
		old, new := d.GetChange("match")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Match = &[]models.UserGroupMatch{}
		}
	}
	if v1, ok := d.GetOk("max_accounts"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_accounts", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxAccounts = &tmp
		}
	}
	if v, ok := d.GetOk("member"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("member", sv)
			diags = append(diags, e)
		}
		t, err := expandUserGroupMember(d, v, "member", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Member = t
		}
	} else if d.HasChange("member") {
		old, new := d.GetChange("member")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Member = &[]models.UserGroupMember{}
		}
	}
	if v1, ok := d.GetOk("mobile_phone"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mobile_phone", sv)
				diags = append(diags, e)
			}
			obj.MobilePhone = &v2
		}
	}
	if v1, ok := d.GetOk("multiple_guest_add"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multiple_guest_add", sv)
				diags = append(diags, e)
			}
			obj.MultipleGuestAdd = &v2
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
	if v1, ok := d.GetOk("sms_custom_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sms_custom_server", sv)
				diags = append(diags, e)
			}
			obj.SmsCustomServer = &v2
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
	if v1, ok := d.GetOk("sponsor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sponsor", sv)
				diags = append(diags, e)
			}
			obj.Sponsor = &v2
		}
	}
	if v1, ok := d.GetOk("sso_attribute_value"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sso_attribute_value", sv)
				diags = append(diags, e)
			}
			obj.SsoAttributeValue = &v2
		}
	}
	if v1, ok := d.GetOk("user_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_id", sv)
				diags = append(diags, e)
			}
			obj.UserId = &v2
		}
	}
	if v1, ok := d.GetOk("user_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_name", sv)
				diags = append(diags, e)
			}
			obj.UserName = &v2
		}
	}
	return &obj, diags
}
