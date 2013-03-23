package main

// https://www.unbound.net/documentation/libunbound-tutorial-1.html

import (
	"fmt"
	"github.com/miekg/unbound"
	"log"
)

func main() {
	u := unbound.New()
	defer u.Destroy()

	u.ResolvConf("/etc/resolv.conf")
	r, err := u.LookupHost("www.nlnetlabs.nl")
	if err != nil {
		log.Fatalf("error %s\n", err.Error())
	}
	fmt.Printf("%+v\n", r)
}