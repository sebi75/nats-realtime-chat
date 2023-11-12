package friends

import (
	"api/app/auth"
	"api/app/friends/dto"
	"api/utils"
	"encoding/json"
	"net/http"
)

type FriendsHandlers struct {
	service     *FriendsService
	authService *auth.AuthService
}

func (h *FriendsHandlers) SendFriendRequest(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")[7:]
	verifyResponse, appErr := h.authService.Verify(token)
	if appErr != nil {
		utils.ResponseWriter(w, http.StatusUnauthorized, appErr.Message)
		return
	}
	var request dto.SendFriendRequestRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.ResponseWriter(w, http.StatusBadRequest, err.Error())
		return
	}

	friend, appErr := h.service.SendFriendRequest(verifyResponse.Id, request.AddresseeId)
	if appErr != nil {
		utils.ResponseWriter(w, appErr.Code, appErr.Message)
		return
	}

	utils.ResponseWriter(w, http.StatusCreated, friend)
	return
}

func (h *FriendsHandlers) FindFriends(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")[7:]
	verifyResponse, err := h.authService.Verify(token)
	if err != nil {
		utils.ResponseWriter(w, http.StatusUnauthorized, err.Message)
		return
	}

	friendsResponse, err := h.service.FindAllFriends(verifyResponse.Id)
	if err != nil {
		utils.ResponseWriter(w, err.Code, err.Message)
		return
	}

	utils.ResponseWriter(w, http.StatusOK, friendsResponse)
	return
}

func NewFriendsHandlers(service *FriendsService, authService *auth.AuthService) *FriendsHandlers {
	return &FriendsHandlers{
		service:     service,
		authService: authService,
	}
}
