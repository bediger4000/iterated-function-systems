# Iterated Function Systems Fractals

More fractals via Iterated Function Systems

* https://en.wikipedia.org/wiki/De_Rham_curve
* http://linas.org/math/de_Rham.pdf

## Build

You need a [Go](https://golang.org/) compiler to compile the [source code](curve.go).
You need [gnuplot](http://www.gnuplot.info/) to create the PNG images.
After that, it's simple:

   $ make all

See the contents of `makefile` for tips on how to invoke and stuff.

### Invoking the program

The program is named `curve`.
It takes 3 command line parameters:

* Number of generations
* real part of complex number A
* imaginary part of complex number A

The number of generations is basically the number of pixels in the final image.
Since the IFS ends up being two functions

    d0(z) = Az
    d1(z) = A + (1 - A)z

Each initial point (1 + 0i and 0 + 0i) ends up yielding 2 subsequent points.
Each "generation" ends up yielding 2 more complex numbers.
I have `curve` not invoking `d0()` and `d1()` on the 2nd complex number
if it's identical to the first one,
and the functions do yield some identical points,
so the final output isn't a multiple of 2.

The real and imaginary parts of A cause different curves.

* .5 + .2886751i (along with the -C flag) make a Koch snowflake
* .5 + .5i make a Levy C Dragon
* .3 + .3i make the cool wiggly Cesaro curve. Spirograph go home!

![Cesaro Curve](c.png?raw=true)

Example Cesaro Curve
