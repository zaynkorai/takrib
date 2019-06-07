package speaker

import (
	"time"

	"github.com/labstack/echo"
	"github.com/zaynkorai/takrib/pkg/api/speaker"
	takrib "github.com/zaynkorai/takrib/pkg/utl/model"
)

// New creates new speaker logging service
func New(svc speaker.Service, logger takrib.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents speaker logging service
type LogService struct {
	speaker.Service
	logger takrib.Logger
}

const name = "speaker"

// Create logging
func (ls *LogService) Create(c echo.Context, req takrib.Speaker) (resp *takrib.Speaker, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Create speaker request", err,
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
func (ls *LogService) List(c echo.Context, req *takrib.Pagination) (resp []takrib.Speaker, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List speaker request", err,
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
func (ls *LogService) View(c echo.Context, req int) (resp *takrib.Speaker, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "View speaker request", err,
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
// 			name, "Delete speaker request", err,
// 			map[string]interface{}{
// 				"req":  req,
// 				"took": time.Since(begin),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.Delete(c, req)
// }

// // Update logging
// func (ls *LogService) Update(c echo.Context, req *speaker.Update) (resp *takrib.Speaker, err error) {
// 	defer func(begin time.Time) {
// 		ls.logger.Log(
// 			c,
// 			name, "Update speaker request", err,
// 			map[string]interface{}{
// 				"req":  req,
// 				"resp": resp,
// 				"took": time.Since(begin),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.Update(c, req)
// }
