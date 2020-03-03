all: levyc.png broc.png koch.png cesaro.png

levyc.png: curve levyc.load
	./curve 15 .5 .5 > levyc.pts
	gnuplot < levyc.load

broc.png: curve broc.load
	./curve -C 17 .6 .37 > broc.pts
	gnuplot < broc.load

koch.png: curve koch.load
	./curve -C 17 .5 .2886751 > koch.pts
	gnuplot < koch.load

cesaro.png: curve cesaro.load
	./curve 17 .3 .3 > cesaro.pts
	gnuplot < cesaro.load

czoom.pts: curve
	./curve 20 .3 .3 > czoom.pts

czoom.png: czoom.pts czoom.load
	gnuplot < czoom.load

curve: curve.go
	go build curve.go

clean:
	-rm -rf curve
	-rm -rf levyc.pts levyc.png
	-rm -rf broc.pts broc.png
	-rm -rf cesaro.pts cesaro.png
	-rm -rf czoom.pts czoom.png
	-rm -rf koch.pts koch.png
