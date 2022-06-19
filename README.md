# Currency alerter
Application allows user to create alert, which are triggered after watched exchange rate of desired currency satisfies set treshold.

It is a server side application which is responsible creating user alerts using http calls, fetching exchange rates from external api and triggering alerts if user's condition is satisfied.

## How to setup
For the development purposes, it uses docker-compose, due to the simplicity of creating a dev environment 

Run the command below in order to create neccesary services
```
docker-compose up -d 
```

By default service with Http endpoint for creating alerts is exposed at `:5000` port