package console

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/vineetksingh/rainy/internal/aws/console"
	tty "github.com/vineetksingh/rainy/internal/console"
	"github.com/vineetksingh/rainy/internal/console/spinner"
)

// Open generates a sign-in URL to the AWS console with an optional service and resource
// If printOnly is true, the URL is printed to the console
// If printOnly is fale, Open attempts to call the OS's browser with the URL
func Open(printOnly bool, service, resource, userName string) {
	spinner.Push("Generating sign-in URL")
	uri, err := console.GetURI(service, resource, userName)
	if err != nil {
		panic(err)
	}
	spinner.Pop()

	if !printOnly {
		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", uri).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", uri).Start()
		case "darwin":
			err = exec.Command("open", uri).Start()
		}
	}

	if printOnly || err != nil {
		if tty.IsTTY {
			fmt.Print("Open the following URL in your browser: ")
		}
		fmt.Println(uri)
	}
}
