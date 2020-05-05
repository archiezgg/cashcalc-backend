## API documentation
This section serves as an API documentation for the frontend side to be able to query data succesfully.

### /countries
Retrieves both air and road countries.
* HTTP method: _GET_
* HTTP response: 
	* _200 if successful_ 
	* _500 if the server cannot process the data properly_
* Sample JSON response:
```
{
	"countriesAir": [{
		"name": "Afganisztán",
		"zoneNumber": 9
	}],
	"countriesRoad": [{
		"name": "Andorra",
		"zoneNumber": 5
	}]
}
```

### /countries/air
Retrieves only air countries.
* HTTP method: _GET_
* HTTP response: 
	* _200 if successful_ 
	* _500 if the server cannot process the data properly_
* Sample JSON response:
```
[{
	"name": "Afganisztán",
	"zoneNumber": 9
}, {
	"name": "Albánia",
	"zoneNumber": 5
}]
```

### /countries/road
Retrieves only road countries.
* HTTP method: _GET_
* HTTP response: 
	* _200 if successful_ 
	* _500 if the server cannot process the data properly_
* Sample JSON response:
```
[{
	"name": "Andorra",
	"zoneNumber": 5
}, {
	"name": "Ausztria",
	"zoneNumber": 1
}]
```

### /pricings
Retrieves both air and road pricings with their zone numbers, weight and baseFare pairings.
* HTTP method: _GET_
* HTTP response: 
  * _200 if successful_
  * _500 if the server cannot process the data properly_
* Sample JSON response:
```
{
	"airPricings": [{
		"zoneNumber": 0,
		"fares": [{
			"weight": 0.5,
			"baseFare": 2950
		}, {
			"weight": 1,
			"baseFare": 3005
		}]
	}],
	"roadPricings": [{
		"zoneNumber": 1,
		"fares": [{
			"weight": 1,
			"baseFare": 18652
		}, {
			"weight": 2,
			"baseFare": 22106
		}]
	}]
}
```

### /pricings/road
Retrieves only the road pricings.
* HTTP method: _GET_
* HTTP response: 
  * _200 if successful_
  * _500 if the server cannot process the data properly_
* Sample JSON response:
```
[{
		"zoneNumber": 1,
		"fares": [{
			"weight": 1,
			"baseFare": 18652
		}]
	},
	{
		"zoneNumber": 5,
		"fares": [{
			"weight": 1,
			"baseFare": 26456
		}]
	}
]
```

### /pricings/air
Retrieves only the air pricings.
* HTTP method: _GET_
* HTTP response: 
  * _200 if successful_
  * _500 if the server cannot process the data properly_
* Sample JSON response:
```
[{
		"zoneNumber": 1,
		"fares": [{
			"weight": 1,
			"baseFare": 18652
		}]
	},
	{
		"zoneNumber": 5,
		"fares": [{
			"weight": 1,
			"baseFare": 26456
		}]
	}
]
```

### /pricings/road/fares/{zoneNumber}
Retrieves the road fares of the zone provided.
* HTTP method: _GET_
* HTTP response: 
  * _200 if successful_
  * _500 if the server cannot process the data properly_
* Zone number: an integer between 1-5
* Queries:
  * weight(optional): _retrieves only the fare of for the weight provided_
* Example: _/pricings/road/fares/4_
* Sample JSON response:
```
[{
	"weight": 1,
	"baseFare": 20674
}, {
	"weight": 2,
	"baseFare": 25454
}]
```
* Example _/pricings/road/fares/4?weight=1_
* Sample JSON response:
```
{"weight":1,"baseFare":20674}
```

### /pricings/air/fares/{zoneNumber}
Retrieves the air fares of the zone provided.
* HTTP method: _GET_
* HTTP response: 
  * _200 if successful_
  * _500 if the server cannot process the data properly_
* Zone number: an integer between 0-9
* Queries:
  * weight(optional): _retrieves only the fare of for the weight provided_
* Example: _/pricings/air/fares/4_
* Sample JSON response:
```
[{
	"weight": 0.5,
	"baseFare": 16224
}, {
	"weight": 1,
	"baseFare": 20534
}]
```
* Example _/pricings/air/fares/4?weight=1_
* Sample JSON response:
```
{"weight":1,"baseFare":20534}
```

### /pricings/air/docfares/{zoneNumber}
Retrieves the air document fares of the zone provided.
* HTTP method: _GET_
* HTTP response: 
  * _200 if successful_
  * _500 if the server cannot process the data properly_
* Zone number: an integer between 5-9
* Queries:
  * weight(optional): _retrieves only the fare of for the weight provided_
* Example: _/pricings/air/docfares/5_
* Sample JSON response:
```
[{"weight":0.5,"baseFare":16329},{"weight":1,"baseFare":20786},{"weight":1.5,"baseFare":24735},{"weight":2,"baseFare":28684}]
```
* Example _/pricings/air/docfares/5?weight=1.5_
* Sample JSON response:
```
{"weight":1.5,"baseFare":24735}
```

### /pricingvariables
Retrieves the pricing variables.
* HTTP method: _GET_
* HTTP response: 
  * _200 if successful_
  * _500 if the server cannot process the data properly_

* Sample JSON response:
```
{"vatPercent":27,"airFuelFarePercent":17.5,"roadFuelFarePercent":10,"express9h":8990,"express9hHun":3300,"express12h":2990,"express12hHun":1575,"insuranceLimit":330000,"minInsurance":3300,"ext":1320,"ras":6600,"tk":990}
```
