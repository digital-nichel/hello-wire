package handler

import (
	"github.com/google/wire"
	"hello-wire/model"
	"hello-wire/service"
)

type mailService interface {
	GetAll() ([]model.Mail, error)
}

type MailHandler struct {
	service mailService
}

func ProvideMailHandler(service mailService) *MailHandler {
	return &MailHandler{
		service: service,
	}
}

func (handler *MailHandler) GetAll() ([]model.Mail, error) {
	return handler.service.GetAll()
}

var MailHandlerSet = wire.NewSet(
	ProvideMailHandler,
	service.MailServiceSet,
	wire.Bind(new(mailService), new(*service.MailService)),
)

var MailHandlerMockedSet = wire.NewSet(
	ProvideMailHandler,
	service.MailServiceMockedSet,
	wire.Bind(new(mailService), new(*service.MailService)),
)
