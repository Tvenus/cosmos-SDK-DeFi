package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"loan/x/loan/types"
)

func (k msgServer) RequestLoan(goCtx context.Context, msg *types.MsgRequestLoan) (*types.MsgRequestLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	var loan = types.Loan {
		Amount: msg.Amount,
		Fee: msg.Fee,
		Collateral: msg.Collateral,
		Deadline: msg.Deadline,
		State: "requested",
		Borrower: msg.Creator,
	}
	borrower, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
        panic(err)
    }

	collateral, err := sdk.ParseCoinsNormalized(loan.Collateral)
	if err != nil {
        panic(err)
    }

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, collateral)
	if sdkError != nil {
        return nil, sdkError
    }

	k.AppendLoan(ctx, loan);

	return &types.MsgRequestLoanResponse{}, nil
}
