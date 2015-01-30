#Serving an http folder

Let's start by serving the public folder over http

In order to achieve this, we must use the go net/http package :
- a ServeMux to serve http files using http.FileServer
- a HandleFunc handler redirects requests to that server 

After you have implemented both of these, you can start the server in plain using ListenAndServe 

