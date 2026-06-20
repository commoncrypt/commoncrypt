package patterns

import "context"

type CommandHandler[Cmd any] func(ctx context.Context, cmd Cmd) error
