#!/usr/bin/env bash
#
# Compare SDK generation output between two Speakeasy CLI versions.
#
# Proves that Speakeasy >= 1.755 converts GET/HEAD requestBody into
# a "query" query parameter, while older versions generate a
# "RequestBody" JSON body field.
#
# Prerequisites:
#   - SPEAKEASY_API_KEY env var set (get from https://app.speakeasy.com)
#   - curl, unzip, go, git, diff
#
# Usage:
#   export SPEAKEASY_API_KEY=your-key
#   bash scripts/compare-speakeasy-versions.sh [old_version] [new_version]
#
# Defaults: old=1.635.1 (pre-3.8.0), new=1.757.1 (used for 3.8.1)

set -euo pipefail

OLD_VERSION="${1:-1.635.1}"
NEW_VERSION="${2:-1.757.1}"

if [ -z "${SPEAKEASY_API_KEY:-}" ]; then
  echo "ERROR: SPEAKEASY_API_KEY is not set."
  echo "Get your key from https://app.speakeasy.com and export it."
  exit 1
fi

WORK_DIR=$(mktemp -d)
trap 'rm -rf "$WORK_DIR"' EXIT
echo "Working in $WORK_DIR"

SPEC_URL="https://github.com/formancehq/stack/releases/download/v3.2.0/generate.json"

install_speakeasy() {
  local version=$1
  local dir="$WORK_DIR/$version/bin"
  mkdir -p "$dir"
  local os_arch
  case "$(uname -s)-$(uname -m)" in
    Linux-x86_64)  os_arch="linux_amd64" ;;
    Linux-aarch64) os_arch="linux_arm64" ;;
    Darwin-x86_64) os_arch="darwin_amd64" ;;
    Darwin-arm64)  os_arch="darwin_arm64" ;;
    *) echo "Unsupported platform"; exit 1 ;;
  esac
  local url="https://github.com/speakeasy-api/speakeasy/releases/download/v${version}/speakeasy_${os_arch}.zip"
  echo "  Downloading Speakeasy $version..."
  curl -sL -o "$dir/speakeasy.zip" "$url"
  unzip -qo "$dir/speakeasy.zip" -d "$dir"
  rm "$dir/speakeasy.zip"
  echo "  Installed: $("$dir/speakeasy" --version 2>&1 | head -1)"
}

setup_workspace() {
  local version=$1
  local ws="$WORK_DIR/$version/workspace"
  mkdir -p "$ws/.speakeasy"

  cp "$WORK_DIR/spec.json" "$ws/"

  cat > "$ws/gen.yaml" << 'YAML'
configVersion: 2.0.0
generation:
  sdkClassName: Formance
  usageSnippets:
    optionalPropertyRendering: withExample
    sdkInitStyle: constructor
  fixes:
    nameResolutionFeb2025: false
    parameterOrderingFeb2024: false
    requestResponseComponentNamesFeb2024: false
    securityFeb2025: false
  auth:
    oAuth2ClientCredentialsEnabled: true
    oAuth2PasswordEnabled: false
    hoistGlobalSecurity: true
  telemetryEnabled: false
go:
  version: 0.0.1
  additionalDependencies: {}
  allowUnknownFieldsInWeakUnions: false
  author: Formance
  clientServerStatusCodesAsErrors: true
  imports:
    option: openapi
    paths:
      callbacks: pkg/models/callbacks
      errors: pkg/models/sdkerrors
      operations: pkg/models/operations
      shared: pkg/models/shared
      webhooks: pkg/models/webhooks
  maxMethodParams: 0
  methodArguments: require-security-and-request
  packageName: github.com/formancehq/formance-sdk-go
  responseFormat: envelope
YAML

  cat > "$ws/.speakeasy/workflow.yaml" << 'YAML'
workflowVersion: 1.0.0
speakeasyVersion: latest
sources:
    stacks-source:
        inputs:
            - location: ./spec.json
targets:
    formance-sdk-go:
        target: go
        source: stacks-source
YAML

  cd "$ws"
  go mod init github.com/formancehq/formance-sdk-go/v3 2>/dev/null || true
  git init -q && git add -A && git commit -q -m "init"
}

run_generation() {
  local version=$1
  local bin="$WORK_DIR/$version/bin/speakeasy"
  local ws="$WORK_DIR/$version/workspace"
  echo "  Generating SDK with Speakeasy $version..."
  cd "$ws"
  "$bin" generate sdk \
    -s ./spec.json \
    -l go \
    -o . \
    -y 2>&1 | tail -5
}

echo "=== Step 1: Download spec ==="
curl -sL "$SPEC_URL" -o "$WORK_DIR/spec.json"
echo "  Spec downloaded ($(wc -l < "$WORK_DIR/spec.json") lines)"

echo ""
echo "=== Step 2: Install Speakeasy versions ==="
install_speakeasy "$OLD_VERSION"
install_speakeasy "$NEW_VERSION"

echo ""
echo "=== Step 3: Set up workspaces ==="
setup_workspace "$OLD_VERSION"
setup_workspace "$NEW_VERSION"

echo ""
echo "=== Step 4: Generate with old version ($OLD_VERSION) ==="
run_generation "$OLD_VERSION"

echo ""
echo "=== Step 5: Generate with new version ($NEW_VERSION) ==="
run_generation "$NEW_VERSION"

echo ""
echo "=== Step 6: Compare key files ==="
echo ""

FILES_TO_CHECK=(
  "pkg/models/operations/v2countaccounts.go"
  "pkg/models/operations/v2listaccounts.go"
  "pkg/models/operations/v2listtransactions.go"
  "pkg/models/operations/v2listlogs.go"
  "pkg/models/operations/v2getbalancesaggregated.go"
  "pkg/models/operations/v2getvolumeswithbalances.go"
  "pkg/models/operations/v3listaccounts.go"
  "pkg/models/operations/v3listpayments.go"
)

for f in "${FILES_TO_CHECK[@]}"; do
  old_file="$WORK_DIR/$OLD_VERSION/workspace/$f"
  new_file="$WORK_DIR/$NEW_VERSION/workspace/$f"

  if [ ! -f "$old_file" ] && [ ! -f "$new_file" ]; then
    echo "--- $f: NOT GENERATED in either version ---"
    continue
  fi

  echo "--- $f ---"

  if [ -f "$old_file" ]; then
    old_body=$(grep -c 'RequestBody\|request:"mediaType' "$old_file" 2>/dev/null || echo 0)
    old_query=$(grep -c 'queryParam.*serialization=json.*name=query' "$old_file" 2>/dev/null || echo 0)
    echo "  OLD ($OLD_VERSION): RequestBody refs=$old_body, Query query param=$old_query"
  else
    echo "  OLD ($OLD_VERSION): file not generated"
  fi

  if [ -f "$new_file" ]; then
    new_body=$(grep -c 'RequestBody\|request:"mediaType' "$new_file" 2>/dev/null || echo 0)
    new_query=$(grep -c 'queryParam.*serialization=json.*name=query' "$new_file" 2>/dev/null || echo 0)
    echo "  NEW ($NEW_VERSION): RequestBody refs=$new_body, Query query param=$new_query"
  else
    echo "  NEW ($NEW_VERSION): file not generated"
  fi

  if [ -f "$old_file" ] && [ -f "$new_file" ]; then
    echo "  DIFF:"
    diff --color=auto -u "$old_file" "$new_file" | grep -E '^\+.*RequestBody|^\+.*Query|^\-.*RequestBody|^\-.*Query|^\+.*request:"media|^\-.*request:"media|^\+.*queryParam.*query|^\-.*queryParam.*query' | head -10 || echo "    (no relevant diff lines)"
  fi
  echo ""
done

echo "=== Summary ==="
echo "If OLD shows RequestBody refs > 0 and NEW shows Query query param > 0,"
echo "the hypothesis is confirmed: Speakeasy $NEW_VERSION converts requestBody"
echo "on GET/HEAD endpoints into a 'query' query parameter automatically."
echo ""
echo "Full generated code is in:"
echo "  OLD: $WORK_DIR/$OLD_VERSION/workspace/"
echo "  NEW: $WORK_DIR/$NEW_VERSION/workspace/"
echo ""
echo "For a complete diff:"
echo "  diff -r $WORK_DIR/$OLD_VERSION/workspace/pkg/models/operations/ $WORK_DIR/$NEW_VERSION/workspace/pkg/models/operations/"
