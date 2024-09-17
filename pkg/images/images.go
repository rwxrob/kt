package images

import "strings"

// ConvertTag converts container image tag to one with for the specified
// namespace and registry.
func ConvertTag(tag, registry, namespace string) string {
	f := strings.Split(tag, `/`)
	switch len(f) {
	case 3:
		tag = f[2]
	case 2:
		tag = f[1]
	}
	return registry + `/` + namespace + `/` + tag
}
