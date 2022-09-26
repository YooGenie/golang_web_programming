package practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestify(t *testing.T) {
	t.Run("일치 여부 확인", func(t *testing.T) {
		num1 := 1
		num2 := 1
		assert.Equal(t, num1, num2) //num1과 num2를 비교하는 것
	})

	t.Run("empty 확인", func(t *testing.T) { //값이 있는지 없는지
		assert.NotEmpty(t, "uuid")
		assert.NotEmpty(t, []int{1, 2, 3})
	})

	t.Run("숫자 값 일치 확인", func(t *testing.T) {
		num1 := 1
		num2 := int64(1)

		//assert.Equal(t, num1, num2) // FAIL
		assert.NotEqual(t, num1, num2)
		assert.EqualValues(t, num1, num2) //int32 1과 1은 비교가 안된다. 타입이 다른 숫자를 확인할 때 사용
	})

	t.Run("숫자 검증", func(t *testing.T) {
		assert.Zero(t, 0)
		assert.Positive(t, 1)
	})

	t.Run("element 확인", func(t *testing.T) {
		type Member struct {
			Name string
			Age  int
		}
		var members1, members2 []Member

		members1 = []Member{
			{Name: "Amy", Age: 10},
			{Name: "John", Age: 20},
		}
		members2 = []Member{
			{Name: "Amy", Age: 10},
			{Name: "John", Age: 20},
		}
		assert.Equal(t, members1, members2)

		members1 = []Member{
			{Name: "Amy", Age: 10},
			{Name: "John", Age: 20},
		}
		members2 = []Member{
			{Name: "John", Age: 20},
			{Name: "Amy", Age: 10},
		}
		//assert.Equal(t, members1, members2) // FAIL
		assert.ElementsMatch(t, members1, members2) // 고루틴같이 순서대로 진행이 안될때에는 순서 상관없이 배열 요소만 똑같은지 확인한다.
	})

	t.Run("len 확인", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		assert.Len(t, numbers, 3)
		assert.Equal(t, 3, len(numbers))
	})
}
