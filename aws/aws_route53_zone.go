// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func ListRoute53Zone(client *Client) ([]Resource, error) {
	req := client.Route53conn.ListHostedZonesRequest(&route53.ListHostedZonesInput{})

	var result []Resource

	p := route53.NewListHostedZonesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.HostedZones {

			result = append(result, Resource{
				Type:      "aws_route53_zone",
				ID:        *r.Id,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
