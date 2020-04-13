# Ubidy_AgencyNotificationAPI

This API is intended to be use for sending notification to agency for specific event in the Ubidy Solution.

## Table of Contents

- [Clean Architecture](#clean-architecture)
- [How To Setup](#how-to-setup)
- [Dependencies](#dependencies)

## Clean Architecture

Just to give some context, Uniplaces was created following Domain Driven Design, which is an approach to design your systems prioritizing your business domain model. Despite this being an architecture that really works, for our needs, something leaner should fit us in some of our upcoming projects, since most of the API’s and jobs we wanted to create weren’t having domain logic. We did like the layered approach and separation of concerns DDD with [Hexagonal Architecture](http://alistair.cockburn.us/Hexagonal+architecture) was giving us so the [Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) was the way to go.

After a decision on where to go, it was time to write our first service working with this architecture, we will show you our first attempt in making the theory something tangible that works for us. But first, just a glimpse of what we’re talking about if you’re out of context.

<img src="https://cdn-images-1.medium.com/max/1200/1*DJANEgMHCy4yBjquq1-smA.png" width="400" height="350">

The overriding rule that makes this architecture work is The Dependency Rule. This rule says that source code dependencies can only point inwards. Nothing in an inner circle can know anything at all about the implementation of something in an outer circle.- [The Clean architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)

Let’s now see an example of how we write a service that gives us information about an EmailPreference (being an EmailPreference, an Entity), consider the following directory structure.

- **Delivery**

Repository will store any Database handler. Querying, or Creating/ Inserting into any database will stored here. This layer will act for CRUD to database only. No business process happen here. Only plain function to Database.

This layer also have responsibility to choose what DB will used in Application. Could be Mysql, MongoDB, MariaDB, Postgresql whatever, will decided here.

If using ORM, this layer will control the input, and give it directly to ORM services.

If calling microservices, will handled here. Create HTTP Request to other services, and sanitize the data. This layer, must fully act as a repository. Handle all data input - output no specific logic happen.

This Repository layer will depends to Connected DB , or other microservices if exists.

- **Domain**

This layer, will store any Object’s Struct and it’s method. Example : Email, Client, Supplier.

```go
package emailpreference

// Client model for client email preference
type Client struct {
	ID                   string `json:"id"`
	UserID               string `json:"userID"`
	NotificationSummary  bool   `json:"notificationSummary"`
	NewApplyingAgencies  bool   `json:"newApplyingAgencies"`
	NewCandidateReceived bool   `json:"newCandidateReceived"`
}

// NewClient create new instance of client email preference model
func NewClient(id string, userID string, notificationSummary bool, newApplyingAgencies bool, newCandidateReceived bool) *Client {
	return &Client{
		ID:                   id,
		UserID:               userID,
		NotificationSummary:  notificationSummary,
		NewApplyingAgencies:  newApplyingAgencies,
		NewCandidateReceived: newCandidateReceived,
	}
}
```

- **Infrastructure**

Infrastructure will store any Database handler. Querying, or Creating/ Inserting into any database will stored here. This layer will act for CRUD to database only. No business process happen here. Only plain function to Database.

This layer also have responsibility to choose what DB will used in Application. Could be Mysql, MongoDB, MariaDB, Postgresql whatever, will decided here.

If using ORM, this layer will control the input, and give it directly to ORM services.

If calling microservices, will handled here. Create HTTP Request to other services, and sanitize the data. This layer, must fully act as a repository. Handle all data input - output no specific logic happen.

This Repository layer will depends to Connected DB , or other microservices if exists.

- **Usecase**

This layer will act as the business process handler. Any process will handled here. This layer will decide, which repository layer will use. And have responsibility to provide data to serve into delivery. Process the data doing calculation or anything will done here.

Usecase layer will accept any input from Delivery layer, that already sanitized, then process the input could be storing into DB , or Fetching from DB ,etc.

This Usecase layer will depends to Repository Layer

### Communications Between Layer

Except Models, each layer will communicate through inteface. For example, Usecase layer need the Repository layer, so how they communicate? Repository will provide an interface to be their contract and communication.

Example of Repository’s Interface

```go
package repository

// Application interface for application repository
type Application interface {
	SendApplicationApprove(auctionID string, tenantID string, auctionNumber string, supplierName string, clientName string) error
	SendApplicationApproveSuccessFee(successFeeID string, tenantID string, successFeeNumber string, supplierName string, clientName string) error
	SendApplicationDecline(auctionID string, tenantID string, auctionNumber string, supplierName string, clientName string) error
	SendApplicationDeclineSuccessFee(successFeeID string, tenantID string, successFeeNumber string, supplierName string, clientName string) error
}
```

Usecase layer will communicate to Repository using this contract, and Repository layer **MUST** implement this interface so can used by Usecase

Example of Usecase’s Interface

```go
package usecase

// ApplicationInteractor interface for application usecase
type ApplicationInteractor interface {
	SendApplicationApprove(auctionID string, tenantID string, auctionNumber string, supplierName string, clientName string) error
	SendApplicationApproveSuccessFee(successFeeID string, tenantID string, successFeeNumber string, supplierName string, clientName string) error
	SendApplicationDecline(auctionID string, tenantID string, auctionNumber string, supplierName string, clientName string) error
	SendApplicationDeclineSuccessFee(successFeeID string, tenantID string, successFeeNumber string, supplierName string, clientName string) error
}
```

Same with Usecase, Delivery layer will use this contract interface. And Usecase layer **MUST** implement this interface.

### Testing Each Layer

As we know, clean means independent. Each layer testable even other layers doesn’t exist yet.

- Domain
  This layer only tested if any function/method declared in any of Struct.
  And can test easily and independent to other layers.
- Repository
  To test this layer, the better ways is doing Integrations testing. But you also can doing mocking for each test.
- Usecase
  Because this layer depends to Repository layer, means this layer need Repository layer for testing . So we must make a mockup of Repository that mocked with mockery, based on the contract interface defined before.
- Delivery
  Same with Usecase, because this layer depends to Usecase layer, means we need Usecase layer for testing. And Usecase layer also must mocked with mockery, based on the contract interface defined before

## How To Setup

1. As Go uses open-source (FREE!) repositories often, be sure to install the Git package [here](https://git-scm.com/download/win) first.
2. Navigate to the Go installation website [here](https://golang.org/doc/install). Download and install the latest 64-bit Go set for Microsoft Windows OS.
3. Follow the instructions on the Go installation program.
4. Run the Command Prompt on your computer by searching for “cmd”. Open the command line and type: “go version”
5. The output after entering go version should look like this:
   <img src="https://cdn-images-1.medium.com/max/1600/1*-j7JjyJSN3DqxEdO4lrjTw.png">

### Phase 2: Creating your Go work-space (Windows Only)

First, confirm your Go binaries: go to your computer’s Control Panel, then to System and Security > System > Advanced system settings, and on the left-hand pane click the Advanced tab. Then click on Environmental Variables on the bottom-right-hand side. Ensure Path under System Variables has the “C:\Go\bin” variable in it.

Then create your Go work-space. This will be in a separate and new folder from where the Go installation files are saved. For example, your G installation files were saved under the path C:\Go and you are creating your Go work-space under C:\Projects\Go

In your new Go work-space folder, set up three new folders:

<img src="https://cdn-images-1.medium.com/max/1600/1*I3BO4S6FQ6keH6o75ATuBg.png">

### Phase 3: Create the GOPATH environment variable

Create the GOPATH variable and reference your newly-created Go work-space. Go back to your Control Panel and navigate to System and then Environmental Variables. Then under System Variables click on New.

Next to Variable Name, enter “GOPATH,” and next to Variable Value enter “C:\Projects\Go”

<img src="https://cdn-images-1.medium.com/max/1600/1*EdndcOEfhY8DWreAWXjung.png">

<img src="https://cdn-images-1.medium.com/max/1600/1*ErNq0vYJQeTJadnJZczBtw.png">

To check that your path has been set correctly, enter “echo %GOPATH%” on the command line.

### Phase 4: Test and ensure

Now you’re ready to verify that all is working correctly by opening the command line and typing: `go get github.com/golang/example/hello`

Wait for the code to be entirely implemented (this could take a few seconds), then enter in the following in the command line: `%GOPATH%/bin/hello`

If the installation was successful, you should get the following return message: “Hello, Go examples!”

<img src="https://cdn-images-1.medium.com/max/1600/1*EXG3IKaDbFqJ3qMpD_n08Q.png">

## Dependencies

Enter the following commands to install dependencies:

```
go get github.com/auth0-community/auth0
go get github.com/denisenkom/go-mssqldb
go get github.com/gin-gonic/gin
go get github.com/google/uuid
go get github.com/gin-contrib/cors
go get "gopkg.in/GetStream/stream-go2.v1"
go get github.com/Microsoft/ApplicationInsights-Go/appinsights
go get gopkg.in/mgo.v2
```
