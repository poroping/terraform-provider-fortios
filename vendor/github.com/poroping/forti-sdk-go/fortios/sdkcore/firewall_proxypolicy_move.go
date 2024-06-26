package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// CreateUpdateFirewallProxypolicyMove API operation for FortiOS moves the specified item.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateFirewallProxypolicyMove(srcId, dstId, mv, vdomparam string, batch int) (err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/proxy-policy"
	path += "/" + srcId

	params := make(map[string][]string)
	params["action"] = []string{"move"}
	params[mv] = []string{dstId}

	req := c.NewRequest(HTTPMethod, path, &params, nil, batch)
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
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

// JSONFirewallProxypolicyItem contains the necessary parameters for each item
type JSONFirewallProxypolicyItem struct {
	Policyid string `json:"policyid"`
}

// GetFirewallProxypolicyList API operation for FortiOS gets the list
// Returns the requested API user value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) GetFirewallProxypolicyList(vdomparam string, batch int) (out []JSONFirewallProxypolicyItem, err error) {

	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/proxy-policy/"

	params := make(map[string][]string)
	params["format"] = []string{"policyid"}

	req := c.NewRequest(HTTPMethod, path, &params, nil, batch)
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
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if fortiAPIHttpStatus404Checking(result) {
		return
	}

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		mapTmp := result["results"].([]interface{}) //)[0].(map[string]interface{})

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		var members []JSONFirewallProxypolicyItem
		for _, v := range mapTmp {
			c := v.(map[string]interface{})

			members = append(members,
				JSONFirewallProxypolicyItem{
					Policyid: strconv.Itoa(int(c["policyid"].(float64))),
				})
		}

		out = members
	}

	return
}
