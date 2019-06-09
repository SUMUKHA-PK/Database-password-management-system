# DATABASE PASSWORD MANAGER

## Problem Statement :
Whenever there are CRUD operations to be done by a web server, the password to the database is needed to log on to it.
This DB password must not be compromised as it can have vital data. So, obtaining passwords for the DB by the web server seems to be a pretty
tough task.


## Proposed solution : 
A secure server,in the local network of the other production servers, that enables easy access of passwords
 to the intended user through API endpoints.


## Implementation details : 
* API :
        * "/storePassword" 
        * "/getPassword"
* No need of config files. 
* No plaintext storage of passwords anywhere.
* Only the server who stored the password can access its own passwords.
