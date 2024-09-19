package main

import "os"

func commandExit(cfg *urlCfg) error {
	os.Exit(0)
	return nil
}
