package admin

import (
	"backstage/diagnostic"
	"backstage/plugin/crypto/rsa"
	"backstage/utils/bcrypt"
	"context"
	"encoding/json"
	"testing"
)

func TestPasswordSignIn(t *testing.T) {
	PublicKey := `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----
`
	passwordPlainText := "1234561"
	bPassword, err := bcrypt.PasswordHash(passwordPlainText)
	if err != nil {
		t.Fatal(err)
	}
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	crypto := rsa.NewCrypto(
		rsa.WithPublicKey([]byte(PublicKey)),
	)
	encryptedPasswordBytes, err := crypto.Encrypt([]byte(bPassword))
	if err != nil {
		t.Fatal(err)
	}

	req := &SignInReq{
		CountryCode: "86",
		PhoneNumber: "18629300170",
		Password:    encryptedPasswordBytes,
	}
	rsp := &SignInRsp{}
	err = SignIn(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("rsp: ", rsp)
	t.Log("Password Plain Text: ", passwordPlainText)
	t.Log("Bcrypt Password: ", []byte(bPassword))
	t.Log("Encrypted Password: ", encryptedPasswordBytes)
}

func TestFetchIdListOfGood(t *testing.T) {
	userId := int64(1)
	behavior := 1
	productName := []byte("product1")
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchIdListOfGoodReq{
		UserId:      userId,
		Behavior:    behavior,
		ProductName: productName,
	}
	rsp := &FetchIdListOfGoodRsp{}
	err := FetchIdListOfGood(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp: ", string(rsp.Body))
}

func TestFetchRecordsOfGood(t *testing.T) {
	userId := int64(1)
	productIdList := []int64{8, 1}
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchRecordsOfGoodReq{
		UserId:        userId,
		ProductIdList: productIdList,
	}
	rsp := &FetchRecordsOfGoodRsp{}
	err := FetchRecordsOfGood(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp: ", string(rsp.Body))
}

func TestInsertRecordOfGood(t *testing.T) {
	name := []byte("某水")
	vendor := []byte("某商家")
	contact := []byte("某商家联系方式")
	description := []byte("某水描述")
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &InsertRecordOfGoodReq{
		UserId:      1,
		Name:        name,
		BuyingPrice: 1000,
		Vendor:      vendor,
		Contact:     contact,
		Description: description,
	}
	rsp := &InsertRecordOfGoodRsp{}
	err := InsertRecordOfGood(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSoftDeleteRecordOfGood(t *testing.T) {
	userId := int64(1)
	productIdList := []int64{9}
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &SoftDeleteRecordsOfGoodReq{
		UserId:        userId,
		ProductIdList: productIdList,
	}
	rsp := &SoftDeleteRecordsOfGoodRsp{}
	err := SoftDeleteRecordOfGood(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Code: ", rsp.Code)
}

func TestUpdateRecordOfGood(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	userId := int64(1)
	productId := int64(9)
	name := "product_name"
	buyingPrice := 100
	status := 0
	vendor := "product_vendor"
	contact := "product_contact"
	description := "product_description"
	req := &UpdateRecordOfGoodReq{
		Name:        []byte(name),
		UserId:      userId,
		Status:      status,
		Vendor:      []byte(vendor),
		Contact:     []byte(contact),
		BuyingPrice: buyingPrice,
		ProductId:   productId,
		Description: []byte(description),
	}
	rsp := &UpdateRecordOfGoodRsp{}
	err := UpdateRecordOfGood(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFetchIdListOfAdvertisement(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	userId := int64(1)
	name := []byte("title1")
	behavior := 1
	req := &FetchIdListOfAdvertisementReq{
		Behavior:          behavior,
		UserId:            userId,
		AdvertisementName: name,
	}
	rsp := &FetchIdListOfAdvertisementRsp{}
	err := FetchIdListOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("bytes: ", string(rsp.Body))
}

func TestFetchRecordsOfAdvertisement(t *testing.T) {
	userId := int64(1)
	idList := []int64{3}
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchRecordsOfAdvertisementReq{
		UserId:              userId,
		AdvertisementIdList: idList,
	}
	rsp := &FetchRecordsOfAdvertisementRsp{}
	err := FetchRecordsOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp: ", string(rsp.Body))
}

func TestInsertRecordOfAdvertisement(t *testing.T) {
	userId := int64(1)
	name := []byte("某水广告")
	productId := int64(1)
	points := []string{"森林1", "大树1", "point3333"}
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()

	sellingPoints, err := json.Marshal(&points)
	if err != nil {
		t.Fatal(err)
	}
	req := &InsertRecordOfAdvertisementReq{
		UserId:       userId,
		Name:         name,
		ProductId:    productId,
		SellingPoint: string(sellingPoints),
	}
	rsp := &InsertRecordOfAdvertisementRsp{}
	err = InsertRecordOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp.code: ", rsp.Code)
	_ = req
	t.Log("SellingPoint: ", sellingPoints)
}

func TestSoftDeleteRecordOfAdvertisement(t *testing.T) {
	userId := int64(1)
	advertisementIdList := []int64{21}
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &SoftDeleteRecordsOfAdvertisementReq{
		UserId:              userId,
		AdvertisementIdList: advertisementIdList,
	}
	rsp := &SoftDeleteRecordsOfAdvertisementRsp{}
	err := SoftDeleteRecordsOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Code: ", rsp.Code)
}

func TestUpdateRecordOfAdvertisement(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	userId := int64(1)
	productId := int64(21)
	name := "product_name"
	status := 0
	req := &UpdateRecordOfAdvertisementReq{
		Name:   []byte(name),
		UserId: userId,
		Status: status,

		ProductId: productId,
	}
	rsp := &UpdateRecordOfAdvertisementRsp{}
	err := UpdateRecordOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp.Code: ", rsp.Code)
}
