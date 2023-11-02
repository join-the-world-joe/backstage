package business

import (
	"encoding/json"
	"fmt"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"testing"
	"time"
	"unicode/utf8"
)

func TestSlicesContains(t *testing.T) {
	userId := int64(10)
	t.Log(testSlicesContains(userId))
}

func TestMarshalEmptySlice(t *testing.T) {
	kk := []string{}
	bytes, err := json.Marshal(&kk)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("bytes: ", string(bytes))
}

func TestTimeFormat(t *testing.T) {
	str := time.Now().Format("2006-01-02 15:04:05")
	t.Log(str)
}

func TestUTF8EncodeAndDecode(t *testing.T) {
	r := '流'
	p := make([]byte, 3)
	n := utf8.EncodeRune(p, r)
	fmt.Printf("bytes: %v, nums: %d\n", p, n)

	newRune, n := utf8.DecodeRune(p)
	fmt.Printf("rune: %c, nums: %d\n", newRune, n)
}

func TestRune(t *testing.T) {
	r := '流'
	p := make([]byte, 3)
	n := utf8.EncodeRune(p, r) // 230, 181, 129
	fmt.Printf("bytes: %v, nums: %d\n", p, n)

	r = '星'
	p = make([]byte, 3)
	n = utf8.EncodeRune(p, r) // 230, 152, 159
	fmt.Printf("bytes: %v, nums: %d\n", p, n)

	bytes := []byte{230, 181, 129, 230, 152, 159} // golang 默认是utf8编码
	t.Log(string(bytes))                          // 流星
}

func TestRoleListOutput(t *testing.T) {
	role1 := Role{
		Rank:        1,
		Department:  "R&D",
		Description: "role1",
	}
	role2 := Role{
		Rank:        2,
		Department:  "R&D",
		Description: "role2",
	}
	role3 := Role{
		Rank:        1,
		Department:  "R&D",
		Description: "role3",
	}
	roleList := []*Role{
		&role1,
		&role2,
		&role3,
	}
	ro := &RoleListOutput{
		Length:   len(roleList),
		RoleList: roleList,
	}
	bytes, err := json.Marshal(ro)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Role List Output: ", string(bytes))
}

func TestTrackListOutput(t *testing.T) {
	track1 := Track{
		Operator: "track1",
	}
	track2 := Track{
		Operator: "track2",
	}
	track3 := Track{
		Operator: "track3",
	}
	trackList := []*Track{
		&track1,
		&track2,
		&track3,
	}
	tlo := &TrackListOutput{
		Length:    len(trackList),
		TrackList: trackList,
	}
	bytes, err := json.Marshal(tlo)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Track List Output: ", string(bytes))
}

func TestUserListOutput(t *testing.T) {
	user1 := UserOutput{
		Id: "1",
	}
	user2 := UserOutput{
		Id: "2",
	}
	user3 := UserOutput{
		Id: "3",
	}
	userList := []*UserOutput{
		&user1,
		&user2,
		&user3,
	}
	tlo := &UserListOutput{
		Length:   len(userList),
		UserList: userList,
	}
	bytes, err := json.Marshal(tlo)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("User List Output: ", string(bytes))
}

func TestPermissionListOutput(t *testing.T) {
	permission1 := Permission{
		Name: "permission1",
	}
	permission2 := Permission{
		Name: "permission2",
	}
	permission3 := Permission{
		Name: "permission2",
	}
	permissionList := []*Permission{
		&permission1,
		&permission2,
		&permission3,
	}
	tlo := &PermissionListOutput{
		Length:         len(permissionList),
		PermissionList: permissionList,
	}
	bytes, err := json.Marshal(tlo)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Permission List Output: ", string(bytes))
}

func TestFuzzy(t *testing.T) {
	input := "工程师"
	source := []string{"软件工程师", "硬件工程师", "财务部主管", "会计专员"}

	matches := fuzzy.Find(input, source)
	t.Log("matches: ", matches)
}

//
//func TestFetchRoleList(t *testing.T) {
//	role := "Manager"
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	req := &backend.FetchRoleListReq{Role: role}
//	rsp := &backend.FetchRoleListRsp{}
//	err := FetchRoleList(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("Role List: ", rsp.RoleList)
//}
//
//func TestFetchMenuList(t *testing.T) {
//	role := "Manager"
//	//role := role2.Administrator
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	req := &backend.FetchMenuListReq{Role: role}
//	rsp := &backend.FetchMenuListRsp{}
//	err := FetchMenuList(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("Code: ", rsp.Code)
//	t.Log("Body: ", string(rsp.Body))
//}
//
//func TestFetchAttributeList(t *testing.T) {
//	role := "Manager"
//	//role := role2.Administrator
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	req := &backend.FetchAttributeListReq{Role: role}
//	rsp := &backend.FetchAttributeListRsp{}
//	err := FetchAttributeList(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("Code: ", rsp.Code)
//	t.Log("Body: ", string(rsp.Body))
//}
//
//func TestFetchPermissionList(t *testing.T) {
//	role := "Manager"
//	//role := role2.Administrator
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	req := &backend.FetchPermissionListReq{Role: role}
//	rsp := &backend.FetchPermissionListRsp{}
//	err := FetchPermissionList(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("Code: ", rsp.Code)
//	t.Log("Permission List: ", rsp.PermissionList)
//}
//
//func TestFetchMenuListOfRole(t *testing.T) {
//	//role := "Manager"
//	role := role2.Administrator
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	req := &backend.FetchMenuListOfRoleReq{Role: role, Query: role}
//	rsp := &backend.FetchMenuListOfRoleRsp{}
//	err := FetchMenuListOfRole(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	bytes, err := json.Marshal(rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log(string(bytes))
//	t.Log(rsp.Body)
//	//t.Log(convert.Bytes2StringArray(rsp.Body))
//	//t.Log(string(rsp.Body))
//}
//
//func TestFetchPermissionListOfRole(t *testing.T) {
//	role := "Manager"
//	//role := role2.Administrator
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	req := &backend.FetchPermissionListOfRoleReq{Role: role, Query: role}
//	rsp := &backend.FetchPermissionListOfRoleRsp{}
//	err := FetchPermissionListOfRole(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("Code: ", rsp.Code)
//	t.Log("Permission List: ", rsp.PermissionList)
//}
//
//func TestFetchAttributeListOfRole(t *testing.T) {
//	role := "Manager"
//	//role := role2.Administrator
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	req := &backend.FetchAttributeListOfRoleReq{Role: role, Query: role}
//	rsp := &backend.FetchAttributeListOfRoleRsp{}
//	err := FetchAttributeListOfRole(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("Code: ", rsp.Code)
//	t.Log("Body: ", string(rsp.Body))
//}
