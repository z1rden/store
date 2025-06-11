package closer_test

import (
	"cart/internal/cart/closer"
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCloser(t *testing.T) {
	closer := closer.NewCloser(os.Interrupt)

	flag := false

	closer.Add(
		func() error {
			flag = true
			return nil
		},
		func() error {
			return errors.New("Имитация неуспешного закрытия")
		})

	// сразу в канал записано, что надо закончить
	closer.Signal()
	// time.Sleep(3 * time.Second)
	// если не сделать вэйт, то не факт, что вторая горутина на closeAll успеет сработать
	closer.Wait()

	assert.True(t, flag, "Должно закрыться")
}
