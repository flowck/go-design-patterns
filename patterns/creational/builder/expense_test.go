package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExpenseBuilder_Build(t *testing.T) {
	b := NewExpenseBuilder()
	b.SetStatus("pending")
	expense, err := b.Build()

	assert.NotNil(t, err, "expense cannot be created without values")
	assert.Nil(t, expense)

	b = NewExpenseBuilder()
	b.SetId("some-unique-id")
	b.SetStatus("pending")
	b.SetAmount(100)
	b.SetAmountTax(0)
	b.SetCurrencyCode("USD")

	expense, err = b.Build()

	assert.Nil(t, err)
	assert.NotNil(t, expense)
}
