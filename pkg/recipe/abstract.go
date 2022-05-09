package recipe

var (
	url                  = "https://www.xiangha.com/caipu/c-jiachang/hot-"
	StorageRecipeInfoArr []*StorageRecipeInfo
)

// StorageRecipeInfo 存储菜谱所需结构体
type StorageRecipeInfo struct {
	// Name 菜品名称
	Name string `json:"name"`
	// ImageAddress 菜品图片
	ImageAddress string `json:"image_address"`
	// 菜品链接
	Link string `json:"link"`
}
