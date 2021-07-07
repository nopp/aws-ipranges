package ipranges

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

const (
	awsURLRanges = "https://ip-ranges.amazonaws.com/ip-ranges.json"
)

type awsRangesInfo struct {
	SyncToken  string `json:"syncToken"`
	CreateDate string `json:"createDate"`
	Prefixes   []struct {
		IPPrefix string `json:"ip_prefix"`
		Region   string `json:"region"`
		Service  string `json:"service"`
	} `json:"prefixes"`
}

type prefixesInfo struct {
	IPPrefix string `json:"ip_prefix"`
	Region   string `json:"region"`
	Service  string `json:"service"`
}

// ReturnRange - Return IP(ranges) based on Region and Service
func ReturnRange(service, region string) {
	var awsRanges awsRangesInfo

	resp, err := http.Get(awsURLRanges)
	if err != nil {
		log.Fatal(err.Error())
	}

	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		log.Println(bodyErr)
	}

	// Json to Struct
	_ = json.Unmarshal(body, &awsRanges)

	// Filter Json by Service and Region
	var listPrefixes []prefixesInfo
	for _, v := range awsRanges.Prefixes {
		if v.Region == region && v.Service == service {
			listPrefixes = append(listPrefixes, v)
		}
	}

	// Struct to Json
	listPrefixesJSON, err := json.Marshal(listPrefixes)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(listPrefixesJSON))
}
