// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure web filter search engines.

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

func resourceWebfilterSearchEngine() *schema.Resource {
	return &schema.Resource{
		Description: "Configure web filter search engines.",

		CreateContext: resourceWebfilterSearchEngineCreate,
		ReadContext:   resourceWebfilterSearchEngineRead,
		UpdateContext: resourceWebfilterSearchEngineUpdate,
		DeleteContext: resourceWebfilterSearchEngineDelete,

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
			"charset": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"utf-8", "gb2312"}, false),

				Description: "Search engine charset.",
				Optional:    true,
				Computed:    true,
			},
			"hostname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Hostname (regular expression).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Search engine name.",
				ForceNew:    true,
				Required:    true,
			},
			"query": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Code used to prefix a query (must end with an equals character).",
				Optional:    true,
				Computed:    true,
			},
			"safesearch": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "url", "header", "translate", "yt-pattern", "yt-scan", "yt-video", "yt-channel"}, false),

				Description: "Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.",
				Optional:    true,
				Computed:    true,
			},
			"safesearch_str": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Safe search parameter used in the URL.",
				Optional:    true,
				Computed:    true,
			},
			"url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "URL (regular expression).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWebfilterSearchEngineCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WebfilterSearchEngine resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWebfilterSearchEngine(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWebfilterSearchEngine(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterSearchEngine")
	}

	return resourceWebfilterSearchEngineRead(ctx, d, meta)
}

func resourceWebfilterSearchEngineUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebfilterSearchEngine(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWebfilterSearchEngine(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WebfilterSearchEngine resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterSearchEngine")
	}

	return resourceWebfilterSearchEngineRead(ctx, d, meta)
}

func resourceWebfilterSearchEngineDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWebfilterSearchEngine(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WebfilterSearchEngine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterSearchEngineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebfilterSearchEngine(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterSearchEngine resource: %v", err)
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

	diags := refreshObjectWebfilterSearchEngine(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWebfilterSearchEngine(d *schema.ResourceData, o *models.WebfilterSearchEngine, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Charset != nil {
		v := *o.Charset

		if err = d.Set("charset", v); err != nil {
			return diag.Errorf("error reading charset: %v", err)
		}
	}

	if o.Hostname != nil {
		v := *o.Hostname

		if err = d.Set("hostname", v); err != nil {
			return diag.Errorf("error reading hostname: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Query != nil {
		v := *o.Query

		if err = d.Set("query", v); err != nil {
			return diag.Errorf("error reading query: %v", err)
		}
	}

	if o.Safesearch != nil {
		v := *o.Safesearch

		if err = d.Set("safesearch", v); err != nil {
			return diag.Errorf("error reading safesearch: %v", err)
		}
	}

	if o.SafesearchStr != nil {
		v := *o.SafesearchStr

		if err = d.Set("safesearch_str", v); err != nil {
			return diag.Errorf("error reading safesearch_str: %v", err)
		}
	}

	if o.Url != nil {
		v := *o.Url

		if err = d.Set("url", v); err != nil {
			return diag.Errorf("error reading url: %v", err)
		}
	}

	return nil
}

func getObjectWebfilterSearchEngine(d *schema.ResourceData, sv string) (*models.WebfilterSearchEngine, diag.Diagnostics) {
	obj := models.WebfilterSearchEngine{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("charset"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("charset", sv)
				diags = append(diags, e)
			}
			obj.Charset = &v2
		}
	}
	if v1, ok := d.GetOk("hostname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hostname", sv)
				diags = append(diags, e)
			}
			obj.Hostname = &v2
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
	if v1, ok := d.GetOk("query"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("query", sv)
				diags = append(diags, e)
			}
			obj.Query = &v2
		}
	}
	if v1, ok := d.GetOk("safesearch"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("safesearch", sv)
				diags = append(diags, e)
			}
			obj.Safesearch = &v2
		}
	}
	if v1, ok := d.GetOk("safesearch_str"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("safesearch_str", sv)
				diags = append(diags, e)
			}
			obj.SafesearchStr = &v2
		}
	}
	if v1, ok := d.GetOk("url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("url", sv)
				diags = append(diags, e)
			}
			obj.Url = &v2
		}
	}
	return &obj, diags
}
