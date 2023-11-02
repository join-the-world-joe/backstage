package rbac

import (
	"backstage/utils/json"
	"fmt"
	"testing"
)

func TestSetPath(t *testing.T) {
	js, err := json.NewJson([]byte("{}"))
	if err != nil {
		panic(err)
	}
	js.SetPath([]string{"role_list", "name_list"}, []string{"role1", "role2", "role3"})
	js.SetPath([]string{"role_list", "description_list"}, []string{"desc1", "desc2", "desc3"})
	bs, err := js.Encode()
	if err == nil {
		fmt.Println(string(bs))
	}
}

func TestJsonEncode(t *testing.T) {
	js := json.New()
	bytes, err := js.Encode()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}
