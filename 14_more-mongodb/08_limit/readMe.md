limit
db.<collection name>.find(<selection criteria>).limit(n)

db.crayons.find().limit(3)

db.customers.find({age:{$gt:32}},{_id:0,name:1,age:1}).limit(2)