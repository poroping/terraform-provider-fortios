package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirewallSecurityPolicySeq() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSecurityPolicySeqCreateUpdate,
		Read:   resourceFirewallSecurityPolicySeqRead,
		Update: resourceFirewallSecurityPolicySeqCreateUpdate,
		Delete: resourceFirewallSecurityPolicySeqDel,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"policy_src_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"policy_dst_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"alter_position": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enable_state_checking": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"state_policy_srcdst_pos": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"state_policy_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policyid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
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

func resourceFirewallSecurityPolicySeqCreateUpdate(d *schema.ResourceData, m interface{}) error {
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

	srcIdPatch := d.Get("policy_src_id").(int)
	dstIdPatch := d.Get("policy_dst_id").(int)
	alterPos := d.Get("alter_position").(string)

	srcId := strconv.Itoa(srcIdPatch)
	dstId := strconv.Itoa(dstIdPatch)

	if alterPos != "before" && alterPos != "after" {
		return fmt.Errorf("<alter_position> param should be only 'after' or 'before'")
	}

	err := c.CreateUpdateFirewallSecurityPolicySeq(srcId, dstId, alterPos, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error Altering Firewall Security Policy Sequence: %s", err)
	}

	d.SetId(srcId)

	return resourceFirewallSecurityPolicySeqRead(d, m)
}

// Not suitable operation
func resourceFirewallSecurityPolicySeqDel(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	return c.DelFirewallSecurityPolicySeq()
}

func resourceFirewallSecurityPolicySeqRead(d *schema.ResourceData, m interface{}) error {
	enable_state_checking := d.Get("enable_state_checking").(bool)

	if enable_state_checking == false {
		d.Set("state_policy_list", nil)
		return nil
	}

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

	sid := d.Get("policy_src_id").(int)
	did := d.Get("policy_dst_id").(int)
	action := d.Get("alter_position").(string)

	o, err := c.GetSecurityPolicyList(vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error reading Firewall Security Policy List: %s", err)
	}

	if o != nil {
		i := 1
		now_sid := -1
		now_did := -1

		var items []interface{}
		for _, z := range o {
			m := make(map[string]interface{})

			idn := z.PolicyID
			if idn == strconv.Itoa(d.Get("policy_src_id").(int)) {
				now_sid = i
				idn = "*" + idn
			}

			if idn == strconv.Itoa(d.Get("policy_dst_id").(int)) {
				now_did = i
				idn = "*" + idn
			}

			m["policyid"] = idn
			m["name"] = z.Name
			m["action"] = z.Action

			items = append(items, m)
			i += 1
		}

		if err := d.Set("state_policy_list", items); err != nil {
			log.Printf("[WARN] Error reading Firewall Security Policy List for (%s): %s", d.Id(), err)
		}

		state_policy_srcdst_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_policy_srcdst_pos = "policies with policy_src_id(" + strconv.Itoa(sid) + ")  and policy_dst_id(" + strconv.Itoa(did) + ") were deleted"
			} else if now_sid == -1 {
				state_policy_srcdst_pos = "policy with policy_src_id(" + strconv.Itoa(sid) + ") was deleted"
			} else if now_did == -1 {
				state_policy_srcdst_pos = "policy with policy_dst_id(" + strconv.Itoa(did) + ") was deleted"
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
					state_policy_srcdst_pos = "policy with policy_src_id(" + strconv.Itoa(sid) + ") is " + strconv.Itoa(relative_pos) + " behind policy with policy_dst_id(" + strconv.Itoa(did) + ")"
				} else {
					state_policy_srcdst_pos = "policy with policy_src_id(" + strconv.Itoa(sid) + ") is " + strconv.Itoa(-relative_pos) + " ahead of policy with policy_dst_id(" + strconv.Itoa(did) + ")"
				}
			}
		}

		d.Set("state_policy_srcdst_pos", state_policy_srcdst_pos)

	}

	return nil
}
