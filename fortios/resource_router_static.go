// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv4 static routing tables.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterStatic() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterStaticCreate,
		Read:   resourceRouterStaticRead,
		Update: resourceRouterStaticUpdate,
		Delete: resourceRouterStaticDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"seq_num": {
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Optional:     true,
				Computed:     true,
			},
			"dst": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateIPv4Classnet,
				Optional:     true,
				Computed:     true,
			},
			"src": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateIPv4Classnet,
				Optional:     true,
				Computed:     true,
			},
			"gateway": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,
				Optional:     true,
				Computed:     true,
			},
			"distance": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"blackhole": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_gateway": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Optional:     true,
				Computed:     true,
			},
			"sdwan": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Optional:     true,
				Computed:     true,
			},
			"virtual_wan_link": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Optional:     true,
				Computed:     true,
			},
			"dstaddr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"internet_service": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"internet_service_custom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"link_monitor_exempt": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Optional:     true,
				Computed:     true,
			},
			"vrf": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"bfd": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
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

func resourceRouterStaticCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterStatic(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating RouterStatic resource while getting object: %v", err)
	}

	o, err := c.CreateRouterStatic(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating RouterStatic resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterStatic")
	}

	return resourceRouterStaticRead(d, m)
}

func resourceRouterStaticUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterStatic(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating RouterStatic resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterStatic(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating RouterStatic resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterStatic")
	}

	return resourceRouterStaticRead(d, m)
}

func resourceRouterStaticDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteRouterStatic(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting RouterStatic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterStaticRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterStatic(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading RouterStatic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterStatic(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading RouterStatic resource from API: %v", err)
	}
	return nil
}

func flattenRouterStaticSeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterStaticSrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterStaticGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticBlackhole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticDynamicGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticSdwan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticVirtualWanLink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticInternetService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticLinkMonitorExempt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterStatic(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("seq_num", flattenRouterStaticSeqNum(o["seq-num"], d, "seq_num", sv)); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("error reading seq_num: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterStaticStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("dst", flattenRouterStaticDst(o["dst"], d, "dst", sv)); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("error reading dst: %v", err)
		}
	}

	if err = d.Set("src", flattenRouterStaticSrc(o["src"], d, "src", sv)); err != nil {
		if !fortiAPIPatch(o["src"]) {
			return fmt.Errorf("error reading src: %v", err)
		}
	}

	if err = d.Set("gateway", flattenRouterStaticGateway(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("error reading gateway: %v", err)
		}
	}

	if err = d.Set("distance", flattenRouterStaticDistance(o["distance"], d, "distance", sv)); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("error reading distance: %v", err)
		}
	}

	if err = d.Set("weight", flattenRouterStaticWeight(o["weight"], d, "weight", sv)); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("error reading weight: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterStaticPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("error reading priority: %v", err)
		}
	}

	if err = d.Set("device", flattenRouterStaticDevice(o["device"], d, "device", sv)); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("error reading device: %v", err)
		}
	}

	if err = d.Set("comment", flattenRouterStaticComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("error reading comment: %v", err)
		}
	}

	if err = d.Set("blackhole", flattenRouterStaticBlackhole(o["blackhole"], d, "blackhole", sv)); err != nil {
		if !fortiAPIPatch(o["blackhole"]) {
			return fmt.Errorf("error reading blackhole: %v", err)
		}
	}

	if err = d.Set("dynamic_gateway", flattenRouterStaticDynamicGateway(o["dynamic-gateway"], d, "dynamic_gateway", sv)); err != nil {
		if !fortiAPIPatch(o["dynamic-gateway"]) {
			return fmt.Errorf("error reading dynamic_gateway: %v", err)
		}
	}

	if err = d.Set("sdwan", flattenRouterStaticSdwan(o["sdwan"], d, "sdwan", sv)); err != nil {
		if !fortiAPIPatch(o["sdwan"]) {
			return fmt.Errorf("error reading sdwan: %v", err)
		}
	}

	if err = d.Set("virtual_wan_link", flattenRouterStaticVirtualWanLink(o["virtual-wan-link"], d, "virtual_wan_link", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-wan-link"]) {
			return fmt.Errorf("error reading virtual_wan_link: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenRouterStaticDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
		if !fortiAPIPatch(o["dstaddr"]) {
			return fmt.Errorf("error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenRouterStaticInternetService(o["internet-service"], d, "internet_service", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service"]) {
			return fmt.Errorf("error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", flattenRouterStaticInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-custom"]) {
			return fmt.Errorf("error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("link_monitor_exempt", flattenRouterStaticLinkMonitorExempt(o["link-monitor-exempt"], d, "link_monitor_exempt", sv)); err != nil {
		if !fortiAPIPatch(o["link-monitor-exempt"]) {
			return fmt.Errorf("error reading link_monitor_exempt: %v", err)
		}
	}

	if err = d.Set("vrf", flattenRouterStaticVrf(o["vrf"], d, "vrf", sv)); err != nil {
		if !fortiAPIPatch(o["vrf"]) {
			return fmt.Errorf("error reading vrf: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterStaticBfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("error reading bfd: %v", err)
		}
	}

	return nil
}

func flattenRouterStaticFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterStaticSeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticBlackhole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticDynamicGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticSdwan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticVirtualWanLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticInternetService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticLinkMonitorExempt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterStatic(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("seq_num"); ok {

		t, err := expandRouterStaticSeqNum(d, v, "seq_num", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandRouterStaticStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok {

		t, err := expandRouterStaticDst(d, v, "dst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok {

		t, err := expandRouterStaticSrc(d, v, "src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {

		t, err := expandRouterStaticGateway(d, v, "gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok {

		t, err := expandRouterStaticDistance(d, v, "distance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOkExists("weight"); ok {

		t, err := expandRouterStaticWeight(d, v, "weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOkExists("priority"); ok {

		t, err := expandRouterStaticPriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok {

		t, err := expandRouterStaticDevice(d, v, "device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandRouterStaticComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("blackhole"); ok {

		t, err := expandRouterStaticBlackhole(d, v, "blackhole", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blackhole"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_gateway"); ok {

		t, err := expandRouterStaticDynamicGateway(d, v, "dynamic_gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-gateway"] = t
		}
	}

	if v, ok := d.GetOk("sdwan"); ok {

		t, err := expandRouterStaticSdwan(d, v, "sdwan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan"] = t
		}
	}

	if v, ok := d.GetOk("virtual_wan_link"); ok {

		t, err := expandRouterStaticVirtualWanLink(d, v, "virtual_wan_link", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-wan-link"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {

		t, err := expandRouterStaticDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOkExists("internet_service"); ok {

		t, err := expandRouterStaticInternetService(d, v, "internet_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok {

		t, err := expandRouterStaticInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("link_monitor_exempt"); ok {

		t, err := expandRouterStaticLinkMonitorExempt(d, v, "link_monitor_exempt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-monitor-exempt"] = t
		}
	}

	if v, ok := d.GetOkExists("vrf"); ok {

		t, err := expandRouterStaticVrf(d, v, "vrf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {

		t, err := expandRouterStaticBfd(d, v, "bfd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	return &obj, nil
}
