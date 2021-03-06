package funcExtractResults

import (
	"fmt"
	"github.com/sp0x/ihcph/sites"
	"github.com/sp0x/torrentd/indexer"
	"github.com/sp0x/torrentd/indexer/definitions"
)

func getEmbeddedDefinitionsSource() indexer.DefinitionLoader {
	x := indexer.CreateEmbeddedDefinitionSource(sites.GzipAssetNames(), func(key string) ([]byte, error) {
		fullname := fmt.Sprintf("sites/%s.yml", key)
		data, err := sites.GzipAsset(fullname)
		if err != nil {
			return nil, err
		}
		data, _ = definitions.UnzipData(data)
		return data, nil
	})
	return x
}

func GetIndexLoader(appName string) *indexer.MultipleDefinitionLoader {
	m := &indexer.MultipleDefinitionLoader{
		getEmbeddedDefinitionsSource(),
		indexer.NewFsLoader(appName),
	}
	return m
}
