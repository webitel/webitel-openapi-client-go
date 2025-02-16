#!/usr/bin/env bash

# https://stackoverflow.com/a/246128
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Pull the schema (last commit from main branch)
SCHEMA="$(curl -s -L https://raw.githubusercontent.com/webitel/protos/refs/heads/main/swagger/api.json)"

# Custom extensions: https://goswagger.io/go-swagger/reference/models/schemas/#custom-extensions
modify() {
    SCHEMA="$(echo "${SCHEMA}" | jq "${1}")"
}

modify '(.paths."/roles/metadata".get.parameters[] | select(.name == "merge.struct_value") | .type) |= "string"'

# Fields "comments", "languages", "photos" doesn't described in Contact GraphQL model (read-only),
# so per now we can skip it and add omitempty=true to json tag in PATCH /contacts/{etag} request body.
modify '(.paths."/contacts/{etag}".patch.parameters[]
        | select(.name == "input")
        | .schema.properties)
        |= if . then with_entries(if .key | IN("photos", "comments", "languages") then
            .value += {"x-omitempty": true} else . end
        ) else . end'




# Write the schema to a file
echo "${SCHEMA}" > "${SCRIPT_DIR}/schema.json"
