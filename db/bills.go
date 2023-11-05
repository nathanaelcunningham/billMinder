package db

import (
	"github.com/VauntDev/tqla"
	"github.com/nathanaelcunningham/billReminder/models"
)

type BillRepository struct {
	DB *DB
}

func NewBillRepository(db *DB) *BillRepository {
	return &BillRepository{db}
}

func (r *BillRepository) GetAll() ([]models.Bill, error) {
	t, err := tqla.New(tqla.WithPlaceHolder(tqla.Dollar))
	if err != nil {
		return nil, err
	}

	stmt, _, err := t.Compile(`SELECT id,name, due_date_day, amount FROM bills`, nil)
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bills := []models.Bill{}
	for rows.Next() {
		var b models.Bill
		err := rows.Scan(&b.ID, &b.Name, &b.DueDateDay, &b.Amount)
		if err != nil {
			return nil, err
		}
		bills = append(bills, b)
	}
	return bills, nil
}

func (r *BillRepository) Create(bill *models.Bill) (int64, error) {
	t, err := tqla.New(tqla.WithPlaceHolder(tqla.Dollar))
	if err != nil {
		return 0, err
	}

	stmt, args, err := t.Compile(
		`INSERT INTO bills (name, due_date_day, amount) VALUES ({{ $.Name }}, {{ $.DueDateDay }}, {{ $.Amount }})`,
		bill,
	)
	if err != nil {
		return 0, err
	}
	res, err := r.DB.Exec(stmt, args...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *BillRepository) Delete(id int64) error {
	t, err := tqla.New(tqla.WithPlaceHolder(tqla.Dollar))
	if err != nil {
		return err
	}

	stmt, args, err := t.Compile(
		`DELETE FROM bills WHERE id = {{ $. }}`,
		id,
	)
	if err != nil {
		return err
	}
	_, err = r.DB.Exec(stmt, args...)
	if err != nil {
		return err
	}
	return nil
}
