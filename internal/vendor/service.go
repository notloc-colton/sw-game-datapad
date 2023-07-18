package vendor

import (
	"sw-game-datapad/pkg/httpclient"
	"sw-game-datapad/pkg/logger"
)

type VendorService struct {
}

func NewVendorService() *VendorService {
	return &VendorService{}
}
func (service *VendorService) GetCharacters() {
	if res, err := httpclient.Get[apiResponse]("https://swapi.dev/api/people/", "luke"); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
	} else {
		logger.Log(logger.LogLevelInfo, "response", *res)
	}
}
