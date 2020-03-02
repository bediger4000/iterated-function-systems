levyc.png: curve levyc.load
	./curve 15 .5 .5 > levyc.pts
	gnuplot < levyc.load

broc.png: curve broc.load
	./curve -C 17 .6 .37 > broc.pts
	gnuplot < broc.load

clean:
	-rm -rf levyc.pts levyc.png
	-rm -rf broc.pts broc.png
