package runtime

import (
	appv1alpha1 "github.com/verzth/cosmos-sdk/api/cosmos/app/v1alpha1"
	autocliv1 "github.com/verzth/cosmos-sdk/api/cosmos/autocli/v1"
	reflectionv1 "github.com/verzth/cosmos-sdk/api/cosmos/reflection/v1"

	"github.com/verzth/cosmos-sdk/runtime/services"
	"github.com/verzth/cosmos-sdk/types/module"
)

func (a *App) registerRuntimeServices(cfg module.Configurator) error {
	appv1alpha1.RegisterQueryServer(cfg.QueryServer(), services.NewAppQueryService(a.appConfig))
	autocliv1.RegisterQueryServer(cfg.QueryServer(), services.NewAutoCLIQueryService(a.ModuleManager.Modules))

	reflectionSvc, err := services.NewReflectionService()
	if err != nil {
		return err
	}
	reflectionv1.RegisterReflectionServiceServer(cfg.QueryServer(), reflectionSvc)

	return nil
}
