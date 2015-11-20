FROM alpine:latest
RUN adduser -D fakeserver
ADD fakeserver /home/fakeserver
USER fakeserver
EXPOSE 8080
CMD /home/fakeserver/fakeserver
