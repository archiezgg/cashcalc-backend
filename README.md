# cashcalc-backend 
[![CircleCI](https://circleci.com/gh/IstvanN/cashcalc-backend.svg?style=svg)](https://circleci.com/gh/IstvanN/cashcalc-backend) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=IstvanN_cashcalc-backend&metric=alert_status)](https://sonarcloud.io/dashboard?id=IstvanN_cashcalc-backend)

Backend for the CashCalc 2020 application.

This project aims to give additional support for people working in logistics to calculate transportation costs.

It is a collaboration between me and [@mark182182](https://github.com/mark182182). The frontend can be found [here](https://github.com/mark182182/cashcalc-frontend).


## API documentation
This section serves as an API documentation for the frontend side to be able to query data succesfully.

### /countries
Retrieves all countries with their zone numbers based on the type (road or air).
* HTTP method: _GET_
* HTTP response: _200 if successful, 400 if the request is badly formed_
* Queries:
   * "type" (mandatory): _road | air_
* Example: _/countries?type=air_
* Sample JSON response:
```
[{"Name":"Andorra","ZoneNumber":5},{"Name":"Ausztria","ZoneNumber":1}]
```

### /pricings
Retrieves all the pricings with their zone numbers, weight and basefare pairings.
* HTTP method: _GET_
* HTTP response: _200 if successful, 400 if the request is badly formed_
* Queries:
  * "type" (mandatory): _road | air_
* Example: _/pricings?type=road_
* Sample JSON response:
```
[{"ZoneNumber":0,"Fares":[{"Weight":0.5,"BaseFare":2950},{"Weight":1,"BaseFare":3005}]
```

### /pricings/fares
Retrieves the weight-basefare pairings based on the queries.
* HTTP method: _GET_
* HTTP response: _200 if successful, 400 if the request is badly formed_
* Queries:
  * "type" (mandatory): _road | air_
  * "zn" (mandatory): _the zone number of the pricing (0-9)_
* Example: _/pricings/fares?type=road&zn=1_
* Sample JSON response:
```
[{"Weight":1,"BaseFare":18652},{"Weight":2,"BaseFare":22106}]
```

### /pricings/docfares
Retrieves the document fares based on the queries.

__NOTE that only air pricings have document fares, and only for the pricings with zone number of 5-9.__

* HTTP method: _GET_
* HTTP response: _200 if successful, 400 if the request is badly formed_
* Queries:
  * "zn" (mandatory): _the zone number of the pricing (5-9)_
* Example: _/pricings/docfares?zn=7_
* Sample JSON response:
```
[{"Weight":0.5,"BaseFare":16329},{"Weight":1,"BaseFare":20786},{"Weight":1.5,"BaseFare":24735},{"Weight":2,"BaseFare":28684}]
```


