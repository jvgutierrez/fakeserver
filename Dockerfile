FROM alpine:latest
RUN adduser -D fakeserver
ADD fakeserver /home/fakeserver
USER fakeserver
CMD /home/fakeserver/fakeserver
