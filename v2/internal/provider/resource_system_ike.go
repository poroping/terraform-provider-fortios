// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IKE global attributes.

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

func resourceSystemIke() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IKE global attributes.",

		CreateContext: resourceSystemIkeCreate,
		ReadContext:   resourceSystemIkeRead,
		UpdateContext: resourceSystemIkeUpdate,
		DeleteContext: resourceSystemIkeDelete,

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
			"dh_group_1": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 1 (MODP-768).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_14": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 14 (MODP-2048).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_15": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 15 (MODP-3072).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_16": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 16 (MODP-4096).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_17": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 17 (MODP-6144).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_18": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 18 (MODP-8192).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_19": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 19 (EC-P256).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_2": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 2 (MODP-1024).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_20": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 20 (EC-P384).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_21": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 21 (EC-P521).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_27": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 27 (EC-P224BP).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_28": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 28 (EC-P256BP).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_29": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 29 (EC-P384BP).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_30": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 30 (EC-P512BP).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_31": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 31 (EC-X25519).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_32": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 32 (EC-X448).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_group_5": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 5 (MODP-1536).",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "custom"}, false),

							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
						"keypair_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),

							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"software", "hardware", "global"}, false),

							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dh_keypair_cache": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Diffie-Hellman key pair cache.",
				Optional:    true,
				Computed:    true,
			},
			"dh_keypair_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 50000),

				Description: "Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).",
				Optional:    true,
				Computed:    true,
			},
			"dh_keypair_throttle": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Diffie-Hellman key pair cache CPU throttling.",
				Optional:    true,
				Computed:    true,
			},
			"dh_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"software", "hardware"}, false),

				Description: "Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations.",
				Optional:    true,
				Computed:    true,
			},
			"dh_multiprocess": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable multiprocess Diffie-Hellman daemon for IKE.",
				Optional:    true,
				Computed:    true,
			},
			"dh_worker_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 32),

				Description: "Number of Diffie-Hellman workers to start.",
				Optional:    true,
				Computed:    true,
			},
			"embryonic_limit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(50, 20000),

				Description: "Maximum number of IPsec tunnels to negotiate simultaneously.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemIkeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemIke(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemIke(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIke")
	}

	return resourceSystemIkeRead(ctx, d, meta)
}

func resourceSystemIkeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemIke(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemIke(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemIke resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIke")
	}

	return resourceSystemIkeRead(ctx, d, meta)
}

func resourceSystemIkeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemIke(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemIke(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemIke resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIkeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemIke(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemIke resource: %v", err)
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

	diags := refreshObjectSystemIke(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemIkeDhGroup1(d *schema.ResourceData, v *models.SystemIkeDhGroup1, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup1{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup14(d *schema.ResourceData, v *models.SystemIkeDhGroup14, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup14{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup15(d *schema.ResourceData, v *models.SystemIkeDhGroup15, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup15{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup16(d *schema.ResourceData, v *models.SystemIkeDhGroup16, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup16{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup17(d *schema.ResourceData, v *models.SystemIkeDhGroup17, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup17{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup18(d *schema.ResourceData, v *models.SystemIkeDhGroup18, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup18{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup19(d *schema.ResourceData, v *models.SystemIkeDhGroup19, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup19{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup2(d *schema.ResourceData, v *models.SystemIkeDhGroup2, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup2{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup20(d *schema.ResourceData, v *models.SystemIkeDhGroup20, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup20{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup21(d *schema.ResourceData, v *models.SystemIkeDhGroup21, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup21{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup27(d *schema.ResourceData, v *models.SystemIkeDhGroup27, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup27{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup28(d *schema.ResourceData, v *models.SystemIkeDhGroup28, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup28{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup29(d *schema.ResourceData, v *models.SystemIkeDhGroup29, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup29{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup30(d *schema.ResourceData, v *models.SystemIkeDhGroup30, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup30{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup31(d *schema.ResourceData, v *models.SystemIkeDhGroup31, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup31{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup32(d *schema.ResourceData, v *models.SystemIkeDhGroup32, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup32{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemIkeDhGroup5(d *schema.ResourceData, v *models.SystemIkeDhGroup5, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemIkeDhGroup5{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.KeypairCache; tmp != nil {
				v["keypair_cache"] = *tmp
			}

			if tmp := cfg.KeypairCount; tmp != nil {
				v["keypair_count"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectSystemIke(d *schema.ResourceData, o *models.SystemIke, sv string, sort bool) diag.Diagnostics {
	var err error

	if _, ok := d.GetOk("dh_group_1"); ok {
		if o.DhGroup1 != nil {
			if err = d.Set("dh_group_1", flattenSystemIkeDhGroup1(d, o.DhGroup1, "dh_group_1", sort)); err != nil {
				return diag.Errorf("error reading dh_group_1: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_14"); ok {
		if o.DhGroup14 != nil {
			if err = d.Set("dh_group_14", flattenSystemIkeDhGroup14(d, o.DhGroup14, "dh_group_14", sort)); err != nil {
				return diag.Errorf("error reading dh_group_14: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_15"); ok {
		if o.DhGroup15 != nil {
			if err = d.Set("dh_group_15", flattenSystemIkeDhGroup15(d, o.DhGroup15, "dh_group_15", sort)); err != nil {
				return diag.Errorf("error reading dh_group_15: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_16"); ok {
		if o.DhGroup16 != nil {
			if err = d.Set("dh_group_16", flattenSystemIkeDhGroup16(d, o.DhGroup16, "dh_group_16", sort)); err != nil {
				return diag.Errorf("error reading dh_group_16: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_17"); ok {
		if o.DhGroup17 != nil {
			if err = d.Set("dh_group_17", flattenSystemIkeDhGroup17(d, o.DhGroup17, "dh_group_17", sort)); err != nil {
				return diag.Errorf("error reading dh_group_17: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_18"); ok {
		if o.DhGroup18 != nil {
			if err = d.Set("dh_group_18", flattenSystemIkeDhGroup18(d, o.DhGroup18, "dh_group_18", sort)); err != nil {
				return diag.Errorf("error reading dh_group_18: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_19"); ok {
		if o.DhGroup19 != nil {
			if err = d.Set("dh_group_19", flattenSystemIkeDhGroup19(d, o.DhGroup19, "dh_group_19", sort)); err != nil {
				return diag.Errorf("error reading dh_group_19: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_2"); ok {
		if o.DhGroup2 != nil {
			if err = d.Set("dh_group_2", flattenSystemIkeDhGroup2(d, o.DhGroup2, "dh_group_2", sort)); err != nil {
				return diag.Errorf("error reading dh_group_2: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_20"); ok {
		if o.DhGroup20 != nil {
			if err = d.Set("dh_group_20", flattenSystemIkeDhGroup20(d, o.DhGroup20, "dh_group_20", sort)); err != nil {
				return diag.Errorf("error reading dh_group_20: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_21"); ok {
		if o.DhGroup21 != nil {
			if err = d.Set("dh_group_21", flattenSystemIkeDhGroup21(d, o.DhGroup21, "dh_group_21", sort)); err != nil {
				return diag.Errorf("error reading dh_group_21: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_27"); ok {
		if o.DhGroup27 != nil {
			if err = d.Set("dh_group_27", flattenSystemIkeDhGroup27(d, o.DhGroup27, "dh_group_27", sort)); err != nil {
				return diag.Errorf("error reading dh_group_27: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_28"); ok {
		if o.DhGroup28 != nil {
			if err = d.Set("dh_group_28", flattenSystemIkeDhGroup28(d, o.DhGroup28, "dh_group_28", sort)); err != nil {
				return diag.Errorf("error reading dh_group_28: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_29"); ok {
		if o.DhGroup29 != nil {
			if err = d.Set("dh_group_29", flattenSystemIkeDhGroup29(d, o.DhGroup29, "dh_group_29", sort)); err != nil {
				return diag.Errorf("error reading dh_group_29: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_30"); ok {
		if o.DhGroup30 != nil {
			if err = d.Set("dh_group_30", flattenSystemIkeDhGroup30(d, o.DhGroup30, "dh_group_30", sort)); err != nil {
				return diag.Errorf("error reading dh_group_30: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_31"); ok {
		if o.DhGroup31 != nil {
			if err = d.Set("dh_group_31", flattenSystemIkeDhGroup31(d, o.DhGroup31, "dh_group_31", sort)); err != nil {
				return diag.Errorf("error reading dh_group_31: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_32"); ok {
		if o.DhGroup32 != nil {
			if err = d.Set("dh_group_32", flattenSystemIkeDhGroup32(d, o.DhGroup32, "dh_group_32", sort)); err != nil {
				return diag.Errorf("error reading dh_group_32: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("dh_group_5"); ok {
		if o.DhGroup5 != nil {
			if err = d.Set("dh_group_5", flattenSystemIkeDhGroup5(d, o.DhGroup5, "dh_group_5", sort)); err != nil {
				return diag.Errorf("error reading dh_group_5: %v", err)
			}
		}
	}

	if o.DhKeypairCache != nil {
		v := *o.DhKeypairCache

		if err = d.Set("dh_keypair_cache", v); err != nil {
			return diag.Errorf("error reading dh_keypair_cache: %v", err)
		}
	}

	if o.DhKeypairCount != nil {
		v := *o.DhKeypairCount

		if err = d.Set("dh_keypair_count", v); err != nil {
			return diag.Errorf("error reading dh_keypair_count: %v", err)
		}
	}

	if o.DhKeypairThrottle != nil {
		v := *o.DhKeypairThrottle

		if err = d.Set("dh_keypair_throttle", v); err != nil {
			return diag.Errorf("error reading dh_keypair_throttle: %v", err)
		}
	}

	if o.DhMode != nil {
		v := *o.DhMode

		if err = d.Set("dh_mode", v); err != nil {
			return diag.Errorf("error reading dh_mode: %v", err)
		}
	}

	if o.DhMultiprocess != nil {
		v := *o.DhMultiprocess

		if err = d.Set("dh_multiprocess", v); err != nil {
			return diag.Errorf("error reading dh_multiprocess: %v", err)
		}
	}

	if o.DhWorkerCount != nil {
		v := *o.DhWorkerCount

		if err = d.Set("dh_worker_count", v); err != nil {
			return diag.Errorf("error reading dh_worker_count: %v", err)
		}
	}

	if o.EmbryonicLimit != nil {
		v := *o.EmbryonicLimit

		if err = d.Set("embryonic_limit", v); err != nil {
			return diag.Errorf("error reading embryonic_limit: %v", err)
		}
	}

	return nil
}

func expandSystemIkeDhGroup1(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup1, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup1

	for i := range l {
		tmp := models.SystemIkeDhGroup1{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup14(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup14, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup14

	for i := range l {
		tmp := models.SystemIkeDhGroup14{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup15(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup15, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup15

	for i := range l {
		tmp := models.SystemIkeDhGroup15{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup16(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup16, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup16

	for i := range l {
		tmp := models.SystemIkeDhGroup16{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup17(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup17, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup17

	for i := range l {
		tmp := models.SystemIkeDhGroup17{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup18(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup18, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup18

	for i := range l {
		tmp := models.SystemIkeDhGroup18{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup19(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup19, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup19

	for i := range l {
		tmp := models.SystemIkeDhGroup19{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup2(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup2, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup2

	for i := range l {
		tmp := models.SystemIkeDhGroup2{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup20(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup20, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup20

	for i := range l {
		tmp := models.SystemIkeDhGroup20{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup21(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup21, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup21

	for i := range l {
		tmp := models.SystemIkeDhGroup21{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup27(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup27, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup27

	for i := range l {
		tmp := models.SystemIkeDhGroup27{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup28(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup28, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup28

	for i := range l {
		tmp := models.SystemIkeDhGroup28{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup29(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup29, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup29

	for i := range l {
		tmp := models.SystemIkeDhGroup29{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup30(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup30, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup30

	for i := range l {
		tmp := models.SystemIkeDhGroup30{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup31(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup31, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup31

	for i := range l {
		tmp := models.SystemIkeDhGroup31{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup32(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup32, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup32

	for i := range l {
		tmp := models.SystemIkeDhGroup32{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemIkeDhGroup5(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemIkeDhGroup5, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIkeDhGroup5

	for i := range l {
		tmp := models.SystemIkeDhGroup5{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keypair_cache", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeypairCache = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keypair_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeypairCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func getObjectSystemIke(d *schema.ResourceData, sv string) (*models.SystemIke, diag.Diagnostics) {
	obj := models.SystemIke{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("dh_group_1"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_1", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup1(d, v, "dh_group_1", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup1 = t
		}
	} else if d.HasChange("dh_group_1") {
		old, new := d.GetChange("dh_group_1")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup1 = &models.SystemIkeDhGroup1{}
		}
	}
	if v, ok := d.GetOk("dh_group_14"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_14", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup14(d, v, "dh_group_14", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup14 = t
		}
	} else if d.HasChange("dh_group_14") {
		old, new := d.GetChange("dh_group_14")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup14 = &models.SystemIkeDhGroup14{}
		}
	}
	if v, ok := d.GetOk("dh_group_15"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_15", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup15(d, v, "dh_group_15", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup15 = t
		}
	} else if d.HasChange("dh_group_15") {
		old, new := d.GetChange("dh_group_15")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup15 = &models.SystemIkeDhGroup15{}
		}
	}
	if v, ok := d.GetOk("dh_group_16"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_16", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup16(d, v, "dh_group_16", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup16 = t
		}
	} else if d.HasChange("dh_group_16") {
		old, new := d.GetChange("dh_group_16")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup16 = &models.SystemIkeDhGroup16{}
		}
	}
	if v, ok := d.GetOk("dh_group_17"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_17", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup17(d, v, "dh_group_17", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup17 = t
		}
	} else if d.HasChange("dh_group_17") {
		old, new := d.GetChange("dh_group_17")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup17 = &models.SystemIkeDhGroup17{}
		}
	}
	if v, ok := d.GetOk("dh_group_18"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_18", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup18(d, v, "dh_group_18", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup18 = t
		}
	} else if d.HasChange("dh_group_18") {
		old, new := d.GetChange("dh_group_18")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup18 = &models.SystemIkeDhGroup18{}
		}
	}
	if v, ok := d.GetOk("dh_group_19"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_19", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup19(d, v, "dh_group_19", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup19 = t
		}
	} else if d.HasChange("dh_group_19") {
		old, new := d.GetChange("dh_group_19")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup19 = &models.SystemIkeDhGroup19{}
		}
	}
	if v, ok := d.GetOk("dh_group_2"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_2", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup2(d, v, "dh_group_2", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup2 = t
		}
	} else if d.HasChange("dh_group_2") {
		old, new := d.GetChange("dh_group_2")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup2 = &models.SystemIkeDhGroup2{}
		}
	}
	if v, ok := d.GetOk("dh_group_20"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_20", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup20(d, v, "dh_group_20", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup20 = t
		}
	} else if d.HasChange("dh_group_20") {
		old, new := d.GetChange("dh_group_20")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup20 = &models.SystemIkeDhGroup20{}
		}
	}
	if v, ok := d.GetOk("dh_group_21"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_21", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup21(d, v, "dh_group_21", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup21 = t
		}
	} else if d.HasChange("dh_group_21") {
		old, new := d.GetChange("dh_group_21")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup21 = &models.SystemIkeDhGroup21{}
		}
	}
	if v, ok := d.GetOk("dh_group_27"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_27", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup27(d, v, "dh_group_27", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup27 = t
		}
	} else if d.HasChange("dh_group_27") {
		old, new := d.GetChange("dh_group_27")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup27 = &models.SystemIkeDhGroup27{}
		}
	}
	if v, ok := d.GetOk("dh_group_28"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_28", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup28(d, v, "dh_group_28", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup28 = t
		}
	} else if d.HasChange("dh_group_28") {
		old, new := d.GetChange("dh_group_28")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup28 = &models.SystemIkeDhGroup28{}
		}
	}
	if v, ok := d.GetOk("dh_group_29"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_29", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup29(d, v, "dh_group_29", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup29 = t
		}
	} else if d.HasChange("dh_group_29") {
		old, new := d.GetChange("dh_group_29")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup29 = &models.SystemIkeDhGroup29{}
		}
	}
	if v, ok := d.GetOk("dh_group_30"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_30", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup30(d, v, "dh_group_30", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup30 = t
		}
	} else if d.HasChange("dh_group_30") {
		old, new := d.GetChange("dh_group_30")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup30 = &models.SystemIkeDhGroup30{}
		}
	}
	if v, ok := d.GetOk("dh_group_31"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_31", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup31(d, v, "dh_group_31", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup31 = t
		}
	} else if d.HasChange("dh_group_31") {
		old, new := d.GetChange("dh_group_31")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup31 = &models.SystemIkeDhGroup31{}
		}
	}
	if v, ok := d.GetOk("dh_group_32"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_32", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup32(d, v, "dh_group_32", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup32 = t
		}
	} else if d.HasChange("dh_group_32") {
		old, new := d.GetChange("dh_group_32")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup32 = &models.SystemIkeDhGroup32{}
		}
	}
	if v, ok := d.GetOk("dh_group_5"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dh_group_5", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIkeDhGroup5(d, v, "dh_group_5", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhGroup5 = t
		}
	} else if d.HasChange("dh_group_5") {
		old, new := d.GetChange("dh_group_5")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhGroup5 = &models.SystemIkeDhGroup5{}
		}
	}
	if v1, ok := d.GetOk("dh_keypair_cache"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dh_keypair_cache", sv)
				diags = append(diags, e)
			}
			obj.DhKeypairCache = &v2
		}
	}
	if v1, ok := d.GetOk("dh_keypair_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dh_keypair_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DhKeypairCount = &tmp
		}
	}
	if v1, ok := d.GetOk("dh_keypair_throttle"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dh_keypair_throttle", sv)
				diags = append(diags, e)
			}
			obj.DhKeypairThrottle = &v2
		}
	}
	if v1, ok := d.GetOk("dh_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dh_mode", sv)
				diags = append(diags, e)
			}
			obj.DhMode = &v2
		}
	}
	if v1, ok := d.GetOk("dh_multiprocess"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dh_multiprocess", sv)
				diags = append(diags, e)
			}
			obj.DhMultiprocess = &v2
		}
	}
	if v1, ok := d.GetOk("dh_worker_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dh_worker_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DhWorkerCount = &tmp
		}
	}
	if v1, ok := d.GetOk("embryonic_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("embryonic_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EmbryonicLimit = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemIke(d *schema.ResourceData, sv string) (*models.SystemIke, diag.Diagnostics) {
	obj := models.SystemIke{}
	diags := diag.Diagnostics{}

	obj.DhGroup1 = &models.SystemIkeDhGroup1{}
	obj.DhGroup14 = &models.SystemIkeDhGroup14{}
	obj.DhGroup15 = &models.SystemIkeDhGroup15{}
	obj.DhGroup16 = &models.SystemIkeDhGroup16{}
	obj.DhGroup17 = &models.SystemIkeDhGroup17{}
	obj.DhGroup18 = &models.SystemIkeDhGroup18{}
	obj.DhGroup19 = &models.SystemIkeDhGroup19{}
	obj.DhGroup2 = &models.SystemIkeDhGroup2{}
	obj.DhGroup20 = &models.SystemIkeDhGroup20{}
	obj.DhGroup21 = &models.SystemIkeDhGroup21{}
	obj.DhGroup27 = &models.SystemIkeDhGroup27{}
	obj.DhGroup28 = &models.SystemIkeDhGroup28{}
	obj.DhGroup29 = &models.SystemIkeDhGroup29{}
	obj.DhGroup30 = &models.SystemIkeDhGroup30{}
	obj.DhGroup31 = &models.SystemIkeDhGroup31{}
	obj.DhGroup32 = &models.SystemIkeDhGroup32{}
	obj.DhGroup5 = &models.SystemIkeDhGroup5{}

	return &obj, diags
}
