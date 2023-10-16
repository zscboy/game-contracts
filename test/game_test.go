package main

import (
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"github.com/zscboy/game-contracts/client"
	contracts "github.com/zscboy/game-contracts/gen"
)

const (
	endpoint  = "https://api.calibration.node.glif.io/"
	networkID = 314159

	contractAddress = "0xCA091fa7b2066A10Bf5d9815881B293c9a2c4E7b"
)

func TestDeployGame(t *testing.T) {
	c, err := client.New(
		client.PrivateKeyOption(os.Getenv("PRIVATE_KEY")),
		client.EndpointOption(endpoint),
		client.NetworkIDOption(networkID),
	)

	if err != nil {
		t.Fatal("new client err", err.Error())
	}

	result, err := c.InvokeContract(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		addr, tr, _, err := contracts.DeployGameReplayContract(opts, c.EthClient())
		if err != nil {
			return nil, err
		}

		t.Logf("deploy contract %s", addr.Hex())
		return tr, nil
	})
	if err != nil {
		t.Fatal("deploy contracts err ", err.Error())
	}

	t.Log("deploy OK: ", string(result))
}

func TestSaveGameInfo(t *testing.T) {
	c, err := client.New(
		client.PrivateKeyOption(os.Getenv("PRIVATE_KEY")),
		client.EndpointOption(endpoint),
		client.NetworkIDOption(networkID),
	)

	if err != nil {
		t.Fatal("new client err ", err.Error())
	}

	myAddress, err := c.Address()
	if err != nil {
		t.Fatal("get address err ", err.Error())
	}

	roundID := uuid.NewString()
	gameContractAddress := common.HexToAddress(contractAddress)
	instance, err := contracts.NewGameReplayContract(gameContractAddress, c.EthClient())
	if err != nil {
		t.Fatal("new contract instance err ", err.Error())
	}

	result, err := c.InvokeContract(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		gameInfo := contracts.GameInfo{
			GameID:    "123",
			RoundID:   roundID,
			PlayerIDs: "1,2,3,4",
			ReplayID:  uuid.NewString(),
		}

		gameResult := contracts.GameResult{}
		gameReplay := contracts.GameReplay{
			DomainSeparationTag: 1,
			Address:             myAddress.Hex(),
			GameInfo:            gameInfo,
			VRFHeight:           1,
			VRFProof:            []byte{},
			GameResult:          gameResult,
		}
		return instance.SaveGameReplay(opts, []contracts.GameReplay{gameReplay})
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("roundID id: ", roundID)
	t.Log("save game replay OK: ", string(result))
	t.Log("querying game replay...")

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-time.After(2 * time.Minute):
			t.Log("query order timeout!")
			return
		case <-ticker.C:
			order, err := instance.GetGameReplay(&bind.CallOpts{
				Pending: true,
			}, roundID)
			if err != nil {
				t.Logf("get game replay err %s", err.Error())
				continue
			}

			t.Log("Query game replay OK: ", order)
			return
		}
	}

}
