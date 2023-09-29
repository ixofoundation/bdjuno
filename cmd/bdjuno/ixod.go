package main

import (
	bondstypes "github.com/ixofoundation/ixo-blockchain/v2/x/bonds/types"
	claimstypes "github.com/ixofoundation/ixo-blockchain/v2/x/claims/types"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v2/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v2/x/iid/types"
	tokentypes "github.com/ixofoundation/ixo-blockchain/v2/x/token/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v4/modules/messages"
	junomessages "github.com/forbole/juno/v4/modules/messages"
)

// customessageAddressesParser represents a parser able to get the addresses of the involved
// account from a Ixo message
var customessageAddressesParser = junomessages.JoinMessageParsers(
	ixoMessageAddressesParser,
)

// ixoMessageAddressesParser represents a MessageAddressesParser for the ixo x/ modules
func ixoMessageAddressesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {

	// Iid Module
	// ======================================
	case *iidtypes.MsgCreateIidDocument:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgUpdateIidDocument:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgAddVerification:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgRevokeVerification:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgSetVerificationRelationships:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgAddService:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgAddLinkedResource:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgAddLinkedClaim:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgAddLinkedEntity:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgDeleteService:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgDeleteLinkedResource:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgDeleteLinkedClaim:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgDeleteLinkedEntity:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgAddAccordedRight:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgDeleteAccordedRight:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgAddController:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgDeleteController:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgAddIidContext:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgDeleteIidContext:
		return []string{msg.Signer}, nil
	case *iidtypes.MsgDeactivateIID:
		return []string{msg.Signer}, nil

	// Entity Module
	// ======================================
	case *entitytypes.MsgCreateEntity:
		return []string{msg.OwnerAddress}, nil
	case *entitytypes.MsgUpdateEntity:
		return []string{msg.ControllerAddress}, nil
	case *entitytypes.MsgUpdateEntityVerified:
		return []string{msg.RelayerNodeAddress}, nil
	case *entitytypes.MsgTransferEntity:
		return []string{msg.OwnerAddress}, nil
	case *entitytypes.MsgCreateEntityAccount:
		return []string{msg.OwnerAddress}, nil
	case *entitytypes.MsgGrantEntityAccountAuthz:
		return []string{msg.OwnerAddress}, nil

	// Bonds Module
	// ======================================
	case *bondstypes.MsgCreateBond:
		return []string{msg.CreatorAddress}, nil
	case *bondstypes.MsgEditBond:
		return []string{msg.EditorAddress}, nil
	case *bondstypes.MsgSetNextAlpha:
		return []string{msg.OracleAddress}, nil
	case *bondstypes.MsgUpdateBondState:
		return []string{msg.EditorAddress}, nil
	case *bondstypes.MsgBuy:
		return []string{msg.BuyerAddress}, nil
	case *bondstypes.MsgSell:
		return []string{msg.SellerAddress}, nil
	case *bondstypes.MsgSwap:
		return []string{msg.SwapperAddress}, nil
	case *bondstypes.MsgMakeOutcomePayment:
		return []string{msg.SenderAddress}, nil
	case *bondstypes.MsgWithdrawShare:
		return []string{msg.RecipientAddress}, nil
	case *bondstypes.MsgWithdrawReserve:
		return []string{msg.WithdrawerAddress}, nil

	// Claims Module
	// ======================================
	case *claimstypes.MsgCreateCollection:
		return []string{msg.Signer}, nil
	case *claimstypes.MsgSubmitClaim:
		return []string{msg.AdminAddress}, nil
	case *claimstypes.MsgEvaluateClaim:
		return []string{msg.AdminAddress, msg.AgentAddress}, nil
	case *claimstypes.MsgDisputeClaim:
		return []string{msg.AgentAddress, msg.AgentAddress}, nil
	case *claimstypes.MsgWithdrawPayment:
		return []string{msg.AdminAddress, msg.FromAddress, msg.ToAddress}, nil

	// Token Module
	// ======================================
	case *tokentypes.MsgCreateToken:
		return []string{msg.Minter}, nil
	case *tokentypes.MsgMintToken:
		return []string{msg.Minter, msg.Owner}, nil
	case *tokentypes.MsgTransferToken:
		return []string{msg.Owner, msg.Recipient}, nil
	case *tokentypes.MsgRetireToken:
		return []string{msg.Owner}, nil
	case *tokentypes.MsgCancelToken:
		return []string{msg.Owner}, nil
	case *tokentypes.MsgPauseToken:
		return []string{msg.Minter}, nil
	case *tokentypes.MsgStopToken:
		return []string{msg.Minter}, nil

	default:
		return nil, messages.MessageNotSupported(cosmosMsg)
	}
}
