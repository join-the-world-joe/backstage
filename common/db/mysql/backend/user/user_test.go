package user

import (
	"backstage/diagnostic"
	"backstage/global/mysql"
	"backstage/utils/bcrypt"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	t.Log(uuid.New().String())
}

func TestAutoMigrate(t *testing.T) {
	diagnostic.SetupMySQL()
	err := mysql.AutoMigrate(GetWhich(), GetDbName(), GetTableName(), &Model{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDropTable(t *testing.T) {
	diagnostic.SetupMySQL()
	err := mysql.DropTable(GetWhich(), GetDbName(), GetTableName())
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	id := int64(1)
	diagnostic.SetupMySQL()
	err := mysql.Delete(GetWhich(), GetDbName(), sqlDeleteAnyById(id))
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsertModel(t *testing.T) {
	name := "流星"
	countryCode := "86"
	phoneNumber := "18629309942"
	password := "123456"
	diagnostic.SetupMySQL()
	encryptedPassword, err := bcrypt.PasswordHash(password)
	if err != nil {
		t.Fatal(err)
	}
	temp, err := InsertModel(&Model{Name: name, CountryCode: countryCode, PhoneNumber: phoneNumber, Password: encryptedPassword, MemberId: uuid.New().String()})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}

func TestGiantInsertModel(t *testing.T) {
	nb := 1000
	for i := 0; i < nb; i++ {
		name := fmt.Sprintf("流星%v", i)
		countryCode := "86"
		phoneNumber := fmt.Sprintf("18622032123%v", i)
		password := "123456"
		diagnostic.SetupMySQL()
		encryptedPassword, err := bcrypt.PasswordHash(password)
		if err != nil {
			t.Fatal(err)
		}
		_, err = InsertModel(&Model{Name: name, CountryCode: countryCode, PhoneNumber: phoneNumber, Password: encryptedPassword})
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGetModelByMobile(t *testing.T) {
	countryCode := "86"
	phoneNumber := "18629309942"
	diagnostic.SetupMySQL()
	m, err := GetModelByMobile(countryCode, phoneNumber)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("m: ", m)
}

func TestGetModelById(t *testing.T) {
	userId := int64(5)
	diagnostic.SetupMySQL()
	m, err := GetModelById(userId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("m: ", m)
}

func TestGetModelByPhoneNumber(t *testing.T) {
	phoneNumber := "18629300170"
	diagnostic.SetupMySQL()
	m, err := GetModelByPhoneNumber(phoneNumber)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("m: ", m)
}

func TestCheckPassword(t *testing.T) {
	countryCode := "86"
	phoneNumber := "18629300170"
	password := "123456"
	diagnostic.SetupMySQL()
	record, err := GetModelByMobile(countryCode, phoneNumber)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(bcrypt.PasswordVerify(password, record.Password))
}

func TestGetByName(t *testing.T) {
	name := "流星"
	diagnostic.SetupMySQL()
	m, err := GetModelByName(name)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("m: ", m)
}

//func TestTransform(t *testing.T) {
//	input := []*Model{
//		&Model{
//			Name: "1",
//		},
//		&Model{
//			Name: "2",
//		},
//	}
//	output := Transform([]string{"name"}, input)
//	bytes, err := json2.Marshal(output)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	js := json.New()
//	js.Set("user_list", json2.RawMessage(bytes))
//	body, err := js.Encode()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	t.Log("body: ", string(body))
//}

func TestGetModelListByName(t *testing.T) {
	diagnostic.SetupMySQL()
	name := "流星"
	userList, err := GetModelListByName(name)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("User List: ", userList)
}

func TestGetModelListByPhoneNumber(t *testing.T) {
	diagnostic.SetupMySQL()
	phoneNumber := "18629300173"
	userList, err := GetModelListByPhoneNumber(phoneNumber)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("User List: ", userList)
}

func TestNotIn(t *testing.T) {
	userIdList := []int64{1, 2, 3, 4}
	sql := "("
	length := len(userIdList)
	for k, v := range userIdList {
		if k < length-1 {
			sql += fmt.Sprintf("%v, ", v)
			continue
		}
		sql += fmt.Sprintf("%v", v)
	}
	sql += ")"
	t.Log("sql: ", sql)
}

func TestGetUserIdListNotInUserIdList(t *testing.T) {
	diagnostic.SetupMySQL()
	userIdList := []int64{1, 2}
	t.Log(GetUserIdListNotInUserIdList(userIdList))
}
