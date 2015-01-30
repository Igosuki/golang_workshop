#Doing pub sub with Redis

We will now use the go wrapper for the redis c client at github.com/gosexy/redis

Here are the new expected features

- Register Web socket connections in a boolean map
- Create a non blocking connection to a redis host using redis.Client
- Create a sub channel to receive subs from redis and send them back to connections which are alive
- Create a pub channel to publish websocket messages to redis
