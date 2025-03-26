package main

import (
	"testing"
)

func TestInitServer(t *testing.T) {
	t.Run("Start Server", func(t *testing.T) {
		InitServer()
	})
}
