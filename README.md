# go-donut-challenge
[![Build Status](https://travis-ci.org/florinutz/go-donut-challenge.svg?branch=master)](https://travis-ci.org/florinutz/go-donut-challenge)
[![codecov](https://codecov.io/gh/florinutz/go-donut-challenge/branch/master/graph/badge.svg)](https://codecov.io/gh/florinutz/go-donut-challenge)

If you find the project structure a bit weird, 
pls check [this](https://fosdem.org/2019/schedule/event/designingcli/) out.

The most complete and well behaved part (read "covered by tests + nice display") 
is the `ticker` command, but everything should work nonetheless.

`make binary` will build for your platform and place a `./bin` symlink 
in the project folder for you to play with. (`./bin help`)

But do run `make` and check out all the targets in the Makefile.

I used the existing 3rd party coinbasepro [package](https://github.com/preichenberger/go-coinbasepro) 
because it provided some stuff that I didn't want to waste time doing myself.

Be sure to have the proper sandbox credentials 
in the required [env vars](https://github.com/preichenberger/go-coinbasepro#setup) 
when placing an order.