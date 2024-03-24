# wechat
A simple chat application


## Folder Struct
---
- /cmd
    - main.go


- /internal
    - /handlers ```Handlers for the HTTP routes```
        - auth_login.go
        - auth_signin.go  

    - /model ```Data models```
        - employee.go ```Hold the all the operation and fields for employee table model```

    - /services ```Business logic layer```

    - /utils ```utilities function```

    - /websocket ```WebSocket handling```


- /pkg
    - /middleware
        - auth_middleware.go 
        ``` authenticate the req which comming from the client side ( frontEnd )```
    
    - /sql ```SQL DB conncetion```
        - conncetion.go
---