package main

import (
	"github.com/mrb/riakpbc"
	"log"
	"runtime"
	"time"
)

type Data struct {
	Data string `json:"data"`
}

func main() {
	runtime.GOMAXPROCS(7)
	cluster := []string{"127.0.0.1:8087", "127.0.0.1:8088", "127.0.0.1:8089", "127.0.0.1:8090"}
	riak := riakpbc.NewClient(cluster)

	err := riak.Dial()
	if err != nil {
		log.Print(err)
	}

	var actionEnd time.Time
	actionBegin := time.Now()

	c := make(chan int)

	for g := 0; g < 7; g++ {
		go func(which int) {
			log.Print("<", which, "> Loaded")
			var times int
			var errs int
			for {
				times = times + 1
				_, err := riak.StoreObject("bucket", "data", "{'ok':'ok'}")
				if err != nil {
					errs = errs + 1
				}

				_, err = riak.SetClientId("coolio")
				if err != nil {
					errs = errs + 1
				}

				_, err = riak.GetClientId()
				if err != nil {
					errs = errs + 1
				}

				data, err := riak.FetchObject("bucket", "data")
				if err != nil {
					errs = errs + 1
				} else {
					if string(data.GetContent()[0].GetValue()) != "{'ok':'ok'}" {
						log.Fatal("!!!")
					}
				}

				_, err = riak.StoreObject("bucket", "moreData", "stringData")
				if err != nil {
					errs = errs + 1
				}

				_, err = riak.FetchObject("bucket", "moreData")
				if err != nil {
					errs = errs + 1
				}

				log.Print("<", which, "> @", times, " ", riak.Pool(), "!<", errs, "> ")
			}
		}(g)
	}
	<-c
	actionEnd = time.Now()
	actionDuration := actionEnd.Sub(actionBegin)
	log.Print("Ran for ", actionDuration)

	riak.Close()
}
