// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/cloud/ml/v1beta1/model_service.proto
// DO NOT EDIT!

package google_cloud_ml_v1beta1 // import "google.golang.org/genproto/googleapis/cloud/ml/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_longrunning "google.golang.org/genproto/googleapis/longrunning"
import _ "github.com/golang/protobuf/ptypes/empty"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents a machine learning solution.
//
// A model can have multiple versions, each of which is a deployed, trained
// model ready to receive prediction requests. The model itself is just a
// container.
type Model struct {
	// Required. The name specified for the model when it was created.
	//
	// The model name must be unique within the project it is created in.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Optional. The description specified for the model when it was created.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Output only. The default version of the model. This version will be used to
	// handle prediction requests that do not specify a version.
	//
	// You can change the default version by calling
	// [projects.methods.versions.setDefault](/ml/reference/rest/v1beta1/projects.models.versions/setDefault).
	DefaultVersion *Version `protobuf:"bytes,3,opt,name=default_version,json=defaultVersion" json:"default_version,omitempty"`
}

func (m *Model) Reset()                    { *m = Model{} }
func (m *Model) String() string            { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()               {}
func (*Model) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Model) GetDefaultVersion() *Version {
	if m != nil {
		return m.DefaultVersion
	}
	return nil
}

// Represents a version of the model.
//
// Each version is a trained model deployed in the cloud, ready to handle
// prediction requests. A model can have multiple versions. You can get
// information about all of the versions of a given model by calling
// [projects.models.versions.list](/ml/reference/rest/v1beta1/projects.models.versions/list).
type Version struct {
	// Required.The name specified for the version when it was created.
	//
	// The version name must be unique within the model it is created in.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Optional. The description specified for the version when it was created.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Output only. If true, this version will be used to handle prediction
	// requests that do not specify a version.
	//
	// You can change the default version by calling
	// [projects.methods.versions.setDefault](/ml/reference/rest/v1beta1/projects.models.versions/setDefault).
	IsDefault bool `protobuf:"varint,3,opt,name=is_default,json=isDefault" json:"is_default,omitempty"`
	// Required. The Google Cloud Storage location of the trained model used to
	// create the version. See the
	// [overview of model deployment](/ml/docs/concepts/deployment-overview) for
	// more informaiton.
	//
	// When passing Version to
	// [projects.models.versions.create](/ml/reference/rest/v1beta1/projects.models.versions/create)
	// the model service uses the specified location as the source of the model.
	// Once deployed, the model version is hosted by the prediction service, so
	// this location is useful only as a historical record.
	DeploymentUri string `protobuf:"bytes,4,opt,name=deployment_uri,json=deploymentUri" json:"deployment_uri,omitempty"`
	// Output only. The time the version was created.
	CreateTime *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// Output only. The time the version was last used for prediction.
	LastUseTime *google_protobuf2.Timestamp `protobuf:"bytes,6,opt,name=last_use_time,json=lastUseTime" json:"last_use_time,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Version) GetCreateTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Version) GetLastUseTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LastUseTime
	}
	return nil
}

// Request message for the CreateModel method.
type CreateModelRequest struct {
	// Required. The project name.
	//
	// Authorization: requires `Editor` role on the specified project.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Required. The model to create.
	Model *Model `protobuf:"bytes,2,opt,name=model" json:"model,omitempty"`
}

func (m *CreateModelRequest) Reset()                    { *m = CreateModelRequest{} }
func (m *CreateModelRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateModelRequest) ProtoMessage()               {}
func (*CreateModelRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *CreateModelRequest) GetModel() *Model {
	if m != nil {
		return m.Model
	}
	return nil
}

// Request message for the ListModels method.
type ListModelsRequest struct {
	// Required. The name of the project whose models are to be listed.
	//
	// Authorization: requires `Viewer` role on the specified project.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. A page token to request the next page of results.
	//
	// You get the token from the `next_page_token` field of the response from
	// the previous call.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// Optional. The number of models to retrieve per "page" of results. If there
	// are more remaining results than this number, the response message will
	// contain a valid value in the `next_page_token` field.
	//
	// The default value is 20, and the maximum page size is 100.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListModelsRequest) Reset()                    { *m = ListModelsRequest{} }
func (m *ListModelsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListModelsRequest) ProtoMessage()               {}
func (*ListModelsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// Response message for the ListModels method.
type ListModelsResponse struct {
	// The list of models.
	Models []*Model `protobuf:"bytes,1,rep,name=models" json:"models,omitempty"`
	// Optional. Pass this token as the `page_token` field of the request for a
	// subsequent call.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListModelsResponse) Reset()                    { *m = ListModelsResponse{} }
func (m *ListModelsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListModelsResponse) ProtoMessage()               {}
func (*ListModelsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ListModelsResponse) GetModels() []*Model {
	if m != nil {
		return m.Models
	}
	return nil
}

// Request message for the GetModel method.
type GetModelRequest struct {
	// Required. The name of the model.
	//
	// Authorization: requires `Viewer` role on the parent project.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetModelRequest) Reset()                    { *m = GetModelRequest{} }
func (m *GetModelRequest) String() string            { return proto.CompactTextString(m) }
func (*GetModelRequest) ProtoMessage()               {}
func (*GetModelRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

// Request message for the DeleteModel method.
type DeleteModelRequest struct {
	// Required. The name of the model.
	//
	// Authorization: requires `Editor` role on the parent project.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteModelRequest) Reset()                    { *m = DeleteModelRequest{} }
func (m *DeleteModelRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteModelRequest) ProtoMessage()               {}
func (*DeleteModelRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

// Uploads the provided trained model version to Cloud Machine Learning.
type CreateVersionRequest struct {
	// Required. The name of the model.
	//
	// Authorization: requires `Editor` role on the parent project.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Required. The version details.
	Version *Version `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *CreateVersionRequest) Reset()                    { *m = CreateVersionRequest{} }
func (m *CreateVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateVersionRequest) ProtoMessage()               {}
func (*CreateVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *CreateVersionRequest) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

// Request message for the ListVersions method.
type ListVersionsRequest struct {
	// Required. The name of the model for which to list the version.
	//
	// Authorization: requires `Viewer` role on the parent project.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. A page token to request the next page of results.
	//
	// You get the token from the `next_page_token` field of the response from
	// the previous call.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// Optional. The number of versions to retrieve per "page" of results. If
	// there are more remaining results than this number, the response message
	// will contain a valid value in the `next_page_token` field.
	//
	// The default value is 20, and the maximum page size is 100.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListVersionsRequest) Reset()                    { *m = ListVersionsRequest{} }
func (m *ListVersionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListVersionsRequest) ProtoMessage()               {}
func (*ListVersionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

// Response message for the ListVersions method.
type ListVersionsResponse struct {
	// The list of versions.
	Versions []*Version `protobuf:"bytes,1,rep,name=versions" json:"versions,omitempty"`
	// Optional. Pass this token as the `page_token` field of the request for a
	// subsequent call.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListVersionsResponse) Reset()                    { *m = ListVersionsResponse{} }
func (m *ListVersionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListVersionsResponse) ProtoMessage()               {}
func (*ListVersionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *ListVersionsResponse) GetVersions() []*Version {
	if m != nil {
		return m.Versions
	}
	return nil
}

// Request message for the GetVersion method.
type GetVersionRequest struct {
	// Required. The name of the version.
	//
	// Authorization: requires `Viewer` role on the parent project.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetVersionRequest) Reset()                    { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()               {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

// Request message for the DeleteVerionRequest method.
type DeleteVersionRequest struct {
	// Required. The name of the version. You can get the names of all the
	// versions of a model by calling
	// [projects.models.versions.list](/ml/reference/rest/v1beta1/projects.models.versions/list).
	//
	// Authorization: requires `Editor` role on the parent project.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteVersionRequest) Reset()                    { *m = DeleteVersionRequest{} }
func (m *DeleteVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteVersionRequest) ProtoMessage()               {}
func (*DeleteVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

// Request message for the SetDefaultVersion request.
type SetDefaultVersionRequest struct {
	// Required. The name of the version to make the default for the model. You
	// can get the names of all the versions of a model by calling
	// [projects.models.versions.list](/ml/reference/rest/v1beta1/projects.models.versions/list).
	//
	// Authorization: requires `Editor` role on the parent project.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *SetDefaultVersionRequest) Reset()                    { *m = SetDefaultVersionRequest{} }
func (m *SetDefaultVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*SetDefaultVersionRequest) ProtoMessage()               {}
func (*SetDefaultVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func init() {
	proto.RegisterType((*Model)(nil), "google.cloud.ml.v1beta1.Model")
	proto.RegisterType((*Version)(nil), "google.cloud.ml.v1beta1.Version")
	proto.RegisterType((*CreateModelRequest)(nil), "google.cloud.ml.v1beta1.CreateModelRequest")
	proto.RegisterType((*ListModelsRequest)(nil), "google.cloud.ml.v1beta1.ListModelsRequest")
	proto.RegisterType((*ListModelsResponse)(nil), "google.cloud.ml.v1beta1.ListModelsResponse")
	proto.RegisterType((*GetModelRequest)(nil), "google.cloud.ml.v1beta1.GetModelRequest")
	proto.RegisterType((*DeleteModelRequest)(nil), "google.cloud.ml.v1beta1.DeleteModelRequest")
	proto.RegisterType((*CreateVersionRequest)(nil), "google.cloud.ml.v1beta1.CreateVersionRequest")
	proto.RegisterType((*ListVersionsRequest)(nil), "google.cloud.ml.v1beta1.ListVersionsRequest")
	proto.RegisterType((*ListVersionsResponse)(nil), "google.cloud.ml.v1beta1.ListVersionsResponse")
	proto.RegisterType((*GetVersionRequest)(nil), "google.cloud.ml.v1beta1.GetVersionRequest")
	proto.RegisterType((*DeleteVersionRequest)(nil), "google.cloud.ml.v1beta1.DeleteVersionRequest")
	proto.RegisterType((*SetDefaultVersionRequest)(nil), "google.cloud.ml.v1beta1.SetDefaultVersionRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ModelService service

type ModelServiceClient interface {
	// Creates a model which will later contain one or more versions.
	//
	// You must add at least one version before you can request predictions from
	// the model. Add versions by calling
	// [projects.models.versions.create](/ml/reference/rest/v1beta1/projects.models.versions/create).
	CreateModel(ctx context.Context, in *CreateModelRequest, opts ...grpc.CallOption) (*Model, error)
	// Lists the models in a project.
	//
	// Each project can contain multiple models, and each model can have multiple
	// versions.
	ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error)
	// Gets information about a model, including its name, the description (if
	// set), and the default version (if at least one version of the model has
	// been deployed).
	GetModel(ctx context.Context, in *GetModelRequest, opts ...grpc.CallOption) (*Model, error)
	// Deletes a model.
	//
	// You can only delete a model if there are no versions in it. You can delete
	// versions by calling
	// [projects.models.versions.delete](/ml/reference/rest/v1beta1/projects.models.versions/delete).
	DeleteModel(ctx context.Context, in *DeleteModelRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error)
	// Creates a new version of a model from a trained TensorFlow model.
	//
	// If the version created in the cloud by this call is the first deployed
	// version of the specified model, it will be made the default version of the
	// model. When you add a version to a model that already has one or more
	// versions, the default version does not automatically change. If you want a
	// new version to be the default, you must call
	// [projects.models.versions.setDefault](/ml/reference/rest/v1beta1/projects.models.versions/setDefault).
	CreateVersion(ctx context.Context, in *CreateVersionRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error)
	// Gets basic information about all the versions of a model.
	//
	// If you expect that a model has a lot of versions, or if you need to handle
	// only a limited number of results at a time, you can request that the list
	// be retrieved in batches (called pages):
	ListVersions(ctx context.Context, in *ListVersionsRequest, opts ...grpc.CallOption) (*ListVersionsResponse, error)
	// Gets information about a model version.
	//
	// Models can have multiple versions. You can call
	// [projects.models.versions.list](/ml/reference/rest/v1beta1/projects.models.versions/list)
	// to get the same information that this method returns for all of the
	// versions of a model.
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*Version, error)
	// Deletes a model version.
	//
	// Each model can have multiple versions deployed and in use at any given
	// time. Use this method to remove a single version.
	//
	// Note: You cannot delete the version that is set as the default version
	// of the model unless it is the only remaining version.
	DeleteVersion(ctx context.Context, in *DeleteVersionRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error)
	// Designates a version to be the default for the model.
	//
	// The default version is used for prediction requests made against the model
	// that don't specify a version.
	//
	// The first version to be created for a model is automatically set as the
	// default. You must make any subsequent changes to the default version
	// setting manually using this method.
	SetDefaultVersion(ctx context.Context, in *SetDefaultVersionRequest, opts ...grpc.CallOption) (*Version, error)
}

type modelServiceClient struct {
	cc *grpc.ClientConn
}

func NewModelServiceClient(cc *grpc.ClientConn) ModelServiceClient {
	return &modelServiceClient{cc}
}

func (c *modelServiceClient) CreateModel(ctx context.Context, in *CreateModelRequest, opts ...grpc.CallOption) (*Model, error) {
	out := new(Model)
	err := grpc.Invoke(ctx, "/google.cloud.ml.v1beta1.ModelService/CreateModel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error) {
	out := new(ListModelsResponse)
	err := grpc.Invoke(ctx, "/google.cloud.ml.v1beta1.ModelService/ListModels", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) GetModel(ctx context.Context, in *GetModelRequest, opts ...grpc.CallOption) (*Model, error) {
	out := new(Model)
	err := grpc.Invoke(ctx, "/google.cloud.ml.v1beta1.ModelService/GetModel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) DeleteModel(ctx context.Context, in *DeleteModelRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error) {
	out := new(google_longrunning.Operation)
	err := grpc.Invoke(ctx, "/google.cloud.ml.v1beta1.ModelService/DeleteModel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) CreateVersion(ctx context.Context, in *CreateVersionRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error) {
	out := new(google_longrunning.Operation)
	err := grpc.Invoke(ctx, "/google.cloud.ml.v1beta1.ModelService/CreateVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) ListVersions(ctx context.Context, in *ListVersionsRequest, opts ...grpc.CallOption) (*ListVersionsResponse, error) {
	out := new(ListVersionsResponse)
	err := grpc.Invoke(ctx, "/google.cloud.ml.v1beta1.ModelService/ListVersions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := grpc.Invoke(ctx, "/google.cloud.ml.v1beta1.ModelService/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) DeleteVersion(ctx context.Context, in *DeleteVersionRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error) {
	out := new(google_longrunning.Operation)
	err := grpc.Invoke(ctx, "/google.cloud.ml.v1beta1.ModelService/DeleteVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) SetDefaultVersion(ctx context.Context, in *SetDefaultVersionRequest, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := grpc.Invoke(ctx, "/google.cloud.ml.v1beta1.ModelService/SetDefaultVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ModelService service

type ModelServiceServer interface {
	// Creates a model which will later contain one or more versions.
	//
	// You must add at least one version before you can request predictions from
	// the model. Add versions by calling
	// [projects.models.versions.create](/ml/reference/rest/v1beta1/projects.models.versions/create).
	CreateModel(context.Context, *CreateModelRequest) (*Model, error)
	// Lists the models in a project.
	//
	// Each project can contain multiple models, and each model can have multiple
	// versions.
	ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error)
	// Gets information about a model, including its name, the description (if
	// set), and the default version (if at least one version of the model has
	// been deployed).
	GetModel(context.Context, *GetModelRequest) (*Model, error)
	// Deletes a model.
	//
	// You can only delete a model if there are no versions in it. You can delete
	// versions by calling
	// [projects.models.versions.delete](/ml/reference/rest/v1beta1/projects.models.versions/delete).
	DeleteModel(context.Context, *DeleteModelRequest) (*google_longrunning.Operation, error)
	// Creates a new version of a model from a trained TensorFlow model.
	//
	// If the version created in the cloud by this call is the first deployed
	// version of the specified model, it will be made the default version of the
	// model. When you add a version to a model that already has one or more
	// versions, the default version does not automatically change. If you want a
	// new version to be the default, you must call
	// [projects.models.versions.setDefault](/ml/reference/rest/v1beta1/projects.models.versions/setDefault).
	CreateVersion(context.Context, *CreateVersionRequest) (*google_longrunning.Operation, error)
	// Gets basic information about all the versions of a model.
	//
	// If you expect that a model has a lot of versions, or if you need to handle
	// only a limited number of results at a time, you can request that the list
	// be retrieved in batches (called pages):
	ListVersions(context.Context, *ListVersionsRequest) (*ListVersionsResponse, error)
	// Gets information about a model version.
	//
	// Models can have multiple versions. You can call
	// [projects.models.versions.list](/ml/reference/rest/v1beta1/projects.models.versions/list)
	// to get the same information that this method returns for all of the
	// versions of a model.
	GetVersion(context.Context, *GetVersionRequest) (*Version, error)
	// Deletes a model version.
	//
	// Each model can have multiple versions deployed and in use at any given
	// time. Use this method to remove a single version.
	//
	// Note: You cannot delete the version that is set as the default version
	// of the model unless it is the only remaining version.
	DeleteVersion(context.Context, *DeleteVersionRequest) (*google_longrunning.Operation, error)
	// Designates a version to be the default for the model.
	//
	// The default version is used for prediction requests made against the model
	// that don't specify a version.
	//
	// The first version to be created for a model is automatically set as the
	// default. You must make any subsequent changes to the default version
	// setting manually using this method.
	SetDefaultVersion(context.Context, *SetDefaultVersionRequest) (*Version, error)
}

func RegisterModelServiceServer(s *grpc.Server, srv ModelServiceServer) {
	s.RegisterService(&_ModelService_serviceDesc, srv)
}

func _ModelService_CreateModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).CreateModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1beta1.ModelService/CreateModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).CreateModel(ctx, req.(*CreateModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_ListModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).ListModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1beta1.ModelService/ListModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).ListModels(ctx, req.(*ListModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_GetModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).GetModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1beta1.ModelService/GetModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).GetModel(ctx, req.(*GetModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_DeleteModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).DeleteModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1beta1.ModelService/DeleteModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).DeleteModel(ctx, req.(*DeleteModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_CreateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).CreateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1beta1.ModelService/CreateVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).CreateVersion(ctx, req.(*CreateVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_ListVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).ListVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1beta1.ModelService/ListVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).ListVersions(ctx, req.(*ListVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1beta1.ModelService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_DeleteVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).DeleteVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1beta1.ModelService/DeleteVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).DeleteVersion(ctx, req.(*DeleteVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_SetDefaultVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDefaultVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).SetDefaultVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1beta1.ModelService/SetDefaultVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).SetDefaultVersion(ctx, req.(*SetDefaultVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.ml.v1beta1.ModelService",
	HandlerType: (*ModelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateModel",
			Handler:    _ModelService_CreateModel_Handler,
		},
		{
			MethodName: "ListModels",
			Handler:    _ModelService_ListModels_Handler,
		},
		{
			MethodName: "GetModel",
			Handler:    _ModelService_GetModel_Handler,
		},
		{
			MethodName: "DeleteModel",
			Handler:    _ModelService_DeleteModel_Handler,
		},
		{
			MethodName: "CreateVersion",
			Handler:    _ModelService_CreateVersion_Handler,
		},
		{
			MethodName: "ListVersions",
			Handler:    _ModelService_ListVersions_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ModelService_GetVersion_Handler,
		},
		{
			MethodName: "DeleteVersion",
			Handler:    _ModelService_DeleteVersion_Handler,
		},
		{
			MethodName: "SetDefaultVersion",
			Handler:    _ModelService_SetDefaultVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/cloud/ml/v1beta1/model_service.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 915 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0x4d, 0x6f, 0xf3, 0x44,
	0x10, 0x96, 0xf3, 0xbe, 0xc9, 0x9b, 0x4e, 0x1a, 0xaa, 0x2c, 0x15, 0x44, 0x46, 0x85, 0xca, 0xa8,
	0x50, 0xa5, 0xad, 0x4d, 0x53, 0x54, 0xa9, 0x29, 0x14, 0xa9, 0x44, 0x42, 0x45, 0x20, 0xaa, 0xb4,
	0xe5, 0x1a, 0x39, 0xc9, 0xd6, 0xb8, 0xd8, 0x5e, 0xe3, 0xdd, 0x54, 0x94, 0x8f, 0x03, 0x3d, 0x20,
	0x4e, 0x1c, 0x40, 0x5c, 0xb9, 0x70, 0xe7, 0xcf, 0xf0, 0x17, 0x10, 0xbf, 0x83, 0xf5, 0xee, 0x3a,
	0x71, 0x3e, 0x1c, 0xbb, 0x48, 0x5c, 0x2c, 0x7b, 0x3c, 0x1f, 0x8f, 0x9f, 0x67, 0x66, 0xbc, 0xf0,
	0xb1, 0x43, 0x88, 0xe3, 0x61, 0xd3, 0x21, 0x9e, 0x1d, 0x38, 0x26, 0x89, 0x1c, 0xcb, 0xc1, 0x41,
	0x18, 0x11, 0x46, 0x2c, 0xf9, 0xca, 0x0e, 0x5d, 0x6a, 0x0d, 0x3d, 0x32, 0x1e, 0x59, 0xbe, 0x67,
	0xdd, 0x1f, 0x0e, 0x30, 0xb3, 0x0f, 0x2d, 0x9f, 0x8c, 0xb0, 0xd7, 0xa7, 0x38, 0xba, 0x77, 0x87,
	0xd8, 0x14, 0xfe, 0xe8, 0x55, 0x95, 0x4b, 0x38, 0x9b, 0xbe, 0x67, 0x2a, 0x67, 0xfd, 0xa2, 0x58,
	0x11, 0x7e, 0xb1, 0x54, 0xc6, 0x21, 0x09, 0x6e, 0x5d, 0xc7, 0xb2, 0x83, 0x80, 0x30, 0x9b, 0xb9,
	0x24, 0xa0, 0xb2, 0x86, 0x7e, 0x5e, 0x2c, 0x95, 0x47, 0x02, 0x27, 0x1a, 0x07, 0x81, 0x1b, 0x38,
	0x16, 0x09, 0x71, 0x34, 0x93, 0xe3, 0xc8, 0x71, 0xd9, 0x17, 0xe3, 0x81, 0x39, 0x24, 0xbe, 0x25,
	0xf3, 0x58, 0xe2, 0xc5, 0x60, 0x7c, 0x6b, 0x85, 0xec, 0x21, 0xc4, 0xd4, 0xc2, 0x3e, 0xbf, 0x91,
	0x57, 0x15, 0x74, 0x9a, 0x1f, 0xc4, 0x5c, 0x1f, 0x53, 0x66, 0xfb, 0xe1, 0xf4, 0x4e, 0x06, 0x1b,
	0x3f, 0x69, 0x50, 0xfe, 0x34, 0x66, 0x0c, 0x21, 0x78, 0x1e, 0xd8, 0x3e, 0x6e, 0x6a, 0xdb, 0xda,
	0xee, 0x5a, 0x4f, 0xdc, 0xa3, 0x6d, 0xa8, 0x8d, 0x30, 0x1d, 0x46, 0x6e, 0x18, 0xa3, 0x6c, 0x96,
	0xc4, 0xab, 0xb4, 0x09, 0x5d, 0xc0, 0xc6, 0x08, 0xdf, 0xda, 0x63, 0x8f, 0xf5, 0xef, 0x71, 0x44,
	0x63, 0xaf, 0x67, 0xdc, 0xab, 0xd6, 0xde, 0x36, 0x33, 0x38, 0x37, 0x3f, 0x97, 0x7e, 0xbd, 0x97,
	0x54, 0xa0, 0x7a, 0x36, 0x1e, 0x4b, 0xf0, 0x42, 0xdd, 0xff, 0x47, 0x30, 0x5b, 0x00, 0x2e, 0xed,
	0xab, 0xb4, 0x02, 0x47, 0xb5, 0xb7, 0xe6, 0xd2, 0xae, 0x34, 0xa0, 0x1d, 0xe0, 0x25, 0x43, 0x8f,
	0x3c, 0xf8, 0x38, 0x60, 0xfd, 0x71, 0xe4, 0x36, 0x9f, 0x8b, 0x1c, 0xf5, 0xa9, 0xf5, 0x26, 0x72,
	0xd1, 0x29, 0xd4, 0x86, 0x11, 0xb6, 0x19, 0xee, 0xc7, 0x64, 0x35, 0xcb, 0xe2, 0x73, 0xf4, 0xe4,
	0x73, 0x12, 0x6a, 0xcd, 0xeb, 0x84, 0xc9, 0x1e, 0x48, 0xf7, 0xd8, 0x80, 0xce, 0xa0, 0xee, 0xd9,
	0x94, 0x67, 0xa7, 0x2a, 0xbc, 0x92, 0x1b, 0x5e, 0x8b, 0x03, 0x6e, 0xa8, 0x88, 0x37, 0x06, 0x80,
	0x3e, 0x14, 0xd9, 0x84, 0x28, 0x3d, 0xfc, 0xd5, 0x98, 0x7b, 0xa1, 0x57, 0xa0, 0x12, 0xda, 0x11,
	0xc7, 0xa7, 0x08, 0x51, 0x4f, 0xe8, 0x5d, 0x28, 0x8b, 0x76, 0x17, 0x64, 0xd4, 0xda, 0xaf, 0x67,
	0x72, 0x2e, 0xb3, 0x49, 0x67, 0xc3, 0x81, 0xc6, 0x27, 0x2e, 0x65, 0xc2, 0x46, 0xf3, 0x4a, 0x70,
	0x4e, 0x43, 0xdb, 0xe1, 0x1f, 0x43, 0xbe, 0xc4, 0x81, 0x22, 0x6c, 0x2d, 0xb6, 0x5c, 0xc7, 0x06,
	0xf4, 0x1a, 0x88, 0x87, 0x3e, 0x75, 0xbf, 0x91, 0x54, 0x95, 0x7b, 0xd5, 0xd8, 0x70, 0xc5, 0x9f,
	0x0d, 0x06, 0x28, 0x5d, 0x88, 0x86, 0xbc, 0xd3, 0x31, 0x3a, 0x86, 0x8a, 0xc0, 0x41, 0x79, 0xa5,
	0x67, 0x05, 0x50, 0x2b, 0x6f, 0xf4, 0x16, 0x6c, 0x04, 0xf8, 0x6b, 0xd6, 0x4f, 0xc1, 0x91, 0x3d,
	0x50, 0x8f, 0xcd, 0x97, 0x09, 0x24, 0x63, 0x07, 0x36, 0x3e, 0xc2, 0x6c, 0x86, 0xbf, 0x25, 0xed,
	0x64, 0xec, 0x02, 0xea, 0x62, 0x0f, 0xcf, 0x31, 0xbd, 0xcc, 0xf3, 0x0e, 0x36, 0xa5, 0x26, 0x49,
	0xe7, 0xe6, 0x50, 0xd6, 0x81, 0x17, 0xc9, 0x2c, 0x94, 0x0a, 0xce, 0x42, 0x12, 0x60, 0xb8, 0xf0,
	0x72, 0x4c, 0x99, 0xb2, 0xff, 0xaf, 0xea, 0x7c, 0x07, 0x9b, 0xb3, 0xa5, 0x94, 0x3e, 0xef, 0x41,
	0x55, 0xa1, 0x49, 0x14, 0xca, 0xc7, 0x3f, 0x89, 0x28, 0xac, 0xd2, 0xdb, 0xd0, 0xe0, 0x2a, 0xcd,
	0x31, 0xba, 0x8c, 0xfd, 0x16, 0x6c, 0x4a, 0x9d, 0x0a, 0xf8, 0x9a, 0xd0, 0xbc, 0xc2, 0xac, 0x3b,
	0xb3, 0x57, 0x56, 0xf8, 0xb7, 0xff, 0x01, 0x58, 0x17, 0xf2, 0x5f, 0xc9, 0xe5, 0x8e, 0x7e, 0xd6,
	0xa0, 0x96, 0x9a, 0x3f, 0xb4, 0x97, 0xf9, 0xe5, 0x8b, 0x53, 0xaa, 0xe7, 0x34, 0xb2, 0xd1, 0x7e,
	0xfc, 0xeb, 0xef, 0x5f, 0x4b, 0xfb, 0xc6, 0x9b, 0x93, 0x7f, 0xd5, 0xb7, 0x52, 0xc6, 0xf7, 0xf9,
	0x56, 0xb8, 0xc3, 0x43, 0x46, 0xad, 0xd6, 0xf7, 0xf2, 0xff, 0x45, 0x3b, 0x72, 0x56, 0xd1, 0x2f,
	0x1a, 0xc0, 0x74, 0x86, 0x50, 0x2b, 0xb3, 0xc4, 0xc2, 0x44, 0xeb, 0x7b, 0x85, 0x7c, 0xa5, 0xe8,
	0xc6, 0x9e, 0xc0, 0xb6, 0x83, 0x8a, 0x60, 0x43, 0x3f, 0x68, 0x50, 0x4d, 0x46, 0x0c, 0xed, 0x66,
	0x96, 0x99, 0x9b, 0xc2, 0x5c, 0x7e, 0x96, 0x60, 0x88, 0x55, 0x4a, 0x21, 0x50, 0x00, 0x38, 0x14,
	0xf4, 0x23, 0x57, 0x2a, 0x35, 0xbf, 0x2b, 0x94, 0x5a, 0x9c, 0x72, 0x7d, 0x2b, 0x71, 0x4e, 0xfd,
	0x8d, 0xcd, 0xcf, 0x92, 0xbf, 0x71, 0x02, 0xa4, 0x55, 0x08, 0xc8, 0xef, 0x1a, 0xd4, 0x67, 0xd6,
	0x03, 0x3a, 0xc8, 0x69, 0x9a, 0xd9, 0xc6, 0xcc, 0x03, 0xf3, 0x81, 0x00, 0x73, 0x62, 0x98, 0x2b,
	0x94, 0x99, 0xc2, 0xb1, 0x92, 0x41, 0xec, 0x24, 0x2b, 0x05, 0xfd, 0xa1, 0xc1, 0x7a, 0x7a, 0xd0,
	0xd1, 0xfe, 0xca, 0xc6, 0x98, 0x5b, 0x3d, 0xfa, 0x41, 0x41, 0x6f, 0xd5, 0x48, 0xc7, 0x02, 0xee,
	0x3b, 0xe8, 0x89, 0x70, 0x45, 0xa3, 0x4f, 0x17, 0xc2, 0x8a, 0x46, 0x5f, 0xd8, 0x1a, 0x7a, 0xee,
	0x7a, 0x5a, 0x06, 0x2a, 0x4b, 0xd0, 0x09, 0xa2, 0x58, 0xdb, 0xdf, 0xb8, 0xb6, 0x33, 0xcb, 0x67,
	0x85, 0xb6, 0xcb, 0x96, 0x54, 0x9e, 0xb6, 0x0a, 0x57, 0xeb, 0xa9, 0xb8, 0xfe, 0xd4, 0xa0, 0xb1,
	0xb0, 0xe8, 0xd0, 0x61, 0x26, 0xb6, 0xac, 0xa5, 0x58, 0x80, 0xba, 0xae, 0x80, 0x78, 0x66, 0x9c,
	0x3c, 0x0d, 0x62, 0x87, 0x4e, 0x4a, 0x76, 0xb4, 0xd6, 0xf9, 0x31, 0xbc, 0xc1, 0x8f, 0xa7, 0x0b,
	0xc5, 0xf8, 0x89, 0x38, 0x29, 0x78, 0xde, 0x48, 0x2f, 0xe2, 0xcb, 0xf8, 0x98, 0x74, 0xa9, 0x0d,
	0x2a, 0xe2, 0xbc, 0x74, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xdd, 0x81, 0x3c, 0x0d,
	0x0c, 0x00, 0x00,
}