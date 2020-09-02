## API documentation
This section serves as an API documentation for the frontend side to be able to query data succesfully.

### /login
Provides interface for login, returns with a JWT access token and a refresh token in cookies if the login was succesful.
* HTTP method: _POST_
* HTTP response: 
	* _200 if successful_ 
	* _422 if data payload is malformed_
	* _401 if the password is incorrect_
* Payload required: _username and password_
* Sample required payload:
```
{
	"username": "some-user",
	"password": "some-pw"
}
```
* Sample errror JSON response:
```
{
  "error": "Unauthorized"
}
```
* Sample JSON response after succesful login:
```
{
  "message": "Logged in successfully",
  "username": "some-user",
  "role": "some-role"
}
```

### /logout
Provides interface for logout, deletes the access and refresh tokens from cookies.
* HTTP method: _POST_
* HTTP response: 
	* _200 if successful_ 
	* _500 if no cookie provided_
* Sample JSON response after succesful login:
```
{
  "message": "Logged out successfully",
}
```

### /is-authorized
Provides interface to check if the user in the token is authorized for given user's contents.
* HTTP method: _GET_
* HTTP response: 
	* _200 if the user is authorized_
	* _401 if user is not authenticated_
	* _403 if the token is not authorized for the given user's contents_
	* _500 if the given role is not valid_
* Queries:
  * role (mandatory): _the role that should be checked for, ("carrier" | "admin" | "superuser")_
* Sample JSON response:
```
{
  "message": "Authorized"
}
```

### /calc
Calculates the resulting fares based on the input.
* HTTP method: _POST_
* HTTP response: 
	* _200 if successful_ 
	* _422 if data payload is malformed_
	* _403 if token is unauthorized for this endpoint_
* Payload required:
	* _transferType: string ("air" | "road"), mandatory_
	* _zoneNumber: integer between 0-9, mandatory_
	* _weight: float between 0.5-200, mandatory_
	* _insurance: integer, optional, defaults to 0_
	* _discountPercent: float, optional, defaults to 0_
	* _expressType: string ("worldwide" | "9h" | "12h"), mandatory_
	* _isDocument: boolean, defaults to false_
	* _isExt: boolean, defaults to false_
	* _isTk: boolean, defaults to false_
	* _isRas: boolean, defaults to false_

* Sample required payload:
```
{
	"transferType": "air",
	"zoneNumber": 5,
	"weight": 0.5,
	"insurance": 1000,
	"discountPercent": 10,
	"expressType": "worldwide",
	"isDocument": true,
	"isExt": false,
	"isTk": true,
	"isRas": false
}
```
* Sample JSON response:
```
{
	"baseFare": 14696,
	"expressFare": 0,
	"insuranceFare": 3300,
	"extFare": 0,
	"rasFare": 0,
	"tkFare": 990,
	"fuelFare": 2572,
	"emergencyFare": 65,
	"result": 21623
}
```

### /countries
Retrieves both air and road countries.
* HTTP method: _GET_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
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
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
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
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
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
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
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
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
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
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
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
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Zone number: an integer between 1-5
* Queries:
  * weight (optional): _retrieves only the fare of for the weight provided_
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
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Zone number: an integer between 0-9
* Queries:
  * weight (optional): _retrieves only the fare of for the weight provided_
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
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Zone number: an integer between 5-9
* Queries:
  * weight (optional): _retrieves only the fare of for the weight provided_
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
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample JSON response:
```
{"vatPercent":27,"airFuelFarePercent":17.5,"roadFuelFarePercent":10,"express9h":8990,"express9hHun":3300,"express12h":2990,"express12hHun":1575,"insuranceLimit":330000,"minInsurance":3300,"ext":1320,"ras":6600,"tk":990,"emergencyFare":65}
```

### /pricingvariables/update
Updates the pricing variables with the given variables.
* HTTP method: _PATCH_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Payload required:
	* _vatPercent: float, mandatory_
	* _airFuelFarePercent: float, mandatory_
	* _roadFuelFarePercent:float, mandatory_
	* _express9h: integer, mandatory_
	* _express9hHun: integer, mandatory_
	* _express12h: integer, mandatory_
	* _express12hHun: integer, mandatory_
	* _insuranceLimit: integer, mandatory_
	* _minInsurance: integer, mandatory_
	* _ext: integer, mandatory_
	* _ras: integer, mandatory_
	* _tk: integer, mandatory_
	* _emergencyFare: integer, mandatory_

* Sample required payload:
```
{
  "vatPercent": 27,
  "airFuelFarePercent": 17.5,
  "roadFuelFarePercent": 10,
  "express9h": 8990,
  "express9hHun": 3300,
  "express12h": 2990,
  "express12hHun": 1575,
  "insuranceLimit": 330000,
  "minInsurance": 3300,
  "ext": 1320,
  "ras": 6600,
  "tk": 990,
  "emergencyFare": 65
}
```
* Sample JSON response:
```
{
  "message": "Pricing variables updated successfully"
}
```

### /tokens
Retrieves the refresh tokens stored in database.
* HTTP method: _GET_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample JSON response:
```
[
  {
    "tokenString": "some-token-string",
    "UserID": 1,
    "CreatedAt": "2020-09-02T21:04:42.770321Z",
    "expiresAt": "2020-09-09T21:04:42Z"
  }
]
```

### /tokens/loggedin
Retrieves the logged in users.
* HTTP method: _GET_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample JSON response:
```
[
  {
    "id": 1,
    "username": "some-user",
    "role": "some-role",
    "createdAt": "2020-09-02T21:00:06.820115Z",
    "updatedAt": "2020-09-02T21:04:42.770679Z",
    "deletedAt": "0001-01-01T00:00:00Z"
  }
]
```

### /tokens/revoke
Revokes a single user's all refresh tokens, esentially logging it out.
* HTTP method: _DELETE_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Queries:
  * userid (mandatory): _deletes the refresh tokens belonging to the user_

* Sample JSON response:
```
{
  "message": "Token revoked successfully"
}
```

### /users/usernames
Retrieves all registered usernames regardless of their roles.
* HTTP method: _GET_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample JSON response:
```
[
  "some-user",
  "another-user"
]
```

### /users/carriers
Retrieves all registered carrier users.
* HTTP method: _GET_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample JSON response:
```
[
  {
    "id": 2,
    "username": "carrier-test",
    "role": "carrier",
    "createdAt": "2020-08-30T21:07:48.541908Z",
    "updatedAt": "2020-08-30T21:07:48.541908Z",
    "deletedAt": "0001-01-01T00:00:00Z"
  },
  {
    "id": 3,
    "username": "carrier-test1",
    "role": "carrier",
    "createdAt": "2020-08-30T21:07:50.705943Z",
    "updatedAt": "2020-08-30T21:07:50.705943Z",
    "deletedAt": "0001-01-01T00:00:00Z"
  }
]
```

### /users/carriers/create
Creates a user with the role carrier.
* HTTP method: _PUT_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample required payload:
```
{
	"username": "some-user",
	"password": "some-password"
}
```
* Sample JSON response:
```
{
  "message": "Carrier created successfully"
}
```

### /users/carriers/delete
Deletes a user with the role carrier.
* HTTP method: _DELETE_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Queries:
  * id (mandatory): _deletes the given carrier by ID_

* Sample JSON response:
```
{
  "message": "Carrier deleted successfully"
}
```

### /users/admins
Retrieves all registered admin users.
* HTTP method: _GET_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample JSON response:
```
[
  {
    "id": 2,
    "username": "admin-test",
    "role": "admin",
    "createdAt": "2020-08-30T21:07:48.541908Z",
    "updatedAt": "2020-08-30T21:07:48.541908Z",
    "deletedAt": "0001-01-01T00:00:00Z"
  },
  {
    "id": 3,
    "username": "admin-test1",
    "role": "admin",
    "createdAt": "2020-08-30T21:07:50.705943Z",
    "updatedAt": "2020-08-30T21:07:50.705943Z",
    "deletedAt": "0001-01-01T00:00:00Z"
  }
]
```

### /users/admins/create
Creates a user with the role admin.
* HTTP method: _PUT_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample required payload:
```
{
	"username": "some-user",
	"password": "some-password"
}
```
* Sample JSON response:
```
{
  "message": "Admin created successfully"
}
```

### /users/admins/delete
Deletes a user with the role admin.
* HTTP method: _DELETE_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Queries:
  * id (mandatory): _deletes the given admin by ID_

* Sample JSON response:
```
{
  "message": "Admin deleted successfully"
}
```

### /users/superusers
Retrieves all registered superuser users.
* HTTP method: _GET_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample JSON response:
```
[
  {
    "id": 2,
    "username": "sudo-test",
    "role": "superuser",
    "createdAt": "2020-08-30T21:07:48.541908Z",
    "updatedAt": "2020-08-30T21:07:48.541908Z",
    "deletedAt": "0001-01-01T00:00:00Z"
  },
  {
    "id": 3,
    "username": "sudo-test1",
    "role": "superuser",
    "createdAt": "2020-08-30T21:07:50.705943Z",
    "updatedAt": "2020-08-30T21:07:50.705943Z",
    "deletedAt": "0001-01-01T00:00:00Z"
  }
]
```

### /users/admins/create
Creates a user with the role superuser.
* HTTP method: _PUT_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample required payload:
```
{
	"username": "some-user",
	"password": "some-password"
}
```
* Sample JSON response:
```
{
  "message": "Superuser created successfully"
}
```