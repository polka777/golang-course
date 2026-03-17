package handlers

import (
	"encoding/json"
	"gateway/internal/usecase"
	"net/http"
	"strings"
)

type RepoHandler struct {
	UseCase usecase.GatewayUsecase
}

func NewRepoHandler(usecase usecase.GatewayUsecase) RepoHandler {
	return RepoHandler{UseCase: usecase}

}

// @Summary Получить информацию о репозитории
// @Description Возвращает данные о репозитории: название, описание, звёзды, форки и дату создания
// @Tags repositories
// @Accept json
// @Produce json
// @Param owner query string true "Владелец репозитория (например, golang)"
// @Param repo query string true "Название репозитория (например, go)"
// @Success 200 {object} domain.Repository "Информация о репозитории"
// @Failure 400 {object} map[string]string "Ошибка: не указаны owner или repo"
// @Failure 404 {object} map[string]string "Репозиторий не найден"
// @Failure 503 {object} map[string]string "Collector service недоступен"
// @Router /api/repo [get]
func (h RepoHandler) GetRepositoryInfo(w http.ResponseWriter, r *http.Request) {
	owner := r.URL.Query().Get("owner")
	name := r.URL.Query().Get("repo")
	if name == "" || owner == "" {
		respondWithError(w, http.StatusBadRequest, "empty fields owner or name")
		return
	}
	repo, err := h.UseCase.GetRepositoryInfo(r.Context(), owner, name)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "not found"):
			respondWithError(w, http.StatusNotFound, "repository not found")
		case strings.Contains(err.Error(), "unavaliable"):
			respondWithError(w, http.StatusServiceUnavailable, "collector service unavaliable")
		case strings.Contains(err.Error(), "timout"):
			respondWithError(w, http.StatusGatewayTimeout, "request timout")
		default:
			respondWithError(w, http.StatusInternalServerError, "server error")

		}
		return
	}
	respondWithJSON(w, http.StatusOK, repo)
	return

}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
