#! /bin/bash

curl -s https://gp.ynov-bordeaux.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"PVERRIER\"}}){id}}"}' | jq '.data.user[0].id'
