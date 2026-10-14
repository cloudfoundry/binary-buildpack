# Cloud Foundry Binary Buildpack

[![CF Slack](https://www.google.com/s2/favicons?domain=www.slack.com) Join us on Slack](https://cloudfoundry.slack.com/messages/buildpacks/)

A Cloud Foundry [buildpack](http://docs.cloudfoundry.org/buildpacks/) for running arbitrary binary web servers.

### Buildpack User Documentation

Official buildpack documentation can be found at [binary buildpack docs](http://docs.cloudfoundry.org/buildpacks/binary/index.html).

### Building the Buildpack

To build this buildpack, run the following command from the buildpack's directory:

1. Source the .envrc file in the buildpack directory.

   ```bash
   source .envrc
   ```
   To simplify the process in the future, install [direnv](https://direnv.net/) which will automatically source .envrc when you change directories.

1. Install buildpack-packager

    ```bash
    go install github.com/cloudfoundry/libbuildpack/packager/buildpack-packager
    ```

1. Build the buildpack

    ```bash
    buildpack-packager build [ --cached=(true|false) ]
    ```

1. Use in Cloud Foundry

   Upload the buildpack to your Cloud Foundry and optionally specify it by name

    ```bash
    cf create-buildpack [BUILDPACK_NAME] [BUILDPACK_ZIP_FILE_PATH] 1
    cf push my_app [-b BUILDPACK_NAME]
    ```

### Testing

Buildpacks use the [Cutlass](https://github.com/cloudfoundry/libbuildpack/tree/master/cutlass) framework for running integration tests.

To test this buildpack, run the following command from the buildpack's directory:

1. Source the .envrc file in the buildpack directory.

   ```bash
   source .envrc
   ```
   To simplify the process in the future, install [direnv](https://direnv.net/) which will automatically source .envrc when you change directories.

1. Run unit tests

    ```bash
    ./scripts/unit.sh
    ```

1. Run integration tests

    ```bash
    ./scripts/integration.sh
    ```

### Dynatrace Integration

This buildpack can automatically inject the [Dynatrace OneAgent](https://www.dynatrace.com/support/help/technology-support/cloud-platforms/cloud-foundry/) into your application. When a Dynatrace service is bound to the app, the buildpack downloads the OneAgent during staging and configures `LD_PRELOAD` (via `profile.d/dynatrace-env.sh`) so the agent is loaded into your binary at launch.

Bind a Dynatrace user-provided service (the service name must contain `dynatrace`):

```bash
cf create-user-provided-service dynatrace -p '{"environmentid":"<env-id>","apitoken":"<paas-token>"}'
cf bind-service my_app dynatrace
cf restage my_app
```

Supported credential fields:

| Key             | Type    | Description                                                                                             | Required | Default         |
| --------------- | ------- | ------------------------------------------------------------------------------------------------------- | -------- | --------------- |
| environmentid   | string  | The ID for the Dynatrace environment.                                                                   | Yes      | N/A             |
| apitoken        | string  | The API Token for the Dynatrace environment.                                                            | Yes      | N/A             |
| apiurl          | string  | Overrides the default Dynatrace API URL to connect to.                                                  | No       | Default API URL |
| skiperrors      | boolean | If `true`, staging does not fail when the OneAgent download fails.                                      | No       | false           |
| networkzone     | string  | If set, the agent is configured to use communication endpoints located in this network zone.           | No       | empty           |
| enablefips      | boolean | If `true`, FIPS 140-2 mode is enabled.                                                                  | No       | false           |
| addtechnologies | string  | Comma-separated list of additional OneAgent code modules to download (e.g. `go`, `java`, `nodejs`).     | No       | empty           |

By default the buildpack downloads the generic `process` code module. Since the
binary buildpack runs an opaque binary, it cannot detect the application's
language. For language-specific code-level insights, set `addtechnologies`
accordingly — e.g. `"addtechnologies":"go"` for a Go binary. Note that OneAgent
injection relies on `LD_PRELOAD`, so the binary must be **dynamically linked**
(for Go, built with `CGO_ENABLED=1`); a fully statically-linked binary ignores
`LD_PRELOAD` and will not be instrumented.

### Contributing

Find our guidelines [here](./CONTRIBUTING.md).

### Help and Support

Join the #buildpacks channel in our [Slack community](http://slack.cloudfoundry.org/) if you need any further assistance.

### Reporting Issues

Open a GitHub issue on this project [here](https://github.com/cloudfoundry/binary-buildpack/issues/new).
