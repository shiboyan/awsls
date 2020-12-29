// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func ListRedshiftSecurityGroup(client *Client) ([]Resource, error) {
	req := client.Redshiftconn.DescribeClusterSecurityGroupsRequest(&redshift.DescribeClusterSecurityGroupsInput{})

	var result []Resource

	p := redshift.NewDescribeClusterSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ClusterSecurityGroups {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:      "aws_redshift_security_group",
				ID:        *r.ClusterSecurityGroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
