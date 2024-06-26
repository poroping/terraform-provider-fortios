package forticlient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type JSONSystemVpnCertificateLocalImport struct {
	Name        string `json:"certname"`
	Type        string `json:"type"`
	Certificate string `json:"file_content"`
	Password    string `json:"password,omitempty"`
	PrivateKey  string `json:"key_file_content,omitempty"`
	Scope       string `json:"scope"`
}

// Uploads certificate to Fortigate
func (c *FortiSDKClient) CreateSystemVpnCertificateLocalImport(data *JSONSystemVpnCertificateLocalImport, vdomparam string, batch int) (res string, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/monitor/vpn-certificate/local/import"
	data.Certificate = base64.StdEncoding.EncodeToString([]byte(data.Certificate))
	if data.PrivateKey != "" {
		data.PrivateKey = base64.StdEncoding.EncodeToString([]byte(data.PrivateKey))
	}
	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes, batch)
	err = req.Send3(vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	if result["status"] == "error" {
		err = fortiAPIErrorFormat(result, string(body))
		if err != nil {
			err = fmt.Errorf("error importing certificate: %s", err)
			return
		}
	}

	res = string(body)

	return
}
