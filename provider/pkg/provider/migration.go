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
