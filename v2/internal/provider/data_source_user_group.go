// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure user groups.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceUserGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure user groups.",

		ReadContext: dataSourceUserGroupRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auth_concurrent_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable overriding the global number of concurrent authentication sessions for this user group.",
				Computed:    true,
			},
			"auth_concurrent_value": {
				Type:        schema.TypeInt,
				Description: "Maximum number of concurrent authenticated connections per user (0 - 100).",
				Computed:    true,
			},
			"authtimeout": {
				Type:        schema.TypeInt,
				Description: "Authentication timeout in minutes for this user group. 0 to use the global user setting auth-timeout.",
				Computed:    true,
			},
			"company": {
				Type:        schema.TypeString,
				Description: "Set the action for the company guest user field.",
				Computed:    true,
			},
			"email": {
				Type:        schema.TypeString,
				Description: "Enable/disable the guest user email address field.",
				Computed:    true,
			},
			"expire": {
				Type:        schema.TypeInt,
				Description: "Time in seconds before guest user accounts expire (1 - 31536000).",
				Computed:    true,
			},
			"expire_type": {
				Type:        schema.TypeString,
				Description: "Determine when the expiration countdown begins.",
				Computed:    true,
			},
			"group_type": {
				Type:        schema.TypeString,
				Description: "Set the group to be for firewall authentication, FSSO, RSSO, or guest users.",
				Computed:    true,
			},
			"guest": {
				Type:        schema.TypeList,
				Description: "Guest User.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": {
							Type:        schema.TypeString,
							Description: "Comment.",
							Computed:    true,
						},
						"company": {
							Type:        schema.TypeString,
							Description: "Set the action for the company guest user field.",
							Computed:    true,
						},
						"email": {
							Type:        schema.TypeString,
							Description: "Email.",
							Computed:    true,
						},
						"expiration": {
							Type:        schema.TypeString,
							Description: "Expire time.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Guest ID.",
							Computed:    true,
						},
						"mobile_phone": {
							Type:        schema.TypeString,
							Description: "Mobile phone.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Guest name.",
							Computed:    true,
						},
						"password": {
							Type:        schema.TypeString,
							Description: "Guest password.",
							Computed:    true,
							Sensitive:   true,
						},
						"sponsor": {
							Type:        schema.TypeString,
							Description: "Set the action for the sponsor guest user field.",
							Computed:    true,
						},
						"user_id": {
							Type:        schema.TypeString,
							Description: "Guest ID.",
							Computed:    true,
						},
					},
				},
			},
			"http_digest_realm": {
				Type:        schema.TypeString,
				Description: "Realm attribute for MD5-digest authentication.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Group ID.",
				Computed:    true,
			},
			"match": {
				Type:        schema.TypeList,
				Description: "Group matches.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_name": {
							Type:        schema.TypeString,
							Description: "Name of matching user or group on remote authentication server.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"server_name": {
							Type:        schema.TypeString,
							Description: "Name of remote auth server.",
							Computed:    true,
						},
					},
				},
			},
			"max_accounts": {
				Type:        schema.TypeInt,
				Description: "Maximum number of guest accounts that can be created for this group (0 means unlimited).",
				Computed:    true,
			},
			"member": {
				Type:        schema.TypeList,
				Description: "Names of users, peers, LDAP severs, or RADIUS servers to add to the user group.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Group member name.",
							Computed:    true,
						},
					},
				},
			},
			"mobile_phone": {
				Type:        schema.TypeString,
				Description: "Enable/disable the guest user mobile phone number field.",
				Computed:    true,
			},
			"multiple_guest_add": {
				Type:        schema.TypeString,
				Description: "Enable/disable addition of multiple guests.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Group name.",
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Guest user password type.",
				Computed:    true,
			},
			"sms_custom_server": {
				Type:        schema.TypeString,
				Description: "SMS server.",
				Computed:    true,
			},
			"sms_server": {
				Type:        schema.TypeString,
				Description: "Send SMS through FortiGuard or other external server.",
				Computed:    true,
			},
			"sponsor": {
				Type:        schema.TypeString,
				Description: "Set the action for the sponsor guest user field.",
				Computed:    true,
			},
			"sso_attribute_value": {
				Type:        schema.TypeString,
				Description: "Name of the RADIUS user group that this local user group represents.",
				Computed:    true,
			},
			"user_id": {
				Type:        schema.TypeString,
				Description: "Guest user ID type.",
				Computed:    true,
			},
			"user_name": {
				Type:        schema.TypeString,
				Description: "Enable/disable the guest user name entry.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadUserGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserGroup dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
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

	d.SetId(mkey)

	return nil
}
