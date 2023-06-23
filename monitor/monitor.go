package monitor

import (
	"context"
	"fmt"
	"log"
	"sync"

	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hemi519/uniswap_monitor/datastore"
)

type EthereumClient interface {
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (EthereumSubscription, error)
}

type EthereumSubscription interface {
	Err() <-chan error
	Unsubscribe()
}

type Header struct {
	Number *big.Int

	// Add other required fields
}

type PoolConfig struct {
	Address string
	// Add other required fields
}

type UniswapMonitor struct {
	client    EthereumClient
	pools     []PoolConfig
	datastore datastore.Datastore
}

func NewUniswapMonitor(client EthereumClient, pools []PoolConfig, datastore datastore.Datastore) *UniswapMonitor {
	return &UniswapMonitor{
		client:    client,
		pools:     pools,
		datastore: datastore,
	}
}

func (m *UniswapMonitor) StartMonitoring() {
	log.Println("Starting Uniswap monitoring...")

	headers := make(chan *types.Header)
	sub, err := m.client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("Failed to subscribe to new headers: %v", err)
	}
	defer sub.Unsubscribe()

	var wg sync.WaitGroup
	wg.Add(1)
	go m.processHeaders(headers, &wg)

	wg.Wait()
}

func (m *UniswapMonitor) processHeaders(headers <-chan *types.Header, wg *sync.WaitGroup) {
	defer wg.Done()

	for header := range headers {
		blockNumber := header.Number

		// Fetch and calculate the required data points for each pool
		for _, pool := range m.pools {
			//	token0Balance, token1Balance, tick, err := m.fetchDataPoints(pool.Address, blockNumber)
			_, _, _, err := m.fetchDataPoints(pool.Address, blockNumber)

			if err != nil {
				log.Printf("Failed to fetch data points for pool %s: %v", pool.Address, err)
				continue
			}

			// Store the data points in the datastore
			// err = m.datastore.SaveDataPoints(pool.Address, token0Balance, token1Balance, tick, blockNumber)
			// if err != nil {
			// 	log.Printf("Failed to save data points for pool %s: %v", pool.Address, err)
			// 	continue
			// }

			log.Printf("Data points saved for pool %s at block %s", pool.Address, blockNumber.String())
		}
	}
}

func (m *UniswapMonitor) fetchDataPoints(poolAddress string, blockNumber *big.Int) (token0Balance, token1Balance *big.Int, tick int64, err error) {
	// Implement the logic to fetch the required data points from the pool contract
	// You can use the Ethereum client or any other necessary packages for this task
	// Return the fetched data points and any potential errors
	return nil, nil, 0, fmt.Errorf("fetchDataPoints not implemented")
}

func (m *UniswapMonitor) GetBalances(poolAddress string, blockNumber *big.Int) (token0Balance, token1Balance *big.Int, tick int64, err error) {
	// Retrieve the balances and tick from the datastore based on the pool address and block number
	// Implement the logic to fetch the required data points from the datastore
	// Return the retrieved balances and tick, along with any potential errors
	return nil, nil, 0, nil
}
