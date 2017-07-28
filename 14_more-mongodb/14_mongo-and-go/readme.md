------------------------
setup
------------------------
go get -u gopkg.in/mgo.v2

-----------------------
setup database
-----------------------
restart mongo
close any open mongo connections, then restart with these commands:
mongo

in a new tab
mongod

-----------------------
create db
-----------------------
use company

-----------------------
create collection employees & insert employees
-----------------------
db.employees.insert([{"name":"Mukusito","score":50,"salary":5.99},{"name":"Josh","score":60,"salary":14.99},{"name":"Beryl Markham","score":70,"salary":14.99}])

----------------------
test
----------------------
db.employees.find()

----------------------
user setup
----------------------
db.createUser(
  {
    user: "admin",
    pwd: "admin",
    roles: [ { role: "readWrite", db: "company" } ]
  }
)

exit mongo & then start again with auth enabled
mongod --auth
mongo -u "admin" -p "admin" --authenticationDatabase "company"

------------------------
test
------------------------
use company
db.employees.find()
db.employees.insert([{"name":"Mukusito","score":50,"salary":5.99},{"name":"Josh","score":60,"salary":14.99},{"name":"Beryl Markham","score":70,"salary":14.99}])
db.employees.find()

------------------------
GO & MONGO
------------------------
db access
mongodb://myuser:mypass@localhost:27017/dbToAccess

If the port number is not provided for a server, it defaults to 27017.
for our example:
mongodb://bond:moneypenny007@localhost:27017/bookstore

models
Update to use mongo. You will use the mgo.Dial to create a session. You can still assign this to the variable DB.

run the application and make a request
curl -i localhost:8080/employees
curl -i -X POST -d "name=Metamorphosis&score=50&salary=5.90" localhost:8080/insert