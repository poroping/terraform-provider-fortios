# FortiOS Terraform Provider

<!-- - Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool) -->

<!-- <img src="https://camo.githubusercontent.com/1a4ed08978379480a9b1ca95d7f4cc8eb80b45ad47c056a7cfb5c597e9315ae5/68747470733a2f2f7777772e6461746f636d732d6173736574732e636f6d2f323838352f313632393934313234322d6c6f676f2d7465727261666f726d2d6d61696e2e737667" width="600px">
<img src ="https://avatars.githubusercontent.com/u/39314919?s=200&v=4"> -->

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.15.x +
- [Go](https://golang.org/doc/install) 1.17.x (to build the provider plugin)

## Building the Provider

1. Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-fortios`.

    ```sh
    mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
    git clone git@github.com:poroping/terraform-provider-fortios
    ```

2. Enter the provider directory and build the provider.

    ```sh
    cd $GOPATH/src/github.com/terraform-providers/terraform-provider-fortios
    go build
    ```

## Using the Provider

If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it into your plugins directory,  run `terraform init --upgrade` to initialize it.

```sh
terraform init --upgrade
```

## Developing the Provider

If you wish to work on the provider, you'll first need Go installed on your machine (version 1.17+ is required). You'll also need to correctly setup a GOPATH, as well as adding $GOPATH/bin to your $PATH.

To compile the provider, run `go install`. This will build the provider and put the provider binary in the $GOPATH/bin directory.

```sh
go install
```

## Versioning

The provider was generated from schemas covering 6.2, 6.4 and 7.0. Acceptance testing is performed with the latest GA release.

## Contributions and Bugs

Feel free to raise an issue for a bug or feature request.

### Collecting debug information

This information will be crucial to quickly resolving any issues.

- SSH to your FortiGate. Close any open GUI windows as this will generate spam during the debug. Run the following commands:

    ```
    diagnose debug enable
    diagnose debug application httpsd -1
    diagnose debug cli 8
    ```

- Run terraform apply with debugging enabled:

    ```sh
    TF_LOG=debug terraform apply
    ```