#!/bin/bash
sleep 10
dithercli tx dither createChannel general --from=me --yes=true
sleep 6
dithercli tx dither createChannel random --from=me --yes=true 
sleep 6
dithercli tx dither createPost 'Hello, world' $(dithercli q dither listChannels | jq -r '.[0].id') --from=me --yes=true
sleep 6
dithercli tx dither createPost 'This is another post.' $(dithercli q dither listChannels | jq -r '.[0].id') --from=me --yes=true
