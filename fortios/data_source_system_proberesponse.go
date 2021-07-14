// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure system probe response.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemProbeResponse() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemProbeResponseRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"http_probe_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ttl_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemProbeResponseRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemProbeResponse"

	o, err := c.ReadSystemProbeResponse(mkey, vdomparam, make(map[string][]string), 0)
	if err != nil {
		return fmt.Errorf("error describing SystemProbeResponse: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemProbeResponse(d, o)
	if err != nil {
		return fmt.Errorf("error describing SystemProbeResponse from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemProbeResponsePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemProbeResponseHttpProbeValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemProbeResponseTtlMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemProbeResponseMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemProbeResponseSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemProbeResponsePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemProbeResponseTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemProbeResponse(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("port", dataSourceFlattenSystemProbeResponsePort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("error reading port: %v", err)
		}
	}

	if err = d.Set("http_probe_value", dataSourceFlattenSystemProbeResponseHttpProbeValue(o["http-probe-value"], d, "http_probe_value")); err != nil {
		if !fortiAPIPatch(o["http-probe-value"]) {
			return fmt.Errorf("error reading http_probe_value: %v", err)
		}
	}

	if err = d.Set("ttl_mode", dataSourceFlattenSystemProbeResponseTtlMode(o["ttl-mode"], d, "ttl_mode")); err != nil {
		if !fortiAPIPatch(o["ttl-mode"]) {
			return fmt.Errorf("error reading ttl_mode: %v", err)
		}
	}

	if err = d.Set("mode", dataSourceFlattenSystemProbeResponseMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("error reading mode: %v", err)
		}
	}

	if err = d.Set("security_mode", dataSourceFlattenSystemProbeResponseSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if !fortiAPIPatch(o["security-mode"]) {
			return fmt.Errorf("error reading security_mode: %v", err)
		}
	}

	if err = d.Set("timeout", dataSourceFlattenSystemProbeResponseTimeout(o["timeout"], d, "timeout")); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("error reading timeout: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemProbeResponseFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
