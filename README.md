# catfactsAPI
Tests for cat facts app's API - Sporty interview assignment.
Implemented for this assignment - tests 1, 2, 5.

Test 1 - Get request for cat facts:
    We want to write an automation that will send a get request to 
    "https://cat-fact.herokuapp.com at the endpoint for cat facts. E.g 
    we would like to perform the equivalent action to putting the endpoint
    address - "https://cat-fact.herokuapp.com/facts?animal_type=cat"
    in the address bar.
    Expected result: a response with a json body that has cat facts.

Test 2 - Get request without authentication:
    we parform the former test without the authentication. 
    Expected result: A message saying that we need to sign in.

Test 3 - Get request for my user:
    Similarly to the former test, we would like to perform user 
    authentication and send a get request that is equivalent to:
    "https://cat-fact.herokuapp.com/users/me".
    Expected result: A json with details about my user.

    ** NOTE: This was not implemented since I didn't have enough time to
    build a google authentication mechanism. **

Test 4 - Get request for a different user than the one authenticated.
    We authenticate with a google account but request a different user
    in the URL.
    Expected result: A message saying we don't have access or an error
    message.

    ** NOTE: This was not implemented since I didn't have enough time to
    build a google authentication mechanism. **

Test 5 - Post request.
    We send a request with a json in its body from port 3333 in our
    client.
    Expected result: An error message or a response id of type 4**.