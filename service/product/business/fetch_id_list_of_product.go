package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/product"
	product2 "backstage/common/protocol/product"
	"backstage/global/log"
	"context"
	"encoding/json"
	"fmt"
)

type OutputOfIdListOfGood struct {
	Behavior        int     `json:"behavior"`
	IdListOfProduct []int64 `json:"id_list_of_product"`
}

func FetchIdListOfProduct(ctx context.Context, req *product2.FetchIdListOfProductReq, rsp *product2.FetchIdListOfProductRsp) error {
	if len(req.ProductName) <= 0 && req.CategoryId <= 0 && (req.Behavior != 1 && req.Behavior != 2) {
		rsp.Code = code.InvalidData
		return nil
	}

	output := &OutputOfIdListOfGood{
		Behavior:        req.Behavior,
		IdListOfProduct: []int64{},
	}

	if req.Behavior == 1 {
		idList, err := product.GetIdListInTable()
		if err != nil {
			rsp.Code = code.DatabaseFailure
			return nil
		}
		for _, v := range idList {
			output.IdListOfProduct = append(output.IdListOfProduct, v)
		}
	} else if req.Behavior == 2 {
		if len(req.ProductName) > 0 {
			ml, err := product.GetModelListByName(string(req.ProductName))
			if err != nil {
				log.Error(fmt.Sprintf("FetchIdListOfGood.product.GetModelListByName failure, err: %v", err.Error()))
			} else {
				for _, v := range ml {
					output.IdListOfProduct = append(output.IdListOfProduct, v.Id)
				}
			}
		}
	} else {
		rsp.Code = code.InvalidData
		return nil
	}

	bytes, err := json.Marshal(output)
	if err != nil {
		rsp.Code = code.InternalError
		return nil
	}

	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}
