package usecase

import "challeng-client-server-api/server/entities"

type FetchQuotationUseCase interface {
	Do() (*entities.Quotation, error)
}

type fetchQuotationUseCase struct {
	repository  QuotationRepository
	integration QuotationIntegration
}

func NewFetchQuotationUseCase(repository QuotationRepository, integration QuotationIntegration) FetchQuotationUseCase {
	return &fetchQuotationUseCase{
		repository:  repository,
		integration: integration,
	}
}

func (uc *fetchQuotationUseCase) Do() (*entities.Quotation, error) {
	integrationResponse, err := uc.integration.GetQuotation()
	if err != nil {
		return nil, err
	}

	return integrationResponse, nil
}
