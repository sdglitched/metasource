package driver

import (
	"fmt"
	"log/slog"
	"metasource/metasource/models/home"
)

func Database(loca string) error {
	var expt error
	var list []home.LinkUnit
	var item home.LinkUnit

	list, expt = PopulateRepositories()
	if expt != nil {
		return expt
	}

	for _, item = range list {
		expt = HandleRepositories(&item, loca)
		if expt != nil {
			slog.Log(nil, slog.LevelWarn, fmt.Sprintf("[%s] Repository handling failed due to %s", item.Name, expt.Error()))
		}
	}

	return nil
}
