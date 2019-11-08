# Terratest example

A simple example of [terrstest]() usage. Contains a simple terraform code and a Golang test that leverages the `terratest` terraform module to test the TF configuration.

## Local usage

* Prerequisite: installed [golang]() >= 1.13
* Clone repo - `git clone https://github.com/slavrd/terraform-terratest-example.git`
* Install `terratest` package and its dependencies - `go get -v -d -t ./test/...`
* Run tests - `go test -v ./test`
