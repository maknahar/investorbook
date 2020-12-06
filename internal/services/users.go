package services

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/raksul-code-review/userapi-candidate-maknahar-a993286a1d8d72e3a9534ec66ef11449/internal/models"
)

type ReportServicer interface {
	GetCompanyReport(ctx context.Context, companyID int64) (*excelize.File, error)
}

func NewReportService(db *sql.DB) ReportServicer {
	return &report{model: models.NewReport(db)}
}

type report struct {
	model models.Reporter
}

func (r *report) GetCompanyReport(ctx context.Context, companyID int64) (*excelize.File, error) {
	f := excelize.NewFile()

	// Create a new sheet.
	investorSheetName, companySheetName := "Investors", "Related Companies"
	investorSheet := f.NewSheet(investorSheetName)
	f.DeleteSheet("Sheet1")

	company, err := r.model.GetCompanyByID(ctx, companyID)
	if err != nil {
		return nil, err
	}

	r.setHeader(f, company, investorSheetName)
	f.NewSheet(companySheetName)
	r.setHeader(f, company, companySheetName)

	// Set headers
	f.SetCellValue(investorSheetName, "A4", "Investor Name")
	f.SetCellValue(investorSheetName, "B4", "Amount Invested")
	f.SetCellValue(companySheetName, "A4", "Company Name")
	f.SetCellValue(companySheetName, "B4", "Shared Investors Name")

	dataRow := 5
	if investments, err := r.model.GetInvestments(ctx, companyID); err != nil {
		return nil, err
	} else {
		for _, investment := range investments {
			f.SetCellValue(investorSheetName, "A"+strconv.Itoa(dataRow), investment.InvestorName)
			f.SetCellValue(investorSheetName, "B"+strconv.Itoa(dataRow), investment.TotalInvestment)
			dataRow++
		}
	}

	dataRow = 5
	if investors, err := r.model.GetCommonInvestor(ctx, companyID); err != nil {
		return nil, err
	} else {
		for _, investor := range investors {
			f.SetCellValue(companySheetName, "A"+strconv.Itoa(dataRow), investor.CompanyName)
			f.SetCellValue(companySheetName, "B"+strconv.Itoa(dataRow), strings.Join(investor.Investors, ", "))
			dataRow++
		}

	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(investorSheet)

	return f, nil
}

func (r *report) setHeader(f *excelize.File, company *models.Company, sheetName string) {
	var headerStyle int

	var err error

	if err = f.SetColWidth(sheetName, "A", "B", 30); err != nil {
		logrus.WithError(err).Error("Could not change width")
	}

	f.SetCellValue(sheetName, "A1", company.Name+" Export")
	f.SetCellValue(sheetName, "B2", fmt.Sprintf("Generated on %s", time.Now().Format("01/02/2006")))

	if headerStyle, err = f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"5C90B9"}, Pattern: 1},
		Font: &excelize.Font{Color: "FFFFFF"},
	}); err != nil {
		logrus.WithError(err).Error("Could not set the  width")
		return
	}

	// define font style for the header
	if err := f.SetCellStyle(sheetName, "A1", "ZZ2", headerStyle); err != nil {
		fmt.Println(err)
		return
	}
}
