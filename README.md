# Pokédex CLI  <img width="32" height="32" alt="pokeball" src="https://github.com/user-attachments/assets/7b6866a8-3154-48bb-aced-3817c2879c0a" />

Pokédex CLI is a command-line (REPL) application written in Go that allows users to explore regions, catch Pokémons, inspect their details, and view all caught Pokémons. It was developed as part of the 'Build a Pokédex' project in the **Backend Developer path** on [Boot.dev](https://www.boot.dev/).

The app interacts with **[PokéAPI](https://pokeapi.co/)** to retreive up-to-date Pokémon data and supports smart caching for API responses to improve performance.


## Supported Commands

| Command                | Description                                                |
|------------------------|------------------------------------------------------------|
| `explore [location]`   | List all Pokémons available in a given location.           |
| `catch [pokemon]`      | Attempt to catch a Pokémon given its name based on calculated probability.                 |
| `inspect [pokemon]`    | Show details of a Pokémon you've caught.                   |
| `pokedex`              | List all Pokémon you've caught in your Pokédex.            |
| `mapn`                 | Move to the next location area.                            |
| `mapp`                 | Move to the previous location area.                        |
| `exit`                 | Exits the Pokédex.                                         |

Refer to the application's `help` for a full list of commands.


## Pokémon Catching Mechanism

Catching a Pokémon depends on randomness and the base experience of the Pokémon; the higher the experience, the harder to catch it.

[Exponential Decay](https://gregorygundersen.com/blog/2022/05/17/exponential-decay/) method was used to calculate the catch chance:
```latex
catch_chance = base_chance * e^(-decay_const * scaled_experience)
```
The Pokémon is caught if the `catch_chance` is greater than a random number generated between `[0.0, 1.0)`.

### Formula:
<img width="291" height="87" alt="image" src="https://github.com/user-attachments/assets/57ac1849-6fb1-44cd-84b2-52a4d201d0fc" />


## Installation

1. **Clone the Repository**
   ```sh
   git clone https://github.com/MaysHroub/pokedexcli.git
   cd pokedexcli
   ```

2. **Build the Application**
   ```sh
   go build -o pokedex
   ```

3. **Run the CLI**
   ```sh
   ./pokedex
   ```

## Notes

Yes, I *over-engineered* it :)




