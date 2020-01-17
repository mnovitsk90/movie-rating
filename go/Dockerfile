# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app
# Set an env var that matches your github repo name
ENV SRC_DIR=/movie-rating/go
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build movierating.go; cp movierating.exe /app/
ENTRYPOINT ["./movierating.exe"]