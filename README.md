# Redesigned Chainsaw

Advent of Code 2021

Some exercises in programming for 2021

## Day 00
Overzealous attempts at trying to relearn idiomatic Go by trying to define reusable interfaces.

## Day 01

About four hours of struggling with trying to do things properly (and doing it wrong), followed by around thirty minutes
beating [zap](https://github.com/uber-go/zap) into submission interspersed with about fifteen minutes of actually
solving the puzzle.

There's a lesson to be learned there.

## Day 02

Nothing much to think about. Could probably run faster with some tweaking of the `bufio.Scanner`, but that's really not
worth any effort. Also, tests weren't really necessary here.

I've learned yesterday's lesson, for now.

## Day 03

The mention of bits was a huge red herring and flipping strings was the easiest. Not having a reduce function also was a
bit of a pain in the ass and probably something worth learning. Having to drop slices ends up triggering the garbage
collector quite often.

## Day 04

Coming up with the data structures for easy access took quite a bit of figuring out. As well as how the `Reader` in
the `encoding/csv` package works. Also learned that you can't change `bufio.Split` functions after using it to
`bufio.Scan()`. Still haven't figured out if it's possible to get a new scanner at the position where you left the
other, but this worked.

Also, realising that more boards can win per roun/d.
