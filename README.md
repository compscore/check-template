# Check Template

This is the Compscore check template, thus this is **not** a working check.

All new checks for compscore should use this repository as a template and build checks from there.

## Check Format

```yaml
- name:
  release:
    org: compscore
    repo: check-template
    tag: latest
  credentials:
    username:
    password:
  target:
  command:
  expectedOutput:
  weight:
  options:
```

## Parameters

|    parameter     |          path           |   type   | default  | required | description                                     |
| :--------------: | :---------------------: | :------: | :------: | :------: | :---------------------------------------------- |
|      `name`      |         `.name`         | `string` |   `""`   |  `true`  | `name of check (must be unique)`                |
|      `org`       |     `.release.org`      | `string` |   `""`   |  `true`  | `organization that check repository belongs to` |
|      `repo`      |     `.release.repo`     | `string` |   `""`   |  `true`  | `repository of the check`                       |
|      `tag`       |     `.release.tag`      | `string` | `latest` | `false`  | `tagged version of check`                       |
|    `username`    | `.credentials.username` | `string` |   `""`   | `false`  | `username to provided to check`                 |
|    `password`    | `.credentials.password` | `string` |   `""`   | `false`  | `default password provided to the check`        |
|     `target`     |        `.target`        | `string` |   `""`   |  `true`  | `network target provide to the check`           |
|    `command`     |       `.command`        | `string` |   `""`   | `false`  | `command provided to the check`                 |
| `expectedOutput` |    `.expectedOutput`    | `string` |   `""`   | `false`  | `expected output for check to measured against` |
|     `weight`     |        `.weight`        |  `int`   |   `0`    |  `true`  | `amount of points a successful check is worth`  |
|    `options`     |       `.options`        |  `map`   |   `{}`   | `false`  | `maps for custom fields for checks`             |
