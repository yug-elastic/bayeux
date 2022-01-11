package main

import (
	"fmt"

	bay "github.com/kush-elastic/bayeux"
)

func Example() {
	out := make(chan bay.TriggerEvent)
	b := bay.Bayeux{}
	creds := bay.GetSalesforceCredentials()
	replay := "-1"
	c := b.Channel(out, replay, creds, "channel")
	for {
		select {
		case e := <-c:
			fmt.Printf("TriggerEvent Received: %+v", e)
		}
	}
}

func main() {
	Example()
}
