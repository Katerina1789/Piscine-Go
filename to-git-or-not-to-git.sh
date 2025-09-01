#! /bin/bash
curl https://platform.zone01.gr/assets/superhero/all.json
jq '.[] | select(.id == 170) | "\(.name) - \(.power) - \(.gender)"'