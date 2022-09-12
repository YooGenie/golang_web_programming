package errors

import "errors"

var (
	ErrParamsBinding =  errors.New("파라미터 바인딩중 오류가 발생했습니다")
	ErrNotExistID = errors.New("입력한 id가 존재하지 않습니다")
	ErrExistUserName = errors.New("이미 등록된 사용자 이름입니다")
	ErrInvalidUserID = errors.New("사용자 아이디가 유효하지 않습니다")
	ErrInvalidMembershipType = errors.New("해당 멤버십 타입은 유효하지 않습니다")
	ErrInputMembershipType = errors.New("멤버십 타입을 입력해주세요")
	ErrInputUserName = errors.New("사용자 이름을 입력해주세요")
)