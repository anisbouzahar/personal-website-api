package user

import (
	"github.com/anisbouzahar/portfolio-api/internal/app/services/user"
	"github.com/anisbouzahar/portfolio-api/pkg/types"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type Handler struct {
	logger      *logrus.Logger
	router      *mux.Router
	userService *user.Service
	validate    *validator.Validate
}

func NewHandler(userS *user.Service) *Handler {
	return &Handler{
		userService: userS,
		validate:    validator.New(),
	}
}

// SubscribeTobeNotified writes a new subscriber to the database
// @Summary add a subscriber to the database
// @Description SubscribeTobeNotified writes a subscriber to the database
// @Description To write a new subscriber,
// @Tags Subscriber
// @Produce json
// @Router /subscriber [post]
// @Success 200 {object} types.Subscriber
// @Failure 400 {object} types.ErrResponse
// @Failure 404 {object} types.ErrResponse
func (h *Handler) SubscribeTobeNotified(w http.ResponseWriter, r *http.Request) {

	subscriberInput := &types.SubscriberInput{}

	err := render.Bind(r, subscriberInput)
	if err := h.validate.Struct(subscriberInput); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make([]string, 0, len(validationErrors))
		for _, e := range validationErrors {
			if e.Tag() == "email" {
				errors = append(errors, "invalid email address")
			} else if e.Tag() == "required" {
				errors = append(errors, e.Field()+"field is missing")
			}
		}
		render.Render(w, r, &types.ErrResponse{
			HTTPStatusCode: 400,
			StatusText:     "Bad Request",
			ErrorText:      strings.Join(errors, ", "),
		})
		return
	}

	isSubscribed, subscriber, err := h.userService.SubscribeToBeNotified(subscriberInput, r.Context())
	if err != nil {
		err := render.Render(w, r, types.ErrInvalidRequest(err))
		if err != nil {
			h.logger.Error("an error has occurred %v", err)
		}
		return
	}

	var email string
	if isSubscribed {
		email = subscriberInput.Email
	} else {
		email = subscriber.Email
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, &types.SubscriberOutput{
		Email:        email,
		IsSubscribed: isSubscribed,
	})
}
