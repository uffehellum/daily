package merit

# CLI tool for Sigma engineers

The task is to create a simple CLI tool described in the "Task" section.
It can be solved in a language of your choice on your laptop/system (or on a laptop/system provided), it's easier to solve it in a scripting language.

## Context

In Merit we use Buildkite CI tool (https://buildkite.com) to create and manage different test and deployment pipelines, with the main one being the master deployment one which automatically deploys successful builds to staging.

When working on the local machine it’s often useful to quickly switch to the desired environment to inspect/test/debug certain features or run certain scripts without manually inspecting which commit is the latest functional on staging.

And so engineers require a simple CLI tool that will allow them to do exactly that.

At this point request a Buildkite overview to understand:
- pipelines
- builds
- jobs
- branches
- statuses
and how it corresponds to commit SHA in GitHub

## Task
Create a CLI tool that for the designated repository will do two things:
1. Print to the console git hash that corresponds to the latest successful Buildkite build on the master branch

2. Switch your git repo configuration to the commit SHA retrieved in the previous step and run a command, for simplicity sake the command to run is just getting top three commits so that the first one corresponds to the checked out commit

Note: in real usage once on the proper configuration we run scripts and/or Scala REPL that interact with current data structures and services (like database services)

It is of the first importance to complete these two conditions.

Extra/bonus: as you work through the implementation consider if your code is readable, maintainable, scoped well with relations to the context and can be expanded further.

## Useful info
The designated repository (can ignore installation instructions there):
https://github.com/LiuVII/pantsbuild_test

- Buildkite API token
```
cd4f8cdb69fdac86ab4cdb7efedb564a0cbc26c8
```
- Organization name in Buildkite
```
sigma
```
- Pipeline name in Buildkite
```
pantsbuild-basics
```
