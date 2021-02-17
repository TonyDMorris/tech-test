## S3 as a database

Considering that the needs of this project are to simply to store json files via a GUID and retreive them the same way, S3 is actually great choice for the POC as it is significantly cheaper than the alternatives.
In fact I have read that orders older than 12 months on amazon are actually stored in S3. 


# cfn 
the cloud formation files are not region agnostic and the instructions do not specifiy that your resources have to be in eu-west-1 , arn's should be used instead so that the we do not need to make considerations about aws regions when editing templates.

# changes

1. the main.yaml cloud formation template was not region agnostic and due to my s3 bucket region being different than the hard coded one this prevented the templates from being built.
I have updated the resources to use the current region in which the template is being executed.

