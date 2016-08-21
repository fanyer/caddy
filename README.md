<a href="https://caddyserver.com"><img src="https://caddyserver.com/resources/images/caddy-lower.png" alt="Caddy" width="350"></a>

[![community](https://img.shields.io/badge/community-forum-ff69b4.svg?style=flat-square)](https://forum.caddyserver.com) [![twitter](https://img.shields.io/badge/twitter-@caddyserver-55acee.svg?style=flat-square)](https://twitter.com/caddyserver) [![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/mholt/caddy) [![Linux Build Status](https://img.shields.io/travis/mholt/caddy.svg?style=flat-square&label=linux+build)](https://travis-ci.org/mholt/caddy) [![Windows Build Status](https://img.shields.io/appveyor/ci/mholt/caddy.svg?style=flat-square&label=windows+build)](https://ci.appveyor.com/project/mholt/caddy)

Caddy is a general-purpose web server for Windows, Mac, Linux, BSD, and
[Android](https://github.com/mholt/caddy/wiki/Running-Caddy-on-Android). It is
a capable but easier alternative to other popular web servers.


----------
##Extension : caddy-ace
Add [Ace](https://github.com/yosssi/ace) template engine directive &amp; plugin for caddy server

Middleware for [Caddy](https://caddyserver.com).



### Usage
Add an ace key into caddyfile as below

```
ace  {
    path /example
}
```
* **path** whose value is the relative path where you store your ace sourcefile 


Then just type 
```
caddy 
```
and visit [http://localhost:2015/example/](http://localhost:2015/example/)

----------------------

For more caddy server docs , please click [here](https://caddyserver.com/) for caddy official guides.

Also you can visit the official [repo](https://github.com/mholt/caddy)