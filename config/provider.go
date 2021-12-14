/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-jet-linode/config/lke"
)

const (
	resourcePrefix = "linode"
	modulePath     = "github.com/crossplane-contrib/provider-jet-linode"
)

// GetProvider returns provider configuration
func GetProvider(resourceMap map[string]*schema.Resource) *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

        // https://github.com/linode/terraform-provider-linode/tree/main/linode
	pc := tjconfig.NewProvider(resourceMap, resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
                tjconfig.WithIncludeList([]string{
                  "lke$",
                }))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
                lke.Customize,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
