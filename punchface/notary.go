package punchface

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/theapemachine/wrkspc/errnie"
	"github.com/valyala/fasthttp"
)

type PunchFaceNotary struct {
	client *Punchface
	err    error
}

func NewNotary(client *Punchface) *PunchFaceNotary {
	return &PunchFaceNotary{
		client: client,
	}
}

func (notary *PunchFaceNotary) Validate(
	path string, ctx *fasthttp.RequestCtx,
) (int, string) {
	switch path {
	case "/admin":
		return notary.admin(ctx, notary.client.Admin)
	case "/balance":
		return notary.call(ctx, notary.client.Balance)
	case "/deposit/:amount":
		return notary.transaction(ctx, notary.client.Deposit)
	case "/withdraw/:amount":
		return notary.transaction(ctx, notary.client.Withdraw)
	}

	return 204, ""
}

func (notary *PunchFaceNotary) admin(
	ctx *fasthttp.RequestCtx, fn func(*bind.CallOpts) (common.Address, error),
) (int, string) {
	var addr common.Address

	if addr, notary.err = fn(&bind.CallOpts{}); notary.err != nil {
		return 500, errnie.Handles(notary.err).Msg
	}

	return 200, addr.String()
}

func (notary *PunchFaceNotary) call(
	ctx *fasthttp.RequestCtx, fn func(*bind.CallOpts) (*big.Int, error),
) (int, string) {
	var amount *big.Int

	if amount, notary.err = fn(&bind.CallOpts{}); notary.err != nil {
		return 500, errnie.Handles(notary.err).Msg
	}

	return 200, amount.String()
}

func (notary *PunchFaceNotary) transaction(
	ctx *fasthttp.RequestCtx,
	fn func(*bind.TransactOpts, *big.Int) (*types.Transaction, error),
) (int, string) {
	var transaction *types.Transaction
	var amount int

	if amount, notary.err = ctx.QueryArgs().GetUint("amount"); notary.err != nil {
		return 500, errnie.Handles(notary.err).Msg
	}

	if transaction, notary.err = fn(
		&bind.TransactOpts{}, big.NewInt(int64(amount)),
	); notary.err != nil {
		return 500, errnie.Handles(notary.err).Msg
	}

	var out []byte

	if out, notary.err = transaction.MarshalJSON(); notary.err != nil {
		return 500, errnie.Handles(notary.err).Msg
	}

	return 200, string(out)
}
