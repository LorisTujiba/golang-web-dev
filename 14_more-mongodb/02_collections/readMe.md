collection commands

create implicitly
db.<collection name>.insert({"name":"McLeod"})

create explicitly
db.createCollection(<name>, {<optional options>})


examples
db.createCollection("customers")
db.createCollection("crs",{capped:true, size:65536,max:1000000})

view collections
show collections

drop
db.<collection name>.drop()