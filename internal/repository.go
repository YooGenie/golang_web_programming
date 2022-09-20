package internal

import (
	customErrors "golang_web_programming/errors"
)

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r Repository) Create(member Membership) (*CreateResponse, error) {

	for _, v := range r.data {
		if v.UserName == member.UserName {
			return nil, customErrors.ErrExistUserName
		}
	}

	r.data[member.ID] = member

	return &CreateResponse{r.data[member.ID].ID, r.data[member.ID].MembershipType}, nil
}

func (r Repository) Update(member Membership) (*UpdateResponse, error) {
	for _, v := range r.data {
		if v.UserName == member.UserName {
			return nil, customErrors.ErrExistUserName
		}
	}

	r.data[member.ID] = member

	return &UpdateResponse{member.ID, member.UserName, member.MembershipType}, nil
}

func (r Repository) Delete(id string) error {
	r.data["3ab365ba-6707-406d-8383-548514e2ecb9"] = Membership{ID: "3ab365ba-6707-406d-8383-548514e2ecb9", UserName: "jenny", MembershipType: "toss"}

	if id != r.data[id].ID {
		return customErrors.ErrNotExistID
	}

	return nil
}

func (r *Repository) GetById(id string) (Membership, error) {
	r.data["3ab365ba-6707-406d-8383-548514e2ecb9"] = Membership{ID: "3ab365ba-6707-406d-8383-548514e2ecb9", UserName: "jenny", MembershipType: "toss"}

	if r.data[id].ID != id {
		return Membership{}, customErrors.ErrNotExistID
	}
	return r.data[id], nil
}
