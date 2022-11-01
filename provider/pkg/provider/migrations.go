package provider

import (
	"os"
	"strconv"

	"github.com/LuxChanLu/pulumi-gomigrate/sdk/go/gomigrate"
	"github.com/golang-migrate/migrate/v4/source"
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
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/database/stub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Migrations struct {
	pulumi.ResourceState
}

func (p *gomigrateProvider) newMigrations(ctx *pulumi.Context, name string, sourceURL, databaseURL pulumi.Input, opts ...pulumi.ResourceOption) (*Migrations, error) {
	migrations := &Migrations{}
	if err := ctx.RegisterComponentResource("gomigrate:index:Migrations", name, migrations, opts...); err != nil {
		return nil, err
	}
	(sourceURL.(pulumi.StringInput)).ToStringOutput().ApplyT(func(src string) (*Migrations, error) {
		sourceDrv, err := source.Open(src)
		if err != nil {
			return nil, err
		}
		var version uint
		var prevVersion int = -1
		var migration *gomigrate.Migration = nil
		for version, err = sourceDrv.First(); err == nil; version, err = sourceDrv.Next(version) {
			opts := []pulumi.ResourceOption{pulumi.Parent(migrations)}
			if migration != nil {
				opts = append(opts, pulumi.DependsOn([]pulumi.Resource{migration}))
			}
			migration, err = gomigrate.NewMigration(ctx, strconv.FormatUint(uint64(version), 10), &gomigrate.MigrationArgs{
				SourceURL: pulumi.ToSecret(src).(pulumi.StringOutput), DatabaseURL: pulumi.ToSecret(databaseURL).(pulumi.StringOutput), Version: pulumi.Int(version),
				PrevVersion: pulumi.Int(prevVersion),
			}, opts...)
			if err != nil {
				return nil, err
			}
			prevVersion = int(version)
		}
		if err != os.ErrNotExist {
			return nil, err
		}
		return migrations, nil
	})
	return migrations, nil
}
