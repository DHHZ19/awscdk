# CDK Go project!

In this project I setup a go backend that to register users add them to the database with increpation utilizing bcyrpt. 

## How It's Made:

**Tech used:**  Golang, AWS cdk, lambda, api gatway and DynamoDB as the database. 

## Lessons Learned:

I learned that when creating an application with many diffrent data inputs there can be many ways to send handle data coming from your database displaying that in the views. There are conventions like MVC that help with the strucure of your web apps speeding up development time and mental overhead.


The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests
