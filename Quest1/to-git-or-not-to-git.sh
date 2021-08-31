#!/bin/bash
var="$(curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"damiron4\"}}){id}}"}')"
echo "${var:23:4}"
