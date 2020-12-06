package services

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/maknahar/investorbook/internal/models"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/sirupsen/logrus"
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

	wg := sync.WaitGroup{}
	wg.Add(2)

	go r.addInvestmentDetails(ctx, &wg, f, company, investorSheetName)
	go r.addCommonInvestor(ctx, &wg, f, company, companySheetName)

	wg.Wait()

	// Set active sheet of the workbook.
	f.SetActiveSheet(investorSheet)

	return f, nil
}

func (r *report) addInvestmentDetails(ctx context.Context, wg *sync.WaitGroup, f *excelize.File, company *models.Company, sheetName string) error {
	dataRow := 5
	defer wg.Done()
	if investments, err := r.model.GetInvestments(ctx, company.ID); err != nil {
		return err
	} else {
		for _, investment := range investments {
			f.SetCellValue(sheetName, "A"+strconv.Itoa(dataRow), investment.InvestorName)
			f.SetCellValue(sheetName, "B"+strconv.Itoa(dataRow), investment.TotalInvestment)
			dataRow++
		}
	}
	return nil
}

func (r *report) addCommonInvestor(ctx context.Context, wg *sync.WaitGroup, f *excelize.File, company *models.Company, sheetName string) error {
	dataRow := 5
	defer wg.Done()
	if investors, err := r.model.GetCommonInvestor(ctx, company.ID); err != nil {
		return err
	} else {
		for _, investor := range investors {
			f.SetCellValue(sheetName, "A"+strconv.Itoa(dataRow), investor.CompanyName)
			f.SetCellValue(sheetName, "B"+strconv.Itoa(dataRow), strings.Join(investor.Investors, ", "))
			dataRow++
		}
	}
	return nil
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
