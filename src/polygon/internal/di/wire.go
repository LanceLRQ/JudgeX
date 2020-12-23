// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/wejudge/wejudge-polygon/src/polygon/internal/dao"
	"github.com/wejudge/wejudge-polygon/src/polygon/internal/service"
	"github.com/wejudge/wejudge-polygon/src/polygon/internal/server/grpc"
	"github.com/wejudge/wejudge-polygon/src/polygon/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
