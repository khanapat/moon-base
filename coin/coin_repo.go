package coin

import (
	"context"
	"database/sql"
	"fmt"
)

type ResetSupplyFn func(ctx context.Context) error

func NewResetSupplyFn(db *sql.DB) ResetSupplyFn {
	return func(ctx context.Context) error {
		_, err := db.ExecContext(ctx, `
			if not exists (select * from sysobjects where name='coin' and xtype='U')
				create table coin (
					id int IDENTITY (1, 1) not null,
					coin_name varchar(255) not null,
					supply float not null,
					primary key(id)
				)
		;`)
		if err != nil {
			return err
		}

		_, err = db.ExecContext(ctx, `
			insert into master.dbo.coin
			(
				coin_name,
				supply
			)
			values
			(
				'MOON',
				1000
			)
		;`)
		if err != nil {
			return err
		}

		_, err = db.ExecContext(ctx, `
			if not exists (select * from sysobjects where name='history' and xtype='U')
			create table history (
				number int IDENTITY (1, 1) not null,
				date_time datetime2 not null,
				user_id varchar(255) not null,
				thbt float not null,
				moon float not null,
				rate varchar(255) not null,
				primary key(number)
			)
		;`)
		if err != nil {
			return err
		}

		return nil
	}
}

type GetSupplyFn func(ctx context.Context, request map[string]interface{}) (*[]Coin, error)

func NewGetSupplyFn(db *sql.DB) GetSupplyFn {
	return func(ctx context.Context, request map[string]interface{}) (*[]Coin, error) {
		coins := make([]Coin, 0)
		param := make([]interface{}, 0)
		query := `
			SELECT	id,
					coin_name,
					supply
			FROM master.dbo.coin
			WHERE 1 = 1
		`
		index := 1
		for key, value := range request {
			query = fmt.Sprintf("%s AND %s = @p%d", query, key, index)
			param = append(param, value)
			index++
		}
		rows, err := db.QueryContext(ctx, query, param...)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var coin Coin
			err := rows.Scan(
				&coin.ID,
				&coin.CoinName,
				&coin.Supply,
			)
			if err != nil {
				return nil, err
			}
			coins = append(coins, coin)
		}
		defer rows.Close()
		return &coins, nil
	}
}

type GetHistoryLogsFn func(ctx context.Context, request map[string]interface{}, orderBy string) (*[]History, error)

func NewGetHistoryLogsFn(db *sql.DB) GetHistoryLogsFn {
	return func(ctx context.Context, request map[string]interface{}, orderBy string) (*[]History, error) {
		historys := make([]History, 0)
		param := make([]interface{}, 0)
		query := `
			SELECT	number,
					date_time,
					user_id,
					thbt,
					moon,
					rate
			FROM master.dbo.history
			WHERE 1 = 1
		`
		index := 1
		for key, value := range request {
			switch key {
			case "history_date_from":
				query = fmt.Sprintf("%s AND date_time >= @p%d", query, index)
			case "history_date_to":
				query = fmt.Sprintf("%s AND date_time < @p%d", query, index)
			default:
				query = fmt.Sprintf("%s AND %s = @p%d", query, key, index)
			}
			param = append(param, value)
			index++
		}
		query = fmt.Sprintf("%s ORDER BY date_time %s", query, orderBy)
		rows, err := db.QueryContext(ctx, query, param...)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var history History
			err := rows.Scan(
				&history.Number,
				&history.DateTime,
				&history.UserID,
				&history.THBT,
				&history.Moon,
				&history.Rate,
			)
			if err != nil {
				return nil, err
			}
			historys = append(historys, history)
		}
		defer rows.Close()
		return &historys, nil
	}
}

type GetSupplyByIDFn func(ctx context.Context, id int) (*float64, error)

func NewGetSupplyByIDFn(db *sql.DB) GetSupplyByIDFn {
	return func(ctx context.Context, id int) (*float64, error) {
		var supply float64
		err := db.QueryRowContext(ctx, `
			SELECT supply
			FROM master.dbo.coin
			WHERE id = @p1
		;`, id).Scan(&supply)
		switch {
		case err == sql.ErrNoRows:
			return nil, nil
		case err != nil:
			return nil, err
		default:
			return &supply, nil
		}
	}
}

type UpdateSupplyByIDFn func(ctx context.Context, id int, supply float64) error

func NewUpdateSupplyByIDFn(db *sql.DB) UpdateSupplyByIDFn {
	return func(ctx context.Context, id int, supply float64) error {
		_, err := db.ExecContext(ctx, `
			UPDATE master.dbo.coin
			SET supply = @p1
			WHERE id = @p2
		;`, supply, id)
		if err != nil {
			return err
		}
		return nil
	}
}

type CreateHistoryLogsFn func(ctx context.Context, date, userId, rate string, thbt, moon float64) error

func NewCreateHistoryLogsFn(db *sql.DB) CreateHistoryLogsFn {
	return func(ctx context.Context, date, userId, rate string, thbt, moon float64) error {
		_, err := db.ExecContext(ctx, `
			INSERT INTO master.dbo.history
			(
				date_time,
				user_id,
				thbt,
				moon,
				rate
			)
			VALUES
			(
				@p1,
				@p2,
				@p3,
				@p4,
				@p5
			)
		;`, date, userId, thbt, moon, rate)
		if err != nil {
			return err
		}
		return nil
	}
}
