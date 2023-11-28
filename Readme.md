### User details consolidation
An api endpoint `/identify` that consolidates user signups. The user is currently tracked using the ip address their request is from.
This tracking can also be done using server set cookies but considering the effort I'm tracking using IP.

The postgres db `user` has the following table `contact`:

| column         | data type                    |
|----------------|------------------------------|
| id             | int                          |
| Email          | string                       |
| phonenumber    | int                          |
| linkedID       | int[]                        |
| linkprecedence | enum('primary', 'secondary') |
| ipv4           | string                       |

LinkedID contains the list of ID of the secondary users if the linkprecedence is primary, 
i.e. contact details of secondary signups.

#### DB setup
[db_init.sh](bin/db_init.sh) contains the script to create db and setup the tables.
This can be improved to goose migrations.

### How it works

When an api call is made to `/identify` the request contains the ip address.
If the email id is not in the db, and there is no `linkprecedance=primary` in the db from that ip
we add that email as `primary` to the db. Subsequent calls from another email/phonenumber combination
from the same ip will create a `linkprecedance=secondary` and update the `linkedID` array in the `primary`.


### Steps to run:
1. Run `./setup.sh`, to bring up the postgres db and initialize the db.
2. This script should also bring up the container that handles the api endpoints

### How to test:
1. First make a call to the api
    ```shell
    curl --location 'localhost:10000/identify' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name":"harshil",
        "email": "h@g.com",
        "phonenumber": 123
    }'
    ```
   Response
   ```json
    {
    "primaryContactID": 1,
    "secondaryContactIDs": null,
    "emails": null,
    "phoneNumbers": null
    }
    ```
2. Another call to the api with different set of email/phonenumber
    ```shell
    curl --location 'localhost:10000/identify' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name":"harshil",
        "email": "h@g1.com",
        "phonenumber": 12321
    }'
    ```
   ```json
   {
    "primaryContactID": 1,
    "secondaryContactIDs": [
        2
    ],
    "emails": [
        "h@g1.com"
    ],
    "phoneNumbers": [
        12321
    ]
   }
   ```
   
Theoritically we should be able to create more docker networks with varying ip. That will give the 
illusion of shopping from different IPs.
