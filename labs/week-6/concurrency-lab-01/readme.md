### Website status checker

Write a program to test the status of the website. You need to make a http GET call and print the status as UP or DOWN

1. Make the program to run using go routines
2. make use of channels to return the status of the website

<details>
  <summary>Not sure how?</summary>

  ```go
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	topDomainNames, err := readDomainsFromCSV()

	t1 := time.Now()
	if err != nil {
		log.Fatal("Error while reading domains ", err.Error())
	}
	statusChecker(topDomainNames)

	fmt.Println("\nTotal time: ", time.Since(t1))
}

func readDomainsFromCSV() ([]string, error) {
	file, err := ioutil.ReadFile("topDomains.csv")
	if err != nil {
		log.Println("Error while reading file...", err.Error())
		return nil, err
	}

	var domains []string
	reader := csv.NewReader(bytes.NewReader(file))
	all, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error while reading records from csv...", err.Error())
		return nil, err
	}
	for _, rec := range all {
		domains = append(domains, rec[1])
	}
	if len(domains) > 0 {
		return domains[1:], nil
	}
	return domains, nil
}

func statusChecker(topDomainNames []string) {
	statusCh := make(chan string)
	wg := &sync.WaitGroup{}

	wg.Add(len(topDomainNames))
	for idx, host := range topDomainNames {
		go lookupHost(host, idx, statusCh, wg)
		// lookupHost(host, idx)
	}

	// Start a goroutine to close the statusCh once all lookups are done
	go func() {
		wg.Wait()
		close(statusCh)
	}()

	for data := range statusCh {
		fmt.Println(data)
	}
	fmt.Println("done here...")
}

func lookupHost(host string, idx int, statusCh chan string, wg *sync.WaitGroup) {
	// func lookupHost(host string, idx int) {
	defer wg.Done()
	client := http.Client{
		Timeout: time.Second * 2,
	}
	_, err := client.Get(host)
	if err != nil {
		// fmt.Println(fmt.Sprintf("%d, %s website is down", idx, host))
		statusCh <- fmt.Sprintf("%d, %s website is down", idx, host)
		// return
	} else {
		// fmt.Println(fmt.Sprintf("%d, %s website is up", idx, host))
		statusCh <- fmt.Sprintf("%d, %s website is up", idx, host)
	}
}
  ```
</details>
<br>
