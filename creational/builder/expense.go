package main

import (
	"errors"
	"time"
)

type Expense struct {
	id           string
	date         time.Time
	attachment   string
	notes        string
	amount       float32
	amountTax    float32
	currencyCode string
	status       string
}

type ExpenseBuilder struct {
	Id           string
	Date         time.Time
	Attachment   string
	Notes        string
	Amount       float32
	AmountTax    float32
	CurrencyCode string
	Status       string
}

func NewExpenseBuilder() *ExpenseBuilder {
	return &ExpenseBuilder{}
}

func (b *ExpenseBuilder) SetId(id string) {
	b.Id = id
}

func (b *ExpenseBuilder) SetDate(date time.Time) {
	b.Date = date
}

func (b *ExpenseBuilder) SetAttachment(attachment string) {
	b.Attachment = attachment
}

func (b *ExpenseBuilder) SetNotes(notes string) {
	b.Notes = notes
}

func (b *ExpenseBuilder) SetAmount(amount float32) {
	b.Amount = amount
}

func (b *ExpenseBuilder) SetAmountTax(amountTax float32) {
	b.AmountTax = amountTax
}

func (b *ExpenseBuilder) SetCurrencyCode(currencyCode string) {
	b.CurrencyCode = currencyCode
}

func (b *ExpenseBuilder) SetStatus(status string) {
	b.Status = status
}

func (b *ExpenseBuilder) Build() (*Expense, error) {

	// if b.Status != "pending" || b.Status != "approved" || b.Status != "rejected" {
	if b.Status == "" {
		return nil, errors.New("status must be: pending, approved or rejected")
	}

	if b.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	if b.AmountTax < 0 {
		return nil, errors.New("amountTax must be a float between 0.0 and 1.0")
	}

	return &Expense{
		id:           b.Id,
		date:         b.Date,
		attachment:   b.Attachment,
		notes:        b.Notes,
		amount:       b.Amount,
		amountTax:    b.AmountTax,
		currencyCode: b.CurrencyCode,
		status:       b.Status,
	}, nil
}
