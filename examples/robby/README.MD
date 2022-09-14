# Robby 

## Table of Contents

- [Robby](#robby)
  - [Table of Contents](#table-of-contents)
  - [Definition](#definition)
  - [Motivation](#motivation)
  - [General Description](#general-description)
  - [Discovery Artifacts](#discovery-artifacts)
  - [Getting Started.](#getting-started)
    - [What we got from discovery.](#what-we-got-from-discovery)

## Definition
Robby is a complete *go-cart application* and is part of the training material for the Purposeful Method,
It illustrates how the various types of application are manufactured and deployed.

## Motivation
- As the **Maintainer** of the go-cart SDK
- I require a **sample application**
- That I **can field-test the SDK in a non-critical environment**
- And **provide the community with a working sample**

## General Description
*Robby* is a turn-based game, that features a playing field on which robots can move around and do battle.

## Discovery Artifacts

- [Big Picture Diagram](https://miro.com/app/board/uXjVPa2FZfk=/)
  - [Execute Game](https://miro.com/app/board/uXjVPa2FWuY=/)

## Getting Started.

### Setting up the local development environment

For local development, we recommend the [compose-dev-stack](https://github.com/rgfaber/compose-dev-stack).
>Please Note: 
>some tests expect this environment to be up.


<br>

usage:
```bash
git clone https://github.com/rgfaber/compose-dev-stack
./run-minimal.sh -d
```




The Purposeful Method starts with a *Discovery Phase* which results in:
- A *Context Map* (aka 'The Big Picture')
- A number of *Contexts* with for each context:
  - The Behavioral Model (from Event Storming)
  - The Structural Model (from DDD)
A *Context*'s boundaries are defined by its behavior, in other words,
  - a Context supports a business process end-to-end and 
  - its schema should support just enough structures, not more,not less than is required to support this process autonomously.

### What we got from discovery.
So far, we have a rough understanding of what the system should look like here: 

- [Big Picture Diagram](https://miro.com/app/board/uXjVPa2FZfk=/)

And we have one Context that is considered ready for implementation: 

- [Execute Game](https://miro.com/app/board/uXjVPa2FWuY=/)








