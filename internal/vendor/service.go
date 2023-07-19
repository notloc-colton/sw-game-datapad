package vendor

import (
	"fmt"
	"sw-game-datapad/pkg/httpclient"
	"sw-game-datapad/pkg/logger"
)

type VendorService struct {
	baseUrl string
}

func NewVendorService(baseUrl string) *VendorService {
	return &VendorService{
		baseUrl: baseUrl,
	}
}
func (service *VendorService) formUrl(resource string) string {
	return fmt.Sprintf("%s/%s/", service.baseUrl, resource)
}
func (service *VendorService) GetCharacters(queryString string) ([]Character, error) {
	url := service.formUrl("people")
	if res, err := httpclient.GetWithSearchString[apiResponse](url, queryString); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
		return nil, err
	} else {
		logger.Log(logger.LogLevelInfo, "response", *res)
		return res.Characters, nil
	}
}
func (service *VendorService) GetPlanet(id string) (*Planet, error) {
	url := service.formUrl("planets/" + id + "/")
	if res, err := httpclient.Get[Planet](url); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
		return nil, err
	} else {
		logger.Log(logger.LogLevelInfo, "response", *res)
		return res, nil
	}
}
func (service *VendorService) GetSpecies(id string) (*Species, error) {
	url := service.formUrl("species/" + id + "/")
	if res, err := httpclient.Get[Species](url); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
		return nil, err
	} else {
		logger.Log(logger.LogLevelInfo, "response", *res)
		return res, nil
	}
}
func (service *VendorService) GetStarShip(id string) (*StarShip, error) {
	url := service.formUrl("starships/" + id + "/")
	if res, err := httpclient.Get[StarShip](url); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
		return nil, err
	} else {
		logger.Log(logger.LogLevelInfo, "response", *res)
		return res, nil
	}
}
