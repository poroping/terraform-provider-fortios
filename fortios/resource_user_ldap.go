// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure LDAP server entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserLdap() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserLdapCreate,
		Read:   resourceUserLdapRead,
		Update: resourceUserLdapUpdate,
		Delete: resourceUserLdapDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
			},
			"secondary_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"tertiary_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"server_identity_check": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cnid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 20),
				Optional:     true,
				Computed:     true,
			},
			"dn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Required:     true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_authentication": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_notification": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"password": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"group_member_check": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_search_base": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"group_object_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional:     true,
				Computed:     true,
			},
			"group_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional:     true,
				Computed:     true,
			},
			"secure": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"password_expiry_warning": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_renewal": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member_attr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"account_key_processing": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional:     true,
				Computed:     true,
			},
			"search_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obtain_user_info": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_info_exchange_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"interface_select_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"batchid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceUserLdapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	urlparams := make(map[string][]string)

	obj, err := getObjectUserLdap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating UserLdap resource while getting object: %v", err)
	}

	o, err := c.CreateUserLdap(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating UserLdap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserLdap")
	}

	return resourceUserLdapRead(d, m)
}

func resourceUserLdapUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	urlparams := make(map[string][]string)

	obj, err := getObjectUserLdap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating UserLdap resource while getting object: %v", err)
	}

	o, err := c.UpdateUserLdap(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating UserLdap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserLdap")
	}

	return resourceUserLdapRead(d, m)
}

func resourceUserLdapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	err := c.DeleteUserLdap(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting UserLdap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserLdapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	urlparams := make(map[string][]string)

	o, err := c.ReadUserLdap(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading UserLdap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserLdap(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading UserLdap resource from API: %v", err)
	}
	return nil
}

func flattenUserLdapName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSecondaryServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapTertiaryServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapCnid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapDn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapTwoFactor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapGroupMemberCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapGroupSearchBase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapGroupObjectFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapGroupFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapCaCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapPasswordExpiryWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapPasswordRenewal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapMemberAttr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapAccountKeyProcessing(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapAccountKeyFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSearchType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapObtainUserInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapUserInfoExchangeServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserLdap(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserLdapName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("server", flattenUserLdapServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("error reading server: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenUserLdapSecondaryServer(o["secondary-server"], d, "secondary_server", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-server"]) {
			return fmt.Errorf("error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenUserLdapTertiaryServer(o["tertiary-server"], d, "tertiary_server", sv)); err != nil {
		if !fortiAPIPatch(o["tertiary-server"]) {
			return fmt.Errorf("error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("server_identity_check", flattenUserLdapServerIdentityCheck(o["server-identity-check"], d, "server_identity_check", sv)); err != nil {
		if !fortiAPIPatch(o["server-identity-check"]) {
			return fmt.Errorf("error reading server_identity_check: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserLdapSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("error reading source_ip: %v", err)
		}
	}

	if err = d.Set("cnid", flattenUserLdapCnid(o["cnid"], d, "cnid", sv)); err != nil {
		if !fortiAPIPatch(o["cnid"]) {
			return fmt.Errorf("error reading cnid: %v", err)
		}
	}

	if err = d.Set("dn", flattenUserLdapDn(o["dn"], d, "dn", sv)); err != nil {
		if !fortiAPIPatch(o["dn"]) {
			return fmt.Errorf("error reading dn: %v", err)
		}
	}

	if err = d.Set("type", flattenUserLdapType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("error reading type: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenUserLdapTwoFactor(o["two-factor"], d, "two_factor", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor"]) {
			return fmt.Errorf("error reading two_factor: %v", err)
		}
	}

	if err = d.Set("two_factor_authentication", flattenUserLdapTwoFactorAuthentication(o["two-factor-authentication"], d, "two_factor_authentication", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor-authentication"]) {
			return fmt.Errorf("error reading two_factor_authentication: %v", err)
		}
	}

	if err = d.Set("two_factor_notification", flattenUserLdapTwoFactorNotification(o["two-factor-notification"], d, "two_factor_notification", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor-notification"]) {
			return fmt.Errorf("error reading two_factor_notification: %v", err)
		}
	}

	if err = d.Set("username", flattenUserLdapUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("error reading username: %v", err)
		}
	}

	if err = d.Set("group_member_check", flattenUserLdapGroupMemberCheck(o["group-member-check"], d, "group_member_check", sv)); err != nil {
		if !fortiAPIPatch(o["group-member-check"]) {
			return fmt.Errorf("error reading group_member_check: %v", err)
		}
	}

	if err = d.Set("group_search_base", flattenUserLdapGroupSearchBase(o["group-search-base"], d, "group_search_base", sv)); err != nil {
		if !fortiAPIPatch(o["group-search-base"]) {
			return fmt.Errorf("error reading group_search_base: %v", err)
		}
	}

	if err = d.Set("group_object_filter", flattenUserLdapGroupObjectFilter(o["group-object-filter"], d, "group_object_filter", sv)); err != nil {
		if !fortiAPIPatch(o["group-object-filter"]) {
			return fmt.Errorf("error reading group_object_filter: %v", err)
		}
	}

	if err = d.Set("group_filter", flattenUserLdapGroupFilter(o["group-filter"], d, "group_filter", sv)); err != nil {
		if !fortiAPIPatch(o["group-filter"]) {
			return fmt.Errorf("error reading group_filter: %v", err)
		}
	}

	if err = d.Set("secure", flattenUserLdapSecure(o["secure"], d, "secure", sv)); err != nil {
		if !fortiAPIPatch(o["secure"]) {
			return fmt.Errorf("error reading secure: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenUserLdapSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("ca_cert", flattenUserLdapCaCert(o["ca-cert"], d, "ca_cert", sv)); err != nil {
		if !fortiAPIPatch(o["ca-cert"]) {
			return fmt.Errorf("error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("port", flattenUserLdapPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("error reading port: %v", err)
		}
	}

	if err = d.Set("password_expiry_warning", flattenUserLdapPasswordExpiryWarning(o["password-expiry-warning"], d, "password_expiry_warning", sv)); err != nil {
		if !fortiAPIPatch(o["password-expiry-warning"]) {
			return fmt.Errorf("error reading password_expiry_warning: %v", err)
		}
	}

	if err = d.Set("password_renewal", flattenUserLdapPasswordRenewal(o["password-renewal"], d, "password_renewal", sv)); err != nil {
		if !fortiAPIPatch(o["password-renewal"]) {
			return fmt.Errorf("error reading password_renewal: %v", err)
		}
	}

	if err = d.Set("member_attr", flattenUserLdapMemberAttr(o["member-attr"], d, "member_attr", sv)); err != nil {
		if !fortiAPIPatch(o["member-attr"]) {
			return fmt.Errorf("error reading member_attr: %v", err)
		}
	}

	if err = d.Set("account_key_processing", flattenUserLdapAccountKeyProcessing(o["account-key-processing"], d, "account_key_processing", sv)); err != nil {
		if !fortiAPIPatch(o["account-key-processing"]) {
			return fmt.Errorf("error reading account_key_processing: %v", err)
		}
	}

	if err = d.Set("account_key_filter", flattenUserLdapAccountKeyFilter(o["account-key-filter"], d, "account_key_filter", sv)); err != nil {
		if !fortiAPIPatch(o["account-key-filter"]) {
			return fmt.Errorf("error reading account_key_filter: %v", err)
		}
	}

	if err = d.Set("search_type", flattenUserLdapSearchType(o["search-type"], d, "search_type", sv)); err != nil {
		if !fortiAPIPatch(o["search-type"]) {
			return fmt.Errorf("error reading search_type: %v", err)
		}
	}

	if err = d.Set("obtain_user_info", flattenUserLdapObtainUserInfo(o["obtain-user-info"], d, "obtain_user_info", sv)); err != nil {
		if !fortiAPIPatch(o["obtain-user-info"]) {
			return fmt.Errorf("error reading obtain_user_info: %v", err)
		}
	}

	if err = d.Set("user_info_exchange_server", flattenUserLdapUserInfoExchangeServer(o["user-info-exchange-server"], d, "user_info_exchange_server", sv)); err != nil {
		if !fortiAPIPatch(o["user-info-exchange-server"]) {
			return fmt.Errorf("error reading user_info_exchange_server: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenUserLdapInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenUserLdapInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	return nil
}

func flattenUserLdapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserLdapName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSecondaryServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTertiaryServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapServerIdentityCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapCnid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTwoFactor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTwoFactorAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTwoFactorNotification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupMemberCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupSearchBase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupObjectFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapCaCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPasswordExpiryWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPasswordRenewal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapMemberAttr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAccountKeyProcessing(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAccountKeyFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSearchType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapObtainUserInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapUserInfoExchangeServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserLdap(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandUserLdapName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {

		t, err := expandUserLdapServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok {

		t, err := expandUserLdapSecondaryServer(d, v, "secondary_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok {

		t, err := expandUserLdapTertiaryServer(d, v, "tertiary_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	if v, ok := d.GetOk("server_identity_check"); ok {

		t, err := expandUserLdapServerIdentityCheck(d, v, "server_identity_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandUserLdapSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("cnid"); ok {

		t, err := expandUserLdapCnid(d, v, "cnid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cnid"] = t
		}
	}

	if v, ok := d.GetOk("dn"); ok {

		t, err := expandUserLdapDn(d, v, "dn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dn"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandUserLdapType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok {

		t, err := expandUserLdapTwoFactor(d, v, "two_factor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_authentication"); ok {

		t, err := expandUserLdapTwoFactorAuthentication(d, v, "two_factor_authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-authentication"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_notification"); ok {

		t, err := expandUserLdapTwoFactorNotification(d, v, "two_factor_notification", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-notification"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {

		t, err := expandUserLdapUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {

		t, err := expandUserLdapPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("group_member_check"); ok {

		t, err := expandUserLdapGroupMemberCheck(d, v, "group_member_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-member-check"] = t
		}
	}

	if v, ok := d.GetOk("group_search_base"); ok {

		t, err := expandUserLdapGroupSearchBase(d, v, "group_search_base", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-search-base"] = t
		}
	}

	if v, ok := d.GetOk("group_object_filter"); ok {

		t, err := expandUserLdapGroupObjectFilter(d, v, "group_object_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-object-filter"] = t
		}
	}

	if v, ok := d.GetOk("group_filter"); ok {

		t, err := expandUserLdapGroupFilter(d, v, "group_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-filter"] = t
		}
	}

	if v, ok := d.GetOk("secure"); ok {

		t, err := expandUserLdapSecure(d, v, "secure", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {

		t, err := expandUserLdapSslMinProtoVersion(d, v, "ssl_min_proto_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok {

		t, err := expandUserLdapCaCert(d, v, "ca_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {

		t, err := expandUserLdapPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("password_expiry_warning"); ok {

		t, err := expandUserLdapPasswordExpiryWarning(d, v, "password_expiry_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-expiry-warning"] = t
		}
	}

	if v, ok := d.GetOk("password_renewal"); ok {

		t, err := expandUserLdapPasswordRenewal(d, v, "password_renewal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("member_attr"); ok {

		t, err := expandUserLdapMemberAttr(d, v, "member_attr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-attr"] = t
		}
	}

	if v, ok := d.GetOk("account_key_processing"); ok {

		t, err := expandUserLdapAccountKeyProcessing(d, v, "account_key_processing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-processing"] = t
		}
	}

	if v, ok := d.GetOk("account_key_filter"); ok {

		t, err := expandUserLdapAccountKeyFilter(d, v, "account_key_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-filter"] = t
		}
	}

	if v, ok := d.GetOk("search_type"); ok {

		t, err := expandUserLdapSearchType(d, v, "search_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["search-type"] = t
		}
	}

	if v, ok := d.GetOk("obtain_user_info"); ok {

		t, err := expandUserLdapObtainUserInfo(d, v, "obtain_user_info", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obtain-user-info"] = t
		}
	}

	if v, ok := d.GetOk("user_info_exchange_server"); ok {

		t, err := expandUserLdapUserInfoExchangeServer(d, v, "user_info_exchange_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-info-exchange-server"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {

		t, err := expandUserLdapInterfaceSelectMethod(d, v, "interface_select_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandUserLdapInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	return &obj, nil
}
