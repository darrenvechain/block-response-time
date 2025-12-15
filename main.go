package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	flagNetworkURL = flag.String("network-url", "https://mainnet.vechain.org", "URL of the network to connect to")
	flagBlocks     = flag.Int("blocks", 100, "Number of blocks to fetch")
)

func main() {
	flag.Parse()

	log.Println(fmt.Sprintf("running block fetch test against %s for %d blocks", *flagNetworkURL, *flagBlocks))

	startBlock := 23_000_000 // using a historic non-cached block for testing

	start := time.Now()
	for i := range *flagBlocks {
		resp, err := http.Get(fmt.Sprintf("%s/blocks/%d?expanded=true", *flagNetworkURL, startBlock+i))
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
	}
	elapsed := time.Since(start)
	avg := elapsed / time.Duration(*flagBlocks)

	log.Printf("Fetched %d blocks sequentially in %s (avg %s per block)", *flagBlocks, elapsed, avg)

	startBlock = 23_001_000 // reset to a different historic block
	var wg sync.WaitGroup
	errs := make(chan error, *flagBlocks)
	start = time.Now()
	for i := range *flagBlocks {
		blockNum := startBlock + i
		wg.Go(func() {
			resp, err := http.Get(fmt.Sprintf("%s/blocks/%d?expanded=true", *flagNetworkURL, blockNum))
			if err == nil {
				resp.Body.Close()
			}
			errs <- err
		})
	}
	wg.Wait()
	close(errs)
	elapsed = time.Since(start)

	for err := range errs {
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Fetched %d blocks concurrently in %s", *flagBlocks, elapsed)
}
