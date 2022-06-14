
1. Request comes in to server on route: /api/project
2. Request is passed through middleware to:
    2.1: Validate the request data (numbers, letters, symbols, etc) 
    2.2: Check in the header what agent type is sent
3. Unmarshal the request into the appropriate struct.
4. Create / generate the appropriate files 
5. Based on the user agent make the files available for the user.


In the future, we will also add authentication to create presets we can store in a 
database of some sorts.

We can also explore how we are going ot handle multiple requests.


A big challenge I have is that while the project is technically small, with a single route, I still want it to function
propertly. 

Yes, a request can come in. The file is created based on some data. What do I do with this data, though? Do I store it somewhere temporarily?
Other than a struct, that is, of course. Do I create a small entry saying "right user with session ID XYZ is requesting this file"?
Do I store these temporary files anywhere?

Middleware design:

Either with reflection/string, so I just provide an array/slice of string names for each middleware.
Then call a function `handle` or `middleware` to handle the middleware



Testing Middleware

What do I even want to test? First, we want to test whether the route
activates all the middleware. Whether that is a responsibility of middleware
or main remains to be discussed.

Middleware should take in mock keys and a handler. Maybe default Mux?
Middleware simply returns a function. That return function runs all the 
middleware. 

So, what do we do there? Do we just check if every function is run?


main contains our listener and router initialiser

main