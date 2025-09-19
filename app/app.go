package app

import (
	"context"

	"github.com/Timasha/core/components"
	"github.com/Timasha/core/log"
	"github.com/pkg/errors"
)

type App struct {
	Components []components.Lifecycle
	name       string
}

func NewApp(name string, components []components.Lifecycle) *App {
	return &App{name: name, Components: components}
}

func (a *App) Start(ctx context.Context) (err error) {
	log.Infof("Starting: %v", a.GetName())

	for _, component := range a.Components {
		if !component.IsEnabled() {
			log.Infof("Component %v is disabled. Skip...", component.GetName())
			continue
		}

		err = component.Start(ctx)
		if err != nil {
			return errors.Errorf("error starting %v: %v", component.GetName(), err)
		}

		log.Infof("%v Started", component.GetName())
	}

	log.Infof("Started: %v", a.GetName())

	return nil
}

func (a *App) Stop(ctx context.Context) (err error) {
	log.Infof("Stopping: %v", a.GetName())

	for i := len(a.Components) - 1; i >= 0; i-- {
		if !a.Components[i].IsEnabled() {
			log.Infof("Component %v is disabled. Skip...", a.Components[i].GetName())
			continue
		}

		err = a.Components[i].Stop(ctx)
		if err != nil {
			return errors.Errorf("error stopping %v: %v", a.Components[i].GetName(), err)
		}

		log.Infof("%v Stopped", a.Components[i].GetName())
	}

	log.Infof("Stopped: %v", a.GetName())

	return nil
}

func (a *App) GetName() string {
	return a.name
}

func (a *App) IsEnabled() bool {
	return true
}

func Run[T any](cfg Config, a components.Lifecycle, sig chan T) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), cfg.StartTimeout.Duration)

	err = a.Start(ctx)
	if err != nil {
		return err
	}

	<-sig

	ctx, _ = context.WithTimeout(context.Background(), cfg.StopTimeout.Duration)

	err = a.Stop(ctx)
	if err != nil {
		return err
	}

	return nil
}
