# Day 5: If You Give A Seed A Fertilizer 

## Part 1

You take the boat and find the gardener right where you were told he would be: managing a giant "garden" that looks more to you like a farm.

"A water source? Island Island **is** the water source!" You point out that Snow Island isn't receiving any water.

"Oh, we had to stop the water because we **ran out of sand** to [filter](https://en.wikipedia.org/wiki/Sand_filter) it with! Can't make snow with dirty water. Don't worry, I'm sure we'll get more sand soon; we only turned off the water a few days... weeks... oh no." His face sinks into a look of horrified realization.

"I've been so busy making sure everyone here has food that I completely forgot to check why we stopped getting more sand! There's a ferry leaving soon that is headed over in that direction - it's much faster than your boat. Could you please go check it out?"

You barely have time to agree to this request when he brings up another. "While you wait for the ferry, maybe you can help us with our **food production problem**. The latest Island Island `Almanac` just arrived and we're having trouble making sense of it."

The almanac (your puzzle input) lists all of the seeds that need to be planted. It also lists what type of soil to use with each kind of seed, what type of fertilizer to use with each kind of soil, what type of water to use with each kind of fertilizer, and so on. Every type of seed, soil, fertilizer and so on is identified with a number, but numbers are reused by each category - that is, soil `123` and fertilizer `123` aren't necessarily related to each other.

For example:

> seeds: 79 14 55 13
>
> seed-to-soil map:
> 50 98 2
> 52 50 48
>  <br>
> soil-to-fertilizer map: <br>
> 0 15 37 <br>
> 37 52 2 <br>
> 39 0 15 <br>
>  <br>
> fertilizer-to-water map: <br>
> 49 53 8 <br>
> 0 11 42 <br>
> 42 0 7 <br>
> 57 7 4 <br>
>  <br>
> water-to-light map: <br>
> 88 18 7 <br>
> 18 25 70 <br>
>  <br>
> light-to-temperature map: <br>
> 45 77 23 <br>
> 81 45 19 <br>
> 68 64 13 <br>
>  <br>
> temperature-to-humidity map: <br>
> 0 69 1 <br>
> 1 0 69 <br>
>  <br>
> humidity-to-location map: <br>
> 60 56 37 <br>
> 56 93 4 <br>

The almanac starts by listing which seeds need to be planted: seeds `79`, `14`, `55`, and `13`.

The rest of the almanac contains a list of **maps** which describe how to convert numbers from a **source category** into numbers in a **destination category**. That is, the section that starts with `seed-to-soil map`: describes how to convert a **seed number** (the source) to a **soil number** (the destination). This lets the gardener and his team know which soil to use with which seeds, which water to use with which fertilizer, and so on.

Rather than list every source number and its corresponding destination number one by one, the maps describe entire **ranges** of numbers that can be converted. Each line within a map contains three numbers: the **destination range start**, the **source range start**, and the **range length**.

Consider again the example seed-to-soil map:

> 50 98 2 <br>
> 52 50 48 <br>

The first line has a **destination range start** of `50`, a **source range start** of `98`, and a **range length** of `2`. This line means that the source range starts at 98 and contains two values: `98` and `99`. The destination range is the same length, but it starts at 50, so its two values are `50` and `51`. With this information, you know that seed number `98` corresponds to soil number `50` and that seed number `99` corresponds to soil number `51`.

The second line means that the source range starts at `50` and contains `48` values: `50`, `51`, ..., `96`, `97`. This corresponds to a destination range starting at `52` and also containing `48` values: `52`, `53`, ..., `98`, `99`. So, seed number `53` corresponds to soil number `55`.

Any source numbers that **aren't mapped** correspond to the **same** destination number. So, seed number `10` corresponds to soil number `10`.

So, the entire list of seed numbers and their corresponding soil numbers looks like this:

> seed  soil <br>
> 0     0 <br>
> 1     1 <br>
> ...   ... <br>
> 48    48 <br>
> 49    49 <br>
> 50    52 <br>
> 51    53 <br>
> ...   ... <br>
> 96    98 <br>
> 97    99 <br>
> 98    50 <br>
> 99    51 <br>

With this map, you can look up the soil number required for each initial seed number:

- Seed number `79` corresponds to soil number `81`.
- Seed number `14` corresponds to soil number `14`.
- Seed number `55` corresponds to soil number `57`.
- Seed number `13` corresponds to soil number `13`.

The gardener and his team want to get started as soon as possible, so they'd like to know the closest location that needs a seed. Using these maps, find **the lowest location number that corresponds to any of the initial seeds**. To do this, you'll need to convert each seed number through other categories until you can find its corresponding **location number**. In this example, the corresponding types are:

- Seed `79`, soil `81`, fertilizer `81`, water `81`, light `74`, temperature `78`, humidity `78`, **location** `82`.
- Seed `14`, soil `14`, fertilizer `53`, water `49`, light `42`, temperature `42`, humidity `43`, **location** `43`.
- Seed `55`, soil `57`, fertilizer `57`, water `53`, light `46`, temperature `82`, humidity `82`, **location** `86`.
- Seed `13`, soil `13`, fertilizer `52`, water `41`, light `34`, temperature `34`, humidity `35`, **location** `35`.

So, the lowest location number in this example is **`35`**.

**What is the lowest location number that corresponds to any of the initial seed numbers?**