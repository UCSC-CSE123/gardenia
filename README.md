# Gardenia

![](https://user-images.githubusercontent.com/13544676/80291661-6dcd2000-8704-11ea-886f-756f41690d1e.png)

A web service that uses the magic of sunflowers to test beavertails

## Building and Running

### Dependencies

This code is reliant on the following projects:

1. [Beavertail](https://github.com/UCSC-CSE123/beavertail)

2. [Sunflower](https://github.com/UCSC-CSE123/sunflower)

You can get them by running via `git clone`:

```bash
$ git clone https://github.com/UCSC-CSE123/beavertail.git

$ git clone https://github.com/UCSC-CSE123/sunflower.git
```

Build and run both programs, documentation for `beavertail` and `sunflower` can be found [here](https://github.com/UCSC-CSE123/beavertail/blob/master/server/README.md) and [here](https://github.com/UCSC-CSE123/sunflower/blob/master/README.md), respectively.

### Writing the config.yaml

For `gardenia` to know where your `beavertail` and `sunflower` servers are running you must write a `config.yaml` file.

The following keys are required to run:

| Key               | Value                                                   |
| ----------------- | ------------------------------------------------------- |
| `Sunflower-Host`  | Where `sunflower` is being hosted on.                   |
| `Sunflower-Port`  | What port `sunflower` is listening on.                  |
| `Sunflower-Calls` | The number of API calls to make to `sunflower`.         |
| `GRPC-Host`       | Where the `beavertail` GRPC server is being hosted on.  |
| `GRPC-Port`       | What port the `beavertail` GRPC server is listening on. |

An example `config.yaml` may look like this:

```yaml
Sunflower-Host: localhost
Sunflower-Port: 8080
Sunflower-Calls: 100
GRPC-Host: localhost
GRPC-Port: 8081
```

### Running

After all that is done and the necessary servers are running run `gardenia` as follows:

```bash
# If [/path/to/config.yaml] is omitted
# gardenia looks for a config.yaml in the pwd.
$ ./gardenia [/path/to/config.yaml]
```

`gardenia` will output csv results to the screen and these can be saved using [UNIX redirects](https://en.wikipedia.org/wiki/Redirection_(computing)) like so:

```
$ ./gardenia > results.csv
```

Sample result:

```csv
Call Number,Duration,Acknowledgment
1,34.729423ms,OK
2,3.531697ms,OK
3,4.722095ms,OK
4,3.868302ms,OK
5,14.315772ms,OK
6,16.453ms,OK
7,19.95669ms,OK
8,12.337507ms,OK
9,11.975442ms,OK
10,18.665213ms,OK
11,19.375854ms,OK
12,11.927659ms,OK
```

## Credits

* Logo made using [LogoMakr](LogoMakr.com)
