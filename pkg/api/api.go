// Copyright 2017 Emir Ribic. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//  TAKRIB - Go(lang) restful starter kit
//
// API Docs for  TAKRIB v1
//
// 	 Terms Of Service:  N/A
//     Schemes: http
//     Version: 2.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Emir Ribic <zaynkorai@gmail.com> https://zaynkorai.ba
//     Host: localhost:8080
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer: []
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package api

import (
	"crypto/sha1"

	"github.com/zaynkorai/takrib/pkg/utl/zlog"

	"github.com/zaynkorai/takrib/pkg/api/auth"
	al "github.com/zaynkorai/takrib/pkg/api/auth/logging"
	at "github.com/zaynkorai/takrib/pkg/api/auth/transport"
	"github.com/zaynkorai/takrib/pkg/api/password"
	pl "github.com/zaynkorai/takrib/pkg/api/password/logging"
	pt "github.com/zaynkorai/takrib/pkg/api/password/transport"
	"github.com/zaynkorai/takrib/pkg/api/user"
	ul "github.com/zaynkorai/takrib/pkg/api/user/logging"
	ut "github.com/zaynkorai/takrib/pkg/api/user/transport"

	"github.com/zaynkorai/takrib/pkg/api/event"
	el "github.com/zaynkorai/takrib/pkg/api/event/logging"
	et "github.com/zaynkorai/takrib/pkg/api/event/transport"

	"github.com/zaynkorai/takrib/pkg/utl/config"
	"github.com/zaynkorai/takrib/pkg/utl/middleware/jwt"
	"github.com/zaynkorai/takrib/pkg/utl/postgres"
	"github.com/zaynkorai/takrib/pkg/utl/rbac"
	"github.com/zaynkorai/takrib/pkg/utl/secure"
	"github.com/zaynkorai/takrib/pkg/utl/server"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {
	db, err := postgres.New(cfg.DB.PSN, cfg.DB.Timeout, cfg.DB.LogQueries)
	if err != nil {
		return err
	}

	sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	rbac := rbac.New()
	jwt := jwt.New(cfg.JWT.Secret, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)
	log := zlog.New()

	e := server.New()
	e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	at.NewHTTP(al.New(auth.Initialize(db, jwt, sec, rbac), log), e, jwt.MWFunc())

	v1 := e.Group("/v1")
	v1.Use(jwt.MWFunc())

	ut.NewHTTP(ul.New(user.Initialize(db, rbac, sec), log), v1)
	pt.NewHTTP(pl.New(password.Initialize(db, rbac, sec), log), v1)

	et.NewHTTP(el.New(event.Initialize(db, rbac, sec), log), v1)


	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
