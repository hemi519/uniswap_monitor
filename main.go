package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hemi519/uniswap_monitor/datastore"
	"github.com/hemi519/uniswap_monitor/monitor"
)

type EthClientWrapper struct {
	client *ethclient.Client
}

type EthSubscriptionWrapper struct {
	sub ethereum.Subscription
}

func (ew *EthClientWrapper) SubscribeNewHead(ch chan<- *monitor.Header) (monitor.Subscription, error) {
	sub, err := ew.client.SubscribeNewHead(context.Background(), ch)
	if err != nil {
		return nil, err
	}
	return &EthSubscriptionWrapper{sub: sub}, nil
}

func (ew *EthClientWrapper) Context() context.Context {
	return context.Background()
}

func (esw *EthSubscriptionWrapper) Unsubscribe() {
	esw.sub.Unsubscribe()
}

func main() {
	// Create an Ethereum client
	ethereumClient, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Create a data store (replace with your preferred data store implementation)
	datastore := datastore.NewMyDatastore()

	// Define the pools to monitor
	pools := []monitor.PoolConfig{
		{
			Address: "0xPOOL_ADDRESS_1",
		},
		{
			Address: "0xPOOL_ADDRESS_2",
		},
		// Add more pool addresses as needed
	}

	// Create an instance of the UniswapMonitor
	uniswapMonitor := monitor.NewUniswapMonitor(&EthClientWrapper{ethereumClient}, pools, datastore)

	// Start monitoring
	go uniswapMonitor.StartMonitoring()

	// Wait for termination signal to gracefully shutdown
	waitForTerminationSignal()
}

func waitForTerminationSignal() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
