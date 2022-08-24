# Livesup Cli

```
livesup-cli get teams                              
┌──────────────────────────────────────┬────────┬─────────────┐
│ #                                    │ NAME   │ DESCRIPTION │
├──────────────────────────────────────┼────────┼─────────────┤
│ ba4405a8-ccf2-4e22-909c-9934fd6f19e9 │ vmware │             │
│ 66fb6b76-ded2-454c-9909-cd77030ecea0 │ what   │             │
│ 8bce96ca-e620-42b2-a0dc-a9febbc7d371 │ ubu    │             │
└──────────────────────────────────────┴────────┴─────────────┘
```

## Testing
You can run the full test suite with go test -v ./.... 

If you'd like to run an individual test, you can run the following

```bash
# Where TestNewTeamService is the name of your test function
$ go test -v ./... -run TestNewTeamService
```

References
* [Go Unit Tests: Tips from the Trenches](https://www.red-gate.com/simple-talk/devops/testing/go-unit-tests-tips-from-the-trenches/#Go_Unit_Tests)
