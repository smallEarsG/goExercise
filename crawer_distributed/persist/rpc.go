package persist

import (
	"github.com/olivere/elastic/v7"
	"reptile/engine"
	"reptile/persist"
)

// 具体功能
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	if err == nil {
		*result = "ok"
	}
	return err
}
