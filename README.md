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

```
$ go run main.go --artist Eminem --token <your_spotify_token_here>
Name: Korn
Popularity: 75
Type: artist
Followers: 3898555
Genres:[alternative metal funk metal nu metal post-grunge rap metal rock]
Image:
LLLLLLLLLLLLLLLLLLLLCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCLLCLLLLLLLLLLLLLLLLLLLffffffffffftttttttt
LLLLLLLLLLLLLLLLCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCLLLLLLLLLLLLLLLLLLLLLfffffffffffftttttt
LLLLLLLLLLLLLCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCLLLLLLLLLLLLLLLLLLLLffffffffffffffftt
LLLLLLLLCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCGGLCCCGCCCCCCCCCCCLLCLLLLLLLLLLLLLLLLLLLLffffffffffffff
LLLCLCCCCCCCCCCCCCCCCCCCCGGGGGGGGGGGGGGGGGGGCt;tCCttCCCCCCCCCCCCCCCCCCLLLLLLLLLLLLLLLLLLffffffffffff
CCCCCCCCCCCCCCCCCCCCCGGGGGGGGGGGGGGGGGGGGGGGi ;ftCGi,CGCCCCCCCCCCCCCCCCLCCLLLLLLLLLLLLLLLLLfffffffff
CCCCCCCCCCCCCCCCCCGGGGGGGGGGGGGGGGGGGGGGGGGL,.,i;tf1.1GCCCCCCCCCCCCCCCCCCCCCLLLLLLLLLLLLLLLLLLLLffff
CCCCCCCCCCCCCGGGGGGGGGGGGGGGG0GCCG0GGGGGGG0f,...,;:,.:CGGGGGGCCCGCCCCCCCCCCCCLCCCLLLLLLLLLLLLLLLLLLf
CCCCCCCCCGCGGGGGGGGGGGGGGGGGG11ti:;LGGG0GL1,.....;,...;ffCGGGGGCLCLLGCCCCCCCCCCLLCCLLLLLLLLLLLLLLLLL
CCCCCCCCCGGGGGGGGGGGCGGGGGG0t ;1ft,.1G0t:. .............,:;CGGf;tLL1;CCCCCCCC1ittifCLLLLLLLLLLLLLLLL
CCCCCCGGGGGGGGGGGC;::;tGGGGC:..,;t:. ;t......,,....:,,,,,:.i0C,;tLCL:1GCCCCG;.itf;.LCLLLLLLLLLLLLLLL
CCCGGGGGGGGGGGGGGG;:tffCGGGf.....,......,.....,,..,:,,:,,,,.fi.,;1fi:.iGGCGt,,,::,.iCCCLLLLLLLLLLLLL
CGGGGGGGGGGGGGG000G::fLfCGf:...................,.,,,,.,:,...,,....,..,.;LCt,,,....,,;i1CCCLLLLLLLLLL
CGGGGGGGGGGGGGGti;,.,t1..:,..........,,,....,,,...::::,...,,,,...........,,,,,,...,,,. 1CLCCLLLLLLLL
GGGGGGGGGGGGG0t   ....... ..........,,,,......,,,,:;:,..;:.,,,,...........,,:.,...,,,,.,LCCCCCCLLLLL
GGGGGGGGGGGGGG,.........,,.........,,,,..  ....,,:..,...f: ...,,:,,.......,,,.....,,,,.,LCCCCCCCLLLL
GGGGGGGGGGGGG1 ............,,. .,:,,,... :; ...,.......,i.......,::;t;:...,,......,,.,.,LCCCCCCCLLLL
GGGGGGGGGGGGGf, ...........,;t1;:,...... ,: ............... ......;fLti;,..,......,,....fCCCCCCCLLLL
CCCCCCGGGGGGGGGt....,;: ...:i11i,......... ................,..........,:............... 1CCCCCCCCLLL
CCCCCCCCCCCCCCCG: . ,i, ......  ...........................i.........   .............., iGCCCCCCCLLL
LLLLLLLLLCCCCCCC: .. ............. ...........,,,,,,.......i,.......,;;::;,..........;i1CCCCCCCCCLLL
LLLLLLLLLLLLLLLC: ..... .. ,,.....,............,,,,........:, ......:it;it:..........:fGCCCCCCCCCCLL
ffffffffffffLLLL, .. ;f....,i.....C, ................................,,.,f; .. 1f... ,CCCCCCCCCCCCLL
ffffftfffffffffL: ...fGi ...L,.. ,G; ...................................:ft....LG1 ..,CCCCCCCCCCCCCL
ttttttttttttttff: . ;LLf....f, . 1G1; ..................................,f1.. ;CCC: ..fCCCCCCCCCCLCL
ttttt1111111tttt,...tLfL1 . ,....LCfL, ...........  ......,,.............i1.. tCLCf . 1GCCCCCCCCCCCC
fttttt1111111ttt,...tfffL,.......fCfLi ..,..... :ttf, ....L1   ........ :i1...tLLCf.. tGCCCCCCCCCCCC
ffffttttttttttt1.. :ffffL1 ......tCtf1 ..:..... :CLC; .. ;Cf11;:..,:...:itt, .fLLCC:  tGCCCCCCCCCCCC
LLLLLfffftttttt1...1ttfff1:,.....tLffL:  1: .., :ffft... tCCCCGf..iL::;tfft...1LLCL,. iGCCCCCCCCCCCC
CCCCCCCLLLLfftfi  .tfttttii, ...,fLLLLi .f1 ..tttfttf: ..tLLLCf;..:G;.;CCt, .,fCLLf,  .LGGGGGGGGGGGG
GGGCGGGGCCCCCL;.,:;tLffttt1:.   .1tttfi  it..,CGGCCCCf..,fii;:  .  i. ,G0LffLG0000GGfttf000000000000
GGGGGGGGGGGGGGCCGGGGGCCCCCC1:,:;i11iii;,,,,,.,ffLLCG0G,..tLfLfLLLLt,..,G0088000000008888000000000000
GGGGGGGGG000000000000000000GGGGGGGCCLffti, ...tCLLCCCf ...L88888888GGG000000000000000000000000000000
GGGGGGG0000000000000000000000000000008C:,,:ifL0880000C1:,. 18000000888800000000000000000000000000000
GGGGGG000000000000000000000000000000000GG008880000008880GGGG0000000000000000000000000000000000000000
GGGGG00000000000000000000000000000000008800000000000000088880000000000000000000000000000000000000000
GGGGGGG000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
GGGGG00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
GGGGG00000000000000000000000000000000000000000000000000000000000000000000000000000000000000800008000
GGGGGGG000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
```

# TODO
- Write some tests
- Setup build (run tests, linting and other go quality tools)
