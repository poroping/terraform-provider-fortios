// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.0 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Configure Access Proxy.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceFirewallAccessProxy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallAccessProxyRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"api_gateway": {
				Type:        schema.TypeList,
				Description: "Set API Gateway.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_cookie_age": {
							Type:        schema.TypeInt,
							Description: "Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.",
							Computed:    true,
						},
						"http_cookie_domain": {
							Type:        schema.TypeString,
							Description: "Domain that HTTP cookie persistence should apply to.",
							Computed:    true,
						},
						"http_cookie_domain_from_host": {
							Type:        schema.TypeString,
							Description: "Enable/disable use of HTTP cookie domain from host field in HTTP.",
							Computed:    true,
						},
						"http_cookie_generation": {
							Type:        schema.TypeInt,
							Description: "Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.",
							Computed:    true,
						},
						"http_cookie_path": {
							Type:        schema.TypeString,
							Description: "Limit HTTP cookie persistence to the specified path.",
							Computed:    true,
						},
						"http_cookie_share": {
							Type:        schema.TypeString,
							Description: "Control sharing of cookies across API Gateway. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.",
							Computed:    true,
						},
						"https_cookie_secure": {
							Type:        schema.TypeString,
							Description: "Enable/disable verification that inserted HTTPS cookies are secure.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "API Gateway ID.",
							Computed:    true,
						},
						"ldb_method": {
							Type:        schema.TypeString,
							Description: "Method used to distribute sessions to real servers.",
							Computed:    true,
						},
						"persistence": {
							Type:        schema.TypeString,
							Description: "Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.",
							Computed:    true,
						},
						"realservers": {
							Type:        schema.TypeList,
							Description: "Select the real servers that this Access Proxy will distribute traffic to.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Description: "Address or address group of the real server.",
										Computed:    true,
									},
									"health_check": {
										Type:        schema.TypeString,
										Description: "Enable to check the responsiveness of the real server before forwarding traffic.",
										Computed:    true,
									},
									"health_check_proto": {
										Type:        schema.TypeString,
										Description: "Protocol of the health check monitor to use when polling to determine server's connectivity status.",
										Computed:    true,
									},
									"http_host": {
										Type:        schema.TypeString,
										Description: "HTTP server domain name in HTTP header.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Real server ID.",
										Computed:    true,
									},
									"ip": {
										Type:        schema.TypeString,
										Description: "IP address of the real server.",
										Computed:    true,
									},
									"mappedport": {
										Type:        schema.TypeString,
										Description: "Port for communicating with the real server.",
										Computed:    true,
									},
									"port": {
										Type:        schema.TypeInt,
										Description: "Port for communicating with the real server.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.",
										Computed:    true,
									},
									"weight": {
										Type:        schema.TypeInt,
										Description: "Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.",
										Computed:    true,
									},
								},
							},
						},
						"saml_server": {
							Type:        schema.TypeString,
							Description: "SAML service provider configuration for VIP authentication.",
							Computed:    true,
						},
						"service": {
							Type:        schema.TypeString,
							Description: "Service.",
							Computed:    true,
						},
						"ssl_algorithm": {
							Type:        schema.TypeString,
							Description: "Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength.",
							Computed:    true,
						},
						"ssl_cipher_suites": {
							Type:        schema.TypeList,
							Description: "SSL/TLS cipher suites to offer to a server, ordered by priority.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cipher": {
										Type:        schema.TypeString,
										Description: "Cipher suite name.",
										Computed:    true,
									},
									"priority": {
										Type:        schema.TypeInt,
										Description: "SSL/TLS cipher suites priority.",
										Computed:    true,
									},
									"versions": {
										Type:        schema.TypeString,
										Description: "SSL/TLS versions that the cipher suite can be used with.",
										Computed:    true,
									},
								},
							},
						},
						"ssl_dh_bits": {
							Type:        schema.TypeString,
							Description: "Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions.",
							Computed:    true,
						},
						"ssl_max_version": {
							Type:        schema.TypeString,
							Description: "Highest SSL/TLS version acceptable from a server.",
							Computed:    true,
						},
						"ssl_min_version": {
							Type:        schema.TypeString,
							Description: "Lowest SSL/TLS version acceptable from a server.",
							Computed:    true,
						},
						"url_map": {
							Type:        schema.TypeString,
							Description: "URL pattern to match.",
							Computed:    true,
						},
						"url_map_type": {
							Type:        schema.TypeString,
							Description: "Type of url-map.",
							Computed:    true,
						},
						"virtual_host": {
							Type:        schema.TypeString,
							Description: "Virtual host.",
							Computed:    true,
						},
					},
				},
			},
			"client_cert": {
				Type:        schema.TypeString,
				Description: "Enable/disable to request client certificate.",
				Computed:    true,
			},
			"empty_cert_action": {
				Type:        schema.TypeString,
				Description: "Action of an empty client certificate.",
				Computed:    true,
			},
			"ldb_method": {
				Type:        schema.TypeString,
				Description: "Method used to distribute sessions to SSL real servers.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Access Proxy name.",
				Required:    true,
			},
			"realservers": {
				Type:        schema.TypeList,
				Description: "Select the SSL real servers that this Access Proxy will distribute traffic to.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Real server ID.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "IP address of the real server.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Port for communicating with the real server.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.",
							Computed:    true,
						},
						"weight": {
							Type:        schema.TypeInt,
							Description: "Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.",
							Computed:    true,
						},
					},
				},
			},
			"server_pubkey_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSH real server public key authentication.",
				Computed:    true,
			},
			"server_pubkey_auth_settings": {
				Type:        schema.TypeList,
				Description: "Server SSH public key authentication settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_ca": {
							Type:        schema.TypeString,
							Description: "Name of the SSH server public key authentication CA.",
							Computed:    true,
						},
						"cert_extension": {
							Type:        schema.TypeList,
							Description: "Configure certificate extension for user certificate.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"critical": {
										Type:        schema.TypeString,
										Description: "Critical option.",
										Computed:    true,
									},
									"data": {
										Type:        schema.TypeString,
										Description: "Name of certificate extension.",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Name of certificate extension.",
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "Type of certificate extension.",
										Computed:    true,
									},
								},
							},
						},
						"permit_agent_forwarding": {
							Type:        schema.TypeString,
							Description: "Enable/disable appending permit-agent-forwarding certificate extension.",
							Computed:    true,
						},
						"permit_port_forwarding": {
							Type:        schema.TypeString,
							Description: "Enable/disable appending permit-port-forwarding certificate extension.",
							Computed:    true,
						},
						"permit_pty": {
							Type:        schema.TypeString,
							Description: "Enable/disable appending permit-pty certificate extension.",
							Computed:    true,
						},
						"permit_user_rc": {
							Type:        schema.TypeString,
							Description: "Enable/disable appending permit-user-rc certificate extension.",
							Computed:    true,
						},
						"permit_x11_forwarding": {
							Type:        schema.TypeString,
							Description: "Enable/disable appending permit-x11-forwarding certificate extension.",
							Computed:    true,
						},
						"source_address": {
							Type:        schema.TypeString,
							Description: "Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address.",
							Computed:    true,
						},
					},
				},
			},
			"vip": {
				Type:        schema.TypeString,
				Description: "Virtual IP name.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallAccessProxyRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""
	key := "name"
	t := d.Get(key)
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing FirewallAccessProxy: type error")
	}

	o, err := c.ReadFirewallAccessProxy(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing FirewallAccessProxy: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallAccessProxy(d, o)
	if err != nil {
		return fmt.Errorf("error describing FirewallAccessProxy from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallAccessProxyApiGateway(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {

			tmp["http_cookie_age"] = dataSourceFlattenFirewallAccessProxyApiGatewayHttpCookieAge(i["http-cookie-age"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {

			tmp["http_cookie_domain"] = dataSourceFlattenFirewallAccessProxyApiGatewayHttpCookieDomain(i["http-cookie-domain"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {

			tmp["http_cookie_domain_from_host"] = dataSourceFlattenFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {

			tmp["http_cookie_generation"] = dataSourceFlattenFirewallAccessProxyApiGatewayHttpCookieGeneration(i["http-cookie-generation"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {

			tmp["http_cookie_path"] = dataSourceFlattenFirewallAccessProxyApiGatewayHttpCookiePath(i["http-cookie-path"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {

			tmp["http_cookie_share"] = dataSourceFlattenFirewallAccessProxyApiGatewayHttpCookieShare(i["http-cookie-share"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {

			tmp["https_cookie_secure"] = dataSourceFlattenFirewallAccessProxyApiGatewayHttpsCookieSecure(i["https-cookie-secure"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenFirewallAccessProxyApiGatewayId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {

			tmp["ldb_method"] = dataSourceFlattenFirewallAccessProxyApiGatewayLdbMethod(i["ldb-method"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {

			tmp["persistence"] = dataSourceFlattenFirewallAccessProxyApiGatewayPersistence(i["persistence"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {

			tmp["realservers"] = dataSourceFlattenFirewallAccessProxyApiGatewayRealservers(i["realservers"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := i["saml-server"]; ok {

			tmp["saml_server"] = dataSourceFlattenFirewallAccessProxyApiGatewaySamlServer(i["saml-server"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {

			tmp["service"] = dataSourceFlattenFirewallAccessProxyApiGatewayService(i["service"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {

			tmp["ssl_algorithm"] = dataSourceFlattenFirewallAccessProxyApiGatewaySslAlgorithm(i["ssl-algorithm"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {

			tmp["ssl_cipher_suites"] = dataSourceFlattenFirewallAccessProxyApiGatewaySslCipherSuites(i["ssl-cipher-suites"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {

			tmp["ssl_dh_bits"] = dataSourceFlattenFirewallAccessProxyApiGatewaySslDhBits(i["ssl-dh-bits"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {

			tmp["ssl_max_version"] = dataSourceFlattenFirewallAccessProxyApiGatewaySslMaxVersion(i["ssl-max-version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {

			tmp["ssl_min_version"] = dataSourceFlattenFirewallAccessProxyApiGatewaySslMinVersion(i["ssl-min-version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := i["url-map"]; ok {

			tmp["url_map"] = dataSourceFlattenFirewallAccessProxyApiGatewayUrlMap(i["url-map"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := i["url-map-type"]; ok {

			tmp["url_map_type"] = dataSourceFlattenFirewallAccessProxyApiGatewayUrlMapType(i["url-map-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := i["virtual-host"]; ok {

			tmp["virtual_host"] = dataSourceFlattenFirewallAccessProxyApiGatewayVirtualHost(i["virtual-host"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}
func dataSourceFlattenFirewallAccessProxyApiGatewayHttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayHttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayHttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayPersistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayRealservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {

			tmp["address"] = dataSourceFlattenFirewallAccessProxyApiGatewayRealserversAddress(i["address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {

			tmp["health_check"] = dataSourceFlattenFirewallAccessProxyApiGatewayRealserversHealthCheck(i["health-check"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {

			tmp["health_check_proto"] = dataSourceFlattenFirewallAccessProxyApiGatewayRealserversHealthCheckProto(i["health-check-proto"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {

			tmp["http_host"] = dataSourceFlattenFirewallAccessProxyApiGatewayRealserversHttpHost(i["http-host"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenFirewallAccessProxyApiGatewayRealserversId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = dataSourceFlattenFirewallAccessProxyApiGatewayRealserversIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {

			tmp["mappedport"] = dataSourceFlattenFirewallAccessProxyApiGatewayRealserversMappedport(i["mappedport"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = dataSourceFlattenFirewallAccessProxyApiGatewayRealserversPort(i["port"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = dataSourceFlattenFirewallAccessProxyApiGatewayRealserversStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = dataSourceFlattenFirewallAccessProxyApiGatewayRealserversWeight(i["weight"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}
func dataSourceFlattenFirewallAccessProxyApiGatewayRealserversAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayRealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayRealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayRealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayRealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayRealserversMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayRealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayRealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayRealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewaySamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewaySslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewaySslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {

			tmp["cipher"] = dataSourceFlattenFirewallAccessProxyApiGatewaySslCipherSuitesCipher(i["cipher"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {

			tmp["priority"] = dataSourceFlattenFirewallAccessProxyApiGatewaySslCipherSuitesPriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {

			tmp["versions"] = dataSourceFlattenFirewallAccessProxyApiGatewaySslCipherSuitesVersions(i["versions"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}
func dataSourceFlattenFirewallAccessProxyApiGatewaySslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewaySslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewaySslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewaySslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewaySslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewaySslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayUrlMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayUrlMapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyApiGatewayVirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyEmptyCertAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyRealservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenFirewallAccessProxyRealserversId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = dataSourceFlattenFirewallAccessProxyRealserversIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = dataSourceFlattenFirewallAccessProxyRealserversPort(i["port"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = dataSourceFlattenFirewallAccessProxyRealserversStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = dataSourceFlattenFirewallAccessProxyRealserversWeight(i["weight"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}
func dataSourceFlattenFirewallAccessProxyRealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyRealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyRealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyRealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyRealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyServerPubkeyAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_ca"
		if _, ok := i["auth-ca"]; ok {

			tmp["auth_ca"] = dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsAuthCa(i["auth-ca"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cert_extension"
		if _, ok := i["cert-extension"]; ok {

			tmp["cert_extension"] = dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtension(i["cert-extension"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_agent_forwarding"
		if _, ok := i["permit-agent-forwarding"]; ok {

			tmp["permit_agent_forwarding"] = dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding(i["permit-agent-forwarding"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_port_forwarding"
		if _, ok := i["permit-port-forwarding"]; ok {

			tmp["permit_port_forwarding"] = dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding(i["permit-port-forwarding"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_pty"
		if _, ok := i["permit-pty"]; ok {

			tmp["permit_pty"] = dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsPermitPty(i["permit-pty"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_user_rc"
		if _, ok := i["permit-user-rc"]; ok {

			tmp["permit_user_rc"] = dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc(i["permit-user-rc"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_x11_forwarding"
		if _, ok := i["permit-x11-forwarding"]; ok {

			tmp["permit_x11_forwarding"] = dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding(i["permit-x11-forwarding"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address"
		if _, ok := i["source-address"]; ok {

			tmp["source_address"] = dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress(i["source-address"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}
func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsAuthCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtension(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "critical"
		if _, ok := i["critical"]; ok {

			tmp["critical"] = dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical(i["critical"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := i["data"]; ok {

			tmp["data"] = dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData(i["data"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType(i["type"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}
func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsPermitPty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAccessProxyVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallAccessProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("api_gateway", dataSourceFlattenFirewallAccessProxyApiGateway(o["api-gateway"], d, "api_gateway")); err != nil {
		if !fortiAPIPatch(o["api-gateway"]) {
			return fmt.Errorf("error reading api_gateway: %v", err)
		}
	}
	if err = d.Set("client_cert", dataSourceFlattenFirewallAccessProxyClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if !fortiAPIPatch(o["client-cert"]) {
			return fmt.Errorf("error reading client_cert: %v", err)
		}
	}
	if err = d.Set("empty_cert_action", dataSourceFlattenFirewallAccessProxyEmptyCertAction(o["empty-cert-action"], d, "empty_cert_action")); err != nil {
		if !fortiAPIPatch(o["empty-cert-action"]) {
			return fmt.Errorf("error reading empty_cert_action: %v", err)
		}
	}
	if err = d.Set("ldb_method", dataSourceFlattenFirewallAccessProxyLdbMethod(o["ldb-method"], d, "ldb_method")); err != nil {
		if !fortiAPIPatch(o["ldb-method"]) {
			return fmt.Errorf("error reading ldb_method: %v", err)
		}
	}
	if err = d.Set("name", dataSourceFlattenFirewallAccessProxyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}
	if err = d.Set("realservers", dataSourceFlattenFirewallAccessProxyRealservers(o["realservers"], d, "realservers")); err != nil {
		if !fortiAPIPatch(o["realservers"]) {
			return fmt.Errorf("error reading realservers: %v", err)
		}
	}
	if err = d.Set("server_pubkey_auth", dataSourceFlattenFirewallAccessProxyServerPubkeyAuth(o["server-pubkey-auth"], d, "server_pubkey_auth")); err != nil {
		if !fortiAPIPatch(o["server-pubkey-auth"]) {
			return fmt.Errorf("error reading server_pubkey_auth: %v", err)
		}
	}
	if err = d.Set("server_pubkey_auth_settings", dataSourceFlattenFirewallAccessProxyServerPubkeyAuthSettings(o["server-pubkey-auth-settings"], d, "server_pubkey_auth_settings")); err != nil {
		if !fortiAPIPatch(o["server-pubkey-auth-settings"]) {
			return fmt.Errorf("error reading server_pubkey_auth_settings: %v", err)
		}
	}
	if err = d.Set("vip", dataSourceFlattenFirewallAccessProxyVip(o["vip"], d, "vip")); err != nil {
		if !fortiAPIPatch(o["vip"]) {
			return fmt.Errorf("error reading vip: %v", err)
		}
	}

	return nil
}
