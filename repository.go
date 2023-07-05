package main

import (
	"database/sql"
	"delivery/config"
	"delivery/entity/model"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func GetDeliveries() ([]model.Delivery, error) {

	// TODO db connect via method param, like ctx context.Context
	db, err := sql.Open(config.DbDriver, config.DataSourceName)
	if err != nil {
		return nil, err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	rows, err := db.Query("select " +
		"id, " +
		"receiver_name, " +
		"receiver_phone, " +
		"receiver_address, " +
		"date_time_from, " +
		"date_time_to, " +
		"comment, " +
		"price, " +
		"status, " +
		"delivered_at, " +
		"non_delivered_reason " +
		" from delivery")
	if err != nil {
		return nil, err
	}

	var deliveries []model.Delivery
	for rows.Next() {
		d := model.Delivery{}
		// TODO see Strup to init or sqlx marsh into structure
		err := rows.Scan(
			&d.ID,
			&d.ReceiverName,
			&d.ReceiverPhone,
			&d.ReceiverAddress,
			&d.DateTimeFrom,
			&d.DateTimeTo,
			&d.Comment,
			&d.Price,
			&d.Status,
			&d.DeliveredAt,
			&d.NonDeliveredReason,
		)
		if err != nil {
			//panic()
			//if r := recover(); r != nil {
			// here processing error
			//}

			// err.Error() interface
			// fmt.Errorf("my error definition: %w", err)
			return deliveries, err
		}
		deliveries = append(deliveries, d)
	}

	return deliveries, nil

}

func CreateDelivery(delivery model.Delivery) (model.Delivery, error) {

	db, err := sql.Open(config.DbDriver, config.DataSourceName)
	if err != nil {
		return delivery, err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	sqlString := "insert into delivery (" +
		"receiver_name, " +
		"receiver_phone, " +
		"receiver_address, " +
		"date_time_from, " +
		"date_time_to, " +
		"comment, " +
		"price," +
		"status" +
		") values ($1, $2, $3, $4, $5, $6, $7, $8)"
	result, err := db.Exec(sqlString,
		delivery.ReceiverName,
		delivery.ReceiverPhone,
		delivery.ReceiverAddress,
		delivery.DateTimeFrom,
		delivery.DateTimeTo,
		delivery.Comment,
		delivery.Price,
		delivery.Status,
	)
	if err != nil {
		return delivery, err
	}

	LastInsertId, LastInsertIdErr := result.LastInsertId()
	if LastInsertIdErr != nil {
		return delivery, LastInsertIdErr
	}

	_, rowsAffectedErr := result.RowsAffected()
	if rowsAffectedErr != nil {
		return delivery, rowsAffectedErr
	}

	delivery.ID = LastInsertId
	delivery.CreatedAt = time.Now().String()

	return delivery, nil
}
