package handlers

import (
	"emperror.dev/errors"
	problemDetails "github.com/duongbuidinh600/core-package/core/http/httperrors/problemdetails"
	"github.com/duongbuidinh600/core-package/core/logger"
	"github.com/labstack/echo/v4"
)

// ProblemDetailErrorHandlerFunc is a custom error handler function for Echo framework.
// It converts the given error to a ProblemDetailErr if it is not already one,
// and writes the problem details to the response if the response has not been committed yet.
func ProblemDetailErrorHandlerFunc(
	err error,
	c echo.Context,
	logger logger.Logger,
) {
	var problem problemDetails.ProblemDetailErr

	// If the error is not of type ProblemDetailErr, convert it to ProblemDetailErr
	if ok := errors.As(err, &problem); !ok {
		problem = problemDetails.ParseError(err)
	}

	// If the response has not been committed and the problem is not nil,
	// write the problem details to the response
	if !c.Response().Committed && problem != nil {
		// WriteTo will set the response status code to the problem details status
		if _, err := problemDetails.WriteTo(problem, c.Response()); err != nil {
			logger.Error(err)
		}
	}
}
