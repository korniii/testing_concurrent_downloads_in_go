package main

import (
	"os"
	"strconv"
	"fmt"
	"net/http"
	"io"
	"sync"
	"time"
	"runtime"
	"log"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(10)
	start_time := time.Now()

	counter := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go download(counter)
		counter += 1
	}
	wg.Wait()

	elapsed_time := time.Since(start_time)
	fmt.Println("%s execution time multi threaded load", elapsed_time)

}

func download(counter int) {

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
	//work_some()
	fmt.Println("Thread " + strconv.Itoa(counter) + ": done")
	wg.Done()
}

func work_some() {
	for i := 0; i < 10000; i++{

	}
}
