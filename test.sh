// Source - https://stackoverflow.com/a/27245610
// Posted by Makpoc
// Retrieved 2026-02-24, License - CC BY-SA 3.0

go test -v ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
