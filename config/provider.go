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
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	/*"github.com/SuryaLeon/provider-apigeebasictwo/config/apigee"
	"github.com/SuryaLeon/provider-apigeebasictwo/config/company"
	"github.com/SuryaLeon/provider-apigeebasictwo/config/developer"
	"github.com/SuryaLeon/provider-apigeebasictwo/config/environment"
	"github.com/SuryaLeon/provider-apigeebasictwo/config/null"
	"github.com/SuryaLeon/provider-apigeebasictwo/config/organization"
	"github.com/SuryaLeon/provider-apigeebasictwo/config/proxy"
	
	"github.com/SuryaLeon/provider-apigeebasictwo/config/shared"
	"github.com/SuryaLeon/provider-apigeebasictwo/config/target" */
	"github.com/SuryaLeon/provider-apigeebasictwo/config/user"
	"github.com/SuryaLeon/provider-apigeebasictwo/config/role"
	//"github.com/SuryaLeon/provider-apigeebasictwo/config/virtual"
)

const (
	resourcePrefix = "apigeebasictwo"
	modulePath     = "github.com/SuryaLeon/provider-apigeebasictwo"
)

//go:embed schema.json
var providerSchema string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList([]string{
			//"apigeebasictwo_apigee$",
			//"apigeebasictwo_company$",
			//"apigeebasictwo_developer$",
		/*	"apigeebasictwo_environment$",
			//"apigeebasictwo_null$",
			"apigeebasictwo_organization$",
			"apigeebasictwo_proxy$",
			
			"apigeebasictwo_shared$",
			"apigeebasictwo_target$", */
			"apigee_user$",
			"apigee_role$",
			"apigee_role_permission$",
			//"apigeebasictwo_virtual$",
		}))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		//apigee.Configure,
		//company.Configure,
		//developer.Configure,
        /*environment.Configure,
		//null.Configure,
		organization.Configure,
		proxy.Configure,
		
		shared.Configure,
		target.Configure, */
		user.Configure,
		role.Configure,
		//rolepermission.Configure,
		//virtual.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
