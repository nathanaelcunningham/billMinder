package db

import (
	"testing"

	"github.com/nathanaelcunningham/billReminder/models"
	"github.com/stretchr/testify/require"
)

func TestBillRepository_Create(t *testing.T) {
	r := &BillRepository{
		DB: database,
	}
	bill := &models.Bill{
		Name:       "Test Bill",
		DueDateDay: 31,
		Amount:     100.00,
	}
	res, err := r.Create(bill)
	require.NoError(t, err)
	want := int64(1)
	got := res

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
