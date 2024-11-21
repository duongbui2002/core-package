package problemDetails

import (
	"context"
	"database/sql"
	"emperror.dev/errors"
	"github.com/duongbui2002/core-package/constants"
	customErrors2 "github.com/duongbui2002/core-package/http/httperrors/customerrors"
	typeMapper "github.com/duongbui2002/core-package/reflection/typemapper"
	errorUtils "github.com/duongbui2002/core-package/utils/errorutils"
	"github.com/go-playground/validator"
	"net/http"
	"reflect"
)

type ProblemDetailParser struct {
	internalErrors map[reflect.Type]func(err error) ProblemDetailErr
}

type ErrorParserFunc func(err error) ProblemDetailErr

func NewProblemDetailParser(
	builder func(builder *OptionBuilder),
) *ProblemDetailParser {
	optionBuilder := NewOptionBuilder()
	builder(optionBuilder)
	items := optionBuilder.Build()
	return &ProblemDetailParser{internalErrors: items}
}

func (p *ProblemDetailParser) ResolveError(err error) ProblemDetailErr {
	errType := typeMapper.GetReflectType(err)
	problem := p.internalErrors[errType]
	if problem != nil {
		return problem(err)
	}
	return nil
}

func ParseError(err error) ProblemDetailErr {
	stackTrace := errorUtils.ErrorsWithStack(err)
	customErr := customErrors2.GetCustomError(err)
	var validatorErr validator.ValidationErrors

	if err != nil && customErr != nil {
		switch {
		case customErrors2.IsDomainError(err, customErr.Status()):
			return NewDomainProblemDetail(
				customErr.Status(),
				customErr.Error(),
				stackTrace,
			)
		case customErrors2.IsApplicationError(err, customErr.Status()):
			return NewApplicationProblemDetail(
				customErr.Status(),
				customErr.Error(),
				stackTrace,
			)
		case customErrors2.IsApiError(err, customErr.Status()):
			return NewApiProblemDetail(
				customErr.Status(),
				customErr.Error(),
				stackTrace,
			)
		case customErrors2.IsBadRequestError(err):
			return NewBadRequestProblemDetail(customErr.Error(), stackTrace)
		case customErrors2.IsNotFoundError(err):
			return NewNotFoundErrorProblemDetail(customErr.Error(), stackTrace)
		case customErrors2.IsValidationError(err):
			return NewValidationProblemDetail(customErr.Error(), stackTrace)
		case customErrors2.IsUnAuthorizedError(err):
			return NewUnAuthorizedErrorProblemDetail(
				customErr.Error(),
				stackTrace,
			)
		case customErrors2.IsForbiddenError(err):
			return NewForbiddenProblemDetail(customErr.Error(), stackTrace)
		case customErrors2.IsConflictError(err):
			return NewConflictProblemDetail(customErr.Error(), stackTrace)
		case customErrors2.IsInternalServerError(err):
			return NewInternalServerProblemDetail(customErr.Error(), stackTrace)
		case customErrors2.IsCustomError(err):
			return NewProblemDetailFromCodeAndDetail(
				customErr.Status(),
				customErr.Error(),
				stackTrace,
			)
		case customErrors2.IsUnMarshalingError(err):
			return NewInternalServerProblemDetail(err.Error(), stackTrace)
		case customErrors2.IsMarshalingError(err):
			return NewInternalServerProblemDetail(err.Error(), stackTrace)

		default:
			return NewInternalServerProblemDetail(err.Error(), stackTrace)
		}
	} else if err != nil && customErr == nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return NewNotFoundErrorProblemDetail(err.Error(), stackTrace)
		case errors.Is(err, context.DeadlineExceeded):
			return NewProblemDetail(
				http.StatusRequestTimeout,
				constants.ErrRequestTimeoutTitle,
				err.Error(),
				stackTrace,
			)
		case errors.As(err, &validatorErr):
			return NewValidationProblemDetail(validatorErr.Error(), stackTrace)
		default:
			return NewInternalServerProblemDetail(err.Error(), stackTrace)
		}
	}

	return nil
}
