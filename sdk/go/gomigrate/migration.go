// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gomigrate

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Migration struct {
	pulumi.CustomResourceState

	// Date of the migration
	MigratedAt pulumi.StringOutput `pulumi:"migratedAt"`
}

// NewMigration registers a new resource with the given unique name, arguments, and options.
func NewMigration(ctx *pulumi.Context,
	name string, args *MigrationArgs, opts ...pulumi.ResourceOption) (*Migration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseURL == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseURL'")
	}
	if args.PrevVersion == nil {
		return nil, errors.New("invalid value for required argument 'PrevVersion'")
	}
	if args.SourceURL == nil {
		return nil, errors.New("invalid value for required argument 'SourceURL'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	if args.DatabaseURL != nil {
		args.DatabaseURL = pulumi.ToSecret(args.DatabaseURL).(pulumi.StringOutput)
	}
	if args.SourceURL != nil {
		args.SourceURL = pulumi.ToSecret(args.SourceURL).(pulumi.StringOutput)
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Migration
	err := ctx.RegisterResource("gomigrate:index:Migration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMigration gets an existing Migration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMigration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrationState, opts ...pulumi.ResourceOption) (*Migration, error) {
	var resource Migration
	err := ctx.ReadResource("gomigrate:index:Migration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Migration resources.
type migrationState struct {
}

type MigrationState struct {
}

func (MigrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationState)(nil)).Elem()
}

type migrationArgs struct {
	// Database URL to run the migrations on
	DatabaseURL string `pulumi:"databaseURL"`
	// Previous version to migrate on undo
	PrevVersion int `pulumi:"prevVersion"`
	// Source URL for the migrations
	SourceURL string `pulumi:"sourceURL"`
	// Version to migrate
	Version int `pulumi:"version"`
}

// The set of arguments for constructing a Migration resource.
type MigrationArgs struct {
	// Database URL to run the migrations on
	DatabaseURL pulumi.StringInput
	// Previous version to migrate on undo
	PrevVersion pulumi.IntInput
	// Source URL for the migrations
	SourceURL pulumi.StringInput
	// Version to migrate
	Version pulumi.IntInput
}

func (MigrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationArgs)(nil)).Elem()
}

type MigrationInput interface {
	pulumi.Input

	ToMigrationOutput() MigrationOutput
	ToMigrationOutputWithContext(ctx context.Context) MigrationOutput
}

func (*Migration) ElementType() reflect.Type {
	return reflect.TypeOf((**Migration)(nil)).Elem()
}

func (i *Migration) ToMigrationOutput() MigrationOutput {
	return i.ToMigrationOutputWithContext(context.Background())
}

func (i *Migration) ToMigrationOutputWithContext(ctx context.Context) MigrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationOutput)
}

// MigrationArrayInput is an input type that accepts MigrationArray and MigrationArrayOutput values.
// You can construct a concrete instance of `MigrationArrayInput` via:
//
//          MigrationArray{ MigrationArgs{...} }
type MigrationArrayInput interface {
	pulumi.Input

	ToMigrationArrayOutput() MigrationArrayOutput
	ToMigrationArrayOutputWithContext(context.Context) MigrationArrayOutput
}

type MigrationArray []MigrationInput

func (MigrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Migration)(nil)).Elem()
}

func (i MigrationArray) ToMigrationArrayOutput() MigrationArrayOutput {
	return i.ToMigrationArrayOutputWithContext(context.Background())
}

func (i MigrationArray) ToMigrationArrayOutputWithContext(ctx context.Context) MigrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationArrayOutput)
}

// MigrationMapInput is an input type that accepts MigrationMap and MigrationMapOutput values.
// You can construct a concrete instance of `MigrationMapInput` via:
//
//          MigrationMap{ "key": MigrationArgs{...} }
type MigrationMapInput interface {
	pulumi.Input

	ToMigrationMapOutput() MigrationMapOutput
	ToMigrationMapOutputWithContext(context.Context) MigrationMapOutput
}

type MigrationMap map[string]MigrationInput

func (MigrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Migration)(nil)).Elem()
}

func (i MigrationMap) ToMigrationMapOutput() MigrationMapOutput {
	return i.ToMigrationMapOutputWithContext(context.Background())
}

func (i MigrationMap) ToMigrationMapOutputWithContext(ctx context.Context) MigrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationMapOutput)
}

type MigrationOutput struct{ *pulumi.OutputState }

func (MigrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Migration)(nil)).Elem()
}

func (o MigrationOutput) ToMigrationOutput() MigrationOutput {
	return o
}

func (o MigrationOutput) ToMigrationOutputWithContext(ctx context.Context) MigrationOutput {
	return o
}

type MigrationArrayOutput struct{ *pulumi.OutputState }

func (MigrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Migration)(nil)).Elem()
}

func (o MigrationArrayOutput) ToMigrationArrayOutput() MigrationArrayOutput {
	return o
}

func (o MigrationArrayOutput) ToMigrationArrayOutputWithContext(ctx context.Context) MigrationArrayOutput {
	return o
}

func (o MigrationArrayOutput) Index(i pulumi.IntInput) MigrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Migration {
		return vs[0].([]*Migration)[vs[1].(int)]
	}).(MigrationOutput)
}

type MigrationMapOutput struct{ *pulumi.OutputState }

func (MigrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Migration)(nil)).Elem()
}

func (o MigrationMapOutput) ToMigrationMapOutput() MigrationMapOutput {
	return o
}

func (o MigrationMapOutput) ToMigrationMapOutputWithContext(ctx context.Context) MigrationMapOutput {
	return o
}

func (o MigrationMapOutput) MapIndex(k pulumi.StringInput) MigrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Migration {
		return vs[0].(map[string]*Migration)[vs[1].(string)]
	}).(MigrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MigrationInput)(nil)).Elem(), &Migration{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrationArrayInput)(nil)).Elem(), MigrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrationMapInput)(nil)).Elem(), MigrationMap{})
	pulumi.RegisterOutputType(MigrationOutput{})
	pulumi.RegisterOutputType(MigrationArrayOutput{})
	pulumi.RegisterOutputType(MigrationMapOutput{})
}
