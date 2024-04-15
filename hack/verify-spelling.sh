#!/usr/bin/env bash

# Copyright 2020 The Kubermatic Kubernetes Platform contributors.
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

set -euo pipefail

cd $(dirname $0)/..
source hack/lib.sh

CONTAINERIZE_IMAGE=quay.io/kubermatic/build:go-1.22-node-20-3 containerize ./hack/verify-spelling.sh

echodate "Running codespell..."

codespell \
  --skip .git,_build,_dist,vendor,go.mod,go.sum,*.jpg,*.jpeg,*.png,*.woff,*.woff2,*.pem,node_modules,src,.angular,package-lock.json,dist,coverage,swagger.json,src,./modules/api/pkg/test/e2e/utils/apiclient \
  --ignore-words .codespell.exclude \
  --check-filenames \
  --check-hidden

echodate "No problems detected."
