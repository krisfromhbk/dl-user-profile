package getuserinfo

import (
	"dl-user-profile/internal/generated"
	"encoding/json"
	"net/http"

	"github.com/samber/lo"
)

type Handler struct {
	s storage
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	result := generated.GetUserInfoResponse{
		Age:             "24",
		City:            "Moscow",
		DanceExperience: "146%",
		DanceStyle:      "Kebab",
		Media: []generated.MediaItem{
			{
				Id:         2579,
				PreviewUrl: "https://github.com/",
				Type:       "video",
				Url:        "https://github.com/",
			},
			{
				Id:         4623,
				PreviewUrl: "https://github.com/",
				Type:       "photo",
				Url:        "https://github.com/",
			},
		},
		Name:              "Johny Seens",
		Note:              lo.ToPtr("Prepare some well-suited cameras"),
		ProfilePictureUrl: "https://github.com/",
	}

	data, _ := json.Marshal(result)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}
