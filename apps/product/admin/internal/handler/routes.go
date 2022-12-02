// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go_zero/apps/product/admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/v1/upload/image",
				Handler: UploadImageHandler(serverCtx),
			},
		},
	)
}
