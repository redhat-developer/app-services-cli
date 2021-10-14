#!/bin/bash

## Kafka ACL grant permissions

## Prerequisites
## Kafka instance created and in context

## Effects
## This command will create various ACLs on the Kafka instance that needs to be cleaned up manually

## Framework

alias rhoas=$(go env GOPATH)/bin/rhoas

## Cases
rhoas kafka acl grant-permissions 


rhoas kafka acl grant-permissions --producer


rhoas kafka acl grant-permissions --producer --consumer


rhoas kafka acl grant-permissions --producer --service-account test_case --user test_case


rhoas kafka acl grant-permissions --consumer --service-account test_case --topic test_case --group test_case --group-prefix test_case


rhoas kafka acl grant-permissions --producer --service-account test_case --topic test_case --topic-prefix test_case


## Cases that require cleanup 

rhoas kafka acl grant-permissions --producer --service-account test_case --topic test_case 


rhoas kafka acl grant-permissions --producer --service-account test_prefix --topic-prefix test_ 


rhoas kafka acl grant-permissions --producer --user test_user --topic-prefix test_ 


rhoas kafka acl grant-permissions --producer --user test_user --topic test_topic

rhoas kafka acl grant-permissions --producer --user all --topic all

rhoas kafka acl grant-permissions --consumer --service-account test_prefix --topic-prefix test_ 


rhoas kafka acl grant-permissions --consumer --service-account test_consumer --topic-prefix test_  --group test_group


rhoas kafka acl grant-permissions --consumer --service-account test_prefix --topic-prefix test_  --group-prefix test_


rhoas kafka acl grant-permissions --consumer --producer --service-account consumer-producer --topic-prefix test_  --group-prefix test_

