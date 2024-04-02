# Go/HTMX practice project

I have been working on different technologies for over 20 years. PHP, Javascript, Java, Clojure and most recent, Python.

Even though I did some Java, it was just for a couple of months and I have been willing to use a statically-typed 
language for some time now. However, it's always difficult to find a job when you don't have experience with the 
language, even if it's quite new.

So I decided that instead of just reading some books, tutorials, videos, or other sources, I will build a side project
using Go.

#### TL DR: I'm creating a project to practice GO and HTMX and document what I learned. Don't use it in real life.

### Why Go?

I tried solving a simple program using different languages (Python, Java, Rust and Go) and I found a very 
good balance between learning curve, ease of coding, documentation and an active community in Go.

### Why HTMX? 
I happen to have a very strong opinion regarding the javascript and frontend direction for the past 6-7
years. I have found over-engineered solutions to very simple problems that we never had in the JQuery era. I've tried
React and Vue and just got very frustrated on how complicated things are and how little benefit you get from using those
tools other than getting hired in some full-stack position where you will kind of understand the front end and noticing
that the rest of the team also kind of understand how it works.

I heard of HTMX a couple of months ago and I loved the idea as soon as I read the documentation. Just go back to what
HTTP and HTML were meant from the beginning and build everything on the server. Use whatever backend language you prefer
and actually enjoy programming it.

If you have never heard of it, I recommend that you read this
[Hypermedia document](https://hypermedia.systems/introduction/).

### Why this project?

It's often difficult to choose a portfolio project that is not too small, or too big. Sure, a hospital can be VERY 
large. However, it can be just a part of it and not the whole thing.

I also happen to be working on a project for a small Hospital that has very specific needs, so I can bring most of it
to his project and do both at the same time. I will be able to play and experiment here while solving a real life
problem.


## Disclaimer

This is not an application intended for real life use. It's mostly a therapy place for me that I can share with the
world as I LOVE programming, and it relaxes me to learn. I have never done something like what I have in mind for this
project.

## So, what is this about?

My idea is to have some kind of blog. At this point is to have log files when I push changes explaining what
I did along with the reasoning around my decision to do what I did. Maybe it can help someone to decide on a pattern or
tool to use.

If there is interest from a single person, I will add something to have discussions on my decisions as this will be a 
learning experience for me. I will investigate if GitHub has something for this.

I will probably do very big changes from time to time as I learn about Go and HTMX specially in the project structure.

## What other tools will you use?

What I'm familiar with.

- Postgres
- Docker
- Bootstrap
- AWS CDK (initially. I might look into Terraform)

And I will add things as they are needed.

## What have you done so far?

I'll try to commit the code that I already have in parts. Some of the tools that I use were not really chosen by me, but
by someone in some mix of tutorials that I read or view. So far I have:

- Base models for addresses
- Migrations using [Fizz](https://github.com/gobuffalo/fizz)
- Docker compose file for development with Postgres

These will probably change over time as I learn about different tools.

## Why is this repository named `core`?

With the little experience that I have with go at this moment, I want to emulate the tools that I'm familiar with.
Something that I like about frameworks like Laravel, Django and Spring is the option to execute commands using all the
service related tools from the application.

I like the minimalism that Go has, so I want to use as minimal tools as possible. My intention is tu have this repo as
a base module to be used in a repo with the HTMX implementation and another repo with a CLI program for different
purposes.
