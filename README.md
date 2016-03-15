## Automatic failover for CloudFlare

This Go program checks a number of endpoints and in case the primary server fails it will automatically update CloudFlare DNS destination of given domain to a backup server.

You will need to create your own checks in checks/ folder and configure your own CloudFlare access in cloudflare.conf

### Dependencies

https://git.rtek.se/rasmus/RTCheck

### Build

[![Build Status](https://travis-ci.org/rasmusj-se/cloudflare-failover.svg?branch=master)](https://travis-ci.org/rasmusj-se/cloudflare-failover)

Latest build is available at https://build.rtek.se/cloudflare-failover and in GitHub releases.
