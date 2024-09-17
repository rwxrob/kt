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
