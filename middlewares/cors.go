package middlewares

import (
	"encoding/json"
	"fmt"
	"main/constants"
	"net/http"
	"runtime/debug"
	"time"
)

func Cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")
		w.Header().Set("Content-Type", "Application/Json")

		if r.Method == "OPTIONS" {
			_ = json.NewEncoder(w).Encode(make(map[string]interface{}))
			return
		}

		defer func(w http.ResponseWriter) {
			if e := recover(); e != nil {
				fmt.Println(time.Now().Format("2006/01/02 15:04:05"), e)
				fmt.Println(string(debug.Stack()))
				fmt.Println()

				_ = json.NewEncoder(w).Encode(map[string]interface{}{
					"status":  false,
					"message": "Terjadi kesalahan dalam memproses permintaan anda",
					"data":    nil,
					"code":    constants.ErrUnknown,
				})
				return
			}
		}(w)

		next.ServeHTTP(w, r)
	}
}
