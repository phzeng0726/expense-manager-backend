package services

import (
	"errors"
	"expense-manager-backend/constants"
	"fmt"
)

func constructURL(language string, newsId *string) (string, error) {
	var paths map[string]string

	domains := map[string]string{
		"en_US": constants.EnDomain,
		"zh_TW": constants.ZhDomain,
	}

	// newsList
	if newsId == nil {
		paths = map[string]string{
			"en_US": "tag/FIN/1",
			"zh_TW": "news/list.php?nt_pk=7",
		}
	} else {
		paths = map[string]string{
			"en_US": "news/",
			"zh_TW": "news/detail.php?",
		}
	}

	if language == "en_US" || language == "zh_TW" {
		domain := domains[language]
		path := paths[language]
		if newsId != nil {
			path += *newsId
		}
		return fmt.Sprintf("%s/%s", domain, path), nil
	}

	return "", errors.New("language not available")
}
