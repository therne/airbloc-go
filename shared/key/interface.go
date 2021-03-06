package key

import (
	"github.com/airbloc/airbloc-go/shared/types"
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type Manager interface {
	NodeKey() *Key
	EncryptData(*types.Data) (*types.EncryptedData, error)
	DecryptData(*types.EncryptedData) (*types.Data, error)
	Encrypt(string) ([]byte, error)
	Decrypt([]byte) (string, error)
	SignEthTx(*ethTypes.Transaction) (*ethTypes.Transaction, error)
	SignBDBTx(*txn.Transaction) error
}
