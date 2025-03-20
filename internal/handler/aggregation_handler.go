package handler

import (
	"log/slog"
	"net/http"

	"ayzhunis/hot-coffee/internal/service"
	"ayzhunis/hot-coffee/models"
	"ayzhunis/hot-coffee/utils"
)

type AggregationHandler struct {
	aggregationService *service.AggregationService
}

func NewAggregationHandler(aggregationService *service.AggregationService) *AggregationHandler {
	return &AggregationHandler{aggregationService: aggregationService}
}

func (a *AggregationHandler) TotalSales(w http.ResponseWriter, r *http.Request) {
	total, err := a.aggregationService.TotalSales()
	if err != nil {
		slog.Error(err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("total salse", utils.ReqGroup())
	utils.RespondWithJSON(w, http.StatusCreated, map[string]float64{"total sales": total})
}

func (a *AggregationHandler) PopularItems(w http.ResponseWriter, r *http.Request) {
	item, err := a.aggregationService.PopularItems()
	if err != nil {
		slog.Error(err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("popular items", utils.ReqGroup())
	utils.RespondWithJSON(w, http.StatusCreated, map[string]models.MenuItem{"popular item": *item})
}
