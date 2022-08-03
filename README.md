## Simple installer helper for GO

**This tool should help Go developers with installing dependencies to their projects**

With installer you can create aliases for multiple package paths at once.
Make sure you have sqlite3 installed, you can test it by running `$ sqlite3` in terminal

**Example:**

```
<!-- Creating alias `fiber` for 3 packages -->
$ ~/installer add fiber github.com/gofiber/fiber gorm.io/gorm gorm.io/driver/sqlite

<!-- In the project folder -->
$ ~/installer install fiber

<!-- github.com/gofiber/fiber gets installed to the project -->
<!-- gorm.io/gorm gets installed to the project -->
<!-- gorm.io/driver/sqlite gets installed to the project -->
```

It's quite handy.

**Usage:**

1. Download or clone the project
2. Build the project to create executable
3. Move the executable to your home directory (or anywhere else, it's about ease of running the executable. You can perhaps create zsh or bash alias so you can run it from anywhere)

**To save an alias run:**
Separate packages paths with space

```
$ ~/installer add <ALIAS> <PATH-TO-DEPENDENCY1> <PATH-TO-DEPENDENCY2> etc...
```

Example:

```
$ ~/installer add gorm gorm.io/gorm gorm.io/driver/sqlite
```

This commands will create alias gorm for gorm and SQLite driver.

**List your current aliases**

```
$ ~/installer list
```

**Install dependencies to Go project by alias**

```
$ ~/installer install gorm
```

**Delete alias**
All of the packages paths will be deleted

```
$ ~/installer delete gorm
```
