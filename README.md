# Bench-Calculations

Script to produce Excel formulas to calculate bench levels for the next 4 weeks

## Use

Available flags

-column = start column (string) for calculation (default=U)
-row = start row (int) for calculation (default=3)
-weeks = number of weeks to derive formula for (default=39)
-equations = the number of weeks to create equations for (default=7)

### Linux
```
./bin/linux/app -column=T -row=3 -weeks=52 -equations=7
```

or

```
make run
```

This will run with defaults of column "T", row 3, 52 weeks and for 7 weeks of bench equations

### Windows
```
.\bin\win\app.exe
```