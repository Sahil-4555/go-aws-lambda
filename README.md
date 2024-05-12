# go-lambda

This project is a simple example of using Go with SQLC for SQL generation with Migrations, MySQL as the database, and integrating with AWS Lambda and AWS API Gateway.

## Setup

### Prerequisites

- Go installed on your system ([Installation Guide](https://golang.org/doc/install))
- Database migrations ([Migration](https://github.com/golang-migrate/migrate))
- SQLC installed and running ([SQLC Installation Guide](https://sqlc.dev/))
- AWS account with access to Lambda and API Gateway ([AWS Free Tier](https://aws.amazon.com/free/))

### Install SQLC

SQLC is a tool for generating type-safe Go code from SQL queries. Install it using the following command:

```
sudo snap install sqlc
```
#### Create sqlc.yaml:

- Create a `sqlc.yaml` file in the root directory of your project.
- Configure your SQLC settings in this file. You can refer to the [SQLC documentation](https://docs.sqlc.dev/en/stable/overview/install.html) for more information on configuring sqlc.yaml.

#### Create schema.sql:

- Create a `schema.sql` file in the root directory of your project and write your database schema.
- Create a `query.sql` file in the root directory and write your SQL queries in this file.

#### Generate Go Code using SQLC

- Generate Go code from SQL queries using SQLC:

```
sqlc generate
```

#### Setup Migrations

```
curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz
```
- Create a new migration file.
```
migrate create -ext sql -dir migrations/ -seq <filename> 
```

- Revert migrations down to a specific version.

```
migrate -path migrations -database "mysql://username:password@tcp(localhost:3306)/databasename" -verbose down <version>
```

- Migrations up to a specific version.
```
migrate -path migrations -database "mysql://username:password@tcp(localhost:3306)/databasename" -verbose up <version>
```

- Force Apply a Specific Migration Version
```
migrate -path migrations -database "mysql://username:password@tcp(localhost:3306)/databasename" -verbose force <version>
```

## AWD Lambda Integartion

Create an Compiled Binary as the Code

```
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -o bootstrap main.go
zip  -o lambda-handler.zip bootstrap
```

### Deploying to AWS Lambda

1. Go to AWS Lambda Console:

2. Create a New Function:

- Click on "Create function".
- Choose "Author from scratch".
- Enter a function name of your choice.
- Select "Amazon Linux 2023" as the runtime.

3. Configure Function:

- Under "Function code", select "Upload a .zip file".
- Upload the compiled binary file (main) that you created earlier.
- Set the handler to the name of your main function (e.g., main).

4. Set Environment Variables (if necessary):
- Configure any necessary environment variables for your Lambda function.

5. Create Function:
- Click on "Create function" to create your Lambda function.

### Setting up API Gateway

1. Go to API Gateway Console:

2. Create a New REST API:
- Click on "Build".
- Choose "REST API".
- Choose "New API".
- Enter an API name of your choice.

3. Create Resources:
- Click on "Actions" and select "Create Resource".
- Enter a resource name and path (e.g., /).

4. Create Method:
- Select the resource you just created.
- Click on "Actions" and select "Create Method".
- Choose the HTTP method you want to use (e.g., POST, GET).
- Select "Lambda Function" as the integration type.
- Choose the Lambda region and enter the Lambda function name.
- Enable "Use Lambda Proxy Integration".

5. Deploy API:
- Click on "Actions" and select "Deploy API".
- Choose the stage (e.g., "prod").
- Click "Deploy".


