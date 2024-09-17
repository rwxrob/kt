package helm

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// Only stuff we care about
type Chart struct {
	Name         string `json:"name"`
	ChartVersion string `json:"version"`
	AppVersion   string `json:"app_version"`
}

// AppVersionFor uses the helm command to lookup the app_version for
// the given chart version.
func AppVersionFor(chart, chartver string) (string, error) {
	var err error

	cmd := exec.Command(
		`helm`, `search`, `repo`, chart, `-o`, `json`, `--version`, chartver)
	jsonData, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("helm search command failed: %w", err)
	}

	var charts []Chart
	err = json.Unmarshal([]byte(jsonData), &charts)
	if err != nil {
		return "", err
	}

	for _, chart := range charts {
		if chart.ChartVersion == chartver {
			return chart.AppVersion, nil
		}
	}

	return "", fmt.Errorf(`app version not found for chart version %q`, chartver)
}

// ChartVersionFor uses the helm command to lookup the version for the
// given app_version.
func ChartVersionFor(chart, appver string) (string, error) {

	cmd := exec.Command(
		`helm`, `search`, `repo`, chart, `-o`, `json`, `--versions`)
	jsonData, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("helm search command failed: %w", err)
	}

	var charts []Chart
	if err := json.Unmarshal([]byte(jsonData), &charts); err != nil {
		return "", err
	}

	for _, chart := range charts {
		if chart.AppVersion == appver {
			return chart.ChartVersion, nil
		}
	}

	return "", fmt.Errorf(`chart version not found for app version %q`, appver)
}
