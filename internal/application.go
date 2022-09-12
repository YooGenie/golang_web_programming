package internal

import (
	"errors"
)

type Application struct {
	service Service
}

func NewApplication(service Service) *Application {
	return &Application{service: service}
}

func (app *Application) Create(request CreateRequest) (*CreateResponse, error) {
	if request.UserName == "" {
		return nil, errors.New("이름을 입력해주세요")
	}
	if request.MembershipType == "" {
		return nil, errors.New("멤버십 타입을 입력해주세요")
	}
	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		return nil, errors.New("해당 멤버십 타입은 유효하지 않습니다")
	}

	createResponse, err := app.service.Create(&request)
	if err != nil {
		return nil, err
	}
	return createResponse, nil
}

func (app *Application) Update(request UpdateRequest) (*UpdateResponse, error) {
	if request.ID == "" {
		err := errors.New("멤버십 아이디를 입력해주세요")
		return nil, err
	}
	if request.UserName == "" {
		err := errors.New("이름을 입력해주세요")
		return nil, err
	}
	if request.MembershipType == "" {
		err := errors.New("멤버십 타입을 입력해주세요")
		return nil, err
	}
	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		err := errors.New("해당 멤버십 타입은 유효하지 않습니다")
		return nil, err
	}
	updateResponse, err := app.service.Update(&request)

	if err != nil {
		return nil, err
	}

	return updateResponse, nil
}

func (app *Application) Delete(id string) error {
	if id == "" {
		return errors.New("삭제할 멤버십 아이디가 유효하지 않습니다")
	}

	if err := app.service.Delete(id); err != nil {
		return err
	}

	return nil
}

func (app *Application) Get(id string) (*GetResponse, error) {

	getResponse, err := app.service.GetByID(id)
	if err != nil {
		return nil, err
	}

	return getResponse, nil
}
