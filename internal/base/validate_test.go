package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("req is nil", func(t *testing.T) {
		t.Parallel()

		var in *base.ValidateRequest
		rsl := base.Validate(in)

		assert.Contains(t, rsl, "Запрос не должен быть пустым")
	})

	t.Run("Not filled name", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			UserID:      "1234",
			Title:       "",
			Description: "Test request",
		}
		rsl := base.Validate(in)

		assert.Contains(t, rsl, "Не указано имя запроса")
	})

	t.Run("Request is correct", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			UserID:      "1234",
			Title:       "NameTest",
			Description: "Test request",
		}
		rsl := base.Validate(in)

		assert.Contains(t, rsl, "Запрос корректный!")
	})
}
