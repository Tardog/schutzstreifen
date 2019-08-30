FROM gobuffalo/buffalo:v0.14.6

RUN go get -v github.com/gobuffalo/suite
RUN apt-get update && apt-get install \
    sassc --no-install-recommends -y \
    && rm -rf /var/lib/apt/lists/*
