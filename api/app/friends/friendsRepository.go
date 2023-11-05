package friends

import (
	"api/app/friends/domain"
	"api/errs"
	"api/utils/logger"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type FriendsRepository struct {
	client *sqlx.DB
}

func (r *FriendsRepository) FindFriendById(id string) (*domain.Friend, *errs.AppError) {
	var friend domain.Friend
	getFriendSql := `SELECT * FROM Friend WHERE id = ?`
	err := r.client.Get(&friend, getFriendSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Friend not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &friend, nil
}

func (r *FriendsRepository) FindFriendByRequesteeAndAddresseeId(requesteeId, addresseeId, status string) (*domain.Friend, *errs.AppError) {
	var friend domain.Friend
	getFriendSql := `SELECT * FROM Friend WHERE requestee_id = ? AND addressee_id = ?`
	// if status is defined, add it to the query
	if status != "" {
		getFriendSql += ` AND status = ?`
	}
	err := r.client.Get(&friend, getFriendSql, requesteeId, addresseeId, status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Friend not found")
		} else {
			logger.Error("Error while getting friend: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &friend, nil
}

func (r *FriendsRepository) FindAllFriendsByUserId(userId string) ([]domain.FriendWithUser, *errs.AppError) {
	var friends []domain.FriendWithUser
	getFriendsSql := `SELECT
						CASE
							WHEN F.requester_id = ? THEN U1
							ELSE U2
						END AS user, F
					FROM
						Friend F
					JOIN
						User U1 ON U1.id = F.requester_id
					JOIN
						User U2 ON U2.id = F.addressee_id
					WHERE
						(F.requester_id = ? OR F.addressee_id = ?)
					AND
						status = "ACCEPTED"
					`
	err := r.client.Select(&friends, getFriendsSql, userId, userId, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Friends not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return friends, nil
}

func (r *FriendsRepository) Save(friend *domain.Friend) (*domain.Friend, *errs.AppError) {
	insertFriendSql := `INSERT INTO Friend (id, requester_id, addressee_id) VALUES (?, ?, ?)`
	_, err := r.client.Exec(insertFriendSql, friend.Id, friend.RequesterId, friend.AddresseeId)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return friend, nil
}

func (r *FriendsRepository) Update(friend *domain.Friend) (*domain.Friend, *errs.AppError) {
	updateFriendSql := `UPDATE Friend SET status = ? WHERE id = ?`
	_, err := r.client.Exec(updateFriendSql, friend.Status, friend.Id)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return friend, nil
}

func New(client *sqlx.DB) *FriendsRepository {
	return &FriendsRepository{
		client: client,
	}
}
