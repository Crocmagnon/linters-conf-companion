# linters-conf-companion

This repo is a companion to a conference about linters.

It's intended to be walked commit by commit, starting with the first one.

First, configure git aliases:

```shell
git config --global alias.next '!git checkout $(git rev-list HEAD..${1:-main} --ancestry-path | tail -1 || echo HEAD)'
git config --global alias.prev 'checkout HEAD^'
git config --global alias.first '!git checkout $(git rev-list --max-parents=0 HEAD)'
```

Then, walk the repo:
```shell
git first
git checkout demo-1
git checkout demo-2
git next
git prev
# ...
```

## Conference list
* [Gophercamp 2025](https://gophercamp.cz/)
