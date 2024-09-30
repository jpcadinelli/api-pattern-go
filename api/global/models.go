package global

import "api_pattern_go/api/models"

func GetModelsList() []interface{} {
	list := []interface{}{
		&models.Teste{},
	}
	return list
}
