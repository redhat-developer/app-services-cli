# Kafka ACL grant permissions

## Requirements
- Kafka instance created and in context
 
## Cases

rhoas kafka acl grant-access 

rhoas kafka acl grant-access --producer


rhoas kafka acl grant-access --producer --consumer


rhoas kafka acl grant-access --producer --service-account test_case --user test_case


rhoas kafka acl grant-access --consumer --service-account test_case --topic test_case --group test_case --group-prefix test_case


rhoas kafka acl grant-access --producer --service-account test_case --topic test_case --topic-prefix test_case


### Cases that require cleanup 

rhoas kafka acl grant-access --producer --service-account test_case --topic test_case 


rhoas kafka acl grant-access --producer --service-account test_prefix --topic-prefix test_ 


rhoas kafka acl grant-access --producer --user test_user --topic-prefix test_ 


rhoas kafka acl grant-access --producer --user test_user --topic test_topic

rhoas kafka acl grant-access --producer --all-accounts --topic all

rhoas kafka acl grant-access --consumer --service-account test_prefix --topic-prefix test_ 


rhoas kafka acl grant-access --consumer --service-account test_consumer --topic-prefix test_  --group test_group


rhoas kafka acl grant-access --consumer --service-account test_prefix --topic-prefix test_  --group-prefix test_


rhoas kafka acl grant-access --consumer --producer --service-account consumer-producer --topic-prefix test_  --group-prefix test_

