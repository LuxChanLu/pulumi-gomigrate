package provider

import (
	"os"
	"strconv"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Migrations struct {
	pulumi.ResourceState
}

type Migration struct {
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
		var migration *Migration = nil
		for version, err = sourceDrv.First(); err == nil; version, err = sourceDrv.Next(version) {
			opts := []pulumi.ResourceOption{pulumi.Parent(migrations)}
			if migration != nil {
				opts = append(opts, pulumi.DependsOn([]pulumi.Resource{migration}))
			}
			migration = &Migration{}
			if err := ctx.RegisterRemoteComponentResource("gomigrate:index:Migration", strconv.FormatUint(uint64(version), 10), pulumi.Map{
				"sourceURL": sourceURL, "databaseURL": databaseURL, "version": pulumi.Int(version),
			}, migration, opts...); err != nil {
				return nil, err
			}
		}
		if err != os.ErrNotExist {
			return nil, err
		}
		return migrations, nil
	})
	return migrations, nil
}
