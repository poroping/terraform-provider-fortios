package fortios

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	fmgclient "github.com/fortinetdev/forti-sdk-go/fortimanager/sdkcore"
	"github.com/poroping/forti-sdk-go/fortios/auth"
	forticlient "github.com/poroping/forti-sdk-go/fortios/sdkcore"
)

// Config gets the authentication information from the given metadata
type Config struct {
	Hostname string
	Token    string
	Insecure *bool
	CABundle string
	Vdom     string

	PeerAuth   string
	CaCert     string
	ClientCert string
	ClientKey  string
}

// FortiClient contains the basic FortiOS SDK connection information to FortiOS
// It can be used to as a client of FortiOS for the plugin
// Now FortiClient contains two kinds of clients:
// Client is for FortiGate
// Client Fottimanager is for FortiManager
type FortiClient struct {
	//to sdk client
	Client             *forticlient.FortiSDKClient
	ClientFortimanager *fmgclient.FmgSDKClient
}

// CreateClient creates a FortiClient Object with the authentication information.
// It returns the FortiClient Object for the use when the plugin is initialized.
func (c *Config) CreateClient() (interface{}, error) {
	var fClient FortiClient

	bFOSExist := bFortiOSHostnameExist(c)

	if bFOSExist {
		err := createFortiOSClient(&fClient, c)
		if err != nil {
			return nil, fmt.Errorf("error create fortios client: %v", err)
		}
	}

	if !bFOSExist {
		return nil, fmt.Errorf("FortiOS hostname should be set")
	}

	return &fClient, nil
}

func bFortiOSHostnameExist(c *Config) bool {
	if c.Hostname == "" {
		if os.Getenv("FORTIOS_ACCESS_HOSTNAME") == "" {
			return false
		}
	}

	return true
}

func createFortiOSClient(fClient *FortiClient, c *Config) error {
	config := &tls.Config{}

	auth := auth.NewAuth(c.Hostname, c.Token, c.CABundle, c.PeerAuth, c.CaCert, c.ClientCert, c.ClientKey, c.Vdom)

	if auth.Hostname == "" {
		_, err := auth.GetEnvHostname()
		if err != nil {
			return fmt.Errorf("error reading Hostname")
		}
	}

	if auth.Token == "" {
		_, err := auth.GetEnvToken()
		if err != nil {
			return fmt.Errorf("error reading Token")
		}
	}

	if auth.CABundle == "" {
		auth.GetEnvCABundle()
	}

	if auth.PeerAuth == "" {
		_, err := auth.GetEnvPeerAuth()
		if err != nil {
			return fmt.Errorf("error reading PeerAuth")
		}
	}
	if auth.CaCert == "" {
		_, err := auth.GetEnvCaCert()
		if err != nil {
			return fmt.Errorf("error reading CaCert")
		}
	}
	if auth.ClientCert == "" {
		_, err := auth.GetEnvClientCert()
		if err != nil {
			return fmt.Errorf("error reading ClientCert")
		}
	}
	if auth.ClientKey == "" {
		_, err := auth.GetEnvClientKey()
		if err != nil {
			return fmt.Errorf("error reading ClientKey")
		}
	}

	pool := x509.NewCertPool()

	if auth.CABundle != "" {
		f, err := os.Open(auth.CABundle)
		if err != nil {
			return fmt.Errorf("error reading CA Bundle: %v", err)
		}
		defer f.Close()

		caBundle, err := ioutil.ReadAll(f)
		if err != nil {
			return fmt.Errorf("error reading CA Bundle: %v", err)
		}

		if !pool.AppendCertsFromPEM([]byte(caBundle)) {
			return fmt.Errorf("error reading CA Bundle")
		}
		config.RootCAs = pool
	}

	if auth.PeerAuth == "enable" {
		if auth.CaCert != "" {
			caCertFile := auth.CaCert
			caCert, err := ioutil.ReadFile(caCertFile)
			if err != nil {
				return fmt.Errorf("client ioutil.ReadFile couldn't load cacert file: %v", err)
			}

			pool.AppendCertsFromPEM(caCert)
		}

		if auth.ClientCert == "" {
			return fmt.Errorf("user cert file doesn't exist")
		}

		if auth.ClientKey == "" {
			return fmt.Errorf("user key file doesn't exist")
		}

		clientCert, err := tls.LoadX509KeyPair(auth.ClientCert, auth.ClientKey)
		if err != nil {
			return fmt.Errorf("client ioutil.ReadFile couldn't load clientCert/clientKey file: %v", err)
		}

		config.Certificates = []tls.Certificate{clientCert}
	}

	if c.Insecure == nil {
		// defaut value for Insecure is false
		b, _ := auth.GetEnvInsecure()
		config.InsecureSkipVerify = b
	} else {
		config.InsecureSkipVerify = *c.Insecure
	}

	if !config.InsecureSkipVerify && auth.CABundle == "" {
		return fmt.Errorf("error getting CA Bundle, CA Bundle should be set when insecure is false")
	}

	tr := &http.Transport{
		TLSClientConfig: config,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 250,
	}

	fc, err := forticlient.NewClient(auth, client)

	if err != nil {
		return fmt.Errorf("connection error: %v", err)
	}

	fClient.Client = fc

	return nil
}
