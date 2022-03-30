package service

import (
	"banking/domain"
	"banking/utils"
	"time"
)

type AccountServiceInterface struct {
	repo domain.AccountRepository
}

func (s AccountServiceInterface) CreateAccount(req *CreateAccountRequest) (*CreateAccountResponse, *utils.AppMess) {
	t := time.Now()
	st := t.Format("2006-02-01 15:04:05")
	a := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: st,
		AccountType: req.AccountType,
		Balance:     req.Balance,
		Status:      domain.ACCT_ACTIVE,
	}

	if a.Balance <= 0 {
		return nil, utils.NewBadRequest("balance should be more than zero")
	}

	accountId, err := s.repo.Create(&a)
	if err != nil {
		return nil, utils.NewBadRequest(err.Error())
	}
	r := CreateAccountResponse{
		AccountId: accountId,
	}
	return &r, nil
}

func (s AccountServiceInterface) GetBalance(req *GetBalanceRequest) (*GetBalanceResponse, *utils.AppMess) {
	balance, err := s.repo.GetBalance(req.AccountId)
	if err != nil {
		return nil, utils.NewNotFound(err.Error())
	}
	r := GetBalanceResponse{
		Balance: balance,
	}
	return &r, nil
}

func (s AccountServiceInterface) LockAccount(req *LockAccountRequest) (*LockAccountResponse, *utils.AppMess) {
	err := s.repo.Lock(req.AccountId)
	if err != nil {
		return nil, utils.NewBadRequest(err.Error())
	}
	return nil, nil
}

func (s AccountServiceInterface) UnlockAccount(req *UnlockAccountRequest) (*UnlockAccountResponse, *utils.AppMess) {
	err := s.repo.Unlock(req.AccountId)
	if err != nil {
		return nil, utils.NewBadRequest(err.Error())
	}
	return nil, nil
}

// contructor
func NewAccountServiceInterface(repo domain.AccountRepository) *AccountServiceInterface {
	s := AccountServiceInterface{
		repo,
	}
	return &s
}
