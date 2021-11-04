#!/bin/bash

## Terms and Conditions acceptance

## Requirements
## Create new account without any terms and conditions accepted.

## Framework

alias rhoas=$(go env GOPATH)/bin/rhoas

## Cases

rhoas service-registry create --name=test -v
## In order to be able to create a new instance, you must first review and accept the terms and conditions:
## https://www.redhat.com/wapps/tnc/ackrequired?site=ocm&event=register

rhoas kafka create --name=test --provider=aws --region=eu-east1 -v
## In order to be able to create a new instance, you must first review and accept the terms and conditions:
## https://www.redhat.com/wapps/tnc/ackrequired?site=ocm&event=onlineService