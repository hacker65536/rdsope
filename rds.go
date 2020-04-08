package main

import "github.com/aws/aws-sdk-go-v2/service/rds"

var (
	rdssvc *rds.Client
)

func init() {
	rdssvc = rds.New(awsCfg)
}
