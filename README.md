## Ads store

*Execute:*
- `docker-compose up` to run postgres db  
  Note: Database storage will create outside the docker container in the `pgdata` directory
- `make` to run main server

---

### Creating the ad

*Do request in another terminal*  
`curl -X POST http://localhost:9000/createAd -H 
"Content-Type: application/json" --data '{
"title": "Title of Ad", 
"content": "Some content", 
"photo": ["some/path/img.jpg", "./another/one.jpg", "./valenkiOptom.jpg"], 
"price": 11.22}'`

You can to post your credentials in this request.  

### Get Ad by ID

*Do request*  
`curl -X POST http://localhost:9000/getAd -H "Content-Type: application/json" --data '{"id": N}'`

Where `N` is a number of Ad ID to get.  

### Get all Ads

*Do request*  
`curl -X POST http://localhost:9000/getAllAds -H 
"Content-Type: application/json" --data 
'{"orderby": "price", 
"direction": "DESC", 
"limit": 10, 
"offset": 10}'`

Note: Then just created my db contains the two test records.  
The `direction` statement can assume "ASC" for ascending of "DESC" for descending only.  
The `limit` statement required to wrapped into double quotation. 
