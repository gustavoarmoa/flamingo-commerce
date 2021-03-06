package states_test

import (
	"context"
	"testing"

	"flamingo.me/flamingo-commerce/v3/checkout/domain/placeorder/process"
	"flamingo.me/flamingo-commerce/v3/checkout/domain/placeorder/states"
	"github.com/stretchr/testify/assert"
)

func TestFailed_IsFinal(t *testing.T) {
	s := states.Failed{}
	assert.True(t, s.IsFinal())
}

func TestFailed_Name(t *testing.T) {
	s := states.Failed{}
	assert.Equal(t, "Failed", s.Name())
}

func TestFailed_Rollback(t *testing.T) {
	s := states.Failed{}
	assert.Nil(t, s.Rollback(context.Background(), nil))
}

func TestFailed_Run(t *testing.T) {
	s := states.Failed{}
	assert.Equal(t, s.Run(context.Background(), &process.Process{}), process.RunResult{})
}
