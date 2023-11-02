package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"loan/x/loan/types"
)

func (k msgServer) ApproveLoan(goCtx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
        return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
    }

	if loan.State != "requested" {
        return nil, sdkerrors.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
    }

	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	amount, err := sdk.ParseCoinsNormalized(loan.Amount)
	if err != nil {
        return nil, sdkerrors.Wrap(types.ErrWrongLoanState, "Cannot parse coins in loan amount")
    }
	
	err = k.bankKeeper.SendCoins(ctx, lender, borrower, amount)
	if err != nil {
        return nil, err
    }

	loan.State = "approved"
	loan.Lender = msg.Creator
	k.SetLoan(ctx, loan)
	return &types.MsgApproveLoanResponse{}, nil
}
