## Automatic failover for CloudFlare

This Go program checks a number of endpoints and in case the primary server fails it will automatically update CloudFlare DNS destination of given domain to a backup server.

You will need to create your own checks in checks/ folder and configure your own CloudFlare access in cloudflare.conf

### Dependencies

https://git.rtek.se/rasmus/RTCheck

### Build

[![Build Status](https://travis-ci.org/rasmusj-se/cloudflare-failover.svg?branch=master)](https://travis-ci.org/rasmusj-se/cloudflare-failover)

Latest build (for linux x86) is available at https://build.rtek.se/cloudflare-failover and in GitHub releases.

### Usage

Download clouflare.conf and update with API key obtained from cloudflare.
Setup your checks in checks/ folder, look at the examples on GitHub for information on how to configure.
