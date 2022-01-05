// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FSSO groups.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceUserAdgrp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FSSO groups.",

		CreateContext: resourceUserAdgrpCreate,
		ReadContext:   resourceUserAdgrpRead,
		UpdateContext: resourceUserAdgrpUpdate,
		DeleteContext: resourceUserAdgrpDelete,

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
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"connector_source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "FSSO connector source.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Group ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"server_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "FSSO agent name.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserAdgrpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating UserAdgrp resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserAdgrp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserAdgrp(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserAdgrp")
	}

	return resourceUserAdgrpRead(ctx, d, meta)
}

func resourceUserAdgrpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserAdgrp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserAdgrp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserAdgrp resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserAdgrp")
	}

	return resourceUserAdgrpRead(ctx, d, meta)
}

func resourceUserAdgrpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserAdgrp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserAdgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserAdgrpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadUserAdgrp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserAdgrp resource: %v", err)
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

	diags := refreshObjectUserAdgrp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectUserAdgrp(d *schema.ResourceData, o *models.UserAdgrp, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ConnectorSource != nil {
		v := *o.ConnectorSource

		if err = d.Set("connector_source", v); err != nil {
			return diag.Errorf("error reading connector_source: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ServerName != nil {
		v := *o.ServerName

		if err = d.Set("server_name", v); err != nil {
			return diag.Errorf("error reading server_name: %v", err)
		}
	}

	return nil
}

func getObjectUserAdgrp(d *schema.ResourceData, sv string) (*models.UserAdgrp, diag.Diagnostics) {
	obj := models.UserAdgrp{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("connector_source"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("connector_source", sv)
				diags = append(diags, e)
			}
			obj.ConnectorSource = &v2
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("server_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_name", sv)
				diags = append(diags, e)
			}
			obj.ServerName = &v2
		}
	}
	return &obj, diags
}
