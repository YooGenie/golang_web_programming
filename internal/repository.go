package internal

import (
	"errors"
	"github.com/google/uuid"
	customErrors "golang_web_programming/errors"
)

var ErrNotFoundMembership = errors.New("not found membership")

var membershipData = Membership{
	ID:             "3ab365ba-6707-406d-8383-548514e2ecb9",
	UserName:       "jenny",
	MembershipType: "toss",
}

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r Repository) Create(member *Request) (CreateResponse, error) {
	r.data["data"] = membershipData

	if member.UserName == r.data["data"].UserName {
		return CreateResponse{}, customErrors.ApiInternalServerError(customErrors.MessageExistUserName)
	}

	return CreateResponse{uuid.New().String(), member.MembershipType}, nil
}

func (r Repository) Update(member *Request) (UpdateResponse, error) {
	r.data["data"] = membershipData

	if member.UserName == r.data["data"].UserName {
		return UpdateResponse{}, customErrors.ApiInternalServerError(customErrors.MessageExistUserName)
	}

	return UpdateResponse{member.ID, member.UserName, member.MembershipType}, nil
}

func (r Repository) Delete(id string) error {
	data := Membership{ID: "1", UserName: "jenny", MembershipType: "toss"}

	if id != data.ID {
		return customErrors.NoResultError(customErrors.MessageNotExistID)
	}

	return nil
}

func (r Repository) GetOne(id string) (GetResponse, error) {
	data := Membership{ID: "1", UserName: "jenny", MembershipType: "toss"}

	if data.ID != id {
		return GetResponse{}, customErrors.NoResultError(customErrors.MessageNotExistID)
	}
	return GetResponse{data.ID, data.UserName, data.MembershipType}, nil
}

func (r *Repository) GetById(id string) (Membership, error) {
	for _, membership := range r.data {
		if membership.ID == id {
			return membership, nil
		}
	}
	return Membership{}, ErrNotFoundMembership
}