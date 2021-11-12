package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/go-playground/assert/v2"
)

func TestCountDay(t *testing.T) {
	year := 2020
	month := 11
	assert.Equal(t, CountTotalDay(year, month), 30)
}
