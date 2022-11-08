package punchface

import (
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/theapemachine/wrkspc/errnie"
	"github.com/theapemachine/wrkspc/sockpuppet"
	"github.com/valyala/fasthttp"
)

type PunchFaceNotary struct {
	conn   *sockpuppet.EthClient
	client *Punchface
	err    error
}

func NewNotary(conn *sockpuppet.EthClient, client *Punchface) *PunchFaceNotary {
	return &PunchFaceNotary{
		conn:   conn,
		client: client,
	}
}

func (notary *PunchFaceNotary) Validate(ctx *fasthttp.RequestCtx) (int, string) {
	split := strings.Split(string(ctx.Path()), "/")
	errnie.Debugs(split)

	switch split[1] {
	case "admin":
		return notary.admin(ctx, notary.client.Admin)
	case "balance":
		return notary.call(ctx, notary.client.Balance)
	case "deposit":
		return notary.transaction(ctx, notary.client.Deposit, split[2])
	case "withdraw":
		return notary.transaction(ctx, notary.client.Withdraw, split[2])
	}

	return http.StatusNotFound, `{"status": 404, "error": "not found"}`
}

func (notary *PunchFaceNotary) admin(
	ctx *fasthttp.RequestCtx, fn func(*bind.CallOpts) (common.Address, error),
) (int, string) {
	errnie.Traces()
	var addr common.Address

	if addr, notary.err = fn(&bind.CallOpts{}); notary.err != nil {
		return 500, errnie.Handles(notary.err).Msg
	}

	errnie.Debugs(addr.String())
	return 200, addr.String()
}

func (notary *PunchFaceNotary) call(
	ctx *fasthttp.RequestCtx, fn func(*bind.CallOpts) (*big.Int, error),
) (int, string) {
	errnie.Traces()
	var amount *big.Int

	if amount, notary.err = fn(&bind.CallOpts{}); notary.err != nil {
		return 500, errnie.Handles(notary.err).Msg
	}

	errnie.Debugs(amount.String())
	return 200, amount.String()
}

func (notary *PunchFaceNotary) transaction(
	ctx *fasthttp.RequestCtx,
	fn func(*bind.TransactOpts, *big.Int) (*types.Transaction, error),
	arg string,
) (int, string) {
	errnie.Traces()
	var transaction *types.Transaction
	var amount int

	if amount, notary.err = strconv.Atoi(arg); notary.err != nil {
		return 500, errnie.Handles(notary.err).Msg
	}

	if transaction, notary.err = fn(
		notary.conn.Auth(), big.NewInt(int64(amount)),
	); notary.err != nil {
		errnie.Inspects(transaction)
		return 500, errnie.Handles(notary.err).Msg
	}

	errnie.Inspects(transaction)

	var out []byte

	if out, notary.err = transaction.MarshalJSON(); notary.err != nil {
		return 500, errnie.Handles(notary.err).Msg
	}

	return 200, string(out)
}
