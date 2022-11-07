package punchface

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/valyala/fasthttp"
)

type PunchFaceNotary struct {
	client *Punchface
}

func NewNotary(client *Punchface) *PunchFaceNotary {
	return &PunchFaceNotary{
		client: client,
	}
}

func (notary *PunchFaceNotary) Validate(
	path string, ctx *fasthttp.RequestCtx,
) (*big.Int, error) {
	var err error
	var transaction *types.Transaction

	switch path {
	case "/admin":
		return big.NewInt(200), nil // notary.client.Admin(&bind.CallOpts{})
	case "/balance":
		return notary.client.Balance(&bind.CallOpts{})
	case "/deposit/:amount":
		transaction, err = notary.client.Deposit(
			&bind.TransactOpts{}, big.NewInt(0),
		)
	case "/withdraw/:amount":
		transaction, err = notary.client.Withdraw(
			&bind.TransactOpts{}, big.NewInt(0),
		)
	}

	transaction.MarshalJSON()
	return big.NewInt(200), err
}
