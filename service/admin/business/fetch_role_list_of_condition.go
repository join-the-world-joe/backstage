package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/macro/abbreviation"
	role2 "backstage/common/macro/role"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/spf13/cast"
	"golang.org/x/exp/slices"
)

type Role struct {
	Name        string `json:"name"`
	Rank        int    `json:"rank"`
	Department  string `json:"department"`
	Description string `json:"description"`
}

type RoleListOutput struct {
	RoleList []*Role `json:"role_list"`
	Length   int     `json:"length"`
}

var chinese = map[string]string{
	"硬件工程师":  "Hardware_Engineer",
	"采购专员":   "Purchasing_Specialist",
	"生产线专员":  "Production_Specialist",
	"研发部主管":  "RD_Manager",
	"财务部主管":  "Finance_Manger",
	"人事部主管":  "HR_Manger",
	"人事行政专员": "HR_Specialist",
	"生产部主管":  "Manufacturing_Manger",
	"系统管理员":  "Administrator",
	"软件工程师":  "Software_Engineer",
	"财务专员":   "Accounting_Specialist",
	"市场部主管":  "Marketing_Manger",
	"销售专员":   "Sales_Specialist",
}

var english = map[string]string{
	"Hardware Engineer":     "Hardware_Engineer",
	"Purchasing Specialist": "Purchasing_Specialist",
	"Production Specialist": "Production_Specialist",
	"RD Manager":            "RD_Manager",
	"Finance Manger":        "Finance_Manger",
	"HR Manger":             "HR_Manger",
	"HR Specialist":         "HR_Specialist",
	"Manufacturing Manger":  "Manufacturing_Manger",
	"Administrator":         "Administrator",
	"Software Engineer":     "Software_Engineer",
	"Accounting Specialist": "Accounting_Specialist",
	"Marketing Manger":      "Marketing_Manger",
	"Sales Specialist":      "Sales_Specialist",
}

func translate(inputs []string) []string {
	outputs := []string{}
	matches := []string{}
	ch := func() []string {
		output := []string{}
		for k, _ := range chinese {
			output = append(output, k)
		}
		return output
	}()
	en := func() []string {
		output := []string{}
		for k, _ := range english {
			output = append(output, k)
		}
		return output
	}()

	for _, input := range inputs {
		if len(input) <= 0 {
			continue
		}
		mts := fuzzy.Find(input, ch)
		for _, v := range mts {
			outputs = append(outputs, v)
		}
	}
	for _, input := range inputs {
		if len(input) <= 0 {
			continue
		}
		mts := fuzzy.Find(input, en)
		for _, v := range mts {
			outputs = append(outputs, v)
		}
	}

	for _, k := range outputs {
		if v, exist := chinese[k]; exist {
			matches = append(matches, v)
		}
		if v, exist := english[k]; exist {
			matches = append(matches, v)
		}
	}

	return matches
}

func FetchRoleListOfCondition(ctx context.Context, req *admin.FetchRoleListOfConditionReq, rsp *admin.FetchRoleListOfConditionRsp) error {
	if req.Id <= 0 {
		log.Error("FetchRoleListOfCondition failure, req.UserId <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	selfRoleList := user_role.GetRoleListByUserId(req.Id)
	if len(selfRoleList) <= 0 {
		log.Error("FetchRoleListOfCondition failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range selfRoleList {
		if rbac.HasPermission(v, cast.ToInt(admin.FetchRoleListOfConditionReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	roleList := []string{}
	if req.Behavior == 1 {
		//roleList = selfRoleList
		if slices.Contains(selfRoleList, role2.Administrator) {
			roleList = rbac.GetRoleListLERank(rbac.GetTopRankOfRoleList(selfRoleList))
		} else {
			for _, role := range selfRoleList {
				if name, department, _, _, b := rbac.GetRole(role); b {
					roleList = rbac.GetRoleListLERankInDepartment(rbac.GetTopRankOfRoleList([]string{name}), department)
				}
			}
		}
	} else if req.Behavior == 2 {
		if len(req.RoleList) <= 0 && req.UserId <= 0 {
			rsp.Code = code.NoData
			return nil
		}
		fmt.Println("req.RoleList1: ", req.RoleList)
		if len(req.RoleList) >= 0 {
			fmt.Println("req.RoleList2: ", req.RoleList)
			for _, inputName := range req.RoleList {
				translatedRoleList := translate([]string{string(inputName)})
				for _, roleName := range translatedRoleList {
					if !slices.Contains(roleList, roleName) {
						roleList = append(roleList, roleName)
					}
				}
			}
		}
		if req.UserId > 0 {
			fmt.Println("roleList1: ", roleList)
			tempRoleList := user_role.GetRoleListByUserId(req.UserId)
			fmt.Println("tempRoleList:", tempRoleList)
			for _, roleName := range tempRoleList {
				if !slices.Contains(roleList, roleName) {
					roleList = append(roleList, roleName)
				}
			}
		}
	} else {
		rsp.Code = code.InvalidData
		return nil
	}

	fmt.Println("roleList2: ", roleList)

	nameList := []string{}
	//descList := []string{}
	//rankList := []string{}
	//departmentList := []string{}
	rlo := &RoleListOutput{RoleList: []*Role{}, Length: 0}
	for _, v := range roleList {
		name, departmnt, desc, rnk, b := rbac.GetRole(v)
		//fmt.Println("n: ", v)
		//fmt.Println("r: ", r)
		//fmt.Println("d: ", d)
		//fmt.Println("b: ", b)
		if b && !slices.Contains(nameList, name) {
			nameList = append(nameList, name)
			description := abbreviation.NA
			if len(desc) > 0 {
				description = desc
			}
			department := abbreviation.NA
			if len(departmnt) > 0 {
				department = departmnt
			}
			role := &Role{
				Name:        name,
				Rank:        rnk,
				Description: description,
				Department:  department,
			}
			rlo.RoleList = append(rlo.RoleList, role)
			rlo.Length += 1
		}
	}

	bytes, err := json2.Marshal(rlo)
	if err != nil {
		rsp.Code = code.InternalError
		return nil
	}

	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}
