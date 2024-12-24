# Package repo

This package provides base services that assume an `Endpoint` object, which includes access to an `sql.DB`, is stored in the Golang context. Below are strategies for providing the `Endpoint` object in the context, depending on your use case.

## Strategies for Providing the `Endpoint` Object

### 1. When the `sql.DB` Object is Not Shared with Other Golang Libraries

In this scenario, you can create and manage a private `sql.DB` connection for use exclusively within this package.

```go
ctx := context.Background()
config := repo.NewConfig()
factory := repo.NewEndpointFactory(config, loggerProvider)
endpoint := factory.CreateEndpoint(ctx)
ctx = repo.NewContextWithEndpoint(ctx, endpoint)
```

#### Explanation:
1. **Line 1**: Creates a context. If you already have a context, you can reuse it.
2. **Line 2**: Initializes a `Config` object by retrieving the connection string from the `SQLITE3_DB` environment variable. The `Config` object can be reused as a singleton.
3. **Line 3**: Creates an `EndpointFactory` object, which can also be reused as a singleton. For the `loggerProvider` argument, refer to [logger_provider.go](./logger_provider.go).
4. **Line 4**: Creates an `Endpoint` object. This object is scoped per request and cannot be a singleton.
5. **Line 5**: Stores the `Endpoint` in the context, returning a new context that should be used for subsequent operations.

### 2. When the `sql.DB` Object is Shared with Other Golang Libraries

In this scenario, you share an existing `sql.DB` connection with other libraries.

```go
var sqlDB *sql.DB
ctx := context.Background()
ctx = repo.NewContextWithSQLDB(ctx, sqlDB)
```

#### Explanation:
1. **Line 1**: Assumes you have already created and provided the `sql.DB` pointer.
2. **Line 2**: Creates a context. If you already have a context, you can reuse it.
3. **Line 3**: Stores the `sql.DB` in the context by creating and embedding the `Endpoint`. A new context is returned and should be used for subsequent operations.

---

### Best Practices
- Ensure that the `sql.DB` object is properly closed when it is no longer needed to avoid connection leaks.
- Use a singleton `Config` object and `EndpointFactory` to optimize resource usage.
- Always use the context returned by `repo.NewContextWithEndpoint` or `repo.NewContextWithSQLDB` for subsequent operations to ensure the correct `Endpoint` is accessible.

For more details, refer to the [logger_provider.go](./logger_provider.go) file and additional documentation in this repository.

