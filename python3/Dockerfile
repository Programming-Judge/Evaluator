FROM python:3.9-slim-bullseye
RUN useradd --create-home --shell /bin/bash execution_user
COPY evaluate.sh /home/execution_user/
RUN chmod +x /home/execution_user/evaluate.sh
USER execution_user
WORKDIR /home/execution_user
ENTRYPOINT [ "./evaluate.sh" ]