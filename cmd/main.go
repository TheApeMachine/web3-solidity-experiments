package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/theapemachine/web3-solidity-experiments/punchface"
	"github.com/theapemachine/wrkspc/errnie"
	"github.com/theapemachine/wrkspc/sockpuppet"
)

type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

func (manager *Manager) Manage() {}

func main() {
	client := sockpuppet.NewEthClient(
		"http://127.0.0.1:7545",
		"fd4eef6dec5575cc78f3f14d4b749094f8b88ad7883caaa8d1d24e9a01e3732d",
	)

	address, tx, instance, err := punchface.DeployPunchface(
		client.Auth(), client.Conn(),
	)
	errnie.Handles(err)
	errnie.Logs(address.Hex())

	_, _ = instance, tx
	errnie.Logs("instance", instance)
	errnie.Logs("tx", tx.Hash().Hex())

	//creating api object to intract with smart contract function
	conn, err := punchface.NewPunchface(
		common.HexToAddress(address.Hex()), client.Conn(),
	)
	errnie.Handles(err)

	srv := sockpuppet.NewContract(
		punchface.NewNotary(conn),
	)
	srv.Up("80")
}
