// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure the MAC address group.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceWirelessControllerAddrgrp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure the MAC address group.",

		CreateContext: resourceWirelessControllerAddrgrpCreate,
		ReadContext:   resourceWirelessControllerAddrgrpRead,
		UpdateContext: resourceWirelessControllerAddrgrpUpdate,
		DeleteContext: resourceWirelessControllerAddrgrpDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_append": {
				Type:         schema.TypeBool,
				Description:  "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"addresses": {
				Type:        schema.TypeList,
				Description: "Manually selected group of addresses.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Address ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"default_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),

				Description: "Allow or block the clients with MAC addresses that are not in the group.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerAddrgrpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	mkey := ""
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating WirelessControllerAddrgrp resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerAddrgrp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerAddrgrp(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerAddrgrp")
	}

	return resourceWirelessControllerAddrgrpRead(ctx, d, meta)
}

func resourceWirelessControllerAddrgrpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
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

	obj, diags := getObjectWirelessControllerAddrgrp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerAddrgrp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerAddrgrp resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerAddrgrp")
	}

	return resourceWirelessControllerAddrgrpRead(ctx, d, meta)
}

func resourceWirelessControllerAddrgrpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	err := c.Cmdb.DeleteWirelessControllerAddrgrp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerAddrgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerAddrgrpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	o, err := c.Cmdb.ReadWirelessControllerAddrgrp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerAddrgrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectWirelessControllerAddrgrp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerAddrgrpAddresses(v *[]models.WirelessControllerAddrgrpAddresses, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectWirelessControllerAddrgrp(d *schema.ResourceData, o *models.WirelessControllerAddrgrp, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Addresses != nil {
		if err = d.Set("addresses", flattenWirelessControllerAddrgrpAddresses(o.Addresses, sort)); err != nil {
			return diag.Errorf("error reading addresses: %v", err)
		}
	}

	if o.DefaultPolicy != nil {
		v := *o.DefaultPolicy

		if err = d.Set("default_policy", v); err != nil {
			return diag.Errorf("error reading default_policy: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerAddrgrpAddresses(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerAddrgrpAddresses, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerAddrgrpAddresses

	for i := range l {
		tmp := models.WirelessControllerAddrgrpAddresses{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerAddrgrp(d *schema.ResourceData, sv string) (*models.WirelessControllerAddrgrp, diag.Diagnostics) {
	obj := models.WirelessControllerAddrgrp{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("addresses"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("addresses", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerAddrgrpAddresses(d, v, "addresses", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Addresses = t
		}
	} else if d.HasChange("addresses") {
		old, new := d.GetChange("addresses")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Addresses = &[]models.WirelessControllerAddrgrpAddresses{}
		}
	}
	if v1, ok := d.GetOk("default_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_policy", sv)
				diags = append(diags, e)
			}
			obj.DefaultPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			obj.Id = &v2
		}
	}
	return &obj, diags
}
