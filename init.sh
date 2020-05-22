#!/bin/bash
rm -r ~/.dithercli
rm -r ~/.ditherd

ditherd init mynode --chain-id dither

sed -i '' 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' ~/.ditherd/config/config.toml

dithercli config keyring-backend test

dithercli keys add me
dithercli keys add you

ditherd add-genesis-account $(dithercli keys show me -a) 1000foo,100000000stake
ditherd add-genesis-account $(dithercli keys show you -a) 1foo

dithercli config chain-id dither
dithercli config output json
dithercli config indent true
dithercli config trust-node true
sed -i '' 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' ~/.dithercli/config/config.toml

ditherd gentx --name me --keyring-backend test
ditherd collect-gentxs