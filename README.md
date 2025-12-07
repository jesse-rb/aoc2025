# Advent of code 2025

(this might be a bit silly)

[advent of code problems](https://adventofcode.com/2025)

**REQUIREMENTS**

- Docker and Docker Compose binaries (Docker Desktop comes with both)
- Ability to run Makefile

**Makefile**
```
> make help

make run DAY='<day>'   - Run a day e.g. make run DAY='day1''
make clean             - Stop docker container and clean images
```

**Running the solutions**

First make sure you have your input to pipe as stdin into the solution

```
make run DAY='<service directory>'

for example:
echo "my input" | make run DAY='day1'
cat input.txt | make run DAY='day1'
make run DAY='day1' < input.txt
```
