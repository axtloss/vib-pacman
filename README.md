
# vib-pacman

A [vib](https://github.com/vanilla-os/vib) plugin to allow using the `pacman(8)` package manager in vib recipes.

## Module structure

```yaml
Name: PacmanModule
Type: pacman
ExtraFlags:
    - "--overwrite=\"*\""
    - "--verbose"
Packages:
    - "bash"
    - "fish"
```

`ExtraFlags` can remain empty, the default flags passed to pacman are `-S --noconfirm`
