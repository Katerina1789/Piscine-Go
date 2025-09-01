#! /bin/bash
jq '.[] | select(.id == 170) | "\(.name) - \(.power) - \(.gender)"' superhero.json 
