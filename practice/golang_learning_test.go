package practice

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"sync"
	"testing"
	"time"
)

// golang 학습 테스트
func TestGolang(t *testing.T) {
	t.Run("string test", func(t *testing.T) {
		str := "Ann,Jenny,Tom,Zico"
		actual := strings.Split(str, ",") // TODO str을 , 단위로 잘라주세요.
		expected := []string{"Ann", "Jenny", "Tom", "Zico"}

		//TODO assert 문을 활용해 actual과 expected를 비교해주세요.
		assert.Equal(t, expected, actual)
	})

	t.Run("goroutine에서 slice에 값 추가해보기", func(t *testing.T) {
		var numbers []int
		var mutex sync.Mutex
		var wg sync.WaitGroup
		wg.Add(100)
		for i := 0; i < 100; i++ {
			go func(i int) {
				// TODO numbers에 i 값을 추가해보세요.
				mutex.Lock()
				numbers = append(numbers, i)
				mutex.Unlock()
				wg.Done()
			}(i)
		}
		wg.Wait()

		var expected []int // actual : [0 1 2 ... 99]

		// TODO expected를 만들어주세요.
		for i := 0; i < 100; i++ {
			expected = append(expected, i)
		}
		assert.ElementsMatch(t, expected, numbers)
	})

	t.Run("fan out, fan in", func(t *testing.T) {
		/*
			TODO 주어진 코드를 수정해서 테스트가 통과하도록 해주세요!
			- inputCh에 1, 2, 3 값을 넣는다.
			- inputCh로 부터 값을 받아와, value * 10 을 한 후 outputCh에 값을 넣어준다.
			- outputCh에서 읽어온 값을 비교한다.
		*/

		inputCh := generate()
		outputCh := make(chan int)
		go func() {
			for {
				select {
				case value, ok := <-inputCh:
					if !ok {
						close(outputCh)
						return
					}
					outputCh <- value * 10
				}
			}
		}()

		var actual []int
		for value := range outputCh {
			actual = append(actual, value)
		}
		expected := []int{10, 20, 30}
		assert.Equal(t, expected, actual)
	})

	t.Run("context timeout", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		ctx := context.TODO() // TODO 3초후에 종료하는 timeout context로 만들어주세요.
		timeOutCtx, _ := context.WithTimeout(ctx, add)

		var endTime time.Time
		select {
		case <-timeOutCtx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context deadline", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		ctx := context.TODO() // TODO 3초후에 종료하는 timeout context로 만들어주세요.
		deadlineCtx, _ := context.WithDeadline(ctx, time.Now().Add(add))

		var endTime time.Time
		select {
		case <-deadlineCtx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context value", func(t *testing.T) {
		// context에 key, value를 추가해보세요.
		ctx := context.Background()
		ctx = context.WithValue(ctx, "name", "유지니")
		// 추가된 key, value를 호출하여 assert로 값을 검증해보세요.
		assert.Equal(t, "유지니", ctx.Value("name"))
		// 추가되지 않은 key에 대한 value를 assert로 검증해보세요.
		assert.Nil(t, ctx.Value("age"))
	})
}

func generate() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()
	return ch
}
