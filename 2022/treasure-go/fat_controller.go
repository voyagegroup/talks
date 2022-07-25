package contoller

import "net/http"

type FatController struct{}

func (c *FatController) UpdateHoge(w http.ResponseWriter, r *http.Request) {
	// リクエストの処理 (クエリパラメータ、リクエストボディの処理など)
	// ビジネスロジック
	// DBとのやり取り
	// レスポンスの処理
}
