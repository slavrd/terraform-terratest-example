# Terratest-GitHub Actions example

A simple example of [terrstest](https://github.com/gruntwork-io/terratest) usage. Contains a simple terraform code and a Golang test that leverages the `terratest` terraform module to test the TF configuration.

## Local usage

* Prerequisite: installed [golang](https://golang.org/dl/) >= 1.13
* Clone repo - `git clone https://github.com/slavrd/terraform-terratest-example.git`
* Install `terratest` package and its dependencies - `go get -v -d -t ./test/...`
* Run tests - `go test -v ./test`

## GitHub Actions

A GitHub Actions workflow is configured to run on all pull requests opened on `master` branch. The workflow will check if `terraform init` and `terraform validate` commands are successful and then execute the `terratest` tests.

Configuration for the workflow is in `.github/workflow/workflow.yml`
