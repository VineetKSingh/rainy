package build_test

import (
	"os"

	"github.com/vineetksingh/rainy/internal/cmd/build"
)

func Example_build_bucket() {
	os.Args = []string{
		os.Args[0],
		"-b",
		"AWS::S3::Bucket",
	}

	build.Cmd.Execute()
	// Output:
	// AWSTemplateFormatVersion: "2010-09-09"
	//
	// Description: Template generated by rain
	//
	// Resources:
	//   MyBucket:
	//     Type: AWS::S3::Bucket
	//     Properties: {}
}
