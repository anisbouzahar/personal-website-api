package types

import (
	"github.com/go-chi/render"
	"net/http"
)

// SubscriberInput is an instance of a subscriber.
type SubscriberInput struct {
	Email string `json:"email" example:"email@example.com" validate:"required,email" error:"The email field is required and must be a valid email address"`
} // @name Subscriber

// SubscriberOutput is an instance of a subscriber.
type SubscriberOutput struct {
	Email        string `json:"email"`
	IsSubscribed bool   `json:"is_subscribed" `
} // @name Subscriber

// Render implements the github.com/go-chi/render.Renderer interface
func (s *SubscriberInput) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind implements the github.com/go-chi/render.Binder interface
func (s *SubscriberInput) Bind(r *http.Request) error {
	return nil
}

// ErrResponse renderer type for handling all sorts of errors.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status" example:"Resource not found."`                                         // user-level status message
	AppCode    int64  `json:"code,omitempty" example:"404"`                                                 // application-specific error code
	ErrorText  string `json:"error,omitempty" example:"The requested resource was not found on the server"` // application-level error message, for debugging
} // @name ErrorResponse

// Render implements the github.com/go-chi/render.Renderer interface for ErrResponse
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest returns a structured http response for invalid requests
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

// ErrRender returns a structured http response in case of rendering errors
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

// ErrNotFound returns a structured http response if a resource couln't be found
func ErrNotFound() render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     "Resource not found.",
	}
}
