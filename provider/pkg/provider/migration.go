package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/aws_s3"
	_ "github.com/golang-migrate/migrate/v4/source/bitbucket"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/golang-migrate/migrate/v4/source/github_ee"
	_ "github.com/golang-migrate/migrate/v4/source/gitlab"
	_ "github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/golang-migrate/migrate/v4/source/godoc_vfs"
	_ "github.com/golang-migrate/migrate/v4/source/google_cloud_storage"
	_ "github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/golang-migrate/migrate/v4/source/pkger"
	_ "github.com/golang-migrate/migrate/v4/source/stub"

	_ "github.com/golang-migrate/migrate/v4/database/cassandra"
	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/database/firebird"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/database/multistmt"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/neo4j"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/ql"
	_ "github.com/golang-migrate/migrate/v4/database/redshift"
	_ "github.com/golang-migrate/migrate/v4/database/snowflake"
	_ "github.com/golang-migrate/migrate/v4/database/spanner"
	_ "github.com/golang-migrate/migrate/v4/database/sqlcipher"
	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/database/stub"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func (p *gomigrateProvider) runMigration(ctx context.Context, inputs resource.PropertyMap, preview bool, outputs map[string]interface{}) (string, error) {
	fields := inputs.Mappable()
	outputs["sourceURL"] = fields["sourceURL"]
	outputs["databaseURL"] = fields["databaseURL"]
	outputs["version"] = fields["version"]
	outputs["prevVersion"] = fields["prevVersion"]
	m, err := migrate.New(fields["sourceURL"].(string), fields["databaseURL"].(string))
	go func() {
		<-ctx.Done()
		m.GracefulStop <- true
	}()
	if err != nil {
		return "", err
	}
	if !preview {
		if err := m.Migrate(uint(fields["version"].(float64))); err != nil {
			return "", err
		}
	}
	outputs["migratedAt"] = time.Now().Format(time.RFC3339)
	return strconv.FormatInt(int64(fields["version"].(float64)), 10), nil
}

func (p *gomigrateProvider) undoMigration(ctx context.Context, inputs resource.PropertyMap) error {
	fields := inputs.Mappable()
	m, err := migrate.New(fields["sourceURL"].(string), fields["databaseURL"].(string))
	go func() {
		<-ctx.Done()
		m.GracefulStop <- true
	}()
	if err != nil {
		return err
	}
	if int(fields["prevVersion"].(float64)) == -1 {
		if err := m.Down(); err != nil {
			return err
		}
		return nil
	}
	if err := m.Migrate(uint(fields["prevVersion"].(float64))); err != nil {
		return err
	}
	return nil
}

func (p *gomigrateProvider) diffMigration(ctx context.Context, diff *resource.ObjectDiff) ([]string, bool) {
	changes := []string{}
	recreate := false
	for _, key := range diff.ChangedKeys() {
		if key == "sourceURL" || key == "databaseURL" || key == "version" || key == "prevVersion" {
			changes = append(changes, string(key))
			recreate = true
		}
	}
	return changes, recreate
}
