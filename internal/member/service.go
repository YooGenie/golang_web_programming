package member

import (
	"github.com/google/uuid"
	"errors"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

var (
	ErrEmptyID = errors.New("empty id")
)

func (s *Service) Create(request *CreateRequest) (*CreateResponse, error) {
	membership := Membership{uuid.New().String(), request.UserName, request.MembershipType}
	createResponse, err := s.repository.Create(membership)
	if err != nil {
		return nil, err
	}
	return createResponse, nil
}

func (s *Service) Update(request *UpdateRequest) (*UpdateResponse, error) {
	membership := Membership{request.ID, request.UserName, request.MembershipType}

	updateResponse, err := s.repository.Update(membership)
	if err != nil {
		return nil, err
	}
	return updateResponse, nil
}

func (s *Service) GetByID(id string) (*GetResponse, error) {
	membership, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	response := GetResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}

	return &response, nil
}

func (s *Service) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

//func (s *Service) Delete(id string) error {
//	if id == "" {
//		return ErrEmptyID
//	}
//	_, err := s.repository.GetById(id)
//	if err != nil {
//		return err
//	}
//	s.repository.DeleteById(id)
//
//	for _, d := range s.repository.data {
//		fmt.Println(d)
//	}
//	return nil
//}


func (s *Service) GetList() (*[]GetResponse, error) {
	members, err := s.repository.GetList()
	if err != nil {
		return nil, err
	}

	var res []GetResponse
	for _, v := range members {
		res = append(res, GetResponse{ID: v.ID, UserName: v.UserName, MembershipType: v.MembershipType})
	}

	return &res, nil
}
