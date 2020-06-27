FILENAME=results.txt
PROCS=$1
echo "PROCS = ${PROCS}" >> $FILENAME
(time GOMAXPROCS=$PROCS go run mandelbrot.go | grep real) &>> $FILENAME
