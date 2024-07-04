package service

import (
	"database/sql"
	"hexagonal/errs"
	"hexagonal/logs"
	"hexagonal/repository"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {

	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custResponses := []CustomerResponse{}

	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}

		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
