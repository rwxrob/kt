package helm_test

import (
	"fmt"

	"github.com/rwxrob/kt/pkg/helm"
)

func ExampleAppVersionFor() {
	version, err := helm.AppVersionFor(`harbor/harbor`, `1.15.1`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(version)
	// Output:
	// 2.11.1
}

func ExampleAppVersionFor_badversion() {
	version, err := helm.AppVersionFor(`harbor/harbor`, `xxx`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(version)
	// Output:
	// helm search command failed: exit status 1
}

func ExampleChartVersionsFor() {
	version, err := helm.ChartVersionsFor(`bitnami/harbor`, `2.5.1`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(version)
	// Output:
	// 14.0.2
	// 14.0.1
	// 13.2.7
	// 13.2.6
	// 13.2.4
	// 13.2.2
}
