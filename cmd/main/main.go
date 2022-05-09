package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/HuangXiaoL/HomeCookingAPI/internal/app/recipe"
	recipe2 "github.com/HuangXiaoL/HomeCookingAPI/pkg/recipe"
)

func main() {
	fmt.Println("开始初始化数据 Please wait for....")
	recipe2.Init()
	fmt.Println("初始化数据完成")
	fmt.Printf("包含: %d 菜品", len(recipe2.StorageRecipeInfoArr))
	webSever()

}
func webSever() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", recipe.GetOneDishInfo)

	http.ListenAndServe(":7788", r)
}
