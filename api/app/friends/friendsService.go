package friends

import (
	"api/app/friends/domain"
	"api/errs"
	"api/utils"
	"api/utils/logger"
	"net/http"
)

type FriendsService struct {
	repo *FriendsRepository
}

func (s *FriendsService) SendFriendRequest(requesterId string, addresseeId string) (*domain.Friend, *errs.AppError) {
	// first check if the requestee and addressee are already friends or there already exists a pending request
	friend, err := s.repo.FindFriendByRequesteeAndAddresseeId(requesterId, addresseeId, "")
	if err != nil && err.Code != http.StatusNotFound {
		return nil, err
	}
	if friend.Status == "PENDING" {
		return nil, errs.NewBadRequestError("Friend request already pending")
	}
	if friend.Status == "ACCEPTED" {
		return nil, errs.NewBadRequestError("Users are already friends")
	}
	// if there isn't any pending request or users are not already friends, create a new friend request
	uuid, uuidErr := utils.GenerateUUID()
	if uuidErr != nil {
		logger.Error("Error while generating uuid: " + err.Message)
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	newFriend := domain.Friend{
		Id:          uuid,
		RequesterId: requesterId,
		AddresseeId: addresseeId,
		// default status is pending at db level
	}
	friend, err = s.repo.Save(&newFriend)
	if err != nil {
		logger.Error("Error while saving friend request: " + err.Message)
		return nil, err
	}

	// send notification to addressee using NATS message

	return friend, nil
}

func (s *FriendsService) RespondToFriendRequest(requesterId string, addresseeId string, status string) (int, *errs.AppError) {
	// todo
	return 0, nil
}

// add pagination and filtering
func (s *FriendsService) FindAllFriends(userId string) (*FindFriendsResponse, *errs.AppError) {
	friendsResponse, err := s.repo.FindAllFriendsByUserId(userId)
	if err != nil {
		return nil, err
	}

	return friendsResponse, nil
}

func NewFriendsService(repo *FriendsRepository) *FriendsService {
	return &FriendsService{
		repo: repo,
	}
}
