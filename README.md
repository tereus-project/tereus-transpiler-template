# Tereus LanguageA to LanguageB transpiler

Converts LanguageA to LanguageB.

## Command-line usage

You can use the following command to convert a LanguageA file to LanguageB:

```shell
go run . <input.la>
```

## Worker usage

The transpiler is designed to listen on a queue for LanguageA files to convert.

First, you need to copy the `.env.example` file to `.env` and fill in the values.

The Kafka endpoint and S3 bucket should be shared with the [Tereus API](https://github.com/tereus-project/tereus-api).

You can then start the worker with:

```shell
➜  go run .
INFO[0000] Connecting to Kafka...
INFO[0000] Connecting to MinIO...
INFO[0000] Starting transpiler job listener...
DEBU[0006] Job '48549e3e-a181-49f5-b204-0fc56df3e319' started
DEBU[0006] Downloading job files...
DEBU[0006] Downloading file 'main.la'
DEBU[0006] Remixing file 'main.la'
DEBU[0006] Uploading file 'main.lb'
DEBU[0007] Job '48549e3e-a181-49f5-b204-0fc56df3e319' completed
```

The worker will get transpilation submissions from a queue and update the status and result in another one.
