FORMAT: 1A


# POSTS CREATION API [/internal/posts]

## Create posts [POST]
The purpose of this api is to create posts given a topic and a body


### Create new posts in the system

+ Request
     + Headers

         Content-Type: application/json


     + Body (application/json)

         ```
         {
            "topic":"Golang Boilerplate",
            "body":"This is about golang boilerplate",
         }
         ```

+ Response 201
     + Headers

            Content-Type: application/json

         + Body

           ```
           {
              "post":{
                  "id": "12345"
                  "topic":"Golang Boilerplate",
                  "body":"This is about golang boilerplate",
                         
               }
              "errors":null
           }
           ```

+ Request
     + Headers

         Content-Type: application/json

     + Body (application/json)
        
        ```
         {
             topic:"Golang Boilerplate",
             body:"This is about golang boilerplate",
         }
         ```

+ Response 400(Malformed JSON)

     + Body

        ```js
        {
           "posts":null,
           "errors":[
              {
                 "code":"900",
                 "entity":"Malformed JSON"
              }
           ]
        }
         ```
+ Request
     + Headers

         Content-Type: application/json

     + Body (application/json)
        
        ```
         {
             "topic":"Golang Boilerplate",
             "body":"This is about golang boilerplate",
         }
         ```

+ Response 500(Internal Server Errors)
     + Headers

        Content-Type: application/json

     + Body

        ```js
        {
           "posts":null,
           "errors":[
              {
                 "code":"900",
                 "entity":"Failed to create post"
              }
           ]
        }
     ```