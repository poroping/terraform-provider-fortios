package provider

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	sdkutils "github.com/poroping/forti-sdk-go/v2/utils"
)

func resourceGenericApi() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 addresses.",

		CreateContext: resourceGenericApiCreate,
		ReadContext:   schema.NoopContext,
		UpdateContext: resourceGenericApiCreate,
		DeleteContext: schema.NoopContext,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"path": {
				Type:        schema.TypeString,
				Description: "Full path including any URL params. Excluding VDOM.",
				Required:    true,
			},
			"method": {
				Type:        schema.TypeString,
				Description: "HTTP Method.",
				Required:    true,
			},
			"force_recreate": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"json": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"response": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGenericApiCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	js := d.Get("json").(string)
	if js != "" {
		var jsnull json.RawMessage
		err := json.Unmarshal([]byte(js), &jsnull)
		if err != nil {
			return diag.Errorf("provided JSON is invalid")
		}
	}
	payload := &sdkutils.UtilsGenericApi{}
	payload.Path = d.Get("path").(string)
	payload.HTTPMethod = d.Get("method").(string)
	payload.Json = js

	o, err := c.Utils.GenericApi(payload, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	d.SetId(uuid.New().String())
	err = d.Set("response", *o)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
