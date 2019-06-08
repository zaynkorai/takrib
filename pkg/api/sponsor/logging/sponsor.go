package sponsor

import (
	"time"

	"github.com/labstack/echo"
	"github.com/zaynkorai/takrib/pkg/api/sponsor"
	takrib "github.com/zaynkorai/takrib/pkg/utl/model"
)

// New creates new sponsor logging service
func New(svc sponsor.Service, logger takrib.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents sponsor logging service
type LogService struct {
	sponsor.Service
	logger takrib.Logger
}

const name = "sponsor"

// Create logging
func (ls *LogService) Create(c echo.Context, req takrib.Sponsor) (resp *takrib.Sponsor, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Create sponsor request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Create(c, req)
}

// List logging
func (ls *LogService) List(c echo.Context, req *takrib.Pagination) (resp []takrib.Sponsor, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List sponsor request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List(c, req)
}

// View logging
func (ls *LogService) View(c echo.Context, req int) (resp *takrib.Sponsor, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "View sponsor request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(c, req)
}

// Delete logging
func (ls *LogService) Delete(c echo.Context, req int) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Delete sponsor request", err,
			map[string]interface{}{
				"req":  req,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Delete(c, req)
}

// Update logging
func (ls *LogService) Update(c echo.Context, req *sponsor.Update) (resp *takrib.Sponsor, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Update sponsor request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Update(c, req)
}
