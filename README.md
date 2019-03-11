Volunteers api app

App is configured using config file. In conjunction with microservices it will be probably
configured in different way (docker allows multiple ways - through env variables for example).

Volunteers api 'description':
Volunteers can 
- create account
- modify account (except password)
- view his/her account.

Authorization/access control:
- implemented simple inmemory session management due to time constraints
- user needs to login first to access restricted endpoints
- in case of not logged in, api returns 400 with error message
- passwords are not encrypted, just saved in DB as plain text. 

Note to volunteers endpoints architecture
There is slight violation, endpoints are not designed as in production they should be this is because
of simplification and ease of use (Postman). 
for example ``GET@api/v1/volunteers`` returns not all accounts but only currently logged in volunteer what should be 
served with ``GET@api/v1/volunteers/{id:int}``


Teams api 'description'
volunteer can
- create team
- view all available teams
- delete team which he created (even when have members/occupants)
- join team
- leave team
- view hom many members (occupants) each team have

I have replaced id:int for namne:string in teams endpoints to simplify use with Postman


Used technologies:
- as DB i used postgres, i od believe exercise also recommends use of it
- ofcourse Go
- as REST framework i have chosen ``iris``, i was deciding between gin and iris, i think there is no 
big difference so i have chosen based on fact that i like iris style more, also it turns out it's prty
easy to use and intuitive, but i dont needed any advanced features tho

Things that are not covered in full or at all:
- returned errors are not fine grained, this requires much more complex approach
- some endpoints 'swallows' some types of errors, this are not critical or dont provide useful feedback (in this app)
- endpoint ``GET@/api/v1/teams/occupants`` should be implemented in different way (for example as secondary api accessible 
without authorization, accessible only internally between microservices) 
- cron logger is very simple, probably too simple...
- session management is very very simplified for time reasons

Resources:
- /postman contains postman collection used to test api (its needed to set URI variable)
- settings.json in root dir contains example config consumed by app
- /sql_scripts contains sql scripts used to create db with tables

How to run:``volunteers-api.exe -s settings.json`` (on windows)

 


   