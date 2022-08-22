#!/usr/bin/env bash

# Copyright 2021 The Crossplane Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Please set ProviderNameLower & ProviderNameUpper environment variables before running this script.
# See: https://github.com/crossplane/terrajet/blob/main/docs/generating-a-provider.md
set -euo pipefail

REPLACE_FILES='./* ./.github :!build/** :!go.* :!hack/prepare.sh'
# shellcheck disable=SC2086
git grep -l 'jet.template' -- ${REPLACE_FILES} | xargs sed -i.bak "s/jet.template/${ProviderNameLower}/g"
git grep -l 'template' -- ${REPLACE_FILES} | xargs sed -i.bak "s/template/${ProviderNameLower}/g"
git grep -l 'template' -- ${REPLACE_FILES} | xargs sed -i.bak "s/jet.//g"
git grep -l 'contrib' -- ${REPLACE_FILES} | xargs sed -i.bak "s/crossplane-contrib/${RepoName}/g"
git grep -l 'jet.crossplane' -- ${REPLACE_FILES} | xargs sed -i.bak "s/jet.//g"
# shellcheck disable=SC2086
git grep -l 'Template' -- ${REPLACE_FILES} | xargs sed -i.bak "s/Template/${ProviderNameUpper}/g"
git grep -l 'Jet.Template' -- ${REPLACE_FILES} | xargs sed -i.bak "s/Template/${ProviderNameUpper}/g"
git grep -l 'Template' -- ${REPLACE_FILES} | xargs sed -i.bak "s/Template/${ProviderNameUpper}/g"
git grep -l 'Template' -- ${REPLACE_FILES} | xargs sed -i.bak "s/jet/apig/g"
git grep -l 'Jet.Template' -- ${REPLACE_FILES} | xargs sed -i.bak "s/Jet.Template/${ProviderNameUpper}/g"
# We need to be careful while replacing "template" keyword in go.mod as it could tamper
# some imported packages under require section.
sed -i.bak "s/provider-jet-template/provider-${ProviderNameLower}/g" go.mod
sed -i.bak "s/crossplane-contrib/${RepoName}/g" go.mod

# Clean up the .bak files created by sed
git clean -fd

git mv "internal/clients/template.go" "internal/clients/${ProviderNameLower}.go"
git mv "cluster/images/provider-jet-template" "cluster/images/provider-${ProviderNameLower}"
git mv "cluster/images/provider-jet-template-controller" "cluster/images/provider-${ProviderNameLower}-controller"

# We need to remove this api folder otherwise first `make generate` fails with
# the following error probably due to some optimizations in go generate with v1.17:
# generate: open /Users/hasanturken/Workspace/crossplane-contrib/provider-jet-template/apis/null/v1alpha1/zz_generated.deepcopy.go: no such file or directory
rm -rf apis/null