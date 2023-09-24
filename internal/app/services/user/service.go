package user

import (
	"context"
	"github.com/anisbouzahar/portfolio-api/internal/app/models"
	"github.com/anisbouzahar/portfolio-api/internal/app/repository/user"
	"github.com/anisbouzahar/portfolio-api/pkg/types"
	"time"
)

type ServiceInterface interface {
	SubscribeToBeNotified(bool, subscriberInput *types.SubscriberInput, ctx *context.Context)
}

type Service struct {
	repo user.Repository
}

func (s Service) SubscribeToBeNotified(subscriberInput *types.SubscriberInput, ctx context.Context) (bool, *models.Subscriber, error) {

	subscriber := &models.Subscriber{Email: subscriberInput.Email, SubscribedAt: time.Now()}
	isSubscribed, err := s.repo.IsSubscribed(subscriberInput.Email, ctx)
	if err != nil {
		return false, nil, err
	}
	println(isSubscribed)
	if isSubscribed {
		return isSubscribed, nil, nil
	}

	var dbSubscriber models.Subscriber
	dbSubscriber, err = s.repo.SaveSubscriber(subscriber, ctx)
	if err != nil {
		return false, nil, err
	}
	return false, &dbSubscriber, nil

}

func NewUserService(r user.Repository) *Service {
	return &Service{
		repo: r,
	}
}
