### Alien invasion
This is a golang simulation for alien invasion.
Cities are stored in a `city_map.txt` file. The number of aliens is passed as an argument.

### Application logic
Logic of the application is split on 5 phases:
1. Read aliens number and the city map. Graph of cities roads is directed, that means that there might be a situation when
there is a road from city A to city B, but not backward.
2. Generate alien names. Just for fun I added logic to generate random names using a set of start names.
3. Aliens are randomly put on the map. I decided to allow several aliens to ground in the same city - that protects 
us from the situations when there are more aliens than cities and someone can't ground. Right after that we check 
cities and aliens state. 
4. In a loop for every alien - move it to the neighbour city. If there's no roads out of the city - leave the alien there with no move.
After all aliens made a move - check cities and aliens state. I decided to add validation for city states in the end of a move,
so that would protect from aliens move ordering question but makes possible a case when a city is destroyed by more than 2 aliens.
5. Print out in console current city map

### Engineering decisions
- For this small test app I decided to use a project structure to which I get used to  -  with `cmd`, and `pkg` folders.
- We can easily add a validation for `trapped` aliens - one that are currently in a city subgraph without any other aliens.
Finding of those might simplify logic of the application, but due to we need 10.000 cycles - decided not to implement this
- Several unit tests were added for core functionality. I decided to limit total work time on this application to 8 hours, so 
there are no big parameterized test suites
- Besides, I added to the printed message information about the step on which city collapsed. So, message looks like: 
  > [grounding] Foo has been destroyed by Helvetios and Deneb!
  > 
  > [step-4] Qu-ux has been destroyed by Wasat and Pleione!
- It was decided for simplicity to support only one-worded cities, without spaces. For example `Paris` is ok, 
when  `San Francisco` is not.

### Application usage
**Prerequsites:** `city_map.txt` file in a root folder. You can use one from the root or update it with a new map 

To run this application run the next command:
```bash
go build -mod vendor -o madaliens cmd/alien-invasion/main.go
./madaliens 3
```

### Map generator
For test purposes I also developed a map generator. To generate a test map you should run
```bash
go build -mod vendor -o map-generator cmd/map-generator/main.go
./map-generator 200
```
Output of this utility can be used by our app
```bash
./map-generator 200 > city_map.txt
```