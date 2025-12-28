#! /usr/env/bin bash

# Copy to clipboard the arguments that can be found in the cmd file in order to help build out the help function
grep "argument ==" cmd/cmd.go \
  | sed -E 's/.*if[[:space:]]*argument[[:space:]]*==[[:space:]]*"([^"]*)".*/aphrodite.PrintInfo("\1")/' \
  | pbcopy
