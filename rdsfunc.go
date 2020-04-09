package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"

	//"github.com/fatih/color"
	color "github.com/logrusorgru/aurora"
)

func getrds(filter string) {
	reg := regexp.MustCompile(filter)

	params := &rds.DescribeDBInstancesInput{}
	req := rdssvc.DescribeDBInstancesRequest(params)

	//resp, err := req.Send(context.TODO())
	p := rds.NewDescribeDBInstancesPaginator(req)

	for p.Next(context.TODO()) {
		page := p.CurrentPage()

		for _, v := range page.DBInstances {
			if reg.MatchString(aws.StringValue(v.DBInstanceIdentifier)) {
				//tags := make(map[string]string)
				fmt.Println(color.Bold(aws.StringValue(v.DBInstanceIdentifier)))

				t := listTagsForResource(v.DBInstanceArn)
				if len(t.ListTagsForResourceOutput.TagList) != 0 {
					//fmt.Printf("[")
					for _, v2 := range t.ListTagsForResourceOutput.TagList {
						//tags[aws.StringValue(v2.Key)] = aws.StringValue(v2.Value)
						fmt.Printf(`"` + aws.StringValue(v2.Key) + `"`)
						fmt.Printf(`:`)
						fmt.Printf(`"` + aws.StringValue(v2.Value) + `"`)
						fmt.Println()
					}
					//fmt.Printf("]")
				}
				//fmt.Println(tags)
			}
		}
	}

	if err := p.Err(); err != nil {
		fmt.Println(err)
	}

}

func listTagsForResource(rdsarn *string) *rds.ListTagsForResourceResponse {
	params := &rds.ListTagsForResourceInput{
		ResourceName: rdsarn,
	}
	req := rdssvc.ListTagsForResourceRequest(params)
	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	return resp
}
