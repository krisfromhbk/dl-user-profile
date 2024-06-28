package selectitems

import (
	"dl-user-profile/internal/generated"
	"encoding/json"
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) SelectItems(w http.ResponseWriter, r *http.Request) {
	result := generated.GetUserInfoSelectItemsResponse{
		Cities: []generated.SlugTitleItem{
			{
				Slug:  "tyumen",
				Title: "Тюмень",
			},
			{
				Slug:  "moscow",
				Title: "Москва",
			},
		},
		DanceExperiences: []generated.SlugTitleItem{
			{
				Slug:  "one_year_and_less",
				Title: "0 - 1 год",
			},
			{
				Slug:  "from_two_to_five_years",
				Title: "2 - 5 лет",
			},
			{
				Slug:  "ten_years_and_more",
				Title: "10 и более лет",
			},
		},
		DanceLevels: []generated.SlugTitleItem{
			{
				Slug:  "beginner",
				Title: "Начинающий",
			},
			{
				Slug:  "advanced",
				Title: "Продвинутый",
			},
		},
		DanceStyles: []generated.SlugTitleItem{},
		Statuses: []generated.SlugTitleItem{
			{
				Slug:  "dancer",
				Title: "Я - танцор",
			},
			{
				Slug:  "teacher",
				Title: "Я - учитель",
			},
		},
	}

	data, _ := json.Marshal(result)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}
