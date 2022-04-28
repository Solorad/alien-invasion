### Alien invasion
This is a golang simulation for alien invasion.
Cities are stored in a `city_map.txt` file. Number of aliens is passed as an argument to this application.

### Application logic
The logic of whole application is split on 5 phases:
1. Read aliens number and the city map. Graph of cities roads is directed, that means that there might be a situation when
there is a road from city A to city B, but not backward.
2. Generate alien names. Just for fun I added a small utility to generate random names using a set of a start names.
3. Aliens are randomly put on a map. I decided to allow several aliens to ground in the same city - that protects 
us from the situations when there are more aliens than cities and someone can't ground. Right after that we check 
cities and aliens state. 
4. In a loop for every alien - move it to the neighbour city. If there's no roads out of the city - leave the alien there.
After all aliens made a move - check cities and aliens state.
5. Print out in console current city map

### Engineering decisions
- For this small test app I decided to use a project structure to which I get used to  -  with `cmd`, and `pkg` folders.
- We can easily add a validation for `trapped` aliens - one that are currently in a city subgraph without any other aliens.
Finding of those might simplify logic of the application, but due to we need 10.000 cycles - decided not to implement this
- Several unit tests were added for core functionality. I decided to limit total work time on this application to 6 hours, so 
there are no big parameterized test suites

I