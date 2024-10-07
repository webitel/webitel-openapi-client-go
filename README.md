# Webitel HTTP OpenAPI Client for Go

This HTTP Go client for [Webitel](https://github.com/webitel) is generated from [Webitel's OpenAPI specification](https://github.com/webitel/protos/blob/main/swagger/api.json) using [swagger for Go](https://github.com/go-swagger/go-swagger).

## Generate the client

Once bingo & swagger are installed (see [Dependencies](#dependencies)), generate the client for _all Webitel APIs_, with the following `make` command:

```bash
make generate-client
```

This runs the Swagger generation command.

To generate the client for a _specific Webitel API_, find the name of its tag and model in the [Webitel OpenAPI specification](https://github.com/webitel/protos/blob/main/swagger/api.json). Then, set those as environment variables and run the command to generate it:
```bash
export API_TAG=CalendarService
export MODEL=engineCalendar
make generate-client
```

### How to use custom templates

In order to generate the client, `go-swagger` uses default templates. These templates can be customised to add custom configuration that are applied each time the client is generated.

For more information, check out the `go-swagger` docs on how to [use custom templates](https://github.com/go-swagger/go-swagger/blob/master/docs/generate/templates.md). The default template definitions for the client can be found in [go-swagger/generator/templates/client/](https://github.com/go-swagger/go-swagger/tree/master/generator/templates/client).

In this project, the custom templates can be found in `templates/`. They are provided to the generation command through the flag `--template-dir=templates`.

The custom templates provide added functionality for things such as authentication, TLS/SSL, retries, and custom error handling.

## Build the client

### Configuration

The client has the following friendly configuration options:

```go
import goapi "github.com/webitel/webitel-openapi-client-go/client"

cfg := &goapi.TransportConfig{
    // Host is the doman name or IP address of the host that serves the API.
    Host:       "dev.webitel.com",
    
    // BasePath is the URL prefix for all API paths, relative to the host root.
    BasePath:   "/api",
    
    // Schemes are the transfer protocols used by the API (http or https).
    Schemes:    []string{"http"},
    
    // APIKey is an optional API key or service account token.
    APIKey:     os.Getenv("WEBITEL_ACCESS_TOKEN"),
    
    // TLSConfig provides an optional configuration for a TLS client
    TLSConfig:  &tls.Config{},
    
    // NumRetries contains the optional number of attempted retries
    NumRetries: 3,
    
    // RetryTimeout sets an optional time to wait before retrying a request
    RetryTimeout: 0,
    
    // RetryStatusCodes contains the optional list of status codes to retry
    // Use "x" as a wildcard for a single digit (default: [429, 5xx])
    RetryStatusCodes: []string{"420", "5xx"},
    
    // HTTPHeaders contains an optional map of HTTP headers to add to each request
    HTTPHeaders: map[string]string{},
}

client := goapi.NewHTTPClientWithConfig(strfmt.Default, cfg)
```

The `goswagger` documentation have more information about how to [build a client](https://goswagger.io/go-swagger/generate/client/).
