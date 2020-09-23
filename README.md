## Echo Validator
EchoValidator is a fail fast extensible validation library only compatible with [Echo Framework](https://github.com/labstack/echo).

#### What does Fail Fast Mean?
When provided a list of validation rules to check, `echovalidate` will bail on the first error and return the error message.
As opposed to running all validations and collecting errors in a bag.

#### Example
```go
    package main

    import (
        "github.com/labstack/echo/v4"
        "github.com/jcobhams/echovalidate/v2"
    )
    
    func main() {
        
        //Initialize Echo Instance
        e := echo.New()
        
        //Register Validator 
        e.Validator = echovalidate.New()
        
        //Register Route Handler
        e.POST("/some-endpoint", handler)
        
        //Start Server
        e.Start(":9009")
    }


    func handler(ctx echo.Content) error {
        username := "jcobhams" // Could come from post request body
    	
        v := echovalidate.Validator{
    		Rules: echovalidate.Rules{
    			{echovalidate.Required, "username", username},
    			/** more validation rules here **/
    			{echovalidate.MinLen, "username", len(username), 5},
    		},
    	}
    	err := ctx.Validate(v)
    	if err != nil {
    		return  ctx.JSON(400, err)
    	}
        
        /** All is well - continue as normal **/
        return ctx.JSON(200, "OK")
    }
```

#### Built In Validation Rules
1. `Required` - Checks that value is not empty if type is string | only string type supported for now - PR for Slices/Maps?
2. `ValidEmail` - Checks that an email is properly formed. Uses regex to pattern match - Valid Email: person@domain.tld - PR for MX validation?
3. `In` - Checks that needle exist in a haystack | Only supports string slices
4. `MinLen` - Checks if the provided length (int) is less than the minimum length and returns error if true.
5. `MaxLen` - Checks if the provided length (int) is greater than the maximum length and returns error if true.
6. `ValidMongoObjectID` - Checks if provided value conforms to MongoDB object ID hexadecimal
7. `ValidDateTime` - Checks if provided date parses according to provided layout string.

#### Extending Validation Rules.
Validation Rules are pure functions and must follow a certain format.
Signature: `func FuncName(args...) error {}`
This means that Your custom validation functions can exists anywhere in your program as long as they can be imported
and used to build the `Rules` list

Your validation functions must always have a return type of `error`, which should contain the error message if validation 
fails or nil if validation passes.

The args to your validation function must be provided to `Rules` in the order in which they are defined at function definition

##### Example Custom Validation Function
```go
func SuperChecker(value string) error {
    /** 
    Do something with value, if successful, return nil else return error
    **/
    
    if len(value) > 1 {
        return nil
    }
    return errors.New("value should be more than 1 char")
}
...

    v := echovalidate.Validator{
    		Rules: echovalidate.Rules{
    			{SuperChecker, username},
    		},
    	}
...
```

#### Why?
Because I needed something simple, lightweight, flexible and most importantly didn't follow the errorbag design.
I needed something that would fail fast on first validation error.

Use if you have a need/like/curious/dislike etc.

PR/Feedback welcomed and encouraged.

#### Where is 1.x
Project was originally on gitlab but was moved to github. To avoid confusion, the github version starts at 2.x