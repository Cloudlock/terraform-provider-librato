Terraform Provider
==================

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/Cloudlock/terraform-provider-librato`

```sh
$ mkdir -p $GOPATH/src/github.com/Cloudlock; cd $GOPATH/src/github.com/Cloudlock
$ git clone git@github.com:Cloudlock/terraform-provider-librato
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/Cloudlock/terraform-provider-librato
$ make build
```

Using the provider
----------------------
```sh
$ make build
cp $GOPATH/bin/terraform-provider-librato ~/.terraform.d/plugins/
```


Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-librato
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

Install it to use with Terraform
```
cp $GOPATH/bin/terraform-provider-librato ~/.terraform.d/plugins/
```

