package recipe

import (
	"encoding/json"
	"net/http"

	"github.com/HuangXiaoL/HomeCookingAPI/pkg/recipe"
)

type GetOneDishInfoRes struct {
	// Name 菜品名称
	Name string `json:"name"`
	// ImageAddress 菜品图片
	ImageAddress string `json:"image_address"`
	// 菜品链接
	Link string `json:"link"`
}

// GetOneDishInfo 返回一个菜谱信息
func GetOneDishInfo(w http.ResponseWriter, r *http.Request) {
	recipeInfo := recipe.ReturnDish()
	res := &GetOneDishInfoRes{}
	res.Name = recipeInfo.Name
	res.ImageAddress = recipeInfo.ImageAddress
	res.Link = recipeInfo.Link
	data, _ := json.Marshal(res)
	w.Write(data)
}
