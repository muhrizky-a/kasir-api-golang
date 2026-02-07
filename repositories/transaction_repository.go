package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
	"strings"
	"time"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		if item.Quantity > stock {
			return nil, fmt.Errorf("insufficient product id %d stock", item.ProductID)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	var (
		placeholders []string
		args         []interface{}
	)
	for i, _ := range details {
		details[i].TransactionID = transactionID
		placeholders = append(
			placeholders,
			fmt.Sprintf("($%d,$%d,$%d,$%d)",
				i*4+1,
				i*4+2,
				i*4+3,
				i*4+4,
			))

		args = append(args,
			transactionID,
			details[i].ProductID,
			details[i].Quantity,
			details[i].Subtotal)
	}
	query := fmt.Sprintf(
		"INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES %s",
		strings.Join(placeholders, ","),
	)

	_, err = tx.Exec(
		query, args...,
	)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

func (repo *TransactionRepository) GetTodayReport() (*models.Report, error) {
	var report models.Report
	var produk models.ProdukTerlaris

	t := time.Now()
	dateStr := t.Format("2006-01-02")

	query := `
	SELECT 
		COALESCE(sum(transactions.total_amount), 0) as total_revenue,
		COUNT(transactions.id) as total_transaksi
		FROM transactions 
		WHERE transactions.created_at >= $1::date
  		AND transactions.created_at < ($1::date + INTERVAL '1 day')
  `

	err := repo.db.QueryRow(query, dateStr).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}
	if report.TotalRevenue < 1 || report.TotalTransaksi < 1 {
		return &report, nil
	}

	query = `
	SELECT
		products.name as nama,
		SUM(transaction_details.quantity) as qty_terjual
		FROM transactions 
		JOIN transaction_details
		ON transactions.id = transaction_details.transaction_id
		JOIN products
		ON products.id = transaction_details.product_id
		WHERE transactions.created_at >= $1::date
		AND transactions.created_at <  ($1::date + INTERVAL '1 day')
		GROUP by nama
		ORDER by qty_terjual DESC
		LIMIT 1
  `

	err = repo.db.QueryRow(query, dateStr).Scan(&produk.Nama, &produk.QtyTerjual)
	if err != nil {
		return nil, err
	}

	report.ProdukTerlaris = produk
	return &report, nil
}
