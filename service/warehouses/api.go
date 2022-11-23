// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package warehouses

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

func NewQueryHistory(client *client.DatabricksClient) *QueryHistoryAPI {
	return &QueryHistoryAPI{
		impl: &queryHistoryImpl{
			client: client,
		},
	}
}

// Access the history of queries through SQL warehouses.
type QueryHistoryAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(QueryHistoryService)
	impl QueryHistoryService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *QueryHistoryAPI) WithImpl(impl QueryHistoryService) *QueryHistoryAPI {
	a.impl = impl
	return a
}

// Impl returns low-level QueryHistory API implementation
func (a *QueryHistoryAPI) Impl() QueryHistoryService {
	return a.impl
}

// List
//
// List the history of queries through SQL warehouses.
//
// You can filter by user ID, warehouse ID, status, and time range.
//
// This method is generated by Databricks SDK Code Generator.
func (a *QueryHistoryAPI) ListQueriesAll(ctx context.Context, request ListQueriesRequest) ([]QueryInfo, error) {
	var results []QueryInfo
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.ListQueries(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Res) == 0 {
			break
		}
		for _, v := range response.Res {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

func NewWarehouses(client *client.DatabricksClient) *WarehousesAPI {
	return &WarehousesAPI{
		impl: &warehousesImpl{
			client: client,
		},
	}
}

// A SQL warehouse is a compute resource that lets you run SQL commands on data
// objects within Databricks SQL. Compute resources are infrastructure resources
// that provide processing capabilities in the cloud.
type WarehousesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(WarehousesService)
	impl WarehousesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *WarehousesAPI) WithImpl(impl WarehousesService) *WarehousesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Warehouses API implementation
func (a *WarehousesAPI) Impl() WarehousesService {
	return a.impl
}

// Create a warehouse
//
// Creates a new SQL warehouse.
func (a *WarehousesAPI) CreateWarehouse(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	return a.impl.CreateWarehouse(ctx, request)
}

// Calls [WarehousesAPI.CreateWarehouse] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetWarehouseResponse](60*time.Minute) functional option.
func (a *WarehousesAPI) CreateWarehouseAndWait(ctx context.Context, createWarehouseRequest CreateWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	createWarehouseResponse, err := a.CreateWarehouse(ctx, createWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: createWarehouseResponse.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    *getWarehouseResponse,
				Timeout: i.Timeout,
			})
		}
		status := getWarehouseResponse.State
		statusMessage := getWarehouseResponse.Health.Summary
		switch status {
		case GetWarehouseResponseStateRunning: // target state
			return getWarehouseResponse, nil
		case GetWarehouseResponseStateStopped, GetWarehouseResponseStateDeleted:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetWarehouseResponseStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Delete a warehouse
//
// Deletes a SQL warehouse.
func (a *WarehousesAPI) DeleteWarehouse(ctx context.Context, request DeleteWarehouseRequest) error {
	return a.impl.DeleteWarehouse(ctx, request)
}

// Calls [WarehousesAPI.DeleteWarehouse] and waits to reach DELETED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetWarehouseResponse](60*time.Minute) functional option.
func (a *WarehousesAPI) DeleteWarehouseAndWait(ctx context.Context, deleteWarehouseRequest DeleteWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.DeleteWarehouse(ctx, deleteWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: deleteWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    *getWarehouseResponse,
				Timeout: i.Timeout,
			})
		}
		status := getWarehouseResponse.State
		statusMessage := getWarehouseResponse.Health.Summary
		switch status {
		case GetWarehouseResponseStateDeleted: // target state
			return getWarehouseResponse, nil
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Delete a warehouse
//
// Deletes a SQL warehouse.
func (a *WarehousesAPI) DeleteWarehouseById(ctx context.Context, id string) error {
	return a.impl.DeleteWarehouse(ctx, DeleteWarehouseRequest{
		Id: id,
	})
}

func (a *WarehousesAPI) DeleteWarehouseByIdAndWait(ctx context.Context, id string, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	return a.DeleteWarehouseAndWait(ctx, DeleteWarehouseRequest{
		Id: id,
	}, options...)
}

// Update a warehouse
//
// Updates the configuration for a SQL warehouse.
func (a *WarehousesAPI) EditWarehouse(ctx context.Context, request EditWarehouseRequest) error {
	return a.impl.EditWarehouse(ctx, request)
}

// Calls [WarehousesAPI.EditWarehouse] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetWarehouseResponse](60*time.Minute) functional option.
func (a *WarehousesAPI) EditWarehouseAndWait(ctx context.Context, editWarehouseRequest EditWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.EditWarehouse(ctx, editWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: editWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    *getWarehouseResponse,
				Timeout: i.Timeout,
			})
		}
		status := getWarehouseResponse.State
		statusMessage := getWarehouseResponse.Health.Summary
		switch status {
		case GetWarehouseResponseStateRunning: // target state
			return getWarehouseResponse, nil
		case GetWarehouseResponseStateStopped, GetWarehouseResponseStateDeleted:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetWarehouseResponseStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Get warehouse info
//
// Gets the information for a single SQL warehouse.
func (a *WarehousesAPI) GetWarehouse(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error) {
	return a.impl.GetWarehouse(ctx, request)
}

// Calls [WarehousesAPI.GetWarehouse] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetWarehouseResponse](60*time.Minute) functional option.
func (a *WarehousesAPI) GetWarehouseAndWait(ctx context.Context, getWarehouseRequest GetWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	getWarehouseResponse, err := a.GetWarehouse(ctx, getWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: getWarehouseResponse.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    *getWarehouseResponse,
				Timeout: i.Timeout,
			})
		}
		status := getWarehouseResponse.State
		statusMessage := getWarehouseResponse.Health.Summary
		switch status {
		case GetWarehouseResponseStateRunning: // target state
			return getWarehouseResponse, nil
		case GetWarehouseResponseStateStopped, GetWarehouseResponseStateDeleted:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetWarehouseResponseStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Get warehouse info
//
// Gets the information for a single SQL warehouse.
func (a *WarehousesAPI) GetWarehouseById(ctx context.Context, id string) (*GetWarehouseResponse, error) {
	return a.impl.GetWarehouse(ctx, GetWarehouseRequest{
		Id: id,
	})
}

func (a *WarehousesAPI) GetWarehouseByIdAndWait(ctx context.Context, id string, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	return a.GetWarehouseAndWait(ctx, GetWarehouseRequest{
		Id: id,
	}, options...)
}

// Get a configuration
//
// Gets the workspace level configuration that is shared by all SQL warehouses
// in a workspace.
func (a *WarehousesAPI) GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error) {
	return a.impl.GetWorkspaceWarehouseConfig(ctx)
}

// List warehouses
//
// Lists all SQL warehouses that a user has manager permissions on.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WarehousesAPI) ListWarehousesAll(ctx context.Context, request ListWarehousesRequest) ([]EndpointInfo, error) {
	response, err := a.impl.ListWarehouses(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Warehouses, nil
}

// Set a configuration
//
// Sets the workspace level configuration that is shared by all SQL warehouses
// in a workspace.
func (a *WarehousesAPI) SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error {
	return a.impl.SetWorkspaceWarehouseConfig(ctx, request)
}

// Start a warehouse
//
// Starts a SQL warehouse.
func (a *WarehousesAPI) StartWarehouse(ctx context.Context, request StartWarehouseRequest) error {
	return a.impl.StartWarehouse(ctx, request)
}

// Calls [WarehousesAPI.StartWarehouse] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetWarehouseResponse](60*time.Minute) functional option.
func (a *WarehousesAPI) StartWarehouseAndWait(ctx context.Context, startWarehouseRequest StartWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.StartWarehouse(ctx, startWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: startWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    *getWarehouseResponse,
				Timeout: i.Timeout,
			})
		}
		status := getWarehouseResponse.State
		statusMessage := getWarehouseResponse.Health.Summary
		switch status {
		case GetWarehouseResponseStateRunning: // target state
			return getWarehouseResponse, nil
		case GetWarehouseResponseStateStopped, GetWarehouseResponseStateDeleted:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetWarehouseResponseStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Stop a warehouse
//
// Stops a SQL warehouse.
func (a *WarehousesAPI) StopWarehouse(ctx context.Context, request StopWarehouseRequest) error {
	return a.impl.StopWarehouse(ctx, request)
}

// Calls [WarehousesAPI.StopWarehouse] and waits to reach STOPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetWarehouseResponse](60*time.Minute) functional option.
func (a *WarehousesAPI) StopWarehouseAndWait(ctx context.Context, stopWarehouseRequest StopWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.StopWarehouse(ctx, stopWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: stopWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    *getWarehouseResponse,
				Timeout: i.Timeout,
			})
		}
		status := getWarehouseResponse.State
		statusMessage := getWarehouseResponse.Health.Summary
		switch status {
		case GetWarehouseResponseStateStopped: // target state
			return getWarehouseResponse, nil
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}
