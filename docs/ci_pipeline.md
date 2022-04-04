
# CI pipeline for LUKSO-cli
## File structure
```
.github  
   |  
   ------ workflows  
   |      |  
   |      ---- release.yml # file for creating release  
   |  
   ------- RELEASE-TEMPLATE.md # file to create any message while we are releasing any tag
   ```

## Supported OS
1. linux
2. darwin

## How CI is working
1. on create any new tag a new github action triggers.
2. github checkouts to our current branch where our new code resides.
3. go version `1.17.5` is downloaded
4. a name for binary is prepared. Right now the naming structure is: {github-repository-name}-{OS}-{Architecture}
5. use `go` to build binary with the above defined name
6. create a release file for this release tag using `RELEASE-TEMPLATE.md` template
7. upload newly generated executables into the asset section of the same tag.