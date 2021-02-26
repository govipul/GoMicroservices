# Problem faced in episode 4
No big problem faced during running the program, but i am unable to decode the json value to product object. This is the curl command i passed 
```curl localhost:9000/1 -XPUT -d '{"name": "tea", "description": "Relaxing tea"}' | jq```
The output message i got is as follow:
```product-api:2021/02/26 17:32:34 Prod: &data.Product{ID:0, Name:"", Description:"", Price:0, SKU:"", CreatedOn:"", UpdatedOn:"", DeleteOn:""}```

So no idea what has went incorrect during the tutorial, but will continuing to debug the code and update the solution as soon i will be able to fix the issue.
