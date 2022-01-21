package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	examplesDir = "../../examples"
)

var providerFactories = map[string]func() (*schema.Provider, error){
	"fortios": func() (*schema.Provider, error) {
		return New("acctest")(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("acc-test")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = New("acc-test")()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("FORTIOS_ACCESS_HOSTNAME"); err == "" {
		t.Fatal("FORTIOS_ACCESS_HOSTNAME must be set for acceptance tests")
	}
	if err := os.Getenv("FORTIOS_ACCESS_TOKEN"); err == "" {
		t.Fatal("FORTIOS_ACCESS_TOKEN must be set for acceptance tests")
	}
	if err := os.Getenv("FORTIOS_VDOM"); err == "" {
		t.Fatal("FORTIOS_VDOM must be set for acceptance tests")
	}
}

func testAccCreateResourceFromExampleStep(rName string) resource.TestStep {
	// skip test if no example is provided
	if testAccExampleResourceConfig(rName) == "" {
		return resource.TestStep{
			SkipFunc: testAccSkipTestStep,
		}
	}
	return resource.TestStep{
		Config: testAccExampleResourceConfig(rName),
		Check:  resource.ComposeTestCheckFunc(),
	}
}

func testAccCreateDataSourceFromExampleStep(rName string) resource.TestStep {
	// skip test if no example is provided
	if testAccExampleDataSourceConfig(rName) == "" {
		return resource.TestStep{
			SkipFunc: testAccSkipTestStep,
		}
	}
	return resource.TestStep{
		Config: testAccExampleDataSourceConfig(rName),
		Check:  resource.ComposeTestCheckFunc(),
	}
}

func testAccSkipTestStep() (bool, error) {
	return true, nil
}

func testAccExampleResourceConfig(rName string) string {
	b, err := os.ReadFile(fmt.Sprintf("%s/resources/%s/resource.tf", examplesDir, rName))
	if err != nil {
		return ""
	}
	return string(b)
}

func testAccExampleDataSourceConfig(rName string) string {
	b, err := os.ReadFile(fmt.Sprintf("%s/data_sources/%s/resource.tf", examplesDir, rName))
	if err != nil {
		return ""
	}
	return string(b)
}

// toFix
// func TestMain(m *testing.M) {
// 	acctest.UseBinaryDriver("fortios", New))
// 	resource.TestMain(m)
// }
