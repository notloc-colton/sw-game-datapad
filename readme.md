Description:
  Simulates a "datapad" for all your geeky star wars needs.

How to run locally:
  -go get -d ./...
  -go run cmd/main.go
  
  Note: server will be running on 127.0.0.1:8080

API Documentation:
  The api is relatively simple, there is one endpoint with one query string parameter ("query") that accepts any number of characters to query the application for
  matching star wars characters. The response will be an array of star wars characters (as seen in the below example).
  Example
    Request: curl --location 'localhost:8080/characters?query=chew'
    Response: [
    {
        "name": "Chewbacca",
        "homePlanet": {
            "name": "Kashyyyk",
            "climate": "tropical",
            "population": "45000000"
        },
        "species": {
            "name": "Wookie",
            "averageLifespan": "400",
            "language": "Shyriiwook"
        },
        "starships": [
            {
                "name": "Millennium Falcon",
                "cargoCapacity": "100000",
                "class": "Light freighter"
            },
            {
                "name": "Imperial shuttle",
                "cargoCapacity": "80000",
                "class": "Armed government transport"
            }
        ]
    }
]



Considerations:
  I felt like the biggest issue that I saw with creating this application was the speed. The requirements specified that the app basically needed to consolidate 4 different swapi endpoints
  into 1. So, as opposed to trying to be super creative with what the endpoint request and response bodies looked like (I feel like simplicity is beauty), I focused on speed. Now, I may have completely 
  missed the mark with how you guys may have wanted this app built (maybe having a super creative request and response body was key), but I personally felt like speed was the most pressing issue in this   app... Additionally, there were many questions that I could have asked with regards to data integrity (ex. "Is the character species `Name` field required to be non empty string?") but I took some 
  liberty and answered most of those questions on my own- I hope that was alright! 


