package dto

import (
	"Status418/go/enums"
	"Status418/go/models"
)

type FiltersDto struct {
	Aproximation string `json:"filter_aproximation" validate:"max=100"`
	Moment       string `json:"filter_moment"`
	Type         string `json:"filter_type"`
	All          bool   `json:"filter_all" validate:"boolean"`
}

func (dto FiltersDto) GetModel() models.Filter {
	return models.Filter{
		Aproximation: dto.Aproximation,
		Moment:       enums.GetMomentEnum(dto.Moment),
		Type:         enums.GetTypeEnum(dto.Type),
		All: 		  dto.All,
	}
}
