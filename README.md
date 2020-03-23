## Go-Microservice

This is an example Golang microservice that was written with sole purpose of building and testikng CI/CD pipelines. Litrally it returns just a string. Thats it.

To over-ride the default message you can get my setting the enviroment variable `GO_MS_MESSAGE`.

### AWS CodeBuild specs

This project also contains two example AWS CodeBuild spec files listed below:

* buildspec.yaml - Builds and pushs the image to ECR
* buildspec-helm.yaml - Deploys the latest image to EKS via Helm chart

### Exmaple Terraform

Example terraform code can be found here: [terraform-examples/go-microservice-tf](https://github.com/photosojourn/terraform-examples/tree/master/go-microservice-tf)