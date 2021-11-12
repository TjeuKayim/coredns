# tunnel

## Name

*tunnel* - returns your resolver's local IP address, port and transport.

## Description

The *tunnel* plugin is not really that useful, but can be used for testing DNS tunneling performance.
Only alters CNAME queries, other types are forwarded to the next plugins.

## Syntax

~~~ txt
tunnel
~~~

## Examples

Start a server on the default port and load the *tunnel* plugin.

~~~ corefile
example.org {
    tunnel
}
~~~

TODO: fix the example

When queried for "example.org CNAME", CoreDNS will respond with:

~~~ txt
;; QUESTION SECTION:
;example.org.   IN       CNAME

;; ADDITIONAL SECTION:
example.org.            0       IN      CNAME       10.240.0.1
~~~
