## Name 
## Requirements

 - Running Kafka instance
 - Create new account 

## Cases

rhoas kafka acl create --all-accounts --permission allow --operation all --topic all -y
> The following ACL rule is going to be created:
  PRINCIPAL      PERMISSION   OPERATION   DESCRIPTION   
 -------------- ------------ ----------- -------------- 
  All accounts   ALLOW        ALL         TOPIC is "*"  

rhoas kafka acl delete --all-accounts --permission allow --operation all --topic "*"
> ? All ACLs matching the criteria provided will be deleted from the Kafka instance "test". Are you sure you want to proceed? Yes 
> Deleting ACLs from Kafka instance "test" 
✔️  Deleted 1 ACL from Kafka instance "test"
The following ACL rules were deleted:

  PRINCIPAL      PERMISSION   OPERATION   DESCRIPTION   
 -------------- ------------ ----------- -------------- 
  All accounts   ALLOW        ALL         TOPIC is "*"  


