package dto

import "api/errs"

type SendFriendRequestRequest struct {
	AddresseeId string `json:"addressee_id"`
}

type RespondToFriendRequestRequest struct {
	RequesterId string `json:"requester_id"`
	Status      string `json:"status"`
}

func (rtfr *RespondToFriendRequestRequest) Validate() *errs.AppError {
	if rtfr.Status != "ACCEPTED" && rtfr.Status != "REJECTED" {
		return errs.NewBadRequestError("Invalid status")
	}
	return nil
}
