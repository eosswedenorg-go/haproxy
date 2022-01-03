# Haproxy

Go module for [Haproxy](https://www.haproxy.com)

### Install package

``` bash
go get -u github.com/eosswedenorg-go/haproxy@latest
```

### Constants

```c
type HealthCheckStatus string
```
Enum for haproxy's agent-check protocol. Read more at the
[official documentation](https://cbonte.github.io/haproxy-dconv/1.7/configuration.html#5.2-agent-check)


```
HealthCheckUp
```
sets back the server's operating state as UP if health checks also report that the service is accessible.

```
HealthCheckDown
HealthCheckFailed
HealthCheckStopped
```
optionally followed by a description string after a sharp ('#').
All of these mark the server's operating state as `DOWN`, but since
the word itself is reported on the haproxy stats page, the difference
allows an administrator to know if the situation was expected or not:

the service may intentionally be stopped, may appear up but fail some
validity tests, or may be seen as down (eg: missing process, or port
not responding)

```
HealthCheckReady
```
This will turn the server's administrative state to the `READY` mode,
thus cancelling any `DRAIN` or `MAINT` state

```
HealthCheckDrain
```
This will turn the server's administrative state to the `DRAIN` mode,
thus it will not accept any new connections other than those that are
accepted via persistence.

```
HealthCheckMaint
```
This will turn the server's administrative state to the `MAINT` mode,
thus it will not accept any new connections at all, and healthchecks
will be stopped.


### Author

Henrik Hautakoski - [Sw/eden](https://eossweden.org/) - [henrik@eossweden.org](mailto:henrik@eossweden.org)
