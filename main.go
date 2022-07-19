package main

import (
	tfsdk "github.com/apparentlymart/terraform-sdk"
	"github.com/cloudandthings/terraform-provider-javascript/internal/provider"
)

func main() {
	tfsdk.ServeProviderPlugin(provider.Provider())
}
