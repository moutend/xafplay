# xafplay

Extended afplay (macOS built-in audio file play) command.

## Install

```console
go install github.com/moutend/xafplay/cmd/xafplay@latest
```

## Usage

```console
xafplay music1.wav music2.wav ... musicN.wav
```

## Why don't you use `xargs`?

After starting playback, you can send the SIGINT signal to control the playback. When you hit the Ctrl-C twice within 1 second, the playback will stop. Otherwise, the next playback will start.

## LICENSE

MIT
