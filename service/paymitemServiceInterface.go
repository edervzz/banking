package service

import (
	"banking/domain"
	"banking/utils"
	"math"
	"time"
)

type PaymitemServiceInterface struct {
	repo domain.PaymItemRepository
}

func (s PaymitemServiceInterface) Create(req *PaymitemCreateRequest) (*PaymitemCreateResponse, *utils.AppMess) {
	var appMess *utils.AppMess

	req.TAmount = -math.Abs(req.TAmount)

	paymitem := domain.PaymItem{
		AccountId: req.AccountId,
		TAmount:   req.TAmount,
		TransType: req.TransType,
		Concept:   req.Concept,
		Status:    domain.PAYM_CREATED,
		DatePost:  time.Now().Local(),
		DateValue: time.Now().Local(),
	}
	docId, err := s.repo.Create(paymitem)
	if err != nil {
		appMess = utils.NewBadRequest(err.Error())
		return nil, appMess
	}
	return &PaymitemCreateResponse{
			DocumentId: docId,
		},
		nil
}

func NewPaymitemServiceInterface(repo domain.PaymItemRepository) *PaymitemServiceInterface {
	return &PaymitemServiceInterface{
		repo,
	}
}
