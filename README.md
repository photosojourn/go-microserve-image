## Go-Microservice

This is an extenstion to the `go-microservice' which now serves an image after pulling it from an S3 bucket. The bucket and region are set via envrionment variables via the Helm chart.

### Environment

* `AWS_REGION` - The AWS region your running in
* `S3_BUCKET` - Name of the S3 bucket to pull the image from

### AWS CodeBuild specs

This project also contains two example AWS CodeBuild spec files listed below:

* buildspec.yaml - Builds and pushs the image to ECR
* buildspec-helm.yaml - Deploys the latest image to EKS via Helm chart

### Exmaple Terraform

Example terraform code can be found here: [terraform-examples/go-microservice-tf](https://github.com/photosojourn/terraform-examples/tree/master/go-microservice-tf)