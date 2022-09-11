package internal

import (
	"errors"
	"github.com/google/uuid"
)

var ErrNotFoundMembership = errors.New("not found membership")

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r Repository) Create(member *CreateRequest) (CreateResponse, error) {
	userName := r.data["data"].UserName

	if member.UserName == userName {
		return CreateResponse{}, errors.New("이미 등록된 사용자 이름입니다")
	}

	return CreateResponse{uuid.New().String(), member.MembershipType}, nil
}

func (r Repository) Update(member *UpdateRequest) (UpdateResponse, error) {
	userName := r.data["data"].UserName

	if member.UserName == userName {
		err := errors.New("사용자의 이름이 이미 존재합니다")
		return UpdateResponse{}, err
	}

	return UpdateResponse{member.ID, member.UserName, member.MembershipType}, nil
}

func (r Repository) Delete(id string) error {
	userID := r.data["data"].ID

	if id != userID {
		return errors.New("입력한 id가 존재하지 않습니다")
	}

	return nil
}

func (r Repository) GetOne(id string) (GetResponse, error) {
	data := r.data["data"]
	if data.ID != id {
		return GetResponse{}, errors.New("입력한 id가 존재하지 않습니다")
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
