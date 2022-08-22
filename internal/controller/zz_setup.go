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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	alias "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/apigee/alias"
	cache "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/apigee/cache"
	company "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/apigee/company"
	developer "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/apigee/developer"
	keystore "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/apigee/keystore"
	product "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/apigee/product"
	proxy "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/apigee/proxy"
	reference "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/apigee/reference"
	role "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/apigee/role"
	user "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/apigee/user"
	app "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/company/app"
	appcredential "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/company/appcredential"
	developercompany "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/company/developer"
	appdeveloper "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/developer/app"
	appcredentialdeveloper "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/developer/appcredential"
	kvm "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/environment/kvm"
	resourcefile "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/environment/resourcefile"
	kvmorganization "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/organization/kvm"
	resourcefileorganization "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/organization/resourcefile"
	providerconfig "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/providerconfig"
	deployment "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/proxy/deployment"
	kvmproxy "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/proxy/kvm"
	policy "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/proxy/policy"
	resourcefileproxy "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/proxy/resourcefile"
	permission "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/role/permission"
	flow "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/shared/flow"
	flowdeployment "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/shared/flowdeployment"
	server "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/target/server"
	roleuser "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/user/role"
	host "github.com/SuryaLeon/provider-apigeebasictwo/internal/controller/virtual/host"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alias.Setup,
		cache.Setup,
		company.Setup,
		developer.Setup,
		keystore.Setup,
		product.Setup,
		proxy.Setup,
		reference.Setup,
		role.Setup,
		user.Setup,
		app.Setup,
		appcredential.Setup,
		developercompany.Setup,
		appdeveloper.Setup,
		appcredentialdeveloper.Setup,
		kvm.Setup,
		resourcefile.Setup,
		kvmorganization.Setup,
		resourcefileorganization.Setup,
		providerconfig.Setup,
		deployment.Setup,
		kvmproxy.Setup,
		policy.Setup,
		resourcefileproxy.Setup,
		permission.Setup,
		flow.Setup,
		flowdeployment.Setup,
		server.Setup,
		roleuser.Setup,
		host.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
