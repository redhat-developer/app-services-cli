
package error

// kafkamgmt error codes 
  
const (

  // Forbidden to perform this action
  ERROR_4 string = "KAFKAS-MGMT-4"

  // Forbidden to create more instances than the maximum allowed
  ERROR_5 string = "KAFKAS-MGMT-5"

  // An entity with the specified unique values already exists
  ERROR_6 string = "KAFKAS-MGMT-6"

  // Resource not found
  ERROR_7 string = "KAFKAS-MGMT-7"

  // General validation failure
  ERROR_8 string = "KAFKAS-MGMT-8"

  // Unspecified error
  ERROR_9 string = "KAFKAS-MGMT-9"

  // HTTP Method not implemented for this endpoint
  ERROR_10 string = "KAFKAS-MGMT-10"

  // Account is unauthorized to perform this action
  ERROR_11 string = "KAFKAS-MGMT-11"

  // Required terms have not been accepted
  ERROR_12 string = "KAFKAS-MGMT-12"

  // Account authentication could not be verified
  ERROR_15 string = "KAFKAS-MGMT-15"

  // Unable to read request body
  ERROR_17 string = "KAFKAS-MGMT-17"

  // Unable to perform this action, as the service is currently under maintenance
  ERROR_18 string = "KAFKAS-MGMT-18"

  // Bad request
  ERROR_21 string = "KAFKAS-MGMT-21"

  // Failed to parse search query
  ERROR_23 string = "KAFKAS-MGMT-23"

  // The maximum number of allowed kafka instances has been reached
  ERROR_24 string = "KAFKAS-MGMT-24"

  // Resource gone
  ERROR_25 string = "KAFKAS-MGMT-25"

  // Provider not supported
  ERROR_30 string = "KAFKAS-MGMT-30"

  // Region not supported
  ERROR_31 string = "KAFKAS-MGMT-31"

  // Kafka cluster name is invalid
  ERROR_32 string = "KAFKAS-MGMT-32"

  // Minimum field length not reached
  ERROR_33 string = "KAFKAS-MGMT-33"

  // Maximum field length has been depassed
  ERROR_34 string = "KAFKAS-MGMT-34"

  // Kafka cluster name is already used
  ERROR_36 string = "KAFKAS-MGMT-36"

  // Field validation failed
  ERROR_37 string = "KAFKAS-MGMT-37"

  // Service account name is invalid
  ERROR_38 string = "KAFKAS-MGMT-38"

  // Service account desc is invalid
  ERROR_39 string = "KAFKAS-MGMT-39"

  // Service account id is invalid
  ERROR_40 string = "KAFKAS-MGMT-40"

  // Instance Type not supported
  ERROR_41 string = "KAFKAS-MGMT-41"

  // Instance plan not supported
  ERROR_42 string = "KAFKAS-MGMT-42"

  // Billing account id missing or invalid
  ERROR_43 string = "KAFKAS-MGMT-43"

  // Enterprise cluster ID is already used
  ERROR_44 string = "KAFKAS-MGMT-44"

  // Enterprise cluster ID is invalid
  ERROR_45 string = "KAFKAS-MGMT-45"

  // Enterprise external cluster ID is invalid
  ERROR_46 string = "KAFKAS-MGMT-46"

  // Dns name is invalid
  ERROR_47 string = "KAFKAS-MGMT-47"

  // Synchronous action is not supported, use async=true parameter
  ERROR_103 string = "KAFKAS-MGMT-103"

  // Failed to create kafka client in the mas sso
  ERROR_106 string = "KAFKAS-MGMT-106"

  // Failed to get kafka client secret from the mas sso
  ERROR_107 string = "KAFKAS-MGMT-107"

  // Failed to get kafka client from the mas sso
  ERROR_108 string = "KAFKAS-MGMT-108"

  // Failed to delete kafka client from the mas sso
  ERROR_109 string = "KAFKAS-MGMT-109"

  // Failed to create service account
  ERROR_110 string = "KAFKAS-MGMT-110"

  // Failed to get service account
  ERROR_111 string = "KAFKAS-MGMT-111"

  // Failed to delete service account
  ERROR_112 string = "KAFKAS-MGMT-112"

  // Failed to find service account
  ERROR_113 string = "KAFKAS-MGMT-113"

  // Max limit for the service account creation has reached
  ERROR_115 string = "KAFKAS-MGMT-115"

  // Insufficient quota
  ERROR_120 string = "KAFKAS-MGMT-120"

  // Failed to check quota
  ERROR_121 string = "KAFKAS-MGMT-121"

  // Too many requests
  ERROR_429 string = "KAFKAS-MGMT-429"

  // An unexpected error happened, please check the log of the service for details
  ERROR_1000 string = "KAFKAS-MGMT-1000"

)