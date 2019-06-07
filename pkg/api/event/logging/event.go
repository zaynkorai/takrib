package event

import (
	"time"
	"github.com/labstack/echo"
	"github.com/zaynkorai/takrib/pkg/api/event"
	"github.com/zaynkorai/takrib/pkg/utl/model"
)

// New creates new event logging service
func New(svc event.Service, logger takrib.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents event logging service
type LogService struct {
	event.Service
	logger takrib.Logger
}

const name = "event"

// Create logging
func (ls *LogService) Create(c echo.Context, req takrib.Event) (resp *takrib.Event, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Create event request", err,
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
func (ls *LogService) List(c echo.Context, req *takrib.Pagination) (resp []takrib.Event, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List event request", err,
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
func (ls *LogService) View(c echo.Context, req int) (resp *takrib.Event, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "View event request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(c, req)
}

// // Delete logging
// func (ls *LogService) Delete(c echo.Context, req int) (err error) {
// 	defer func(begin time.Time) {
// 		ls.logger.Log(
// 			c,
// 			name, "Delete event request", err,
// 			map[string]interface{}{
// 				"req":  req,
// 				"took": time.Since(begin),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.Delete(c, req)
// }

// // Update logging
// func (ls *LogService) Update(c echo.Context, req *event.Update) (resp *takrib.Event, err error) {
// 	defer func(begin time.Time) {
// 		ls.logger.Log(
// 			c,
// 			name, "Update event request", err,
// 			map[string]interface{}{
// 				"req":  req,
// 				"resp": resp,
// 				"took": time.Since(begin),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.Update(c, req)
// }
