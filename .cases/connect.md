## Name 

Verify if cli uses custom namespace 

## Requirements

 - Running Kafka instance

## Cases

kubectl create namespace new-test  

rhoas cluster connect --namespace=new-test --ignore-context --service-type=kafka -v

kubectl get KafkaConnection