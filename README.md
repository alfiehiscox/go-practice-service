# Go Practice Service

This is a simple go service that I'd like to try. 
The idea came from a [reddit post](https://www.reddit.com/r/golang/comments/ojjgum/what_are_some_good_project_ideas_for_a_go_beginner/)
and I like how simple yet core these little tasks are. The project is broken down into a series of sub-goals with the eventual project
being a simple service that is reachable over both HTTP and gRPC protocols, that mimics the functionality of a simple web service.

The sub projects are as such:
- [] Build a simple http server that returns 'Hello, world'
- [] Build a CLI that makes a HTTP request to said server
- [] Modify the server to return JSON (and perhaps add some logic)
- [] Modify the CLI to parse the JSON object (and perhaps add some logic)
- [] Rework both to support gRPC instead of HTTP

Each of the above goals has an Issue that holds dicussions on each topic. Each also has a branch that will be used to carry out the sub-goal. 
Any feedback is welcome, please just a Pull Request or comment suggested changes on the relavent Issue. 

Alfie 👾
