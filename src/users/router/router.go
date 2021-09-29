package router

import (
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/internal/web"
)

type RouterStruct struct {
	web.RouterStruct
	jwtAuth auth.JwtTokenInterface
}
