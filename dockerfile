# escape=`

FROM ubuntu:latest

RUN apt-get update

RUN apt-get install g++ -y

RUN mkdir submissions

COPY evaluate.sh .

RUN chmod +x evaluate.sh

ENTRYPOINT ["./evaluate.sh"] 