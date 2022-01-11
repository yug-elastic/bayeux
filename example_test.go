package bayeux

import (
	"fmt"
)

func Example() {
	out := make(chan TriggerEvent)
	replay := "-1"
	b := Bayeux{}
	creds := GetSalesforceCredentials()
	c := b.Channel(out, replay, creds, "channel")
	for {
		select {
		case e := <-c:
			fmt.Printf("TriggerEvent Received: %+v", e)
		}
	}
}
