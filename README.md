# amazon-scraper-sellerapp
This is a scraper built as a part of the Seller App assignment.

# Task

Create a POST (REST) API in Golang
- Which scrapes an Amazon web page given its URL
- The code should scrape following data from the page
    - Product Name/Title
    - Product image url
    - Product description
    - Product price
    - Product total number of reviews.
- Use Colly library for golang to simplify scraping.
- Calls the below SEPARATE API to persist the information 

Create another POST (REST) API in Golang and in a SEPARATE service.
- API handler takes the JSON structure from the request payload
- Handler writes the obtained payload to a new/updated Document (Use your favourite Document store)
- Also write timestamp of the create/update within the document
- Refer to the sample JSON below, but add/update fields as you see fit.

Containerize,
- Dockerize each service (both APIs are deployed as separate service) into a separate image,
- Use Docker Compose to include 2 services and any DB that is required
- Ensure that ‘docker compose up’ can bring up all the services together.

Push to Github
- Create a project/repo on github
- Add a simple README demonstrating how to build and run this code
- Ensure the access to the repo is provided 

# Solution

There are two Services built as a part of the assignment ScraperAPI and StoreAPI.

The Scraper API, uses colly to scrape the Product name, description, imageURL, price and totalReviews from the passed URL.

POST /scrape

```
{
    "URL" : "https://www.amazon.com/Corsair-Premium-Gaming-Headset-Surround/dp/B07X8SFYJV"
}
```

On Success, returns a 201.

The Store API, takes a document to write to into the MongoDB store.

POST /product

```
{
	"url": "https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/",
	"product": {
		"name": "PlayStation 4 Pro 1TB Console",
		"imageURL": "https://images-na.ssl-images-amazon.com/images/I/41GGPRqTZtL._AC_.jpg",
		"description": "Heighten your experiences.\n Enrich your adventures.",
		"price": "$348.00",
		"totalReviews": 4436
	}
}
```

On Success, returns a 201.

# Running the Scraper System

The system contains a docker-compose.yml with all the config required for running the services and mongoDB.

Running 
    ```docker compose up``` should have everything up and running.

# Contact

Please reach out to adithdevreddy@gmail.com for any queries regaring the same.