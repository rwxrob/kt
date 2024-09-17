package helm

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// AppVersionFor uses the helm command to lookup the app_version for
// the given chart version.
func AppVersionFor(chart, chartver string) (string, error) {
	var err error

	type achart struct {
		Name       string `json:"name"`
		AppVersion string `json:"app_version"`
	}

	cmd := exec.Command(
		`helm`, `search`, `repo`, chart, `-o`, `json`, `--version`, chartver)
	jsonData, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("helm search command failed: %w", err)
	}

	var charts []achart
	err = json.Unmarshal([]byte(jsonData), &charts)
	if err != nil {
		return "", err
	}

	switch len(charts) {
	case 0:
		return "", fmt.Errorf(`no charts found`)
	case 1:
		return charts[0].AppVersion, nil
	default:
		return charts[0].AppVersion,
			fmt.Errorf(`more than one chart has same chart version`)
	}

}
