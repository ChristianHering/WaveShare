WaveShare
===========

This repository is a functional library for 7.5" WaveShare ePaper displays.

It provides:

  * A way of displaying content on WaveShare ePaper displays
  * An example of how rPi GPIO/I2C can be used in go

Table of Contents:

  * [About](#about)
  * [Compiling from Source](#compiling-from-source)
  * [Contributing](#contributing)
  * [License](#license)

About
-----

I originally wrote this so I didn't have to use [WaveShare's C lib](https://github.com/waveshare/e-Paper) in my [thermostat project](https://github.com/ChristianHering) in an attempt to avoid using cgo.

This Repo was originally made around [WaveShare's 7.5" eInk Display](www.waveshare.com/7.5inch-e-paper-hat.htm), but could be adapted to other resolution displays that use the same protocol. Feel free to take a look at the documentation I used:

  * [WaveShare Display Documentation](7.5inch_e-Paper_V2_Specification.pdf) - [Source](https://www.waveshare.com/w/upload/6/60/7.5inch_e-Paper_V2_Specification.pdf)

Compiling from Source
------------

If you're looking to compile from source, you'll need the following:

  * [Go](https://golang.org) installed and [configured](https://golang.org/doc/install)
  * A [Raspberry Pi](https://www.raspberrypi.org/), [Arduino](https://www.arduino.cc/), or another I2C controller
  * A [WaveShare 7.5" E-Ink Display](www.waveshare.com/7.5inch-e-paper-hat.htm) or another one that uses the same protocol
  * Some patience, as this repository isn't really intended for public use. (There may be unnoticed bugs/random rough edges)

Contributing
------------

Contributions are always welcome. If you're interested in contributing, send me an email or submit a PR.

The following things need work:

  * Partial updating needs to be added
  * Timings should be tightened up and unnecessary sleeping should be removed
  * The project should be written in a procedural way... the Go way!
  * Other, less important features need to be implemented

License
-------

This project is currently licensed under GPLv3. This means you may use our source for your own project, so long as it remains open source and is licensed under GPLv3.

Please refer to the [license](/LICENSE) file for more information.
