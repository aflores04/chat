package handler

import (
	"errors"
	"fmt"
	"github.com/aflores04/chat/backend/src/chat_messages/request"
	"github.com/aflores04/chat/backend/src/utils"
	"net/http"
	"strconv"
)

func (h chatMessagesHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	req, err := parseQueryParams(r)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.service.GetMessages(r.Context(), req)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	utils.HttpResponse(w, http.StatusOK, resp.Messages)
}

func parseQueryParams(r *http.Request) (*request.ListMessagesRequest, error) {
	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	sortOrder, err := strconv.Atoi(r.URL.Query().Get("sort_order"))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot parse amount or sort_order: %s", err.Error()))
	}
	sortKey := r.URL.Query().Get("sort_key")

	return &request.ListMessagesRequest{
		Amount:    int64(amount),
		SortOrder: int64(sortOrder),
		SortKey:   sortKey,
	}, nil
}
