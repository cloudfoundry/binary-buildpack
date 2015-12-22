# Cloud Foundry buildpack: Binary Files

A Cloud Foundry [buildpack](http://docs.cloudfoundry.org/buildpacks/) for running arbitrary binary web servers.

Additional information can be found at [CloudFoundry.org](http://docs.cloudfoundry.org/buildpacks/).

## Usage

Unlike most other Cloud Foundry buildpacks, the binary buildpack *must* be specified in order for it to be used in staging your binary file. You can specify the buildpack your app should use with the `-b` option for `cf push`:

```bash
cf push my_app -b https://github.com/cloudfoundry/binary-buildpack.git
```

There are two ways to provide Cloud Foundry with the shell command to execute your binary:

### Procfile

Include in your app's root directory a `Procfile` that specifies a `web` task:

```yaml
web: ./app
```

### Command line

Alternatively, you can provide the start command when deploying your app with `cf push` by including a `-c` option:

```bash
cf push my_app -c './app' -b binary-buildpack
```

## Compiling your Binary

Your binary is expected to bind to the port specified in the `PORT` environment variable. Here is an example in [Go](https://golang.org/):

```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", "world!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
```

Your binary should run without any additional runtime dependencies on the cflinuxfs2 or lucid64 root filesystem (rootfs). Any such dependencies should be statically linked to the binary.

To boot a docker container running the cflinuxfs2 filesystem:

```bash
docker run -it cloudfoundry/cflinuxfs2 bash
```

To boot a docker container running the lucid64 filesystem:

```bash
docker run -it cloudfoundry/lucid64 bash
```

To compile the above Go application on the rootfs, golang must be installed. `apt-get install golang` and `go build app.go` will produce an `app` binary.

When deploying your binary to Cloud Foundry, you can specify the root filesystem it should run against with `cf push`'s `-s` option flag:

```bash
cf push my_app -s (cflinuxfs2|lucid64)
```

To run docker on Mac OS X, we recommend [boot2docker](http://boot2docker.io/).

## Contributing

Find our guidelines [here](./CONTRIBUTING.md).

## Help and Support

Join the #buildpacks channel in our [Slack community] (http://slack.cloudfoundry.org/) 

## Reporting Issues

To report an issue with the buildpack, open an issue on this GitHub repo.
