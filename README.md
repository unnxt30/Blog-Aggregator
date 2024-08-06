# Blog-Aggregator

A RSS-Feed Aggregator implemented in GO, which lets you get the data of your favorite RSS feed with a few endpoint requests. It comes with the feature of creating users and authorizing them with ApiKeys. Go-routines and waitgroups utilized to fetch multiple feeds together at once.


## Quick Start 🚀
- Clone the repository
- Install [PostgreSQL](https://www.postgresql.org/download/)

### Install Goose
- A tool to manage your database schema for GO.
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest 
```

### Install sqlc
- A tool to convert sql queries into GO functions to easily access the database.
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### Start the server
```bash
go run .
```



## Endpoints

### POST /v1/users
Endpoint to create user in the database.\
Body:
```json
{"name": "your_name"}
```
### GET /v1/users
This is an authorized endpoint, which uses an API key generated by the POST Users endpoint to give back the details of that specific user.\
Header:
```
Authorization : Bearer {ApiKey}
```

### POST /v1/feed
This endpoint lets you enter your desired feed that you will get the data from.This too is an authorized endpoint which will map to the user who is the owner of the ApiKey.\
Body:
```json
{"name": "name of your blog",
 "url": "https://abc.com/index.xml"}
 ```
Header:
```
Authorization: Bearer {ApiKey}
```
 ### GET /v1/feed
 Authorized endpoint which gets the all the feeds of a the user verified by ApiKey.\
 Header:
```
Authorization : Bearer {ApiKey}
```

### POST /v1/feed_follows
This is an authorized endpoint lets you follow feeds created by other users.\
Body:
```json
{"feed_id" : "{id of the feed you have to follow}"}
```
Header:
```
Authorization: Bearer {ApiKey}
```

### GET /v1/feed_follows
Get all the feed a user follows from the database.
Header:
```
Authorization: Bearer {ApiKey}
```

### DELETE /v1/feed_follows/{feedFollowID}
Delete the following of a feed with the given id.

### GET /v1/posts
Authorized endpoint, which shows the posts of the feed the user is following or has posted.
Header:
```
Authorization: Bearer {ApiKey}
```