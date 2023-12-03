Advent of code 2023
by: @i-Welch
in: GoLang

## Day 1

Fun string manipulation stuff. Playing with regex and string parsing.
Interesting question I've come to ask myself: Is all regex universally reversible? ie: a "reversal" of regex "\d" is "\d" or a reversal of regex "nine" is "enin". You see that I used that to solve part 2. I did a cursory google search of "is all regex reversible", but the results were not helpful. I'll have to look into this more, there is probably a academic paper on this hiding somewhere. I assume once you get into needing look backs and such this gets to be a more complicated problem but for a subset of the regex feature set it seems easily true.

I also played around with using a trie data structure, due to the specifics of the problem I decided it would actually be easier to not so I abandoned that idea part way through, but I was surprised by how easy it was to make a data structure like that in GoLang. I'd likely not hand roll my own in any production code, but know I can fairly easily is nice all the same. The fact that recursive data structures seem to work pretty easily is also nice. Looking online I was led to believe that GoLang was going to make me make the recursive structure out of pointers but it seems like I didn't have to most likely due to how the internal map[] data structure works.

The module system in go is decent. I don't love it nor hate it yet. I think coming from the hell that is javascript module resolution and the whole ecosystem around that my instincts for that are probably off. I'll keep this in mind as this project grows.

Standard libraries are nice.

## Day 2

(cont.)
