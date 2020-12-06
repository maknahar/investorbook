package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/raksul-code-review/userapi-candidate-maknahar-a993286a1d8d72e3a9534ec66ef11449/internal/services"
)

type ReportHandler struct {
	service services.ReportServicer
}

func NewReportHandler(db *sql.DB) *ReportHandler {
	return &ReportHandler{service: services.NewReportService(db)}
}

func (u *ReportHandler) GenerateCompanyReport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := logrus.WithField("companyId", chi.URLParam(r, "companyId"))
	l.Debug("Generating the report")

	companyID, err := strconv.ParseInt(chi.URLParam(r, "companyId"), 10, 64)
	if err != nil {
		l.WithError(err).Debug("bad request")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	file, err := u.service.GetCompanyReport(ctx, companyID)
	if err != nil {
		logrus.WithError(err).WithField("companyId", chi.URLParam(r, "companyId")).
			Error("unable to generate the request")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", `attachment;filename="companyDataExport.xlsx"`)
	w.Header().Set("File-Name", "companyDataExport.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")

	err = file.Write(w)
	if err != nil {
		logrus.WithError(err).WithField("companyId", chi.URLParam(r, "companyId")).
			Error("unable to write report")

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	return
}
