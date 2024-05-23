Features
--------

-   Submit one or more orders in a single request
-   Retrieve a list of items purchased by an individual customer
-   Generate summaries including all customers with the number of purchased items and total amount spent

Usage
-----

### Installation

1.  Clone the repository:

bash

Copy code

`git clone https://github.com/kowalskyyy/golang_rest_api.git`

Install gin package

`go get -u github.com/gin-gonic/gin`


### Running the Server

1.  Run the project:

`go run .`


### Making Requests

-   Submit Orders:
    -   Assumption: This endpoint on default will accept only valid orders and reject those with invalid data structure. Response will include valid and invalid orders in separate objects. For more strict validation, add query parameter `?strict=true` to the endpoint - it will reject the entire request if at least one order is incorrect.
    -   Endpoint: `/submit-orders`
    -   Method: `POST`
    -   Payload: JSON array of orders
    -   Example:

        `[
            {
                "customerId": "01",
                "orderId": "50",
                "timestamp": "1637245070513",
                "items": [
                    {
                        "itemId": "20201",
                        "costEur": 2
                    }
                ]
            }
        ]`
-   Get All Orders:
    -   Endpoint: `/get-orders`
    -   Method: `GET`
-   Get Customer Items:
    -   Endpoint: `/get-items/{customerId}`
    -   Method: `GET`
-   Get Customer Summaries:
    -   Endpoint: `/get-summary`
    -   Method: `GET`
