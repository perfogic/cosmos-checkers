#!/bin/bash

Sleeptime=1
Gameid=$1

Players=($2 $3)

Moves1=(
"1 2 2 3" 
"0 5 1 4" 
"2 3 0 5" 
"4 5 3 4" 
"3 2 2 3" 
"3 4 1 2" 
"0 1 2 3" 
"2 5 3 4" 
"2 3 4 5" 
"5 6 3 4" 
"5 2 4 3" 
"3 4 5 2" 
"6 1 4 3" 
"6 5 5 4" 
"4 3 6 5" 
"7 6 5 4" 
"7 2 6 3" 
"5 4 7 2" 
"4 1 3 2" 
"3 6 4 5" 
"5 0 4 1" 
"2 7 3 6" 
"0 5 2 7" 
"4 5 3 4" 
"2 7 4 5" 
)

Moves2=(
"4 5 2 3" 
"6 7 5 6" 
"2 3 3 4" 
"0 7 1 6" 
"3 2 4 3" 
"7 2 6 1" 
"7 0 5 2" 
"1 6 2 5" 
"3 4 1 6" 
"4 7 3 6" 
"4 3 3 4" 
"5 6 4 5" 
"3 4 5 6" 
"3 6 2 5" 
"1 6 3 4" 
)

cosmos-checkersd tx cosmoscheckers create-game ${Players[1]} ${Players[0]} 0 token --from ${Players[1]} --chain-id cosmoscheckers --yes

for i in "${!Moves1[@]}"
do
    cmd="cosmos-checkersd tx cosmoscheckers play-move $Gameid ${Moves1[i]} --from ${Players[$(((i+1)%2))]} --chain-id cosmoscheckers --yes"
    $cmd
    sleep $Sleeptime
done

for i in "${!Moves2[@]}"
do
    cmd="cosmos-checkersd tx cosmoscheckers play-move $Gameid ${Moves2[i]} --from ${Players[$(((i+1)%2))]} --chain-id cosmoscheckers --yes"
    $cmd
    sleep $Sleeptime
done