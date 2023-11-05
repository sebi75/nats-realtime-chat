package friends

import "api/errs"

type FriendsService struct {
	repo FriendsRepository
}

func (s *FriendsService) SendFriendRequest(requesterId string, addresseeId string) (int, *errs.AppError) {
	// todo
	return 0, nil
}

func (s *FriendsService) RespondToFriendRequest(requesterId string, addresseeId string, status string) (int, *errs.AppError) {
	// todo
	return 0, nil
}

func NewFriendsService(repo FriendsRepository) FriendsService {
	return FriendsService{
		repo: repo,
	}
}
