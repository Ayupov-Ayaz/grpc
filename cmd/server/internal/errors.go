package internal

import (
	"errors"
	"fmt"

	"log"

	"github.com/Ayupov-Ayaz/grpc/cmd/server/internal/transaction"
	"github.com/Ayupov-Ayaz/grpc/cmd/server/internal/wallet"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/protoadapt"

	api "github.com/Ayupov-Ayaz/grpc/gen/go/aayupov/wallet/v1alpha1"
)

type walletError interface {
	UserID() string
	Error() string
}

type transactionError interface {
	TransactionID() transaction.ID
	Error() string
}

const (
	walletResourceType      = api.ResourceType_RESOURCE_TYPE_WALLET
	transactionResourceType = api.ResourceType_RESOURCE_TYPE_TRANSACTION
)

const (
	insufficientFundViolationType = api.ViolationType_VIOLATION_TYPE_INSUFFICIENT_FUNDS
)

func wrapCreateWalletError(err error) error {
	if e := new(wallet.NotFoundError); errors.As(err, &e) {
		return errWalletNotFoundGRPCError(e)
	}

	if e := new(wallet.AlreadyExistsError); errors.As(err, &e) {
		return errWalletAlreadyExistsGRPC(e)
	}

	return err
}

func wrapChangeWalletError(err error) error {
	if e := new(wallet.NotFoundError); errors.As(err, &e) {
		return errWalletNotFoundGRPCError(e)

	}

	if e := new(wallet.InsufficientFundsError); errors.As(err, &e) {
		return errInsufficientFundsGRPC(e)
	}

	return err
}

func wrapCheckBetTransactionError(err error) error {
	if e := new(transaction.NotFoundError); errors.As(err, &e) {
		return errTransactionNotFoundGRPC(e)
	}

	return err
}

func wrapSetTransactionError(err error) error {
	if e := new(transaction.AlreadyExistsError); errors.As(err, &e) {
		return errTransactionAlreadyExists(e)
	}

	return err
}

func errWalletAlreadyExistsGRPC(err *wallet.AlreadyExistsError) error {
	return NewGRPCError(codes.AlreadyExists, "wallet already exists",
		newWalletResourceInfo(err))
}

func errWalletNotFoundGRPCError(err *wallet.NotFoundError) error {
	return NewGRPCError(codes.NotFound, "wallet not found",
		newWalletResourceInfo(err))
}

func errInsufficientFundsGRPC(err *wallet.InsufficientFundsError) error {
	desc := fmt.Sprintf("not enough money for bet=%d with user balance=%d",
		err.Bet(), err.Balance())

	violation := &errdetails.PreconditionFailure{
		Violations: []*errdetails.PreconditionFailure_Violation{
			{
				Type:        insufficientFundViolationType.String(),
				Subject:     walletResourceName(err.UserID()),
				Description: desc,
			},
		},
	}

	return NewGRPCError(codes.FailedPrecondition, "insufficient funds", violation)
}

func errTransactionNotFoundGRPC(err *transaction.NotFoundError) error {
	return NewGRPCError(codes.NotFound, "transaction not found",
		newTransactionResourceInfo(err))
}

func errTransactionAlreadyExists(err *transaction.AlreadyExistsError) error {
	return NewGRPCError(codes.AlreadyExists, "transaction already exists",
		newTransactionResourceInfo(err))
}

func NewGRPCError(
	code codes.Code,
	message string,
	details ...protoadapt.MessageV1,
) error {
	s, err := status.New(code, message).WithDetails(details...)
	if err != nil {
		log.Println("failed to create gRPC error with details",
			"error", err,
			"code", code,
			"message", message)

		// create a status without details
		return status.Error(code, message)
	}

	return s.Err()
}

func newTransactionResourceInfo(tErr transactionError) *errdetails.ResourceInfo {
	return &errdetails.ResourceInfo{
		ResourceType: transactionResourceType.String(),
		ResourceName: transactionResourceName(tErr.TransactionID()),
		Owner:        tErr.TransactionID().GetUserID(),
		Description:  tErr.Error(),
	}
}

func newWalletResourceInfo(wErr walletError) *errdetails.ResourceInfo {
	return &errdetails.ResourceInfo{
		ResourceType: walletResourceType.String(),
		ResourceName: walletResourceName(wErr.UserID()),
		Owner:        wErr.UserID(),
		Description:  wErr.Error(),
	}
}

func transactionResourceName(id transaction.ID) string {
	return "user_id=" + id.GetUserID() +
		"/tr_id=" + id.GetOperationID()
}

func walletResourceName(userID string) string {
	return "user_id=" + userID
}
