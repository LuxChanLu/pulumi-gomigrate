package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func (p *gomigrateProvider) runMigration(ctx context.Context, inputs resource.PropertyMap, preview bool, outputs map[string]interface{}) (string, error) {
	fields := inputs.Mappable()
	m, err := migrate.New(fields["sourceURL"].(string), fields["databaseURL"].(string))
	go func() {
		<-ctx.Done()
		m.GracefulStop <- true
	}()
	if err != nil {
		return "", err
	}
	if !preview {
		if err := m.Migrate(uint(fields["version"].(int))); err != nil {
			return "", err
		}
	}
	outputs["migratedAt"] = time.Now().Format(time.RFC3339)
	return strconv.FormatInt(int64(fields["version"].(int)), 10), nil
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
	if err := m.Down(); err != nil {
		return err
	}
	return nil
}
