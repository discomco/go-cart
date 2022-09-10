# The GO-CART SDK

## OUTLINE

- [The GO-CART SDK](#the-go-cart-sdk)
  - [OUTLINE](#outline)
  - [DESCRIPTION](#description)
  - [MOTIVATIONS](#motivations)
  - [DEVELOPING GO-CART Applications](#developing-go-cart-applications)
    - [Setting up the local development environment](#setting-up-the-local-development-environment)
    - [Reference Implementation](#reference-implementation)
  - [SOME REFERENCES](#some-references)
  - [SERVERLESS DEPLOYMENT](#serverless-deployment)


## DESCRIPTION
This SDK provides a basis for developing modular applications for EventSourcing/CQRS in Golang, using Purposeful Architecture.

## MOTIVATIONS
- as a **Software Engineering Manager**
  - I require **a way to normalize the software manufacturing process** 
  - so that **software creation becomes a predictable, repeatable and scalable activity, both in velocity and quality**
- as a **Software Architect**
  - I want to be able **to provide a unified approach to solving non-functional requirements** 
  - and **empower software engineers with tools that minimize boilerplate code**
  - and **offer tools that normalize the software manufacturing process**
  - in order to **provide a consistent level of quality**
  - and **achieve a predictable manufacturing process**
- as a **Software Engineer**
   - I require **a set of base artifacts and patterns**
   - so I will **be able to focus on implementing business logic**
   - and **not have to worry about non-functional requirements**
   - and can **avoid to implement recurring boilerplate code**


## DEVELOPING GO-CART Applications



### Setting up the local development environment

For local development, we recommend the [compose-dev-stack](https://github.com/rgfaber/compose-dev-stack). 

<br>

usage:
```bash
git clone https://github.com/rgfaber/compose-dev-stack
./run-minimal.sh -d
```

### Reference Implementation

We are working on training material, but for now, 
please check out the *./robby* folder for a reference implementation. 




## SOME REFERENCES
- [Event Sourcing/CQRS - Greg Young](https://youtu.be/8JKjvY4etTY)
- [Event Storming - Alberto Brandolini](https://youtu.be/mLXQIYEwK24)
- [Microservices - Allard Buijze](https://youtu.be/jrbWIS7BH70)


## SERVERLESS DEPLOYMENT
One may notice that this architecture, due to its decoupled and modular nature, has many similarities with the concept of **Functions as a Service** (FaaS). One of the major issues that FaaS platforms struggle with is how to implement event sourced systems. This architecture does not pretend to have the solution for this problem, but **Behavior Composition** may offer a path forward. Should decomposing the Screaming Monolith into functional nanoservices become an option or a necessity one day, packaging each feature as a function and deploying it to a "serverless" infrastructure like OpenFaaS, should require minimal effort.