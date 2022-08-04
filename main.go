package main

import (
	tfsdk "github.com/apparentlymart/terraform-sdk"
	"github.com/lewispeckover/terraform-provider-javascript/internal/provider"
)

func main() {
	tfsdk.ServeProviderPlugin(provider.Provider())
}
