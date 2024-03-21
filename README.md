# CDK, Go Backend for Registering Users, handles logins, and returns a Bearer token for protected routes.

In this project I setup a go backend to register users add them to the database with encryption utilizing bcyrpt. 

## How It's Made:

**Tech used:**  Golang, AWS cdk, lambda, api gatway and DynamoDB as the database. 

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests
