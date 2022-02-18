# budget-bot
A cli that automates keeping a financial budget using the Plaid API &amp; Google Sheets API

## Install
1. Build the latest version of the budget-bot binary (macOS Mach-O 64-bit executable x86_64 file)
```bash
git pull origin main
make build
```
2. Copy the executable to `/usr/local/bin`. For example:
```bash
cp budget-bot /usr/local/bin
```
3. If running `budget-bot` results in a `permission denied` error, 
make it executable by running
```bash
chmod +x /usr/local/bin/budget-bot
```
