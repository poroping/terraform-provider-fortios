package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
)

type policySortFirewallProxypolicy struct {
	policyid int
}

func getPolicyListFirewallProxypolicy(c *FortiSDKClient, vdomparam string, batch int) (idlist []policySortFirewallProxypolicy, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/proxy-policy"

	params := make(map[string][]string)
	params["format"] = []string{"policy|name"}

	req := c.NewRequest(HTTPMethod, path, &params, nil, batch)
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
		mapTmp := result["results"].([]interface{})

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		for _, v := range mapTmp {
			c := v.(map[string]interface{})

			idlist = append(idlist, policySortFirewallProxypolicy{policyid: int(c["policyid"].(float64))})
		}
	}

	return
}

func bPolicyListSortedFirewallProxypolicy(idlist []policySortFirewallProxypolicy, sortby, sortdirection string) (bsorted bool) {
	bsorted = true

	if sortby == "policyid" {
		for i := 0; i < len(idlist)-1; i++ {
			if sortdirection == "ascending" {
				if idlist[i].policyid > idlist[i+1].policyid {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if idlist[i].policyid < idlist[i+1].policyid {
					bsorted = false
					return
				}
			}
		}
	}

	return
}

func moveAfterFirewallProxypolicy(idbefore, idafter int, c *FortiSDKClient, vdomparam string, batch int) (err error) {
	idbefores := strconv.Itoa(idbefore)
	idafters := strconv.Itoa(idafter)

	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/proxy-policy/"
	path += idbefores

	params := make(map[string][]string)
	params["action"] = []string{"move"}
	params["after"] = []string{idafters}

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

func sortPolicyListFirewallProxypolicy(idlist []policySortFirewallProxypolicy, sortby, sortdirection string, c *FortiSDKClient, vdomparam string, batch int) (err error) {
	if sortby == "policyid" {
		if sortdirection == "ascending" {
			sort.Slice(idlist, func(i, j int) bool {
				return idlist[i].policyid < idlist[j].policyid
			})
		} else if sortdirection == "descending" {
			sort.Slice(idlist, func(i, j int) bool {
				return idlist[i].policyid > idlist[j].policyid
			})
		}

		for i := 0; i < len(idlist)-1; i++ {
			err = moveAfterFirewallProxypolicy(idlist[i+1].policyid, idlist[i].policyid, c, vdomparam, batch)
			if err != nil {
				err = fmt.Errorf("sort err %s", err)
				return
			}
		}
	}

	return nil
}

// CreateUpdateFirewallProxypolicySort API operation for FortiOS to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateFirewallProxypolicySort(sortby, sortdirection, vdomparam string, batch int) (err error) {
	idlist, err := getPolicyListFirewallProxypolicy(c, vdomparam, batch)
	log.Printf("shengh: %v", idlist)
	if err != nil {
		err = fmt.Errorf("sort err %s", err)
		return
	}

	bsorted := bPolicyListSortedFirewallProxypolicy(idlist, sortby, sortdirection)
	if bsorted {
		return
	}

	err = sortPolicyListFirewallProxypolicy(idlist, sortby, sortdirection, c, vdomparam, batch)
	if err != nil {
		err = fmt.Errorf("sort err %s", err)
		return
	}

	return
}

// ReadFirewallProxypolicySort API operation for FortiOS to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadFirewallProxypolicySort(sortby, sortdirection string, vdomparam string, batch int) (sorted bool, err error) {
	idlist, err := getPolicyListFirewallProxypolicy(c, vdomparam, batch)
	if err != nil {
		err = fmt.Errorf("sort err %s", err)
		return
	}

	bsorted := bPolicyListSortedFirewallProxypolicy(idlist, sortby, sortdirection)
	log.Printf("shengh: %v", bsorted)
	if bsorted {
		sorted = true
		return
	}

	sorted = false

	return
}
