package usecase

import "challeng-client-server-api/internal/entities"

type FetchAndSaveQuotationUseCase interface {
	Do() (*entities.Quotation, error)
}

type fetchAndSaveQuotationUseCase struct {
	repository  QuotationRepository
	integration EconomiaIntegration
}

func NewFetchAndSaveQuotationUseCase(repository QuotationRepository, integration EconomiaIntegration) FetchAndSaveQuotationUseCase {
	return &fetchAndSaveQuotationUseCase{
		repository:  repository,
		integration: integration,
	}
}

func (uc *fetchAndSaveQuotationUseCase) Do() (*entities.Quotation, error) {
	integrationResponse, err := uc.integration.FetchQuotation()
	if err != nil {
		return nil, err
	}

	err = uc.repository.SaveQuotation(integrationResponse)
	if err != nil {
		return nil, err
	}

	return integrationResponse, nil
}
