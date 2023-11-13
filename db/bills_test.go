package db

import (
	"fmt"
	"reflect"
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

func TestBillRepository_Get(t *testing.T) {
	TestBillRepository_Create(t)
	r := &BillRepository{
		DB: database,
	}
	id := 1
	bill, err := r.Get(int64(id))
	require.NoError(t, err)
	fmt.Println(bill)
}

func TestBillRepository_GetAll(t *testing.T) {
	type fields struct {
		DB *DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Bill
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &BillRepository{
				DB: tt.fields.DB,
			}
			got, err := r.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("BillRepository.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BillRepository.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
