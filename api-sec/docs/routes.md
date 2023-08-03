## Routes

### **GET `/home`**
Interface for testing login with Google.

### **GET `/api/`**
Health check.
-   **Auth Required:** No
-   **Allowed Roles:** All
-   **Response:** `200`

### **GET `/api/google/callback`**
Callback endpoint for Google login, used during the OAuth2 authentication flow with Google.

### **POST `/api/login`**
Login with username and password.
-   **Request Body:**

```json
{
	"username": "mikaella",
	"password": "mikaellapw"
}
``` 
-   **Response:** `200`

```json
{
	"token": {
		"Username": "mikaella",
		"ExpireAt": "2023-07-26T21:30:36-03:00",
		"Attributes": {}
	}
}
```

### **POST `/api/logout`**
Logout user.

-   **Auth Required:** Yes
-   **Allowed Roles:** All
-   **Response:** `204`

### **POST `/api/token/refresh`**
Refresh token when expired. Refresh token must be in cookies.

-   **Auth Required:** Yes
-   **Allowed Roles:** All
-   **Response:** `200`

```json
{
	"token": {
		"Username": "mikaella",
		"ExpireAt": "2023-07-26T21:30:36-03:00",
		"Attributes": {}
	}
}
```

### **POST `/api/users`**
Create a new user.

-   **Request Body:**
```json
{
	"name": "Mikaella",
	"username": "mikaella",
	"password": "mikaellapw"
}
```

-   **Response:** `201`

### **DELETE `/api/users`**
Delete the logged-in user.

-   **Response:** `204`

### **GET `/api/hello`**
Say hello to the logged-in user. Any user can call this.

-   **Auth Required:** Yes
-   **Allowed Roles:** All
-   **Response:** `200`
``` 
"Hi, Mikaella!"
``` 
### **GET  `/api/logs`**
View system logs. Only users with the admin role can access this.

-   **Auth Required:** Yes
-   **Allowed Roles:** Admin
-   **Response:** `200`

```json
[
	{
		"id": 4,
		"method": "GET",
		"path": "/api/hello",
		"username": "geraldo",
		"status": 200,
		"audit_time": "2023-07-07T22:27:13.276794121-03:00"
	},
	{
		"id": 5,
		"method": "GET",
		"path": "/api/logs",
		"username": "geraldo",
		"status": -1,
		"audit_time": "2023-07-07T22:27:21.661260968-03:00"
	}
]
```