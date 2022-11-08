package main

import (
	"errors"

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
		"2ade3131f1c0bf782b96f3e098e1433bb563684b66086b741d11779b0af2db20",
	).Dial()

	if client == nil {
		errnie.Handles(errors.New("no client"))
		return
	}

	if client.Auth() == nil {
		errnie.Handles(errors.New("auth fail"))
		return
	}

	// Deploy the smart contract function.
	address, tx, _, err := punchface.DeployPunchface(
		client.Auth(), client.Conn(),
	)

	if err := errnie.Handles(err); err.Type != errnie.NIL {
		return
	}

	errnie.Informs(address.Hex())
	errnie.Informs("tx", tx.Hash().Hex())

	// Get a handle on the smart contract function.
	conn, err := punchface.NewPunchface(
		common.HexToAddress(address.Hex()), client.Conn(),
	)

	if errnie.Handles(err).Type != errnie.NIL {
		return
	}

	// Setup the smart contract server.
	sockpuppet.NewContract(punchface.NewNotary(client, conn)).Up("1323")
}
