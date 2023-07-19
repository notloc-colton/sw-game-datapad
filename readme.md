Summary
Using the Star Wars API (located here: https://swapi.dev/), create an application that
can provide various required information on a provided Star Wars character

Requirements
The application should:
Take a Star Wars character name, and returns the following information about the
character:

Starship
Starship name, cargo capacity, and Starship class
Home Planet
Planet name, population, and climate
Species
Name, language, and average lifespan


Handle as few as one letter in the input
e.g. if the user inputs c, return information for all characters with a c in the name
If multiple characters are found, return in alphabetical order by character name
Do Not:
Make use of any of the SWAPI helper libraries

You are able to use whichever language you are most comfortable with. How you do
this, how you take the input, or display the output is all up to you. Creativity is
encouraged

BE Take Home Technical Interview 2
Additional Considerations
What do you do if you get back multiple characters?
What do you do if you get no characters back?
How to handle multiple starships?
How to handle if a section (starship, home planet, species) is empty?





Bench marks


// Benchmarks
// "l" with everything 6.49, 5.42, 4.73
// "l" with planets ONLY 3.43, 2.27, 5.12
// "p" with planets & species 15.1, 4.56, 4.58
// "p" with everything 18.73, 17.44



// Benchmarks
// "l" with everything
// "l" with planets ONLY
// "p" with planets & species 1.6, 2.24, 2.0
// "p" with everything 2.87, 4.0, 5.73



