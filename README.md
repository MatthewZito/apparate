# apparate

## Table of Contents

- [Installation](#install)
- [Usage](#usage)

## <a name="install"></a> Installation

1. Clone:

```shell
git clone https://github.com/MatthewZito/apparate.git && cd apparate
```

2. Compile the app:

```shell
make build
```

3. Add the wrapper to your PATH:

```shell
echo "source $PWD/apparate.bash" >> ~./bashrc # replace w/ your shell config if different
```

One-liner:

```shell
git clone https://github.com/MatthewZito/apparate.git && \
cd apparate && \
make build && \
echo "source $PWD/apparate.bash" >> ~./bashrc
```

## <a name="usage"></a> Usage

Add a warp-point to the current path:

```shell
apparate add <alias>
```

Apparate to an existing warp-point:

```shell
apparate goto <alias>
```

Remove an existing warp-point:

```shell
apparate remove <alias>
```
