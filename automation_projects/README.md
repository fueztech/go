OVERVIEW:
This application automates the scripture treasure hunt I do for my children each week.
I pick a random but well known scriture and hide it in the house for them to discover.
The kids really enjoy this and I do as well, this gave me the idea to create this simple automation.


INSTALLATION:



APPLICATION STEPS:
A Scripture is randomly selected from a struct of scriptures
The gin http framework, handles the http routes.
The Scripture comes down in a json form when a curl command is run aginst the endpoint. 
The Json message will give a location of were to hide the scripture in the house.



FUTURE ENHANCEMENTS:
All of the scriptures reside in memory instead of a database currently.
I plan to add a database to store the spiritual treasures and location of where the treasures were hidden.