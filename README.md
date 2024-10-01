# Usage of Kompack Go API

## 1. Run HTTP Server Command

The rest command to handle RESTful operations. Below are the descriptions and usages of each flag.

#### Example to run rest

```bash
 $ main rest
```

## 2. Migration Command

This command provides several flags to handle database migrations. Below are the descriptions and usages of each flag.

### Flags

- **`--mmf`**: This flag is used to create a migration file name.

  - **Usage**: `--mmf=<migration_file_name>`
  - **Example**: `--mmf=initial_migration`

- **`--file`, `-f`**: This flag is used to specify a file for migration.

  - **Usage**: `--file=<file_name>` or `-f <file_name>`
  - **Example**: `--file=migration.sql` or `-f migration.sql`

- **`--init-migration`, `-i`**: This flag is used for the first migration.
  - **Usage**: `--init-migration` or `-i`
  - **Example**: `--init-migration` or `-i`

### Exclusive Flags

The following flags are mutually exclusive, meaning you cannot use them together in a single command:

- `--mmf`
- `--file` (or `-f`)
- `--init-migration` (or `-i`)

#### Example to run migrate

```bash
 $ main migrate [flags]
```
