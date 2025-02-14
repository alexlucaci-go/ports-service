package web

import (
	"context"
	"log"
	"net/http"
)

// middleware is a wrapper function that can will execute some code before or after another Handler
type middleware func(Handler) Handler

func wrapMiddleware(mw []middleware, handler Handler) Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}

	return handler
}

// errorsMid is a middleware that catches any errors that is being returned from the handler
func errorsMid() middleware {
	// This is the actual middleware function to be executed.
	m := func(before Handler) Handler {
		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// catch any errors that is being returned from the handler
			if err := before(ctx, w, r); err != nil {
				log.Printf("ERROR : %v\n", err)

				if err := RespondError(ctx, w, err); err != nil {
					return err
				}
			}

			// The error has been handled so we can stop propagating it.
			return nil
		}

		return h
	}

	return m
}

// panicsMid is a middleware that catches any panics, recovers from them and returns an error to the client
func panicsMid() middleware {
	// not implemented
	return nil
}

// loggerMid is a middleware that logs information about the request
func loggerMid() middleware {
	// not implemented
	return nil
}

// authMid is a middleware that checks if the user is authenticated
func authMid() middleware {
	// not implemented
	return nil
}
