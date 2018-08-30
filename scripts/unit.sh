#!/usr/bin/env bash
set -euo pipefail

export ROOT="$( dirname "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" )"
$ROOT/scripts/install_tools.sh

cd $ROOT/src/binary/
ginkgo -r -skipPackage=brats,integration
