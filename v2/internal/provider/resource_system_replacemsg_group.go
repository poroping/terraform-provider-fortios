// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure replacement message groups.

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

func resourceSystemReplacemsgGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure replacement message groups.",

		CreateContext: resourceSystemReplacemsgGroupCreate,
		ReadContext:   resourceSystemReplacemsgGroupRead,
		UpdateContext: resourceSystemReplacemsgGroupUpdate,
		DeleteContext: resourceSystemReplacemsgGroupDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"admin": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"alertmail": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"auth": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"automation": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"custom_message": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"device_detection_portal": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fortiguard_wf": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ftp": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"group_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "utm", "auth"}, false),

				Description: "Group type.",
				Optional:    true,
				Computed:    true,
			},
			"http": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"icap": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"mail": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"nac_quar": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Group name.",
				ForceNew:    true,
				Required:    true,
			},
			"nntp": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"spam": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"sslvpn": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"traffic_quota": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"utm": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"webproxy": {
				Type:        schema.TypeList,
				Description: "Replacement message table entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),

							Description: "Message string.",
							Optional:    true,
							Computed:    true,
						},
						"format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

							Description: "Format flag.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

							Description: "Header flag.",
							Optional:    true,
							Computed:    true,
						},
						"msg_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),

							Description: "Message type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemReplacemsgGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemReplacemsgGroup resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemReplacemsgGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemReplacemsgGroup(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemReplacemsgGroup")
	}

	return resourceSystemReplacemsgGroupRead(ctx, d, meta)
}

func resourceSystemReplacemsgGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemReplacemsgGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemReplacemsgGroup(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemReplacemsgGroup resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemReplacemsgGroup")
	}

	return resourceSystemReplacemsgGroupRead(ctx, d, meta)
}

func resourceSystemReplacemsgGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemReplacemsgGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemReplacemsgGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemReplacemsgGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemReplacemsgGroup resource: %v", err)
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

	diags := refreshObjectSystemReplacemsgGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemReplacemsgGroupAdmin(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupAdmin, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupAlertmail(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupAlertmail, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupAuth(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupAuth, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupAutomation(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupAutomation, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupCustomMessage(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupCustomMessage, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupDeviceDetectionPortal(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupDeviceDetectionPortal, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupFortiguardWf(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupFortiguardWf, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupFtp(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupFtp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupHttp(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupHttp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupIcap(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupIcap, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupMail(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupMail, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupNacQuar(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupNacQuar, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupNntp(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupNntp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupSpam(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupSpam, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupSslvpn(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupSslvpn, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupTrafficQuota(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupTrafficQuota, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupUtm(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupUtm, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func flattenSystemReplacemsgGroupWebproxy(d *schema.ResourceData, v *[]models.SystemReplacemsgGroupWebproxy, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Buffer; tmp != nil {
				v["buffer"] = *tmp
			}

			if tmp := cfg.Format; tmp != nil {
				v["format"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.MsgType; tmp != nil {
				v["msg_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "msg_type")
	}

	return flat
}

func refreshObjectSystemReplacemsgGroup(d *schema.ResourceData, o *models.SystemReplacemsgGroup, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Admin != nil {
		if err = d.Set("admin", flattenSystemReplacemsgGroupAdmin(d, o.Admin, "admin", sort)); err != nil {
			return diag.Errorf("error reading admin: %v", err)
		}
	}

	if o.Alertmail != nil {
		if err = d.Set("alertmail", flattenSystemReplacemsgGroupAlertmail(d, o.Alertmail, "alertmail", sort)); err != nil {
			return diag.Errorf("error reading alertmail: %v", err)
		}
	}

	if o.Auth != nil {
		if err = d.Set("auth", flattenSystemReplacemsgGroupAuth(d, o.Auth, "auth", sort)); err != nil {
			return diag.Errorf("error reading auth: %v", err)
		}
	}

	if o.Automation != nil {
		if err = d.Set("automation", flattenSystemReplacemsgGroupAutomation(d, o.Automation, "automation", sort)); err != nil {
			return diag.Errorf("error reading automation: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.CustomMessage != nil {
		if err = d.Set("custom_message", flattenSystemReplacemsgGroupCustomMessage(d, o.CustomMessage, "custom_message", sort)); err != nil {
			return diag.Errorf("error reading custom_message: %v", err)
		}
	}

	if o.DeviceDetectionPortal != nil {
		if err = d.Set("device_detection_portal", flattenSystemReplacemsgGroupDeviceDetectionPortal(d, o.DeviceDetectionPortal, "device_detection_portal", sort)); err != nil {
			return diag.Errorf("error reading device_detection_portal: %v", err)
		}
	}

	if o.FortiguardWf != nil {
		if err = d.Set("fortiguard_wf", flattenSystemReplacemsgGroupFortiguardWf(d, o.FortiguardWf, "fortiguard_wf", sort)); err != nil {
			return diag.Errorf("error reading fortiguard_wf: %v", err)
		}
	}

	if o.Ftp != nil {
		if err = d.Set("ftp", flattenSystemReplacemsgGroupFtp(d, o.Ftp, "ftp", sort)); err != nil {
			return diag.Errorf("error reading ftp: %v", err)
		}
	}

	if o.GroupType != nil {
		v := *o.GroupType

		if err = d.Set("group_type", v); err != nil {
			return diag.Errorf("error reading group_type: %v", err)
		}
	}

	if o.Http != nil {
		if err = d.Set("http", flattenSystemReplacemsgGroupHttp(d, o.Http, "http", sort)); err != nil {
			return diag.Errorf("error reading http: %v", err)
		}
	}

	if o.Icap != nil {
		if err = d.Set("icap", flattenSystemReplacemsgGroupIcap(d, o.Icap, "icap", sort)); err != nil {
			return diag.Errorf("error reading icap: %v", err)
		}
	}

	if o.Mail != nil {
		if err = d.Set("mail", flattenSystemReplacemsgGroupMail(d, o.Mail, "mail", sort)); err != nil {
			return diag.Errorf("error reading mail: %v", err)
		}
	}

	if o.NacQuar != nil {
		if err = d.Set("nac_quar", flattenSystemReplacemsgGroupNacQuar(d, o.NacQuar, "nac_quar", sort)); err != nil {
			return diag.Errorf("error reading nac_quar: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Nntp != nil {
		if err = d.Set("nntp", flattenSystemReplacemsgGroupNntp(d, o.Nntp, "nntp", sort)); err != nil {
			return diag.Errorf("error reading nntp: %v", err)
		}
	}

	if o.Spam != nil {
		if err = d.Set("spam", flattenSystemReplacemsgGroupSpam(d, o.Spam, "spam", sort)); err != nil {
			return diag.Errorf("error reading spam: %v", err)
		}
	}

	if o.Sslvpn != nil {
		if err = d.Set("sslvpn", flattenSystemReplacemsgGroupSslvpn(d, o.Sslvpn, "sslvpn", sort)); err != nil {
			return diag.Errorf("error reading sslvpn: %v", err)
		}
	}

	if o.TrafficQuota != nil {
		if err = d.Set("traffic_quota", flattenSystemReplacemsgGroupTrafficQuota(d, o.TrafficQuota, "traffic_quota", sort)); err != nil {
			return diag.Errorf("error reading traffic_quota: %v", err)
		}
	}

	if o.Utm != nil {
		if err = d.Set("utm", flattenSystemReplacemsgGroupUtm(d, o.Utm, "utm", sort)); err != nil {
			return diag.Errorf("error reading utm: %v", err)
		}
	}

	if o.Webproxy != nil {
		if err = d.Set("webproxy", flattenSystemReplacemsgGroupWebproxy(d, o.Webproxy, "webproxy", sort)); err != nil {
			return diag.Errorf("error reading webproxy: %v", err)
		}
	}

	return nil
}

func expandSystemReplacemsgGroupAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupAdmin, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupAdmin

	for i := range l {
		tmp := models.SystemReplacemsgGroupAdmin{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupAlertmail(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupAlertmail, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupAlertmail

	for i := range l {
		tmp := models.SystemReplacemsgGroupAlertmail{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupAuth, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupAuth

	for i := range l {
		tmp := models.SystemReplacemsgGroupAuth{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupAutomation(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupAutomation, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupAutomation

	for i := range l {
		tmp := models.SystemReplacemsgGroupAutomation{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupCustomMessage(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupCustomMessage, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupCustomMessage

	for i := range l {
		tmp := models.SystemReplacemsgGroupCustomMessage{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupDeviceDetectionPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupDeviceDetectionPortal, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupDeviceDetectionPortal

	for i := range l {
		tmp := models.SystemReplacemsgGroupDeviceDetectionPortal{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupFortiguardWf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupFortiguardWf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupFortiguardWf

	for i := range l {
		tmp := models.SystemReplacemsgGroupFortiguardWf{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupFtp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupFtp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupFtp

	for i := range l {
		tmp := models.SystemReplacemsgGroupFtp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupHttp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupHttp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupHttp

	for i := range l {
		tmp := models.SystemReplacemsgGroupHttp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupIcap(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupIcap, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupIcap

	for i := range l {
		tmp := models.SystemReplacemsgGroupIcap{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupMail(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupMail, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupMail

	for i := range l {
		tmp := models.SystemReplacemsgGroupMail{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupNacQuar(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupNacQuar, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupNacQuar

	for i := range l {
		tmp := models.SystemReplacemsgGroupNacQuar{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupNntp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupNntp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupNntp

	for i := range l {
		tmp := models.SystemReplacemsgGroupNntp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupSpam(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupSpam, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupSpam

	for i := range l {
		tmp := models.SystemReplacemsgGroupSpam{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupSslvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupSslvpn, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupSslvpn

	for i := range l {
		tmp := models.SystemReplacemsgGroupSslvpn{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupTrafficQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupTrafficQuota, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupTrafficQuota

	for i := range l {
		tmp := models.SystemReplacemsgGroupTrafficQuota{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupUtm(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupUtm, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupUtm

	for i := range l {
		tmp := models.SystemReplacemsgGroupUtm{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemReplacemsgGroupWebproxy(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemReplacemsgGroupWebproxy, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemReplacemsgGroupWebproxy

	for i := range l {
		tmp := models.SystemReplacemsgGroupWebproxy{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.buffer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Buffer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Format = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.msg_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MsgType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemReplacemsgGroup(d *schema.ResourceData, sv string) (*models.SystemReplacemsgGroup, diag.Diagnostics) {
	obj := models.SystemReplacemsgGroup{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("admin"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("admin", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupAdmin(d, v, "admin", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Admin = t
		}
	} else if d.HasChange("admin") {
		old, new := d.GetChange("admin")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Admin = &[]models.SystemReplacemsgGroupAdmin{}
		}
	}
	if v, ok := d.GetOk("alertmail"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("alertmail", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupAlertmail(d, v, "alertmail", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Alertmail = t
		}
	} else if d.HasChange("alertmail") {
		old, new := d.GetChange("alertmail")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Alertmail = &[]models.SystemReplacemsgGroupAlertmail{}
		}
	}
	if v, ok := d.GetOk("auth"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("auth", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupAuth(d, v, "auth", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Auth = t
		}
	} else if d.HasChange("auth") {
		old, new := d.GetChange("auth")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Auth = &[]models.SystemReplacemsgGroupAuth{}
		}
	}
	if v, ok := d.GetOk("automation"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("automation", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupAutomation(d, v, "automation", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Automation = t
		}
	} else if d.HasChange("automation") {
		old, new := d.GetChange("automation")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Automation = &[]models.SystemReplacemsgGroupAutomation{}
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v, ok := d.GetOk("custom_message"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("custom_message", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupCustomMessage(d, v, "custom_message", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.CustomMessage = t
		}
	} else if d.HasChange("custom_message") {
		old, new := d.GetChange("custom_message")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.CustomMessage = &[]models.SystemReplacemsgGroupCustomMessage{}
		}
	}
	if v, ok := d.GetOk("device_detection_portal"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("device_detection_portal", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupDeviceDetectionPortal(d, v, "device_detection_portal", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DeviceDetectionPortal = t
		}
	} else if d.HasChange("device_detection_portal") {
		old, new := d.GetChange("device_detection_portal")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DeviceDetectionPortal = &[]models.SystemReplacemsgGroupDeviceDetectionPortal{}
		}
	}
	if v, ok := d.GetOk("fortiguard_wf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("fortiguard_wf", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupFortiguardWf(d, v, "fortiguard_wf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FortiguardWf = t
		}
	} else if d.HasChange("fortiguard_wf") {
		old, new := d.GetChange("fortiguard_wf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FortiguardWf = &[]models.SystemReplacemsgGroupFortiguardWf{}
		}
	}
	if v, ok := d.GetOk("ftp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ftp", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupFtp(d, v, "ftp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ftp = t
		}
	} else if d.HasChange("ftp") {
		old, new := d.GetChange("ftp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ftp = &[]models.SystemReplacemsgGroupFtp{}
		}
	}
	if v1, ok := d.GetOk("group_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_type", sv)
				diags = append(diags, e)
			}
			obj.GroupType = &v2
		}
	}
	if v, ok := d.GetOk("http"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("http", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupHttp(d, v, "http", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Http = t
		}
	} else if d.HasChange("http") {
		old, new := d.GetChange("http")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Http = &[]models.SystemReplacemsgGroupHttp{}
		}
	}
	if v, ok := d.GetOk("icap"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("icap", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupIcap(d, v, "icap", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Icap = t
		}
	} else if d.HasChange("icap") {
		old, new := d.GetChange("icap")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Icap = &[]models.SystemReplacemsgGroupIcap{}
		}
	}
	if v, ok := d.GetOk("mail"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("mail", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupMail(d, v, "mail", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Mail = t
		}
	} else if d.HasChange("mail") {
		old, new := d.GetChange("mail")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Mail = &[]models.SystemReplacemsgGroupMail{}
		}
	}
	if v, ok := d.GetOk("nac_quar"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("nac_quar", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupNacQuar(d, v, "nac_quar", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NacQuar = t
		}
	} else if d.HasChange("nac_quar") {
		old, new := d.GetChange("nac_quar")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NacQuar = &[]models.SystemReplacemsgGroupNacQuar{}
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
	if v, ok := d.GetOk("nntp"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("nntp", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupNntp(d, v, "nntp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Nntp = t
		}
	} else if d.HasChange("nntp") {
		old, new := d.GetChange("nntp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Nntp = &[]models.SystemReplacemsgGroupNntp{}
		}
	}
	if v, ok := d.GetOk("spam"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("spam", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupSpam(d, v, "spam", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Spam = t
		}
	} else if d.HasChange("spam") {
		old, new := d.GetChange("spam")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Spam = &[]models.SystemReplacemsgGroupSpam{}
		}
	}
	if v, ok := d.GetOk("sslvpn"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("sslvpn", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupSslvpn(d, v, "sslvpn", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Sslvpn = t
		}
	} else if d.HasChange("sslvpn") {
		old, new := d.GetChange("sslvpn")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Sslvpn = &[]models.SystemReplacemsgGroupSslvpn{}
		}
	}
	if v, ok := d.GetOk("traffic_quota"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("traffic_quota", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupTrafficQuota(d, v, "traffic_quota", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.TrafficQuota = t
		}
	} else if d.HasChange("traffic_quota") {
		old, new := d.GetChange("traffic_quota")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.TrafficQuota = &[]models.SystemReplacemsgGroupTrafficQuota{}
		}
	}
	if v, ok := d.GetOk("utm"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("utm", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupUtm(d, v, "utm", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Utm = t
		}
	} else if d.HasChange("utm") {
		old, new := d.GetChange("utm")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Utm = &[]models.SystemReplacemsgGroupUtm{}
		}
	}
	if v, ok := d.GetOk("webproxy"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("webproxy", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemReplacemsgGroupWebproxy(d, v, "webproxy", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Webproxy = t
		}
	} else if d.HasChange("webproxy") {
		old, new := d.GetChange("webproxy")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Webproxy = &[]models.SystemReplacemsgGroupWebproxy{}
		}
	}
	return &obj, diags
}
