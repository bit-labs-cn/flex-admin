package owl

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm/schema"
)

type Handle[T schema.Tabler] struct {
	form T
	repo *Repository[T]
}

func NewHandle[T schema.Tabler](repo *Repository[T]) Handle[T] {
	return Handle[T]{
		repo: repo,
	}
}
func (i *Handle[T]) Create(ctx *gin.Context) {
	form := new(T)
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	err := i.repo.Create(*form)
	ctx.JSON(200, gin.H{"msg": err})
}

func (i *Handle[T]) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := i.repo.Delete(cast.ToInt64(id))
	ctx.JSON(200, gin.H{"msg": err})
}

func (i *Handle[T]) List(ctx *gin.Context) {

}
func (i *Handle[T]) Retrieve(ctx *gin.Context) {

}
func (i *Handle[T]) Update(ctx *gin.Context) {
	form := new(T)
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	err := i.repo.Update(cast.ToInt64(ctx.Param("id")), *form)
	ctx.JSON(200, gin.H{"msg": err})
}
func (i *Handle[T]) Detail(ctx *gin.Context) {

}
