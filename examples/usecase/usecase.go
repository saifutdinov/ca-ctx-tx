package usecase

import (
	"context"
	ctxtx "saifutdinov/ca-ctx-tx"
	"saifutdinov/ca-ctx-tx/examples/domain"
)

type Usecase struct {
	ctxtx.TXI
	repository domain.Repository
}

func NewUsecase(
	txs ctxtx.TXI,
	repository domain.Repository,
) domain.Usecase {
	// here we have repo with 'wrapped' DBI and txs - out transactions DBI.
	return &Usecase{
		TXI:        txs,
		repository: repository,
	}
}

// default method without transaction
func (u *Usecase) DoSomethingAgain(ctx context.Context) {
	u.repository.DoSomethingElseAgain(ctx)
}

// method with transaction
func (u *Usecase) DoSomething(ctx context.Context) {

	// main context "copied" with transaction
	txCtx, err := u.BeginTx(ctx)
	if err != nil {
		panic(err)
	}
	// check err if you want!:)
	defer u.RollbackTx(txCtx)

	// inside of each method woriking our methods with transacton
	u.repository.DoSomething(txCtx)
	u.repository.DoSomethingElse(txCtx)
	u.repository.DoSomethingAgain(txCtx)
	u.repository.DoSomethingElseAgain(txCtx)
	u.repository.DoSomethingElseAgainAgain(txCtx)

	if err := u.CommitTx(txCtx); err != nil {
		panic(err)
	}
}
