# CrashDB

CrashDB is an ephemeral in-memory NoSQL database for the confident developer.

## Features

- CrashDB is webscale.
- Data is not persisted between process restarts.
- Using the wrong HTTP method will crash the database.
- Getting a key that isn't in the database will crash the database.
- Using malformed JSON will crash the database.
- Any system error will crash the database.

For privacy and security, CrashDB does not provide methods that reveal what
keys are currently stored in the database.

![That's a paddlin'](https://raw.githubusercontent.com/gunnihinn/crashdb/master/static/crashdb.jpg)

## Build

    $ make

CrashDB accepts no command-line options or arguments.

## Use

Health check:

    $ curl localhost:8080/ping

Get a value from the DB:

    $ curl -X GET --data '{"Key": "mykey"}' localhost:8080/

Put a value in the DB:

    $ curl -X POST --data '{"Key": "mykey", "Value": "myvalue"}' localhost:8080/

The keys must be strings, but CrashDB supports arbitrary values that can be
encoded in JSON:

    $ curl -X POST --data '{"Key": "mykey", "Value": {"foo": ["bar"]}}' localhost:8080/

## But why?

You have to learn how to improve and debug things somehow.

CrashDB has no shortage of problems. Discover them, fix them, and by improving
CrashDB, improve yourself as a developer.

The author recommends doing any or all of the following. Some of those points
are aimed at making any single instance of CrashDB less terrible, and some of
them are aimed at treating CrashDB like a distributed system.

- Run `sadness.sh`. Then Google [`delve`](https://github.com/derekparker/delve)
  to find out what you're looking at.  Become sad that
  [this](https://fntlnz.wtf/post/gopostmortem/) was not a part of your life
  until now. (The author notes that this only really works on Linux, as he
  sadly discovered while in the last throes of pretending that OSX was a decent
  operating system for doing development.)
- Improve the error messages to get a better idea of what went wrong. Collect
  those error messages somewhere so you can search them.
- Use an issue tracker - any issue tracker - to keep track of what you think is
  wrong with this thing. A folder with two folders called "open" and "closed"
  and some text files under version control is a perfectly good issue tracker
  for beginners.
- Decide whether to fix the race conditions (you did run `sadness.sh`, right?)
  by embracing channels or just using a fucking mutex like everyone in the
  world apparently does. Reflect on Erlang/Elixir and debate whether you would
  have this problem there.
- Find your own things to care about and improve CrashDB in those directions.
  Realize after weeks of work that some of those things were bad ideas and all
  that work needs to be thrown away. Be OK with that.
- Question some of the original design decisions. Is it really necessary to
  serialize and deserialize the values all the time? Does this add unacceptable
  overhead for intricate keys? Is it a potential DOS vector? Can you craft
  key/value pairs that result in arbitrarily bad performance?
- [Profile](https://blog.golang.org/profiling-go-programs) CrashDB. Get your
  own torture tests with sample queries that exhibit some of the problems you
  want to fix, or some of the performance characteristics of the program.
  Profile before and after fixes to see their impact.
- Make CrashDB persistent. Figure out how and where to store its data between
  restarts or crashes. Do you need to worry about data getting corrupted?
- Add monitoring. CrashDB is a service, so the
  [RED method](https://www.weave.works/blog/the-red-method-key-metrics-for-microservices-architecture/)
  seems appropriate. Decide whether to use "push" monitoring (à la Graphite) or
  "pull" monitoring (à la Prometheus).
- Support alternative serialization formats. JSON is fine for some things (the
  more you use numbers, the less fine it is), but consider binary formats like
  Protocol buffers. Use schemas to autogenerate a client for CrashDB for at
  least one language.
- Run CrashDB in Docker. Now run two of them at once and setup a master/slave
  replication chain. Direct writes to the master and reads to the slave. Do
  service discovery for the outside world. Run three instances at once and have
  the slaves elect a new master if the original one becomes sad. Find out why
  Zookeeper is a thing when you try to do any of this or query the cluster from
  the outside.

