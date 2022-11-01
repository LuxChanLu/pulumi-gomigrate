package provider

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiProvider "github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	pbempty "github.com/golang/protobuf/ptypes/empty"
)

type gomigrateProvider struct {
	host    *provider.HostClient
	name    string
	version string
	schema  []byte
}

func makeProvider(host *provider.HostClient, name, version string, pulumiSchema []byte) (pulumirpc.ResourceProviderServer, error) {
	// Return the new provider
	return &gomigrateProvider{
		host:    host,
		name:    name,
		version: version,
		schema:  pulumiSchema,
	}, nil
}

// Attach sends the engine address to an already running plugin.
func (p *gomigrateProvider) Attach(context context.Context, req *pulumirpc.PluginAttach) (*emptypb.Empty, error) {
	host, err := provider.NewHostClient(req.GetAddress())
	if err != nil {
		return nil, err
	}
	p.host = host
	return &pbempty.Empty{}, nil
}

// Call dynamically executes a method in the provider associated with a component resource.
func (p *gomigrateProvider) Call(ctx context.Context, req *pulumirpc.CallRequest) (*pulumirpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "call is not yet implemented")
}

// Configure configures the resource provider with "globals" that control its behavior.
func (p *gomigrateProvider) Configure(_ context.Context, req *pulumirpc.ConfigureRequest) (*pulumirpc.ConfigureResponse, error) {
	return &pulumirpc.ConfigureResponse{
		AcceptSecrets:   true,
		AcceptResources: true,
		AcceptOutputs:   true,
		SupportsPreview: true,
	}, nil
}

// CheckConfig validates the configuration for this provider.
func (p *gomigrateProvider) CheckConfig(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (p *gomigrateProvider) DiffConfig(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	return &pulumirpc.DiffResponse{Changes: pulumirpc.DiffResponse_DIFF_NONE, Replaces: []string{}}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (p *gomigrateProvider) Invoke(ctx context.Context, req *pulumirpc.InvokeRequest) (*pulumirpc.InvokeResponse, error) {
	tok := req.GetTok()
	return nil, fmt.Errorf("unknown Invoke token '%s'", tok)
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (p *gomigrateProvider) StreamInvoke(req *pulumirpc.InvokeRequest, server pulumirpc.ResourceProvider_StreamInvokeServer) error {
	tok := req.GetTok()
	return fmt.Errorf("unknown StreamInvoke token '%s'", tok)
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are using for detecting and rendering diffs.
func (p *gomigrateProvider) Check(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (p *gomigrateProvider) Diff(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	d := olds.Diff(news)
	changes := pulumirpc.DiffResponse_DIFF_NONE

	if d.AnyChanges() {
		changes = pulumirpc.DiffResponse_DIFF_SOME
	}

	return &pulumirpc.DiffResponse{
		Changes:  changes,
		Replaces: []string{},
	}, nil
}

// Construct creates a new component resource.
func (p *gomigrateProvider) Construct(ctx context.Context, req *pulumirpc.ConstructRequest) (*pulumirpc.ConstructResponse, error) {
	return pulumiProvider.Construct(ctx, req, p.host.EngineConn(), func(ctx *pulumi.Context, typ, name string, inputs pulumiProvider.ConstructInputs, options pulumi.ResourceOption) (*pulumiProvider.ConstructResult, error) {
		var component pulumi.ComponentResource
		fields, err := inputs.Map()
		if err != nil {
			return nil, err
		}
		switch typ {
		case "gomigrate:index:Migrations":
			component, err = p.newMigrations(ctx, name, fields["sourceURL"], fields["databaseURL"])
			if err != nil {
				return nil, err
			}
		default:
			return nil, status.Error(codes.Unimplemented, fmt.Sprintf("%s does not exist", typ))
		}
		return pulumiProvider.NewConstructResult(component)
	})
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (p *gomigrateProvider) Create(ctx context.Context, req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	id := ""

	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	outputs := map[string]interface{}{}

	switch urn.Type() {
	case "gomigrate:index:Migration":
		id, err = p.runMigration(ctx, inputs, req.GetPreview(), outputs)
		if err != nil {
			return nil, err
		}
	default:
		return nil, status.Error(codes.Unimplemented, fmt.Sprintf("%s does not exist", urn.Type()))
	}

	outputProperties, err := plugin.MarshalProperties(resource.NewPropertyMapFromMap(outputs), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}
	return &pulumirpc.CreateResponse{Id: id, Properties: outputProperties}, nil
}

// Read the current live state associated with a resource.
func (p *gomigrateProvider) Read(ctx context.Context, req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	var err error
	urn := resource.URN(req.GetUrn())
	id := ""

	outputs := map[string]interface{}{}

	switch urn.Type() {
	case "gomigrate:index:Migration":
		outputs["migratedAt"] = req.Properties.GetFields()["migratedAt"]
	default:
		return nil, status.Error(codes.Unimplemented, fmt.Sprintf("%s does not exist", urn.Type()))
	}
	outputProperties, err := plugin.MarshalProperties(resource.NewPropertyMapFromMap(outputs), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}
	return &pulumirpc.ReadResponse{Id: id, Properties: outputProperties}, nil
}

// Update updates an existing resource with new values.
func (p *gomigrateProvider) Update(ctx context.Context, req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	return nil, status.Error(codes.Unimplemented, "update is not yet implemented")
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed
// to still exist.
func (p *gomigrateProvider) Delete(ctx context.Context, req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())

	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	switch urn.Type() {
	case "gomigrate:index:Migration":
		return &pbempty.Empty{}, p.undoMigration(ctx, inputs)
	default:
		return nil, status.Error(codes.Unimplemented, fmt.Sprintf("%s does not exist", urn.Type()))
	}
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (p *gomigrateProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*pulumirpc.PluginInfo, error) {
	return &pulumirpc.PluginInfo{Version: p.version}, nil
}

// GetSchema returns the JSON-serialized schema for the provider.
func (p *gomigrateProvider) GetSchema(ctx context.Context, req *pulumirpc.GetSchemaRequest) (*pulumirpc.GetSchemaResponse, error) {
	if v := req.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}
	return &pulumirpc.GetSchemaResponse{Schema: string(p.schema)}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (p *gomigrateProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	return &pbempty.Empty{}, nil
}
