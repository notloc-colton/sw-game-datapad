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
	baseUrl       string
	cache         tCache[[]Character]
	planetCache   tCache[Planet]
	speciesCache  tCache[Species]
	starshipCache tCache[StarShip]
}

func NewVendorService(baseUrl string) *vendorService {
	return &vendorService{
		baseUrl:       baseUrl,
		cache:         *newTCache[[]Character](),
		planetCache:   *newTCache[Planet](),
		speciesCache:  *newTCache[Species](),
		starshipCache: *newTCache[StarShip](),
	}
}
func (service *vendorService) formUrl(resource string) string {
	return fmt.Sprintf("%s/%s/", service.baseUrl, resource)
}
func (service *vendorService) GetCharacters(queryString string) ([]Character, error) {
	url := service.formUrl("people")
	if cached, found := service.cache.read(queryString); found {
		return cached, nil
	}
	if res, err := httpclient.GetWithSearchString[multipleCharacters](url, queryString); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
		return nil, err
	} else {
		logger.Log(logger.LogLevelInfo, "response", *res)
		go func() {
			service.cache.update(queryString, res.Characters)
		}()
		return res.Characters, nil
	}
}
func (service *vendorService) GetPlanet(id string) (*Planet, error) {
	url := service.formUrl("planets/" + id + "/")
	if cached, found := service.planetCache.read(id); found {
		return &cached, nil
	}
	if res, err := httpclient.Get[Planet](url); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
		return nil, err
	} else {
		go func() {
			service.planetCache.update(id, *res)
		}()
		return res, nil
	}
}
func (service *vendorService) GetSpecies(id string) (*Species, error) {
	url := service.formUrl("species/" + id + "/")
	if cached, found := service.speciesCache.read(id); found {
		return &cached, nil
	}
	if res, err := httpclient.Get[Species](url); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
		return nil, err
	} else {
		go func() {
			service.speciesCache.update(id, *res)
		}()
		return res, nil
	}
}
func (service *vendorService) GetStarShip(id string) (*StarShip, error) {
	url := service.formUrl("starships/" + id + "/")
	if cached, found := service.starshipCache.read(id); found {
		return &cached, nil
	}
	if res, err := httpclient.Get[StarShip](url); err != nil {
		logger.Log(logger.LogLevelError, "error", err.Error())
		return nil, err
	} else {
		go func() {
			service.starshipCache.update(id, *res)
		}()
		return res, nil
	}
}
