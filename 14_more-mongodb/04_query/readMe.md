find (aka, query)
setup
use store
db
show dbs
db.customers.insert([{"role":"double-zero","name": "Bond","age": 32},{"role":"citizen","name": "Moneypenny","age":32},{"role":"citizen","name": "Q","age":67},{"role":"citizen","name": "M","age":57},{"role":"citizen","name": "Dr. No","age":52}])

find
db.<collection name>.find()
db.customers.find()

find one
db.<collection name>.findOne()
db.customers.findOne()

find specific
db.customers.find({"name":"Bond"})
db.customers.find({name:"Bond"})
You can do it either way: "name" or name. JSON specification is to enclose name (object name-value pair) in double qoutes

and
db.customers.find({$and: [{name:"Bond"}, {age:32}]})
db.customers.find({$and: [{name:"Bond"}, {age:{$lt:20}}]})
db.customers.find({$and: [{name:"Bond"}, {age:{$gt:20}}]})

or
db.customers.find({$or: [{name:"Bond"}, {age:67}]})
db.customers.find({$or: [{name:"Bond"}, {age:{$lt:20}}]})
db.customers.find({$or: [{name:"Bond"}, {age:{$gt:32}}]})

and or
db.customers.find({$or:[
{ $and : [ { role : "citizen" }, { age : 32 } ] },
{ $and : [ { role : "citizen" }, { age : 67 } ] }
]})

regex
db.customers.find({name: {$regex: '^M'}})