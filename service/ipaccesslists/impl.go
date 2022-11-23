// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ipaccesslists

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just IpAccessLists API methods
type ipAccessListsImpl struct {
	client *client.DatabricksClient
}

func (a *ipAccessListsImpl) CreateIpAccessList(ctx context.Context, request CreateIpAccessListRequest) (*IpAccessListInfo, error) {
	var ipAccessListInfo IpAccessListInfo
	path := "/api/2.0/ip-access-lists"
	err := a.client.Post(ctx, path, request, &ipAccessListInfo)
	return &ipAccessListInfo, err
}

func (a *ipAccessListsImpl) DeleteIpAccessList(ctx context.Context, request DeleteIpAccessListRequest) error {
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *ipAccessListsImpl) FetchIpAccessList(ctx context.Context, request FetchIpAccessListRequest) (*IpAccessListInfo, error) {
	var ipAccessListInfo IpAccessListInfo
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Get(ctx, path, request, &ipAccessListInfo)
	return &ipAccessListInfo, err
}

func (a *ipAccessListsImpl) ListIpAccessLists(ctx context.Context) (*GetIpAccessListResponse, error) {
	var getIpAccessListResponse GetIpAccessListResponse
	path := "/api/2.0/ip-access-lists"
	err := a.client.Get(ctx, path, nil, &getIpAccessListResponse)
	return &getIpAccessListResponse, err
}

func (a *ipAccessListsImpl) ReplaceIpAccessList(ctx context.Context, request ReplaceIpAccessListRequest) error {
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Put(ctx, path, request)
	return err
}

func (a *ipAccessListsImpl) UpdateIpAccessList(ctx context.Context, request UpdateIpAccessListRequest) error {
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Patch(ctx, path, request)
	return err
}
