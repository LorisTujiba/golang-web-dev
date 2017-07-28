sort
db.<collection name>.find().sort(<field to sort on>:<1 for ascend, -1 descend>)

db.oscars.find({},{_id:0,year:1,title:1}).limit(10).sort({title:1})
db.oscars.find({},{_id:0,year:1,title:1}).limit(10).sort({title:-1})
