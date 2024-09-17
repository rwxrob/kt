package helm_test

import (
	"fmt"

	"github.com/rwxrob/kutil/helm"
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
