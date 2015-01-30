#Serving websocket requests

We will now create a basic message handling for websockets using the gorilla web tooltip, which is a wrapper around the standard net packages.

The you can find the documentation at [http://www.gorillatoolkit.org/](http://www.gorillatoolkit.org/).

Complete the following steps :

- Import the websocket package
- Create a struct to hold the websocket serving logic
- Implement the net/http.Handler interface on that struct

Your serve http method must do the following :
- Upgrade the http connection to a websocket connection using the upgrader type. You can return true for all requests in your CheckOrigin function
- Cleanly close it in case you exit the method
- Use a generic httpError handling function for the various websocket error you can encounter 
- Loop as long as  you can read messages from the connection, return otherwise. Within the loop :
    - Read text messages as json (other messages are simply pings form the browser to keep the connection alive)
    - Write back an ack in any form (preferably another structured json message)

