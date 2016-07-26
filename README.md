# mantra

A simple cron-like scheduler for a single command

## Motivation

Running a process in a Docker container on a fixed schedule is fiddly. Most solutions seem to involve installing `cron`, then installing a `crontab` and tailing `cron`'s output. Yuck.

`mantra` is more like [`at`](https://en.wikipedia.org/wiki/At_(Unix)), in that it allows you to specify the schedule as an argument in a `cron` format, along with the executable name and any arguments. It will then run that command, and only that command, on the schedule provided. STDOUT and STDERR are shared, making it easier to capture the program output.

## Usage

Execute `my_prog` with some arguments at 1am every day:

```
mantra "0 1 * * * *" my_prog arg1 arg2
```

## In Docker

In your `Dockerfile`, download the latest binary, for example:

```
RUN curl -o /usr/local/bin/mantra -L https://github.com/pugnascotia/mantra/releases/download/0.0.1/mantra && \
    chmod +x /usr/local/bin/mantra
```

Then call it:

```
CMD mantra schedule my_prog arg1 arg2 ...
```

## See Also

Uses a [slightly modified](https://github.com/pugnascotia/cron) version of the excellent [robfig/cron](https://github.com/robfig/cron) package.

