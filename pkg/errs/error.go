package errr

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/kurdilesmana/backend-product-service/pkg/constants"
	"github.com/kurdilesmana/backend-product-service/pkg/logger"
)

func New(ctx context.Context, errType string, message string, err error) error {
	var errFormat error

	if err != nil {
		if err.Error() != "" {
			errFormat = fmt.Errorf("%s | %s: %w", errType, message, err)
		}

	} else {
		errFormat = fmt.Errorf("%s | %s", errType, message)
	}

	logger.LogError(ctx, constants.Err, errType, message)
	return errFormat
}

func TrimErrorMessage(err error) (errType, errMessage string, newErr error) {
	errs := strings.Split(err.Error(), "|")
	errType = strings.TrimSpace(errs[0])
	errMessage = strings.TrimSpace(errs[1])

	newErr = errors.New(strings.TrimSpace(errs[1]))
	if len(errs)-1 == 2 {
		newErr = errors.New(strings.TrimSpace(errs[2]))
	} else if len(errs) > 1 {
		newErr = errors.New(strings.TrimSpace(errs[1]))
	}

	return
}
