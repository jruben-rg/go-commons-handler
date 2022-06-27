package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

func WithCommandDecorator[H any](handler CommandHandler[H], logger *logrus.Entry) CommandHandler[H] {
	return commandLoggingDecorator[H]{
		base:   handler,
		logger: logger,
	}
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
