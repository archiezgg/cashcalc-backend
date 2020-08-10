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
  "message": "Logged in succesfully",
  "role": "some-role"
}
```

### /refresh
Provides interface to refresh the access token.
* HTTP method: _POST_
* HTTP response: 
	* _200 if successful_ 
	* _422 if data payload is malformed_
	* _401 if the refresh token is not valid_
* Sample required payload:
```
{
	"refreshToken": "some-refresh-token"
}
```
* Sample errror JSON response:
```
{
  "error": "Unauthorized"
}
```
* Sample JSON response after succesful token refreshing:
```
{
  "message": "Token refreshed succesfully"
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
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
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
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
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
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample JSON response:
```
{"vatPercent":27,"airFuelFarePercent":17.5,"roadFuelFarePercent":10,"express9h":8990,"express9hHun":3300,"express12h":2990,"express12hHun":1575,"insuranceLimit":330000,"minInsurance":3300,"ext":1320,"ras":6600,"tk":990}
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
  "tk": 990
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
    "username": "some-user",
    "role": "carrier",
    "tokenString": "some-refresh-token",
    "issuedAt": 1590740566,
    "expiresAt": 1590935428
  }
]
```

### /tokens/revoke
Revokes a single user's refresh token.
* HTTP method: _DELETE_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample required payload:
```
{
	"username": "some-user"
}
```
* Sample JSON response:
```
{
  "message": "Token revoked successfully"
}
```

### /tokens/revokeBulk
Revokes multiple users' refresh tokens.
* HTTP method: _DELETE_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample required payload:
```
{
	"usernames": [
		"some-username",
		"another-username"
	]
}
```
* Sample JSON response:
```
{
  "message": "Multiple tokens revoked successfully"
}
```

### /tokens/revokeAll
Revokes all refresh tokens.
* HTTP method: _DELETE_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample JSON response:
```
{
  "message": "All tokens revoked successfully"
}
```

### /users/usernames
Retrieves all registered usernames, including all roles.
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
Retrieves all registered carrier usernames.
* HTTP method: _GET_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample JSON response:
```
[
  "some-carrier",
  "another-carrier"
]
```

### /users/admins
Retrieves all registered admin usernames.
* HTTP method: _GET_
* HTTP response: 
	* _200 if successful_
	* _401 if no valid token is provided_
	* _403 if token is unathorized for this endpoint_
* Sample JSON response:
```
[
  "some-admin",
  "another-admin"
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
* Sample required payload:
```
{
	"username": "some-user"
}
```
* Sample JSON response:
```
{
  "message": "Carrier deleted successfully"
}
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
* Sample required payload:
```
{
	"username": "some-user"
}
```
* Sample JSON response:
```
{
  "message": "Admin deleted successfully"
}
```
