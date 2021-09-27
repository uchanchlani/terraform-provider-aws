package aws

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/provider"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

func TestAccDataSourceAwsCognitoUserPools_basic(t *testing.T) {
	rName := fmt.Sprintf("tf_acc_ds_cognito_user_pools_%s", sdkacctest.RandString(7))
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(t); testAccPreCheckAWSCognitoIdentityProvider(t) },
		ErrorCheck: acctest.ErrorCheck(t, cognitoidentityprovider.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsCognitoUserPoolsConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.aws_cognito_user_pools.selected", "ids.#", "2"),
					resource.TestCheckResourceAttr("data.aws_cognito_user_pools.selected", "arns.#", "2"),
				),
			},
			{
				Config:      testAccDataSourceAwsCognitoUserPoolsConfig_notFound(rName),
				ExpectError: regexp.MustCompile(`No cognito user pool found with name:`),
			},
		},
	})
}

func testAccDataSourceAwsCognitoUserPoolsConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_cognito_user_pool" "main" {
  count = 2
  name  = "%s"
}

data "aws_cognito_user_pools" "selected" {
  name = aws_cognito_user_pool.main.*.name[0]
}
`, rName)
}

func testAccDataSourceAwsCognitoUserPoolsConfig_notFound(rName string) string {
	return fmt.Sprintf(`
data "aws_cognito_user_pools" "selected" {
  name = "%s-not-found"
}
`, rName)
}