package src

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHelloFrom(t *testing.T) {
	t.Run("returns wanted salutation", func(t *testing.T) {
		want := "Hello Mars! How are you?"
		got := GetHelloFrom("Mars")

		assert.Equal(t, want, got)
	})
}
