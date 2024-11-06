package problemdetail

import (
	problemDetails "github.com/duongbui2002/core-package/core/http/httperrors/problemdetails"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ProblemDetail returns an Echo middleware function that processes errors and converts them to problem details.
func ProblemDetail(opts ...Option) echo.MiddlewareFunc {
	cfg := config{}
	// Apply all provided options to the config
	for _, opt := range opts {
		opt.apply(&cfg)
	}

	// Set the default skipper if none is provided
	if cfg.Skipper == nil {
		cfg.Skipper = middleware.DefaultSkipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Skip middleware if the skipper function returns true
			if cfg.Skipper(c) {
				return next(c)
			}

			// Call the next handler in the chain
			err := next(c)

			// Parse the error into a problem detail
			prbError := problemDetails.ParseError(err)

			// Use custom problem parser if provided
			if cfg.ProblemParser != nil {
				prbError = cfg.ProblemParser(prbError)
			}

			// If a problem detail error exists, handle it
			if prbError != nil {
				// Handle echo error in this middleware and raise echo error handler func and our custom error handler
				// When we call c.Error more than once, `c.Response().Committed` becomes true and response doesn't write to client again in our error handler
				// Error will update response status with occurred error object status code
				c.Error(prbError)
			}

			// Return the problem detail error
			return prbError
		}
	}
}
