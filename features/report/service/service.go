package service

import (
	"time"

	"github.com/roihan12/technical-test-kalbe/features/report"
	"github.com/roihan12/technical-test-kalbe/utils"
)

type attendanceReportService struct {
	query report.AttendanceReportData
}

func New(q report.AttendanceReportData) report.AttendanceReportService {
	return &attendanceReportService{
		query: q,
	}
}

func (s *attendanceReportService) GenerateReport(startTime, endTime time.Time) ([]report.AttendanceReport, error) {
	// Panggil method dari data layer untuk mendapatkan data report absen
	reports, err := s.query.GenerateReport(startTime, endTime)
	if err != nil {
		return nil, utils.ErrInternal
	}

	return reports, nil
}
