// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IKE global attributes.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemIke() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IKE global attributes.",

		ReadContext: dataSourceSystemIkeRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"dh_group_1": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 1 (MODP-768).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_14": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 14 (MODP-2048).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_15": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 15 (MODP-3072).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_16": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 16 (MODP-4096).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_17": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 17 (MODP-6144).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_18": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 18 (MODP-8192).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_19": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 19 (EC-P256).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_2": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 2 (MODP-1024).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_20": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 20 (EC-P384).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_21": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 21 (EC-P521).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_27": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 27 (EC-P224BP).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_28": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 28 (EC-P256BP).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_29": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 29 (EC-P384BP).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_30": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 30 (EC-P512BP).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_31": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 31 (EC-X25519).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_32": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 32 (EC-X448).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_group_5": {
				Type:        schema.TypeList,
				Description: "Diffie-Hellman group 5 (MODP-1536).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_cache": {
							Type:        schema.TypeString,
							Description: "Configure custom key pair cache size for this Diffie-Hellman group.",
							Computed:    true,
						},
						"keypair_count": {
							Type:        schema.TypeInt,
							Description: "Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group.",
							Computed:    true,
						},
					},
				},
			},
			"dh_keypair_cache": {
				Type:        schema.TypeString,
				Description: "Enable/disable Diffie-Hellman key pair cache.",
				Computed:    true,
			},
			"dh_keypair_count": {
				Type:        schema.TypeInt,
				Description: "Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).",
				Computed:    true,
			},
			"dh_keypair_throttle": {
				Type:        schema.TypeString,
				Description: "Enable/disable Diffie-Hellman key pair cache CPU throttling.",
				Computed:    true,
			},
			"dh_mode": {
				Type:        schema.TypeString,
				Description: "Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations.",
				Computed:    true,
			},
			"dh_multiprocess": {
				Type:        schema.TypeString,
				Description: "Enable/disable multiprocess Diffie-Hellman daemon for IKE.",
				Computed:    true,
			},
			"dh_worker_count": {
				Type:        schema.TypeInt,
				Description: "Number of Diffie-Hellman workers to start.",
				Computed:    true,
			},
			"embryonic_limit": {
				Type:        schema.TypeInt,
				Description: "Maximum number of IPsec tunnels to negotiate simultaneously.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemIkeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemIke"

	o, err := c.Cmdb.ReadSystemIke(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemIke dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
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

	d.SetId(mkey)

	return nil
}
