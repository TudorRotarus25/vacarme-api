# Vacarme API

A Golang API for the Vacarme website. It's using MongoDB to serve projects data

## Usage

### Running locally

You can run locally using docker-compose. It will spin up a local MongoDB database to connect to
```
$> docker-compose up
```

You can also simulate the eb environment by using
```
$> eb local run --port 5000 --envvars MONGO_URL='<mongo_url>'
```

### Deploying

The application is deployed to Elastic Beanstalk. You can deploy using the command:
```
$> eb deploy --profile tudorrr
```
