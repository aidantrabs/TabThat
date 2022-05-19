package controllers

import (
	"example/bookmark-api/services"
)

type BookmarkController struct {
	BookmarkService services.BookmarkService
}

func New(bmService services.BookmarkService) BookmarkController {
	return BookmarkController {
		BookmarkService: bmService,
	}
}

func (bmc *BookmarkController) CreateBM(ctx *gin.Context) {
	
	var bm models.Bookmark
	if err := ctx.ShouldBindJSON(&bm); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bmc.BookmarkService.CreateBM(&bm); err != nil {
		ctx.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Bookmark created successfully!"})
}

func (bmc *BookmarkController) GetBM(ctx *gin.Context) {
	
	var id string:= ctx.Param("id")

	bm, err := bmc.BookmarkService.GetBM(&id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, bm)
}

func (bmc *BookmarkController) GetAllBM(ctx *gin.Context) {
	
	bms, err := bmc.BookmarkService.GetAllBM()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, bms)
}

func (bmc *BookmarkController) UpdateBM(ctx *gin.Context) {
	
	var bm models.Bookmark
	if err := ctx.ShouldBindJSON(&bm); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bmc.BookmarkService.UpdateBM(&bm); err != nil {
		ctx.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Bookmark updated successfully!"})
}

func (bmc *BookmarkController) DeleteBM(ctx *gin.Context) {
	
	var id string:= ctx.Param("id")

	if err := bmc.BookmarkService.DeleteBM(&id); err != nil {
		ctx.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Bookmark deleted successfully!"})
}

func (bmc *BookmarkController) RegisterRoutes(r *gin.RouterGroup) {
	
	bmroute := r.Group("/bookmark")
	bookmark.POST("/create", bmc.CreateBM)
	bookmark.GET("/get/:id", bmc.GetBM)
	bookmark.GET("/getall", bmc.GetAllBM)
	bookmark.PATCH("/update", bmc.UpdateBM)
	bookmark.DELETE("/delete/:id", bmc.DeleteBM)
}