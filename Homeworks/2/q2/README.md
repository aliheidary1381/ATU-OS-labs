## AWK and jq scripting
Python libraries have their own default package manager called `pip`,
but `pip` does not support upgrading all packages with one command.

You are assigned to write one-liner scripts that upgrade all outdated python libraries.
### Note
Running `pip list --outdated`, gives us something like the below output:
```text
Package            Version     Latest       Type
------------------ ----------- ------------ -----
numpy              1.25.2      1.26.0       wheel
pandas             2.1.0       2.1.1        wheel
```
Running `pip list --outdated --format=json`, gives us something like the below output:
```json
[{"name": "numpy", "version": "1.25.2", "latest_version": "1.26.0", "latest_filetype": "wheel"}, {"name": "pandas", "version": "2.1.0", "latest_version": "2.1.1", "latest_filetype": "wheel"}]
```
