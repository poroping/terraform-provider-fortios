// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.6,v7.0.2 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure USB auto installation.

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

func resourceSystemAutoInstall() *schema.Resource {
	return &schema.Resource{
		Description: "Configure USB auto installation.",

		CreateContext: resourceSystemAutoInstallCreate,
		ReadContext:   resourceSystemAutoInstallRead,
		UpdateContext: resourceSystemAutoInstallUpdate,
		DeleteContext: resourceSystemAutoInstallDelete,

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
			"auto_install_config": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable auto install the config in USB disk.",
				Optional:    true,
				Computed:    true,
			},
			"auto_install_image": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable auto install the image in USB disk.",
				Optional:    true,
				Computed:    true,
			},
			"default_config_file": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Default config file name in USB disk.",
				Optional:    true,
				Computed:    true,
			},
			"default_image_file": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Default image file name in USB disk.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemAutoInstallCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAutoInstall(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemAutoInstall(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAutoInstall")
	}

	return resourceSystemAutoInstallRead(ctx, d, meta)
}

func resourceSystemAutoInstallUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAutoInstall(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemAutoInstall(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemAutoInstall resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAutoInstall")
	}

	return resourceSystemAutoInstallRead(ctx, d, meta)
}

func resourceSystemAutoInstallDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemAutoInstall(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemAutoInstall resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoInstallRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAutoInstall(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAutoInstall resource: %v", err)
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

	diags := refreshObjectSystemAutoInstall(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemAutoInstall(d *schema.ResourceData, o *models.SystemAutoInstall, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AutoInstallConfig != nil {
		v := *o.AutoInstallConfig

		if err = d.Set("auto_install_config", v); err != nil {
			return diag.Errorf("error reading auto_install_config: %v", err)
		}
	}

	if o.AutoInstallImage != nil {
		v := *o.AutoInstallImage

		if err = d.Set("auto_install_image", v); err != nil {
			return diag.Errorf("error reading auto_install_image: %v", err)
		}
	}

	if o.DefaultConfigFile != nil {
		v := *o.DefaultConfigFile

		if err = d.Set("default_config_file", v); err != nil {
			return diag.Errorf("error reading default_config_file: %v", err)
		}
	}

	if o.DefaultImageFile != nil {
		v := *o.DefaultImageFile

		if err = d.Set("default_image_file", v); err != nil {
			return diag.Errorf("error reading default_image_file: %v", err)
		}
	}

	return nil
}

func getObjectSystemAutoInstall(d *schema.ResourceData, sv string) (*models.SystemAutoInstall, diag.Diagnostics) {
	obj := models.SystemAutoInstall{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auto_install_config"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_install_config", sv)
				diags = append(diags, e)
			}
			obj.AutoInstallConfig = &v2
		}
	}
	if v1, ok := d.GetOk("auto_install_image"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_install_image", sv)
				diags = append(diags, e)
			}
			obj.AutoInstallImage = &v2
		}
	}
	if v1, ok := d.GetOk("default_config_file"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_config_file", sv)
				diags = append(diags, e)
			}
			obj.DefaultConfigFile = &v2
		}
	}
	if v1, ok := d.GetOk("default_image_file"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_image_file", sv)
				diags = append(diags, e)
			}
			obj.DefaultImageFile = &v2
		}
	}
	return &obj, diags
}
