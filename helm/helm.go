package helm

import (
	"encoding/json"
	"fmt"

	"github.com/rwxrob/kutil/cli"
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

	jsonData := cli.Out(
		`helm`, `search`, `repo`, chart, `-o`, `json`, `--version`, chartver)

	var charts []Chart
	if err := json.Unmarshal([]byte(jsonData), &charts); err != nil {
		return "", err
	}

	for _, chart := range charts {
		if chart.ChartVersion == chartver {
			return chart.AppVersion, nil
		}
	}

	return "", fmt.Errorf(`app version not found for chart version %q`, chartver)
}
