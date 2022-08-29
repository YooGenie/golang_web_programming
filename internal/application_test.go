package internal

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{"data": {uuid.New().String(), "jenny", "naver"}}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Empty(t, res)
		assert.EqualError(t, err, "이미 등록된 사용자 이름입니다")

	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"", "naver"}
		res, err := app.Create(req)
		assert.Empty(t, res)
		assert.EqualError(t, err, "이름을 입력해주세요")
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", ""}
		res, err := app.Create(req)

		assert.Empty(t, res)
		assert.EqualError(t, err, "멤버십 타입을 입력해주세요")
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "kakao"}
		res, err := app.Create(req)

		assert.Empty(t, res)
		assert.EqualError(t, err, "해당 멤버십 타입은 유효하지 않습니다")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("internal 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{"data": {uuid.New().String(), "jenny", "naver"}}))
		req := UpdateRequest{uuid.New().String(), "genie", "naver"}
		res, err := app.Update(req)

		assert.Nil(t, err)
		assert.Equal(t, req.ID, res.ID)
		assert.Equal(t, req.UserName, res.UserName)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{"data": {uuid.New().String(), "genie", "naver"}}))
		req := UpdateRequest{uuid.New().String(), "genie", "naver"}
		res, err := app.Update(req)
		assert.Empty(t, res)
		assert.EqualError(t, err, "사용자의 이름이 이미 존재합니다")
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{"data": {uuid.New().String(), "jenny", "naver"}}))
		req := UpdateRequest{"", "genie", "naver"}
		res, err := app.Update(req)
		assert.Empty(t, res)
		assert.EqualError(t, err, "멤버십 아이디를 입력해주세요")

	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := UpdateRequest{uuid.New().String(), "", "naver"}
		res, err := app.Update(req)
		assert.Empty(t, res)
		assert.EqualError(t, err, "이름을 입력해주세요")
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := UpdateRequest{uuid.New().String(), "genie", ""}
		res, err := app.Update(req)
		assert.Empty(t, res)
		assert.EqualError(t, err, "멤버십 타입을 입력해주세요")
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := UpdateRequest{uuid.New().String(), "jenny", "kakao"}
		res, err := app.Update(req)
		assert.Empty(t, res)
		assert.EqualError(t, err, "해당 멤버십 타입은 유효하지 않습니다")
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{"data": {uuid.New().String(), "jenny", "naver"}}))
		params := app.repository.data["data"].ID
		err := app.Delete(params)
		assert.Nil(t, err)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{"data": {uuid.New().String(), "jenny", "naver"}}))
		params := ""
		err := app.Delete(params)
		assert.EqualError(t, err, "삭제할 멤버십 아이디가 유효하지 않습니다")
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{"data": {uuid.New().String(), "jenny", "naver"}}))
		params := uuid.New().String()
		err := app.Delete(params)
		assert.EqualError(t, err, "입력한 id가 존재하지 않습니다")
	})
}

func TestGet(t *testing.T) {
	t.Run("멤버십을 조회한다", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{"data": {uuid.New().String(), "jenny", "naver"}}))
		params := app.repository.data["data"].ID
		res, err := app.Get(params)
		assert.Nil(t, err)
		assert.Equal(t, params, res.ID)
		assert.Equal(t, "jenny", res.UserName)
		assert.Equal(t, "naver", res.MembershipType)

	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{"data": {uuid.New().String(), "jenny", "naver"}}))
		params := uuid.New().String()
		res, err := app.Get(params)
		assert.Empty(t, res)
		assert.EqualError(t, err, "입력한 id가 존재하지 않습니다")
	})
}
