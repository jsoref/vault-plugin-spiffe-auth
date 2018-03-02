#!/bin/bash

plugin_directory=$(pwd)

tee /tmp/vault.hcl <<EOF
plugin_directory = "$plugin_directory/bin/"
EOF

