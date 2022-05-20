# NGINX
NGINX is open source software for web serving, reverse proxying, caching, load balancing, media streaming, and more. It started out as a web server designed for maximum performance and stability. In addition to its HTTP server capabilities, NGINX can also function as a proxy server for email (IMAP, POP3, and SMTP) and a reverse proxy and load balancer for HTTP, TCP, and UDP servers.

## Table of Contents
- [Command Line Options](#command-line-options)
- [Getting Started](#getting-started)
- [Contexts](#contexts)
- [Serving Static Content](#serving-static-content)
    - [Single Domain](#single-domain)
    - [Multiple Domain](#multiple-domain)
- [Best Practice](#best-practice)

## Command Line Options
```sh
nginx version: nginx/1.21.6
Usage: nginx [-?hvVtTq] [-s signal] [-p prefix]
             [-e filename] [-c filename] [-g directives]

Options:
  -?,-h         : this help
  -v            : show version and exit
  -V            : show version and configure options then exit
  -t            : test configuration and exit
  -T            : test configuration, dump it and exit
  -q            : suppress non-error messages during configuration testing
  -s signal     : send signal to a master process: stop, quit, reopen, reload
  -p prefix     : set prefix path (default: /etc/nginx/)
  -e filename   : set error log file (default: /var/log/nginx/error.log)
  -c filename   : set configuration file (default: /etc/nginx/nginx.conf)
  -g directives : set global directives out of configuration file

```

## Getting Started
change directory to `nginx/1-getting-started`

```sh
# to run
docker compose -f compose.yaml up -d

# to remove
docker compose -f compose.yaml down
```

## Contexts
```conf
# The most general context is the “main” or “global” context. 
# It is the only context that is not contained within the typical context blocks.
# Examples:

user                www www;
worker_processes    2;
pid                 /var/run/nginx.pid;
error_log           /var/log/nginx.error_log    info;

# The “events” context is contained within the “main” context. 
# It is used to set global options that affect how Nginx handles connections at a general level. 
# There can only be a single events context defined within the Nginx configuration.
events {
    worker_connections 768;   
    multi_accept on; 
}

# The main function of the mail context is to provide an area for configuring a mail proxying solution on the server. 
# Nginx has the ability to redirect authentication requests to an external authentication server. 
# It can then provide access to POP3 and IMAP mail servers for serving the actual mail data. 
# The mail context can also be configured to connect to an SMTP Relayhost if desired.
mail {
    server_name mail.example.com;  
    auth_http   localhost:9000/cgi-bin/nginxauth.cgi;  
    proxy_pass_error_message on;  
}

# This context is to define how the program will handle HTTP or HTTPS connections.
# The http context is a sibling of the events context, so they should be listed side-by-side, rather than nested. 
# They both are children of the main context
http {
    ...

    # The upstream context is used to configure and define an upstream server. 
    # This context is allowed to define a pool of backend servers that Nginx can proxy the used when request. 
    # This context is usually we are configuring proxies of various types.
    # Upstream context enables Nginx to perform load balancing while proxying the request. 
    # This context is defined inside the HTTP context and outside any server context.
    # The upstream context is referenced by name within server or location blocks. 
    # And then pass the requests of a certain type to the pool of defined servers. 
    # The upstream will then use an algorithm (by default round-robin) to determine which specific server need to be used to handle the request.
    upstream <custom_name> {  
        ...  
    }  

    # The server context is declared within the http context. 
    # The server context is used to define the Nginx virtual host settings. 
    # There can be multiple server contexts inside the HTTP context.
    server {
        ...

        # Location contexts define directives to handle the request of the client. 
        # When any request for resource arrives at Nginx, 
        # it will try to match the URI (Uniform Resource Identifier) to one of the locations and handle it accordingly.
        location <modifier> path{
            ...
        }

        ...
    }

    ...
}
```

change directory to `nginx/2-contexts`

```sh
# to run
docker compose -f compose.yaml up -d

# to remove
docker compose -f compose.yaml down
```

## Serving Static Content
In this section, we will discuss how to configure the Nginx Plus and Nginx open source to serve static content.

### Single Domain
it will serve static content on one server and domain

change directory to `nginx/3-3-serving-static-content/1-single-domain`

```sh
# to run
docker compose up -d

# to remove
docker compose down
```

### Multiple Domain
it will serve static content on one server and multiple domain

change directory to `nginx/3-3-serving-static-content/2-multiple-domain`

before run command below you have to add alias domain in your hosts file, example 
```hosts
12.0.0.1     domain1.localhost
12.0.0.1     domain2.localhost
```

```sh
# to run
docker compose up -d

# to remove
docker compose down
```

## Best Practice
- use "`nginx -s reload`" to reload configuration