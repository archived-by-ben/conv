#Readme
conv is a terminal tool to convert everyday units of measurements.

###Downloads
* [Linux](https://drive.google.com/uc?export=download&id=0B-JKAyjDRrBtM1ZtY1p2SVd0a3c)
* [OS X](https://drive.google.com/uc?export=download&id=0B-JKAyjDRrBtVVdBcnpBVkhQSkE)
* [Windows](https://drive.google.com/uc?export=download&id=0B-JKAyjDRrBtMFd3Unc1eWZiam8)
* [Linux for ARM (Raspberry Pi, Arduino, etc.)](https://drive.google.com/uc?export=download&id=0B-JKAyjDRrBtUjk5Q2xrU25xNEE)

BSD, Plan 9, ARM and 32-bit binaries can be found at [Google Drive](https://drive.google.com/folderview?id=0B-JKAyjDRrBtfmVWa3RPbEFEdjRPdDMzTi1LcmdYMzhBM05NWjB5THlQek5zajIyclR5YTg&usp=drive_web#list).

###Instructions
conv is a self-contained, portable program with no dependencies. Simply download the appropriate archive for your operating system. Unzip it to a directory of your choosing and run it from your terminal or command prompt.

```shell
Usage:  conv (measurement)(unit)
        conv -h | --help
        conv -v | --version

Example:    conv 100f
            100°F is about 37.8°C (Celsius)

The permitted units are:

    bbl     » bbl   petroleum barrel
    bps     » bps   bits per second
    bs      » B/s   bytes per second
    c       » °C    Celsius
    cm      » cm    centimetres
    ct      » ct    carats
    cum     » m³    cubic metre
    f       » °F    Fahrenheit
    ft      » ′     feet
    g       » g     grams
    guk     » gal   imperial gallon
    gus     » gal   us gallon
    hp      » hp    horsepower
    in      » ″     inches
    kbps    » kib/s kilobit per second
    kbs     » kB/s  kilobyte per second
    kg      » kg    kilograms
    km      » km    kilometres
    kmh     » km/h  kilometres per hour
    kn      » kn    knots
    l       » L     litres
    lb      » ℔     pounds
    m       » m     metres
    mbps    » Mib/s megabit per second
    mbs     » MB/s  megabyte per second
    mi      » mi    miles
    mph     » mph   miles per hour
    mps     » m/s   metres per second
    nm      » NM    nautical miles
    oz      » oz    ounces
    st      » st    stone
    w       » W     watts
    yd      » yd    yards
```

###Future additions
* Add support for fractions to use with select Imperial units; feet + inches; stone + pounds.
* Add measurements.

##### Areas
* acre; hectare; square km; square mile; square yard

##### Liquid volumes
* ounce; pint; quart; pony

###Licence
Copyright © 2015 Ben Garrett. [The MIT License (MIT)](http://choosealicense.com/licenses/mit/)
