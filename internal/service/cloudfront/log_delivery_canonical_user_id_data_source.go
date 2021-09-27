package cloudfront

import (
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

const (
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html#AccessLogsBucketAndFileOwnership.
	defaultCloudFrontLogDeliveryCanonicalUserId = "c4c1ede66af53448b93c283ce9448c4ba468c9432aa01d700d3878632f77d2d0"

	// See https://docs.amazonaws.cn/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html#AccessLogsBucketAndFileOwnership.
	cnCloudFrontLogDeliveryCanonicalUserId = "a52cb28745c0c06e84ec548334e44bfa7fc2a85c54af20cd59e4969344b7af56"
)

func DataSourceLogDeliveryCanonicalUserID() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLogDeliveryCanonicalUserIDRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceLogDeliveryCanonicalUserIDRead(d *schema.ResourceData, meta interface{}) error {
	canonicalId := defaultCloudFrontLogDeliveryCanonicalUserId

	region := meta.(*conns.AWSClient).Region
	if v, ok := d.GetOk("region"); ok {
		region = v.(string)
	}

	if v, ok := endpoints.PartitionForRegion(endpoints.DefaultPartitions(), region); ok && v.ID() == endpoints.AwsCnPartitionID {
		canonicalId = cnCloudFrontLogDeliveryCanonicalUserId
	}

	d.SetId(canonicalId)

	return nil
}