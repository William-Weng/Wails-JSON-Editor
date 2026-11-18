package backend

import (
	"context"
	"wails-json-editor/backend/update"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type UpdateService struct {
	lastReleaseURL string
}

func (service *UpdateService) CheckUpdate(ctx context.Context) (*update.UpdateInfo, error) {

	const currentVersion = "0.2.1"
	const repository = "William-Weng/Wails-JSON-Editor"

	info, err := update.CheckUpdate(ctx, currentVersion, repository)
	if err != nil {
		return nil, err
	}

	service.lastReleaseURL = info.ReleaseURL
	return info, nil
}

func (service *UpdateService) OpenReleasePage() error {

	app := application.Get()
	if app == nil || service.lastReleaseURL == "" {
		return nil
	}

	return app.Browser.OpenURL(service.lastReleaseURL)
}
