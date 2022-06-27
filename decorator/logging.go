package decorator

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type commandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger *logrus.Entry
}

type queryLoggingDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	logger *logrus.Entry
}

func (cl commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	handlerType := generateActionName(cmd)

	logger := cl.logger.WithFields(logrus.Fields{
		"command":      handlerType,
		"command_body": fmt.Sprintf("%#v", cmd),
	})

	logger.Debug("Executing command")

	defer func() {
		if err == nil {
			logger.Info("Command executed ok")
		} else {
			logger.WithError(err).Error("Failed to execute command")
		}
	}()

	return cl.base.Handle(ctx, cmd)
}

func (ql queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {

	logger := ql.logger.WithFields(logrus.Fields{
		"query":      generateActionName(cmd),
		"query_body": fmt.Sprintf("%#v", cmd),
	})

	logger.Debug("Executing query")
	defer func() {
		if err == nil {
			logger.Info("Query executed ok")
		} else {
			logger.WithError(err).Error("Failed to execute query")
		}
	}()

	return ql.base.Handle(ctx, cmd)
}
