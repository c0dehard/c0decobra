Simple and easy example for **[spf13/cobra](https://github.com/spf13/cobra/)**</br>

# Initialize
``` 
cobra init --pkg-name <NameOfYourApp>
```
See also why it isn't just `cobra init` anymore: https://github.com/spf13/cobra/pull/888#issuecomment-503311013 

## Want to add more commands?
### No problem!

> if [cobra](https://github.com/spf13/cobra/) == installed **{** 
```
cobra add <commandname>
```
**}** else **{**
> install([cobra](https://github.com/spf13/cobra/))
```
go get -u github.com/spf13/cobra/cobra
```
```golang
// Don't forget to init like mentioned
```
**}**
