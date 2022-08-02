## Simple installer helper for GO

**This tool should help Go developers with installing dependencies to their projects**

Let's say you want to install a dependency with a long name. With this tool, you can create alias for the dependency path and then install it to the project through the alias. Installer is using SQLite DB behind the scenes.

**Usage:**

1. Download or clone the project
2. Build the project to create executable
3. Move the executable to your home directory (or anywhere else, it's about ease of running the executable. You can perhaps create alias so you can run it from anywhere)
4. To save an alias run:

```
$ ~/installer add <ALIAS> <PATH-TO-DEPENDENCY>
```

Example:

```
$ ~/installer add gorm gorm.io/gorm
$ ~/installer add gormsqlite gorm.io/driver/sqlite
```

This commands will create aliases for gorm and SQLite driver.

5. List your current aliases

```
$ ~/installer list
```

6. Install dependency to Go project with installer

```
$ ~/installer install gorm
$ ~/installer install gormsqlite
```

7. Remove aliases

```
$ ~/installer remove gorm
```
