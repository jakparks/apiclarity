module github.com/apiclarity/apiclarity/plugins/gateway/kong

go 1.16

require (
	github.com/Kong/go-pdk v0.6.0
	github.com/apiclarity/apiclarity/plugins/api v0.0.0
	github.com/apiclarity/apiclarity/plugins/common v0.0.0
	github.com/go-openapi/strfmt v0.21.0
)

replace github.com/apiclarity/apiclarity/plugins/api v0.0.0 => ./../../api

replace github.com/apiclarity/apiclarity/plugins/common v0.0.0 => ./../../common
