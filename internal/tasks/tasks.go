package tasks

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kamikazechaser/disposable-email-cache/pkg/util"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
)

func DownloadData() {
	log.Info().Msg("tasks: attempting to download data file")

	if err := util.DownloadFile("index.json", "https://raw.githubusercontent.com/ivolo/disposable-email-domains/master/index.json"); err != nil {
		log.Err(err).Msg("tasks: could not download index.json")
	}

	log.Info().Msg("tasks: successfully downloaded data files")
}

func LoadCache() *cache.Cache {
	var indexDomains []string

	indexData, err := ioutil.ReadFile("index.json")
	if err != nil {
		log.Err(err).Msg("tasks: could no read index.json")
		log.Info().Msg("tasks: attempting redownload of data")
		DownloadData()
		LoadCache()
		return nil
	}

	if err := json.Unmarshal(indexData, &indexDomains); err != nil {
		log.Err(err).Msg("tasks: could not unmarshal index data")
	}

	log.Info().Msgf("tasks: read %d domains from data files", len(indexDomains))

	cacheMap := map[string]cache.Item{}
	for _, m := range indexDomains {
		for range m {
			cacheMap[m] = cache.Item{
				Object:     true,
				Expiration: 0,
			}
		}
	}

	log.Info().Msg("tasks: loaded data into in-memory cache")

	return cache.NewFrom(cache.NoExpiration, 0, cacheMap)
}
