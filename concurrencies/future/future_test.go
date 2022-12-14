package future

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}
	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(s string) {
			t.Fail()
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "Hello World", nil
		})
		wg.Wait()
	})

	t.Run("Error result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)
		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(s string) {
			t.Log(s)
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "", errors.New("error ocurred")
		})
		wg.Wait()
	})
}

func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("Timeout!")

	t.Fail()
	wg.Done()
}
