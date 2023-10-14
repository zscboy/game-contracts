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
	contracts "github.com/zscboy/game-contracts/contracts"
)

func TestSetGameInfo(t *testing.T) {
	c, err := client.New(
		client.PrivateKeyOption(os.Getenv("PRIVATE_KEY")),
	)
	if err != nil {
		t.Fatal("new client err ", err.Error())
	}

	roundID := uuid.NewString()
	vpsContractAddress := common.HexToAddress("0x0f35f7BE96936D68612d438302D1De7065376A2d")
	result, err := c.InvokeContract(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		instance, err := contracts.NewGame(vpsContractAddress, c.Client)
		if err != nil {
			return nil, err
		}

		gameInfo := contracts.GameInfoContractGameInfo{
			DomainSeparationTag: 1,
			ReplayCID:           uuid.NewString(),
			VRFHeight:           1,
		}
		return instance.SetGameInfo(opts, roundID, gameInfo)
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("roundID id: ", roundID)
	t.Log("deploy OK: ", string(result))
	t.Log("querying game info...")

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-time.After(2 * time.Minute):
			t.Log("query order timeout!")
			return
		case <-ticker.C:
			instance, err := contracts.NewGame(vpsContractAddress, c.Client)
			if err != nil {
				t.Logf("new game instance err %s", err.Error())
				continue
			}
			order, err := instance.GetGameInfo(&bind.CallOpts{
				Pending: true,
			}, roundID)
			if err != nil {
				t.Logf("get game info err %s", err.Error())
				continue
			}

			t.Log("Query game info OK: ", order)
			return
		}
	}

}

func TestDeployGame(t *testing.T) {
	c, err := client.New(
		client.PrivateKeyOption(os.Getenv("PRIVATE_KEY")),
	)
	if err != nil {
		t.Fatal("new client err", err.Error())
	}

	result, err := c.InvokeContract(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		addr, tr, _, err := contracts.DeployGame(opts, c.Client)
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