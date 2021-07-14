package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirewallCentralsnatmapMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallCentralsnatmapMoveCreateUpdate,
		Read:   resourceFirewallCentralsnatmapMoveRead,
		Update: resourceFirewallCentralsnatmapMoveCreateUpdate,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"policyid_src": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"policyid_dst": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"move": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state_policy_srcdst_pos": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"batchid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceFirewallCentralsnatmapMoveCreateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

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

	srcIdPatch := d.Get("policyid_src").(int)
	dstIdPatch := d.Get("policyid_dst").(int)
	mv := d.Get("move").(string)

	srcId := strconv.Itoa(srcIdPatch)
	dstId := strconv.Itoa(dstIdPatch)

	if mv != "before" && mv != "after" {
		return fmt.Errorf("<move> param should be only 'after' or 'before'")
	}

	err := c.CreateUpdateFirewallCentralsnatmapMove(srcId, dstId, mv, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error Altering FirewallCentralsnatmap Moveuence: %s", err)
	}

	d.SetId(srcId)

	return resourceFirewallCentralsnatmapMoveRead(d, m)
}

func resourceFirewallCentralsnatmapMoveRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

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

	sid := d.Get("policyid_src").(int)
	did := d.Get("policyid_dst").(int)
	action := d.Get("move").(string)

	o, err := c.GetFirewallCentralsnatmapList(vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error reading FirewallCentralsnatmap List: %s", err)
	}

	if o != nil {
		i := 1
		now_sid := -1
		now_did := -1

		for _, z := range o {
			idn := z.Policyid
			if idn == strconv.Itoa(d.Get("policyid_src").(int)) {
				now_sid = i
			}

			if idn == strconv.Itoa(d.Get("policyid_dst").(int)) {
				now_did = i
			}

			i += 1
		}

		state_policy_srcdst_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_policy_srcdst_pos = "FirewallCentralsnatmap with policyid_src(" + strconv.Itoa(sid) + ")  and policyid_dst(" + strconv.Itoa(did) + ") were deleted"
			} else if now_sid == -1 {
				state_policy_srcdst_pos = "FirewallCentralsnatmap with policyid_src(" + strconv.Itoa(sid) + ") was deleted"
			} else if now_did == -1 {
				state_policy_srcdst_pos = "FirewallCentralsnatmap with policyid_dst(" + strconv.Itoa(did) + ") was deleted"
			}
		} else {
			bconsistent := true
			if action == "before" {
				if now_sid != now_did-1 {
					bconsistent = false
				}
			}

			if action == "after" {
				if now_sid != now_did+1 {
					bconsistent = false
				}
			}

			if bconsistent == false {
				relative_pos := now_sid - now_did

				if relative_pos > 0 {
					state_policy_srcdst_pos = "FirewallCentralsnatmap with policyid_src(" + strconv.Itoa(sid) + ") is " + strconv.Itoa(relative_pos) + " behind FirewallCentralsnatmap with policyid_dst(" + strconv.Itoa(did) + ")"
				} else {
					state_policy_srcdst_pos = "FirewallCentralsnatmap with policyid_src(" + strconv.Itoa(sid) + ") is " + strconv.Itoa(-relative_pos) + " ahead of FirewallCentralsnatmap with policyid_dst(" + strconv.Itoa(did) + ")"
				}
			}
		}

		d.Set("state_policy_srcdst_pos", state_policy_srcdst_pos)

	}

	return nil
}
