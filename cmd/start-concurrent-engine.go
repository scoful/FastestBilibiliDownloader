package main

import (
	"fmt"
	"log"
	"os"
	"simple-golang-crawler/engine"
	"simple-golang-crawler/parser"
	"simple-golang-crawler/persist"
	"simple-golang-crawler/scheduler"
	"strconv"
	"sync"
)

func main() {
	itemProcessFun := persist.GetItemProcessFun()
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	itemChan, err := itemProcessFun(&wg)
	if err != nil {
		panic(err)
	}

	var idType string
	var id string
	var req *engine.Request
	fmt.Println("Please enter your id type(`aid` or `upid` or `bvid`)")
	fmt.Scan(&idType)
	fmt.Println("Please enter your id")
	fmt.Scan(&id)

	if idType == "aid" {
		intId, _ := strconv.ParseInt(id, 10, 64)
		req = parser.GetRequestByAid(intId)
	} else if idType == "upid" {
		intId, _ := strconv.ParseInt(id, 10, 64)
		req = parser.GetRequestByUpId(intId)
	} else if idType == "bvid" {
		req = parser.GetRequestByBvid(id)
	} else {
		log.Fatalln("Wrong type you enter")
		os.Exit(1)
	}

	queueScheduler := scheduler.NewConcurrentScheduler()
	conEngine := engine.NewConcurrentEngine(30, queueScheduler, itemChan)
	log.Println("Start working.")
	conEngine.Run(req)
	wg.Wait()
	log.Println("All work has done")
}
