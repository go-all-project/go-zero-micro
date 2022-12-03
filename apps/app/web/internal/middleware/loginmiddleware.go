package middleware

import "net/http"

type LoginMiddleware struct {
}

func NewLoginMiddleware() *LoginMiddleware {
	return &LoginMiddleware{}
}

func (m *LoginMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 中间件
		next(w, r)
	}
}
