package helm_test

import (
	"fmt"
	"log"
	"os"

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

func ExampleFlatten() {
	err := helm.Flatten(`bitnami/harbor`, `2.11.1`, `testdata/flattentest`)
	if err != nil {
		fmt.Println(err)
	}
	// TODO output something to validate testdata
	// Output:
	// something
}

func ExampleFlatten_impliedPath() {
	if err := os.Chdir(`testdata/flattentest`); err != nil {
		log.Fatal(err)
	}
	if err := helm.Flatten(`bitnami/harbor`, `2.11.1`); err != nil {
		fmt.Println(err)
	}
	// TODO output something to validate testdata
	// Output:
	// something
}

func ExampleDetected() {
	fmt.Println(helm.Detected())
	// Output:
	// true
}
