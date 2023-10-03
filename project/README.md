## Intro
This is a series of labs in which you'll build a somewhat fault-tolerant in-memory key-value database in Go,
with journaling to support persistence (like [Redis](https://en.wikipedia.org/wiki/Redis)).

In this project ([inspired](https://pdos.csail.mit.edu/6.824/labs/lab-raft.html) by two labs from
[MIT](https://pdos.csail.mit.edu/6.824/labs/lab-kvraft.html) and [Berkeley](https://inst.eecs.berkeley.edu/~cs162/sp23/static/hw/lab-grpc-rs)),
you'll deploy [Ceph](https://en.wikipedia.org/wiki/Ceph_(software)), a replicated storage platform, on Docker.
Next, you'll build a key/value service on top of Ceph.
Then you will “cache” your service for higher performance.

Keep in mind that the most challenging part of this lab may not be implementing your solution, but debugging it.
To help address this challenge, you may wish to spend time thinking about how to make your implementation more easily debuggable.

This lab is due in two to four parts. You should submit each part on the corresponding due date (optional).

### How to begin
Examine the `api` dir, Then you'll be good to go.
We suggest that you implement the back-end `server` first.
The front-end load-test has no constraints on the programming language of your choice (JS is recommended).

### Roadmap
1. [ ] Implement `DBServer`, passing all the tests.
2. [ ] Write a load test (`DBClient`) and benchmark your app (ideally with k6).
3. [ ] Write a `Dockerfile` and build a containerized image, run it, and test it with **Postman**.
4. [ ] Implement transaction journaling for persistence.
5. [ ] Use [Ceph](https://hub.docker.com/r/ceph/ceph) along with the container you built.
6. [ ] Utilize caching and pipelining techniques to prevent the journaling procedure from slowing your program down.
7. [ ] (optional) deploy the whole package in kubernetes.
8. [ ] Stonks.

## Pair programming guide
Working well with your group is crucial for success. One way to facilitate group work is by pair programming. Pair programming is a way to program collaboratively with a partner. It's a great approach that is used by many companies in the tech industry.

In pair programming, partners are working together at the same time. One partner is the "driver," who actually types the code. The other partner is the "navigator," who observes, asks questions, suggests solutions, and thinks about slightly longer-term strategies. We will continue to discuss pair programming in the canonical sense (with two people), but note that the idea can easily extend to a larger group of people, with one driver and multiple navigators.

It is crucial for partners to switch roles throughout an assignment, either every 10-20 minutes or alternating on each subtask/problem.

#### Working with a partner
We sometimes see feedback like:

> My partner didn't do any work, or didn't do their share of the work, or didn't communicate or meet with me, etc. What can I do?

Or:

> My partner did too much! They hogged the keyboard, or they did the whole assignment without waiting for me, or they didn't communicate with me, etc. I feel that I didn't get a real chance to help in solving the assignment. What can I do?

Have you tried speaking to your partner about your expectations beforehand? Many group issues can be resolved by better communication and setting expectations. If talking to your partner does not resolve the situation, speak to course staff and explain the details of what has happened. We will try to help you resolve the issue.

## Additional resources
- National Suicide Prevention Hotline: 021-123