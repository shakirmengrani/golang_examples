package main

import (
	"log"
	"time"
	"github.com/nats-io/go-nats"
)

func main(){
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subj, msg := "foo", []byte("Hello World")
	nc.Publish(subj, msg)
	nc.Flush()
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}

	subj, i := "foo", 0
	nc.Subscribe(subj, func(msg *nats.Msg) {
		i += 1
		log.Printf("[#%d] Received on [%s]: '%s'\n", i, msg.Subject, string(msg.Data))
	})
	nc.Flush()
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on [%s]\n", subj)
	
	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe("foo", ch)
	ch_msg := <- ch
	log.Printf("Received from channel subscriber on [%s]: '%s'\n", "foo", string(ch_msg.Data))
	sub.Unsubscribe()

	// Requests
	r_msg, err := nc.Request("help", []byte("help me"), 10*time.Second)
	if err != nil{
		log.Fatal(err)
	}
	log.Printf("Requested to [%s], '%s'\n", "help", string(r_msg.Data))
	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})	

}