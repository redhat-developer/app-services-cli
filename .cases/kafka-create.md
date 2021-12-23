## Kafka provider verification

## Requirements

 - Create new account 

## Cases

rhoas command
> output

rhoas kafka create --name=wtrocki --provider=azure
❌ azure is not a valid or currently enabled cloud provider name. Valid names are: aws 

rhoas kafka create --name=wtrocki --provider=aws --region="magentadrive"
❌ Magentadrive is not a valid or enabled region name.
Valid regions: eu-west-1,us-east-1
Run the command in verbose mode using the -v flag to see more information

