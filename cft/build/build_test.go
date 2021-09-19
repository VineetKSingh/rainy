package build_test

import (
	"testing"

	"github.com/vineetksingh/rainy/cft/build"
	"github.com/vineetksingh/rainy/cft/spec"
)

var allResourceTypes map[string]string

func init() {
	allResourceTypes = make(map[string]string)

	for resourceType := range spec.Cfn.ResourceTypes {
		allResourceTypes[resourceType] = resourceType
	}
}

func TestAllResourceTypes(t *testing.T) {
	for resourceType := range spec.Cfn.ResourceTypes {
		build.Template(map[string]string{
			"Res": resourceType,
		}, true)
	}
}

func BenchmarkAllResourceTypesIndividually(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for resourceType := range allResourceTypes {
			build.Template(map[string]string{
				"Res": resourceType,
			}, true)
		}
	}
}

func BenchmarkAllResourceTypesInOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		build.Template(allResourceTypes, true)
	}
}
