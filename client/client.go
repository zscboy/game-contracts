package client

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type client struct {
	Cfg    Config
	Client *ethclient.Client
}

func New(opts ...Option) (*client, error) {
	cfg := defaultConfig()

	for _, opt := range opts {
		opt(&cfg)
	}

	c, err := ethclient.Dial(cfg.endpoint)
	if err != nil {
		return nil, err
	}

	return &client{
		Cfg:    cfg,
		Client: c,
	}, nil
}

// InvokeContract Invoke an EVM smart contract
func (c *client) InvokeContract(invokeFunc func(opts *bind.TransactOpts) (*types.Transaction, error)) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(c.Cfg.privateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	signer := types.LatestSignerForChainID(big.NewInt(c.Cfg.networkID))
	opts := &bind.TransactOpts{
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(transaction, signer, privateKey)
		},
		From:    fromAddress,
		Context: context.Background(),
	}

	tx, err := invokeFunc(opts)
	if err != nil {
		return nil, err
	}

	return tx.MarshalJSON()
}
