package format_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vineetksingh/rainy/cft/format"
	"github.com/vineetksingh/rainy/cft/parse"
)

const input = `
Outputs:
  Bucket1:
    Value: !GetAtt Bucket1.Arn # Short GetAtt
  Bucket2: # Bucket comment
    Value:
      Fn::GetAtt: # GetAtt comment
        - Bucket2
        - Arn # Arn comment

Resources:
  Bucket2:
    Properties:
      BucketName: !Ref Name # Ref: comment
    Type: "AWS::S3::Bucket"
  Bucket1:
    Type: AWS::S3::Bucket
    Properties:
      BucketName: !Sub ${Bucket2}-newer
Parameters:
  Name:
    Type: String
`

const expectedYaml = `Parameters:
  Name:
    Type: String

Resources:
  Bucket2:
    Type: AWS::S3::Bucket
    Properties:
      BucketName: !Ref Name # Ref: comment

  Bucket1:
    Type: AWS::S3::Bucket
    Properties:
      BucketName: !Sub ${Bucket2}-newer

Outputs:
  Bucket1:
    Value: !GetAtt Bucket1.Arn # Short GetAtt

  Bucket2: # Bucket comment
    Value: !GetAtt Bucket2.Arn # GetAtt comment Arn comment
`

const expectedYamlUnsorted = `Outputs:
  Bucket1:
    Value: !GetAtt Bucket1.Arn # Short GetAtt

  Bucket2: # Bucket comment
    Value: !GetAtt Bucket2.Arn # GetAtt comment Arn comment

Resources:
  Bucket2:
    Properties:
      BucketName: !Ref Name # Ref: comment
    Type: AWS::S3::Bucket

  Bucket1:
    Type: AWS::S3::Bucket
    Properties:
      BucketName: !Sub ${Bucket2}-newer

Parameters:
  Name:
    Type: String
`

const expectedJson = `{
    "Parameters": {
        "Name": {
            "Type": "String"
        }
    },
    "Resources": {
        "Bucket2": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "BucketName": {
                    "Ref": "Name"
                }
            }
        },
        "Bucket1": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "BucketName": {
                    "Fn::Sub": "${Bucket2}-newer"
                }
            }
        }
    },
    "Outputs": {
        "Bucket1": {
            "Value": {
                "Fn::GetAtt": "Bucket1.Arn"
            }
        },
        "Bucket2": {
            "Value": {
                "Fn::GetAtt": "Bucket2.Arn"
            }
        }
    }
}
`

const expectedUnsortedJson = `{
    "Outputs": {
        "Bucket1": {
            "Value": {
                "Fn::GetAtt": "Bucket1.Arn"
            }
        },
        "Bucket2": {
            "Value": {
                "Fn::GetAtt": "Bucket2.Arn"
            }
        }
    },
    "Resources": {
        "Bucket2": {
            "Properties": {
                "BucketName": {
                    "Ref": "Name"
                }
            },
            "Type": "AWS::S3::Bucket"
        },
        "Bucket1": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "BucketName": {
                    "Fn::Sub": "${Bucket2}-newer"
                }
            }
        }
    },
    "Parameters": {
        "Name": {
            "Type": "String"
        }
    }
}
`

func checkMatch(t *testing.T, expected string, opt format.Options) {
	template, err := parse.String(input)
	if err != nil {
		t.Fatal(err)
	}

	actual := format.String(template, opt)

	if d := cmp.Diff(expected, actual); d != "" {
		t.Errorf(d)
	}
}

func TestFormatDefault(t *testing.T) {
	checkMatch(t, expectedYaml, format.Options{})
	checkMatch(t, expectedYamlUnsorted, format.Options{
		Unsorted: true,
	})
	checkMatch(t, expectedJson, format.Options{
		JSON: true,
	})
	checkMatch(t, expectedUnsortedJson, format.Options{
		JSON:     true,
		Unsorted: true,
	})
}
