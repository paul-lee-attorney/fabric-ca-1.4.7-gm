/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package factory

import (
	"context"

	"github.com/paul-lee-attorney/fabric-2.1-gm/bccsp"
	"github.com/paul-lee-attorney/fabric-ca-1.4.7-gm/lib/server/db"
	"github.com/paul-lee-attorney/fabric-ca-1.4.7-gm/lib/server/db/mysql"
	"github.com/paul-lee-attorney/fabric-ca-1.4.7-gm/lib/server/db/postgres"
	"github.com/paul-lee-attorney/fabric-ca-1.4.7-gm/lib/server/db/sqlite"
	"github.com/paul-lee-attorney/fabric-ca-1.4.7-gm/lib/tls"
	"github.com/pkg/errors"
)

// DB is interface that defines the functions on a database
type DB interface {
	Connect() error
	PingContext(ctx context.Context) error
	Create() (*db.DB, error)
}

// New returns a DB interface for the request database type
func New(
	dbType,
	datasource,
	caName string,
	tlsConfig *tls.ClientTLSConfig,
	csp bccsp.BCCSP,
	metrics *db.Metrics,
) (DB, error) {
	switch dbType {
	case "sqlite3":
		return sqlite.NewDB(datasource, caName, metrics), nil
	case "postgres":
		return postgres.NewDB(datasource, caName, tlsConfig, metrics), nil
	case "mysql":
		return mysql.NewDB(datasource, caName, tlsConfig, csp, metrics), nil
	default:
		return nil, errors.Errorf("Invalid db.type in config file: '%s'; must be 'sqlite3', 'postgres', or 'mysql'", dbType)
	}
}
