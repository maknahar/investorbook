package models

import (
	"context"

	"github.com/lib/pq"

	"github.com/sirupsen/logrus"

	//nolint:gosec
	"database/sql"
	"time"
)

// Company represents all user info. Secret is populated if accessToken is given.
type Company struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Investment struct {
	InvestorID      int64
	InvestorName    string
	TotalInvestment int
}

type CommonInvestor struct {
	CompanyName string
	Investors   []string
}

type Reporter interface {
	GetInvestments(ctx context.Context, companyID int64) ([]Investment, error)
	GetCompanyByID(ctx context.Context, id int64) (*Company, error)
	GetCommonInvestor(ctx context.Context, id int64) (commonInvestors []CommonInvestor, err error)
}

type report struct {
	db *sql.DB
}

func (u report) GetCompanyByID(ctx context.Context, id int64) (*Company, error) {
	user := Company{ID: id}

	query := "SELECT name, created_at, updated_at FROM company WHERE id=$1"

	err := u.db.QueryRowContext(ctx, query, id).Scan(&user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u report) GetInvestments(ctx context.Context, id int64) (investments []Investment, err error) {
	query := `WITH investments AS (
					SELECT investor_id, SUM(amount) AS amount
					FROM investment
					WHERE company_id = $1 GROUP BY investor_id
				)
				SELECT id, name, amount
				FROM investments
						 INNER JOIN investor ON investor.id = investments.investor_id`

	rows, err := u.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			logrus.WithError(err).Error("failed to close the row")
		}
	}()

	for rows.Next() {
		var investment Investment

		err = rows.Scan(&investment.InvestorID, &investment.InvestorName, &investment.TotalInvestment)
		if err != nil {
			return nil, err
		}

		investments = append(investments, investment)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return investments, nil
}

func (u report) GetCommonInvestor(ctx context.Context, id int64) (commonInvestors []CommonInvestor, err error) {
	query := `WITH common_company AS (SELECT i2.company_id, i2.investor_id
                        FROM investment i1
                                 INNER JOIN investment i2
                                            ON i1.investor_id = i2.investor_id AND i1.company_id = $1 AND
                                               i2.company_id <> $1
                        GROUP BY i2.company_id, i2.investor_id
			)
			SELECT company.name AS company_name, array_agg(investor.name) investors
			FROM common_company
					 INNER JOIN investor ON investor.id = common_company.investor_id
					 INNER JOIN company ON common_company.company_id = company.id
			GROUP BY company.id`

	rows, err := u.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			logrus.WithError(err).Error("failed to close the row")
		}
	}()

	for rows.Next() {
		var commonInvestor CommonInvestor

		err = rows.Scan(&commonInvestor.CompanyName, pq.Array(&commonInvestor.Investors))
		if err != nil {
			return nil, err
		}

		commonInvestors = append(commonInvestors, commonInvestor)
	}

	if rows.Err() != nil {
		return nil, err
	}

	logrus.Info("Found common investor", len(commonInvestors))

	return commonInvestors, nil
}

func NewReport(d *sql.DB) Reporter {
	return report{db: d}
}
