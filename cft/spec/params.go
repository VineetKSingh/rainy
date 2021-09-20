package spec

import (
	"strings"
)

var Parameters = map[string]map[string]interface{}{
	"String": {
		"Type":                  "String",
		"AllowedValues":         "",
		"AllowedPattern":        "",
		"ConstraintDescription": "",
		"Default":               "",
		"Description":           "",
		"MaxLength":             0,
		"MinLength":             0,
		"NoEcho":                "",
	},
	"Number": {
		"AllowedValues": "",
		"Default":       "",
		"MaxValue":      0,
		"MinValue":      0,
		"Description":   "",
	},
	"List<Number>": {
		"AllowedValues": "",
		"Default":       "",
		"Description":   "",
	},
	"CommaDelimitedList": {
		"Default":     "",
		"Description": "",
	},
	"AWS::EC2::AvailabilityZone::Name": {
		"Default":       "",
		"AllowedValues": "",
	},
	"AWS::EC2::Image::Id": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"AWS::EC2::Instance::Id": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"AWS::EC2::KeyPair::KeyName": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"AWS::EC2::SecurityGroup::GroupName": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"AWS::EC2::SecurityGroup::Id": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"AWS::EC2::Subnet::Id": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": ""},
	"AWS::EC2::Volume::Id": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": ""},
	"AWS::EC2::VPC::Id": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"AWS::Route53::HostedZone::Id": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"List<AWS::EC2::AvailabilityZone::Name>": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"List<AWS::EC2::Image::Id>": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"List<AWS::EC2::Instance::Id>": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"List<AWS::EC2::SecurityGroup::GroupName>": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"List<AWS::EC2::SecurityGroup::Id>": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"List<AWS::EC2::Subnet::Id>": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"List<AWS::EC2::Volume::Id>": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"List<AWS::EC2::VPC::Id>": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"List<AWS::Route53::HostedZone::Id>": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"AWS::SSM::Parameter::Name": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"AWS::SSM::Parameter::Value<String>": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"AWS::SSM::Parameter::Value<List<String>>": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
	"AWS::SSM::Parameter::Value<CommaDelimitedList>": {
		"Default":               "",
		"Description":           "",
		"AllowedValues":         "",
		"ConstraintDescription": "",
	},
}

func ResolveParams(paramName string) []string {

	paramName = strings.ToLower(paramName)

	options := make([]string, 0)

	for p := range Parameters {
		if strings.HasSuffix(strings.ToLower(p), paramName) {
			options = append(options, p)
		}
	}

	return options
}
