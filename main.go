package main

import (
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/hemi519/uniswap_monitor/datastore"
	"github.com/hemi519/uniswap_monitor/middleware"
	"github.com/hemi519/uniswap_monitor/monitor"
)

const apiKey = "d890adccbc6d4fec887c8d0f23f5e5b7"

func main() {
	// Create an Ethereum client
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Define the pool configurations
	pools := []monitor.PoolConfig{
		{PoolID: "0xPOOL_ADDRESS"},
		// Add more pools as required
	}

	// Create a datastore instance
	datastore := datastore.NewDatastore()

	// Create an instance of the UniswapMonitor
	uniswapMonitor := monitor.NewUniswapMonitor(client, pools, datastore)

	// Start monitoring the pools
	uniswapMonitor.StartMonitoring()

	// Set up the REST API
	router := mux.NewRouter()

	// Apply logging middleware to all API routes
	router.Use(middleware.LoggingMiddleware)

	// Define API routes
	router.HandleFunc("/v1/api/pool/{pool_id}", uniswapMonitor.GetPoolData).Methods("GET")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}
