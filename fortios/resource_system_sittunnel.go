// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 tunnel over IPv4.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSitTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSitTunnelCreate,
		Read:   resourceSystemSitTunnelRead,
		Update: resourceSystemSitTunnelUpdate,
		Delete: resourceSystemSitTunnelDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"source": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip6": {
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

func resourceSystemSitTunnelCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSitTunnel(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating SystemSitTunnel resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSitTunnel(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating SystemSitTunnel resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSitTunnel")
	}

	return resourceSystemSitTunnelRead(d, m)
}

func resourceSystemSitTunnelUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSitTunnel(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating SystemSitTunnel resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSitTunnel(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating SystemSitTunnel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSitTunnel")
	}

	return resourceSystemSitTunnelRead(d, m)
}

func resourceSystemSitTunnelDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteSystemSitTunnel(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting SystemSitTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSitTunnelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSitTunnel(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading SystemSitTunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSitTunnel(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading SystemSitTunnel resource from API: %v", err)
	}
	return nil
}

func flattenSystemSitTunnelName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSitTunnelSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSitTunnelDestination(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSitTunnelIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSitTunnelInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSitTunnel(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemSitTunnelName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("source", flattenSystemSitTunnelSource(o["source"], d, "source", sv)); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("error reading source: %v", err)
		}
	}

	if err = d.Set("destination", flattenSystemSitTunnelDestination(o["destination"], d, "destination", sv)); err != nil {
		if !fortiAPIPatch(o["destination"]) {
			return fmt.Errorf("error reading destination: %v", err)
		}
	}

	if err = d.Set("ip6", flattenSystemSitTunnelIp6(o["ip6"], d, "ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("error reading ip6: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemSitTunnelInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	return nil
}

func flattenSystemSitTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSitTunnelName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSitTunnelSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSitTunnelDestination(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSitTunnelIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSitTunnelInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSitTunnel(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemSitTunnelName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {

		t, err := expandSystemSitTunnelSource(d, v, "source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("destination"); ok {

		t, err := expandSystemSitTunnelDestination(d, v, "destination", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {

		t, err := expandSystemSitTunnelIp6(d, v, "ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandSystemSitTunnelInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	return &obj, nil
}
