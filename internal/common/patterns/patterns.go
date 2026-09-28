package patterns

import "context"

// CommandHandler handles a command whose effect is a pure side effect; it
// reports only whether the effect succeeded.
type CommandHandler[Cmd any] interface {
	Handle(ctx context.Context, cmd Cmd) error
}

// CommandProducer handles a command whose side effect produces a result
// intrinsic to that effect (e.g. a generated ID or token). It is still a
// command: the result describes what the mutation did, not an independent
// read of unrelated state.
type CommandProducer[Cmd, Result any] interface {
	Produce(ctx context.Context, cmd Cmd) (Result, error)
}

// QueryProducer handles a query, returning a result without producing any
// side effects.
type QueryProducer[Query, Result any] interface {
	Produce(ctx context.Context, query Query) (Result, error)
}
