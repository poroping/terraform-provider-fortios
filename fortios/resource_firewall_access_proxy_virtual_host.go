// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.0
// Template Authors:
// Justin Roberts (@poroping)

// Description: Configure Access Proxy virtual hosts.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallAccessProxyVirtualHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxyVirtualHostCreate,
		Read:   resourceFirewallAccessProxyVirtualHostRead,
		Update: resourceFirewallAccessProxyVirtualHostUpdate,
		Delete: resourceFirewallAccessProxyVirtualHostDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"batchid": {
				Type:        schema.TypeInt,
				Description: "Associate with batch. From 6.4.x+. Currently a WIP and broken.",
				Optional:    true,
				Default:     0,
			},
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"host": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Description:  "The host name.",
				Optional:     true,
				Computed:     true,
			},
			"host_type": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"sub-string", "wildcard"}),
				Description:  "Type of host pattern.",
				Optional:     true,
				Computed:     true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Description:  "Virtual host name.",
				ForceNew:     true,
				Required:     true,
			},
			"ssl_certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "SSL certificate for this host.",
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceFirewallAccessProxyVirtualHostCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

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

	allow_append := false

	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}

	urlparams["allow_append"] = []string{strconv.FormatBool(allow_append)}

	key := "name"
	mkey := ""
	if v, ok := d.GetOk(key); ok {
		if s, ok := v.(string); ok {
			mkey = s
		}
	}

	obj, err := getObjectFirewallAccessProxyVirtualHost(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating FirewallAccessProxyVirtualHost resource while getting object: %v", err)
	}

	if mkey == "" && allow_append {
		return fmt.Errorf("error creating FirewallAccessProxyVirtualHost resource: %q must be set if \"allow_append\" is true", key)
	}

	o := make(map[string]interface{})
	if mkey != "" && allow_append {
		o, err = c.UpdateFirewallAccessProxyVirtualHost(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.CreateFirewallAccessProxyVirtualHost(obj, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error creating FirewallAccessProxyVirtualHost resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAccessProxyVirtualHost")
	}

	return resourceFirewallAccessProxyVirtualHostRead(d, m)
}

func resourceFirewallAccessProxyVirtualHostUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

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

	obj, err := getObjectFirewallAccessProxyVirtualHost(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating FirewallAccessProxyVirtualHost resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallAccessProxyVirtualHost(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating FirewallAccessProxyVirtualHost resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAccessProxyVirtualHost")
	}

	return resourceFirewallAccessProxyVirtualHostRead(d, m)
}

func resourceFirewallAccessProxyVirtualHostDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteFirewallAccessProxyVirtualHost(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting FirewallAccessProxyVirtualHost resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxyVirtualHostRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

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

	o, err := c.ReadFirewallAccessProxyVirtualHost(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading FirewallAccessProxyVirtualHost resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxyVirtualHost(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading FirewallAccessProxyVirtualHost resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxyVirtualHostHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyVirtualHostHostType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyVirtualHostName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyVirtualHostSslCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallAccessProxyVirtualHost(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("host", flattenFirewallAccessProxyVirtualHostHost(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("error reading host: %v", err)
		}
	}
	if err = d.Set("host_type", flattenFirewallAccessProxyVirtualHostHostType(o["host-type"], d, "host_type", sv)); err != nil {
		if !fortiAPIPatch(o["host-type"]) {
			return fmt.Errorf("error reading host_type: %v", err)
		}
	}
	if err = d.Set("name", flattenFirewallAccessProxyVirtualHostName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}
	if err = d.Set("ssl_certificate", flattenFirewallAccessProxyVirtualHostSslCertificate(o["ssl-certificate"], d, "ssl_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-certificate"]) {
			return fmt.Errorf("error reading ssl_certificate: %v", err)
		}
	}

	return nil
}

func expandFirewallAccessProxyVirtualHostHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyVirtualHostHostType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyVirtualHostName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyVirtualHostSslCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAccessProxyVirtualHost(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("host"); ok {

		t, err := expandFirewallAccessProxyVirtualHostHost(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}
	if v, ok := d.GetOk("host_type"); ok {

		t, err := expandFirewallAccessProxyVirtualHostHostType(d, v, "host_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-type"] = t
		}
	}
	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallAccessProxyVirtualHostName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}
	if v, ok := d.GetOk("ssl_certificate"); ok {

		t, err := expandFirewallAccessProxyVirtualHostSslCertificate(d, v, "ssl_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	return &obj, nil
}
