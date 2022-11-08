package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/theapemachine/web3-solidity-experiments/punchface"
	"github.com/theapemachine/wrkspc/errnie"
	"github.com/theapemachine/wrkspc/sockpuppet"
)

func main() {
	errnie.Tracing(true)
	errnie.Debugs(true)

	client := sockpuppet.NewEthClient(
		"http://127.0.0.1:7545",
		"fd4eef6dec5575cc78f3f14d4b749094f8b88ad7883caaa8d1d24e9a01e3732d",
	).Dial()

	// Deploy the smart contract function.
	address, tx, instance, err := punchface.DeployPunchface(
		client.Auth(), client.Conn(),
	)

	if err := errnie.Handles(err); err.Type != errnie.NIL {
		return
	}

	errnie.Logs(address.Hex())

	errnie.Logs("instance", instance)
	errnie.Logs("tx", tx.Hash().Hex())

	// Get a handle on the smart contract function.
	conn, err := punchface.NewPunchface(
		common.HexToAddress(address.Hex()), client.Conn(),
	)
	errnie.Handles(err)

	// Setup the smart contract server.
	sockpuppet.NewContract(punchface.NewNotary(conn)).Up("1323")
}
