package v1

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/service"
	"bit-labs.cn/owl"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var _ owl.Handler = (*DictHandle)(nil)
var _ owl.CrudHandler = (*DictHandle)(nil)

type DictHandle struct {
	dictSvc *service.DictService
}

func (i *DictHandle) ModuleName() (string, string) {
	return "dict", "字典管理"
}

func NewDictHandle(dictService *service.DictService) *DictHandle {
	return &DictHandle{
		dictSvc: dictService,
	}
}

func (i *DictHandle) Create(ctx *gin.Context) {
	var req service.CreateDictReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		owl.Fail(ctx, "参数绑定失败")
		return
	}

	err := i.dictSvc.CreateDict(&req)
	owl.Auto(ctx, nil, err)
}

func (i *DictHandle) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	err := i.dictSvc.DeleteDict(id)
	owl.Auto(ctx, nil, err)
}

func (i *DictHandle) Detail(ctx *gin.Context) {

}

func (i *DictHandle) Update(ctx *gin.Context) {
	var req service.UpdateDictReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		owl.Fail(ctx, "参数绑定失败")
		return
	}
	id := cast.ToUint(ctx.Param("id"))
	req.ID = id

	err := i.dictSvc.UpdateDict(&req)
	owl.Auto(ctx, nil, err)
}

func (i *DictHandle) Retrieve(ctx *gin.Context) {
	var req service.RetrieveDictReq
	if err := ctx.ShouldBind(&req); err != nil {
		owl.Fail(ctx, "参数绑定失败")
		return
	}

	_, list, err := i.dictSvc.RetrieveDicts(&req)
	owl.Auto(ctx, list, err)
}

func (i *DictHandle) CreateItem(ctx *gin.Context) {
	var req model.DictItem
	if err := ctx.ShouldBindJSON(&req); err != nil {
		//owl.Fail(ctx, "参数绑定失败")
		//return
	}
	req.DictID = cast.ToUint(ctx.Param("id"))
	err := i.dictSvc.CreateItem(&req)
	owl.Auto(ctx, nil, err)
}
func (i *DictHandle) UpdateItem(ctx *gin.Context) {
	var req model.DictItem
	if err := ctx.ShouldBindJSON(&req); err != nil {
		owl.Fail(ctx, "参数绑定失败")
		return
	}
	err := i.dictSvc.UpdateItem(&req)
	owl.Auto(ctx, nil, err)
}
func (i *DictHandle) RetrieveItems(ctx *gin.Context) {
	dictID := ctx.Param("id")
	_, list, err := i.dictSvc.RetrieveItems(dictID)
	owl.Auto(ctx, list, err)
}

func (i *DictHandle) DeleteItem(ctx *gin.Context) {
	dictID := ctx.Param("id")
	itemID := ctx.Param("itemID")
	err := i.dictSvc.DeleteItems(dictID, itemID)
	owl.Auto(ctx, nil, err)
}
