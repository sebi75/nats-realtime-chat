package friends

import "net/http"

type FriendsHandlers struct {
	service FriendsService
}

func (h *FriendsHandlers) SendFriendRequest(w http.ResponseWriter, r *http.Request) {
	// todo
}

func NewFriendsHandlers(service FriendsService) FriendsHandlers {
	return FriendsHandlers{
		service: service,
	}
}
