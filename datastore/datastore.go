package datastore

//	"github.com/hemi519/uniswap_monitor/monitor"

type Datastore interface {
	//SaveData(data monitor.DataPoint) error
	//GetDataByBlock(poolID string, blockNumber uint64) (monitor.DataPoint, error)
}

type YourDatastore struct {
	// Add necessary fields for your chosen datastore implementation
}

func NewDatastore() *YourDatastore {
	return &YourDatastore{}
	// Initialize and return your datastore instance
}

// func (d *YourDatastore) SaveData(data monitor.DataPoint) error {
// 	// Implement saving the data to your chosen datastore
// }

// func (d *YourDatastore) GetDataByBlock(poolID string, blockNumber uint64) (monitor.DataPoint, error) {
// 	// Implement retrieving data from your chosen datastore based on the provided pool ID and block number
// }
