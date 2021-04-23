package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	getAccountInfo = "http://127.0.0.1:5001/tapi/v3/?tapi_method=get_account_info"
)

type Level struct {
	fires             int
	requestsPerSecond int
}
type Specs struct {
	name   string
	levels []Level
}

var (
	successCount int
	errorCount   int
)

func main() {
	var title string
	flag.StringVar(&title, "title", "?", "Name of the test")
	flag.Parse()

	log.Println("Stress Tester:", title)

	successCount = 0
	errorCount = 0

	startTime := time.Now()
	specs := Specs{
		name: title,
		levels: []Level{
			{fires: 100, requestsPerSecond: 1},
			{fires: 100, requestsPerSecond: 5},
			{fires: 100, requestsPerSecond: 10},
			{fires: 100, requestsPerSecond: 15},
			{fires: 100, requestsPerSecond: 20},
			{fires: 100, requestsPerSecond: 25},
			{fires: 100, requestsPerSecond: 30},
		},
	}
	RunStressTest(specs)

	log.Println("Time ", time.Since(startTime))
}

func RunStressTest(specs Specs) {
	var wg sync.WaitGroup
	for l := 0; l < len(specs.levels); l++ {
		level := specs.levels[l]
		wg.Add(level.fires * level.requestsPerSecond)
		for f := 0; f < level.fires; f++ {
			for i := 1; i <= level.requestsPerSecond; i++ {
				go Request(&wg)
			}
			time.Sleep(time.Second)
		}
		wg.Wait()
		totalReq := errorCount + successCount
		errorPercentage := float32(errorCount) * 100 / float32(totalReq)
		log.Printf("Level: %d | Fires: %d | Req/Second: %d | Total: %d | Errors: %d (%.2f%%) \n", l, level.fires, level.requestsPerSecond, totalReq, errorCount, errorPercentage)
		successCount = 0
		errorCount = 0
	}
}

func Request(wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(getAccountInfo)
	if err != nil {
		// log.Println("Get:", err.Error())
		errorCount += 1
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// log.Println("ReadAll:", err.Error())
		errorCount += 1
		return
	}

	var jsonBody map[string]interface{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		// log.Println("Unmarshal:", err.Error())
		errorCount += 1
		return
	}

	statusCode := int(jsonBody["status_code"].(float64))
	if statusCode == 100 {
		successCount += 1
	} else {
		errorCount += 1
	}
}
