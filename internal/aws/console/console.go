package console

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/aws/smithy-go/ptr"
	"github.com/vineetksingh/rainy/internal/aws"
	"github.com/vineetksingh/rainy/internal/aws/cfn"
	"github.com/vineetksingh/rainy/internal/aws/sts"
)

const signinURI = "https://signin.aws.amazon.com/federation"
const issuer = "https://aws-cloudformation.github.io/rain/rain_console.html"
const consoleURI = "https://console.aws.amazon.com"
const defaultService = "cloudformation"
const sessionDuration = 43200

func buildSessionString(sessionName string) (string, error) {
	if sessionName == "" {
		id, err := sts.GetCallerID()
		if err != nil {
			return "", err
		}

		idParts := strings.Split(ptr.ToString(id.Arn), ":")
		nameParts := strings.Split(idParts[len(idParts)-1], "/")

		if nameParts[0] == "user" {
			panic(errors.New("sign-in URLs can only be constructed for assumed roles"))
		}

		sessionName = nameParts[1]
	}

	creds, err := aws.NamedConfig(sessionName).Credentials.Retrieve(context.Background())
	if err != nil {
		return "", err
	}

	return url.QueryEscape(fmt.Sprintf(`{"sessionId": "%s", "sessionKey": "%s", "sessionToken": "%s"}`,
		creds.AccessKeyID,
		creds.SecretAccessKey,
		creds.SessionToken,
	)), nil
}

func getSigninToken(userName string) (string, error) {
	sessionString, err := buildSessionString(userName)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(fmt.Sprintf("%s?Action=getSigninToken&Session=%s&SessionDuration=%d", signinURI, sessionString, sessionDuration))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var out map[string]string
	err = json.Unmarshal(body, &out)
	if err != nil {
		return "", err
	}

	token, ok := out["SigninToken"]
	if !ok {
		return "", errors.New("No token present in the response")
	}

	return token, nil
}

// GetURI returns a sign-in uri for the current credentials and region
func GetURI(service, stackName, userName string) (string, error) {
	token, err := getSigninToken(userName)
	if err != nil {
		return "", err
	}

	if service == "" {
		service = defaultService
	}

	destination := fmt.Sprintf("%s/%s/home?region=%s", consoleURI, service, aws.Config().Region)

	if service == defaultService && stackName != "" {
		if stack, err := cfn.GetStack(stackName); err == nil {
			if stack.StackId != nil {
				destination += fmt.Sprintf("#/stacks/stackinfo?stackId=%s&hideStacks=false&viewNested=true",
					ptr.ToString(stack.StackId),
				)
			}
		}
	}

	return fmt.Sprintf("%s?Action=login&Issuer=%s&Destination=%s&SigninToken=%s",
		signinURI,
		url.QueryEscape(issuer),
		url.QueryEscape(destination),
		url.QueryEscape(token),
	), nil
}
