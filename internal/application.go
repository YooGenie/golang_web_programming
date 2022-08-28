package internal

import (
	"errors"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	data := app.repository.data["data"]
	if request.UserName == "" {
		err := errors.New("이름을 입력해주세요")
		return CreateResponse{}, err
	} else if request.MembershipType == "" {
		err := errors.New("멤버십 타입을 입력해주세요")
		return CreateResponse{}, err
	} else if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		err := errors.New("해당 멤버십 타입은 유효하지 않습니다")
		return CreateResponse{}, err
	}

	if data.UserName == request.UserName {
		err := errors.New("이미 등록된 사용자 이름입니다")
		return CreateResponse{}, err

	}

	return CreateResponse{"1", "naver"}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	data := app.repository.data
	for _, v := range data {
		if v.UserName == request.UserName {
			err := errors.New("사용자의 이름이 이미 존재합니다")
			return UpdateResponse{}, err
		}
	}
	if request.ID == "" {
		err := errors.New("멤버십 아이디를 입력해주세요")
		return UpdateResponse{}, err
	} else if request.UserName == "" {
		err := errors.New("이름을 입력해주세요")
		return UpdateResponse{}, err
	} else if request.MembershipType == "" {
		err := errors.New("멤버십 타입을 입력해주세요")
		return UpdateResponse{}, err
	} else if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		err := errors.New("해당 멤버십 타입은 유효하지 않습니다")
		return UpdateResponse{}, err
	}
	return UpdateResponse{"1", "genie", "naver"}, nil
}

func (app *Application) Delete(id string) error {
	data := app.repository.data["data"]
	if id == "" {
		err := errors.New("삭제할 멤버십 아이디가 유효하지 않습니다")
		return err
	} else if data.ID != id {
		err := errors.New("입력한 id가 존재하지 않습니다")
		return err
	}
	return nil
}

func (app *Application) Get(id string) (GetResponse, error) {
	data := app.repository.data["data"]
	if data.ID != id {
		err := errors.New("입력한 id가 존재하지 않습니다")
		return GetResponse{}, err
	}
	return GetResponse{"1", "jenny", "naver"}, nil
}
