# Chain of responsibility

Use this to execute chain of actions.

1. Action interface

+ execute()
+ set next()

2. Action interface implemented by chain of actions

+ Execute() is called to execute the next execute()

3. Run it in main.go
