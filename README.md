## Ads store

*Execute:*
- `docker-compose up` to run postgres db  
  Note: Database storage will create outside the docker container in the `pgdata` directory
- `make` to run main server

---

### Creating the ad

*Do request in another terminal*  
`curl -X POST http://localhost:9000/createAd -H "Content-Type: application/json" --data '{"title": "Title of Ad", "photo": "some/path/img.jpg", "price": 11.22}'`

You can to post your credentials in this request.  

