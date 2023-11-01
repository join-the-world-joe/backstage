package business

import (
	role2 "backstage/common/macro/role"
	"backstage/common/service/backend"
	"backstage/diagnostic"
	"context"
	"encoding/json"
	"testing"
)

func TestFetchRoleList(t *testing.T) {
	role := "Manager"
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	req := &backend.FetchRoleListReq{Role: role}
	rsp := &backend.FetchRoleListRsp{}
	err := FetchRoleList(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Role List: ", rsp.RoleList)
}

func TestFetchMenuList(t *testing.T) {
	role := "Manager"
	//role := role2.Administrator
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	req := &backend.FetchMenuListReq{Role: role}
	rsp := &backend.FetchMenuListRsp{}
	err := FetchMenuList(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Code: ", rsp.Code)
	t.Log("Body: ", string(rsp.Body))
}

func TestFetchAttributeList(t *testing.T) {
	role := "Manager"
	//role := role2.Administrator
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	req := &backend.FetchAttributeListReq{Role: role}
	rsp := &backend.FetchAttributeListRsp{}
	err := FetchAttributeList(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Code: ", rsp.Code)
	t.Log("Body: ", string(rsp.Body))
}

func TestFetchPermissionList(t *testing.T) {
	role := "Manager"
	//role := role2.Administrator
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	req := &backend.FetchPermissionListReq{Role: role}
	rsp := &backend.FetchPermissionListRsp{}
	err := FetchPermissionList(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Code: ", rsp.Code)
	t.Log("Permission List: ", rsp.PermissionList)
}

func TestFetchMenuListOfRole(t *testing.T) {
	//role := "Manager"
	role := role2.Administrator
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	req := &backend.FetchMenuListOfRoleReq{Role: role, Query: role}
	rsp := &backend.FetchMenuListOfRoleRsp{}
	err := FetchMenuListOfRole(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := json.Marshal(rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
	t.Log(rsp.Body)
	//t.Log(convert.Bytes2StringArray(rsp.Body))
	//t.Log(string(rsp.Body))
}

func TestFetchPermissionListOfRole(t *testing.T) {
	role := "Manager"
	//role := role2.Administrator
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	req := &backend.FetchPermissionListOfRoleReq{Role: role, Query: role}
	rsp := &backend.FetchPermissionListOfRoleRsp{}
	err := FetchPermissionListOfRole(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Code: ", rsp.Code)
	t.Log("Permission List: ", rsp.PermissionList)
}

func TestFetchAttributeListOfRole(t *testing.T) {
	role := "Manager"
	//role := role2.Administrator
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	req := &backend.FetchAttributeListOfRoleReq{Role: role, Query: role}
	rsp := &backend.FetchAttributeListOfRoleRsp{}
	err := FetchAttributeListOfRole(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Code: ", rsp.Code)
	t.Log("Body: ", string(rsp.Body))
}
