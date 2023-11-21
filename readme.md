Description:
  The wanna-be-chuck-norris-app is a single endpoint API that allows you to query for a random name to be inserted into a random chuck norris joke

How to run locally:
  - go get -d ./...
  - add a .env file to your top level directory (same level as this readme)
    - Note: an example .env is included in the "local.env" file
  - go run cmd/main.go
     Note: server will be running on 127.0.0.1:8080 (unless you changed the .env from its default port and address)



How to run unit tests:
 - Unit tests are written in the ginkgo testing framework, which is 100% compatible with native Go testing package. As such, there are two options for running tests:
    - go test ./...
    OR
    - ginkgo -r
        Note: If you don't have ginkgo installed, you will need to install it!
            go install github.com/onsi/ginkgo/ginkgo


API documentation:
 The api is relatively simple, there is one endpoint that, upon request, will provide a randomly-generated-chuck-norris-joke ®.

 Example:
    Request: curl --location --request GET 'http://127.0.0.1:8080/joke'
    Response: 
    {"Joke": "Sulayman Coltrain protocol design method has no status, requests or responses only commands." }



Considerations:

 The prompt for this application asked for a production ready web app, with TODO comments throughout the code in places where additional work will be needed. Given the considerable work that is needed to create a production ready distributed web service, I was not able to create a 100% production ready app (I maxed out my 4 hour time allotment and stopped where I was).
 
 In addition to the comments spread throughout the application, here are some other elements that need to be completed before the web service is ready for production:

  1.) Documentation - While it's true we only have one endpoint with no request payload, it's best to start on the right foot and have some form of API documentation (swagger, postman collection, etc)

  2.) Performance - The need to call two separate API endpoints to supply data for each inbound http request can be costly to our response time. As such, I would recommend some form of data caching if we are allowed to cache data. I stubbed out a simple cache in the application (that caches responses from the random name API) as an example of what I'm thinking here. Ideally, we would want a distributed cache like memcached or Redis, caching responses(from both external APIs) so that our clients have shorter response times.

  3.) Security - I didn't do too much work in terms of security and application hardening. That's a whole other can of worms that we can open and discuss if you'd like.

  4.) Handling failures - I wasn't opposed to handling failures more effectively, but I have a good number of questions that will need to be answered before doing so.

  5.) Logging and cloud monitoring - We will need to decide on what our cloud monitoring solution is going to be. Likewise, we will need to establish a pattern in our logging/tracing so we can easily monitor our applications (search logs, setup alerts, etc)

  6.) Deployment environment - We'll need to decide how our app will be deployed and create the appropriate environments, CI/CD pipelines, Dockerfile (if that's what we are using), helm files (if using Kubernetes), etc.

  7.) Testing - I did some basic unit tests but wasn't able to complete unit tests for all files (or even all of the most important files/functions for that matter). In addition to more unit tests, we will also need to consider other types of tests such as end-to-end tests, smoke tests, static analysis, etc.

