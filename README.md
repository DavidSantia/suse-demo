# suse-demo
SUSE Linux Enterprise Server 12 SP3 (x86_64) New Relic agent demo

## How to build docker image
The build script compiles a small Go application, and places it on a SUSE linux container, along with the New Relic linux agent.
```sh
./build.sh
```

## How to configure monitoring
First copy the template to newrelic.env
```sh
cp newrelic.template.env newrelic.env
```

Then edit `newrelic.env` and set the license key on the first line:
```
NEW_RELIC_LICENSE_KEY=YOUR_LICENSE_KEY_HERE
```

The run script will then automatically use this key.  This will enable monitoring for the Linux agent, and the Go server.
```sh
./run.sh
```

