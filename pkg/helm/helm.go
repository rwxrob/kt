package helm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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

// ChartVersionFor uses the helm command to lookup all chart versions
// for which the app_version is the passed argument.
func ChartVersionsFor(chart, appver string) (string, error) {

	helmout := new(bytes.Buffer)
	cmd := exec.Command(`helm`, `search`, `repo`, chart, `--versions`, `-o`, `json`)
	cmd.Stdout = helmout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	type achart struct {
		Name         string `json:"name"`
		ChartVersion string `json:"version"`
		AppVersion   string `json:"app_version"`
	}

	var charts []achart
	if err := json.Unmarshal(helmout.Bytes(), &charts); err != nil {
		return "", err
	}

	var out string
	for _, chart := range charts {
		if chart.AppVersion == appver {
			out += chart.ChartVersion + "\n"
		}
	}

	return out, nil
}

func Flatten(args ...string) error {
	log.Print(`unimplemented call to Flatten`)
	return nil
}

// Detected returns true if helm command is found in executable path.
func Detected() bool {
	out, _ := exec.LookPath(`helm`)
	return len(out) > 0
}
