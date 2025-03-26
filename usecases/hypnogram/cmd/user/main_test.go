package main

import (
	"testing"
)

func TestOpenDataset(t *testing.T) {
	t.Run("open dataset", func(t *testing.T) {
		OpenDataset()
	})
}

func TestInitUsers(t *testing.T) {
	t.Run("Start User", func(t *testing.T) {
		InitUsers()
	})
}
