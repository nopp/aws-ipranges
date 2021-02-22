package main

import (
	"aws-ipranges/ipranges"
	"flag"
	"os"
)

func main() {

	service := flag.String("service", "", "ex: API_GATEWAY")
	region := flag.String("region", "", "ex: sa-east-1, us-east-1")

	flag.Parse()

	if (*service == "") || (*region == "") {
		flag.PrintDefaults()
		os.Exit(1)
	}

	ipranges.ReturnRange(*service, *region)
	os.Exit(0)

}
