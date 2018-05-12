package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
	"strconv"
	"time"
	"log"
)

func main() {

	start_time := time.Now()
	counter := 0

	for i := 0; i < 10; i++ {

		url := "http://ipv4.download.thinkbroadband.com:8080/5MB.zip"
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create(path + "/data_dump/" + strconv.Itoa(counter))
		if err != nil {
			fmt.Println("path not created")
		}

		requester, err := http.Get(url)
		if err != nil {
			fmt.Println("no get request possible on url")
		}

		defer requester.Body.Close()
		_, err = io.Copy(out, requester.Body)
		if err != nil {
			fmt.Println("failure on write to data_dump")
		}

		fmt.Println("Thread " + strconv.Itoa(counter) + ": done")
		//work()
		counter += 1
	}

	elapsed_time := time.Since(start_time)
	fmt.Println("%s execution time single threaded load", elapsed_time)

}

func work() {
	for i := 0; i < 10000; i++{

	}
}
