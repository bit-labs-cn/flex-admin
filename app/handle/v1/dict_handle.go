package v1

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/service"
	"bit-labs.cn/owl/provider/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var _ router.Handler = (*DictHandle)(nil)
var _ router.CrudHandler = (*DictHandle)(nil)

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

// Create 创建字典
//
//	@Summary		创建字典
//	@Description	创建新的字典项
//	@Tags			字典管理
//	@Router			/api/v1/dict [POST]

// @Name			创建字典
// @Param			createDictReq	body		service.CreateDictReq	true	"字典创建请求参数"
// @Success		200				{object}	router.RouterInfo		"字典创建成功"
// @Failure		400				{object}	router.RouterInfo		"请求参数错误"
// @Failure		500				{object}	router.RouterInfo		"服务器内部错误"
func (i *DictHandle) Create(ctx *gin.Context) {
	var req service.CreateDictReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		router.Fail(ctx, "参数绑定失败")
		return
	}

	err := i.dictSvc.CreateDict(&req)
	router.Auto(ctx, nil, err)
}

// Delete 删除字典
//
//	@Summary		删除字典
//	@Description	根据字典ID删除指定字典
//	@Tags			字典管理
//	@Router			/api/v1/dict/{id} [DELETE]

// @Name			删除字典
// @Param			id	path		string				true	"字典ID"
// @Success		200	{object}	router.RouterInfo	"字典删除成功"
// @Failure		400	{object}	router.RouterInfo	"请求参数错误"
// @Failure		500	{object}	router.RouterInfo	"服务器内部错误"
func (i *DictHandle) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	err := i.dictSvc.DeleteDict(id)
	router.Auto(ctx, nil, err)
}

func (i *DictHandle) Detail(ctx *gin.Context) {

}

// Update 更新字典
//
//	@Summary		更新字典
//	@Description	根据字典ID更新字典信息
//	@Tags			字典管理
//	@Router			/api/v1/dict/{id} [PUT]

// @Name			更新字典
// @Param			id				path		int						true	"字典ID"
// @Param			updateDictReq	body		service.UpdateDictReq	true	"字典更新请求参数"
// @Success		200				{object}	router.RouterInfo		"字典更新成功"
// @Failure		400				{object}	router.RouterInfo		"请求参数错误"
// @Failure		500				{object}	router.RouterInfo		"服务器内部错误"
func (i *DictHandle) Update(ctx *gin.Context) {
	var req service.UpdateDictReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		router.Fail(ctx, "参数绑定失败")
		return
	}
	id := cast.ToUint(ctx.Param("id"))
	req.ID = id

	err := i.dictSvc.UpdateDict(&req)
	router.Auto(ctx, nil, err)
}

// Retrieve 获取字典列表
//
//	@Summary		获取字典列表
//	@Description	分页获取字典列表，支持搜索和筛选
//	@Tags			字典管理
//	@Name			获取字典列表
//	@Param			page		query		int					false	"页码"	default(1)
//	@Param			pageSize	query		int					false	"每页数量"	default(10)
//	@Param			keyword		query		string				false	"搜索关键词"
//	@Success		200			{object}	router.RouterInfo	"字典列表获取成功"
//	@Failure		400			{object}	router.RouterInfo	"请求参数错误"
//	@Failure		500			{object}	router.RouterInfo	"服务器内部错误"
//	@Router			/api/v1/dict [GET]

func (i *DictHandle) Retrieve(ctx *gin.Context) {
	var req service.RetrieveDictReq
	if err := ctx.ShouldBind(&req); err != nil {
		router.Fail(ctx, "参数绑定失败")
		return
	}

	_, list, err := i.dictSvc.RetrieveDicts(&req)
	router.Auto(ctx, list, err)
}

// CreateItem 创建字典项
//
//	@Summary		创建字典项
//	@Description	为指定字典创建新的字典项
//	@Tags			字典管理
//	@Router			/api/v1/dict/{id}/items [POST]

// @Name			创建字典项
// @Param			id			path		int					true	"字典ID"
// @Param			dictItem	body		model.DictItem		true	"字典项创建请求参数"
// @Success		200			{object}	router.RouterInfo	"字典项创建成功"
// @Failure		400			{object}	router.RouterInfo	"请求参数错误"
// @Failure		500			{object}	router.RouterInfo	"服务器内部错误"
func (i *DictHandle) CreateItem(ctx *gin.Context) {
	var req model.DictItem
	if err := ctx.ShouldBindJSON(&req); err != nil {
		//router.Fail(ctx, "参数绑定失败")
		//return
	}
	req.DictID = cast.ToUint(ctx.Param("id"))
	err := i.dictSvc.CreateItem(&req)
	router.Auto(ctx, nil, err)
}

// UpdateItem 更新字典项
//
//	@Summary		更新字典项
//	@Description	更新指定字典项的信息
//	@Tags			字典管理
//	@Router			/api/v1/dict/{id}/items/{itemID} [PUT]

// @Name			更新字典项
// @Param			id			path		int					true	"字典ID"
// @Param			itemID		path		int					true	"字典项ID"
// @Param			dictItem	body		model.DictItem		true	"字典项更新请求参数"
// @Success		200			{object}	router.RouterInfo	"字典项更新成功"
// @Failure		400			{object}	router.RouterInfo	"请求参数错误"
// @Failure		500			{object}	router.RouterInfo	"服务器内部错误"
func (i *DictHandle) UpdateItem(ctx *gin.Context) {
	var req model.DictItem
	if err := ctx.ShouldBindJSON(&req); err != nil {
		router.Fail(ctx, "参数绑定失败")
		return
	}
	err := i.dictSvc.UpdateItem(&req)
	router.Auto(ctx, nil, err)
}

// RetrieveItems 获取字典项列表
//
//	@Summary		获取字典项列表
//	@Description	获取指定字典的所有字典项
//	@Tags			字典管理
//	@Router			/api/v1/dict/{id}/items [GET]

// @Name			获取字典项列表
// @Param			id	path		string				true	"字典ID"
// @Success		200	{object}	router.RouterInfo	"字典项列表获取成功"
// @Failure		400	{object}	router.RouterInfo	"请求参数错误"
// @Failure		500	{object}	router.RouterInfo	"服务器内部错误"
func (i *DictHandle) RetrieveItems(ctx *gin.Context) {
	dictID := ctx.Param("id")
	_, list, err := i.dictSvc.RetrieveItems(dictID)
	router.Auto(ctx, list, err)
}

// DeleteItem 删除字典项
//
//	@Summary		删除字典项
//	@Description	删除指定字典的指定字典项
//	@Tags			字典管理
//	@Router			/api/v1/dict/{id}/items/{itemID} [DELETE]

// @Name			删除字典项
// @Param			id		path		string				true	"字典ID"
// @Param			itemID	path		string				true	"字典项ID"
// @Success		200		{object}	router.RouterInfo	"字典项删除成功"
// @Failure		400		{object}	router.RouterInfo	"请求参数错误"
// @Failure		500		{object}	router.RouterInfo	"服务器内部错误"
func (i *DictHandle) DeleteItem(ctx *gin.Context) {
	dictID := ctx.Param("id")
	itemID := ctx.Param("itemID")
	err := i.dictSvc.DeleteItems(dictID, itemID)
	router.Auto(ctx, nil, err)
}
