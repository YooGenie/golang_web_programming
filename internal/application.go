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
	if request.UserName == "" {
		return CreateResponse{}, errors.New("이름을 입력해주세요")
	}
	if request.MembershipType == "" {
		return CreateResponse{}, errors.New("멤버십 타입을 입력해주세요")
	}
	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		return CreateResponse{}, errors.New("해당 멤버십 타입은 유효하지 않습니다")
	}

	createResponse, err := app.repository.Create(&request)
	if err != nil {
		return CreateResponse{}, err
	}

	return createResponse, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	if request.ID == "" {
		err := errors.New("멤버십 아이디를 입력해주세요")
		return UpdateResponse{}, err
	}
	if request.UserName == "" {
		err := errors.New("이름을 입력해주세요")
		return UpdateResponse{}, err
	}
	if request.MembershipType == "" {
		err := errors.New("멤버십 타입을 입력해주세요")
		return UpdateResponse{}, err
	}
	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		err := errors.New("해당 멤버십 타입은 유효하지 않습니다")
		return UpdateResponse{}, err
	}

	updateResponse, err := app.repository.Update(&request)
	if err != nil {
		return UpdateResponse{}, err
	}

	return updateResponse, nil
}

func (app *Application) Delete(id string) error {
	if id == "" {
		return errors.New("삭제할 멤버십 아이디가 유효하지 않습니다")
	}

	if err := app.repository.Delete(id); err != nil {
		return err
	}

	return nil
}

func (app *Application) Get(id string) (GetResponse, error) {

	getResponse, err := app.repository.GetOne(id)
	if err != nil {
		return GetResponse{}, err
	}

	return getResponse, nil
}
