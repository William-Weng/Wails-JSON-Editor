package update

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
)

type UpdateInfo struct {
	HasUpdate   bool   `json:"hasUpdate"`
	CurrentVer  string `json:"currentVersion"`
	LatestVer   string `json:"latestVersion"`
	ReleaseURL  string `json:"releaseUrl"`
	DownloadURL string `json:"downloadUrl"`
}

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	HTMLURL string `json:"html_url"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func CheckUpdate(ctx context.Context, currentVersion string, repository string) (*UpdateInfo, error) {

	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repository)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "Wails-JSON-Editor")

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf(
			"GitHub API 回應 %d：%s；remaining=%s，reset=%s",
			resp.StatusCode,
			string(body),
			resp.Header.Get("X-RateLimit-Remaining"),
			resp.Header.Get("X-RateLimit-Reset"),
		)
	}

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	latestVersion := strings.TrimPrefix(release.TagName, "v")
	currentSemver, err := semver.NewVersion(
		strings.TrimPrefix(currentVersion, "v"),
	)
	if err != nil {
		return nil, fmt.Errorf("無效的目前版本 %q: %w", currentVersion, err)
	}

	latestSemver, err := semver.NewVersion(latestVersion)
	if err != nil {
		return nil, fmt.Errorf("無效的最新版本 %q: %w", latestVersion, err)
	}

	info := &UpdateInfo{
		CurrentVer: currentSemver.String(),
		LatestVer:  latestSemver.String(),
		ReleaseURL: release.HTMLURL,
	}

	if latestSemver.GreaterThan(currentSemver) {
		info.HasUpdate = true

		for _, asset := range release.Assets {
			if isCurrentPlatformAsset(asset.Name) {
				info.DownloadURL = asset.BrowserDownloadURL
				break
			}
		}
	}

	return info, nil
}

// 依你的實際 release 檔名調整
func isCurrentPlatformAsset(name string) bool {

	if runtime.GOOS == "darwin" && runtime.GOARCH == "arm64" {
		return strings.Contains(name, "macOS") && strings.Contains(name, "arm64")
	}

	if runtime.GOOS == "windows" && runtime.GOARCH == "amd64" {
		return strings.Contains(name, "Win11") && strings.Contains(name, "x86_64")
	}

	return false
}
