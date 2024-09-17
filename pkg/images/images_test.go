package images_test

import (
	"fmt"

	"github.com/rwxrob/kt/pkg/images"
)

func ExampleConvertTag_tagOnly() {
	fmt.Println(
		images.ConvertTag(`my-app:v1.0.0`, `example.com`, `ns`))
	// Output:
	// example.com/ns/my-app:v1.0.0
}

func ExampleConvertTag_tagAndNamespace() {
	fmt.Println(
		images.ConvertTag(`ns/my-app:v1.0.0`, `example.com`, `myns`))
	// Output:
	// example.com/myns/my-app:v1.0.0
}

func ExampleConvertTag_tagRegistryAndNamespace() {
	fmt.Println(
		images.ConvertTag(`example.org/ns/my-app:v1.0.0`, `example.com`, `myns`))
	// Output:
	// example.com/myns/my-app:v1.0.0
}
