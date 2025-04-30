package usecase

type FetchQuotationUseCase interface {
	Do() error
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

func (uc *fetchQuotationUseCase) Do() error {
	integrationResponse, err := uc.integration.FetchQuotation()
	if err != nil {
		return err
	}

	err = uc.repository.SaveQuotation(integrationResponse)
	if err != nil {
		return err
	}
	
	return nil
}
