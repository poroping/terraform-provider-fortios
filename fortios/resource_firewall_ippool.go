// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv4 IP pools.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallIppool() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIppoolCreate,
		Read:   resourceFirewallIppoolRead,
		Update: resourceFirewallIppoolUpdate,
		Delete: resourceFirewallIppoolDelete,

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
				Optional:     true,
				Computed:     true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"startip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"endip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"startport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5117, 65533),
				Optional:     true,
				Computed:     true,
			},
			"endport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5117, 65533),
				Optional:     true,
				Computed:     true,
			},
			"source_startip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_endip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(64, 4096),
				Optional:     true,
				Computed:     true,
			},
			"port_per_user": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_blocks_per_user": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 128),
				Optional:     true,
				Computed:     true,
			},
			"pba_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 300),
				Optional:     true,
				Computed:     true,
			},
			"permit_any_host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp_reply": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp_intf": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"associated_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"batchid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"allow_append": {
				Type:         schema.TypeBool,
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"name"},
			},
		},
	}
}

func resourceFirewallIppoolCreate(d *schema.ResourceData, m interface{}) error {
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

	allow_append := false

	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}

	urlparams := make(map[string][]string)
	urlparams["allow_append"] = []string{strconv.FormatBool(allow_append)}

	key := "name"
	mkey := ""
	if v, ok := d.GetOk(key); ok {
		if s, ok := v.(string); ok {
			mkey = s
		}
	}

	obj, err := getObjectFirewallIppool(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating FirewallIppool resource while getting object: %v", err)
	}

	if mkey == "" && allow_append {
		return fmt.Errorf("error creating FirewallIppool resource: %q must be set if \"allow_append\" is true", key)
	}

	o := make(map[string]interface{})
	if mkey != "" && allow_append {
		o, err = c.UpdateFirewallIppool(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.CreateFirewallIppool(obj, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error creating FirewallIppool resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallIppool")
	}

	return resourceFirewallIppoolRead(d, m)
}

func resourceFirewallIppoolUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallIppool(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating FirewallIppool resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallIppool(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating FirewallIppool resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallIppool")
	}

	return resourceFirewallIppoolRead(d, m)
}

func resourceFirewallIppoolDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteFirewallIppool(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting FirewallIppool resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIppoolRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallIppool(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading FirewallIppool resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIppool(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading FirewallIppool resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIppoolName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolStartip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolEndip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolStartport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolEndport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolSourceStartip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolSourceEndip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolBlockSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolPortPerUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolNumBlocksPerUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolPbaTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolPermitAnyHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolArpReply(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolArpIntf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolAssociatedInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppoolComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallIppool(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallIppoolName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallIppoolType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("error reading type: %v", err)
		}
	}

	if err = d.Set("startip", flattenFirewallIppoolStartip(o["startip"], d, "startip", sv)); err != nil {
		if !fortiAPIPatch(o["startip"]) {
			return fmt.Errorf("error reading startip: %v", err)
		}
	}

	if err = d.Set("endip", flattenFirewallIppoolEndip(o["endip"], d, "endip", sv)); err != nil {
		if !fortiAPIPatch(o["endip"]) {
			return fmt.Errorf("error reading endip: %v", err)
		}
	}

	if err = d.Set("startport", flattenFirewallIppoolStartport(o["startport"], d, "startport", sv)); err != nil {
		if !fortiAPIPatch(o["startport"]) {
			return fmt.Errorf("error reading startport: %v", err)
		}
	}

	if err = d.Set("endport", flattenFirewallIppoolEndport(o["endport"], d, "endport", sv)); err != nil {
		if !fortiAPIPatch(o["endport"]) {
			return fmt.Errorf("error reading endport: %v", err)
		}
	}

	if err = d.Set("source_startip", flattenFirewallIppoolSourceStartip(o["source-startip"], d, "source_startip", sv)); err != nil {
		if !fortiAPIPatch(o["source-startip"]) {
			return fmt.Errorf("error reading source_startip: %v", err)
		}
	}

	if err = d.Set("source_endip", flattenFirewallIppoolSourceEndip(o["source-endip"], d, "source_endip", sv)); err != nil {
		if !fortiAPIPatch(o["source-endip"]) {
			return fmt.Errorf("error reading source_endip: %v", err)
		}
	}

	if err = d.Set("block_size", flattenFirewallIppoolBlockSize(o["block-size"], d, "block_size", sv)); err != nil {
		if !fortiAPIPatch(o["block-size"]) {
			return fmt.Errorf("error reading block_size: %v", err)
		}
	}

	if err = d.Set("port_per_user", flattenFirewallIppoolPortPerUser(o["port-per-user"], d, "port_per_user", sv)); err != nil {
		if !fortiAPIPatch(o["port-per-user"]) {
			return fmt.Errorf("error reading port_per_user: %v", err)
		}
	}

	if err = d.Set("num_blocks_per_user", flattenFirewallIppoolNumBlocksPerUser(o["num-blocks-per-user"], d, "num_blocks_per_user", sv)); err != nil {
		if !fortiAPIPatch(o["num-blocks-per-user"]) {
			return fmt.Errorf("error reading num_blocks_per_user: %v", err)
		}
	}

	if err = d.Set("pba_timeout", flattenFirewallIppoolPbaTimeout(o["pba-timeout"], d, "pba_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["pba-timeout"]) {
			return fmt.Errorf("error reading pba_timeout: %v", err)
		}
	}

	if err = d.Set("permit_any_host", flattenFirewallIppoolPermitAnyHost(o["permit-any-host"], d, "permit_any_host", sv)); err != nil {
		if !fortiAPIPatch(o["permit-any-host"]) {
			return fmt.Errorf("error reading permit_any_host: %v", err)
		}
	}

	if err = d.Set("arp_reply", flattenFirewallIppoolArpReply(o["arp-reply"], d, "arp_reply", sv)); err != nil {
		if !fortiAPIPatch(o["arp-reply"]) {
			return fmt.Errorf("error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("arp_intf", flattenFirewallIppoolArpIntf(o["arp-intf"], d, "arp_intf", sv)); err != nil {
		if !fortiAPIPatch(o["arp-intf"]) {
			return fmt.Errorf("error reading arp_intf: %v", err)
		}
	}

	if err = d.Set("associated_interface", flattenFirewallIppoolAssociatedInterface(o["associated-interface"], d, "associated_interface", sv)); err != nil {
		if !fortiAPIPatch(o["associated-interface"]) {
			return fmt.Errorf("error reading associated_interface: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallIppoolComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("error reading comments: %v", err)
		}
	}

	return nil
}

func flattenFirewallIppoolFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallIppoolName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolStartip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolEndip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolStartport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolEndport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolSourceStartip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolSourceEndip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolBlockSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolPortPerUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolNumBlocksPerUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolPbaTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolPermitAnyHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolArpReply(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolArpIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolAssociatedInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIppool(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallIppoolName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandFirewallIppoolType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("startip"); ok {

		t, err := expandFirewallIppoolStartip(d, v, "startip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startip"] = t
		}
	}

	if v, ok := d.GetOk("endip"); ok {

		t, err := expandFirewallIppoolEndip(d, v, "endip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endip"] = t
		}
	}

	if v, ok := d.GetOk("startport"); ok {

		t, err := expandFirewallIppoolStartport(d, v, "startport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startport"] = t
		}
	}

	if v, ok := d.GetOk("endport"); ok {

		t, err := expandFirewallIppoolEndport(d, v, "endport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endport"] = t
		}
	}

	if v, ok := d.GetOk("source_startip"); ok {

		t, err := expandFirewallIppoolSourceStartip(d, v, "source_startip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-startip"] = t
		}
	}

	if v, ok := d.GetOk("source_endip"); ok {

		t, err := expandFirewallIppoolSourceEndip(d, v, "source_endip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-endip"] = t
		}
	}

	if v, ok := d.GetOk("block_size"); ok {

		t, err := expandFirewallIppoolBlockSize(d, v, "block_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-size"] = t
		}
	}

	if v, ok := d.GetOkExists("port_per_user"); ok {

		t, err := expandFirewallIppoolPortPerUser(d, v, "port_per_user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-per-user"] = t
		}
	}

	if v, ok := d.GetOk("num_blocks_per_user"); ok {

		t, err := expandFirewallIppoolNumBlocksPerUser(d, v, "num_blocks_per_user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["num-blocks-per-user"] = t
		}
	}

	if v, ok := d.GetOk("pba_timeout"); ok {

		t, err := expandFirewallIppoolPbaTimeout(d, v, "pba_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pba-timeout"] = t
		}
	}

	if v, ok := d.GetOk("permit_any_host"); ok {

		t, err := expandFirewallIppoolPermitAnyHost(d, v, "permit_any_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-any-host"] = t
		}
	}

	if v, ok := d.GetOk("arp_reply"); ok {

		t, err := expandFirewallIppoolArpReply(d, v, "arp_reply", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("arp_intf"); ok {

		t, err := expandFirewallIppoolArpIntf(d, v, "arp_intf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-intf"] = t
		}
	}

	if v, ok := d.GetOk("associated_interface"); ok {

		t, err := expandFirewallIppoolAssociatedInterface(d, v, "associated_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["associated-interface"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandFirewallIppoolComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	return &obj, nil
}
