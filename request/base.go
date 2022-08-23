package request

import (
	"hr/config"
	"hr/repository"
	"strconv"

	"github.com/kataras/iris/v12"
)

func GetSearchParams(ctx iris.Context, search func(map[string]string) []repository.Condition) (page, perPage int, condition []repository.Condition) {
	query := ctx.URLParams()
	page = 1
	if _, exist := query["page"]; exist {
		page, _ = strconv.Atoi(query["page"])
		delete(query, "page")
	}
	perPage = config.PERPAGE
	if _, exist := query["per_page"]; exist {
		perPage, _ = strconv.Atoi(query["per_page"])
		delete(query, "per_page")
	}
	condition = search(query)
	return
}
