
# Pokedex

A REPL project for traversing the pokemon world, exploring areas, catching Pokemon, storing them in your Pokedex and checking the stats of all pokemon you have in your pokedex. This is just a project for and enhancing my skillset.


## Usage/Examples
For traversing the map use 'map' to move forwards and mapb to move back
```
Pokdex > map
Displaying locations 21 to 40 (Page 2)
mt-coronet-1f-route-216
mt-coronet-1f-route-211
mt-coronet-b1f
great-marsh-area-1
great-marsh-area-2
...

Pokedex > mapb
Displaying locations 1 to 20 (Page 1)
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
...
```
For seeing pokemon encounters in a specific area use 'explore'
```
Pokedex > explore great-marsh-area-1
- arbok
- psyduck
- tangela
- magikarp
- gyarados
...
```
For catching a pokemon use 'catch'. If caught, it gets addd to your pokedex
```
Pokedex > catch arbok
Throwing a pokeball at arbok
arbok broke free!
Pokedex > catch arbok
Throwing a pokeball at arbok
Gotcha! arbok was caught!
```
For seeing all the pokemon in your pokedex use 'pokedex'
```
Pokedex > pokedex
Pokedex:
arbok
```
For inpecting the stats of the pokemon in your pokedex use 'inspect'
```
Pokedex > inspect arbok
Name: arbok
Height: 35
Weight: 650
Stats:
  -hp: 60
  -attack: 95
  -defense: 69
  -special-attack: 65
  -special-defense: 79
  -speed: 80
Types:
  - poison
  ```


