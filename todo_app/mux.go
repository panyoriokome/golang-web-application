// ルーティング定義を担当する
// どのようなハンドラーの実装をどんなURLパスで公開するかルーティングする
package main

import "net/http"

// 戻り値を*http.ServeMux型の値ではなくhttp.Handlerインタフェース
// にすることで内部実装に依存しない関数シグネチャにする
func NewMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	return mux
}
