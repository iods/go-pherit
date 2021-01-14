## Installation

How to install the Go language.

### On Linux

Navigate to your home directory, and grab the Go source:

``sh
wget -c https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
``

https://www.tecmint.com/install-go-in-linux/


The go build command compiles source code into a binary format that computers can execute.

The go run command compiles the source code but does not save an executable locally.


## FAQs

**Where do I save my source code?**

Under my $GOPATH/src ($GOPATH can be anywhere)

**What does $GOPATH mean?**

Stores the Go source code files and compiled packages.

**Do you need to set this $GOPATH?**

Yes, set it yourself.

**How can you print it?**

Using the `go env GOPATH` command.