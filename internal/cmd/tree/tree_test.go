package tree_test

import (
	"os"

	"github.com/vineetksingh/rainy/internal/cmd/tree"
	"github.com/vineetksingh/rainy/internal/console"
)

func Example_tree() {
	os.Args = []string{
		os.Args[0],
		"../../../test/templates/success.template",
	}

	console.NoColour = true

	tree.Cmd.Execute()
	// Output:
	// Resources:
	//   Bucket1:
	//     DependsOn:
	//       Parameters:
	//         - BucketName
}
