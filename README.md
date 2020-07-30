# go-fetch-spotify
Goes to Spotify and pulls back details on an artist of your choice, and displays an ASCII art image of the artist.

Using this exercise to build my very first Golang app, after having completed [A Tour Of Go](https://tour.golang.org/).

# Getting started
```
$ go run main.go --help
Usage:
  -artist string
        Artist name to lookup
  -token string
        Auth token to use for query
```

## Get your Spotify token
```
curl -X "POST" -H "Authorization: Basic <base64 encoded client_id:client_secret>" -d grant_type=client_credentials https://accounts.spotify.com/api/token
```
As per [Spotify documentation here](https://developer.spotify.com/documentation/general/guides/authorization-guide/#client-credentials-flow)

```
$ go run main.go --artist Korn --token <your_spotify_token_here>
Name: Korn
Popularity: 75
Type: artist
Followers: 3898555
Genres:[alternative metal funk metal nu metal post-grunge rap metal rock]
Image:
LLLLLLLLLCCCCCCCCCCCCCCCCCCCCCLLLLLLLLLLfffffftttt
LLLLLCCCCCCCCCCCCCCCCCGCCCCCCCCCLLLLLLLLLLffffffff
LCCCCCCCCCCCGGGGGGGGGG11CffGCCCCCCCCLLLLLLLLLfffff
CCCCCCCGGGGGGGGGCGGG0C.,i;:GGGCGCCCCCCCCCLLLLLLLff
CCCCCGGGGCCGG0111;CG1:..,..:1LGfLLLGCCLffLCLLLLLLL
CCGGGGGGGiiL0C,.;.,: ..,.,,,.:f,1t:LGL,;;;CCLLLLLL
CGGGGGGCf;i11;...........:,,,,,....,1;...,:1CLLLLL
GGGGGGG:  .. ....,,,...,,,,,;.,,,....,,..,..LCCCLL
GGGGGGL.......;:,,. ,,.....,, ..,i1:.,...,..LCCCLL
CCCCCCGL. :, .;:.............,....,,........fCCCLL
LLLLLLCC, ....  .......,,,...:,...:;;,....,tCCCCCL
ffffffff, ;1 ,; ,1 ...............,.1; it ,CCCCCCL
ttt11ttt, tL,.: iC: .......... .... ;; tG: fCCCCCC
ftt111tt.,fL1 . ;Lt.,.. iL, ,f;:.,..i; fC1 LCCCCCC
CLLLffti ;ftt,. ;LL,::.;tLi ;LLt 1:1L:.fC1 1GGGGGG
GGGGGGCt1fCLL1,,it1::,.LCCL ;f1ii;.t0LG00GCC000000
GGGG0000800000GGGGCt,,iCGGC:.t8880G088800888000000
GGG00000000000000080G0888000GG00008000000000000000
GGGG0000000000000000800000008800000000000000000000
GGG00000000000000000000000000000000000000000000000
```

# TODO
- Write some tests
- Pipeline: Add linting and other go quality tools
