FROM scratch

ADD canary /

ENV PORT=80

CMD ["/canary"]