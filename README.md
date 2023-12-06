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

There seems to be a lot of different ways to do the string iteration in Go. I'm probably need to do a quick survey of all the string manipulation functions in the standard library. I appreciate the ease of which I can go from string to runes to byte arrays, but that they also have to be cast explicitly. AFAIK the conversions don't require any runtime overhead, but I'm not sure. I'll have to look into that.

I need to go back and fix up how these are structured. Right now I'm just calling the functions of each day module in main.go while also directly referencing the input files in each day. Ideally I'd go back and read the files in from main.go and pass them into the day modules. I don't know what the best format for that would be yet. The three options I have right now are: - pass in the built in `os.File` type - pass in a `bufio.Reader` type - pass in a `string` type

Will make a decision on day 5 once I have enough examples to make an informed decision.

## Day 3

Decided better late than never to start working with different string manipulation functions in the std lib so I tried out scanner. Seemingly less powerful than the `bufio.Reader` but if it does what you need it to specifically then it gives better dx. I'll probably use it when possible but it's nice to know that I can fall back on the `bufio.Reader` if I need to.

Regex is a very powerful tool. In fact it's so powerful I feel as if a lot of the problem solving for a certain class of problems can be done just by reaching for regex. Which when trying to get more comfortable with a tool is actually a pretty bad thing. Like why would I even need to learn the ins and outs of doing something in a language when instead I can just learn how to call into the regex library for that language and then everything else is pretty much handled? I think probably either on day 4 or 5 I'm going to stop using regex all together in order to try and get more comfortable with different aspects of GoLang. I may also start playing with some of the IO functionality to turn a lot of the day solutions into a CLI tool which I think will help expose me to things I haven't had to worry about as much (as most of the GoLang stuff I've done so far has been web apis, JSON parsing, and database stuff).

(cont.)

Yeah that pointer stuff I did was crazy. I would never recommend someone actually do something like that but being able to easily change the reference there was cool. Thinking about how to do something similar in typescript you would have to do something like make the values of the object map a getter or a function or something and then have some datastructure which easily allows you to reference the underlying number which drives the getter/function and sub it out. That obviously would be a lot more work than just holding reference to a pointer and also less performant. Obviously this would have an issue if the gears in part 2 could be 0 without rendering the instance invalid, but that wasn't the case so I didn't have to worry about it. That also could be easily solved by making the underlying "calibration number" stored in the map more complex to track the double counting issue without the hack I did of just setting it to 0.

All in all fun stuff. O(2n) solution complexity depending on how exactly the regex is handling the string finding underneath the hood which I suppose I'm not exactly sure on the specifics of. Maybe something to look into.

Also I did a quick test with 1 million rows and it wasn't that fast honestly. Took about 8 seconds are so which isn't great imo. I may come back and try and do some performance benchmarking and tuning.
