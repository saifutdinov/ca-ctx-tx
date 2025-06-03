package domain

import "context"

type (
	Usecase interface {
		DoSomething(ctx context.Context)

		DoSomethingAgain(ctx context.Context)
	}

	Repository interface {
		DoSomething(ctx context.Context)

		DoSomethingAgain(ctx context.Context)

		DoSomethingElse(ctx context.Context)

		DoSomethingElseAgain(ctx context.Context)

		DoSomethingElseAgainAgain(ctx context.Context)
	}
)
