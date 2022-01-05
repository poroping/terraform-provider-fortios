// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiGuard - AntiSpam.

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

func resourceEmailfilterFortishield() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiGuard - AntiSpam.",

		CreateContext: resourceEmailfilterFortishieldCreate,
		ReadContext:   resourceEmailfilterFortishieldRead,
		UpdateContext: resourceEmailfilterFortishieldUpdate,
		DeleteContext: resourceEmailfilterFortishieldDelete,

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
			"spam_submit_force": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable force insertion of a new mime entity for the submission text.",
				Optional:    true,
				Computed:    true,
			},
			"spam_submit_srv": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Hostname of the spam submission server.",
				Optional:    true,
				Computed:    true,
			},
			"spam_submit_txt2htm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable conversion of text email to HTML email.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceEmailfilterFortishieldCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectEmailfilterFortishield(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateEmailfilterFortishield(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EmailfilterFortishield")
	}

	return resourceEmailfilterFortishieldRead(ctx, d, meta)
}

func resourceEmailfilterFortishieldUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectEmailfilterFortishield(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateEmailfilterFortishield(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating EmailfilterFortishield resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EmailfilterFortishield")
	}

	return resourceEmailfilterFortishieldRead(ctx, d, meta)
}

func resourceEmailfilterFortishieldDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectEmailfilterFortishield(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateEmailfilterFortishield(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting EmailfilterFortishield resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterFortishieldRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadEmailfilterFortishield(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading EmailfilterFortishield resource: %v", err)
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

	diags := refreshObjectEmailfilterFortishield(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectEmailfilterFortishield(d *schema.ResourceData, o *models.EmailfilterFortishield, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.SpamSubmitForce != nil {
		v := *o.SpamSubmitForce

		if err = d.Set("spam_submit_force", v); err != nil {
			return diag.Errorf("error reading spam_submit_force: %v", err)
		}
	}

	if o.SpamSubmitSrv != nil {
		v := *o.SpamSubmitSrv

		if err = d.Set("spam_submit_srv", v); err != nil {
			return diag.Errorf("error reading spam_submit_srv: %v", err)
		}
	}

	if o.SpamSubmitTxt2htm != nil {
		v := *o.SpamSubmitTxt2htm

		if err = d.Set("spam_submit_txt2htm", v); err != nil {
			return diag.Errorf("error reading spam_submit_txt2htm: %v", err)
		}
	}

	return nil
}

func getObjectEmailfilterFortishield(d *schema.ResourceData, sv string) (*models.EmailfilterFortishield, diag.Diagnostics) {
	obj := models.EmailfilterFortishield{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("spam_submit_force"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spam_submit_force", sv)
				diags = append(diags, e)
			}
			obj.SpamSubmitForce = &v2
		}
	}
	if v1, ok := d.GetOk("spam_submit_srv"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spam_submit_srv", sv)
				diags = append(diags, e)
			}
			obj.SpamSubmitSrv = &v2
		}
	}
	if v1, ok := d.GetOk("spam_submit_txt2htm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spam_submit_txt2htm", sv)
				diags = append(diags, e)
			}
			obj.SpamSubmitTxt2htm = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectEmailfilterFortishield(d *schema.ResourceData, sv string) (*models.EmailfilterFortishield, diag.Diagnostics) {
	obj := models.EmailfilterFortishield{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
