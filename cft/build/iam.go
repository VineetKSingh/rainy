package build

import (
	"github.com/vineetksingh/rainy/cft"
	"github.com/vineetksingh/rainy/cft/spec"
)

// iamBuilder contains specific code for building IAM policies
type iamBuilder struct {
	builder
}

// newIamBuilder creates a new iamBuilder
func newIamBuilder() iamBuilder {
	var b iamBuilder
	b.Spec = spec.Iam

	return b
}

// Policy generates a an IAM policy body
func (b iamBuilder) Policy() (interface{}, []*cft.Comment) {
	return b.newPropertyType("", "Policy")
}
