// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: taskService.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TaskService service

type TaskService interface {
	CreateTask(ctx context.Context, in *TaskRequest, opts ...client.CallOption) (*TaskDetailResponse, error)
	GetTasksList(ctx context.Context, in *TaskRequest, opts ...client.CallOption) (*TaskListResponse, error)
	GetTask(ctx context.Context, in *TaskRequest, opts ...client.CallOption) (*TaskDetailResponse, error)
	UpdateTask(ctx context.Context, in *TaskRequest, opts ...client.CallOption) (*TaskDetailResponse, error)
	DeleteTask(ctx context.Context, in *TaskRequest, opts ...client.CallOption) (*TaskDetailResponse, error)
}

type taskService struct {
	c    client.Client
	name string
}

func NewTaskService(name string, c client.Client) TaskService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "services"
	}
	return &taskService{
		c:    c,
		name: name,
	}
}

func (c *taskService) CreateTask(ctx context.Context, in *TaskRequest, opts ...client.CallOption) (*TaskDetailResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.CreateTask", in)
	out := new(TaskDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetTasksList(ctx context.Context, in *TaskRequest, opts ...client.CallOption) (*TaskListResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetTasksList", in)
	out := new(TaskListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetTask(ctx context.Context, in *TaskRequest, opts ...client.CallOption) (*TaskDetailResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetTask", in)
	out := new(TaskDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) UpdateTask(ctx context.Context, in *TaskRequest, opts ...client.CallOption) (*TaskDetailResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.UpdateTask", in)
	out := new(TaskDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) DeleteTask(ctx context.Context, in *TaskRequest, opts ...client.CallOption) (*TaskDetailResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.DeleteTask", in)
	out := new(TaskDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskService service

type TaskServiceHandler interface {
	CreateTask(context.Context, *TaskRequest, *TaskDetailResponse) error
	GetTasksList(context.Context, *TaskRequest, *TaskListResponse) error
	GetTask(context.Context, *TaskRequest, *TaskDetailResponse) error
	UpdateTask(context.Context, *TaskRequest, *TaskDetailResponse) error
	DeleteTask(context.Context, *TaskRequest, *TaskDetailResponse) error
}

func RegisterTaskServiceHandler(s server.Server, hdlr TaskServiceHandler, opts ...server.HandlerOption) error {
	type taskService interface {
		CreateTask(ctx context.Context, in *TaskRequest, out *TaskDetailResponse) error
		GetTasksList(ctx context.Context, in *TaskRequest, out *TaskListResponse) error
		GetTask(ctx context.Context, in *TaskRequest, out *TaskDetailResponse) error
		UpdateTask(ctx context.Context, in *TaskRequest, out *TaskDetailResponse) error
		DeleteTask(ctx context.Context, in *TaskRequest, out *TaskDetailResponse) error
	}
	type TaskService struct {
		taskService
	}
	h := &taskServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TaskService{h}, opts...))
}

type taskServiceHandler struct {
	TaskServiceHandler
}

func (h *taskServiceHandler) CreateTask(ctx context.Context, in *TaskRequest, out *TaskDetailResponse) error {
	return h.TaskServiceHandler.CreateTask(ctx, in, out)
}

func (h *taskServiceHandler) GetTasksList(ctx context.Context, in *TaskRequest, out *TaskListResponse) error {
	return h.TaskServiceHandler.GetTasksList(ctx, in, out)
}

func (h *taskServiceHandler) GetTask(ctx context.Context, in *TaskRequest, out *TaskDetailResponse) error {
	return h.TaskServiceHandler.GetTask(ctx, in, out)
}

func (h *taskServiceHandler) UpdateTask(ctx context.Context, in *TaskRequest, out *TaskDetailResponse) error {
	return h.TaskServiceHandler.UpdateTask(ctx, in, out)
}

func (h *taskServiceHandler) DeleteTask(ctx context.Context, in *TaskRequest, out *TaskDetailResponse) error {
	return h.TaskServiceHandler.DeleteTask(ctx, in, out)
}
