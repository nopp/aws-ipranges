package main

import (
	"aws-ipranges/ipranges"
	"flag"
	"fmt"
	"os"
)

func main() {

	service := flag.String("service", "", "ex: API_GATEWAY")
	region := flag.String("region", "", "ex: sa-east-1, us-east-1")
	json := flag.Bool("json", true, "ex: true or false")

	flag.Parse()

	if (*service == "") || (*region == "") {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *json {
		fmt.Println(string(ipranges.ReturnRangeJSON(*service, *region)))
		os.Exit(0)
	}

	ipranges.ReturnRange(*service, *region)
	os.Exit(0)

}
