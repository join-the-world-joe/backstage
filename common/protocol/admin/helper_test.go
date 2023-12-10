package admin

import (
	"backstage/diagnostic"
	"backstage/plugin/crypto/rsa"
	"backstage/utils/bcrypt"
	"context"
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
	productIdList := []int64{1, 2, 3}
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
	userId := int64(1)
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	name := []byte("蒙牛酸酸乳")
	vendor := []byte("汕头市蒙牛奶业有限公司")
	contact := []byte("0756-88788371")
	description := []byte("正规渠道、国产大牌子、大人小孩都爱")
	req := &InsertRecordOfGoodReq{
		UserId:      userId,
		Name:        name,
		BuyingPrice: 100,
		Status:      0, // 为0时, 采用数据库设定的默认值
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

func TestSoftDeleteRecordsOfGood(t *testing.T) {
	userId := int64(1)
	productIdList := []int64{1, 2, 3, 4}
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
	productId := int64(4)
	name := []byte("product_name")
	buyingPrice := 100
	status := 1
	vendor := []byte("product_vendor")
	contact := []byte("product_contact")
	description := []byte("product_description")
	req := &UpdateRecordOfGoodReq{
		Name:        name,
		UserId:      userId,
		Status:      status,
		Vendor:      vendor,
		Contact:     contact,
		BuyingPrice: buyingPrice,
		ProductId:   productId,
		Description: description,
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
	idList := []int64{4}
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
	name := []byte("name")
	title := []byte("title")
	sellingPrice := 100
	sellingPoints := [][]byte{[]byte("2"), []byte("大2")}
	placeOfOrigin := []byte("地要工")
	url := []byte("urlfd在")
	stock := 10
	productId := int64(1)
	status := 1
	description := []byte("描述")
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()

	req := &InsertRecordOfAdvertisementReq{
		UserId:        userId,
		Url:           url,
		Name:          name,
		Title:         title,
		ProductId:     productId,
		SellingPoints: sellingPoints,
		PlaceOfOrigin: placeOfOrigin,
		SellingPrice:  sellingPrice,
		Stock:         stock,
		Status:        status,
		Description:   description,
	}
	rsp := &InsertRecordOfAdvertisementRsp{}
	err := InsertRecordOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp.code: ", rsp.Code)
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
	id := int64(3)
	name := []byte("name new")
	title := []byte("title new")
	sellingPrice := 100
	sellingPoints := [][]byte{[]byte("11"), []byte("大小")}
	placeOfOrigin := []byte("地要工22")
	url := []byte("urlfd在111")
	stock := 10
	productId := int64(1)
	status := 1
	description := []byte("描述 new")
	req := &UpdateRecordOfAdvertisementReq{
		Id:            id,
		Url:           url,
		Stock:         stock,
		Name:          name,
		Title:         title,
		UserId:        userId,
		Status:        status,
		SellingPrice:  sellingPrice,
		SellingPoints: sellingPoints,
		PlaceOfOrigin: placeOfOrigin,
		Description:   description,
		ProductId:     productId,
	}
	rsp := &UpdateRecordOfAdvertisementRsp{}
	err := UpdateRecordOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp.Code: ", rsp.Code)
}

func TestInsertRecordOfADOfCarousel(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	userId := int64(1)
	advertisementIdList := []int64{1, 2, 3}
	req := &InsertRecordOfADOfCarouselReq{
		UserId:              userId,
		AdvertisementIdList: advertisementIdList,
	}
	rsp := &InsertRecordOfADOfCarouselRsp{}
	err := InsertRecordOfADOfCarousel(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp.Code: ", rsp.Code)
}

func TestRemoveOutdatedRecordsOfADOfCarousel(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &RemoveOutdatedRecordsOfADOfCarouselReq{
		UserId: 1,
	}
	rsp := &RemoveOutdatedRecordsOfADOfCarouselRsp{}
	err := RemoveOutdatedRecordsOfADOfCarousel(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Code: ", rsp.Code)
}
