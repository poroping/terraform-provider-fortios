// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Remote certificate as a PEM file.

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

func resourceVpncertificateRemote() *schema.Resource {
	return &schema.Resource{
		Description: "Remote certificate as a PEM file.",

		CreateContext: resourceVpncertificateRemoteCreate,
		ReadContext:   resourceVpncertificateRemoteRead,
		UpdateContext: resourceVpncertificateRemoteUpdate,
		DeleteContext: resourceVpncertificateRemoteDelete,

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
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"range": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"global", "vdom"}, false),

				Description: "Either the global or VDOM IP address range for the remote certificate.",
				Optional:    true,
				Computed:    true,
			},
			"remote": {
				Type: schema.TypeString,

				Description: "Remote certificate.",
				Optional:    true,
				Computed:    true,
			},
			"source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"factory", "user", "bundle"}, false),

				Description: "Remote certificate source type.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpncertificateRemoteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating VpncertificateRemote resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpncertificateRemote(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpncertificateRemote(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpncertificateRemote")
	}

	return resourceVpncertificateRemoteRead(ctx, d, meta)
}

func resourceVpncertificateRemoteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpncertificateRemote(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpncertificateRemote(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpncertificateRemote resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpncertificateRemote")
	}

	return resourceVpncertificateRemoteRead(ctx, d, meta)
}

func resourceVpncertificateRemoteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVpncertificateRemote(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpncertificateRemote resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpncertificateRemoteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpncertificateRemote(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpncertificateRemote resource: %v", err)
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

	diags := refreshObjectVpncertificateRemote(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectVpncertificateRemote(d *schema.ResourceData, o *models.VpncertificateRemote, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Range != nil {
		v := *o.Range

		if err = d.Set("range", v); err != nil {
			return diag.Errorf("error reading range: %v", err)
		}
	}

	if o.Remote != nil {
		v := *o.Remote

		if err = d.Set("remote", v); err != nil {
			return diag.Errorf("error reading remote: %v", err)
		}
	}

	if o.Source != nil {
		v := *o.Source

		if err = d.Set("source", v); err != nil {
			return diag.Errorf("error reading source: %v", err)
		}
	}

	return nil
}

func getObjectVpncertificateRemote(d *schema.ResourceData, sv string) (*models.VpncertificateRemote, diag.Diagnostics) {
	obj := models.VpncertificateRemote{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("range"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("range", sv)
				diags = append(diags, e)
			}
			obj.Range = &v2
		}
	}
	if v1, ok := d.GetOk("remote"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remote", sv)
				diags = append(diags, e)
			}
			obj.Remote = &v2
		}
	}
	if v1, ok := d.GetOk("source"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source", sv)
				diags = append(diags, e)
			}
			obj.Source = &v2
		}
	}
	return &obj, diags
}
