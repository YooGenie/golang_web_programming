package member

import (
	customErrors "golang_web_programming/errors"
)

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r Repository) Create(membership Membership) (*CreateResponse, error) {
	//r.data[membership.ID] = membership

	for _, v := range r.data {
		if v.UserName == membership.UserName {
			return nil, customErrors.ErrExistUserName
		}
	}

	r.data[membership.ID] = membership

	return &CreateResponse{r.data[membership.ID].ID, r.data[membership.ID].UserName, r.data[membership.ID].MembershipType}, nil
}

func (r Repository) Update(membership Membership) (*UpdateResponse, error) {
	//r.data[membership.ID] = membership

	for _, v := range r.data {
		if v.UserName == membership.UserName {
			return nil, customErrors.ErrExistUserName
		}
	}

	r.data[membership.ID] = membership

	return &UpdateResponse{membership.ID, membership.UserName, membership.MembershipType}, nil
}

func (r Repository) Delete(id string) error {
	if id != r.data[id].ID {
		return customErrors.ErrNotExistID
	}

	return nil
}

func (r *Repository) GetById(id string) (Membership, error) {
	r.data = map[string]Membership{}
	r.data[id] = Membership{ID: id, UserName: "kim", MembershipType: "toss"}

	if r.data[id].ID != id {
		return Membership{}, customErrors.ErrNotExistID
	}
	return r.data[id], nil
}

func (r *Repository) GetList() (map[string]Membership, error) {
	r.data["3ab365ba-6707-406d-8383-548514e2ecb9"] = Membership{ID: "3ab365ba-6707-406d-8383-548514e2ecb9", UserName: "jenny", MembershipType: "toss"}
	r.data["3ab365ba-6707-406d-8383-548514e2ecb5"] = Membership{ID: "3ab365ba-6707-406d-8383-548514e2ecb5", UserName: "jay", MembershipType: "naver"}

	return r.data, nil
}
