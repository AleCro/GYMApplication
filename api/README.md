# ⚡️ YSvelGoK: API

This is the **Go** ([Gin](https://github.com/gin-gonic/gin)) backend API for the [YSvelGoK](../) project. It serves as the middleware between the front-end and the database, responsible for session integrity, authentication, and logic before data is persisted.

# 🔐 Authentication Middleware
The API provides a tiered middleware system to handle [JWT](https://www.jwt.io/)-based sessions.
## `SoftSession` ➡️
This middleware is intended to run on all handlers by default, and 
- Checks if a JWT token is present in the request.  a token is found, it is parsed and the claims are stored using [`c.Set`](https://pkg.go.dev/github.com/gin-gonic/gin#Context.Set) for optional use in the handler.
- If no token is found, or if any parsing errors occur, the middleware silently ignores the error and proceeds to the original handler.

## `RequireSession` 🔒
This middleware enforces a **valid**, **authenticated session**.

It ensures that a JWT token is **available** and **valid**.

The handler is only reached if the following conditions are met:

- A JWT token is **present**.
- The token signature **can be verified**.
- The token is **not expired**.
- The token claims **can be successfully parsed**.

If successful, the claims are stored using [`c.Set`](https://pkg.go.dev/github.com/gin-gonic/gin#Context.Set).

## `RequireSessionValidate` 🕵️

This is the strictest middleware, validating the token against the database.

It performs all checks from `RequireSession`.

It then queries the database to perform additional validation(s):
- Checks if the underlying **session** (on which the JWT was based) is still **valid in the database**.
- Verifies that the session belongs to the **same user** specified **in the JWT claims**.
- Confirms that the **user account** is still **valid**.

If all conditions are met, both the user and session objects fetched from the database are stored using [`c.Set`](https://pkg.go.dev/github.com/gin-gonic/gin#Context.Set).

# 🧭 Routing Examples

Here is how you can apply the middleware to your Gin router:

```go
// Runs `SoftSession` by default (or no auth)
// 💡 Good for public info
r.GET("/public-info", MyFunction) 

// Requires a valid, unexpired token.
// 🔒 Good for user-specific, non-critical data
r.GET("/my-account", Routes.RequireSession(), MyFunction) 

// Requires a token that is also validated against the live database session.
// ✅ Good for sensitive actions or data
r.GET("/my-sensitive-data", Routes.RequireSessionValidate(), MyFunction)
```

# 💾 Database
Once the API makes a connection to the database, it will create a database, user collection, and session collection (if it doesn't exist yet). It will then verify if a `TTL` index exists, and make sure there is only one which is the same as specified in the envionment variable.

The `Db` module exposes `Connection`, set when `Connect` is executed. Allowing easy access through the entire application. This variable allows you to interact with the application's Data-Access-Layer (DAL). Frequently used and neccesary functions. Refer to [handler.go](./Database/handler.go), and [methods.go](./Database/methods.go) for more details.