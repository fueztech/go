## OVERVIEW:
 -- This application automates the scripture treasure hunt I do for my children each week.
 -- I pick a random but well known scriture and hide it in the house for them to discover.
 -- The kids really enjoy this and I do as well, this gave me the idea to create this simple automation.

_______________________________________________________________________________________________________

## APPLICATION STEPS:
 -- A Scripture is randomly selected from a struct of scriptures
 -- The gin http framework, handles the http routes.
 -- The Scripture comes down in a json form when a curl command is run aginst the endpoint. 
 -- The Json message will give a location of were to hide the scripture in the house along with the date and time.

_______________________________________________________________________________________________________

## FUTURE ENHANCEMENTS:
 -- Add a database to store the spiritual treasures and location of where the treasures were hidden.
 
 -- Send the Scripture via sms message to my phone, using aws or someother cloud provided solution.
 
 -- Add jwt authentication to the endpoint where this is deployed for security.
 
 -- Add a cron job to run my curl command weekly vs manually running the command.
