package vendor

import (
	"fmt"
	"sw-game-datapad/pkg/httpclient"
	"sw-game-datapad/pkg/logger"
)

type VendorService interface {
	GetCharacters(queryString string) ([]Character, error)
	GetPlanet(id string) (*Planet, error)
	GetSpecies(id string) (*Species, error)
	GetStarShip(id string) (*StarShip, error)
}
type vendorService struct {
	baseUrl string
}

func NewVendorService(baseUrl string) *vendorService {
	return &vendorService{
		baseUrl: baseUrl,
	}
}
func (service *vendorService) formUrl(resource string) string {
	return fmt.Sprintf("%s/%s/", service.baseUrl, resource)
}
func (service *vendorService) GetCharacters(queryString string) ([]Character, error) {
	url := service.formUrl("people")
	if res, err := httpclient.GetWithSearchString[multipleCharacters](url, queryString); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
		return nil, err
	} else {
		logger.Log(logger.LogLevelInfo, "response", *res)
		return res.Characters, nil
	}
}
func (service *vendorService) GetPlanet(id string) (*Planet, error) {
	url := service.formUrl("planets/" + id + "/")
	if res, err := httpclient.Get[Planet](url); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
		return nil, err
	} else {
		logger.Log(logger.LogLevelInfo, "response", *res)
		return res, nil
	}
}
func (service *vendorService) GetSpecies(id string) (*Species, error) {
	url := service.formUrl("species/" + id + "/")
	if res, err := httpclient.Get[Species](url); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
		return nil, err
	} else {
		logger.Log(logger.LogLevelInfo, "response", *res)
		return res, nil
	}
}
func (service *vendorService) GetStarShip(id string) (*StarShip, error) {
	url := service.formUrl("starships/" + id + "/")
	if res, err := httpclient.Get[StarShip](url); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
		return nil, err
	} else {
		logger.Log(logger.LogLevelInfo, "response", *res)
		return res, nil
	}
}
