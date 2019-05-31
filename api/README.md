# API

In this folder, different-flavoured interfaces for our backend are defined.

The implementations are split first by flavour (REST, gRPC, etc.), and second by service (customer, car, etc.).

Each service-specific API implementation is a razor-thin wrapper around the actual service implementation, as defined in the respective top-level packages (customer, car, etc.).

## Example Customer

Let's use the below function, found in `customer/register.go`, as our example.

```go
func (cust *Customer) Register(ctx context.Context, req *customerpb.RegisterRequest) (*customerpb.RegisterResponse, error)
```

### REST

Our rest endpoint requires a handler function.  That is, a function of the type `func(http.ResponseWriter, *http.Request)`, for every method.

To implement this register function, we create a function, `func Register (w http.ResponseWriter, req *http.Request)`, to wrap the register function in `customer/register.go`.

This function might look like this:

```go
func (rapi *customerRESTAPI) Register(w http.ResponseWriter, req *http.Request) {
  // set the timeout to whatever your SLA for registration is
  regCtx, cancel := context.WithTimeout(100 * time.Millisecond)
  defer cancel()

  customerReq := *customerpb.RegisterRequest{}
  // ... fill `customerReq` with information from req ...

  resp, err := rapi.customer.Register(regCtx, customerReq)

  // ... use resp and err to write an appropriate response with w.
}
```

### gRPC

It should turn out that your implementation of the service already implements a gRPC service.

We can leverage this fact to make the implementation of a gRPC-based API extremely thin:

```go
func (gapi *ElasticAuthGRPCAPI) Register(ctx context.Context, req *customerpb.RegisterRequest) (*customerpb.RegisterResponse, error) {
  return gapi.customer.Register(ctx, customerReq)
}
```
