// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.0 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Configure Access Proxy virtual hosts.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceFirewallAccessProxyVirtualHost() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallAccessProxyVirtualHostRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"host": {
				Type:        schema.TypeString,
				Description: "The host name.",
				Computed:    true,
			},
			"host_type": {
				Type:        schema.TypeString,
				Description: "Type of host pattern.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Virtual host name.",
				Required:    true,
			},
			"ssl_certificate": {
				Type:        schema.TypeString,
				Description: "SSL certificate for this host.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallAccessProxyVirtualHostRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""
	key := "name"
	t := d.Get(key)
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing FirewallAccessProxyVirtualHost: type error")
	}

	o, err := c.ReadFirewallAccessProxyVirtualHost(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing FirewallAccessProxyVirtualHost: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallAccessProxyVirtualHost(d, o)
	if err != nil {
		return fmt.Errorf("error describing FirewallAccessProxyVirtualHost from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallAccessProxyVirtualHostHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyVirtualHostHostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyVirtualHostName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyVirtualHostSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallAccessProxyVirtualHost(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("host", dataSourceFlattenFirewallAccessProxyVirtualHostHost(o["host"], d, "host")); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("error reading host: %v", err)
		}
	}
	if err = d.Set("host_type", dataSourceFlattenFirewallAccessProxyVirtualHostHostType(o["host-type"], d, "host_type")); err != nil {
		if !fortiAPIPatch(o["host-type"]) {
			return fmt.Errorf("error reading host_type: %v", err)
		}
	}
	if err = d.Set("name", dataSourceFlattenFirewallAccessProxyVirtualHostName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}
	if err = d.Set("ssl_certificate", dataSourceFlattenFirewallAccessProxyVirtualHostSslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if !fortiAPIPatch(o["ssl-certificate"]) {
			return fmt.Errorf("error reading ssl_certificate: %v", err)
		}
	}

	return nil
}
