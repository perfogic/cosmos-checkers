#!/bin/sh

rly config init
rly chains add-dir configs
rly paths add-dir paths

rly keys restore cosmoscheckers alice "cinnamon legend sword giant master simple visit action level ancient day rubber pigeon filter garment hockey stay water crawl omit airport venture toilet oppose"
rly keys restore cosmoscheckers bob "stamp later develop betray boss ranch abstract puzzle calm right bounce march orchard edge correct canal fault miracle void dutch lottery lucky observe armed"

rly keys restore leaderboard alice "cinnamon legend sword giant master simple visit action level ancient day rubber pigeon filter garment hockey stay water crawl omit airport venture toilet oppose"
rly keys restore leaderboard bob "stamp later develop betray boss ranch abstract puzzle calm right bounce march orchard edge correct canal fault miracle void dutch lottery lucky observe armed"

rly tx link cosmoscheckers -d -t 3s --src-port leaderboard --dst-port leaderboard --version leaderboard-1
rly start cosmoscheckers