# kafka-cli-utility
- A Golang CLI tool to manage Apache Kafka

## How to configure
Once you have started your cluster, you can use the CLI tool to easily manage it. Just connect to the cluster by providing localhost to bootstrap.servers in librdkafka.config properties file

## Release v1
1. Create and Delete Topics
2. Create and manage Consumer groups
3. Create and manage Producer

We are using confluent-kafka-go, a Confluent's Golang client for Apache Kafka and the Confluent Platform.

Features:

**High performance** - confluent-kafka-go is a lightweight wrapper around librdkafka, a finely tuned C client.

**Reliability** - There are a lot of details to get right when writing an Apache Kafka client. We get them right in one place (librdkafka) and leverage this work across all of our clients (also confluent-kafka-python and confluent-kafka-dotnet).

**Supported** - Commercial support is offered by Confluent.

Future proof - Confluent, founded by the creators of Kafka, is building a streaming platform with Apache Kafka at its core. It's high priority for us that client features keep pace with core Apache Kafka and components of the Confluent Platform.

The Golang bindings provides a high-level Producer and Consumer with support for the balanced consumer groups of Apache Kafka 0.9 and above.

See the [API documentation](https://docs.confluent.io/platform/current/clients/confluent-kafka-go/index.html) for more information.

## Getting Started
Supports Go 1.11+ and librdkafka 1.6.0+.

Starting with Go 1.13, you can use Go Modules to install confluent-kafka-go.

1. Download the zip file or clone the repository on your system
2. Build your project:

```java
go build ./...
```
A dependency to the latest stable version of confluent-kafka-go should be automatically added to your go.mod file.

If Go modules can't be used we recommend manually installing gopkg.in:
```java
go get -u gopkg.in/confluentinc/confluent-kafka-go.v1/kafka
```
## librdkafka
Prebuilt librdkafka binaries are included with the Go client and librdkafka does not need to be installed separately on the build or target system. The following platforms are supported by the prebuilt librdkafka binaries:

Mac OSX x64
glibc-based Linux x64 (e.g., RedHat, Debian, CentOS, Ubuntu, etc) - without GSSAPI/Kerberos support
musl-based Linux 64 (Alpine) - without GSSAPI/Kerberos support

## Installing librdkafka
If the bundled librdkafka build is not supported on your platform, or you need a librdkafka with **GSSAPI/Kerberos** support, you must install librdkafka manually on the build and target system using one of the following alternatives:

- For Debian and Ubuntu based distros, install librdkafka-dev from the standard repositories or using Confluent's Deb repository.
- For Redhat based distros, install librdkafka-devel using Confluent's YUM repository.
- For MacOS X, install librdkafka from Homebrew. You may also need to brew install pkg-config if you don't already have it: brew install librdkafka pkg-config.
- For Alpine: apk add librdkafka-dev pkgconf

**Note**: confluent-kafka-go is not supported on Windows.
For source builds, see instructions below.
```java
git clone https://github.com/edenhill/librdkafka.git
cd librdkafka
./configure
make
sudo make install
```
After installing librdkafka you will need to build your Go application with -tags dynamic.

**Note**: If you use the master branch of the Go client, then you need to use the master branch of librdkafka.
***confluent-kafka-go requires librdkafka v1.6.0 or later.**
