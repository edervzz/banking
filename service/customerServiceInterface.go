package service

import (
	"banking/domain"
	"banking/logger"
	"banking/utils"
	"net/http"
)

type CustomerServiceInterface struct {
	repo domain.CustomerRepository
}

func (d *CustomerServiceInterface) CreateCustomer(req *CustomerCreateRequest) (*CustomerCreateResponse, *utils.AppMess) {
	c := &domain.Customer{
		Fullname: req.Fullname,
		City:     req.City,
		Zipcode:  req.Zipcode,
	}

	customerId, err := d.repo.Create(c)
	if err != nil {
		logger.Warn(err.Error())
		e := &utils.AppMess{
			Code:    http.StatusInternalServerError,
			Message: "cannot create customer.",
		}
		return nil, e
	}

	response := CustomerCreateResponse{
		CustomerId: customerId,
	}

	return &response, nil
}

func (d *CustomerServiceInterface) GetCustomer(req *CustomerGetRequest) (*CustomerGetResponse, *utils.AppMess) {
	c, err := d.repo.Get(req.CustomerId)
	if err != nil {
		logger.Warn(err.Error())
		e := &utils.AppMess{
			Code:    http.StatusBadRequest,
			Message: "customer not found",
		}
		return nil, e
	}
	res := CustomerGetResponse{
		CustomerId: c.CustomerId,
		Fullname:   c.Fullname,
		City:       c.City,
		Zipcode:    c.Zipcode,
	}
	return &res, nil
}

// constructor
func NewCustomerServiceInterface(r domain.CustomerRepository) *CustomerServiceInterface {
	return &CustomerServiceInterface{
		repo: r,
	}
}
