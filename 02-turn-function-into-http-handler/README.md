# Turn function into HTTP handler

The puzzle already contains the code to run an http server listening on port 9090. You are provided with a struct representing the root handler, which should be passed as argument to the `mux.Handle()` method. You should complete the code such that the provided go sample runs correctly, printing "handler hit" whenever a request is sent to the server.

<details>
<summary><b>Hint</b></summary>

go doc net/http.Handler

</details>